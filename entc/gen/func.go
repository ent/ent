// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"bytes"
	"errors"
	"fmt"
	"go/token"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"entgo.io/ent/schema/field"

	"github.com/go-openapi/inflect"
)

var (
	// Funcs are the predefined template
	// functions used by the codegen.
	Funcs = template.FuncMap{
		"ops":           fieldOps,
		"add":           add,
		"append":        reflect.Append,
		"appends":       reflect.AppendSlice,
		"order":         order,
		"camel":         camel,
		"snake":         snake,
		"pascal":        pascal,
		"extend":        extend,
		"xrange":        xrange,
		"receiver":      receiver,
		"plural":        plural,
		"aggregate":     aggregate,
		"primitives":    primitives,
		"singular":      rules.Singularize,
		"quote":         quote,
		"base":          filepath.Base,
		"keys":          keys,
		"indexOf":       indexOf,
		"join":          join,
		"joinWords":     joinWords,
		"isNil":         isNil,
		"lower":         strings.ToLower,
		"upper":         strings.ToUpper,
		"trim":          strings.Trim,
		"hasField":      hasField,
		"hasImport":     hasImport,
		"indirect":      indirect,
		"hasPrefix":     strings.HasPrefix,
		"hasSuffix":     strings.HasSuffix,
		"trimPackage":   trimPackage,
		"xtemplate":     xtemplate,
		"hasTemplate":   hasTemplate,
		"matchTemplate": matchTemplate,
		"split":         strings.Split,
		"tagLookup":     tagLookup,
		"toString":      toString,
		"dict":          dict,
		"get":           get,
		"set":           set,
		"unset":         unset,
		"hasKey":        hasKey,
		"list":          list[any],
		"slist":         list[string],
		"fail":          fail,
		"replace":       strings.ReplaceAll,
	}
	rules    = ruleset()
	acronyms = make(map[string]struct{})
)

// joinWords with spaces and add linebreaks to ensure lines do not exceed the given maxSize.
func joinWords(words []string, maxSize int) string {
	if len(words) == 0 {
		return ""
	}
	b := &strings.Builder{}
	b.WriteString(words[0])
	n := len(words[0])
	for _, w := range words[1:] {
		if n+len(w)+1 > maxSize {
			b.WriteByte('\n')
			n = 0
		}
		b.WriteString(" ")
		b.WriteString(w)
		n += len(w) + 1
	}
	return b.String()
}

// indexOf returns the index of the given string in the slice.
func indexOf(s []string, v string) int {
	for i, x := range s {
		if x == v {
			return i
		}
	}
	return -1
}

// quote only strings.
func quote(v any) any {
	if s, ok := v.(string); ok {
		return strconv.Quote(s)
	}
	return v
}

// fieldOps returns all predicate operations for a given field.
func fieldOps(f *Field) (ops []Op) {
	switch t := f.Type.Type; {
	case f.HasGoType() && !f.ConvertedToBasic() && !f.Type.Valuer():
	case t == field.TypeJSON:
	case t == field.TypeBool:
		ops = boolOps
	case t == field.TypeString && strings.ToLower(f.Name) != "id":
		ops = stringOps
		if f.HasGoType() && !f.ConvertedToBasic() && (f.Type.Valuer() || f.HasValueScanner()) {
			ops = numericOps
		}
	case t == field.TypeEnum || f.IsEdgeField():
		ops = enumOps
	default:
		ops = numericOps
	}
	if f.Optional {
		ops = append(ops, nillableOps...)
	}
	return ops
}

// xrange generates a slice of len n.
func xrange(n int) (a []int) {
	for i := 0; i < n; i++ {
		a = append(a, i)
	}
	return
}

// plural a name.
func plural(name string) string {
	p := rules.Pluralize(name)
	if p == name {
		p += "Slice"
	}
	return p
}

func isSeparator(r rune) bool {
	return r == '_' || r == '-' || unicode.IsSpace(r)
}

func pascalWords(words []string) string {
	for i, w := range words {
		upper := strings.ToUpper(w)
		if _, ok := acronyms[upper]; ok {
			words[i] = upper
		} else {
			words[i] = rules.Capitalize(w)
		}
	}
	return strings.Join(words, "")
}

// pascal converts the given name into a PascalCase.
//
//	user_info 	=> UserInfo
//	full_name 	=> FullName
//	user_id   	=> UserID
//	full-admin	=> FullAdmin
func pascal(s string) string {
	words := strings.FieldsFunc(s, isSeparator)
	return pascalWords(words)
}

