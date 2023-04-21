# JSON Encode Extension

`EncodeExtension` is an implementation of entc.Extension that adds a `MarshalJSON`
method to each generated type `<T>` and inlines the Edges field to the top level JSON.

### Generate Assets

```console
go generate ./...
```

### Run Examples

```console
go test
```
