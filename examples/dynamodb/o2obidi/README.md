# User-Spouse Bidirectional O2O Relation

An example for a reflexive O2O relation between a User to its spouse (also a User).    
Each user can have only one spouse. If a user A sets its spouse (using `spouse`) to B,
B can get its spouse using the `spouse` edge.

### Generate Assets

```console
go generate ./...
```

### Run Examples

```console
go test
```
