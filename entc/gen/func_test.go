// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"testing"
)

func Test_receiver(t *testing.T) {
	tests := []struct {
		arg  string
		want string
	}{
		{
			"[]T",
			"_t",
		},
		{
			"[1]T",
			"_t",
		},
		{
			"User",
			"_u",
		},
		{
			"UserQuery",
			"_uq",
		},
		{
			"UserSelect",
			"_us",
		},
		{
			"OS",
			"_o",
		},
		{
			"OSQuery",
			"_oq",
		},
		{
			"OSSelect",
			"_os",
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

func Test_snake(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		arg  string
		want string
	}{
		{
			"Username",
			"username",
		},
		{
			"FullName",
			"full_name",
		},
		{
			"HTTPCode",
			"http_code",
		},
		{
			"OS",
			"os",
		},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := snake(tt.arg); got != tt.want {
				t.Errorf("snake() = %v, want %v", got, tt.want)
			}
		})
	}
}
