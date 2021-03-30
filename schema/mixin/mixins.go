// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/privacy"
	"entgo.io/ent/schema"
)

func NewMixins(mixins ...ent.Mixin) ent.Mixin {
	return Mixins(mixins)
}

type Mixins []ent.Mixin

func (m Mixins) Fields() []ent.Field {
	var o []ent.Field
	for _, v := range m {
		o = append(o, v.Fields()...)
	}
	return o
}

func (m Mixins) Edges() []ent.Edge {
	var o []ent.Edge
	for _, v := range m {
		o = append(o, v.Edges()...)
	}
	return o
}

func (m Mixins) Indexes() []ent.Index {
	var o []ent.Index
	for _, v := range m {
		o = append(o, v.Indexes()...)
	}
	return o
}

func (m Mixins) Hooks() []ent.Hook {
	var o []ent.Hook
	for _, v := range m {
		o = append(o, v.Hooks()...)
	}
	return o
}
func (m Mixins) Policy() ent.Policy {
	var o []ent.Policy
	for _, v := range m {
		if v.Policy() != nil {
			o = append(o, v.Policy())
		}
	}
	switch len(o) {
	case 0:
		return nil
	case 1:
		return o[0]
	}
	return privacy.Policies(o)
}

func (m Mixins) Annotations() []schema.Annotation {
	var o []schema.Annotation
	for _, v := range m {
		o = append(o, v.Annotations()...)
	}
	return o
}
