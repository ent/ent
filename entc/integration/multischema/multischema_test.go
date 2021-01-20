// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package json

import (
	"context"
	"strings"
	"testing"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/multischema/ent"
	"github.com/facebook/ent/entc/integration/multischema/ent/migrate"
	"github.com/stretchr/testify/require"

	_ "github.com/go-sql-driver/mysql"
)

func TestMySQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:pass@tcp(localhost:3308)/")
	require.NoError(t, err)
	defer db.Close()
	ctx := context.Background()
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS db1")
	require.NoError(t, err, "creating database")
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS db2")
	require.NoError(t, err, "creating database")
	defer db.ExecContext(ctx, "DROP DATABASE IF EXISTS db1")
	defer db.ExecContext(ctx, "DROP DATABASE IF EXISTS db2")
	setupSchema(t, db)

	client := ent.NewClient(ent.Driver(db), ent.AlternateSchema(ent.SchemaConfig{
		Pet:        "db1",
		User:       "db2",
		Group:      "db1",
		GroupUsers: "db2",
	}))
	pedro := client.Pet.Create().SaveX(ctx)
	groups := client.Group.CreateBulk(
		client.Group.Create().SetName("GitHub"),
		client.Group.Create().SetName("GitLab"),
	).SaveX(ctx)
	client.User.Create().AddPets(pedro).AddGroups(groups...).SaveX(ctx)

	// id := client.Group.Query().
	//	Where(group.HasUsersWith(user.ID(usr.ID))).
	//	OnlyIDX(ctx)
	//	require.Equal(t, pedro.ID, id)
}

func setupSchema(t *testing.T, drv *sql.Driver) {
	client := ent.NewClient(ent.Driver(&rewriter{drv}))
	err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false))
	require.NoError(t, err)
}

type rewriter struct {
	dialect.Driver
}

func (r *rewriter) Tx(ctx context.Context) (dialect.Tx, error) {
	tx, err := r.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txRewriter{tx}, nil
}

type txRewriter struct {
	dialect.Tx
}

func (r *txRewriter) Exec(ctx context.Context, query string, args, v interface{}) error {
	rp := strings.NewReplacer("`groups`", "`db1`.`groups`", "`pets`", "`db1`.`pets`", "`users`", "`db2`.`users`", "`group_users`", "`db2`.`group_users`")
	query = rp.Replace(query)
	return r.Tx.Exec(ctx, query, args, v)
}
