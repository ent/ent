package graphson

import (
	"strings"
)

type tagOptions string

// parseTag splits a struct field's graphson tag into its type and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, ""
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (opts tagOptions) Contains(opt string) bool {
	s := string(opts)
	for s != "" {
		var next string
		i := strings.Index(s, ",")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}
		if s == opt {
			return true
		}
		s = next
	}
	return false
}
