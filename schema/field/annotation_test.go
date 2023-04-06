// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package field_test

import (
	"testing"

	"entgo.io/ent/schema/field"

	"github.com/stretchr/testify/assert"
)

func TestAnnotation_Merge(t *testing.T) {
	ant := field.Annotation{}
	a := ant.Merge(field.Annotation{
		StructTag: map[string]string{"foo": "bar"},
	})
	assert.Equal(t, a.(field.Annotation).StructTag["foo"], "bar")
	a = ant.Merge(&field.Annotation{
		StructTag: map[string]string{"foo": "baz", "baz": "qux"},
	})
	assert.Equal(t, a.(field.Annotation).StructTag["foo"], "baz")
	assert.Equal(t, a.(field.Annotation).StructTag["baz"], "qux")
}
