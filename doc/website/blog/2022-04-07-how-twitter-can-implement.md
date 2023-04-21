---
title: How to implement the Twitter edit button with Ent
author: Amit Shani
authorURL: "https://github.com/hedwigz"
authorImageURL: "https://avatars.githubusercontent.com/u/8277210?v=4"
authorTwitter: itsamitush
image: "https://entgo.io/images/assets/enthistory/share.png"
---

Twitter's "Edit Button" feature has reached the headlines with Elon Musk's poll tweet asking whether users want the feature or not.

[![Elons Tweet](https://entgo.io/images/assets/enthistory/enthistory2.webp)](https://twitter.com/elonmusk/status/1511143607385874434)

Without a doubt, this is one of Twitter's most requested features.

As a software developer, I immediately began to think about how I would implement this myself. The tracking/auditing problem is very common in many applications. If you have an entity (say, a `Tweet`) and you want to track changes to one of its fields (say, the `content` field), there are many common solutions. Some databases even have proprietary solutions like Microsoft's change tracking and MariaDB's System Versioned Tables. However, in most use-cases you'd have to "stitch" it yourself. Luckily, Ent provides a modular extensions system that lets you plug in features like this with just a few lines of code.

![Twitter+Edit Button](https://entgo.io/images/assets/enthistory/enthistory3.gif)

<div style={{textAlign: 'center'}}>
  <p style={{fontSize: 12}}>if only</p>
</div>

### Introduction to Ent
Ent is an Entity framework for Go that makes developing large applications a breeze. Ent comes pre-packed with awesome features out of the box, such as:
* Type-safe generated [CRUD API](https://entgo.io/docs/crud)
* Complex [Graph traversals](https://entgo.io/docs/traversals) (SQL joins made easy)
* [Paging](https://entgo.io/docs/paging)
* [Privacy](https://entgo.io/docs/privacy)
* Safe DB [migrations](https://entgo.io/blog/2022/03/14/announcing-versioned-migrations).
  
With Ent's code generation engine and advanced [extensions system](https://entgo.io/blog/2021/09/02/ent-extension-api/), you can easily modularize your Ent's client with advanced features that are usually time-consuming to implement manually. For example:
* Generate [REST](https://entgo.io/blog/2022/02/15/generate-rest-crud-with-ent-and-ogen), [gRPC](https://entgo.io/docs/grpc-intro), and [GraphQL](https://entgo.io/docs/graphql) server.
* [Caching](http://entgo.io/blog/2021/10/14/introducing-entcache)
* Monitoring with [sqlcommenter](https://entgo.io/blog/2021/10/19/sqlcomment-support-for-ent)

### Enthistory
`enthistory` is an extension that we started developing when we wanted to add an "Activity & History" panel to one of our web services. The panel's role is to show who changed what and when (aka auditing). In [Atlas](https://atlasgo.io/), a tool for managing databases using declarative HCL files, we have an entity called "schema" which is essentially a large text blob. Any change to the schema is logged and can later be viewed in the "Activity & History" panel.

![Activity and History](https://entgo.io/images/assets/enthistory/enthistory1.gif)

<div style={{textAlign: 'center'}}>
  <p style={{fontSize: 12}}>The "Activity & History" screen in Atlas</p>
</div>

This feature is very common and can be found in many apps, such as Google docs, GitHub PRs, and Facebook posts, but is unfortunately missing in the very popular and beloved Twitter.

Over 3 million people voted in favor of adding the "edit button" to Twitter, so let me show you how Twitter can make their users happy without breaking a sweat!

With Enthistory, all you have to do is simply annotate your Ent schema like so:

```go
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").
			Annotations(enthistory.TrackField()),
		field.Time("created").
			Default(time.Now),
	}
}
```

Enthistory hooks into your Ent client to ensure that every CRUD operation to "Tweet" is recorded into the "tweets_history" table, with no code modifications and provides an API to consume these records:

```go
// Creating a new Tweet doesn't change. enthistory automatically modifies
// your transaction on the fly to record this event in the history table
client.Tweet.Create().SetContent("hello world!").SaveX(ctx)

// Querying history changes is as easy as querying any other entity's edge.
t, _ := client.Tweet.Get(ctx, id)
hs := client.Tweet.QueryHistory(t).WithChanges().AllX(ctx)
```

Let's see what you'd have to do if you weren't using Enthistory: For example, consider an app similar to Twitter. It has a table called "tweets" and one of its columns is the tweet content.

| id      | content | created_at | author_id |
| ----------- | ----------- | ----------- | ----------- |
| 1      | Hello Twitter!       | 2022-04-06T13:45:34+00:00       | 123       |
| 2      | Hello Gophers!       | 2022-04-06T14:03:54+00:00       | 456       |

Now, assume that we want to allow users to edit the content, and simultaneously display the changes in the frontend. There are several common approaches for solving this problem, each with its own pros and cons, but we will dive into those in another technical post. For now, a possible solution for this is to create a table "tweets_history" which records the changes of a tweet:

| id      | tweet_id | timestamp | event | content |
| ----------- | ----------- | ----------- | ----------- | ----------- |
| 1      | 1       | 2022-04-06T12:30:00+00:00       | CREATED       | hello world!       |
| 2      | 2       | 2022-04-06T13:45:34+00:00       | UPDATED       | hello Twitter!       |

With a table similar to the one above, we can record changes to the original tweet "1" and if requested, we can show that it was originally tweeted at 12:30:00 with the content "hello world!" and was modified at 13:45:34 to "hello Twitter!".  

To implement this, we will have to change every `UPDATE` statement for "tweets" to include an `INSERT` to "tweets_history". For correctness, we will need to wrap both statements in a transaction to avoid corrupting the history. in case the first statement succeeds but the subsequent one fails. We'd also need to make sure every `INSERT` to "tweets" is coupled with an `INSERT` to "tweets_history"

```diff
# INSERT is logged as "CREATE" history event
- INSERT INTO tweets (`content`) VALUES ('Hello World!');
+BEGIN;
+INSERT INTO tweets (`content`) VALUES ('Hello World!');
+INSERT INTO tweets_history (`content`, `timestamp`, `record_id`, `event`)
+VALUES ('Hello World!', NOW(), 1, 'CREATE');
+COMMIT;

# UPDATE is logged as "UPDATE" history event
- UPDATE tweets SET `content` = 'Hello World!' WHERE id = 1;
+BEGIN;
+UPDATE tweets SET `content` = 'Hello World!' WHERE id = 1;
+INSERT INTO tweets_history (`content`, `timestamp`, `record_id`, `event`)
+VALUES ('Hello World!', NOW(), 1, 'UPDATE');
+COMMIT;
```

This method is nice but you'd have to create another table for different entities ("comment_history", "settings_history"). To prevent that, Enthistory creates a single "history" and a single "changes" table and records all the tracked fields there. It also supports many type of fields without needing to add more columns.

### Pre release
Enthistory is still in early design stages and is being internally tested. Therefore, we haven't released it to open-source yet, though we plan to do so very soon.
If you want to play with a pre-release version of Enthistory, I wrote a simple React application with GraphQL+Enthistory to demonstrate how a tweet edit could look like. You can check it out [here](https://github.com/hedwigz/edit-twitter-example-app). Please feel free to share your feedback.

### Wrapping up
We saw how Ent's modular extension system lets you streamline advanced features as if they were just a package install away. Developing your own extension [is fun, easy and educating](https://entgo.io/blog/2021/12/09/contributing-my-first-feature-to-ent-grpc-plugin)! I invite you to try it yourself!
In the future, Enthistory will be used to track changes to Edges (aka foreign-keyed tables), integrate with OpenAPI and GraphQL extensions, and provide more methods for its underlying implementation.

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
