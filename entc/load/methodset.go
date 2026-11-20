// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package load

import (
	"reflect"

	"entgo.io/ent/internal/methodset"
	"entgo.io/ent/schema/field"
)

func init() {
	methodset.Implements = func(r any, typ reflect.Type) bool {
		return implements(r.(*field.RType), typ)
	}
}

// implements reports whether r ~implements the interface type typ,
// by comparing the method signatures of typ with r.Methods.
func implements(r *field.RType, typ reflect.Type) bool {
	n := typ.NumMethod()
	for i := 0; i < n; i++ {
		m0 := typ.Method(i)
		m1, ok := r.Methods[m0.Name]
		if !ok || len(m1.In) != m0.Type.NumIn() || len(m1.Out) != m0.Type.NumOut() {
			return false
		}
		in := m0.Type.NumIn()
		for j := 0; j < in; j++ {
			if !m1.In[j].TypeEqual(m0.Type.In(j)) {
				return false
			}
		}
		out := m0.Type.NumOut()
		for j := 0; j < out; j++ {
			if !m1.Out[j].TypeEqual(m0.Type.Out(j)) {
				return false
			}
		}
	}
	return true
}
