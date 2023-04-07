// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package packagealias

import (
	"testing"

	entuser "entgo.io/ent/entc/integration/packagealias/ent/user"
	"entgo.io/ent/entc/integration/packagealias/user"
	"github.com/stretchr/testify/require"
)

func TestPackageAlias(t *testing.T) {
	require.Equal(t, entuser.State("on"), entuser.StateOn)
	require.Equal(t, entuser.State("off"), entuser.StateOff)
	entuser.Account(&user.User{Name: "test"})
}
