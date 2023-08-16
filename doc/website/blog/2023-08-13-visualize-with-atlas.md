---
title: "Quickly Generate ERDs from your Ent Schemas (Updated)" 
author: Rotem Tamir
authorURL: "https://github.com/rotemtam"
authorImageURL: "https://s.gravatar.com/avatar/36b3739951a27d2e37251867b7d44b1a?s=80"
authorTwitter: _rtam
image: "https://atlasgo.io/uploads/ent/inspect/entviz.png"
---

### TL;DR

Create a visualization of your Ent schema with one command:

```
atlas schema inspect \
  -u ent://ent/schema \
  --dev-url "sqlite://demo?mode=memory&_fk=1" \
  --visualize
```

![](https://entgo.io/images/assets/erd/edges-quick-summary.png)


Hi Everyone!

A few months ago, we shared [entviz](/blog/2023/01/26/visualizing-with-entviz), a cool
tool that enables you to visualize your Ent schemas. Due to its success and popularity,
we decided to integrate it directly into [Atlas](https://atlasgo.io), the migration engine
that Ent uses. 

Since the release of [v0.13.0](https://atlasgo.io/blog/2023/08/06/atlas-v-0-13) of Atlas,
you can now visualize your Ent schemas directly from Atlas without needing to install an
additional tool.

### Private vs. Public Visualizations

Previously, you could only share a visualization of your schema to the
[Atlas Public Playground](https://gh.atlasgo.cloud/explore). While this is convenient
for sharing your schema with others, it is not acceptable for many teams who maintain
schemas that themselves are sensitive and cannot be shared publicly.

With this new release, you can easily publish your schema directly to your private
workspace on [Atlas Cloud](https://atlasgo.cloud). This means that only you and your
team can access the visualization of your schema.

### Visualizing your Ent Schema with Atlas

To visualize your Ent schema with Atlas, first install its latest version:

```
curl -sSfL https://atlasgo.io/install.sh | sh
```
For other installation options, see the [Atlas installation docs](https://atlasgo.io/getting-started#installation).

Next, run the following command to generate a visualization of your Ent schema:

```
atlas schema inspect \
  -u ent://ent/schema \
  --dev-url "sqlite://demo?mode=memory&_fk=1" \
  --visualize
```

Let's break this command down:
* `atlas schema inspect` - this command can be used to inspect schemas from a variety of sources and outputs
  them in various formats. In this case, we are using it to inspect an Ent schema.
* `-u ent://ent/schema` - this is the URL to the Ent schema we want to inspect. In this case, we are using the
  `ent://` schema loader to point to a local Ent schema in the `./ent/schema` directory.
* `--dev-url "sqlite://demo?mode=memory&_fk=1"` - Atlas relies on having an empty database called the
  [Dev Database](https://atlasgo.io/concepts/dev-database) to normalize schemas and make various calculations.
In this case, we are using an in memory SQLite database; but, if you are using a different driver, you can use
  `docker://mysql/8/dev` (for MySQL) or `docker://postgres/15/?search_path=public` (for PostgreSQL).

Once you run this command, you should see the following output:

```text
Use the arrow keys to navigate: ↓ ↑ → ←
? Where would you like to share your schema visualization?:
  ▸ Publicly (gh.atlasgo.cloud)
    Your personal workspace (requires 'atlas login')
```

If you want to share your schema publicly, you can select the first option. If you want to share it privately, you
can select the second option and then run `atlas login` to log in to your (free) Atlas account.

### Wrapping up

In this post, we showed how you can easily visualize your Ent schema with Atlas. We hope you find this feature useful
and we look forward to hearing your feedback!


:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
