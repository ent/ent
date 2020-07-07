// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/facebookincubator/ent/entc/integration/json/ent/user"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	url           **url.URL
	raw           *json.RawMessage
	dirs          *[]http.Dir
	ints          *[]int
	floats        *[]float64
	strings       *[]string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows to manage the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the id field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetURL sets the url field.
func (m *UserMutation) SetURL(u *url.URL) {
	m.url = &u
}

// URL returns the url value in the mutation.
func (m *UserMutation) URL() (r *url.URL, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old url value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldURL(ctx context.Context) (v *url.URL, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldURL is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ClearURL clears the value of url.
func (m *UserMutation) ClearURL() {
	m.url = nil
	m.clearedFields[user.FieldURL] = struct{}{}
}

// URLCleared returns if the field url was cleared in this mutation.
func (m *UserMutation) URLCleared() bool {
	_, ok := m.clearedFields[user.FieldURL]
	return ok
}

// ResetURL reset all changes of the "url" field.
func (m *UserMutation) ResetURL() {
	m.url = nil
	delete(m.clearedFields, user.FieldURL)
}

// SetRaw sets the raw field.
func (m *UserMutation) SetRaw(jm json.RawMessage) {
	m.raw = &jm
}

// Raw returns the raw value in the mutation.
func (m *UserMutation) Raw() (r json.RawMessage, exists bool) {
	v := m.raw
	if v == nil {
		return
	}
	return *v, true
}

// OldRaw returns the old raw value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldRaw(ctx context.Context) (v json.RawMessage, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRaw is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRaw requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRaw: %w", err)
	}
	return oldValue.Raw, nil
}

// ClearRaw clears the value of raw.
func (m *UserMutation) ClearRaw() {
	m.raw = nil
	m.clearedFields[user.FieldRaw] = struct{}{}
}

// RawCleared returns if the field raw was cleared in this mutation.
func (m *UserMutation) RawCleared() bool {
	_, ok := m.clearedFields[user.FieldRaw]
	return ok
}

// ResetRaw reset all changes of the "raw" field.
func (m *UserMutation) ResetRaw() {
	m.raw = nil
	delete(m.clearedFields, user.FieldRaw)
}

// SetDirs sets the dirs field.
func (m *UserMutation) SetDirs(h []http.Dir) {
	m.dirs = &h
}

// Dirs returns the dirs value in the mutation.
func (m *UserMutation) Dirs() (r []http.Dir, exists bool) {
	v := m.dirs
	if v == nil {
		return
	}
	return *v, true
}

// OldDirs returns the old dirs value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldDirs(ctx context.Context) (v []http.Dir, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDirs is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDirs requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDirs: %w", err)
	}
	return oldValue.Dirs, nil
}

// ClearDirs clears the value of dirs.
func (m *UserMutation) ClearDirs() {
	m.dirs = nil
	m.clearedFields[user.FieldDirs] = struct{}{}
}

// DirsCleared returns if the field dirs was cleared in this mutation.
func (m *UserMutation) DirsCleared() bool {
	_, ok := m.clearedFields[user.FieldDirs]
	return ok
}

// ResetDirs reset all changes of the "dirs" field.
func (m *UserMutation) ResetDirs() {
	m.dirs = nil
	delete(m.clearedFields, user.FieldDirs)
}

// SetInts sets the ints field.
func (m *UserMutation) SetInts(i []int) {
	m.ints = &i
}

// Ints returns the ints value in the mutation.
func (m *UserMutation) Ints() (r []int, exists bool) {
	v := m.ints
	if v == nil {
		return
	}
	return *v, true
}

// OldInts returns the old ints value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldInts(ctx context.Context) (v []int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldInts is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldInts requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInts: %w", err)
	}
	return oldValue.Ints, nil
}

// ClearInts clears the value of ints.
func (m *UserMutation) ClearInts() {
	m.ints = nil
	m.clearedFields[user.FieldInts] = struct{}{}
}

// IntsCleared returns if the field ints was cleared in this mutation.
func (m *UserMutation) IntsCleared() bool {
	_, ok := m.clearedFields[user.FieldInts]
	return ok
}

// ResetInts reset all changes of the "ints" field.
func (m *UserMutation) ResetInts() {
	m.ints = nil
	delete(m.clearedFields, user.FieldInts)
}

// SetFloats sets the floats field.
func (m *UserMutation) SetFloats(f []float64) {
	m.floats = &f
}

// Floats returns the floats value in the mutation.
func (m *UserMutation) Floats() (r []float64, exists bool) {
	v := m.floats
	if v == nil {
		return
	}
	return *v, true
}

