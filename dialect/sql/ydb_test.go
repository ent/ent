package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYDB_Placeholder(t *testing.T) {
	d := YDB{}
	assert.Equal(t, "$", d.Placeholder())
}

func TestYDB_Array(t *testing.T) {
	d := YDB{}
	assert.Equal(t, "?", d.Array())
}

func TestYDB_Drivers(t *testing.T) {
	d := YDB{}
	drivers := d.Drivers()
	assert.Equal(t, []string{"ydb"}, drivers)
}

func TestYDB_Schema(t *testing.T) {
	d := YDB{}
	assert.Equal(t, "ydb", d.Schema())
}

func TestYDB_ConvertType(t *testing.T) {
	d := YDB{}
	tests := []struct {
		name     string
		input    interface{}
		expected YDBType
		wantErr  bool
	}{
		{"int8", int8(1), YDBTypeInt8, false},
		{"int16", int16(1), YDBTypeInt16, false},
		{"int32", int32(1), YDBTypeInt32, false},
		{"int64", int64(1), YDBTypeInt64, false},
		{"uint8", uint8(1), YDBTypeUint8, false},
		{"uint16", uint16(1), YDBTypeUint16, false},
		{"uint32", uint32(1), YDBTypeUint32, false},
		{"uint64", uint64(1), YDBTypeUint64, false},
		{"float32", float32(1), YDBTypeFloat, false},
		{"float64", float64(1), YDBTypeDouble, false},
		{"string", "test", YDBTypeString, false},
		{"[]byte", []byte("test"), YDBTypeBytes, false},
		{"bool", true, YDBTypeBool, false},
		{"unsupported", struct{}{}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.ConvertType(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestYDBDriver_ModifyQuery(t *testing.T) {
	d := &YDBDriver{}
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "replace LIMIT",
			input:    "SELECT * FROM users LIMIT 10",
			expected: "SELECT * FROM users TOP 10",
		},
		{
			name:     "no modification needed",
			input:    "SELECT * FROM users",
			expected: "SELECT * FROM users",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := d.modifyQuery(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}
