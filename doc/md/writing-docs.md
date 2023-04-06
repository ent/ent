---
id: writing-docs
title: Writing Docs
---

This document contains guidelines for contributing changes to the Ent documentation website.

The Ent documentation website is generated from the project's main [GitHub repo](https://github.com/ent/ent).

Follow this short guide to contribute documentation improvements and additions:

### Setting Up

1\. [Fork and clone locally](https://docs.github.com/en/github/getting-started-with-github/quickstart/fork-a-repo) the
[main repository](https://github.com/ent/ent).

2\. The documentation site uses [Docusaurus](https://docusaurus.io/). To run it you will need [Node.js installed](https://nodejs.org/en/).

3\. Install the dependencies:
```shell
cd doc/website && npm install
```

4\. Run the website in development mode:

```shell
cd doc/website && npm start
```

5\. Open you browser at [http://localhost:3000](http://localhost:3000).

### General Guidelines

* Documentation files are located in `doc/md`, they are [Markdown-formatted](https://en.wikipedia.org/wiki/Markdown)
  with "front-matter" style annotations at the top. [Read more](https://docusaurus.io/docs/docs-introduction) about
  Docusaurus's document format.
* Ent uses [Golang CommitMessage](https://github.com/golang/go/wiki/CommitMessage) formats to keep the repository's
  history nice and readable. As such, please use a commit message such as:
```text
doc/md: adding a guide on contribution of docs to ent
```

### Adding New Documents

1\. Add a new Markdown file in the `doc/md` directory, for example `doc/md/writing-docs.md`.

2\. The file should be formatted as such:

```markdown
---
id: writing-docs
title: Writing Docs
---
...
```
Where `id` should be a unique identifier for the document, should be the same as the filename without the `.md` suffix,
and `title` is the title of the document as it will appear in the page itself and any navigation element on the site.

3\. If you want the page to appear in the documentation website's sidebar, add its `id` to `website/sidebars.js`, for example:
```diff
{
      type: 'category',
      label: 'Misc',
      items: [
        'templates',
        'graphql',
        'sql-integration',
        'testing',
        'faq',
        'generating-ent-schemas',
        'feature-flags',
        'translations',
        'contributors',
+       'writing-docs',
        'slack'
      ],
      collapsed: false,
    },
```
