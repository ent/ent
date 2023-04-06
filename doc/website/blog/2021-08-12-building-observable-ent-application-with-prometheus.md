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
visible way a quantitative description of the running system's state. For instance, to expose an HTTP endpoint where we 
can see how many errors occurred since the process has started. In this post, we will explore how to build more
observable Ent applications using Prometheus.

### What is Ent?

[Ent](https://entgo.io/docs/getting-started/), is a simple, yet powerful entity framework for Go, that makes it easy
to build and maintain applications with large data models.

### What is Prometheus?

[Prometheus](https://prometheus.io/) is an open source monitoring system developed by engineering at SoundCloud in 2012.
It includes an embedded time series database and many integrations to third-party systems.
The Prometheus client exposes the process's metrics via an HTTP endpoint (usually `/metrics`), this endpoint is
discovered by the Prometheus scraper which polls the endpoint every interval (typically 30s) and writes it
into a time-series database.

Prometheus is just an example of a class of metric collection backends. Many others, such as AWS CloudWatch, InfluxDB
and others exist and are in wide use in the industry. Towards the end of this post, we will discuss a possible path to
a unified, standards-based integration with any such backend.

### Working with Prometheus

To expose an application's metrics using Prometheus, we need to create a
Prometheus [Collector](https://prometheus.io/docs/introduction/glossary/#collector), a collector collects
a set of metrics from your server.

In our example, we will be using [two types of metrics](https://prometheus.io/docs/concepts/metric_types/)
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
They function similar to the popular [HTTP middleware pattern](https://github.com/go-chi/chi#middleware-handlers).

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
			// Do something before mutation.
			v, err := next.Mutate(ctx, m)
			if err != nil {
				// Do something if error after mutation.
			}
			// Do something after mutation.
			return v, err
		})
	}
}
```

In Ent, there are two types of mutation hooks - schema hooks and runtime hooks. Schema hooks are mainly used for defining
custom mutation logic on a specific entity type, for example, syncing entity creation to another system. Runtime hooks, on the other hand, are used
to define more global logic for adding things like logging, metrics, tracing, etc.

For our use case, we should definitely use runtime hooks, because to be valuable we want to export metrics on all
operations on all entity types:

```go
package example

import (
	"entprom/ent"
	"entprom/ent/hook"
)

func main() {
	client, _ := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	// Add a hook only on user mutations.
	client.User.Use(exampleHook())

	// Add a hook only on update operations.
	client.Use(hook.On(exampleHook(), ent.OpUpdate|ent.OpUpdateOne))
}
```

### Exporting Prometheus Metrics for an Ent Application

With all of the introductions complete, let’s cut to the chase and show how to use Prometheus and Ent hooks together to
create an observable application. Our goal with this example is to export these metrics using a hook:

| Metric Name                    | Description                              |
|--------------------------------|------------------------------------------|
| ent_operation_total            | Number of ent mutation operations        |
| ent_operation_error            | Number of failed ent mutation operations |
| ent_operation_duration_seconds | Time in seconds per operation            |


Each of these metrics will be broken down by labels into two dimensions:
* `mutation_type`: Entity type that is being mutated (User, BlogPost, Account etc.).
* `mutation_op`: The operation that is being performed (Create, Delete etc.).

Let’s start by defining our collectors:

```go
//Ent dynamic dimensions
const (
	mutationType = "mutation_type"
	mutationOp   = "mutation_op"
)

var entLabels = []string{mutationType, mutationOp}

// Create a collector for total operations counter
func initOpsProcessedTotal() *prometheus.CounterVec {
	return promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ent_operation_total",
			Help: "Number of ent mutation operations",
		},
		entLabels,
	)
}

// Create a collector for error counter
func initOpsProcessedError() *prometheus.CounterVec {
	return promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ent_operation_error",
			Help: "Number of failed ent mutation operations",
		},
		entLabels,
	)
}

// Create a collector for duration histogram collector
func initOpsDuration() *prometheus.HistogramVec {
	return promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ent_operation_duration_seconds",
			Help: "Time in seconds per operation",
		},
		entLabels,
	)
}
```
Next, let’s define our new hook:

```go
// Hook init collectors, count total at beginning error on mutation error and duration also after.
func Hook() ent.Hook {
	opsProcessedTotal := initOpsProcessedTotal()
	opsProcessedError := initOpsProcessedError()
	opsDuration := initOpsDuration()
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			// Before mutation, start measuring time.
			start := time.Now()
			// Extract dynamic labels from mutation.
			labels := prometheus.Labels{mutationType: m.Type(), mutationOp: m.Op().String()}
			// Increment total ops counter.
			opsProcessedTotal.With(labels).Inc()
			// Execute mutation.
			v, err := next.Mutate(ctx, m)
			if err != nil {
				// In case of error increment error counter.
				opsProcessedError.With(labels).Inc()
			}
			// Stop time measure.
			duration := time.Since(start)
			// Record duration in seconds.
			opsDuration.With(labels).Observe(duration.Seconds())
			return v, err
		})
	}
}
```

### Connecting the Prometheus Collector to our Service

After defining our hook, let’s see next how to connect it to our application and how to use Prometheus to serve
an endpoint that exposes the metrics in our collectors:

```go
package main

