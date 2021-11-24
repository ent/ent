// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package ocgremlin

import (
	"context"
	"strconv"
	"time"

	"entgo.io/ent/dialect/gremlin"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// The following measures are supported for use in custom views.
var (
	RequestCount = stats.Int64(
		"gremlin/request_count",
		"Number of Gremlin requests started",
		stats.UnitDimensionless,
	)
	ResponseBytes = stats.Int64(
		"gremlin/response_bytes",
		"Total number of bytes in response data",
		stats.UnitBytes,
	)
	RoundTripLatency = stats.Float64(
		"gremlin/roundtrip_latency",
		"End-to-end latency",
		stats.UnitMilliseconds,
	)
)

// The following tags are applied to stats recorded by this package.
var (
	// StatusCode is the numeric Gremlin response status code,
	// or "error" if a transport error occurred and no status code was read.
	StatusCode, _ = tag.NewKey("gremlin_status_code")
)

// Default distributions used by views in this package.
var (
	DefaultSizeDistribution    = view.Distribution(32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576)
	DefaultLatencyDistribution = view.Distribution(1, 2, 3, 4, 5, 6, 8, 10, 13, 16, 20, 25, 30, 40, 50, 65, 80, 100, 130, 160, 200, 250, 300, 400, 500, 650, 800, 1000, 2000, 5000, 10000, 20000, 50000, 100000)
)

// Package ocgremlin provides some convenience views for measures.
// You still need to register these views for data to actually be collected.
var (
	RequestCountView = &view.View{
		Name:        "gremlin/request_count",
		Measure:     RequestCount,
		Aggregation: view.Count(),
		Description: "Count of Gremlin requests started",
	}

	ResponseCountView = &view.View{
		Name:        "gremlin/response_count",
		Measure:     RoundTripLatency,
		Aggregation: view.Count(),
		Description: "Count of responses received, by response status",
		TagKeys:     []tag.Key{StatusCode},
	}

	ResponseBytesView = &view.View{
		Name:        "gremlin/response_bytes",
		Measure:     ResponseBytes,
		Aggregation: DefaultSizeDistribution,
		Description: "Total number of bytes in response data",
	}

	RoundTripLatencyView = &view.View{
		Name:        "gremlin/roundtrip_latency",
		Measure:     RoundTripLatency,
		Aggregation: DefaultLatencyDistribution,
		Description: "End-to-end latency, by response code",
		TagKeys:     []tag.Key{StatusCode},
	}
)

// Views are the default views provided by this package.
func Views() []*view.View {
	return []*view.View{
		RequestCountView,
		ResponseCountView,
		ResponseBytesView,
		RoundTripLatencyView,
	}
}

// statsTransport is an gremlin.RoundTripper that collects stats for the outgoing requests.
type statsTransport struct {
	base gremlin.RoundTripper
}

func (t statsTransport) RoundTrip(ctx context.Context, req *gremlin.Request) (*gremlin.Response, error) {
	stats.Record(ctx, RequestCount.M(1))
	start := time.Now()
	rsp, err := t.base.RoundTrip(ctx, req)
	latency := float64(time.Since(start)) / float64(time.Millisecond)
	var (
		tags = make([]tag.Mutator, 1)
		ms   = []stats.Measurement{RoundTripLatency.M(latency)}
	)
	if err == nil {
		tags[0] = tag.Upsert(StatusCode, strconv.Itoa(rsp.Status.Code))
		ms = append(ms, ResponseBytes.M(int64(len(rsp.Result.Data))))
	} else {
		tags[0] = tag.Upsert(StatusCode, "error")
	}
	_ = stats.RecordWithTags(ctx, tags, ms...)
	return rsp, err
}
