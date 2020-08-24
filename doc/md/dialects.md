---
id: dialects
title: Supported Dialects
---

## MySQL

MySQL supports all the features that are mentioned in the [Migration](migrate.md) section,
and it's being tested constantly on the following 3 versions: `5.6.35`, `5.7.26` and `8`. 


## PostgreSQL

PostgreSQL supports all the features that are mentioned in the [Migration](migrate.md) section,
and it's being tested constantly on the following 3 versions: `10`, `11` and `12`. 

## SQLite

SQLite supports all _"append-only"_ features mentioned in the [Migration](migrate.md) section. 
However, dropping or modifying resources, like [drop-index](migrate.md#drop-resources) are not
supported by default by SQLite, and will be added in the future using a [temporary table](https://www.sqlite.org/lang_altertable.html#otheralter).

## Gremlin

Gremlin does not support migration nor indexes, and **<ins>it's considered experimental</ins>**.
