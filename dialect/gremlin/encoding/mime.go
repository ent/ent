// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package encoding

import (
	"bytes"
)

// Mime defines a gremlin mime type.
type Mime []byte

// GraphSON3Mime mime headers.
var (
	GraphSON3Mime = NewMime("application/vnd.gremlin-v3.0+json")
)

// NewMime creates a wire format mime header.
func NewMime(s string) Mime {
	var buf bytes.Buffer
	buf.WriteByte(byte(len(s)))
	buf.WriteString(s)
	return buf.Bytes()
}

// String implements fmt.Stringer interface.
func (m Mime) String() string {
	return string(m[1:])
}
