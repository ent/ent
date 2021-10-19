---
title: sqlcomment support for ent
author: Amit Shani
authorURL: "https://github.com/hedwigz"
authorImageURL: "https://avatars.githubusercontent.com/u/8277210?v=4"
authorTwitter: itsamitush
---

Ent is a powerful Entity Framework that helps developers write neat code that is translated into (possibly complex) database queries. As the usage of your application grows, it doesn’t take long until you stumble upon performance issues with your database.
Troubleshooting database performance issues is notoriously hard, especially when you’re not equipped with the right tools.  

The following example shows how ent query code is translated into an SQL query.

<div style={{textAlign: 'center'}}>
  <img alt="ent example 1" src="https://entgo.io/images/assets/sqlcomment/pipeline.png" />
  <p style={{fontSize: 12}}>Example 1 - ent code is translated to SQL query</p>
</div>

Traditionally, it has been very difficult to correlate between poorly performing database queries and the application code that is generating them. Database performance analysis tools could help point out slow queries by analyzing database server logs, but how could they be traced back to the application?

### Sqlcommenter
Earlier this year, [Google introduced](https://cloud.google.com/blog/topics/developers-practitioners/introducing-sqlcommenter-open-source-orm-auto-instrumentation-library) Sqlcommenter. Sqlcommenter is 

> <em>an open source library that addresses the gap between the ORM libraries and understanding database performance. Sqlcommenter gives application developers visibility into which application code is generating slow queries and maps application traces to database query plans</em>

In other words, sqlcommenter adds application context metadata to SQL queries. This information can then be used to provide meaningful insights. It does so by adding [https://en.wikipedia.org/wiki/SQL_syntax#Comments](SQL comments) to the query that carry metadata but are ignored by the database during query execution. 
For example, the following query contains a comment that carries metadata about the application that issued it (`users-mgr`), which controller and route triggered it (`users` and `user_rename`, respectively), and the database driver that was used (`ent:v0.9.1`):

```SQL
update users set username = ‘hedwigz’ where id = 88
/*application='users-mgr',controller='users',route='user_rename',db_driver='ent:v0.9.1'*/
```

In the following example, we see Cloud SQL Insights Dashboard and we can see that the HTTP route “demo/charge” is causing many locks on the database. We can also see that this query got called ~500,000 times in the last hour.

<div style={{textAlign: 'center'}}>
  <img alt="Cloud SQL insights" src="https://storage.googleapis.com/gweb-cloudblog-publish/images/query_insights.max-1300x1300.png" />
  <p style={{fontSize: 12}}>illustration from <a href="https://cloud.google.com/blog/topics/developers-practitioners/introducing-sqlcommenter-open-source-orm-auto-instrumentation-library">Google's announcement</a></p>
</div>

This is the power of SQL tags - they provide you correlation between your application-level information and your Database monitors.

### sqlcomm**ent**
[sqlcomment](https://github.com/ariga/sqlcomment) is an ent driver that adds SQL tags following the [sqlcommenter specification](https://google.github.io/sqlcommenter/spec/).
Without further ado, let’s see sqlcomment in action.  
First, to install sqlcomment run:
```bash
go get ariga.io/sqlcomment
```

`sqlcomment` is wrapping an underlying SQL driver, therefore, we need to open our SQL connection using ent’s `sql` module, instead of ents popular helper `ent.Open`

```go
import (
	"ariga.io/sqlcomment"
	"entgo.io/ent/dialect/sql"
)

// Create db driver.
db, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
if err != nil {
	log.Fatalf("Failed to connect to database: %v", err)
}
// create sqlcomment driver which wraps sqlite driver.
drv := sqlcomment.NewDriver(db),
	sqlcomment.WithDriverVerTag(),
	sqlcomment.WithTags(sqlcomment.Tags{
		sqlcomment.KeyApplication: "my-app",
		sqlcomment.KeyFramework: "net/http",
	}),
)
// create and configure ent client
client := ent.NewClient(ent.Driver(drv))
```

Now, whenever we execute a query, `sqlcomment` will suffix our SQL query with the tags we set up.

![sqlcomment pipeline](https://entgo.io/images/assets/sqlcomment/pipeline2.png)

As you can see, ent outputted an SQL query with a comment at the end, containing all the relevant information associated with that query.  

For more advanced examples, please visit the [github repo](https://github.com/ariga/sqlcomment).

### Wrapping-Up

In this post we introduced the concept of SQL tags and we saw how they help correlate between source code and database queries. Next, we introduced `sqlcomment` - an Ent driver that adds SQL tags to all of your queries. Finally, we got to see `sqlcomment` in action, by installing and configuring it with Ent. If you like the code and/or want to contribute - feel free to checkout the [project on github](https://github.com/ariga/sqlcomment).

Have questions? Need help with getting started? Feel free to [join our Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://www.getrevue.co/profile/ent)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
