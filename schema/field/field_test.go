package field_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/facebookincubator/ent/schema/field"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	f := field.Int("age").Positive()
	fd := f.Descriptor()
	assert.Equal(t, "age", fd.Name)
	assert.Equal(t, field.TypeInt, fd.Type)
	assert.Len(t, fd.Validators, 1)

	f = field.Int("age").Default(10).Min(10).Max(20)
	fd = f.Descriptor()
	assert.NotNil(t, fd.Default)
	assert.Equal(t, 10, fd.Default)
	assert.Len(t, fd.Validators, 2)

	f = field.Int("age").Range(20, 40).Nillable()
	fd = f.Descriptor()
	assert.Nil(t, fd.Default)
	assert.True(t, fd.Nillable)
	assert.False(t, fd.Immutable)
	assert.Len(t, fd.Validators, 1)

	assert.Equal(t, field.TypeInt8, field.Int8("age").Descriptor().Type)
	assert.Equal(t, field.TypeInt16, field.Int16("age").Descriptor().Type)
	assert.Equal(t, field.TypeInt32, field.Int32("age").Descriptor().Type)
	assert.Equal(t, field.TypeInt64, field.Int64("age").Descriptor().Type)
}

func TestFloat(t *testing.T) {
	f := field.Float("age").Positive()
	fd := f.Descriptor()
	assert.Equal(t, "age", fd.Name)
	assert.Equal(t, field.TypeFloat64, fd.Type)
	assert.Len(t, fd.Validators, 1)

	f = field.Float("age").Min(2.5).Max(5)
	fd = f.Descriptor()
	assert.Len(t, fd.Validators, 2)
}

func TestBool(t *testing.T) {
	f := field.Bool("active").Default(true).Immutable()
	fd := f.Descriptor()
	assert.Equal(t, "active", fd.Name)
	assert.Equal(t, field.TypeBool, fd.Type)
	assert.NotNil(t, fd.Default)
	assert.True(t, fd.Immutable)
	assert.Equal(t, true, fd.Default)
}

func TestBytes(t *testing.T) {
	f := field.Bytes("active").Default([]byte("{}"))
	fd := f.Descriptor()
	assert.Equal(t, "active", fd.Name)
	assert.Equal(t, field.TypeBytes, fd.Type)
	assert.NotNil(t, fd.Default)
	assert.Equal(t, []byte("{}"), fd.Default)
}

func TestString(t *testing.T) {
	re := regexp.MustCompile("[a-zA-Z0-9]")
	f := field.String("name").Unique().Match(re).Validate(func(string) error { return nil })
	fd := f.Descriptor()
	assert.Equal(t, field.TypeString, fd.Type)
	assert.Equal(t, "name", fd.Name)
	assert.True(t, fd.Unique)
	assert.Len(t, fd.Validators, 2)
}

func TestCharset(t *testing.T) {
	fd := field.String("name").
		Charset("utf8").
		Descriptor()
	assert.Equal(t, "utf8", fd.Charset)
}

func TestTime(t *testing.T) {
	now := time.Now()
	fd := field.Time("created_at").
		Default(func() time.Time {
			return now
		}).
		Descriptor()
	assert.Equal(t, "created_at", fd.Name)
	assert.Equal(t, field.TypeTime, fd.Type)
	assert.Equal(t, "time.Time", fd.Type.String())
	assert.NotNil(t, fd.Default)
	assert.Equal(t, now, fd.Default.(func() time.Time)())
}

func TestField_Tag(t *testing.T) {
	fd := field.Bool("expired").
		StructTag(`json:"expired,omitempty"`).
		Descriptor()
	assert.Equal(t, `json:"expired,omitempty"`, fd.Tag)
}
