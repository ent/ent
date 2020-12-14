// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package integration

import (
	"context"
	"github.com/facebook/ent/entc/integration/ent"
	"github.com/facebook/ent/entc/integration/ent/node"
	"github.com/stretchr/testify/require"
	"testing"
)

func Issue1(t *testing.T, client *ent.Client) {
	require := require.New(t)
	ctx := context.Background()

	nd := client.Node.Create().SetValue(1e3).SaveX(ctx)
	err := client.Node.Update().Where(node.ID(nd.ID)).Exec(ctx)
	require.NoError(err)
}
