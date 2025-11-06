# Contributing to ent
We want to make contributing to this project as easy and transparent as
possible.

# Project structure

- `dialect` - Contains SQL and Gremlin code used by the generated code.
  - `dialect/sql/schema` - Auto migration logic resides there.
  - `dialect/sql/sqljson` - JSON extension for SQL.

- `schema` - User schema API.
  - `schema/{field, edge, index, mixin}` - provides schema builders API.
  - `schema/field/gen` - Templates and codegen for numeric builders.

- `entc` - Codegen of `ent`.
  - `entc/load` - `entc` loader API for loading user schemas into a Go objects at runtime.
  - `entc/gen` - The actual code generation logic resides in this package (and its `templates` package).
  - `integration` - Integration tests for `entc`.

- `privacy` - Runtime code for [privacy layer](https://entgo.io/docs/privacy/).

- `doc` - Documentation code for `entgo.io` (uses [Docusaurus](https://docusaurus.io)).
  - `doc/md` - Markdown files for documentation.
  - `doc/website` - Website code and assets.

  In order to test your documentation changes, run `npm start` from the `doc/website` directory, and open [localhost:3000](http://localhost:3000/).

# Run integration tests
If you touch any file in `entc`, run the following commands in `entc/integration` and 'examples' dirs:

```
go generate ./...
go mod tidy
```

Then, in `entc/integration` run `docker-compose` in order to spin-up all database containers:

```
docker-compose -f docker-compose.yaml up -d
```

Then, run `go test ./...` to run all integration tests.


## Pull Requests
We actively welcome your pull requests.

1. Fork the repo and create your branch from `master`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. If you haven't already, complete the Contributor License Agreement ("CLA").

## License
By contributing to ent, you agree that your contributions will be licensed
under the LICENSE file in the root directory of this source tree.
