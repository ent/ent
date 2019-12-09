# Tree O2M Relation

In this example, we have a recursive O2M relation between tree's nodes and their children (or their parent).  
Each node in the tree **has many** children, and **has one** parent. If node A adds B to its children,
B can get its owner using the `owner` edge.

### Generate Assets

```console
go generate ./...
```

### Run Examples

```console
go test
```
