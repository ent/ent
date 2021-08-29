package otel

import (
	"context"

	sc "entgo.io/ent/dialect/sqlcommenter"
	"go.opentelemetry.io/otel"
)

type (
	HttpCommenter  struct{}
	CommentCarrier sc.SqlComments
)

func NewCommentCarrier() CommentCarrier {
	return make(CommentCarrier)
}

// Get returns the value associated with the passed key.
func (c CommentCarrier) Get(key string) string {
	return string(c[sc.CommentKey(key)])
}

// Set stores the key-value pair.
func (c CommentCarrier) Set(key string, value string) {
	c[sc.CommentKey(key)] = sc.CommentValue(value)
}

// Keys lists the keys stored in this carrier.
func (c CommentCarrier) Keys() []string {
	keys := make([]string, 0, len(c))
	for k := range c {
		keys = append(keys, string(k))
	}
	return keys
}

// ref: https://github1s.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/emicklei/go-restful/otelrestful/restful_test.go
func (hc HttpCommenter) GetComments(ctx context.Context) sc.SqlComments {
	comments := NewCommentCarrier()
	otel.GetTextMapPropagator().Inject(ctx, comments)
	return sc.SqlComments(comments)
}
