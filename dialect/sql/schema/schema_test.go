package schema

import (
	"testing"

	"fbc/ent/field"

	"github.com/stretchr/testify/require"
)

func TestColumn_ConvertibleTo(t *testing.T) {
	c1 := &Column{Type: field.TypeString, Size: 10}
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 10}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 255}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 9}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeFloat32}))

	c1 = &Column{Type: field.TypeFloat32}
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeFloat32}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeFloat64}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint}))

	c1 = &Column{Type: field.TypeFloat64}
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeFloat32}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeFloat64}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint}))

	c1 = &Column{Type: field.TypeUint}
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeUint}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeInt}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeInt64}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeUint64}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeInt8}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint8}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint16}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint32}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString}))

	c1 = &Column{Type: field.TypeInt}
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeInt}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeInt64}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeInt8}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeInt32}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint8}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint16}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint32}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString}))
}
