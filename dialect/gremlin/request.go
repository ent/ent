// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"bytes"
	"encoding/base64"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	// A Request models a request message sent to the server.
	Request struct {
		RequestID string         `json:"requestId" graphson:"g:UUID"`
		Operation string         `json:"op"`
		Processor string         `json:"processor"`
		Arguments map[string]any `json:"args"`
	}

	// RequestOption enables request customization.
	RequestOption func(*Request)

	// Credentials holds request plain auth credentials.
	Credentials struct{ Username, Password string }
)

// NewEvalRequest returns a new evaluation request request.
func NewEvalRequest(query string, opts ...RequestOption) *Request {
	r := &Request{
		RequestID: uuid.New().String(),
		Operation: OpsEval,
		Arguments: map[string]any{
			ArgsGremlin:  query,
			ArgsLanguage: "gremlin-groovy",
		},
	}
	for i := range opts {
		opts[i](r)
	}
	return r
}

// NewAuthRequest returns a new auth request.
func NewAuthRequest(requestID, username, password string) *Request {
	return &Request{
		RequestID: requestID,
		Operation: OpsAuthentication,
		Arguments: map[string]any{
			ArgsSasl: Credentials{
				Username: username,
				Password: password,
			},
			ArgsSaslMechanism: "PLAIN",
		},
	}
}

// WithBindings sets request bindings.
func WithBindings(bindings map[string]any) RequestOption {
	return func(r *Request) {
		r.Arguments[ArgsBindings] = bindings
	}
}

// WithEvalTimeout sets script evaluation timeout.
func WithEvalTimeout(timeout time.Duration) RequestOption {
	return func(r *Request) {
		r.Arguments[ArgsEvalTimeout] = int64(timeout / time.Millisecond)
	}
}

// MarshalText implements encoding.TextMarshaler interface.
func (c Credentials) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte(0)
	buf.WriteString(c.Username)
	buf.WriteByte(0)
	buf.WriteString(c.Password)

	enc := base64.StdEncoding
	text := make([]byte, enc.EncodedLen(buf.Len()))
	enc.Encode(text, buf.Bytes())
	return text, nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (c *Credentials) UnmarshalText(text []byte) error {
	enc := base64.StdEncoding
	data := make([]byte, enc.DecodedLen(len(text)))

	n, err := enc.Decode(data, text)
	if err != nil {
		return err
	}
	data = data[:n]

	parts := bytes.SplitN(data, []byte{0}, 3)
	if len(parts) != 3 {
		return errors.New("bad credentials data")
	}

	c.Username = string(parts[1])
	c.Password = string(parts[2])
	return nil
}
