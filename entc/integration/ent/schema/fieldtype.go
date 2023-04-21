// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"bytes"
	"database/sql/driver"
	"encoding/gob"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/role"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// FieldType holds the schema definition for the FieldType entity.
// used for testing field types.
type FieldType struct {
	ent.Schema
}

// Fields of the File.
func (FieldType) Fields() []ent.Field { //nolint:funlen
	return []ent.Field{
		// ----------------------------------------------------------------------------
		// Basic types

		field.Int("int"),
		field.Int8("int8"),
		field.Int16("int16"),
		field.Int32("int32"),
		field.Int64("int64").
			UpdateDefault(func() int64 {
				return 100
			}),
		field.Int("optional_int").
			Optional(),
		field.Int8("optional_int8").
			Optional(),
		field.Int16("optional_int16").
			Optional(),
		field.Int32("optional_int32").
			Optional(),
		field.Int64("optional_int64").
			Optional(),
		field.Int("nillable_int").
			Optional().
			Nillable(),
		field.Int8("nillable_int8").
			Optional().
			Nillable(),
		field.Int16("nillable_int16").
			Optional().
			Nillable(),
		field.Int32("nillable_int32").
			Optional().
			Nillable(),
		field.Int64("nillable_int64").
			Optional().
			Nillable(),
		field.Int32("validate_optional_int32").
			Optional().
			Max(100),
		field.Uint("optional_uint").
			Optional(),
		field.Uint8("optional_uint8").
			Optional(),
		field.Uint16("optional_uint16").
			Optional(),
		field.Uint32("optional_uint32").
			Optional(),
		field.Uint64("optional_uint64").
			Optional(),
		field.Enum("state").
			Values("on", "off").
			Optional(),
		field.Float("optional_float").
			Optional(),
		field.Float32("optional_float32").
			Optional(),

		// ----------------------------------------------------------------------------
		// Dialect-specific types

		field.Text("text").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "mediumtext",
			}),
		field.Time("datetime").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "datetime",
				dialect.Postgres: "date",
			}),
		field.Float("decimal").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)",
				dialect.Postgres: "numeric",
			}),
		field.Other("link_other", &Link{}).
			SchemaType(map[string]string{
				dialect.Postgres: "varchar",
				dialect.MySQL:    "varchar(255)",
				dialect.SQLite:   "varchar(255)",
			}).
			Optional().
			Default(DefaultLink()),
		field.Other("link_other_func", &Link{}).
			SchemaType(map[string]string{
				dialect.Postgres: "varchar",
				dialect.MySQL:    "varchar(255)",
				dialect.SQLite:   "varchar(255)",
			}).
			Optional().
			Default(DefaultLink),
		field.String("mac").
			Optional().
			GoType(MAC{}).
			SchemaType(map[string]string{
				dialect.Postgres: "macaddr",
			}).
			Validate(func(s string) error {
				_, err := net.ParseMAC(s)
				return err
			}),
		field.Other("string_array", Strings{}).
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "text[]",
				dialect.SQLite:   "json",
				dialect.MySQL:    "blob",
			}),
		field.String("password").
			Optional().
			Sensitive().
			SchemaType(map[string]string{
				dialect.MySQL: "char(32)",
			}),

		// ----------------------------------------------------------------------------
		// Custom Go types

		field.String("string_scanner").
			GoType(StringScanner("")).
			Nillable().
			Optional(),
		field.Int64("duration").
			GoType(time.Duration(0)).
			UpdateDefault(func() time.Duration {
				return time.Duration(100)
			}).
			Optional(),
		field.String("dir").
			GoType(http.Dir("dir")).
			DefaultFunc(func() http.Dir {
				return "unknown"
			}),
		field.String("ndir").
			Optional().
			Nillable().
			NotEmpty().
			GoType(http.Dir("ndir")),
		field.String("str").
			Optional().
			GoType(sql.NullString{}).
			DefaultFunc(func() sql.NullString {
				return sql.NullString{String: "default", Valid: true}
			}),
		field.String("null_str").
			Optional().
			Nillable().
			GoType(&sql.NullString{}).
			DefaultFunc(func() *sql.NullString {
				return &sql.NullString{String: "default", Valid: true}
			}),
		field.String("link").
			Optional().
			NotEmpty().
			GoType(Link{}),
		field.String("null_link").
			Optional().
			Nillable().
			GoType(&Link{}),
		field.Bool("active").
			Optional().
			GoType(Status(false)),
		field.Bool("null_active").
			Optional().
			Nillable().
			GoType(Status(false)),
		field.Bool("deleted").
			Optional().
			Nillable().
			GoType(&sql.NullBool{}),
		field.Time("deleted_at").
			Optional().
			GoType(&sql.NullTime{}).
			Default(func() *sql.NullTime {
				return &sql.NullTime{Time: time.Now(), Valid: true}
			}).
			UpdateDefault(func() *sql.NullTime {
				return &sql.NullTime{Time: time.Now(), Valid: true}
			}),
		field.Bytes("raw_data").
			Optional().
			MaxLen(20).
			MinLen(3),
		field.Bytes("sensitive").
			Optional().
			Sensitive(),
		field.Bytes("ip").
			Optional().
			GoType(net.IP("127.0.0.1")).
			DefaultFunc(func() net.IP {
				return net.IP("127.0.0.1")
			}).
			Validate(func(i []byte) error {
				if net.ParseIP(string(i)) == nil {
					return fmt.Errorf("ent/schema: invalid ip %q", string(i))
				}
				return nil
			}),
		field.Int("null_int64").
			Optional().
			GoType(&sql.NullInt64{}),
		field.Int("schema_int").
			Optional().
			GoType(Int(0)),
		field.Int8("schema_int8").
			Optional().
			GoType(Int8(0)),
		field.Int64("schema_int64").
			Optional().
			GoType(Int64(0)),
		field.Float("schema_float").
			Optional().
			GoType(Float64(0)),
		field.Float32("schema_float32").
			Optional().
			GoType(Float32(0)),
		field.Float("null_float").
			Optional().
			GoType(&sql.NullFloat64{}),
		field.Enum("role").
			Default(string(role.Read)).
			GoType(role.Role("role")),
		field.Enum("priority").
			Optional().
			GoType(role.Priority(0)),
		field.UUID("optional_uuid", uuid.UUID{}).
			Optional(),
		field.UUID("nillable_uuid", uuid.UUID{}).
			Optional().
			Nillable(),
		field.Strings("strings").
			Optional(),
		field.Bytes("pair").
			GoType(Pair{}).
			DefaultFunc(func() Pair {
				return Pair{K: []byte("K"), V: []byte("V")}
			}),
		field.Bytes("nil_pair").
			GoType(&Pair{}).
			Optional().
			Nillable(),
		field.String("vstring").
			GoType(VString("")).
			DefaultFunc(func() VString {
				return "value scanner string"
			}),
		field.String("triple").
			GoType(Triple{}).
			DefaultFunc(func() Triple {
				return Triple{E: [3]string{"A", "B", "C"}}
			}),
		field.Int("big_int").
			Optional().
			GoType(BigInt{}),
		field.Other("password_other", Password("")).
			Optional().
			Sensitive().
			SchemaType(map[string]string{
				dialect.MySQL:    "char(32)",
				dialect.SQLite:   "char(32)",
				dialect.Postgres: "varchar",
			}),
	}
}

