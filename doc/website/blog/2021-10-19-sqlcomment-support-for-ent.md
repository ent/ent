---
title: sqlcomment support for ent
author: Amit Shani
authorURL: "https://github.com/hedwigz"
authorImageURL: "https://avatars.githubusercontent.com/u/8277210?v=4"
authorTwitter: itsamitush
---

Ent is a powerful Entity Framework that helps developers write neat code that is translated into (possibly complex) database queries. As the usage of your application grows, it doesn’t take long until you stumble upon performance issues with your database.
Troubleshooting database performance issues is notoriously hard, especially when you’re not using the right tools/services.
The following example shows how ent query code is translated into an SQL query.

<div style={{textAlign: 'center'}}>
  <img alt="ent example 1" src="https://entgo.io/images/assets/entviz/datagrip_er_diagram.png" />
  <p style={{fontSize: 12}}>Example 1 - ent code is translated to SQL query</p>
</div>

### Sqlcommenter
Earlier this year, [Google introduced](https://cloud.google.com/blog/topics/developers-practitioners/introducing-sqlcommenter-open-source-orm-auto-instrumentation-library) Sqlcommenter. Sqlcommenter is 

> <em>an open source library that addresses the gap between the ORM libraries and understanding database performance. Sqlcommenter gives application developers visibility into which application code is generating slow queries and maps application traces to database query plans</em>

In the following example, we see Cloud SQL Insights Dashboard and we can see that the HTTP route “demo/charge” is causing many locks on the database. We can also see that this query got called ~500,000 times in the last hour.

<div style={{textAlign: 'center'}}>
  <img alt="Cloud SQL insights" src="https://entgo.io/images/assets/entviz/datagrip_er_diagram.png" />
  <p style={{fontSize: 12}}>illustration for <a href="https://entgo.io/images/assets/entviz/datagrip_er_diagram.png">Google's announcement</a></p>
</div>

This is the power of SQL tags - they provide you correlation between your application-level information and your Database monitors.

### sqlcomm**ent**
[sqlcomment](https://github.com/ariga/sqlcomment) is an ent driver that adds SQL tags following the [sqlcommenter specification](https://google.github.io/sqlcommenter/spec/). It also supports OpenTelemetry and OpenCensus.  

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

![sqlcomment pipeline](https://entgo.io/images/assets/entviz/entviz-tutorial-1.png)

### Wrapping-Up

In this post we introduced the concept of SQL tags and we saw how they help correlate between source code and database queries. Next, we introduced `sqlcomment` - an Ent driver that adds SQL tags to all of your queries. Finally, we got to see `sqlcomment` in action, by installing and configuring it with Ent. If you like the code and/or want to contribute - feel free to checkout the [project on github](https://github.com/ariga/sqlcomment).

Have questions? Need help with getting started? Feel free to [join our Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://www.getrevue.co/profile/ent)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