// OldFloats returns the old floats value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldFloats(ctx context.Context) (v []float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldFloats is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldFloats requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFloats: %w", err)
	}
	return oldValue.Floats, nil
}

// ClearFloats clears the value of floats.
func (m *UserMutation) ClearFloats() {
	m.floats = nil
	m.clearedFields[user.FieldFloats] = struct{}{}
}

// FloatsCleared returns if the field floats was cleared in this mutation.
func (m *UserMutation) FloatsCleared() bool {
	_, ok := m.clearedFields[user.FieldFloats]
	return ok
}

// ResetFloats reset all changes of the "floats" field.
func (m *UserMutation) ResetFloats() {
	m.floats = nil
	delete(m.clearedFields, user.FieldFloats)
}

// SetStrings sets the strings field.
func (m *UserMutation) SetStrings(s []string) {
	m.strings = &s
}

// Strings returns the strings value in the mutation.
func (m *UserMutation) Strings() (r []string, exists bool) {
	v := m.strings
	if v == nil {
		return
	}
	return *v, true
}

// OldStrings returns the old strings value of the User.
// If the User object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *UserMutation) OldStrings(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStrings is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStrings requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStrings: %w", err)
	}
	return oldValue.Strings, nil
}

// ClearStrings clears the value of strings.
func (m *UserMutation) ClearStrings() {
	m.strings = nil
	m.clearedFields[user.FieldStrings] = struct{}{}
}

// StringsCleared returns if the field strings was cleared in this mutation.
func (m *UserMutation) StringsCleared() bool {
	_, ok := m.clearedFields[user.FieldStrings]
	return ok
}

// ResetStrings reset all changes of the "strings" field.
func (m *UserMutation) ResetStrings() {
	m.strings = nil
	delete(m.clearedFields, user.FieldStrings)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.url != nil {
		fields = append(fields, user.FieldURL)
	}
	if m.raw != nil {
		fields = append(fields, user.FieldRaw)
	}
	if m.dirs != nil {
		fields = append(fields, user.FieldDirs)
	}
	if m.ints != nil {
		fields = append(fields, user.FieldInts)
	}
	if m.floats != nil {
		fields = append(fields, user.FieldFloats)
	}
	if m.strings != nil {
		fields = append(fields, user.FieldStrings)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldURL:
		return m.URL()
	case user.FieldRaw:
		return m.Raw()
	case user.FieldDirs:
		return m.Dirs()
	case user.FieldInts:
		return m.Ints()
	case user.FieldFloats:
		return m.Floats()
	case user.FieldStrings:
		return m.Strings()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldURL:
		return m.OldURL(ctx)
	case user.FieldRaw:
		return m.OldRaw(ctx)
	case user.FieldDirs:
		return m.OldDirs(ctx)
	case user.FieldInts:
		return m.OldInts(ctx)
	case user.FieldFloats:
		return m.OldFloats(ctx)
	case user.FieldStrings:
		return m.OldStrings(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldURL:
		v, ok := value.(*url.URL)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case user.FieldRaw:
		v, ok := value.(json.RawMessage)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRaw(v)
		return nil
	case user.FieldDirs:
		v, ok := value.([]http.Dir)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDirs(v)
		return nil
	case user.FieldInts:
		v, ok := value.([]int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInts(v)
		return nil
	case user.FieldFloats:
		v, ok := value.([]float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFloats(v)
		return nil
	case user.FieldStrings:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStrings(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldURL) {
		fields = append(fields, user.FieldURL)
	}
	if m.FieldCleared(user.FieldRaw) {
		fields = append(fields, user.FieldRaw)
	}
	if m.FieldCleared(user.FieldDirs) {
		fields = append(fields, user.FieldDirs)
	}
	if m.FieldCleared(user.FieldInts) {
		fields = append(fields, user.FieldInts)
	}
	if m.FieldCleared(user.FieldFloats) {
		fields = append(fields, user.FieldFloats)
	}
	if m.FieldCleared(user.FieldStrings) {
		fields = append(fields, user.FieldStrings)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldURL:
		m.ClearURL()
		return nil
	case user.FieldRaw:
		m.ClearRaw()
		return nil
	case user.FieldDirs:
		m.ClearDirs()
		return nil
	case user.FieldInts:
		m.ClearInts()
		return nil
	case user.FieldFloats:
		m.ClearFloats()
		return nil
	case user.FieldStrings:
		m.ClearStrings()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldURL:
		m.ResetURL()
		return nil
	case user.FieldRaw:
		m.ResetRaw()
		return nil
	case user.FieldDirs:
		m.ResetDirs()
		return nil
	case user.FieldInts:
		m.ResetInts()
		return nil
	case user.FieldFloats:
		m.ResetFloats()
		return nil
	case user.FieldStrings:
		m.ResetStrings()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
