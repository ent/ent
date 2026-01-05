---
id: dialects
title: Supported Dialects
---

## MySQL

MySQL supports all the features that are mentioned in the [Migration](migrate.md) section,
and it's being tested constantly on the following 3 versions: `5.6.35`, `5.7.26` and `8`. 

## MariaDB

MariaDB supports all the features that are mentioned in the [Migration](migrate.md) section,
and it's being tested constantly on the following 3 versions: `10.2`, `10.3` and latest version.

## PostgreSQL

PostgreSQL supports all the features that are mentioned in the [Migration](migrate.md) section,
and it's being tested constantly on the following 8 versions: `10`, `11`, `12`, `13`, `14`, `15`, `16`, and `17`.

## CockroachDB **(<ins>preview</ins>)**

CockroachDB support is in preview and requires the [Atlas migration engine](migrate.md#atlas-integration).  
The integration with CRDB is currently tested on versions `v21.2.11`.

## SQLite

Using [Atlas](https://github.com/ariga/atlas), the SQLite driver supports all the features that
are mentioned in the [Migration](migrate.md) section. Note that some changes, like column modification,
are performed on a temporary table using the sequence of operations described in [SQLite official documentation](https://www.sqlite.org/lang_altertable.html#otheralter).

## Gremlin

Gremlin does not support migration nor indexes, and **<ins>it's considered experimental</ins>**.

## TiDB **(<ins>preview</ins>)**

TiDB support is in preview and requires the [Atlas migration engine](migrate.md#atlas-integration).  
TiDB is MySQL compatible and thus any feature that works on MySQL _should_ work on TiDB as well.  
For a list of known compatibility issues, visit: https://docs.pingcap.com/tidb/stable/mysql-compatibility  
The integration with TiDB is currently tested on versions `5.4.0`, `6.0.0`.
