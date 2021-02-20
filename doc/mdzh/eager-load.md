---
id: eager-load
title: Eager Loading
---

## 概述

`ent`支持查询实体及其关联(通过`edges`)。相关的实体
被填充到返回对象的`Edges`字段中。

看如下的例子：

![er-group-users](https://s3.eu-central-1.amazonaws.com/entgo.io/assets/er_user_pets_groups.png)



**查询所有的用户及其他们的pet**
```go
users, err := client.User.
	Query().
	WithPets().
	All(ctx)
if err != nil {
	return err
}
// The returned users look as follows:
//
//	[
//		User {
//			ID:   1,
//			Name: "a8m",
//			Edges: {
//				Pets: [Pet(...), ...]
//				...
//			}
//		},
//		...
//	]
//
for _, u := range users {
	for _, p := range u.Edges.Pets {
		fmt.Printf("User(%v) -> Pet(%v)\n", u.ID, p.ID)
		// Output:
		// User(...) -> Pet(...)
	}
} 
```

即时加载允许查询多个关联(包括嵌套)，而且允许
筛选、排序或限制它们的结果。例如:

```go
admins, err := client.User.
	Query().
	Where(user.Admin(true)).
	// Populate the `pets` that associated with the `admins`.
	WithPets().
	// Populate the first 5 `groups` that associated with the `admins`.
	WithGroups(func(q *ent.GroupQuery) {
		q.Limit(5) 				// Limit to 5.
		q.WithUsers().Limit(5)	// Populate the `users` of each `groups`. 
	}).
	All(ctx)
if err != nil {
	return err
}

// The returned users look as follows:
//
//	[
//		User {
//			ID:   1,
//			Name: "admin1",
//			Edges: {
//				Pets:   [Pet(...), ...]
//				Groups: [
//					Group {
//						ID:   7,
//						Name: "GitHub",
//						Edges: {
//							Users: [User(...), ...]
//							...
//						}
//					}
//				]
//			}
//		},
//		...
//	]
//
for _, admin := range admins {
	for _, p := range admin.Edges.Pets {
		fmt.Printf("Admin(%v) -> Pet(%v)\n", u.ID, p.ID)
		// Output:
		// Admin(...) -> Pet(...)
	}
	for _, g := range admin.Edges.Groups {
		for _, u := range g.Edges.Users {
			fmt.Printf("Admin(%v) -> Group(%v) -> User(%v)\n", u.ID, g.ID, u.ID)
			// Output:
			// Admin(...) -> Group(...) -> User(...)
		}
	}
} 
```

## API

每一个query-builder的edges都有一个方法列表，形如：`With<E>(...func(<N>Query))`。
`<E>` 表示边缘名称(如`WithGroups`)， `<N>`表示边缘类型(如`GroupQuery`)。

注意只有SQL方言支持这些特征。

## 实现

自从query-builder能够载入超过一个关联之后，它就不能使用`JSON`操作。
因此，`ent` 执行加载的额外查询。一个查询`M2O/O2M`和`O2O`的边缘，和
2个加载`M2M`边缘的查询。

注意，我们期待在下一个`ent`版本升级这些特性。