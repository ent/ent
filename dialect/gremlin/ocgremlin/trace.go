// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package ocgremlin

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/gremlin"

	"go.opencensus.io/trace"
)

// Attributes recorded on the span for the requests.
const (
	RequestIDAttribute = "gremlin.request_id"
	OperationAttribute = "gremlin.operation"
	QueryAttribute     = "gremlin.query"
	BindingAttribute   = "gremlin.binding"
	CodeAttribute      = "gremlin.code"
	MessageAttribute   = "gremlin.message"
)

type traceTransport struct {
	base           gremlin.RoundTripper
	startOptions   trace.StartOptions
	formatSpanName func(context.Context, *gremlin.Request) string
	withQuery      bool
}

func (t *traceTransport) RoundTrip(ctx context.Context, req *gremlin.Request) (*gremlin.Response, error) {
	ctx, span := trace.StartSpan(ctx,
		t.formatSpanName(ctx, req),
		trace.WithSampler(t.startOptions.Sampler),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer span.End()

	span.AddAttributes(requestAttrs(req, t.withQuery)...)
	rsp, err := t.base.RoundTrip(ctx, req)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeUnknown, Message: err.Error()})
		return rsp, err
	}

	span.AddAttributes(responseAttrs(rsp)...)
	span.SetStatus(TraceStatus(rsp.Status.Code))
	return rsp, err
}

func requestAttrs(req *gremlin.Request, withQuery bool) []trace.Attribute {
	attrs := []trace.Attribute{
		trace.StringAttribute(RequestIDAttribute, req.RequestID),
		trace.StringAttribute(OperationAttribute, req.Operation),
	}
	if withQuery {
		query, _ := req.Arguments[gremlin.ArgsGremlin].(string)
		attrs = append(attrs, trace.StringAttribute(QueryAttribute, query))
		if bindings, ok := req.Arguments[gremlin.ArgsBindings].(map[string]any); ok {
			attrs = append(attrs, bindingsAttrs(bindings)...)
		}
	}
	return attrs
}

func bindingsAttrs(bindings map[string]any) []trace.Attribute {
	attrs := make([]trace.Attribute, 0, len(bindings))
	for key, val := range bindings {
		key = BindingAttribute + "." + key
		attrs = append(attrs, bindingToAttr(key, val))
	}
	return attrs
}

func bindingToAttr(key string, val any) trace.Attribute {
	switch v := val.(type) {
	case nil:
		return trace.StringAttribute(key, "")
	case int64:
		return trace.Int64Attribute(key, v)
	case float64:
		return trace.Float64Attribute(key, v)
	case string:
		return trace.StringAttribute(key, v)
	case bool:
		return trace.BoolAttribute(key, v)
	default:
		s := fmt.Sprintf("%v", v)
		if len(s) > 256 {
			s = s[:256]
		}
		return trace.StringAttribute(key, s)
	}
}

func responseAttrs(rsp *gremlin.Response) []trace.Attribute {
	attrs := []trace.Attribute{
		trace.Int64Attribute(CodeAttribute, int64(rsp.Status.Code)),
	}
	if rsp.Status.Message != "" {
		attrs = append(attrs, trace.StringAttribute(MessageAttribute, rsp.Status.Message))
	}
	return attrs
}

// TraceStatus is a utility to convert the gremlin status code to a trace.Status.
func TraceStatus(status int) trace.Status {
	var code int32
	switch status {
	case gremlin.StatusSuccess,
		gremlin.StatusNoContent,
		gremlin.StatusPartialContent:
		code = trace.StatusCodeOK
	case gremlin.StatusUnauthorized:
		code = trace.StatusCodePermissionDenied
	case gremlin.StatusAuthenticate:
		code = trace.StatusCodeUnauthenticated
	case gremlin.StatusMalformedRequest,
		gremlin.StatusInvalidRequestArguments,
		gremlin.StatusScriptEvaluationError:
		code = trace.StatusCodeInvalidArgument
	case gremlin.StatusServerError,
		gremlin.StatusServerSerializationError:
		code = trace.StatusCodeInternal
	case gremlin.StatusServerTimeout:
		code = trace.StatusCodeDeadlineExceeded
	default:
		code = trace.StatusCodeUnknown
	}
	return trace.Status{Code: code, Message: gremlin.StatusText(status)}
}
