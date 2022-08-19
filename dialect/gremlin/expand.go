// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"context"
	"fmt"
	"sort"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

// ExpandBindings expands the given RoundTripper and expands the request bindings into the Gremlin traversal.
func ExpandBindings(rt RoundTripper) RoundTripper {
	return RoundTripperFunc(func(ctx context.Context, r *Request) (*Response, error) {
		bindings, ok := r.Arguments[ArgsBindings]
		if !ok {
			return rt.RoundTrip(ctx, r)
		}
		query, ok := r.Arguments[ArgsGremlin]
		if !ok {
			return rt.RoundTrip(ctx, r)
		}
		{
			query, bindings := query.(string), bindings.(map[string]any)
			keys := make(sort.StringSlice, 0, len(bindings))
			for k := range bindings {
				keys = append(keys, k)
			}
			sort.Sort(sort.Reverse(keys))
			kv := make([]string, 0, len(bindings)*2)
			for _, k := range keys {
				s, err := jsoniter.MarshalToString(bindings[k])
				if err != nil {
					return nil, fmt.Errorf("marshal bindings value for key %s: %w", k, err)
				}
				kv = append(kv, k, s)
			}
			delete(r.Arguments, ArgsBindings)
			r.Arguments[ArgsGremlin] = strings.NewReplacer(kv...).Replace(query)
		}
		return rt.RoundTrip(ctx, r)
	})
}
