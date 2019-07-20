package encoding

import (
	"bytes"
)

// Mime defines a gremlin mime type.
type Mime []byte

// Graphson mime headers.
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
