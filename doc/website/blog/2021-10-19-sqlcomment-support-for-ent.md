---
title: Introducing sqlcomment - Database Performance Analysis with Ent and Google's Sqlcommenter
author: Amit Shani
authorURL: "https://github.com/hedwigz"
authorImageURL: "https://avatars.githubusercontent.com/u/8277210?v=4"
authorTwitter: itsamitush
image: https://entgo.io/images/assets/sqlcomment/share.png
---

Ent is a powerful Entity framework that helps developers write neat code that is translated into (possibly complex) database queries. As the usage of your application grows, it doesn’t take long until you stumble upon performance issues with your database.
Troubleshooting database performance issues is notoriously hard, especially when you’re not equipped with the right tools.  

The following example shows how Ent query code is translated into an SQL query.

<div style={{textAlign: 'center'}}>
  <img alt="ent example 1" src="https://entgo.io/images/assets/sqlcomment/pipeline.png" />
  <p style={{fontSize: 12}}>Example 1 - ent code is translated to SQL query</p>
</div>

Traditionally, it has been very difficult to correlate between poorly performing database queries and the application code that is generating them. Database performance analysis tools could help point out slow queries by analyzing database server logs, but how could they be traced back to the application?

### Sqlcommenter
Earlier this year, [Google introduced](https://cloud.google.com/blog/topics/developers-practitioners/introducing-sqlcommenter-open-source-orm-auto-instrumentation-library) Sqlcommenter. Sqlcommenter is 

> <em>an open source library that addresses the gap between the ORM libraries and understanding database performance. Sqlcommenter gives application developers visibility into which application code is generating slow queries and maps application traces to database query plans</em>

In other words, Sqlcommenter adds application context metadata to SQL queries. This information can then be used to provide meaningful insights. It does so by adding [SQL comments](https://en.wikipedia.org/wiki/SQL_syntax#Comments) to the query that carry metadata but are ignored by the database during query execution. 
For example, the following query contains a comment that carries metadata about the application that issued it (`users-mgr`), which controller and route triggered it (`users` and `user_rename`, respectively), and the database driver that was used (`ent:v0.9.1`):

```sql
update users set username = ‘hedwigz’ where id = 88
/*application='users-mgr',controller='users',route='user_rename',db_driver='ent:v0.9.1'*/
```

To get a taste of how the analysis of metadata collected from Sqlcommenter metadata can help us better understand performance issues of our application, consider the following example: Google Cloud recently launched [Cloud SQL Insights](https://cloud.google.com/blog/products/databases/get-ahead-of-database-performance-issues-with-cloud-sql-insights), a cloud-based SQL performance analysis product.  In the image below, we see a screenshot from the Cloud SQL Insights Dashboard that shows that the HTTP route 'api/users' is causing many locks on the database. We can also see that this query got called 16,067 times in the last 6 hours.

<div style={{textAlign: 'center'}}>
  <img alt="Cloud SQL insights" src="https://entgo.io/images/assets/sqlcomment/ginsights.png" />
  <p style={{fontSize: 12}}>Screenshot from Cloud SQL Insights Dashboard</p>
</div>

This is the power of SQL tags - they provide you correlation between your application-level information and your Database monitors.

### sqlcomment

[sqlcomm**ent**](https://github.com/ariga/sqlcomment) is an Ent driver that adds metadata to SQL queries using comments following the [sqlcommenter specification](https://google.github.io/sqlcommenter/spec/). By wrapping an existing Ent driver with `sqlcomment`,  users can leverage any tool that supports the standard to triage query performance issues.
Without further ado, let’s see `sqlcomment` in action.  
  
First, to install sqlcomment run:
```bash
go get ariga.io/sqlcomment
```

`sqlcomment` is wrapping an underlying SQL driver, therefore, we need to open our SQL connection using ent’s `sql` module, instead of Ent's popular helper `ent.Open`.

:::info
Make sure to import `entgo.io/ent/dialect/sql` in the following snippet
:::

```go
// Create db driver.
db, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
if err != nil {
	log.Fatalf("Failed to connect to database: %v", err)
}

// Create sqlcomment driver which wraps sqlite driver.
drv := sqlcomment.NewDriver(db,
	sqlcomment.WithDriverVerTag(),
	sqlcomment.WithTags(sqlcomment.Tags{
		sqlcomment.KeyApplication: "my-app",
		sqlcomment.KeyFramework:   "net/http",
	}),
)

// Create and configure ent client.
client := ent.NewClient(ent.Driver(drv))
```

Now, whenever we execute a query, `sqlcomment` will suffix our SQL query with the tags we set up. If we were to run the following query:

```go
client.User.
	Update().
	Where(
		user.Or(
			user.AgeGT(30),
			user.Name("bar"),
		),
		user.HasFollowers(),
	).
	SetName("foo").
	Save()
```

Ent would output the following commented SQL query:

```sql
UPDATE `users`
SET `name` = ?
WHERE (
    `users`.`age` > ?
    OR `users`.`name` = ?
  )
  AND `users`.`id` IN (
    SELECT `user_following`.`follower_id`
    FROM `user_following`
  )
  /*application='my-app',db_driver='ent:v0.9.1',framework='net%2Fhttp'*/
```

As you can see, Ent outputted an SQL query with a comment at the end, containing all the relevant information associated with that query.  

sqlcomm**ent** supports more tags, and has integrations with [OpenTelemetry](https://opentelemetry.io) and [OpenCensus](https://opencensus.io).
To see more examples and scenarios, please visit the [github repo](https://github.com/ariga/sqlcomment).

### Wrapping-Up

In this post I showed how adding metadata to queries using SQL comments can help correlate between source code and database queries. Next, I introduced `sqlcomment` - an Ent driver that adds SQL tags to all of your queries. Finally, I got to see `sqlcomment` in action, by installing and configuring it with Ent. If you like the code and/or want to contribute - feel free to checkout the [project on GitHub](https://github.com/ariga/sqlcomment).

Have questions? Need help with getting started? Feel free to join our [Discord server](https://discord.gg/qZmPgTE6RX) or [Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
