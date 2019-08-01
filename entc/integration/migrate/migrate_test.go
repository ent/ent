package migrate

import (
	"context"
	"strconv"
	"testing"

	"fbc/ent/dialect/sql"
	"fbc/ent/dialect/sql/schema"
	"fbc/ent/entc/integration/migrate/entv1"
	"fbc/ent/entc/integration/migrate/entv2"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestMySQL(t *testing.T) {
	root, err := sql.Open("mysql", "root:pass@tcp(localhost:3306)/")
	require.NoError(t, err)
	defer root.Close()
	ctx := context.Background()
	err = root.Exec(ctx, "CREATE DATABASE migrate", []interface{}{}, new(sql.Result))
	require.NoError(t, err, "creating database")
	defer root.Exec(ctx, "DROP DATABASE migrate", []interface{}{}, new(sql.Result))

	drv, err := sql.Open("mysql", "root:pass@tcp(localhost:3306)/migrate?parseTime=True")
	require.NoError(t, err, "connecting to migrate database")

	// run migration and execute queries on v1.
	clientv1 := entv1.NewClient(entv1.Driver(drv))
	require.NoError(t, clientv1.Schema.Create(ctx, schema.WithGlobalUniqueID(true)))
	SanityV1(t, clientv1)

	// run migration and execute queries on v2.
	clientv2 := entv2.NewClient(entv2.Driver(drv))
	require.NoError(t, clientv2.Schema.Create(ctx, schema.WithGlobalUniqueID(true)))
	SanityV2(t, clientv2)

	// since "users" created in the migration of v1, it will occupy the range of 0 ... 1<<32-1,
	// even though they are ordered differently in the migration of v2 (groups, pets, users).
	idRange(t, clientv2.User.Create().SetAge(1).SetName("foo").SetPhone("100").SaveX(ctx).ID, 0, 1<<32)
	idRange(t, clientv2.Group.Create().SaveX(ctx).ID, 1<<32-1, 2<<32)
	idRange(t, clientv2.Pet.Create().SaveX(ctx).ID, 2<<32-1, 3<<32)
}

func TestSQLite(t *testing.T) {
	drv, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer drv.Close()

	ctx := context.Background()
	client := entv2.NewClient(entv2.Driver(drv))
	require.NoError(t, client.Schema.Create(ctx, schema.WithGlobalUniqueID(true)))

	SanityV2(t, client)
	idRange(t, client.Group.Create().SaveX(ctx).ID, 0, 1<<32)
	idRange(t, client.Pet.Create().SaveX(ctx).ID, 1<<32-1, 2<<32)
	idRange(t, client.User.Create().SetAge(1).SetName("x").SetPhone("y").SaveX(ctx).ID, 2<<32, 3<<32-1)
}

func SanityV1(t *testing.T, client *entv1.Client) {
	ctx := context.Background()
	u := client.User.Create().SetAge(1).SetName("foo").SaveX(ctx)
	require.Equal(t, int32(1), u.Age)
	require.Equal(t, "foo", u.Name)

	_, err := client.User.Create().SetAge(1).SetName("foobarbazqux").Save(ctx)
	require.Error(t, err, "name is limited to 10 chars")
}

func SanityV2(t *testing.T, client *entv2.Client) {
	ctx := context.Background()
	u := client.User.Create().SetAge(1).SetName("bar").SetPhone("100").SaveX(ctx)
	require.Equal(t, 1, u.Age)
	require.Equal(t, "bar", u.Name)
	require.Equal(t, []byte("{}"), u.Buffer)
	u = u.Update().SetBuffer([]byte("[]")).SaveX(ctx)
	require.Equal(t, []byte("[]"), u.Buffer)

	_, err := client.User.Create().SetAge(1).SetName("foobarbazqux").SetPhone("100").Save(ctx)
	require.NoError(t, err, "name is not limited to 10 chars")
}

func idRange(t *testing.T, s string, l, h int) {
	id, err := strconv.Atoi(s)
	require.NoError(t, err)
	require.Truef(t, id > l && id < h, "id %s should be between %d to %d", s, l, h)
}
