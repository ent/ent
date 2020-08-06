// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import "github.com/facebook/ent"

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Config of the User.
func (User) Config() ent.Config {
	return ent.Config{
		Table: "Users",
	}
}
