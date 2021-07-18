# Recursive Traversal Using [CTE](https://en.wikipedia.org/wiki/Hierarchical_and_recursive_queries_in_SQL#Common_table_expression)

In this example, we create a file system with a tree structure, and want to query all "undeleted" files.
A file is considered as "deleted", if it's marked as "deleted" (a bool field), or any of its parents is
marked as "deleted".

Given the following tree structure:

```console
a/
├─ b/
│  ├─ ba
│  ├─ bb
│  └─ bc (deleted)
├─ c/ (deleted)
│  ├─ ca
│  └─ cb
└─ d (deleted)
```

Query "undeleted" files should return the following structure:

```console
a/
└─ b/
   ├─ ba
   └─ bb
```

As you can see, in order to check if "cb" (or "ca") is "deleted", we need to "look behind" recursively
until we find a "deleted" parent, or reach the root ("a").


### Generate Assets

```console
go generate ./...
```

### Run Example

```console
go test
```
