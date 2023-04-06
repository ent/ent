---
title: Announcing preview support for TiDB
author: Amit Shani
authorURL: "https://github.com/hedwigz"
authorImageURL: "https://avatars.githubusercontent.com/u/8277210?v=4"
authorTwitter: itsamitush
---

We [previously announced](2022-01-20-announcing-new-migration-engine.md) Ent's new migration engine - Atlas.
Using Atlas, it has become easier than ever to add support for new databases to Ent.
Today, I am happy to announce that preview support for [TiDB](https://en.pingcap.com/tidb/) is now available, using the latest version of Ent with Atlas enabled.  
  
Ent can be used to access data in many types of databases, both graph-oriented and relational.  Most commonly, users have been using standard open-source relational databases such as MySQL, MariaDB, and PostgreSQL.  As teams building Ent-based applications become more successful and need to deal with  traffic on larger scales, these single-node databases often become the bottleneck for scaling out. For this reason, many members of the Ent community have requested support for [NewSQL](https://en.wikipedia.org/wiki/NewSQL) databases such as TiDB.

### TiDB
[TiDB](https://en.pingcap.com/tidb/) is an [open-source](https://github.com/pingcap/tidb) NewSQL database. It provides many features that traditional databases don't, such as:
1. **Horizontal scaling** - for many years software architects needed to choose between the familiarity and guarantees that relational databases provide and the scaling-out capability of _NoSQL_ databases (such as MongoDB or Cassandra). TiDB supports horizontal scaling while maintaining good compatibility with MySQL features. 
2. **HTAP (Hybrid transactional/analytical processing)** - In addition, databases are traditionally divided into analytical (OLAP) and transactional (OLTP) databases. TiDB breaks this dichotomy by enabling both analytics and transactional workloads on the same database. 
3. **Pre-packed monitoring** w/ Prometheus+Grafana - TiDB is built on Cloud-native paradigms from the ground up, and natively supports the standard CNCF observability stack. 
  
To read more about it, check out the official [TiDB Introduction](https://docs.pingcap.com/tidb/stable).

### Hello World with TiDB

For a quick "Hello World" application with Ent+TiDB, follow these steps:  
1. Spin up a local TiDB server by using Docker:
 ```shell
 docker run -p 4000:4000 pingcap/tidb
 ```
 Now you should have a running instance of TiDB listening on port 4000.

2. Clone the example [`hello world` repository](https://github.com/hedwigz/tidb-hello-world):
 ```shell
 git clone https://github.com/hedwigz/tidb-hello-world.git
 ```
 In this example repository we defined a simple schema `User`:
 ```go title="ent/schema/user.go"
 func (User) Fields() []ent.Field {
 	return []ent.Field{
 		field.Time("created_at").
 			Default(time.Now),
 		field.String("name"),
 		field.Int("age"),
 	}
 }
 ```
 Then, we connected Ent with TiDB:
 ```go title="main.go"
 client, err := ent.Open("mysql", "root@tcp(localhost:4000)/test?parseTime=true")
 if err != nil {
 	log.Fatalf("failed opening connection to tidb: %v", err)
 }
 defer client.Close()
 // Run the auto migration tool, with Atlas.
 if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
 	log.Fatalf("failed printing schema changes: %v", err)
 }
	```
 Note that in line `1` we connect to the TiDB server using a `mysql` dialect. This is possible due to the fact that TiDB is [MySQL compatible](https://docs.pingcap.com/tidb/stable/mysql-compatibility), and it does not require any special driver.  
 Having said that, there are some differences between TiDB and MySQL, especially when pertaining to schema migrations, such as information schema inspection and migration planning. For this reason, `Atlas` automatically detects if it is connected to `TiDB` and handles the migration accordingly.  
 In addition, note that in line `7` we used `schema.WithAtlas(true)`, which flags Ent to use `Atlas` as its 
 migration engine.  
   
 Finally, we create a user and save the record to TiDB to later be queried and printed.
 ```go title="main.go"
 client.User.Create().
		SetAge(30).
		SetName("hedwigz").
		SaveX(context.Background())
 user := client.User.Query().FirstX(context.Background())
 fmt.Printf("the user: %s is %d years old\n", user.Name, user.Age)
 ```
 3. Run the example program:
 ```go
 $ go run main.go
 the user: hedwigz is 30 years old
 ```

Woohoo! In this quick walk-through we managed to:
* Spin up a local instance of TiDB.
* Connect Ent with TiDB.
* Migrate our Ent schema with Atlas.
* Insert and query from TiDB using Ent.

### Preview support
The integration of Atlas with TiDB is well tested with TiDB version `v5.4.0` (at the time of writing, `latest`) and we will extend that in the future.
If you're using other versions of TiDB or looking for help, don't hesitate to [file an issue](https://github.com/ariga/atlas/issues) or join our [Discord channel](https://discord.gg/zZ6sWVg6NT).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
