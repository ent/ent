// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build goexperiment.jsonv2

package schema

import (
	"encoding/json"
	"encoding/json/jsontext"
	jsonv2 "encoding/json/v2"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Record struct{ ent.Schema }

func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.JSON("legacy", &Payload{}).Optional(),
		field.JSON("modern", &Payload{}).JSONOptions(jsonv2.DefaultOptionsV2()).Optional().StorageKey("modern_data"),
		field.JSON("entries", []Payload{}).JSONOptions(jsonv2.DefaultOptionsV2()).Optional(),
		field.Strings("tags").JSONOptions(jsonv2.DefaultOptionsV2()).Default(nil),
		field.JSON("dictionary", map[string]string{}).JSONOptions(jsonv2.DefaultOptionsV2()).Optional(),
		field.JSON("stream", &Stream{}).JSONOptions(jsonv2.DefaultOptionsV2()).Optional(),
		field.JSON("raw", json.RawMessage{}).JSONOptions(jsonv2.DefaultOptionsV2()).Optional(),
	}
}

type Payload struct {
	Name  string   `json:"name"`
	Count int      `json:"count,omitempty"`
	Items []string `json:"items"`
}

type Stream struct{ Value string }

func (s Stream) MarshalJSONTo(enc *jsontext.Encoder) error {
	return enc.WriteToken(jsontext.String("v2:" + s.Value))
}

func (s *Stream) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	token, err := dec.ReadToken()
	if err != nil {
		return err
	}
	if token.Kind() != '"' || !strings.HasPrefix(token.String(), "v2:") {
		return fmt.Errorf("invalid stream value")
	}
	s.Value = strings.TrimPrefix(token.String(), "v2:")
	return nil
}

type Configured struct{ ent.Schema }

func (Configured) Mixin() []ent.Mixin { return []ent.Mixin{OptionsMixin{}} }

func (Configured) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.JSON("payload", &Payload{}).JSONOptions(jsonv2.StringifyNumbers(true)).Optional(),
		field.JSON("ordered", map[string]int{}).JSONOptions(jsonv2.Deterministic(true)).Optional(),
		field.JSON("entries", []Custom{}).JSONOptions(
			jsonv2.DefaultOptionsV2(),
			jsonv2.WithMarshalers(jsonv2.MarshalFunc(func(v Custom) ([]byte, error) {
				return jsonv2.Marshal("custom:" + string(v))
			})),
			jsonv2.WithUnmarshalers(jsonv2.UnmarshalFunc(func(data []byte, v *Custom) error {
				var s string
				if err := jsonv2.Unmarshal(data, &s); err != nil {
					return err
				}
				*v = Custom(strings.TrimPrefix(s, "custom:"))
				return nil
			})),
		).Optional(),
		field.JSON("strict", &Payload{}).JSONOptions(jsonv2.DefaultOptionsV2(), jsonv2.RejectUnknownMembers(true)).Optional(),
		field.JSON("raw", json.RawMessage{}).JSONOptions(jsontext.AllowDuplicateNames(true)).Optional(),
	}
}

type Custom string

type OptionsMixin struct{ mixin.Schema }

func (OptionsMixin) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("mixed_legacy", &Payload{}).Optional(),
		field.JSON("mixed", []int{}).JSONOptions(jsonv2.StringifyNumbers(true)).Optional(),
	}
}
