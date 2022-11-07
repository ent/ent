// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"net/http"

	"entgo.io/ent/entc/integration/json/ent/schema"
	"entgo.io/ent/entc/integration/json/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescDirs is the schema descriptor for dirs field.
	userDescDirs := userFields[4].Descriptor()
	// user.DefaultDirs holds the default value on creation for the dirs field.
	user.DefaultDirs = userDescDirs.Default.(func() []http.Dir)
	// userDescInts is the schema descriptor for ints field.
	userDescInts := userFields[5].Descriptor()
	// user.DefaultInts holds the default value on creation for the ints field.
	user.DefaultInts = userDescInts.Default.([]int)
}
