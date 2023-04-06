// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"reflect"
	"testing"

	"entgo.io/ent/dialect/gremlin/encoding/graphson"
	"entgo.io/ent/dialect/gremlin/graph"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeResponse(t *testing.T) {
	in := `{
		"requestId": "a65f2d39-1efa-45d2-a06a-c736476500fc",
		"result": {
			"data": {
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Map",
						"@value": [
							{
								"@type": "g:T",
								"@value": "id"
							},
							{
								"@type": "g:Int64",
								"@value": 1
							},
							{
								"@type": "g:T",
								"@value": "label"
							},
							"person",
							"name",
							{
								"@type": "g:List",
								"@value": [
									"marko"
								]
							},
							"age",
							{
								"@type": "g:List",
								"@value": [
									{
										"@type": "g:Int32",
										"@value": 29
									}
								]
							}
						]
					},
					{
						"@type": "g:Map",
						"@value": [
							{
								"@type": "g:T",
								"@value": "id"
							},
							{
								"@type": "g:Int64",
								"@value": 6
							},
							{
								"@type": "g:T",
								"@value": "label"
							},
							"person",
							"name",
							{
								"@type": "g:List",
								"@value": [
									"peter"
								]
							},
							"age",
							{
								"@type": "g:List",
								"@value": [
									{
										"@type": "g:Int32",
										"@value": 35
									}
								]
							}
						]
					}
				]
			},
			"meta": {
				"@type": "g:Map",
				"@value": []
			}
		},
		"status": {
			"attributes": {
				"@type": "g:Map",
				"@value": []
			},
			"code": 200,
			"message": ""
		}
	}`

	var rsp Response
	err := graphson.UnmarshalFromString(in, &rsp)
	require.NoError(t, err)

	assert.Equal(t, "a65f2d39-1efa-45d2-a06a-c736476500fc", rsp.RequestID)
	assert.Equal(t, 200, rsp.Status.Code)
	assert.Empty(t, rsp.Status.Message)
	assert.Empty(t, rsp.Status.Attributes)
	assert.Empty(t, rsp.Result.Meta)

	var vm graph.ValueMap
	err = graphson.Unmarshal(rsp.Result.Data, &vm)
	require.NoError(t, err)
	require.Len(t, vm, 2)

	type person struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var people []person
	err = vm.Decode(&people)
	require.NoError(t, err)
	assert.Equal(t, []person{
		{1, "marko", 29},
		{6, "peter", 35},
	}, people)
}

func TestDecodeResponseWithError(t *testing.T) {
	in := `{
		"requestId": "41d2e28a-20a4-4ab0-b379-d810dede3786",
		"result": {
			"data": null,
			"meta": {
				"@type": "g:Map",
				"@value": []
			}
		},
		"status": {
			"attributes": {
				"@type": "g:Map",
				"@value": []
			},
			"code": 500,
			"message": "Database Down"
		}
	}`
	var rsp Response
	err := graphson.UnmarshalFromString(in, &rsp)
	require.NoError(t, err)

	err = rsp.Err()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Database Down")

	rsp = Response{}
	err = graphson.UnmarshalFromString(`{"status": null}`, &rsp)
	require.NoError(t, err)
	assert.Error(t, rsp.Err())
}

func TestResponseReadVal(t *testing.T) {
	var rsp Response
	rsp.Status.Code = StatusSuccess
	rsp.Result.Data = []byte(`{"@type": "g:Int32", "@value": 15}`)

	var v int32
	err := rsp.ReadVal(&v)
	assert.NoError(t, err)
	assert.Equal(t, int32(15), v)

	var s string
	err = rsp.ReadVal(&s)
	assert.Error(t, err)

	rsp.Status.Code = StatusServerError
	err = rsp.ReadVal(&v)
	assert.Error(t, err)
}

