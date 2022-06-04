# [Optimistic Lock](https://en.wikipedia.org/wiki/Optimistic_concurrency_control)

In this example, we implement an optimistic locking mechanism using the technique mentioned in 
[Ent Blog](https://entgo.io/blog/2021/07/22/database-locking-techniques-with-ent/).

The idea is to add to our schema a `version` field that holds the Unix time of when the latest update occurred.
When an `Update` operation is executed, the hook updates the `version` field with the new value and adds a predicate
to verify that the `version` wasn't updated by another process/transaction during the mutation.

An error is returned if the versions are mismatched, and the user should reload the entity and retry the mutation.

### Generate Assets

```console
go generate ./...
```

### Run Example

```console
go test
```
