// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package field

import (
	"go/token"
	"path"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// typeName renders generic instantiations as Go source. Reflection qualifies
// their type arguments with import paths, which need explicit import aliases.
func typeName(t reflect.Type) (string, map[string]string) {
	if !hasTypeArgs(t) {
		return t.String(), nil
	}
	n := &typeNamer{
		imports: make(map[string]string),
		aliases: make(map[string]bool),
	}
	ident := n.typeExpr(t)
	return ident, n.imports
}

func hasTypeArgs(t reflect.Type) bool {
	if name := t.Name(); name != "" {
		return strings.ContainsRune(name, '[')
	}
	switch t.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Array:
		return hasTypeArgs(t.Elem())
	case reflect.Map:
		return hasTypeArgs(t.Key()) || hasTypeArgs(t.Elem())
	default:
		return false
	}
}

type typeNamer struct {
	imports map[string]string
	aliases map[string]bool
}

func (n *typeNamer) typeExpr(t reflect.Type) string {
	if name := t.Name(); name != "" {
		var prefix string
		if p := t.PkgPath(); p != "" {
			prefix = n.importAlias(p, pkgName(t)) + "."
		}
		if base, args, ok := strings.Cut(name, "["); ok {
			return prefix + base + "[" + n.typeArgs(args)
		}
		return prefix + name
	}
	switch t.Kind() {
	case reflect.Ptr:
		return "*" + n.typeExpr(t.Elem())
	case reflect.Slice:
		return "[]" + n.typeExpr(t.Elem())
	case reflect.Array:
		return "[" + strconv.Itoa(t.Len()) + "]" + n.typeExpr(t.Elem())
	case reflect.Map:
		return "map[" + n.typeExpr(t.Key()) + "]" + n.typeExpr(t.Elem())
	default:
		return t.String()
	}
}

// typeArgs replaces qualified names, including those in nested type arguments,
// while preserving punctuation and quoted struct tags.
func (n *typeNamer) typeArgs(s string) string {
	var b strings.Builder
	for len(s) > 0 {
		if s[0] == '"' || s[0] == '`' || s[0] == '\'' {
			if quoted, err := strconv.QuotedPrefix(s); err == nil {
				b.WriteString(quoted)
				s = s[len(quoted):]
				continue
			}
		}
		r, size := utf8.DecodeRuneInString(s)
		if !unicode.IsLetter(r) && r != '_' {
			b.WriteString(s[:size])
			s = s[size:]
			continue
		}
		i := strings.IndexFunc(s, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r) && !strings.ContainsRune("_./-~", r)
		})
		if i == -1 {
			i = len(s)
		}
		word := s[:i]
		if dot := strings.LastIndexByte(word, '.'); dot != -1 {
			word = n.importAlias(word[:dot], "") + word[dot:]
		}
		b.WriteString(word)
		s = s[i:]
	}
	return b.String()
}

func (n *typeNamer) importAlias(pkg, name string) string {
	if alias, ok := n.imports[pkg]; ok {
		return alias
	}
	if name == "" {
		name = strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
				return r
			}
			return '_'
		}, path.Base(pkg))
		if !token.IsIdentifier(name) || name == "_" {
			name = "pkg_" + name
		}
	}
	alias := name
	for i := 2; n.aliases[alias]; i++ {
		alias = name + strconv.Itoa(i)
	}
	n.imports[pkg] = alias
	n.aliases[alias] = true
	return alias
}
