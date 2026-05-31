---
id: testing
title: Testing
---

If you're using `ent.Client` in your unit-tests, you can use the generated `enttest`
package for creating a client and auto-running the schema migration as follows:

```go
package main

import (
	"testing"

	"<project>/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
)

func TestXXX(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	// ...
}
```

In order to pass functional options to `Open`, use `enttest.Option`:

```go
func TestXXX(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()
	// ...
}
```

In order to use [`go-sqlmock`](https://github.com/DATA-DOG/go-sqlmock) for testing, use `ent.Driver` to create your `ent.Client`:

```go
func CreateDBDriver(cfgStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfgString) // pg example
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	return db, nil
}

func CreateEntClient(dbDriver *sql.DB) (*ent.Client,error) {
	return ent.NewClient(ent.Driver(
		entsql.OpenDB("postgres", dbDriver)),
	), nil
}

func QueryPets(ctx context.Context, client *ent.Client) ([]*ent.Pet, error) {
	return client.Pets.Query().All(ctx)
}

func TestQueryPets(t *testing.T) {
// create mock db driver and sql mocker
	sqlmockDbDriver, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error create mock %v", err)
	}

	// create the db struct
	db, err := CreateEntClient(sqlmockDbDriver)
	if err != nil {
		t.Fatalf("error create db %v", err)
	}
	
	// mock rows
	mockRows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test")
	// mock query
	mock.ExpectQuery(`SELECT (.+) FROM "pets"`).WillReturnRows(mockRows)

	// do query
	pets, err := QueryPets(context.TODO(), db)
	if err != nil {
		t.Errorf("fail %v", err)
	}

	t.Log("pass")
}
```
