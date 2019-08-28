---
id: crud
title: CRUD API
---

As mentioned in the [introduction](code-gen.md) section, running `entc` on the schemas,
will generate the following assets:

- `Client` and `Tx` objects used for interacting with the graph.
- CRUD builders for each schema type. See [CRUD](crud.md) for more info.
- Entity object (Go struct) for each of the schema type.
- Package contains constants and predicates used for interacting with the builders.
- A `migrate` package, for SQL dialects. See [Migration](migrate.md) for more info.

## Create A New Client

**MySQL**

```go
package main

import (
	"log"

	"<project>/ent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/facebookincubator/ent/dialect/sql"
)

func main() {
	drv, err := sql.Open("mysql", "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer drv.Close()
	client := ent.NewClient(ent.Driver(drv))
}
```

**SQLite**

```go
package main

import (
	"log"

	"<project>/ent"

	_ "github.com/mattn/go-sqlite3"
	"github.com/facebookincubator/ent/dialect/sql"
)

func main() {
	drv, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer drv.Close()
	client := ent.NewClient(ent.Driver(drv))
}
```


**Gremlin (AWS Neptune)**

```go
package main

import (
	"log"
	"net/url"

	"<project>/ent"

	"github.com/facebookincubator/ent/dialect/gremlin"
)

func main() {
	c, err := gremlin.NewClient(gremlin.Config{
		Endpoint: gremlin.Endpoint{
			URL: &url.URL{
				Scheme: "http",
				Host:   "localhost:8182",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	client := ent.NewClient(ent.Driver(gremlin.NewDriver(c)))
}
```

## Create An Entity

## Update An Entity

## Update Many Entities

## Query An Entity

## Delete An Entity

## Delete Entities