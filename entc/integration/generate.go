package integration

//go:generate go run ../cmd/entc/entc.go generate --storage=sql,gremlin ./ent/schema
//go:generate go run ../cmd/entc/entc.go generate --storage=sql,gremlin ./plugin/ent/schema
//go:generate go run ../cmd/entc/entc.go generate ./migrate/entv1/schema
//go:generate go run ../cmd/entc/entc.go generate ./migrate/entv2/schema
