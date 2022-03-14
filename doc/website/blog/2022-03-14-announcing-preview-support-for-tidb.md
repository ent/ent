---
title: Announcing preview support for TiDB
author: Amit Shani
authorURL: "https://github.com/hedwigz"
authorImageURL: "https://avatars.githubusercontent.com/u/8277210?v=4"
authorTwitter: itsamitush
---

We [previously announced](https://entgo.io/blog/2022/01/20/announcing-new-migration-engine) the new migration engine - `Atlas`.
With `Atlas`'s new design, it became easier than ever to add support for new databases for Ent.
Today, I am happy to announce that a preview support for [TiDB](https://en.pingcap.com/tidb/) is now available, using the latest version of Ent with `Atlas` enabled.  
For a quick `Hello World` with `Ent`+`TiDB`, follow the following steps:  
1. Spin up a local TiDB server by using Docker:
 ```shell
 docker run -p 4000:4000 pingcap/tidb
 ```
 You should now have a running instance of TiDB listening on port 4000.

2. Clone the example [`hello world` repository](https://github.com/hedwigz/tidb-hello-world):
 ```shell
 git clone https://github.com/hedwigz/tidb-hello-world.git
 ```
 
 Connecting Ent with TiDB is easy, just peek at `main.go`
 ```go {1,7} title="main.go"
 client, err := ent.Open("mysql", "root@tcp(localhost:4000)/test?parseTime=true")
 if err != nil {
 	log.Fatalf("failed opening connection to sqlite: %v", err)
 }
 defer client.Close()
 // Run the auto migration tool, with Atlas.
 if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
 	log.Fatalf("failed printing schema changes: %v", err)
 }
 ```
 Note that in line `1` we connect to the TiDB server using a `mysql` dialect. This is due to the fact that TiDB is [MySQL compatible](https://docs.pingcap.com/tidb/stable/mysql-compatibility#:~:text=TiDB%20is%20highly%20compatible%20with,can%20be%20used%20for%20TiDB.), and it does not require any proprietary driver.  
 With that said, `Atlas` automatically detects when it is connected to `TiDB` and handles that accordingly.  
 In addition to that, note that in line `7` we used `schema.WithAtlas(true)`, which flags Ent to use `Atlas` as its 
 migration engine.  
  
You can now run the example program:
```go
go run main.go
```
and get the following output:
```shell
the user: hedwigz is 30 years old
```

For more in-depth content about this, and future databases support - subscribe to our [Newsletter](https://www.getrevue.co/profile/ariga)

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://www.getrevue.co/profile/ent)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
