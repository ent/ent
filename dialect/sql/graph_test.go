// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNeighbors(t *testing.T) {
	tests := []struct {
		name      string
		input     *Step
		wantQuery string
		wantArgs  []interface{}
	}{
		{
			name: "O2M/2types",
			input: func() *Step {
				step := &Step{}
				step.From.V = 1
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "pets"
				step.To.Column = "id"
				step.Edge.Rel = O2M
				step.Edge.Table = "pets"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `pets` WHERE `owner_id` = ?",
			wantArgs:  []interface{}{1},
		},
		{
			name: "O2M/1type",
			input: func() *Step {
				step := &Step{}
				step.From.V = 1
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = O2M
				step.Edge.Table = "users"
				step.Edge.Columns = []string{"parent_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` WHERE `parent_id` = ?",
			wantArgs:  []interface{}{1},
		},
		{
			name: "M2O/2types/inverse",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "pets"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = M2O
				step.Edge.Inverse = true
				step.Edge.Table = "pets"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `owner_id` FROM `pets` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`owner_id`",
			wantArgs:  []interface{}{2},
		},
		{
			name: "M2O/1type/inverse",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = M2O
				step.Edge.Inverse = true
				step.Edge.Table = "users"
				step.Edge.Columns = []string{"parent_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `parent_id` FROM `users` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`parent_id`",
			wantArgs:  []interface{}{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selector := Neighbors("", tt.input)
			query, args := selector.Query()
			require.Equal(t, tt.wantQuery, query)
			require.Equal(t, tt.wantArgs, args)
		})
	}
}
