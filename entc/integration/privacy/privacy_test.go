// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package privacy

import (
	"context"
	"errors"
	"testing"

	"github.com/facebook/ent/entc/integration/privacy/ent/enttest"
	"github.com/facebook/ent/entc/integration/privacy/ent/galaxy"
	"github.com/facebook/ent/entc/integration/privacy/ent/planet"
	"github.com/facebook/ent/entc/integration/privacy/ent/privacy"
	"github.com/facebook/ent/entc/integration/privacy/rule"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestPrivacyRules(t *testing.T) {
	client := enttest.Open(t, "sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	defer client.Close()
	logf := rule.SetMutationLogFunc(t.Logf)
	defer rule.SetMutationLogFunc(logf)

	ctx := context.Background()
	earth, err := client.Planet.Create().
		SetName("Earth").
		SetAge(4_540_000_000).
		Save(ctx)
	require.NoError(t, err)
	mars := client.Planet.Create().
		SetName("Mars").
		SaveX(ctx)
	err = earth.Update().
		AddNeighbors(mars).
		Exec(ctx)
	require.NoError(t, err)

	logf = rule.SetMutationLogFunc(func(string, ...interface{}) {
		require.FailNow(t, "hook called on privacy deny")
	})
	err = client.Planet.Update().
		Where(planet.ID(earth.ID)).
		SetAge(4_600_000_000).
		Exec(ctx)
	require.True(t, errors.Is(err, privacy.Deny))
	err = earth.Update().
		AddNeighbors(earth).
		Exec(ctx)
	require.True(t, errors.Is(err, privacy.Deny))
	rule.SetMutationLogFunc(logf)

	err = client.Planet.Update().
		Where(planet.ID(earth.ID)).
		SetAge(4_600_000_000).
		Exec(privacy.DecisionContext(ctx, privacy.Allow))
	require.NoError(t, err)

	count := client.Planet.Query().CountX(ctx)
	require.Equal(t, 1, count)
	mars.Update().SetAge(6_000_000_000).ExecX(ctx)
	count = client.Planet.Query().CountX(ctx)
	require.Equal(t, 2, count)

	client.Galaxy.Create().
		SetName("Milky Way").
		SetType(galaxy.TypeBarredSpiral).
		AddPlanets(earth, mars).
		SaveX(ctx)
	client.Galaxy.Create().
		SetName("IC 3583").
		SetType(galaxy.TypeIrregular).
		SaveX(ctx)
	count = client.Galaxy.Query().CountX(ctx)
	require.Equal(t, 1, count)
}
