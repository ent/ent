package privacy

import (
	"context"
	"errors"
	"testing"

	"github.com/facebookincubator/ent/entc/integration/privacy/ent"
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
	ctx := context.Background()
	err = client.Schema.Create(ctx)
	require.NoError(t, err)
	logf := rule.SetMutationLogFunc(t.Logf)
	defer rule.SetMutationLogFunc(logf)

	earth, err := client.Planet.Create().SetName("Earth").SetAge(4_540_000_000).Save(ctx)
	require.NoError(t, err)
	err = client.Planet.Update().Where(planet.ID(earth.ID)).SetAge(4_600_000_000).Exec(ctx)
	require.True(t, errors.Is(err, privacy.Deny))
	err = earth.Update().AddNeighbors(earth).Exec(ctx)
	require.True(t, errors.Is(err, privacy.Deny))
	mars := client.Planet.Create().SetName("Mars").SaveX(ctx)
	err = earth.Update().AddNeighbors(mars).Exec(ctx)
	require.NoError(t, err)
}
