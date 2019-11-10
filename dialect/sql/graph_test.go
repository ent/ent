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
			name: "O2O/1type",
			input: func() *Step {
				step := &Step{}
				// Since the relation is on the same table,
				// V used as a reference value.
				step.From.V = 1
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Table = "users"
				step.Edge.Columns = []string{"spouse_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` WHERE `spouse_id` = ?",
			wantArgs:  []interface{}{1},
		},
		{
			name: "O2O/1type/inverse",
			input: func() *Step {
				step := &Step{}
				step.From.V = 1
				step.From.Table = "nodes"
				step.From.Column = "id"
				step.To.Table = "nodes"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Table = "nodes"
				step.Edge.Inverse = true
				step.Edge.Columns = []string{"prev_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `nodes` JOIN (SELECT `prev_id` FROM `nodes` WHERE `id` = ?) AS `t1` ON `nodes`.`id` = `t1`.`prev_id`",
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
			name: "O2O/2types",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "card"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Table = "cards"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `card` WHERE `owner_id` = ?",
			wantArgs:  []interface{}{2},
		},
		{
			name: "O2O/2types/inverse",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "card"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Table = "cards"
				step.Edge.Inverse = true
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `owner_id` FROM `cards` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`owner_id`",
			wantArgs:  []interface{}{2},
		},
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
		{
			name: "M2M/2type",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "groups"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = M2M
				step.Edge.Table = "user_groups"
				step.Edge.Columns = []string{"group_id", "user_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `user_groups`.`user_id` FROM `user_groups` WHERE `user_groups`.`group_id` = ?) AS `t1` ON `users`.`id` = `t1`.`user_id`",
			wantArgs:  []interface{}{2},
		},
		{
			name: "M2M/2type/inverse",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "groups"
				step.To.Column = "id"
				step.Edge.Rel = M2M
				step.Edge.Inverse = true
				step.Edge.Table = "user_groups"
				step.Edge.Columns = []string{"group_id", "user_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `groups` JOIN (SELECT `user_groups`.`group_id` FROM `user_groups` WHERE `user_groups`.`user_id` = ?) AS `t1` ON `groups`.`id` = `t1`.`group_id`",
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
