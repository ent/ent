---
id: dialects
title: Supported Dialects
---

## MySQL

MySQL supports all the features that are mentioned in the [Migration](migrate.md) section,
and it's being tested constantly on the following versions: `5.6.51`, `5.7.36` and `8.0.27`. 

## MariaDB

MariaDB supports all the features that are mentioned in the [Migration](migrate.md) section,
and it's being tested constantly on the following versions: `10.2.40`, `10.3.31`, `10.4.21`, `10.5.12` and `10.6.4`.

## PostgreSQL

PostgreSQL supports all the features that are mentioned in the [Migration](migrate.md) section,
and it's being tested constantly on the following versions: `10.18`, `11.13`, `12.8`, `13.4` and `14.0`.

## SQLite

SQLite supports all _"append-only"_ features mentioned in the [Migration](migrate.md) section. 
However, dropping or modifying resources, like [drop-index](migrate.md#drop-resources) are not
supported by default by SQLite, and will be added in the future using a [temporary table](https://www.sqlite.org/lang_altertable.html#otheralter).

## Gremlin

Gremlin does not support migration nor indexes, and **<ins>it's considered experimental</ins>**.
