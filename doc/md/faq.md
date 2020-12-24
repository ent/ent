---
id: faq
title: Frequently Asked Questions (FAQ)
sidebar_label: FAQ
---

## Questions

[How to create an entity from a struct `T`?](#how-to-create-an-entity-from-a-struct-t)  
[How to create a struct (or a mutation) level validator?](#how-to-create-a-mutation-level-validator)  
[How to write an audit-log extension?](#how-to-write-an-audit-log-extension)  
[How to write custom predicates?](#how-to-write-custom-predicates)  
[How to add custom predicates to the codegen assets?](#how-to-add-custom-predicates-to-the-codegen-assets)  
[How to define a network address field in PostgreSQL?](#how-to-define-a-network-address-field-in-postgresql)  
[How to customize time fields to type `DATETIME` in MySQL?](#how-to-customize-time-fields-to-type-datetime-in-mysql)

## Answers

#### How to create an entity from a struct `T`?

The different builders don't support the option of setting the entity fields (or edges) from a given struct `T`.
The reason is that there's no way to distinguish between zero/real values when updating the database (for example, `&ent.T{Age: 0, Name: ""}`).
Setting these values, may set incorrect values in the database or update unnecessary columns.

However, the [external template](templates.md) option lets you extend the default code-generation assets by adding custom logic.
For example, in order to generate a method for each of the create-builders, that accepts a struct as an input and configure the builder,
use the following template:

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

#### How to create a mutation level validator?

In order to implement a mutation-level validator, you can either use [schema hooks](hooks.md#schema-hooks) for validating
changes applied on one entity type, or use [transaction hooks](transactions.md#hooks) for validating mutations that being
applied on multiple entity types (e.g. a GraphQL mutation). For example:

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

#### How to write an audit-log extension?

The preferred way for writing such an extension is to use [ent.Mixin](schema-mixin.md). Use the `Fields` option for
setting the fields that are shared between all schemas that import the mixed-schema, and use the `Hooks` option for
attaching a mutation-hook for all mutations that are being applied on these schemas. Here's an example, based on a
discussion in the [repository issue-tracker](https://github.com/facebook/ent/issues/830):

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

#### How to write custom predicates?

Users can provide custom predicates to apply on the query before it's executed. For example:

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

For more examples, go to the [predicates](predicates.md#custom-predicates) page, or search in the repository
issue-tracker for more advance examples like [issue-842](https://github.com/facebook/ent/issues/842#issuecomment-707896368).

#### How to add custom predicates to the codegen assets?

The [template](templates.md) option enables the capability for extending or overriding the default codegen assets.
In order to generate a type-safe predicate for the [example above](#how-to-write-custom-predicates), use the template
option for doing it as follows:

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

#### How to define a network address field in PostgreSQL?

The [GoType](http://localhost:3000/docs/schema-fields#go-type) and the [SchemaType](http://localhost:3000/docs/schema-fields#database-type)
options allow users to define database-specific fields. For example, in order to define a
 [`macaddr`](https://www.postgresql.org/docs/13/datatype-net-types.html#DATATYPE-MACADDR) field, use the following configuration:

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
Note that, if the database doesn't support the `macaddr` type (e.g. SQLite on testing), the field fallback to its
native type (i.e. `string`).

`inet` example:

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

#### How to customize time fields to type `DATETIME` in MySQL?

`Time` fields use the MySQL `TIMESTAMP` type in the schema creation by default, and this type
 has a range of '1970-01-01 00:00:01' UTC to '2038-01-19 03:14:07' UTC (see, [MySQL docs](https://dev.mysql.com/doc/refman/5.6/en/datetime.html)). 

In order to customize time fields for a wider range, use the MySQL `DATETIME` as follows:
```go
field.Time("birth_date").
	Optional().
	SchemaType(map[string]string{
		dialect.MySQL: "datetime",
	}),
```