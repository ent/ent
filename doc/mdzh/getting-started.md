---
id: getting-started
title: Quick Introduction
sidebar_label: Quick Introduction
---

**ent**是一个简单而又功能强大的Go语言实体框架，ent易于构建和维护应用程序与大数据模型。它严格遵循以下原则

- 图就是代码 - 将任何数据库表建模为Go对象。
- 轻松地遍历任何图形 - 可以轻松地运行查询、聚合和遍历任何图形结构。
- 静态类型和显式API - 使用代码生成静态类型和显式API，查询数据更加便捷。
- 多存储驱动程序 - 支持MySQL, PostgreSQL, SQLite 和 Gremlin。
- 可扩展 - 简单地扩展和使用Go模板自定义。


<br/>

![gopher-schema-as-code](https://s3.eu-central-1.amazonaws.com/entgo.io/assets/gopher-schema-as-code.png)

## 下载

```console
go get entgo.io/ent/cmd/ent
```

在你下载`ent`代码生成工具后，你应该有你的`PATH`。
如果找不到你的的路径，你也可以运行`go run entgo.io/ent/cmd/ent <command>`

## 设置一个Go语言环境

如果你的项目在[GOPATH](https://github.com/golang/go/wiki/GOPATH)的外部，或者你没有相似的GOPATH，设置一个如下的[Go module](https://github.com/golang/go/wiki/Modules#quick-start)
项目。

```console
go mod init <project>
```

## 创建你的第一个Schema

去项目根目录执行如下命令：

```console
ent init User
```
命令将会在`<project>/ent/schema/` 路径下生成`User`的schema

```go
// <project>/ent/schema/user.go

package schema

import "entgo.io/ent"

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

```

添加2个字段到`User` schema:

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)


// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}
```

在项目根目录下运行`go generate`命令：

```go
go generate ./ent
```

运行后将会产生如下文件

```
ent
├── client.go
├── config.go
├── context.go
├── ent.go
├── migrate
│   ├── migrate.go
│   └── schema.go
├── predicate
│   └── predicate.go
├── schema
│   └── user.go
├── tx.go
├── user
│   ├── user.go
│   └── where.go
├── user.go
├── user_create.go
├── user_delete.go
├── user_query.go
└── user_update.go
```


## 创建你的第一个实体

首先，创建一个新的`ent.Client`，如下所示，我们将使用SQLite3。

```go
package main

import (
	"context"
	"log"

	"<project>/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
```

现在，已经创建了我们的user，调用这个`CreateUser`函数:

```go
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
```

## 查询你的实体

`ent`对于每个实体都生成包含它断言、校验、关于实体存储的附加信息的包（列名，主键等等）

```go
package main

import (
	"log"

	"<project>/ent"
	"<project>/ent/user"
)

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.NameEQ("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

```


## 添加你的第一个关系
在这部分的内容中，我们想要声明一个在同一schema中和其他实体的关系。
让我们创建2个名为`Car` 和 `Group`的有一些字段的实体。我们使用`ent`命令生成初始化的schema: 

```console
go run entgo.io/ent/cmd/ent init Car Group
```

然后手动添加其余字段：

```go
import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Time("registered_at"),
	}
}


// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			// Regexp validation for group name.
			Match(regexp.MustCompile("[a-zA-Z_]+$")),
	}
}
```

让我们定义第一个关系，从`User`到`Car`的关系，1个user有1个或者更多的cars，但是一个car仅仅只属于一个所属者。(1对多)

![er-user-cars](https://s3.eu-central-1.amazonaws.com/entgo.io/assets/re_user_cars.png)

让我们添加`"cars"`的关系到 `User`，然后运行`go generate ./ent`:

 ```go
 import (
 	"log"

 	"entgo.io/ent"
 	"entgo.io/ent/schema/edge"
 )

 // Edges of the User.
 func (User) Edges() []ent.Edge {
 	return []ent.Edge{
		edge.To("cars", Car.Type),
 	}
 }
 ```

我们继续通过2个cars添加到所属的user的例子来学习。

```go
func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// Create a new car with model "Tesla".
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}

	// Create a new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}
	log.Println("car was created: ", ford)

	// Create a new user, and add it the 2 cars.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}
```

但是有关`cars`关系的查询怎么做？这里展示怎么做：

```go
import (
	"log"

	"<project>/ent"
	"<project>/ent/car"
)

