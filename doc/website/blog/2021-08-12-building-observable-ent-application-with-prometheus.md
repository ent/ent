---
title: Building Observable Ent Applications with Prometheus
author: Yoni Davidson
authorURL: "https://github.com/yonidavidson"
authorImageURL: "https://avatars0.githubusercontent.com/u/5472778"
authorTwitter: yonidavidson
---

Observability is a quality of a system that refers to how well its internal state can be measured externally.
As a computer program evolves into a full-blown production system this quality becomes increasingly important.
One of the ways to make a software system more observable is to export metrics, that is, to report in some externally 
visible way a quantitative description of the running system’s state. For instance, to expose an HTTP endpoint where we 
can see how many errors occurred since the process has started. In this post, we will explore how to build more
observable Ent applications using Prometheus.

### What is Ent?

[Ent](https://entgo.io/docs/getting-started/), is a simple, yet powerful entity framework for Go, that makes it easy
to build and maintain applications with large data models.

### What is Prometheus?

[Prometheus](https://prometheus.io/) is an open source monitoring system developed by engineering at SoundCloud in 2012.
It includes an embedded time series database and many integrations to third-party systems.
The Prometheus client exposes the process's metrics via an HTTP endpoint (usually /metrics), this endpoint is
discovered by the Prometheus scraper which polls the endpoint every interval (typically 30s) and writes it
into a time-series database.

Prometheus is just an example of a class of metric collection backends. Many others, such as AWS CloudWatch, InfluxDB
and others exist and are in wide use in the industry. Towards the end of this post, we will discuss a possible path to
a unified, standards-based integration with any such backend.

### Working with Prometheus

To expose an application’s metrics using Prometheus, we need to create a
prometheus [Collector](https://prometheus.io/docs/introduction/glossary/#collector), a collector collects
a set of metrics from your server.

In our example, we will be using [two types of metrics](https://prometheus.io/docs/concepts/metric_types/#histogram)
that can be stored in a collector: Counters and Histograms. Counters are monotonically increasing cumulative metrics
that represent how many times something has happened, commonly used to count the number of requests a server has
processed or errors that have occurred. Histograms sample observations into buckets of configurable sizes and are
commonly used to represent latency distributions (i.e how many requests returned in under 5ms, 10ms, 100ms, 1s, etc.)
In addition, Prometheus allows metrics to be broken down into labels.  This is useful for example for counting requests
but breaking down the counter by endpoint name.

Let’s see how to create such a collector using the [official Go client](https://github.com/prometheus/client_golang).
To do so, we will use a package in the client called [promauto](https://pkg.go.dev/github.com/prometheus/client_golang@v1.11.0/prometheus/promauto) that simplifies the processes of creating collectors.
A simple example of a collector that counts (for example, total request or number or request error):

```go
package example

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// List of dynamic labels
	labelNames = []string{"endpoint", "error_code"}

	// Create a counter collector
	exampleCollector = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "endpoint_errors",
			Help: "Number of errors in endpoints",
		},
		labelNames,
	)
)

// When using you set the values of the dynamic labels and then increment the counter
func incrementError() {
	exampleCollector.WithLabelValues("/create-user", "400").Inc()
}
```

### Ent Hooks

[Hooks](https://entgo.io/docs/hooks) are a feature of Ent that allows adding custom logic before and after operations that change the data entities.

A mutation is an operation that changes something in the database.
There are 5 types of mutations:
1. Create.
2. UpdateOne.
3. Update.
4. DeleteOne.
5. Delete.

Hooks are functions that get an [ent.Mutator](https://pkg.go.dev/entgo.io/ent#Mutator) and return a mutator back.
They function similar to the popular [HTTP middleware pattern](https://dzone.com/articles/understanding-middleware-pattern-in-expressjs).

```go
package example

import (
	"context"

	"entgo.io/ent"
)

func exampleHook() ent.Hook {
	//use this to init your hook
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			// Do something before mutation
			v, err := next.Mutate(ctx, m)
			if err != nil {
				// Do something if error after mutation
			}
			// Do something always after mutation
			return v, err
		})
	}
}
```

### Wrapping Up

In this post, we presented the Upsert API, a long-anticipated capability, that is available by feature-flag in Ent v0.9.0.
We discussed where upserts are commonly used in applications and the way they are implemented using common relational databases.
Finally, we showed a simple example of how to get started with the Upsert API using Ent.

Have questions? Need help with getting started? Feel free to [join our Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://www.getrevue.co/profile/ent)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)

:::
