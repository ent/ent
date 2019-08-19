### entc integration tests

#### Regenerating new templates

If you edited one of the files in `entc/gen/template` or `entc/build/template`,
please run the following command:

For `entc/gen` 
```
cd ~/fbsource/fbcode/github.com/facebookincubator/ent/entc/gen && go generate && cd -
``` 

For `entc/build`

```
cd ~/fbsource/fbcode/github.com/facebookincubator/ent/entc/gen && go generate && cd -
```

Then, regenerate new assets for your schema:
```
go run ~/fbsource/fbcode/github.com/facebookincubator/ent/entc/cmd/entc/entc.go generate ./ent/schema
```

#### Running the integration tests

```
docker-compose -f compose/docker-compose.yaml up -d
go test 
```

Use the `-run` flag for running specific test or set of tests. For example:
```
go test -run=MySQL

go test -run=SQLite/Sanity
```
