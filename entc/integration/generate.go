package integration

//go:generate go run ../cmd/entc/entc.go generate --storage=sql,gremlin --idtype string --header "Code generated (@generated) by entc, DO NOT EDIT." ./ent/schema
//go:generate go run ../cmd/entc/entc.go generate --header "Code generated (@generated) by entc, DO NOT EDIT." ./migrate/entv1/schema
//go:generate go run ../cmd/entc/entc.go generate --header "Code generated (@generated) by entc, DO NOT EDIT." ./migrate/entv2/schema
