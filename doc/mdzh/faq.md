---
id: faq
title: Frequently Asked Questions (FAQ)
sidebar_label: FAQ
---

## 常见问题

[如何从结构体`T`中创建一个实体](#how-to-create-an-entity-from-a-struct-t) \
[如何创建一个结构(或多样)级别的验证器?](#how-to-create-a-mutation-level-validator) \
[如何创建一个可扩展的日志系统?](#how-to-write-an-audit-log-extension) \
[如何写一个自定义的断言?](#how-to-write-custom-predicates) \
[如何给代码生成器写一个自定义的断言?](#how-to-add-custom-predicates-to-the-codegen-assets) \
[如何在PostgreSQL中定义一个网络地址字段?](#how-to-define-a-network-address-field-in-postgresql) \
[如何给MySQL中的`DATETIME`字段定义一个时间字段?](#how-to-customize-time-fields-to-type-datetime-in-mysql) \
[如何使用自定义的ID生成器?](#how-to-use-a-custom-generator-of-ids)

## 回答

#### 如何从结构体`T`中创建一个实体?

不同的构造器不支持从给定的结构体`T`中设置实体字段的选项。
原因是因为它在更新数据库的时候没有办法区分0值和实际的值(例如：`&ent.T{Age: 0, Name: ""}`)
设置这些值，可能在数据中设置错误的值或者更新不必要的列。

然而，这个[外部模板](templates.md)选项让你通过自定义逻辑继承默认的代码生成断言，
例如，为了生成每个创建构造器的方法，它接受一个结构体作为输入，并且配置构造器。
使用如下模板：

```gotemplate
{{ range $n := $.Nodes }}
    {{ $builder := $n.CreateName }}
    {{ $receiver := receiver $builder }}

    func ({{ $receiver }} *{{ $builder }}) Set{{ $n.Name }}(input *{{ $n.Name }}) *{{ $builder }} {
        {{- range $f := $n.Fields }}
            {{- $setter := print "Set" $f.StructField }}
            {{ $receiver }}.{{ $setter }}(input.{{ $f.StructField }})
        {{- end }}
        return {{ $receiver }}
    }
{{ end }}
```

#### 如何创建一个多级验证器?

为了实现多级验证器，你能够同时使用[schema hooks](hooks.md#schema-hooks)进行验证应用于一个实体类型的更改，也可以使用[transaction hooks](transactions.md#hooks)
来验证存在的变化,应用于多种实体类型(例如，GraphQL 变化)。例如

```go
// A VersionHook is a dummy example for a hook that validates the "version" field
// is incremented by 1 on each update. Note that this is just a dummy example, and
// it doesn't promise consistency in the database.
func VersionHook() ent.Hook {
	type OldSetVersion interface {
		SetVersion(int)
		Version() (int, bool)
		OldVersion(context.Context) (int, error)
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			ver, ok := m.(OldSetVersion)
			if !ok {
				return next.Mutate(ctx, m)
			}
			oldV, err := ver.OldVersion(ctx)
			if err != nil {
				return nil, err
			}
			curV, exists := ver.Version()
			if !exists {
				return nil, fmt.Errorf("version field is required in update mutation")
			}
			if curV != oldV+1 {
				return nil, fmt.Errorf("version field must be incremented by 1")
			}
			// Add an SQL predicate that validates the "version" column is equal
			// to "oldV" (ensure it wasn't changed during the mutation by others).
			return next.Mutate(ctx, m)
		})
	}
}
```

#### 如何写一个可扩展的日志系统?

编写这样一个扩展的首选方法是使用[ent.Mixin]。使用`Fields`选项设置在导入混合模式的所有模式之间的共享字段。使用`Hooks`选项
为应用于这些模式的所有变化附加一个mutation-hook,这里有个例子，基于[repository issue-tracker](https://entgo.io/ent/issues/830) 的讨论：

```go
// AuditMixin implements the ent.Mixin for sharing
// audit-log capabilities with package schemas.
type AuditMixin struct{
	mixin.Schema
}

// Fields of the AuditMixin.
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Int("created_by").
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("updated_by").
			Optional(),
	}
}

// Hooks of the AuditMixin.
func (AuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.AuditHook,
	}
}

// A AuditHook is an example for audit-log hook.
func AuditHook(next ent.Mutator) ent.Mutator {
	// AuditLogger wraps the methods that are shared between all mutations of
	// schemas that embed the AuditLog mixin. The variable "exists" is true, if
	// the field already exists in the mutation (e.g. was set by a different hook).
	type AuditLogger interface {
		SetCreatedAt(time.Time)
		CreatedAt() (value time.Time, exists bool)
		SetCreatedBy(int)
		CreatedBy() (id int, exists bool)
		SetUpdatedAt(time.Time)
		UpdatedAt() (value time.Time, exists bool)
		SetUpdatedBy(int)
		UpdatedBy() (id int, exists bool)
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}
		usr, err := viewer.UserFromContext(ctx)
		if err != nil {
			return nil, err
		}
		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetCreatedAt(time.Now())
			if _, exists := ml.CreatedBy(); !exists {
				ml.SetCreatedBy(usr.ID)
			}
		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			ml.SetUpdatedAt(time.Now())
			if _, exists := ml.UpdatedBy(); !exists {
				ml.SetUpdatedBy(usr.ID)
			}
		}
		return next.Mutate(ctx, m)
	})
}
```

#### 如何写一个自定义断言?

在执行查询之前，用户可以提供自定义的断言应用于查询。例如：

```go
pets := client.Pet.
	Query().
	Where(predicate.Pet(func(s *sql.Selector) {
		s.Where(sql.InInts(pet.OwnerColumn, 1, 2, 3))
	})).
	AllX(ctx)

users := client.User.
	Query().
	Where(predicate.User(func(s *sql.Selector) {
		s.Where(sqljson.ValueContains(user.FieldTags, "tag"))
	})).
	AllX(ctx)
```

获取更多的案例，到[predicates](predicates.md#custom-predicates) ，或者在仓库搜索更多高级的案例，如[issue-842](https://entgo.io/ent/issues/842#issuecomment-707896368).

#### 如何给代码生成器写一个自定义的断言?

[template](templates.md)选项能够扩展或者覆盖默认的代码生成器断言。
为了给[example above](#how-to-write-custom-predicates)生成类型安全的断言。使用模板选项做如下操作：

```gotemplate
{{/* A template that adds the "<F>Glob" predicate for all string fields. */}}
{{ define "where/additional/strings" }}
    {{ range $f := $.Fields }}
        {{ if $f.IsString }}
            {{ $func := print $f.StructField "Glob" }}
            // {{ $func }} applies the Glob predicate on the {{ quote $f.Name }} field.
            func {{ $func }}(pattern string) predicate.{{ $.Name }} {
                return predicate.{{ $.Name }}(func(s *sql.Selector) {
                    s.Where(sql.P(func(b *sql.Builder) {
                        b.Ident(s.C({{ $f.Constant }})).WriteString(" glob" ).Arg(pattern)
                    }))
                })
            }
        {{ end }}
    {{ end }}
{{ end }}
```

#### 如何在PostgreSQL中定义一个网络地址字段?

[GoType](http://localhost:3000/docs/schema-fields#go-type) 和 [SchemaType](http://localhost:3000/docs/schema-fields#database-type)
选项允许用户定义特殊的数据库字段，例如，为了定义一个 [`macaddr`](https://www.postgresql.org/docs/13/datatype-net-types.html#DATATYPE-MACADDR)字段
使用如下配置：

```go
func (T) Fields() []ent.Field {
	return []ent.Field{
		field.String("mac").
			GoType(&MAC{}).
			SchemaType(map[string]string{
				dialect.Postgres: "macaddr",
			}).
			Validate(func(s string) error {
				_, err := net.ParseMAC(s)
				return err
			}),
	}
}

// MAC represents a physical hardware address.
type MAC struct {
	net.HardwareAddr
}

// Scan implements the Scanner interface.
func (m *MAC) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		m.HardwareAddr, err = net.ParseMAC(string(v))
	case string:
		m.HardwareAddr, err = net.ParseMAC(v)
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements the driver Valuer interface.
func (m MAC) Value() (driver.Value, error) {
	return m.HardwareAddr.String(), nil
}
```

注意，如果数据库不支持`macaddr`类型（例如SQLite的测试）,这个字段回调它自己的本地类型（也就是`string`）

`inet` 案例:

```go
func (T) Fields() []ent.Field {
    return []ent.Field{
		field.String("ip").
			GoType(&Inet{}).
			SchemaType(map[string]string{
				dialect.Postgres: "inet",
			}).
			Validate(func(s string) error {
				if net.ParseIP(s) == nil {
					return fmt.Errorf("invalid value for ip %q", s)
				}
				return nil
			}),
    }
}

// Inet represents a single IP address
type Inet struct {
    net.IP
}

// Scan implements the Scanner interface
func (i *Inet) Scan(value interface{}) (err error) {
    switch v := value.(type) {
    case nil:
    case []byte:
        if i.IP = net.ParseIP(string(v)); i.IP == nil {
            err = fmt.Errorf("invalid value for ip %q", s)
        }
    case string:
        if i.IP = net.ParseIP(v); i.IP == nil {
            err = fmt.Errorf("invalid value for ip %q", s)
        }
    default:
        err = fmt.Errorf("unexpected type %T", v)
    }
    return
}

// Value implements the driver Valuer interface
func (i Inet) Value() (driver.Value, error) {
    return i.IP.String(), nil
}
```

#### 如何给MySQL中的`DATETIME`字段定义一个时间字段?

`Time`字段默认使用MySQL中的`TIMESTAMP` 类型创建。此类型的范围是'1970-01-01 00:00:01' UTC 到 '2038-01-19 03:14:07' UTC ，
（详情请看[MySQL docs](https://dev.mysql.com/doc/refman/5.6/en/datetime.html)）

为了自定义更长的时间字段，使用MySQL的`DATETIME`。
```go
field.Time("birth_date").
	Optional().
	SchemaType(map[string]string{
		dialect.MySQL: "datetime",
	}),
```

#### 如何使用自定义的ID生成器?

如果你使用定制的ID生成器代替数据库中的自增（例如Twitter公司的[Snowflake](https://github.com/twitter-archive/snowflake/tree/snowflake-2010)）
你将需要编写一个自定义ID字段，该字段在创建资源时自动调用生成器。

要实现这一点，你可以使用`DefaultFunc`或schema hooks这取决于你的用例。
如果生成器不返回错误。`DefaultFunc`会更加简洁。而在资源创建上设置一个hook也将允许你捕获错误。
链接中展示了如何用`DefaultFunc`实现([the ID field](schema-fields.md#id-field))。

这里的例子展示了如何使用HooK来写自定义的生成器，举个例子：[sonyflake](https://github.com/sony/sonyflake).

```go
// BaseMixin to be shared will all different schemas.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the Mixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
	}
}

// Hooks of the Mixin.
func (BaseMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(IDHook(), ent.OpCreate),
	}
}

func IDHook() ent.Hook {
    sf := sonyflake.NewSonyflake(sonyflage.Settings{})
	type IDSetter interface {
		SetID(uint64)
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			is, ok := m.(IDSetter)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation %T", m)
			}
			id, err := sf.NextID()
			if err != nil {
				return nil, err
			}
			is.SetID(id)
			return next.Mutate(ctx, m)
		})
	}
}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}
```

