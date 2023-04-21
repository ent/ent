// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("id").
			MaxLen(64).
			GoType(ID{}).
			DefaultFunc(NewID),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("device", Device.Type).
			Ref("sessions").
			Unique(),
	}
}

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("id").
			MaxLen(64).
			GoType(ID{}).
			DefaultFunc(NewID),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("active_session", Session.Type).
			Unique(),
		edge.To("sessions", Session.Type),
	}
}

type ID [64]byte

func NewID() ID {
	var id [64]byte
	copy(id[:], uuid.NewString()+uuid.NewString()+uuid.NewString()+uuid.NewString())
	return id
}

func (i ID) String() string {
	return string(i[:])
}

func (i *ID) Scan(v any) error {
	switch v := v.(type) {
	case []byte:
		copy(i[:], v)
	case string:
		copy(i[:], v)
	default:
		return fmt.Errorf("unexpected type: %T", v)
	}
	return nil
}

func (i ID) Value() (driver.Value, error) {
	return string(i[:]), nil
}