type Password string

func (p Password) Value() (driver.Value, error) {
	return string(p), nil
}

func (p *Password) Scan(src any) error {
	switch src := src.(type) {
	case nil:
		return nil
	case string:
		*p = Password(src)
		return nil
	case []byte:
		*p = Password(src)
		return nil
	default:
		return fmt.Errorf("scan: unable to scan type %T into string", src)
	}
}

type Strings []string

func (s *Strings) Scan(v any) (err error) {
	switch v := v.(type) {
	case nil:
	case []byte:
		err = s.scan(string(v))
	case string:
		err = s.scan(v)
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

func (s *Strings) scan(v string) error {
	if v == "" {
		return nil
	}
	if l := len(v); l < 2 || v[0] != '{' && v[l-1] != '}' {
		return fmt.Errorf("unexpected array format %q", v)
	}
	*s = strings.Split(v[1:len(v)-1], ",")
	return nil
}

func (s Strings) Value() (driver.Value, error) {
	return "{" + strings.Join(s, ",") + "}", nil
}

type VString string

func (s *VString) Scan(v any) (err error) {
	switch v := v.(type) {
	case nil:
	case string:
		*s = VString(v)
	case []byte:
		*s = VString(v)
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

func (s VString) Value() (driver.Value, error) {
	return string(s), nil
}

type Triple struct {
	E [3]string
}

// Value implements the driver Valuer interface.
func (t Triple) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%s,%s)", t.E[0], t.E[1], t.E[2]), nil
}

// Scan implements the Scanner interface.
func (t *Triple) Scan(value any) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		es := strings.Split(strings.TrimPrefix(string(v), "()"), ",")
		t.E[0], t.E[1], t.E[2] = es[0], es[1], es[2]
	case string:
		es := strings.Split(strings.TrimPrefix(v, "()"), ",")
		t.E[0], t.E[1], t.E[2] = es[0], es[1], es[2]
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

type Pair struct {
	K, V []byte
}

// Value implements the driver Valuer interface.
func (p Pair) Value() (driver.Value, error) {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(p); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Scan implements the Scanner interface.
func (p *Pair) Scan(value any) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		err = gob.NewDecoder(bytes.NewBuffer(v)).Decode(p)
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

type (
	Int     int
	Int8    int8
	Int64   int64
	Status  bool
	Float64 float64
	Float32 float32
)

type Link struct {
	*url.URL
}

func DefaultLink() *Link {
	u, _ := url.Parse("127.0.0.1")
	return &Link{URL: u}
}

// Scan implements the Scanner interface.
func (l *Link) Scan(value any) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		l.URL, err = url.Parse(string(v))
	case string:
		l.URL, err = url.Parse(v)
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements the driver Valuer interface.
func (l Link) Value() (driver.Value, error) {
	if l.URL == nil {
		return nil, nil
	}
	return l.String(), nil
}

type MAC struct {
	net.HardwareAddr
}

// Scan implements the Scanner interface.
func (m *MAC) Scan(value any) (err error) {
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

type StringScanner string

// Scan implements the Scanner interface.
func (s *StringScanner) Scan(value any) (err error) {
	switch v := value.(type) {
	case nil:
	case string:
		*s = StringScanner(v)
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements the driver Valuer interface.
func (s StringScanner) Value() (driver.Value, error) {
	return string(s), nil
}

type BigInt struct {
	*big.Int
}

func NewBigInt(i int64) BigInt {
	return BigInt{Int: big.NewInt(i)}
}

func (b *BigInt) Scan(src any) error {
	var i sql.NullString
	if err := i.Scan(src); err != nil {
		return err
	}
	if !i.Valid {
		return nil
	}
	if b.Int == nil {
		b.Int = big.NewInt(0)
	}
	// Value came in a floating point format.
	if strings.ContainsAny(i.String, ".+e") {
		f := big.NewFloat(0)
		if _, err := fmt.Sscan(i.String, f); err != nil {
			return err
		}
		b.Int, _ = f.Int(b.Int)
	} else if _, err := fmt.Sscan(i.String, b.Int); err != nil {
		return err
	}
	return nil
}

func (b BigInt) Value() (driver.Value, error) {
	return b.String(), nil
}

func (b BigInt) Add(c BigInt) BigInt {
	b.Int = b.Int.Add(b.Int, c.Int)
	return b
}