// camel converts the given name into a camelCase.
//
//	user_info  => userInfo
//	full_name  => fullName
//	user_id    => userID
//	full-admin => fullAdmin
func camel(s string) string {
	words := strings.FieldsFunc(s, isSeparator)
	if len(words) == 1 {
		return strings.ToLower(words[0])
	}
	return strings.ToLower(words[0]) + pascalWords(words[1:])
}

// snake converts the given struct or field name into a snake_case.
//
//	Username => username
//	FullName => full_name
//	HTTPCode => http_code
func snake(s string) string {
	var (
		j int
		b strings.Builder
	)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		// Put '_' if it is not a start or end of a word, current letter is uppercase,
		// and previous is lowercase (cases like: "UserInfo"), or next letter is also
		// a lowercase and previous letter is not "_".
		if i > 0 && i < len(s)-1 && unicode.IsUpper(r) {
			if unicode.IsLower(rune(s[i-1])) ||
				j != i-1 && unicode.IsLower(rune(s[i+1])) && unicode.IsLetter(rune(s[i-1])) {
				j = i
				b.WriteString("_")
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

// receiver returns the receiver name of the given type.
//
//	[]T       => t
//	[1]T      => t
//	User      => u
//	UserQuery => uq
func receiver(s string) (r string) {
	// Trim invalid tokens for identifier prefix.
	s = strings.Trim(s, "[]*&0123456789")
	parts := strings.Split(snake(s), "_")
	min := len(parts[0])
	for _, w := range parts[1:] {
		if len(w) < min {
			min = len(w)
		}
	}
	for i := 1; i < min; i++ {
		r := parts[0][:i]
		for _, w := range parts[1:] {
			r += w[:i]
		}
		if _, ok := importPkg[r]; !ok {
			s = r
			break
		}
	}
	name := strings.ToLower(s)
	if token.Lookup(name).IsKeyword() {
		name = "_" + name
	}
	return name
}

// typeScope wraps the Type object with extended scope.
type typeScope struct {
	*Type
	Scope map[any]any
}

// graphScope wraps the Graph object with extended scope.
type graphScope struct {
	*Graph
	Scope map[any]any
}

// extend extends the parent block with a KV pairs.
//
//	{{ with $scope := extend $ "key" "value" }}
//		{{ template "setters" $scope }}
//	{{ end}}
func extend(v any, kv ...any) (any, error) {
	if len(kv)%2 != 0 {
		return nil, fmt.Errorf("invalid number of parameters: %d", len(kv))
	}
	scope := make(map[any]any, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		scope[kv[i]] = kv[i+1]
	}
	switch v := v.(type) {
	case *Type:
		return &typeScope{Type: v, Scope: scope}, nil
	case *Graph:
		return &graphScope{Graph: v, Scope: scope}, nil
	case *typeScope:
		for k := range v.Scope {
			scope[k] = v.Scope[k]
		}
		return &typeScope{Type: v.Type, Scope: scope}, nil
	case *graphScope:
		for k := range v.Scope {
			scope[k] = v.Scope[k]
		}
		return &graphScope{Graph: v.Graph, Scope: scope}, nil
	default:
		return nil, fmt.Errorf("invalid type for extend: %T", v)
	}
}

// add calculates summarize list of variables.
func add(xs ...int) (n int) {
	for _, x := range xs {
		n += x
	}
	return
}

func ruleset() *inflect.Ruleset {
	rules := inflect.NewDefaultRuleset()
	// Add common initialism from golint and more.
	for _, w := range []string{
		"ACL", "API", "ASCII", "AWS", "CPU", "CSS", "DNS", "EOF", "GB", "GUID",
		"HCL", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "KB", "LHS", "MAC",
		"MB", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "SSO",
		"TCP", "TLS", "TTL", "UDP", "UI", "UID", "URI", "URL", "UTF8", "UUID",
		"VM", "XML", "XMPP", "XSRF", "XSS",
	} {
		acronyms[w] = struct{}{}
		rules.AddAcronym(w)
	}
	return rules
}

// AddAcronym adds initialism to the global ruleset.
func AddAcronym(word string) {
	acronyms[word] = struct{}{}
	rules.AddAcronym(word)
}

// order returns a map of sort orders.
// The key is the function name, and the value its database keyword.
func order() map[string]string {
	return map[string]string{
		"asc":  "incr",
		"desc": "decr",
	}
}

// aggregate returns a map between all agg-functions and if they accept a field name as a parameter or not.
func aggregate() map[string]bool {
	return map[string]bool{
		"min":   true,
		"max":   true,
		"sum":   true,
		"mean":  true,
		"count": false,
	}
}

// keys returns the given map keys.
func keys(v reflect.Value) ([]string, error) {
	v = indirect(v)
	if k := v.Type().Kind(); k != reflect.Map {
		return nil, fmt.Errorf("expect map for keys, got: %s", k)
	}
	keys := make([]string, v.Len())
	for i, v := range v.MapKeys() {
		keys[i] = v.String()
	}
	sort.Strings(keys)
	return keys, nil
}

// primitives returns all primitives types.
func primitives() []string {
	return []string{field.TypeString.String(), field.TypeInt.String(), field.TypeFloat64.String(), field.TypeBool.String()}
}

// join is a wrapper around strings.Join to provide consistent output.
func join(a []string, sep string) string {
	sort.Strings(a)
	return strings.Join(a, sep)
}

// xtemplate dynamically executes templates by their names.
func xtemplate(name string, v any) (string, error) {
	buf := bytes.NewBuffer(nil)
	if err := templates.ExecuteTemplate(buf, name, v); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// hasTemplate checks whether a template exists in the loaded templates.
func hasTemplate(name string) bool {
	for _, t := range templates.Templates() {
		if t.Name() == name {
			return true
		}
	}
	return false
}

// matchTemplate returns all template names that match the given patterns.
func matchTemplate(patterns ...string) []string {
	var (
		names  []string
		exists = make(map[string]struct{})
	)
	for _, pattern := range patterns {
		for _, t := range templates.Templates() {
			name := t.Name()
			if _, ok := exists[name]; ok {
				continue
			}
			if match, _ := filepath.Match(pattern, name); match {
				names = append(names, name)
				exists[name] = struct{}{}
			}
		}
	}
	sort.Strings(names)
	return names
}

// hasField determines if a struct has a field with the given name.
func hasField(v any, name string) bool {
	vr := reflect.Indirect(reflect.ValueOf(v))
	return vr.FieldByName(name).IsValid()
}

// hasImport reports if the package name exists in the predefined import packages.
func hasImport(name string) bool {
	_, ok := importPkg[name]
	return ok
}

// trimPackage trims the package name from the given identifier.
func trimPackage(ident, pkg string) string {
	return strings.TrimPrefix(ident, pkg+".")
}

// isNil reports whether its argument v is nil.
func isNil(v any) bool {
	rv := indirect(reflect.ValueOf(v))
	if !rv.IsValid() {
		return true
	}
	switch rv.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

// indirect returns the item at the end of indirection.
func indirect(v reflect.Value) reflect.Value {
	for ; v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface; v = v.Elem() {
	}
	return v
}

// tagLookup returns the value associated with key in the tag string.
func tagLookup(tag, key string) string {
	v, _ := reflect.StructTag(tag).Lookup(key)
	return v
}

// toString converts `v` to a string.
func toString(v any) string {
	switch v := v.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprint(v)
	}
}

// dict creates a dictionary from a list of pairs.
func dict(v ...any) map[string]any {
	lenv := len(v)
	dict := make(map[string]any, lenv/2)
	for i := 0; i < lenv; i += 2 {
		key := toString(v[i])
		if i+1 >= lenv {
			dict[key] = ""
			continue
		}
		dict[key] = v[i+1]
	}
	return dict
}

// get the value from the dict for key.
func get(d map[string]any, key string) any {
	if val, ok := d[key]; ok {
		return val
	}
	return ""
}

// set adds a value to the dict for key.
func set(d map[string]any, key string, value any) map[string]any {
	d[key] = value
	return d
}

// unset removes a key from the dict.
func unset(d map[string]any, key string) map[string]any {
	delete(d, key)
	return d
}

// hasKey tests whether a key is found in dict.
func hasKey(d map[string]any, key string) bool {
	_, ok := d[key]
	return ok
}

// list creates a list from values.
func list[T any](v ...T) []T {
	return v
}

// fail unconditionally returns an empty string and an error with the specified text.
func fail(msg string) (string, error) {
	return "", errors.New(msg)
}
