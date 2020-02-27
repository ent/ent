package field

import (
	"database/sql/driver"
	"reflect"
)

// customBuilder is the builder for int field.
type customBuilder struct {
	desc *Descriptor
}

// Int returns a new Field with type int.
func Custom(name string, typeName string, typ driver.Valuer) *customBuilder {
	rt := reflect.TypeOf(typ)
	return &customBuilder{&Descriptor{
		Name: name,
		Info: &TypeInfo{
			Type:       TypeCustom,
			Nillable:   true,
			Ident:      rt.String(),
			PkgPath:    rt.PkgPath(),
			CustomType: typeName,
		},
	}}
}

// Descriptor implements the ent.Field interface by returning its descriptor.
func (b *customBuilder) Descriptor() *Descriptor {
	return b.desc
}
