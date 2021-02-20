---
id: dialects
title: Supported Dialects
---

## MySQL

MySQL支持所有的在[Migration](migrate.md)章节涉及到的特性。
它不断地测试以下3个版本`5.6.35`, `5.7.26` 和 `8`。

## PostgreSQL

PostgreSQL支持所有的在[Migration](migrate.md)章节涉及到的特性。
它不断地测试以下3个版本`10`, `11` 和 `12`。

## SQLite

SQLite支持所有的在[Migration](migrate.md)章节涉及到的 _"append-only"_ 特性。
然而，删除或者修改资源，像[drop-index](migrate.md#drop-resources)默认不支持SQLite，
并将在以后的使用 [temporary table](https://www.sqlite.org/lang_altertable.html#otheralter)中添加。

## Gremlin

Gremlin不支持迁移和索引。**<ins>还在实验中</ins>**.