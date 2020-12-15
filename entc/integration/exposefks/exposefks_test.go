// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package exposefks

import (
	"context"
	"testing"

	"github.com/facebook/ent/entc/integration/exposefks/ent/enttest"
	"github.com/facebook/ent/entc/integration/exposefks/ent/migrate"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestExposeFKs(t *testing.T) {
	client := enttest.Open(t, "sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	defer client.Close()
	ctx := context.Background()
	require.NoError(t, client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)))

	a8m := client.User.Create().SetName("a8m").SaveX(ctx)
	dogFood := client.Food.Create().SetName("dog food").SaveX(ctx)

	errorhandler := client.User.Create().SetName("errorhandler").SaveX(ctx)
	mittens := client.Pet.Create().SetName("mittens").SetOwner(errorhandler).SaveX(ctx)
	fred := client.Pet.Create().SetName("fred").SetOwner(a8m).SetFavouriteFood(dogFood).SaveX(ctx)

	require.Nil(t, mittens.PetFavouriteFood)
	require.Equal(t, errorhandler.ID, mittens.UserPets)

	require.Equal(t, &dogFood.ID, fred.PetFavouriteFood)
	require.Equal(t, a8m.ID, fred.UserPets)
}
