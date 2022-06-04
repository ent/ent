# Linked-List O2O Relation Example

An example for a O2O recursive relation between linked-list nodes.  
Each node in the list can have only of `next`. If a node A points (using `next`) to a node B,
B can get its pointer using `prev`.
   
### Generate Assets

```console
go generate ./...
```

### Run Example

```console
go test
```
