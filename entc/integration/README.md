### entc integration tests

#### Regenerating new templates

If you edited one of the files in `entc/gen/template` or `entc/load/template`,
run the following command to from `entc` directory:

```
go generate ./...
```

#### Running the integration tests

```
docker-compose -f compose/docker-compose.yaml up -d --scale gremlin=0
go test 
```

Use the `-run` flag for running specific test or set of tests. For example:
```
go test -run=MySQL

go test -run=SQLite/Sanity
```

