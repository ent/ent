// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"github.com/google/uuid"
)

// TweetTag is the model entity for the TweetTag schema.
type TweetTag struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AddedAt holds the value of the "added_at" field.
	AddedAt time.Time `json:"added_at,omitempty"`
	// TagID holds the value of the "tag_id" field.
	TagID int `json:"tag_id,omitempty"`
	// TweetID holds the value of the "tweet_id" field.
	TweetID int `json:"tweet_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TweetTagQuery when eager-loading is set.
	Edges TweetTagEdges `json:"edges"`
}

// TweetTagEdges holds the relations/edges for other nodes in the graph.
type TweetTagEdges struct {
	// Tag holds the value of the tag edge.
	Tag *Tag `json:"tag,omitempty"`
	// Tweet holds the value of the tweet edge.
	Tweet *Tweet `json:"tweet,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TweetTagEdges) TagOrErr() (*Tag, error) {
	if e.loadedTypes[0] {
		if e.Tag == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tag.Label}
		}
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// TweetOrErr returns the Tweet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TweetTagEdges) TweetOrErr() (*Tweet, error) {
	if e.loadedTypes[1] {
		if e.Tweet == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tweet.Label}
		}
		return e.Tweet, nil
	}
	return nil, &NotLoadedError{edge: "tweet"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TweetTag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tweettag.FieldTagID, tweettag.FieldTweetID:
			values[i] = new(sql.NullInt64)
		case tweettag.FieldAddedAt:
			values[i] = new(sql.NullTime)
		case tweettag.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TweetTag", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TweetTag fields.
func (tt *TweetTag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tweettag.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tt.ID = *value
			}
		case tweettag.FieldAddedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field added_at", values[i])
			} else if value.Valid {
				tt.AddedAt = value.Time
			}
		case tweettag.FieldTagID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tag_id", values[i])
			} else if value.Valid {
				tt.TagID = int(value.Int64)
			}
		case tweettag.FieldTweetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tweet_id", values[i])
			} else if value.Valid {
				tt.TweetID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTag queries the "tag" edge of the TweetTag entity.
func (tt *TweetTag) QueryTag() *TagQuery {
	return NewTweetTagClient(tt.config).QueryTag(tt)
}

// QueryTweet queries the "tweet" edge of the TweetTag entity.
func (tt *TweetTag) QueryTweet() *TweetQuery {
	return NewTweetTagClient(tt.config).QueryTweet(tt)
}

// Update returns a builder for updating this TweetTag.
// Note that you need to call TweetTag.Unwrap() before calling this method if this TweetTag
// was returned from a transaction, and the transaction was committed or rolled back.
func (tt *TweetTag) Update() *TweetTagUpdateOne {
	return NewTweetTagClient(tt.config).UpdateOne(tt)
}

// Unwrap unwraps the TweetTag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tt *TweetTag) Unwrap() *TweetTag {
	_tx, ok := tt.config.driver.(*txDriver)
	if !ok {
		panic("ent: TweetTag is not a transactional entity")
	}
	tt.config.driver = _tx.drv
	return tt
}

// String implements the fmt.Stringer.
func (tt *TweetTag) String() string {
	var builder strings.Builder
	builder.WriteString("TweetTag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tt.ID))
	builder.WriteString("added_at=")
	builder.WriteString(tt.AddedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("tag_id=")
	builder.WriteString(fmt.Sprintf("%v", tt.TagID))
	builder.WriteString(", ")
	builder.WriteString("tweet_id=")
	builder.WriteString(fmt.Sprintf("%v", tt.TweetID))
	builder.WriteByte(')')
	return builder.String()
}

// TweetTags is a parsable slice of TweetTag.
type TweetTags []*TweetTag

func (tt TweetTags) config(cfg config) {
	for _i := range tt {
		tt[_i].config = cfg
	}
}

func (tt TweetTags) IDs() []uuid.UUID {
	ids := make([]uuid.UUID, len(tt))
	for _i := range tt {
		ids[_i] = tt[_i].ID
	}
	return ids
}