func TestResponseReadGraphElements(t *testing.T) {
	tests := []struct {
		method string
		data   string
		want   any
	}{
		{
			method: "ReadVertices",
			data: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Vertex",
						"@value": {
							"id": {
								"@type": "g:Int64",
								"@value": 1
							},
							"label": "person"
						}
					},
					{
						"@type": "g:Vertex",
						"@value": {
							"id": {
								"@type": "g:Int64",
								"@value": 6
							},
							"label": "person"
						}
					}
				]
			}`,
			want: []graph.Vertex{
				graph.NewVertex(int64(1), "person"),
				graph.NewVertex(int64(6), "person"),
			},
		},
		{
			method: "ReadVertexProperties",
			data: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:VertexProperty",
						"@value": {
							"id": {
								"@type": "g:Int64",
								"@value": 0
							},
							"label": "name",
							"value": "marko"
						}
					},
					{
						"@type": "g:VertexProperty",
						"@value": {
							"id": {
								"@type": "g:Int64",
								"@value": 2
							},
							"label": "age",
							"value": {
								"@type": "g:Int32",
								"@value": 29
							}
						}
					}
				]
			}`,
			want: []graph.VertexProperty{
				graph.NewVertexProperty(int64(0), "name", "marko"),
				graph.NewVertexProperty(int64(2), "age", int32(29)),
			},
		},
		{
			method: "ReadEdges",
			data: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Edge",
						"@value": {
							"id": {
								"@type": "g:Int32",
								"@value": 12
							},
							"inV": {
								"@type": "g:Int64",
								"@value": 3
							},
							"inVLabel": "software",
							"label": "created",
							"outV": {
								"@type": "g:Int64",
								"@value": 6
							},
							"outVLabel": "person"
						}
					}
				]
			}`,
			want: []graph.Edge{
				graph.NewEdge(int32(12), "created",
					graph.NewVertex(int64(6), "person"),
					graph.NewVertex(int64(3), "software"),
				),
			},
		},
		{
			method: "ReadProperties",
			data: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Property",
						"@value": {
							"key": "weight",
							"value": {
								"@type": "g:Double",
								"@value": 0.2
							}
						}
					}
				]
			}`,
			want: []graph.Property{
				graph.NewProperty("weight", float64(0.2)),
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.method, func(t *testing.T) {
			t.Parallel()
			var rsp Response
			rsp.Status.Code = StatusSuccess
			rsp.Result.Data = []byte(tc.data)
			vals := reflect.ValueOf(&rsp).MethodByName(tc.method).Call(nil)
			require.Len(t, vals, 2)
			require.True(t, vals[1].IsNil())
			assert.Equal(t, tc.want, vals[0].Interface())
		})
	}
}

func TestResponseReadValueMap(t *testing.T) {
	t.Parallel()
	var rsp Response
	rsp.Status.Code = StatusSuccess
	rsp.Result.Data = []byte(`{
		"@type": "g:List",
		"@value": [
			{
				"@type": "g:Map",
				"@value": [
					"name",
					{
						"@type": "g:List",
						"@value": [
							"alex"
						]
					}
				]
			}
		]
	}`)
	m, err := rsp.ReadValueMap()
	require.NoError(t, err)

	var name string
	err = m.Decode(&struct {
		Name *string `json:"name"`
	}{&name})
	require.NoError(t, err)
	assert.Equal(t, "alex", name)
}

func TestResponseReadBool(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    bool
		wantErr bool
	}{
		{
			name: "Simple",
			data: `{
				"@type": "g:List",
				"@value": [
					true
				]
			}`,
			want: true,
		},
		{
			name: "Multi",
			data: `{
				"@type": "g:List",
				"@value": [
					false,
					true
				]
			}`,
			want: false,
		},
		{
			name: "Empty",
			data: `{
				"@type": "g:List",
				"@value": []
			}`,
			wantErr: true,
		},
		{
			name: "BadType",
			data: `{
				"@type": "g:List",
				"@value": [
					"user"
				]
			}`,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var rsp Response
			rsp.Status.Code = StatusSuccess
			rsp.Result.Data = []byte(tc.data)
			got, err := rsp.ReadBool()
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestResponseReadInt(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    int
		wantErr bool
	}{
		{
			name: "Simple",
			data: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Int64",
						"@value": 42
					}
				]
			}`,
			want: 42,
		},
		{
			name: "Multi",
			data: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Int64",
						"@value": 55
					},
					{
						"@type": "g:Int64",
						"@value": 13
					}
				]
			}`,
			want: 55,
		},
		{
			name: "Empty",
			data: `{
				"@type": "g:List",
				"@value": []
			}`,
			wantErr: true,
		},
		{
			name: "BadType",
			data: `{
				"@type": "g:List",
				"@value": [
					true
				]
			}`,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var rsp Response
			rsp.Status.Code = StatusSuccess
			rsp.Result.Data = []byte(tc.data)
			got, err := rsp.ReadInt()
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestResponseReadString(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    string
		wantErr bool
	}{
		{
			name: "Simple",
			data: `{
				"@type": "g:List",
				"@value": ["foo"]
			}`,
			want: "foo",
		},
		{
			name: "Empty",
			data: `{
				"@type": "g:List",
				"@value": []
			}`,
			wantErr: true,
		},
		{
			name: "BadType",
			data: `{
				"@type": "g:List",
				"@value": [
					true
				]
			}`,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var rsp Response
			rsp.Status.Code = StatusSuccess
			rsp.Result.Data = []byte(tc.data)
			got, err := rsp.ReadString()
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