import (
	"context"
	"log"
	"net/http"

	"entprom"
	"entprom/ent"

	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func createClient() *ent.Client {
	c, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	ctx := context.Background()
	// Run the auto migration tool.
	if err := c.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return c
}

func handler(client *ent.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		// Run operations.
		_, err := client.User.Create().SetName("a8m").Save(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func main() {
	// Create Ent client and migrate
	client := createClient()
	// Use the hook
	client.Use(entprom.Hook())
	// Simple handler to run actions on our DB.
	http.HandleFunc("/", handler(client))
	// This endpoint sends metrics to the prometheus to collect
	http.Handle("/metrics", promhttp.Handler())
	log.Println("server starting on port 8080")
	// Run the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

After a few times of accessing `/` on our server (using `curl` or a browser), go to `/metrics`. There you will see the output from the Prometheus client:

```
# HELP ent_operation_duration_seconds Time in seconds per operation
# TYPE ent_operation_duration_seconds histogram
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="0.005"} 2
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="0.01"} 2
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="0.025"} 2
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="0.05"} 2
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="0.1"} 2
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="0.25"} 2
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="0.5"} 2
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="1"} 2
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="2.5"} 2
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="5"} 2
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="10"} 2
ent_operation_duration_seconds_bucket{mutation_op="OpCreate",mutation_type="User",le="+Inf"} 2
ent_operation_duration_seconds_sum{mutation_op="OpCreate",mutation_type="User"} 0.000265669
ent_operation_duration_seconds_count{mutation_op="OpCreate",mutation_type="User"} 2
# HELP ent_operation_error Number of failed ent mutation operations
# TYPE ent_operation_error counter
ent_operation_error{mutation_op="OpCreate",mutation_type="User"} 1
# HELP ent_operation_total Number of ent mutation operations
# TYPE ent_operation_total counter
ent_operation_total{mutation_op="OpCreate",mutation_type="User"} 2
```

In the top part, we can see the histogram calculated, it calculates the number of operations in each “bucket”.
After that, we can see the number of total operations and the number of errors.
Each metric is followed by its description that can be seen when querying with Prometheus dashboard.

The Prometheus client is only one component of the Prometheus architecture.
To run a complete system including a scraper that will poll your endpoint, a Prometheus that will store your
metrics and can answer queries, and a simple UI to interact with it, I recommend reading the official
documentation or use the docker-compose.yaml in this example [repo](https://github.com/yonidavidson/ent-prometheus-example).

### Future Work on Observability in Ent

As we’ve mentioned above, there is an abundance of metric collections backends available today,
Prometheus being just one of many successful projects. While these solutions differ in many dimensions
(self-hosted vs SaaS, different storage engines with different query languages, and more) - from the metric
reporting client perspective, they are virtually identical.

In cases like these, good software engineering principles suggest that the concrete backend should be abstracted away
from the client using an interface. This interface can then be implemented by backends so client applications can easily
switch between the different implementations. Such changes are happening in recent years in our industry. Consider,
for example, the [Open Container Initiative](https://opencontainers.org/) or the
[Service Mesh Interface](https://smi-spec.io/): both are initiatives that strive to define a standard interface
for a problem space. This interface is supposed to create an ecosystem of implementations of
the standard.  In the observability space, the exact same convergence is occurring with [OpenCensus](https://opencensus.io/) and 
[OpenTracing](https://opentracing.io/) currently merging into [OpenTelemetry](https://opentelemetry.io/).

As nice as it would be to publish an Ent + Prometheus extension similar to the one presented in this post, we are firm
believers that observability should be solved with a standards-based approach.
We invite everyone to [join the discussion](https://github.com/ent/ent/discussions/1819) on what is the right way to do this for Ent.


### Wrap-Up

We started this post by presenting Prometheus, a popular open-source monitoring solution. Next, we reviewed “Hooks”, a feature
of Ent that allows adding custom logic before and after operations that change the data entities. We then showed how
to integrate the two to create observable applications using Ent.  Finally, we discussed the future of observability in Ent
and invited everyone to join the discussion [to shape it](https://github.com/ent/ent/discussions/1819).


Have questions? Need help with getting started? Feel free to join our [Discord server](https://discord.gg/qZmPgTE6RX) or [Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
