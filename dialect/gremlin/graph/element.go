// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graph

// Element defines a base struct for graph elements.
type Element struct {
	ID    any    `json:"id"`
	Label string `json:"label"`
}

// NewElement create a new graph element.
func NewElement(id any, label string) Element {
	return Element{id, label}
}
