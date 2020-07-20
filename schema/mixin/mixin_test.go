// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package mixin_test

import (
	"testing"

	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/mixin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeMixin(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		t.Parallel()
		fields := mixin.CreateTime{}.Fields()
		require.Len(t, fields, 1)
		desc := fields[0].Descriptor()
		assert.Equal(t, "create_time", desc.Name)
		assert.True(t, desc.Immutable)
		assert.NotNil(t, desc.Default)
		assert.Nil(t, desc.UpdateDefault)
	})
	t.Run("Update", func(t *testing.T) {
		t.Parallel()
		fields := mixin.UpdateTime{}.Fields()
		require.Len(t, fields, 1)
		desc := fields[0].Descriptor()
		assert.Equal(t, "update_time", desc.Name)
		assert.True(t, desc.Immutable)
		assert.NotNil(t, desc.Default)
		assert.NotNil(t, desc.UpdateDefault)
	})
	t.Run("Compose", func(t *testing.T) {
		t.Parallel()
		fields := mixin.Time{}.Fields()
		require.Len(t, fields, 2)
		assert.Equal(t, "create_time", fields[0].Descriptor().Name)
		assert.Equal(t, "update_time", fields[1].Descriptor().Name)
	})
}

type annotation string

func (annotation) Name() string { return "" }

func TestAnnotate(t *testing.T) {
	annotations := []field.Annotation{
		annotation("foo"),
		annotation("bar"),
		annotation("baz"),
	}
	fields := mixin.Annotate(
		mixin.Time{}, annotations...,
	).Fields()
	for _, f := range fields {
		desc := f.Descriptor()
		require.Len(t, desc.Annotations, len(annotations))
		for i := range desc.Annotations {
			assert.Equal(t, annotations[i], desc.Annotations[i])
		}
	}
}
