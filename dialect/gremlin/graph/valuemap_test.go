// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValueMapDecodeOne(t *testing.T) {
	vm := ValueMap{map[string]any{
		"id":    int64(1),
		"label": "person",
		"name":  []any{"marko"},
		"age":   []any{int32(29)},
	}}

	var ent struct {
		ID    uint64 `json:"id"`
		Label string `json:"label"`
		Name  string `json:"name"`
		Age   uint8  `json:"age"`
	}
	err := vm.Decode(&ent)
	require.NoError(t, err)

	assert.Equal(t, uint64(1), ent.ID)
	assert.Equal(t, "person", ent.Label)
	assert.Equal(t, "marko", ent.Name)
	assert.Equal(t, uint8(29), ent.Age)
}

func TestValueMapDecodeMany(t *testing.T) {
	vm := ValueMap{
		map[string]any{
			"id":    int64(1),
			"label": "person",
			"name":  []any{"chico"},
		},
		map[string]any{
			"id":    int64(2),
			"label": "person",
			"name":  []any{"dico"},
		},
	}

	ents := []struct {
		ID    int    `json:"id"`
		Label string `json:"label"`
		Name  string `json:"name"`
	}{}
	err := vm.Decode(&ents)
	require.NoError(t, err)

	require.Len(t, ents, 2)
	assert.Equal(t, 1, ents[0].ID)
	assert.Equal(t, "person", ents[0].Label)
	assert.Equal(t, "chico", ents[0].Name)
	assert.Equal(t, 2, ents[1].ID)
	assert.Equal(t, "person", ents[1].Label)
	assert.Equal(t, "dico", ents[1].Name)
}

func TestValueMapDecodeBadInput(t *testing.T) {
	type s struct{ Name string }
	err := ValueMap{}.Decode(s{})
	assert.Error(t, err)
	err = ValueMap{}.Decode((*s)(nil))
	assert.Error(t, err)
}
