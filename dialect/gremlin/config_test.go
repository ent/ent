// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestConfigParsing(t *testing.T) {
	var cfg Config
	_, err := flags.ParseArgs(&cfg, []string{
		"--disable-expansion",
		"--endpoint", "http://localhost:8182/gremlin",
	})
	assert.NoError(t, err)
	assert.True(t, cfg.DisableExpansion)
	assert.Equal(t, "http", cfg.Endpoint.Scheme)
	assert.Equal(t, "http://localhost:8182/gremlin", cfg.Endpoint.String())

	cfg = Config{}
	_, err = flags.ParseArgs(&cfg, nil)
	assert.NoError(t, err)
	assert.NotNil(t, cfg.Endpoint.URL)
}

func TestConfigBuild(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		opts    []Option
		wantErr bool
	}{
		{
			name: "HTTP",
			cfg: Config{
				Endpoint: Endpoint{
					URL: func() *url.URL {
						u, _ := url.Parse("http://gremlin-server/gremlin")
						return u
					}(),
				},
			},
		},
		{
			name: "NoScheme",
			cfg: Config{
				Endpoint: Endpoint{
					URL: &url.URL{},
				},
			},
			wantErr: true,
		},
		{
			name: "BadScheme",
			cfg: Config{
				Endpoint: Endpoint{
					URL: &url.URL{
						Scheme: "bad",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "WithOptions",
			cfg: Config{
				Endpoint: Endpoint{
					URL: func() *url.URL {
						u, _ := url.Parse("http://gremlin-server/gremlin")
						return u
					}(),
				},
				DisableExpansion: true,
			},
			opts: []Option{WithHTTPClient(&http.Client{})},
		},
		{
			name: "NoExpansion",
			cfg: Config{
				Endpoint: Endpoint{
					URL: &url.URL{
						Scheme: "bad",
					},
				},
				DisableExpansion: true,
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			client, err := tc.cfg.Build(tc.opts...)
			if !tc.wantErr {
				assert.NotNil(t, client)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

type testRoundTripper struct{ mock.Mock }

func (rt *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	args := rt.Called(req)
	rsp, _ := args.Get(0).(*http.Response)
	return rsp, args.Error(1)
}

func TestBuildWithHTTPClient(t *testing.T) {
	var transport testRoundTripper
	transport.On("RoundTrip", mock.Anything).
		Return(nil, errors.New("noop")).
		Once()
	defer transport.AssertExpectations(t)

	u, err := url.Parse("http://gremlin-server:8182/gremlin")
	require.NoError(t, err)

	client, err := Config{Endpoint: Endpoint{u}}.
		Build(WithHTTPClient(&http.Client{Transport: &transport}))
	require.NoError(t, err)
	_, _ = client.Do(context.Background(), NewEvalRequest("g.V()"))
}

func TestExpandOrdering(t *testing.T) {
	var cfg Config
	cfg.Endpoint.URL, _ = url.Parse("http://gremlin-server/gremlin")
	interceptor := func(RoundTripper) RoundTripper {
		return RoundTripperFunc(func(ctx context.Context, req *Request) (*Response, error) {
			assert.Equal(t, `g.V().hasLabel("user")`, req.Arguments[ArgsGremlin])
			assert.Nil(t, req.Arguments[ArgsBindings])
			return nil, errors.New("noop")
		})
	}
	c, err := cfg.Build(WithInterceptor(interceptor))
	require.NoError(t, err)
	req := NewEvalRequest("g.V().hasLabel($1)", WithBindings(map[string]any{"$1": "user"}))
	_, _ = c.Do(context.Background(), req)
}
