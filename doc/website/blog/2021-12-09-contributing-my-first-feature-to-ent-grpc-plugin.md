---
title: "What I learned contributing my first feature to Ent's gRPC plugin"
author: Jeremy Vesperman
authorURL: "https://github.com/jeremyv2014"
authorImageURL: "https://avatars.githubusercontent.com/u/9276415?v=4"
image: https://entgo.io/images/assets/grpc/ent_party.png
---

I've been writing software for years, but, until recently, I didn't know what an ORM was. I learned many things 
obtaining my B.S. in Computer Engineering, but Object-Relational Mapping was not one of those; I was too focused on 
building things out of bits and bytes to be bothered with something that high-level. It shouldn't be too surprising 
then, that when I found myself tasked with helping to build a distributed web application, I ended up outside my comfort 
zone.

One of the difficulties with developing software for someone else is, that you aren't able to see inside their head. The 
requirements aren't always clear and asking questions only helps you understand so much of what they are looking for. 
Sometimes, you just have to build a prototype and demonstrate it to get useful feedback.

The issue with this approach, of course, is that it takes time to develop prototypes, and you need to pivot frequently. 
If you were like me and didn't know what an ORM was, you would waste a lot of time doing simple, but time-consuming 
tasks:
1. Re-define the data model with new customer feedback.
2. Re-create the test database.
3. Re-write the SQL statements for interfacing with the database.
4. Re-define the gRPC interface between the backend and frontend services.
5. Re-design the frontend and web interface.
6. Demonstrate to customer and get feedback
7. Repeat

Hundreds of hours of work only to find out that everything needs to be re-written. So frustrating! I think you can 
imagine my relief (and also embarrassment), when a senior developer asked me why I wasn't using an ORM 
like Ent. 


### Discovering Ent
It only took one day to re-implement our current data model with Ent. I couldn't believe I had been doing all this work 
by hand when such a framework existed! The gRPC integration through entproto was the icing on the cake! I could perform 
basic CRUD operations over gRPC just by adding a few annotations to my schema. This allows me to skip all the steps 
between data model definition and re-designing the web interface! There was, however, just one problem for my use case: 
How do you get the details of entities over the gRPC interface if you don't know their IDs ahead of time? I see that 
Ent can query all, but where is the `GetAll` method for entproto?

### Becoming an Open-Source Contributor
I was surprised to find it didn't exist! I could have added it to my project by implementing the feature in a separate 
service, but it seemed like a generic enough method to be generally useful. For years, I had wanted 
to find an open-source project that I could meaningfully contribute to; this seemed like the perfect opportunity!

So, after poking around entproto's source into the early morning hours, I managed to hack the feature in! Feeling 
accomplished, I opened a pull request and headed off to sleep, not realizing the learning experience I had just signed 
myself up for.

