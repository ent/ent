// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("url", &url.URL{}).
			Optional(),
		field.JSON("raw", json.RawMessage{}).
			Optional(),
		field.JSON("dirs", []http.Dir{}).
			Optional(),
		field.Ints("ints").
			Optional(),
		field.Floats("floats").
			Optional(),
		field.Strings("strings").
			Optional(),
	}
}
