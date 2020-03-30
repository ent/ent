package privacy

import (
	"context"
	"errors"
	"testing"

	"github.com/facebookincubator/ent/entc/integration/privacy/ent"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/galaxy"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/planet"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/privacy"
	_ "github.com/facebookincubator/ent/entc/integration/privacy/ent/runtime"
	"github.com/facebookincubator/ent/entc/integration/privacy/rule"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestPrivacyRules(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	ctx := context.Background()
	err = client.Schema.Create(ctx)
	require.NoError(t, err)
	logf := rule.SetMutationLogFunc(t.Logf)
	defer rule.SetMutationLogFunc(logf)

	earth, err := client.Planet.Create().SetName("Earth").SetAge(4_540_000_000).Save(ctx)
	require.NoError(t, err)
	mars := client.Planet.Create().SetName("Mars").SaveX(ctx)
	err = earth.Update().AddNeighbors(mars).Exec(ctx)
	require.NoError(t, err)

	logf = rule.SetMutationLogFunc(func(string, ...interface{}) {
		require.FailNow(t, "hook called on privacy deny")
	})
	err = client.Planet.Update().Where(planet.ID(earth.ID)).SetAge(4_600_000_000).Exec(ctx)
	require.True(t, errors.Is(err, privacy.Deny))
	err = earth.Update().AddNeighbors(earth).Exec(ctx)
	require.True(t, errors.Is(err, privacy.Deny))
	rule.SetMutationLogFunc(logf)

	count := client.Planet.Query().CountX(ctx)
	require.Equal(t, 1, count)
	mars.Update().SetAge(6_000_000_000).ExecX(ctx)
	count = client.Planet.Query().CountX(ctx)
	require.Equal(t, 2, count)

	client.Galaxy.Create().SetName("Milky Way").SetType(galaxy.TypeBarredSpiral).AddPlanets(earth, mars).SaveX(ctx)
	client.Galaxy.Create().SetName("IC 3583").SetType(galaxy.TypeIrregular).SaveX(ctx)
	count = client.Galaxy.Query().CountX(ctx)
	require.Equal(t, 1, count)
}
