// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package customid

import (
	"context"
	"database/sql"
	"fmt"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/nativeenum/ent/car"
	"github.com/facebook/ent/entc/integration/nativeenum/ent/mood"
	"github.com/facebook/ent/entc/integration/nativeenum/ent/person"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/entc/integration/nativeenum/ent"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestPostgresNativeEnum(t *testing.T) {
	for version, port := range map[string]int{
		//"10": 5430,
		//"11": 5431,
		"12": 5433,
	} {
		t.Run(version, func(t *testing.T) {
			for _, tt := range tests {
				name := runtime.FuncForPC(reflect.ValueOf(tt).Pointer()).Name()
				t.Run(name[strings.LastIndex(name, ".")+1:], func(t *testing.T) {
					dsn := fmt.Sprintf("host=localhost port=%d user=postgres password=pass sslmode=disable", port)
					db, err := sql.Open(dialect.Postgres, dsn)
					require.NoError(t, err)
					defer db.Close()
					_, err = db.Exec("DROP DATABASE IF EXISTS native_enum")
					require.NoError(t, err, "dropping database database")
					_, err = db.Exec("CREATE DATABASE native_enum")
					require.NoError(t, err, "creating database")
					defer db.Exec("DROP DATABASE native_enum")

					nativeEnumDb, err := sql.Open(dialect.Postgres, dsn+" dbname=native_enum")
					require.NoError(t, err, "opening connection to native_enum db")
					nativeEnumDriver := entsql.OpenDB(dialect.Postgres, nativeEnumDb)

					client := ent.NewClient(ent.Driver(nativeEnumDriver))
					require.NoError(t, err, "connecting to json database")
					defer client.Close()
					err = client.Schema.Create(context.Background())
					require.NoError(t, err)

					tt(t, client, nativeEnumDb)
				})
			}
		})
	}
}

var (
	tests = [...]func(*testing.T, *ent.Client, *sql.DB){
		SaveAndQuery,
		MigrateStringToEnum,
		MigrateEnumTypeModified,
		MigrateEnumToString,
		MigrateAddEnumColumn,
	}
)

func SaveAndQuery(t *testing.T, client *ent.Client, db *sql.DB) {
	require := require.New(t)

	p := client.Person.Create().SetMood(mood.Happy).SaveX(context.Background())
	require.Equal(mood.Happy, p.Mood)

	pQuery := client.Person.Query().Where(person.MoodEQ(mood.Happy)).OnlyX(context.Background())
	require.Equal(p.ID, pQuery.ID)
}

func MigrateStringToEnum(t *testing.T, client *ent.Client, db *sql.DB) {
	require := require.New(t)

	client.Person.Create().SetMood(mood.Happy).SaveX(context.Background())
	client.Person.Create().SetMood(mood.Sad).SaveX(context.Background())
	client.Person.Create().SetMood(mood.Ok).SaveX(context.Background())

	_, err := db.Exec(`ALTER TABLE "persons" ALTER COLUMN "mood" TYPE varchar(255)`)
	require.NoError(err, "altering table")
	_, err = db.Exec(`DROP TYPE mood`)
	require.NoError(err, "drop mood type")

	err = client.Schema.Create(context.Background())
	require.NoError(err, "migrate string to enum")

	p := client.Person.Create().SetMood(mood.Happy).SaveX(context.Background())
	require.Equal(mood.Happy, p.Mood)

	count := client.Person.Query().Where(person.MoodEQ(mood.Happy)).CountX(context.Background())
	require.Equal(2, count)
}

func MigrateEnumTypeModified(t *testing.T, client *ent.Client, db *sql.DB) {
	require := require.New(t)

	client.Person.Create().SetMood(mood.Happy).SaveX(context.Background())
	client.Person.Create().SetMood(mood.Sad).SaveX(context.Background())

	_, err := db.Exec(`ALTER TYPE "mood" ADD VALUE 'vibing'`)
	require.NoError(err, "altering type")

	err = client.Schema.Create(context.Background())
	require.NoError(err, "migrate enum renamed")

	p := client.Person.Create().SetMood(mood.Happy).SaveX(context.Background())
	require.Equal(mood.Happy, p.Mood)

	count := client.Person.Query().Where(person.MoodEQ(mood.Happy)).CountX(context.Background())
	require.Equal(2, count)
}

func MigrateEnumToString(t *testing.T, client *ent.Client, db *sql.DB) {
	require := require.New(t)

	client.Car.Create().SetTransmission("automatic").SaveX(context.Background())
	client.Car.Create().SetTransmission("manual").SaveX(context.Background())

	_, err := db.Exec(`CREATE TYPE "transmission" AS ENUM ('automatic', 'manual')`)
	require.NoError(err, "creating type")
	_, err = db.Exec(`ALTER TABLE "cars" ALTER COLUMN "transmission" TYPE transmission using (transmission::transmission)`)
	require.NoError(err, "altering table")

	err = client.Schema.Create(context.Background())
	require.NoError(err, "migrate enum to string")

	count := client.Car.Query().Where(car.Transmission("automatic")).CountX(context.Background())
	require.Equal(1, count)
}

func MigrateAddEnumColumn(t *testing.T, client *ent.Client, db *sql.DB) {
	require := require.New(t)

	_, err := db.Exec(`ALTER TABLE "persons" DROP COLUMN "mood"`)
	require.NoError(err, "dropping column")

	err = client.Schema.Create(context.Background())
	require.NoError(err, "migrate enum column added")

	p := client.Person.Create().SetMood(mood.Happy).SaveX(context.Background())
	require.Equal(mood.Happy, p.Mood)

	count := client.Person.Query().Where(person.MoodEQ(mood.Happy)).CountX(context.Background())
	require.Equal(1, count)
}