In the morning, I awoke to the disappointment of my pull request being closed by [Rotem](https://github.com/rotemtam), 
but with an invitation to collaborate further to refine the idea. The reason for closing the request was obvious, my 
implementation of `GetAll` was dangerous. Returning an entire table's worth of data is only feasible if the table is 
small. Exposing this interface on a large table could have disastrous results!

### Optional Service Method Generation
My solution was to make the `GetAll` method optional by passing an argument into `entproto.Service()`. This 
provides control over whether this feature is exposed. We decided that this was a desirable feature, but that 
it should be more generic. Why should `GetAll` get special treatment just because it was added last? It would be better 
if all methods could be optionally generated. Something like:
```go
entproto.Service(entproto.Methods(entproto.Create | entproto.Get))
```
However, to keep everything backwards-compatible, an empty `entproto.Service()` annotation would also need to generate 
all methods. I'm not a Go expert, so the only way I knew of to do this was with a variadic function:
```go
func Service(methods ...Method)
```
The problem with this approach is that you can only have one argument type that is variable length. What if we wanted to 
add additional options to the service annotation later on? This is where I was introduced to the powerful design pattern 
of [functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis):

```go
// ServiceOption configures the entproto.Service annotation.
type ServiceOption func(svc *service)

// Service annotates an ent.Schema to specify that protobuf service generation is required for it.
func Service(opts ...ServiceOption) schema.Annotation {
	s := service{
		Generate: true,
	}
	for _, apply := range opts {
		apply(&s)
	}
	// Default to generating all methods
	if s.Methods == 0 {
		s.Methods = MethodAll
	}
	return s
}
```
This approach takes in a variable number of functions that are called to set options on a struct, in this case, our 
service annotation. With this approach, we can implement any number of other options functions aside from `Methods`. 
Very cool!

### List: The Superior GetAll
With optional method generation out of the way, we could return our focus to adding `GetAll`. How could we implement 
this method in a safe fashion? Rotem suggested we base the method off of Google's API Improvement Proposal (AIP) for List, 
[AIP-132](https://google.aip.dev/132). This approach allows a client to retrieve all entities, but breaks the retrieval 
up into pages. As an added bonus, it also sounds better than "GetAll"!


### List Request
With this design, a request message would look like:
```protobuf
message ListUserRequest {
  int32 page_size = 1;

  string page_token = 2;

  View view = 3;

  enum View {
    VIEW_UNSPECIFIED = 0;

    BASIC = 1;

    WITH_EDGE_IDS = 2;
  }
}
```

#### Page Size
The `page_size` field allows the client to specify the maximum number of entries they want to receive in the 
response message, subject to a maximum page size of 1000. This eliminates the issue of returning more results than the 
client can handle in the initial `GetAll` implementation. Additionally, the maximum page size was implemented to prevent 
a client from overburdening the server.

#### Page Token
The `page_token` field is a base64-encoded string utilized by the server to determine where the next page begins. An 
empty token means that we want the first page.

#### View
The `view` field is used to specify whether the response should return the edge IDs associated with the entities.


### List Response
The response message would look like:
```protobuf
message ListUserResponse {
  repeated User user_list = 1;

  string next_page_token = 2;
}
```

#### List
The `user_list` field contains page entities.

#### Next Page Token
The `next_page_token` field is a base64-encoded string that can be utilized in another List request to retrieve the next 
page of entities. An empty token means that this response contains the last page of entities.


### Pagination
With the gRPC interface determined, the challenge of implementing it began. One of the most critical design decisions 
was how to implement the pagination. The naive approach would be to use `LIMIT/OFFSET` pagination to skip over 
the entries we've already seen. However, this approach has massive [drawbacks](https://use-the-index-luke.com/no-offset); 
the most problematic being that the database has to _fetch all the rows it is skipping_ to get the rows we want. 

#### Keyset Pagination
Rotem proposed a much better approach: keyset pagination. This approach is slightly more 
complicated since it requires the use of a unique column (or combination of columns) to order the rows. But 
in exchange we gain a significant performance improvement. This is because we can take advantage of the sorted rows to select only entries with 
unique column(s) values that are greater (ascending order) or less (descending order) than / equal to the value(s) in 
the client-provided page token. Thus, the database doesn't have to fetch the rows we want to skip over, significantly 
speeding up queries on large tables!

With keyset pagination selected, the next step was to determine how to order the entities. The most straightforward 
approach for Ent was to use the `id` field; every schema will have this, and it is guaranteed to be unique for the schema. 
This is the approach we chose to use for the initial implementation. Additionally, a decision needed to be made regarding 
whether ascending or descending order should be employed. Descending order was chosen for the initial release.


### Usage
Let's take a look at how to actually use the new `List` feature:

```go
package main

import (
  "context"
  "log"

  "ent-grpc-example/ent/proto/entpb"
  "google.golang.org/grpc"
  "google.golang.org/grpc/status"
)

func main() {
  // Open a connection to the server.
  conn, err := grpc.Dial(":5000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("failed connecting to server: %s", err)
  }
  defer conn.Close()
  // Create a User service Client on the connection.
  client := entpb.NewUserServiceClient(conn)
  ctx := context.Background()
  // Initialize token for first page.
  pageToken := ""
  // Retrieve all pages of users.
  for {
    // Ask the server for the next page of users, limiting entries to 100.
    users, err := client.List(ctx, &entpb.ListUserRequest{
        PageSize:  100,
        PageToken: pageToken,
    })
    if err != nil {
        se, _ := status.FromError(err)
        log.Fatalf("failed retrieving user list: status=%s message=%s", se.Code(), se.Message())
    }
    // Check if we've reached the last page of users.
    if users.NextPageToken == "" {
        break
    }
    // Update token for next request.
    pageToken = users.NextPageToken
    log.Printf("users retrieved: %v", users)
  }
}
```


### Looking Ahead
The current implementation of `List` has a few limitations that can be addressed in future revisions. First, sorting is 
limited to the `id` column. This makes `List` compatible with any schema, but it isn't very flexible. Ideally, the client 
should be able to specify what columns to sort by. Alternatively, the sort column(s) could be defined in the schema. 
Additionally, `List` is restricted to descending order. In the future, this could be an option specified in the request. 
Finally, `List` currently only works with schemas that use `int32`, `uuid`, or `string` type `id` fields. This is because 
a separate conversion method to/from the page token must be defined for each type that Ent supports in the code generation 
template (I'm only one person!).


### Wrap-up
I was pretty nervous when I first embarked on my quest to contribute this functionality to entproto; as a newbie open-source 
contributor, I didn't know what to expect. I'm happy to share that working on the Ent project was a ton of fun! 
I got to work with awesome, knowledgeable people while helping out the open-source community. From functional 
options and keyset pagination to smaller insights gained through PR review, I learned so much about Go 
(and software development in general) in the process! I'd highly encourage anyone thinking they might want to contribute 
something to take that leap! You'll be surprised with how much you gain from the experience.

Have questions? Need help with getting started? Feel free to join our [Discord server](https://discord.gg/qZmPgTE6RX) or [Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::