// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import "testing"

func Test_receiver(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{
			"[]T",
			"t",
		},
		{
			"[1]T",
			"t",
		},
		{
			"User",
			"u",
		},
		{
			"UserQuery",
			"uq",
		},
		{
			"UserSelect",
			"us",
		},
		{
			"OSQuery",
			"oq",
		},
		{
			"OSSelect",
			"os",
		},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := receiver(tt.arg); got != tt.want {
				t.Errorf("receiver() = %v, want %v", got, tt.want)
			}
		})
	}
}
