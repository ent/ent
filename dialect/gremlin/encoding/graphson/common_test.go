package graphson

import (
	"unsafe"

	"github.com/json-iterator/go"
	"github.com/stretchr/testify/mock"
)

type mocker struct {
	mock.Mock
}

// Encode belongs to jsoniter.ValEncoder interface.
func (m *mocker) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	m.Called(ptr, stream)
}

// IsEmpty belongs to jsoniter.ValEncoder interface.
func (m *mocker) IsEmpty(ptr unsafe.Pointer) bool {
	args := m.Called(ptr)
	return args.Bool(0)
}

// Decode implements jsoniter.ValDecoder interface.
func (m *mocker) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	m.Called(ptr, iter)
}

// CheckType implements typeChecker interface.
func (m *mocker) CheckType(typ Type) error {
	args := m.Called(typ)
	return args.Error(0)
}

// MarshalGraphson implements Marshaler interface.
func (m *mocker) MarshalGraphson() ([]byte, error) {
	args := m.Called()
	data, err := args.Get(0), args.Error(1)
	if data == nil {
		return nil, err
	}
	return data.([]byte), err
}

// UnmarshalGraphson implements Unmarshaler interface.
func (m *mocker) UnmarshalGraphson(data []byte) error {
	args := m.Called(data)
	return args.Error(0)
}
