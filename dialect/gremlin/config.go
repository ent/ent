// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"fmt"
	"net/http"
	"net/url"
)

type (
	// Config offers a declarative way to construct a client.
	Config struct {
		Endpoint         Endpoint `env:"ENDPOINT" long:"endpoint" default:"" description:"gremlin endpoint to connect to"`
		DisableExpansion bool     `env:"DISABLE_EXPANSION" long:"disable-expansion" description:"disable bindings expansion"`
	}

	// An Option configured client.
	Option func(*options)

	options struct {
		interceptors []Interceptor
		httpClient   *http.Client
	}

	// Endpoint wraps a url to add flag unmarshalling.
	Endpoint struct {
		*url.URL
	}
)

// WithInterceptor adds interceptors to the client's transport.
func WithInterceptor(interceptors ...Interceptor) Option {
	return func(opts *options) {
		opts.interceptors = append(opts.interceptors, interceptors...)
	}
}

// WithHTTPClient assigns underlying http client to be used by http transport.
func WithHTTPClient(client *http.Client) Option {
	return func(opts *options) {
		opts.httpClient = client
	}
}

// Build constructs a client from Config.
func (cfg Config) Build(opt ...Option) (c *Client, err error) {
	opts := cfg.buildOptions(opt)
	switch cfg.Endpoint.Scheme {
	case "http", "https":
		c, err = NewHTTPClient(cfg.Endpoint.String(), opts.httpClient)
	default:
		err = fmt.Errorf("unsupported endpoint scheme: %s", cfg.Endpoint.Scheme)
	}
	if err != nil {
		return nil, err
	}

	for i := len(opts.interceptors) - 1; i >= 0; i-- {
		c.Transport = opts.interceptors[i](c.Transport)
	}
	if !cfg.DisableExpansion {
		c.Transport = ExpandBindings(c.Transport)
	}
	return c, nil
}

func (Config) buildOptions(opts []Option) options {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

// UnmarshalFlag implements flag.Unmarshaler interface.
func (ep *Endpoint) UnmarshalFlag(value string) (err error) {
	ep.URL, err = url.Parse(value)
	return
}
