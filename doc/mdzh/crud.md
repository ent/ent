---
id: crud
title: CRUD API
---

之前提到的[介绍](code-gen.md)。运行`ent`将会生成如下资源：

- `Client` 和 `Tx` 用于图形之间的交互
- 每个schema类型的CRUD构建器。更多内容阅读[CRUD](crud.md) 
- 每个schema类型的实体对象(Go struct)
- 包含用于与构建器交互的常量和断言的包
- `migrate` 一个进行数据迁移的包，更多内容阅读[Migration](migrate.md)

## 创建一个新的Client

**MySQL**

```go
package main

import (
	"log"

	"<project>/ent"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
```

**PostgreSQL**

```go
package main

import (
	"log"

	"<project>/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres","host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
```

**SQLite**

```go
package main

import (
	"log"

	"<project>/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
```


**Gremlin (AWS Neptune)**

```go
package main

import (
	"log"

	"<project>/ent"
)

func main() {
	client, err := ent.Open("gremlin", "http://localhost:8182")
	if err != nil {
		log.Fatal(err)
	}
}
```

## 创建一个实体

**保存** 一个实体.

```go
a8m, err := client.User.	// UserClient.
	Create().				// User create builder.
	SetName("a8m").			// Set field value.
	SetNillableAge(age).	// Avoid nil checks.
	AddGroups(g1, g2).		// Add many edges.
	SetSpouse(nati).		// Set unique edge.
	Save(ctx)				// Create and return.
```

**SaveX** a pet; Unlike **Save**, **SaveX** panics if an error occurs.

```go
pedro := client.Pet.	// PetClient.
	Create().			// Pet create builder.
	SetName("pedro").	// Set field value.
	SetOwner(a8m).		// Set owner (unique edge).
	SaveX(ctx)			// Create and return.
```

## 创建多个实体

**保存**多个实体.

```go
names := []string{"pedro", "xabi", "layla"}
bulk := make([]*ent.PetCreate, len(names))
for i, name := range names {
    bulk[i] = client.Pet.Create().SetName(name).SetOwner(a8m)
}
pets, err := client.Pet.CreateBulk(bulk...).Save(ctx)
```

## 更新一个实体

Update an entity that was returned from the database.

```go
a8m, err = a8m.Update().	// User update builder.
	RemoveGroup(g2).		// Remove specific edge.
	ClearCard().			// Clear unique edge.
	SetAge(30).				// Set field value
	Save(ctx)				// Save and return.
```


## 通过ID更新实体

```go
pedro, err := client.Pet.	// PetClient.
	UpdateOneID(id).		// Pet update builder.
	SetName("pedro").		// Set field name.
	SetOwnerID(owner).		// Set unique edge, using id.
	Save(ctx)				// Save and return.
```

## 更新多个实体 

使用判断语句筛选

```go
n, err := client.User.			// UserClient.
	Update().					// Pet update builder.
	Where(						//
		user.Or(				// (age >= 30 OR name = "bar") 
			user.AgeEQ(30), 	//
			user.Name("bar"),	// AND
		),						//  
		user.HasFollowers(),	// UserHasFollowers()  
	).							//
	SetName("foo").				// Set field name.
	Save(ctx)					// exec and return.
```

查询边缘判断

```go
n, err := client.User.			// UserClient.
	Update().					// Pet update builder.
	Where(						// 
		user.HasFriendsWith(	// UserHasFriendsWith (
			user.Or(			//   age = 20
				user.Age(20),	//      OR
				user.Age(30),	//   age = 30
			)					// )
		), 						//
	).							//
	SetName("a8m").				// Set field name.
	Save(ctx)					// exec and return.
```

## 查询图

获取所有的用户和关注者
```go
users, err := client.User.		// UserClient.
	Query().					// User query builder.
	Where(user.HasFollowers()).	// filter only users with followers.
	All(ctx)					// query and return.
```

获取到特定用户的跟随者。从图中的一个节点开始遍历。

```go
users, err := a8m.
	QueryFollowers().
	All(ctx)
```

获得一个用户的所有追随者的pet。
```go
users, err := a8m.
	QueryFollowersd).
	QueryPets().
	All(ctx)
```

更多高级的遍历在[下一章](traversals.md). 

## 字段选择

获取所有pet的名称

```go
names, err := client.Pet.
	Query().
	Select(pet.FieldName).
	Strings(ctx)
```

选择部分对象和部分关联关系。
获取所有的pet和他们的拥有者，但是选择和填充仅能使用`ID` 和 `Name` 字段。

```go
pets, err := client.Pet.
    Query().
    Select(pet.FieldName).
    WithOwner(func (q *ent.UserQuery) {
        q.Select(user.FieldName)
    }).
    All(ctx)
```

扫描所有的pet名称和年龄到自定义的结构体中。

```go
var v []struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}
err := client.Pet.
	Query().
	Select(pet.FieldAge, pet.FieldName).
	Scan(ctx, &v)
if err != nil {
	log.Fatal(err)
}
```

## 删除一个实体

删除一个实体

```go
err := client.User.
	DeleteOne(a8m).
	Exec(ctx)
```

通过`ID`删除

```go
err := client.User.
	DeleteOneID(id).
	Exec(ctx)
```

## 删除多个

通过条件删除

```go
_, err := client.File.
	Delete().
	Where(file.UpdatedAtLT(date)).
	Exec(ctx)
```

## 变化

Each generated node type has its own type of mutation. For example, all [`User` builders](crud.md#create-an-entity), share
the same generated `UserMutation` object.
However, all builder types implement the generic <a target="_blank" href="https://pkg.go.dev/entgo.io/ent?tab=doc#Mutation">`ent.Mutation`<a> interface.

For example, in order to write a generic code that apply a set of methods on both `ent.UserCreate`
and `ent.UserUpdate`, use the `UserMutation` object:

```go
func Do() {
    creator := client.User.Create()
    SetAgeName(creator.Mutation())
	updater := client.User.UpdateOneID(id)
	SetAgeName(updater.Mutation())
}

// SetAgeName sets the age and the name for any mutation.
func SetAgeName(m *ent.UserMutation) {
    m.SetAge(32)
    m.SetName("Ariel")
}
```

In some cases, you want to apply a set of methods on multiple types.
For cases like this, either use the generic `ent.Mutation` interface,
or create your own interface.

```go
func Do() {
	creator1 := client.User.Create()
	SetName(creator1.Mutation(), "a8m")

	creator2 := client.Pet.Create()
	SetName(creator2.Mutation(), "pedro")
}

// SetNamer wraps the 2 methods for getting
// and setting the "name" field in mutations.
type SetNamer interface {
	SetName(string)
	Name() (string, bool)
}

func SetName(m SetNamer, name string) {
    if _, exist := m.Name(); !exist {
    	m.SetName(name)
    }
}
```
