package field_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	f := field.Int("age").Positive()
	assert.Equal(t, "age", f.Name())
	assert.Equal(t, field.TypeInt, f.Type())
	assert.Len(t, f.Validators(), 1)

	f = field.Int("age").Default(10).Min(10).Max(20)
	assert.True(t, f.HasDefault())
	assert.Equal(t, 10, f.Value())
	assert.Len(t, f.Validators(), 2)

	f = field.Int("age").Range(20, 40).Nillable()
	assert.False(t, f.HasDefault())
	assert.True(t, f.IsNillable())
	assert.Len(t, f.Validators(), 1)

	assert.Equal(t, field.TypeInt8, field.Int8("age").Type())
	assert.Equal(t, field.TypeInt16, field.Int16("age").Type())
	assert.Equal(t, field.TypeInt32, field.Int32("age").Type())
	assert.Equal(t, field.TypeInt64, field.Int64("age").Type())
}

func TestFloat(t *testing.T) {
	f := field.Float("age").Positive()
	assert.Equal(t, "age", f.Name())
	assert.Equal(t, field.TypeFloat64, f.Type())
	assert.Len(t, f.Validators(), 1)

	f = field.Float("age").Min(2.5).Max(5)
	assert.Len(t, f.Validators(), 2)
}

func TestBool(t *testing.T) {
	f := field.Bool("active").Default(true)
	assert.Equal(t, "active", f.Name())
	assert.Equal(t, field.TypeBool, f.Type())
	assert.True(t, f.HasDefault())
	assert.Equal(t, true, f.Value())
}

func TestBytes(t *testing.T) {
	f := field.Bytes("active").Default([]byte("{}"))
	assert.Equal(t, "active", f.Name())
	assert.Equal(t, field.TypeBytes, f.Type())
	assert.True(t, f.HasDefault())
	assert.Equal(t, []byte("{}"), f.Value())
}

func TestString(t *testing.T) {
	re := regexp.MustCompile("[a-zA-Z0-9]")
	f := field.String("name").Unique().Match(re).Validate(func(string) error { return nil })
	assert.Equal(t, field.TypeString, f.Type())
	assert.Equal(t, "name", f.Name())
	assert.True(t, f.IsUnique())
	assert.Len(t, f.Validators(), 2)
}

func TestCharset(t *testing.T) {
	var f ent.Field = field.String("name").SetCharset("utf8")
	cs, ok := f.(field.Charseter)
	assert.True(t, ok, "string field implements the Charseter interface")
	assert.Equal(t, "utf8", cs.Charset())
}

func TestTime(t *testing.T) {
	now := time.Now()
	f := field.Time("created_at").
		Default(func() time.Time {
			return now
		})
	assert.Equal(t, "created_at", f.Name())
	assert.Equal(t, field.TypeTime, f.Type())
	assert.Equal(t, "time.Time", f.Type().String())
	assert.NotNil(t, f.Value())
	assert.Equal(t, now, f.Value().(func() time.Time)())
}

func TestField_Tag(t *testing.T) {
	f := field.Bool("expired").StructTag(`json:"expired,omitempty"`)
	assert.Equal(t, `json:"expired,omitempty"`, f.Tag())
}
