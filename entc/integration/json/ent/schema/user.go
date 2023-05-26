// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("t", &T{}).
			Optional(),
		field.JSON("url", &url.URL{}).
			Optional(),
		field.JSON("URLs", []*url.URL{}).
			StorageKey("urls").
			StructTag(`json:"urls,omitempty"`).
			Optional(),
		field.JSON("raw", json.RawMessage{}).
			Optional(),
		field.JSON("dirs", []http.Dir{}).
			Default(func() []http.Dir {
				return []http.Dir{"/tmp"}
			}),
		field.Ints("ints").
			Optional().
			Default([]int{1, 2, 3}),
		field.Floats("floats").
			Optional(),
		field.Strings("strings").
			Optional(),
		field.Ints("ints_validate").
			Optional().
			Validate(validate[int]),
		field.Floats("floats_validate").
			Optional().
			Validate(validate[float64]),
		field.Strings("strings_validate").
			Optional().
			Validate(validate[string]),
		field.JSON("addr", Addr{}).
			Sensitive().
			Optional(),
		field.Any("unknown").
			Optional(),
	}
}

type T struct {
	I  int      `json:"i,omitempty"`
	F  float64  `json:"f,omitempty"`
	B  bool     `json:"b,omitempty"`
	S  string   `json:"s,omitempty"`
	T  *T       `json:"t,omitempty"`
	Li []int    `json:"li"`
	Ls []string `json:"ls"`
	// Do not omit empty or null maps.
	M map[string]any `json:"m"`
}

type Addr struct{ net.Addr }

func (a *Addr) UnmarshalJSON(data []byte) error {
	var types struct {
		TCP *net.TCPAddr `json:"tcp,omitempty"`
		UDP *net.UDPAddr `json:"udp,omitempty"`
	}
	if err := json.Unmarshal(data, &types); err != nil {
		return err
	}
	switch {
	case types.TCP != nil && types.UDP != nil:
		return errors.New("TCP and UDP addresses are mutually exclusive")
	case types.TCP != nil:
		a.Addr = types.TCP
	case types.UDP != nil:
		a.Addr = types.UDP
	}
	return nil
}

func (a Addr) MarshalJSON() ([]byte, error) {
	var types struct {
		TCP *net.TCPAddr `json:"tcp,omitempty"`
		UDP *net.UDPAddr `json:"udp,omitempty"`
	}
	switch a := a.Addr.(type) {
	case *net.TCPAddr:
		types.TCP = a
	case *net.UDPAddr:
		types.UDP = a
	default:
		return nil, fmt.Errorf("unsupported address type: %T", a)
	}
	return json.Marshal(types)
}

func (a Addr) String() string {
	if a.Addr == nil {
		return ""
	}
	return a.Addr.String()
}

var ErrValidate = errors.New("validation error")

func validate[T int | string | float64](xs []T) error {
	if xs != nil {
		return ErrValidate
	}
	return nil
}
