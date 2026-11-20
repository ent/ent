// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package load

import (
	"database/sql/driver"
	"testing"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/stretchr/testify/require"
)

type msLevel int

func (l msLevel) String() string               { return "" }
func (l msLevel) Add(o msLevel) msLevel        { return l + o }
func (l msLevel) Value() (driver.Value, error) { return int64(l), nil }
func (l *msLevel) Scan(any) error              { return nil }

type MethodSet struct{ ent.Schema }

func (MethodSet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("level").GoType(msLevel(0)),
		field.Int("plain"),
	}
}

func TestMarshalSchema_MethodSet(t *testing.T) {
	// In-memory descriptors answer from the Go type.
	fd := field.Int("level").GoType(msLevel(0)).Descriptor()
	require.True(t, fd.Info.ValueScanner())
	require.True(t, fd.Info.Stringer())
	require.False(t, fd.Info.Validator())

	buf, err := MarshalSchema(&MethodSet{})
	require.NoError(t, err)
	s, err := UnmarshalSchema(buf)
	require.NoError(t, err)
	info := s.Fields[0].Info
	require.Contains(t, info.RType.Methods, "Scan", "pointer-receiver methods are recorded")
	require.Contains(t, info.RType.Methods, "Add")
	require.True(t, info.ValueScanner())
	require.True(t, info.Valuer())
	require.True(t, info.Stringer())
	require.False(t, info.Validator())
	require.Nil(t, s.Fields[1].Info.RType)
}
