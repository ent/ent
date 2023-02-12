---
title: Visualizing your Data Graph Using entviz
author: Amit Shani
authorURL: "https://github.com/hedwigz"
authorImageURL: "https://avatars.githubusercontent.com/u/8277210?v=4"
authorTwitter: itsamitush
---

Joining an existing project with a large codebase can be a daunting task.  

Understanding the data model of an application is key for developers to start working on an existing project. One  commonly used tool to help overcome this challenge, and enable developers to grasp an application's data model is an [ER (Entity Relation) diagram](https://en.wikipedia.org/wiki/Entity%E2%80%93relationship_model).  

ER diagrams provide a visual representation of your data model, and details each field of the entities. Many tools can help create these, where one example is Jetbrains DataGrip, that can generate an ER diagram by connecting to and inspecting an existing database:

<div style={{textAlign: 'center'}}>
  <img alt="Datagrip ER diagram" src="https://entgo.io/images/assets/entviz/datagrip_er_diagram.png" />
  <p style={{fontSize: 12}}>DataGrip ER diagram example</p>
</div>

[Ent](https://entgo.io/docs/getting-started/), a simple, yet powerful entity framework for Go, was originally developed inside Facebook specifically for dealing with projects with large and complex data models.
This is why Ent uses code generation - it gives type-safety and code-completion out-of-the-box which helps explain the data model and improves developer velocity.
On top of all of this, wouldn't it be great to automatically generate ER diagrams that maintain a high-level view of the data model in a visually appealing representation? (I mean, who doesn't love visualizations?) 

### Introducing entviz
[entviz](https://github.com/hedwigz/entviz) is an ent extension that automatically generates a static HTML page that visualizes your data graph.

<div style={{textAlign: 'center'}}>
  <img width="600px" alt="Entviz example output" src="https://entgo.io/images/assets/entviz/entviz-example-visualization.png" />
  <p style={{fontSize: 12}}>Entviz example output</p>
</div>
Most ER diagram generation tools need to connect to your database and introspect it, which makes it harder to maintain an up-to-date diagram of the database schema. Since entviz integrates directly to your Ent schema, it does not need to connect to your database, and it automatically generates fresh visualization every time you modify your schema.

If you want to know more about how entviz was implemented, checkout the [implementation section](#implementation).

  
### See it in action
First, let's add the entviz extension to our entc.go file:
```bash
go get github.com/hedwigz/entviz
```
:::info
If you are not familiar with `entc` you're welcome to read [entc documentation](https://entgo.io/docs/code-gen#use-entc-as-a-package) to learn more about it.
:::
```go title="ent/entc.go"
import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{}, entc.Extensions(entviz.Extension{}))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
```
Let's say we have a simple schema with a user entity and some fields:
```go title="ent/schema/user.go"
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.Time("created").
			Default(time.Now),
	}
}
```
Now, entviz will automatically generate a visualization of our graph everytime we run: 
```bash
go generate ./...
```
You should now see a new file called `schema-viz.html` in your ent directory:
```bash
$ ll ./ent/schema-viz.html
-rw-r--r-- 1 hedwigz hedwigz 7.3K Aug 27 09:00 schema-viz.html
```
Open the html file with your favorite browser to see the visualization

![tutorial image](https://entgo.io/images/assets/entviz/entviz-tutorial-1.png)

Next, let's add another entity named Post, and see how our visualization changes:
```bash
ent new Post
```
```go title="ent/schema/post.go"
// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
		field.Time("created").
			Default(time.Now),
	}
}
```
Now we add an ([O2M](https://entgo.io/docs/schema-edges/#o2m-two-types)) edge from User to Post:
```go title="ent/schema/post.go"
// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
	}
}
```
Finally, regenerate the code:
```bash
go generate ./...
```
Refresh your browser to see the updated result!

![tutorial image 2](https://entgo.io/images/assets/entviz/entviz-tutorial-2.png)


### Implementation
Entviz was implemented by extending ent via its [extension API](https://github.com/ent/ent/blob/1304dc3d795b3ea2de7101c7ca745918def668ef/entc/entc.go#L197).
The Ent extension API lets you aggregate multiple [templates](https://entgo.io/docs/templates/), [hooks](https://entgo.io/docs/hooks/), [options](https://entgo.io/docs/code-gen/#code-generation-options) and [annotations](https://entgo.io/docs/templates/#annotations).
For instance, entviz uses templates to add another go file, `entviz.go`, which exposes the `ServeEntviz` method that can be used as an http handler, like so:
```go
func main() {
	http.ListenAndServe("localhost:3002", ent.ServeEntviz())
}
```
We define an extension struct which embeds the default extension, and we export our template via the `Templates` method:
```go
//go:embed entviz.go.tmpl
var tmplfile string
 
type Extension struct {
	entc.DefaultExtension
}
 
func (Extension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("entviz").Parse(tmplfile)),
	}
}
```
The template file is the code that we want to generate:
```gotemplate
{{ define "entviz"}}
 
{{ $pkg := base $.Config.Package }}
{{ template "header" $ }}
import (
	_ "embed"
	"net/http"
	"strings"
	"time"
)

//go:embed schema-viz.html
var html string

func ServeEntviz() http.Handler {
	generateTime := time.Now()
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		http.ServeContent(w, req, "schema-viz.html", generateTime, strings.NewReader(html))
	})
}
{{ end }}
```
That's it! now we have a new method in ent package.  

### Wrapping-Up

We saw how ER diagrams help developers keep track of their data model. Next, we introduced entviz - an Ent extension that automatically generates an ER diagram for Ent schemas. We saw how entviz utilizes Ent's extension API to extend the code generation and add extra functionality. Finally, you got to see it in action by installing and use entviz in your own project. If you like the code and/or want to contribute - feel free to checkout the [project on github](https://github.com/hedwigz/entviz).

Have questions? Need help with getting started? Feel free to join our [Discord server](https://discord.gg/qZmPgTE6RX) or [Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
