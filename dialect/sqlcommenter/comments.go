package sqlcommenter

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

const (
	DbDriverCommentKey    = CommentKey("db_driver")
	FrameworkCommentKey   = CommentKey("framework")
	ApplicationCommentKey = CommentKey("application")
	RouteCommentKey       = CommentKey("route")
	ControllerCommentKey  = CommentKey("controller")
	ActionCommentKey      = CommentKey("action")
)

type (
	CommentKey   string
	CommentValue string
	SqlComments  map[CommentKey]CommentValue

	// used for sort
	Comment struct {
		CommentKey   CommentKey
		CommentValue CommentValue
	}
	Comments []Comment
)

func (c Comments) Len() int           { return len(c) }
func (c Comments) Less(i, j int) bool { return c[i].CommentKey < c[j].CommentKey }
func (c Comments) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (v CommentValue) marshal() string {
	urlEscape := strings.ReplaceAll(url.PathEscape(string(v)), "+", "%20")
	return fmt.Sprintf("'%s'", urlEscape)
}

func (v CommentKey) marshal() string {
	return url.QueryEscape(string(v))
}

func (sc SqlComments) Marshal() string {
	cmts := make(Comments, 0, len(sc))
	for k, v := range sc {
		cmts = append(cmts, Comment{k, v})
	}
	sort.Sort(cmts)
	var b strings.Builder
	for _, c := range cmts {
		fmt.Fprintf(&b, "%s=%s,", c.CommentKey.marshal(), c.CommentValue.marshal())
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
