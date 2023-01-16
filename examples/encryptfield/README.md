# Encrypted field example using go.dev

### Setup 

Add the `secrets.Keeper` as a dependency to your project and enable the `intercept` feature flag.

```diff
func main() {
+	opts := []entc.Option{
+		entc.Dependency(
+			entc.DependencyType(&secrets.Keeper{}),
+		),
+		entc.FeatureNames("intercept"),
+	}
	if err := entc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
```

### Generate Assets

```console
go generate ./...
```

### Update the schema with secret field.

See `ent/schema/user.go` for full example.

### Run Example

```console
go test
```
