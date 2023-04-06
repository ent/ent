### entc integration tests

#### Regenerating assets

If you edited one of the files in `entc/gen/template` or `entc/load/template`,
run the following command:

```
go generate ./...
```

#### Running the integration tests

```
docker-compose up -d --scale gremlin=0
go test .
```

In order to run the Gremlin tests, run:

```
docker-compose up -d gremlin
go test ./gremlin/...
```

Use the `-run` flag for running specific test or set of tests. For example:

```
go test -run=MySQL

go test -run=MySQL/8/Sanity

go test -run=SQLite/Sanity
```
