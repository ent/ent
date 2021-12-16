---
title: Serverless GraphQL using with AWS and ent
author: Bodo Kaiser
authorURL: "https://github.com/bodokaiser"
authorImageURL: "https://avatars.githubusercontent.com/u/1780466?v=4"
---

[Graphql][1] is a query language for HTTP APIs, providing a statically-typed interface to conveniently represent today's complex data hierarchies.
One way to use GraphQL is to import a library implementing a GraphQL server to which one registers custom resolvers implementing the database interface.
An alternative way is to use a GraphQL cloud service to implement the GraphQL server and register serverless cloud functions as resolvers.
Among the many benefits of cloud services, one of the biggest practical advantages is the resolvers' independence and composability.
For example, we can write one resolver to a relational database and another to a search database.

We consider such a setup using [Amazon Web Services (AWS)][2] in the following. In particular, we use [AWS AppSync][3] as the GraphQL service and [AWS Lambda][4] to run a relational database resolver, which we implement using [Go][5] with [ent][6] as the entity framework.
Compared to Nodejs, the most popular runtime for AWS Lambda, Go offers faster start times, higher performance, and, in my opinion, an improved developer experience.
On the other hand, Ent presents an innovative approach towards type-safe access to relational databases, which in my opinion, is unmatched in the Go ecosystem.
In conclusion, running Ent with AWS Lambda as AWS AppSync resolvers is an extremely powerful setup to face today's demanding API requirements.

### Getting Started

### Deploying AWS Lambda

### Setting up AWS AppSync

### Wrapping Up

Have questions? Need help with getting started? Feel free to [join our Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://www.getrevue.co/profile/ent)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::

[1]: https://graphql.org
[2]: https://aws.amazon.com
[3]: https://aws.amazon.com/appsync/
[4]: https://aws.amazon.com/lambda/
[5]: https://go.dev
[6]: https://entgo.io
