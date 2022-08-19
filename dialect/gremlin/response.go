// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"errors"
	"fmt"

	"entgo.io/ent/dialect/gremlin/encoding/graphson"
	"entgo.io/ent/dialect/gremlin/graph"
)

// A Response models a response message received from the server.
type Response struct {
	RequestID string `json:"requestId" graphson:"g:UUID"`
	Status    struct {
		Code       int            `json:"code"`
		Attributes map[string]any `json:"attributes"`
		Message    string         `json:"message"`
	} `json:"status"`
	Result struct {
		Data graphson.RawMessage `json:"data"`
		Meta map[string]any      `json:"meta"`
	} `json:"result"`
}

// IsErr returns whether response indicates an error.
func (rsp *Response) IsErr() bool {
	switch rsp.Status.Code {
	case StatusSuccess, StatusNoContent, StatusPartialContent:
		return false
	default:
		return true
	}
}

// Err returns an error representing response status.
func (rsp *Response) Err() error {
	if rsp.IsErr() {
		return fmt.Errorf("gremlin: code=%d, message=%q", rsp.Status.Code, rsp.Status.Message)
	}
	return nil
}

// ReadVal reads gremlin response data into v.
func (rsp *Response) ReadVal(v any) error {
	if err := rsp.Err(); err != nil {
		return err
	}
	if err := graphson.Unmarshal(rsp.Result.Data, v); err != nil {
		return fmt.Errorf("gremlin: unmarshal response data: type=%T: %w", v, err)
	}
	return nil
}

// ReadVertices returns response data as slice of vertices.
func (rsp *Response) ReadVertices() ([]graph.Vertex, error) {
	var v []graph.Vertex
	err := rsp.ReadVal(&v)
	return v, err
}

// ReadVertexProperties returns response data as slice of vertex properties.
func (rsp *Response) ReadVertexProperties() ([]graph.VertexProperty, error) {
	var vp []graph.VertexProperty
	err := rsp.ReadVal(&vp)
	return vp, err
}

// ReadEdges returns response data as slice of edges.
func (rsp *Response) ReadEdges() ([]graph.Edge, error) {
	var e []graph.Edge
	err := rsp.ReadVal(&e)
	return e, err
}

// ReadProperties returns response data as slice of properties.
func (rsp *Response) ReadProperties() ([]graph.Property, error) {
	var p []graph.Property
	err := rsp.ReadVal(&p)
	return p, err
}

// ReadValueMap returns response data as a value map.
func (rsp *Response) ReadValueMap() (graph.ValueMap, error) {
	var m graph.ValueMap
	err := rsp.ReadVal(&m)
	return m, err
}

// ReadBool returns response data as a bool.
func (rsp *Response) ReadBool() (bool, error) {
	var b [1]*bool
	if err := rsp.ReadVal(&b); err != nil {
		return false, err
	}
	if b[0] == nil {
		return false, errors.New("gremlin: no boolean value")
	}
	return *b[0], nil
}

// ReadInt returns response data as an int.
func (rsp *Response) ReadInt() (int, error) {
	var v [1]*int
	if err := rsp.ReadVal(&v); err != nil {
		return 0, err
	}
	if v[0] == nil {
		return 0, errors.New("gremlin: no integer value")
	}
	return *v[0], nil
}

// ReadString returns response data as a string.
func (rsp *Response) ReadString() (string, error) {
	var v [1]*string
	if err := rsp.ReadVal(&v); err != nil {
		return "", err
	}
	if v[0] == nil {
		return "", errors.New("gremlin: no string value")
	}
	return *v[0], nil
}
