package field_test

import (
	"testing"

	"github.com/facebook/ent/schema/field"

	"github.com/stretchr/testify/assert"
)

func TestAnnotation_Merge(t *testing.T) {
	ant := field.Annotation{}
	ant.Merge(field.Annotation{
		StructTag: map[string]string{"foo": "bar"},
	})
	assert.Equal(t, ant.StructTag["foo"], "bar")
	ant.Merge(&field.Annotation{
		StructTag: map[string]string{"foo": "baz", "baz": "qux"},
	})
	assert.Equal(t, ant.StructTag["foo"], "baz")
	assert.Equal(t, ant.StructTag["baz"], "qux")
}
