package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"entgo.io/ent/examples/softdelete/ent"
	"entgo.io/ent/examples/softdelete/ent/user"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&_fk=1", ent.Debug(), ent.Log(func(s ...any) {
		fmt.Println(s...)
	}))
	if err != nil {
		panic(err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	ctx := context.Background()

	p1, err := client.Pet.Create().SetName("p1").Save(ctx)
	if err != nil {
		panic(err)
	}
	p2, err := client.Pet.Create().SetName("p2").Save(ctx)
	if err != nil {
		panic(err)
	}
	p3, err := client.Pet.Create().SetName("p3").Save(ctx)
	if err != nil {
		panic(err)
	}
	p4, err := client.Pet.Create().SetName("p4").Save(ctx)
	if err != nil {
		panic(err)
	}
	u1, err := client.User.Create().SetName("u1").AddPets(p1, p2).Save(ctx)
	if err != nil {
		panic(err)
	}
	u2, err := client.User.Create().SetName("u2").AddPets(p3, p4).Save(ctx)
	if err != nil {
		panic(err)
	}

	g, err := client.Group.Create().SetName("group").AddUsers(u1, u2).Save(ctx)
	if err != nil {
		panic(err)
	}

	{
		u1, err = client.User.Query().Where(user.Name("u1")).First(ctx)
		if err != nil {
			panic(err)
		}
		if err := client.User.DeleteOne(u1).Exec(ctx); err != nil {
			panic(err)
		}

		_, err = client.User.Query().Where(user.Name("u1")).First(ctx)
		if err == nil {
			panic("found no soft delete user")
		} else {
			if !ent.IsNotFound(err) {
				panic(err)
			}
		}

		n, err := client.User.Update().Where(user.Name("u1")).SetName("nu1").Save(ctx)
		if err != nil {
			panic(err)
		}
		if n > 0 {
			panic("set name to soft delete user")
		}
		n, err = client.User.Update().Real().Where(user.Name("u1")).SetName("nu1").Save(ctx)
		if err != nil {
			panic(err)
		}
		if n == 0 {
			panic("set name to soft delete user not work")
		}
	}
	{
		if err := client.Pet.DeleteOne(p1).Exec(ctx); err != nil {
			panic(err)
		}
		if err := client.Pet.DeleteOne(p3).Exec(ctx); err != nil {
			panic(err)
		}

		{
			// normal query
			// {
			//   "id": 1,
			//   "name": "group",
			//   "edges": {
			//     "users": [
			//       {
			//         "id": 2,
			//         "name": "u2",
			//         "edges": {
			//           "pets": [
			//             {
			//               "id": 4,
			//               "name": "p4",
			//               "edges": {}
			//             }
			//           ]
			//         }
			//       }
			//     ]
			//   }
			// }
			g, err = client.Group.Query().WithUsers(func(uq *ent.UserQuery) {
				uq.WithPets()
			}).First(ctx)
			if err != nil {
				panic(err)
			}
			data, err := json.MarshalIndent(g, "", "  ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(data))
		}
		{
			// query real user
			// {
			//   "id": 1,
			//   "name": "group",
			//   "edges": {
			//     "users": [
			//       {
			//         "id": 1,
			//         "deleted_at": "2022-12-09T09:22:23.609494085+08:00",
			//         "name": "u1",
			//         "edges": {
			//           "pets": [
			//             {
			//               "id": 2,
			//               "name": "p2",
			//               "edges": {}
			//             }
			//           ]
			//         }
			//       },
			//       {
			//         "id": 2,
			//         "name": "u2",
			//         "edges": {
			//           "pets": [
			//             {
			//               "id": 4,
			//               "name": "p4",
			//               "edges": {}
			//             }
			//           ]
			//         }
			//       }
			//     ]
			//   }
			// }
			g, err = client.Group.Query().WithUsers(func(uq *ent.UserQuery) {
				uq.Real().WithPets()
			}).First(ctx)
			if err != nil {
				panic(err)
			}
			data, err := json.MarshalIndent(g, "", "  ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(data))
		}
		{
			// query real pet
			// {
			//   "id": 1,
			//   "name": "group",
			//   "edges": {
			//     "users": [
			//       {
			//         "id": 2,
			//         "name": "u2",
			//         "edges": {
			//           "pets": [
			//             {
			//               "id": 3,
			//               "deleted_at": "2022-12-09T09:23:15.245042904+08:00",
			//               "name": "p3",
			//               "edges": {}
			//             },
			//             {
			//               "id": 4,
			//               "name": "p4",
			//               "edges": {}
			//             }
			//           ]
			//         }
			//       }
			//     ]
			//   }
			// }
			g, err = client.Group.Query().WithUsers(func(uq *ent.UserQuery) {
				uq.WithPets(func(pq *ent.PetQuery) {
					pq.Real()
				})
			}).First(ctx)
			if err != nil {
				panic(err)
			}
			data, err := json.MarshalIndent(g, "", "  ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(data))
		}
		{
			// query all real
			// {
			//   "id": 1,
			//   "name": "group",
			//   "edges": {
			//     "users": [
			//       {
			//         "id": 1,
			//         "deleted_at": "2022-12-09T09:23:15.244923613+08:00",
			//         "name": "u1",
			//         "edges": {
			//           "pets": [
			//             {
			//               "id": 1,
			//               "deleted_at": "2022-12-09T09:23:15.245015666+08:00",
			//               "name": "p1",
			//               "edges": {}
			//             },
			//             {
			//               "id": 2,
			//               "name": "p2",
			//               "edges": {}
			//             }
			//           ]
			//         }
			//       },
			//       {
			//         "id": 2,
			//         "name": "u2",
			//         "edges": {
			//           "pets": [
			//             {
			//               "id": 3,
			//               "deleted_at": "2022-12-09T09:23:15.245042904+08:00",
			//               "name": "p3",
			//               "edges": {}
			//             },
			//             {
			//               "id": 4,
			//               "name": "p4",
			//               "edges": {}
			//             }
			//           ]
			//         }
			//       }
			//     ]
			//   }
			// }
			g, err = client.Group.Query().WithUsers(func(uq *ent.UserQuery) {
				uq.Real().WithPets(func(pq *ent.PetQuery) {
					pq.Real()
				})
			}).First(ctx)
			if err != nil {
				panic(err)
			}
			data, err := json.MarshalIndent(g, "", "  ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(data))
		}

	}
	{
		nu, err := client.User.Query().Real().Where(user.Name("nu1")).First(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(nu)
		if err := client.User.DeleteOne(nu).Real().Exec(ctx); err != nil {
			panic(err)
		}
		_, err = client.User.Query().Real().Where(user.Name("nu1")).First(ctx)
		if err == nil {
			panic("found no delete user")
		} else {
			if !ent.IsNotFound(err) {
				panic(err)
			}
		}
	}
}
