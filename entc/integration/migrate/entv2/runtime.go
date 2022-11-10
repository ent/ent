// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"time"

	"entgo.io/ent/entc/integration/migrate/entv2/schema"
	"entgo.io/ent/entc/integration/migrate/entv2/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescMixedString is the schema descriptor for mixed_string field.
	userDescMixedString := userMixinFields0[0].Descriptor()
	// user.DefaultMixedString holds the default value on creation for the mixed_string field.
	user.DefaultMixedString = userDescMixedString.Default.(string)
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[1].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
	// userDescNickname is the schema descriptor for nickname field.
	userDescNickname := userFields[5].Descriptor()
	// user.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	user.NicknameValidator = userDescNickname.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[6].Descriptor()
	// user.DefaultPhone holds the default value on creation for the phone field.
	user.DefaultPhone = userDescPhone.Default.(string)
	// userDescBuffer is the schema descriptor for buffer field.
	userDescBuffer := userFields[7].Descriptor()
	// user.DefaultBuffer holds the default value on creation for the buffer field.
	user.DefaultBuffer = userDescBuffer.Default.(func() []byte)
	// userDescTitle is the schema descriptor for title field.
	userDescTitle := userFields[8].Descriptor()
	// user.DefaultTitle holds the default value on creation for the title field.
	user.DefaultTitle = userDescTitle.Default.(string)
	// userDescNewToken is the schema descriptor for new_token field.
	userDescNewToken := userFields[10].Descriptor()
	// user.DefaultNewToken holds the default value on creation for the new_token field.
	user.DefaultNewToken = userDescNewToken.Default.(func() string)
	// userDescBlob is the schema descriptor for blob field.
	userDescBlob := userFields[11].Descriptor()
	// user.BlobValidator is a validator for the "blob" field. It is called by the builders before save.
	user.BlobValidator = userDescBlob.Validators[0].(func([]byte) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[16].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescDropOptional is the schema descriptor for drop_optional field.
	userDescDropOptional := userFields[17].Descriptor()
	// user.DefaultDropOptional holds the default value on creation for the drop_optional field.
	user.DefaultDropOptional = userDescDropOptional.Default.(func() string)
}
