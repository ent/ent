# User-Friends Bidirectional M2M Relation

In this user-friends example, we have a **symmetric M2M relation** named `friends`.
Each user can **have many** friends. If user A becomes a friend of B, B is also a friend of A.

### Generate Assets

```console
go generate ./...
```

### Run Example

```console
go test
```
