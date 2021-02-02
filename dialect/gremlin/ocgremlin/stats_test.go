// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package ocgremlin

import (
	"context"
	"strings"
	"testing"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/encoding/graphson"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opencensus.io/stats/view"
)

func TestStatsCollection(t *testing.T) {
	err := view.Register(
		RequestCountView,
		ResponseCountView,
		ResponseBytesView,
		RoundTripLatencyView,
	)
	require.NoError(t, err)

	req := gremlin.NewEvalRequest("g.E()")
	rsp := &gremlin.Response{RequestID: req.RequestID}
	rsp.Status.Code = gremlin.StatusSuccess
	rsp.Result.Data = graphson.RawMessage(
		`{"@type": "g:List", "@value": [{"@type": "g:Int32", "@value": 42}]}`,
	)

	transport := &mockTransport{}
	transport.On("RoundTrip", mock.Anything, mock.Anything).
		Return(rsp, nil).
		Once()
	defer transport.AssertExpectations(t)

	rt := &statsTransport{transport}
	_, _ = rt.RoundTrip(context.Background(), req)

	tests := []struct {
		name   string
		expect func(*testing.T, *view.Row)
	}{
		{
			name: "gremlin/request_count",
			expect: func(t *testing.T, row *view.Row) {
				count, ok := row.Data.(*view.CountData)
				require.True(t, ok)
				assert.Equal(t, int64(1), count.Value)
			},
		},
		{
			name: "gremlin/response_count",
			expect: func(t *testing.T, row *view.Row) {
				count, ok := row.Data.(*view.CountData)
				require.True(t, ok)
				assert.Equal(t, int64(1), count.Value)
			},
		},
		{
			name: "gremlin/response_bytes",
			expect: func(t *testing.T, row *view.Row) {
				data, ok := row.Data.(*view.DistributionData)
				require.True(t, ok)
				assert.EqualValues(t, len(rsp.Result.Data), data.Sum())
			},
		},
		{
			name: "gremlin/roundtrip_latency",
			expect: func(t *testing.T, row *view.Row) {
				data, ok := row.Data.(*view.DistributionData)
				require.True(t, ok)
				assert.NotZero(t, data.Sum())
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name[strings.Index(tt.name, "/")+1:], func(t *testing.T) {
			v := view.Find(tt.name)
			assert.NotNil(t, v)
			rows, err := view.RetrieveData(tt.name)
			require.NoError(t, err)
			require.Len(t, rows, 1)
			tt.expect(t, rows[0])
		})
	}
}
