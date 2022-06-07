# User-Pets O2M Relation

An example for a O2M (one-to-many) relation between a user and its pets.  
Each user **has many** pets, and a pet **has one** owner. If a user A adds
a pet B using the pets edge, B can get its owner using the owner edge.


### Generate Assets

```console
go generate ./...
```

### Run Example

```console
go test
```