func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println("returned cars:", cars)

	// What about filtering specific cars.
	ford, err := a8m.QueryCars().
		Where(car.ModelEQ("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println(ford)
	return nil
}
```

## 添加第一个反向关系(BackRef)

假设我们有一个`Car`对象，我们想得到它的所属者，对于这些，我们有另一种类型叫做"inverse edge"，这种类型使用`edge.From`函数被定义。

![er-cars-owner](https://s3.eu-central-1.amazonaws.com/entgo.io/assets/re_cars_owner.png)

在上图中创建的新边是半透明的，它强调我们不在数据库中创建其他关系。仅仅是对于真实关系的反向。

添加一个`owner` 到 `Car`的反向关系，参考它在`User` schema中到`cars`的关系。然后运行`go generate ./ent`:

```go
import (
	"log"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
	 	// and reference it to the "cars" edge (in User schema)
	 	// explicitly using the `Ref` method.
	 	edge.From("owner", User.Type).
	 		Ref("cars").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}
```

我们将继续通过查询user/cars的反向关系：

```go
import (
	"log"

	"<project>/ent"
)

func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	// Query the inverse edge.
	for _, ca := range cars {
		owner, err := ca.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %v", ca.Model, err)
		}
		log.Printf("car %q owner: %q\n", ca.Model, owner.Name)
	}
	return nil
}
```

## 创建第二个关系

我们将继续我的案例，通过创建users和groups之间多对多的关系来说明。

![er-group-users](https://s3.eu-central-1.amazonaws.com/entgo.io/assets/re_group_users.png)

如你所见，每个group实体**有多个**users，一个user能**被多个group关联**;
一个简单的“多对多”的关系，在上面的图示中，`Group` schema是`users`的拥有者。`User`实体对于`groups`有一个反向的关系，
让我们一起定义这个关系吧：

- `<project>/ent/schema/group.go`:

	```go
	 import (
		"log"
	
		"entgo.io/ent"
		"entgo.io/ent/schema/edge"
	 )
	
	 // Edges of the Group.
	 func (Group) Edges() []ent.Edge {
		return []ent.Edge{
			edge.To("users", User.Type),
		}
	 }
	```

- `<project>/ent/schema/user.go`:   
	```go
	 import (
	 	"log"
	
	 	"entgo.io/ent"
	 	"entgo.io/ent/schema/edge"
	 )
	
	 // Edges of the User.
	 func (User) Edges() []ent.Edge {
	 	return []ent.Edge{
			edge.To("cars", Car.Type),
		 	// Create an inverse-edge called "groups" of type `Group`
		 	// and reference it to the "users" edge (in Group schema)
		 	// explicitly using the `Ref` method.
			edge.From("groups", Group.Type).
				Ref("users"),
	 	}
	 }
	```

在schema目录下运行`ent`重新生成断言。
```console
go generate ./ent
```

## 运行第一次图遍历

为了运行第一次图遍历，我们需要生成一些数据（node和关系，换句话说就是实体和关系），让我们创建使用框架创建下面的图：

![re-graph](https://s3.eu-central-1.amazonaws.com/entgo.io/assets/re_graph_getting_started.png)


```go

func CreateGraph(ctx context.Context, client *ent.Client) error {
	// First, create the users.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.
		Create().
		SetAge(28).
		SetName("Neta").
		Save(ctx)
	if err != nil {
		return err
	}
	// Then, create the cars, and attach them to the users in the creation.
	_, err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(neta).              // attach this graph to Neta.
		Save(ctx)
	if err != nil {
		return err
	}
	// Create the groups, and add their users in the creation.
	_, err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}
```

现在有了一个有数据的图，我们可以运行下面这些查询：

1. 在名为"Github"的group中查询user的car:

	```go
	import (
		"log"
		
		"<project>/ent"
		"<project>/ent/group"
	)

	func QueryGithub(ctx context.Context, client *ent.Client) error {
		cars, err := client.Group.
			Query().
			Where(group.Name("GitHub")). // (Group(Name=GitHub),)
			QueryUsers().                // (User(Name=Ariel, Age=30),)
			QueryCars().                 // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
			All(ctx)
		if err != nil {
			return fmt.Errorf("failed getting cars: %v", err)
		}
		log.Println("cars returned:", cars)
		// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
		return nil
	}
	```

2. 修改上面的查询，遍历用户*Ariel*的资源:

	```go
	import (
		"log"
		
		"<project>/ent"
		"<project>/ent/car"
	)

	func QueryArielCars(ctx context.Context, client *ent.Client) error {
		// Get "Ariel" from previous steps.
		a8m := client.User.
			Query().
			Where(
				user.HasCars(),
				user.Name("Ariel"),
			).
			OnlyX(ctx)
		cars, err := a8m. 						// Get the groups, that a8m is connected to:
				QueryGroups(). 					// (Group(Name=GitHub), Group(Name=GitLab),)
				QueryUsers().  					// (User(Name=Ariel, Age=30), User(Name=Neta, Age=28),)
				QueryCars().   					//
				Where(         					//
					car.Not( 					//	Get Neta and Ariel cars, but filter out
						car.ModelEQ("Mazda"),	//	those who named "Mazda"
					), 							//
				). 								//
				All(ctx)
		if err != nil {
			return fmt.Errorf("failed getting cars: %v", err)
		}
		log.Println("cars returned:", cars)
		// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Ford, RegisteredAt=<Time>),)
		return nil
	}
	```

3. 查询所有的有用户的group(使用备用predicate进行查询):

	```go
	import (
		"log"
		
		"<project>/ent"
		"<project>/ent/group"
	)

	func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
    	groups, err := client.Group.
    		Query().
    		Where(group.HasUsers()).
    		All(ctx)
    	if err != nil {
    		return fmt.Errorf("failed getting groups: %v", err)
    	}
    	log.Println("groups returned:", groups)
    	// Output: (Group(Name=GitHub), Group(Name=GitLab),)
    	return nil
    }
    ```

完整的例子在[GitHub](https://github.com/facebook/ent/tree/master/examples/start).
