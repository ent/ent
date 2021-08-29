package sqlcommenter

import (
	"fmt"
	"net/url"
	"strings"
)

type (
	CommentKey   string
	CommentValue string
	SqlComments  map[CommentKey]CommentValue
)

func (sc SqlComments) Marshal() string {
	var b strings.Builder
	for k, v := range sc {
		fmt.Fprintf(&b, "%s=%s,", url.PathEscape(string(k)), url.PathEscape(string(v)))
	}
	s := b.String()
	// remove trailing ","
	return s[:b.Len()-1]
}

func (sc SqlComments) Append(comments ...SqlComments) SqlComments {
	for _, c := range comments {
		for k, v := range c {
			sc[k] = v
		}
	}
	return sc
}
