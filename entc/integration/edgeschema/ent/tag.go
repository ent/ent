// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
)

// Tag is the model entity for the Tag schema.
type Tag struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TagQuery when eager-loading is set.
	Edges        TagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TagEdges holds the relations/edges for other nodes in the graph.
type TagEdges struct {
	// Tweets holds the value of the tweets edge.
	Tweets []*Tweet `json:"tweets,omitempty"`
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// TweetTags holds the value of the tweet_tags edge.
	TweetTags []*TweetTag `json:"tweet_tags,omitempty"`
	// GroupTags holds the value of the group_tags edge.
	GroupTags []*GroupTag `json:"group_tags,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TweetsOrErr returns the Tweets value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) TweetsOrErr() ([]*Tweet, error) {
	if e.loadedTypes[0] {
		return e.Tweets, nil
	}
	return nil, &NotLoadedError{edge: "tweets"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[1] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// TweetTagsOrErr returns the TweetTags value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) TweetTagsOrErr() ([]*TweetTag, error) {
	if e.loadedTypes[2] {
		return e.TweetTags, nil
	}
	return nil, &NotLoadedError{edge: "tweet_tags"}
}

// GroupTagsOrErr returns the GroupTags value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) GroupTagsOrErr() ([]*GroupTag, error) {
	if e.loadedTypes[3] {
		return e.GroupTags, nil
	}
	return nil, &NotLoadedError{edge: "group_tags"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			values[i] = new(sql.NullInt64)
		case tag.FieldValue:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tag fields.
func (t *Tag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tag.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				t.Value = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Tag.
// This includes values selected through modifiers, order, etc.
func (t *Tag) GetValue(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryTweets queries the "tweets" edge of the Tag entity.
func (t *Tag) QueryTweets() *TweetQuery {
	return NewTagClient(t.config).QueryTweets(t)
}

// QueryGroups queries the "groups" edge of the Tag entity.
func (t *Tag) QueryGroups() *GroupQuery {
	return NewTagClient(t.config).QueryGroups(t)
}

// QueryTweetTags queries the "tweet_tags" edge of the Tag entity.
func (t *Tag) QueryTweetTags() *TweetTagQuery {
	return NewTagClient(t.config).QueryTweetTags(t)
}

// QueryGroupTags queries the "group_tags" edge of the Tag entity.
func (t *Tag) QueryGroupTags() *GroupTagQuery {
	return NewTagClient(t.config).QueryGroupTags(t)
}

// Update returns a builder for updating this Tag.
// Note that you need to call Tag.Unwrap() before calling this method if this Tag
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tag) Update() *TagUpdateOne {
	return NewTagClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tag) Unwrap() *Tag {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tag is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("value=")
	builder.WriteString(t.Value)
	builder.WriteByte(')')
	return builder.String()
}

// Tags is a parsable slice of Tag.
type Tags []*Tag

// Len returns length of Tags.
func (t Tags) Len() int { return len(t) }
