---
title: Announcing preview support for TiDB
author: Amit Shani
authorURL: "https://github.com/hedwigz"
authorImageURL: "https://avatars.githubusercontent.com/u/8277210?v=4"
authorTwitter: itsamitush
---

Dear community,

We [previously announced](https://entgo.io/blog/2022/01/20/announcing-new-migration-engine) the new migration engine `Atlas`.
With `Atlas`'s new design, it became easier than ever to add support for new databases for Ent.
Today, I am happy to announce that a preview support for [TiDB](https://en.pingcap.com/tidb/) is now available, using `ent@master` and `Atlas` enabled.  
  
For a quick `Hello World` with `Ent`+`TiDB`, follow the following steps:  
1. Spin up a local TiDB server by using [`TiUP`](https://tiup.io/):
 ```bash
 curl --proto '=https' --tlsv1.2 -sSf https://tiup-mirrors.pingcap.com/install.sh | sh
 ```
 followed by:
 ```
 tiup playground v5.4.0
 ```
 You should now have a running instance of TiDB listening on port 4000.

2. Clone the example hello world repository:
 ```bash
 git clone github.com/hedwigz/tidb-hello-world
 ```
 
 Connecting Ent with TiDB is easy, just peek at `main.go`
 ```go {1,7}
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
 Note that in line 1 we connect to the TiDB server using a `mysql` dialect. This is due to the fact that TiDB is MySQL compatible, and it does not require any special driver.  
 With that said, `Atlas` automatically detects when it is connected to `TiDB` and handles that accordingly.  
 In addition to that, note that in line 7 we used `schema.WithAtlas(true)`, which flags Ent to use `Atlas` as its 
 migration engine.  
  
You can now run the example program:
```go
go run main.go
```
and get the following output:
```bash
the user: hedwigz is 30 years old
```


<!-- TODO: call for action to subscribe to ariga blog for more content about this announcement -->

Have questions? Need help with getting started? Feel free to [join our Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://www.getrevue.co/profile/ent)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
