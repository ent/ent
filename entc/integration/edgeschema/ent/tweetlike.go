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
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweetlike"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
)

// TweetLike is the model entity for the TweetLike schema.
type TweetLike struct {
	config `json:"-"`
	// LikedAt holds the value of the "liked_at" field.
	LikedAt time.Time `json:"liked_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// TweetID holds the value of the "tweet_id" field.
	TweetID int `json:"tweet_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TweetLikeQuery when eager-loading is set.
	Edges TweetLikeEdges `json:"edges"`
}

// TweetLikeEdges holds the relations/edges for other nodes in the graph.
type TweetLikeEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Tweet holds the value of the tweet edge.
	Tweet *Tweet `json:"tweet,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TweetLikeEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// TweetOrErr returns the Tweet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TweetLikeEdges) TweetOrErr() (*Tweet, error) {
	if e.loadedTypes[1] {
		if e.Tweet == nil {
			// The edge tweet was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: tweet.Label}
		}
		return e.Tweet, nil
	}
	return nil, &NotLoadedError{edge: "tweet"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TweetLike) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tweetlike.FieldUserID, tweetlike.FieldTweetID:
			values[i] = new(sql.NullInt64)
		case tweetlike.FieldLikedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TweetLike", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TweetLike fields.
func (tl *TweetLike) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tweetlike.FieldLikedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field liked_at", values[i])
			} else if value.Valid {
				tl.LikedAt = value.Time
			}
		case tweetlike.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				tl.UserID = int(value.Int64)
			}
		case tweetlike.FieldTweetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tweet_id", values[i])
			} else if value.Valid {
				tl.TweetID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the TweetLike entity.
func (tl *TweetLike) QueryUser() *UserQuery {
	return (&TweetLikeClient{config: tl.config}).QueryUser(tl)
}

// QueryTweet queries the "tweet" edge of the TweetLike entity.
func (tl *TweetLike) QueryTweet() *TweetQuery {
	return (&TweetLikeClient{config: tl.config}).QueryTweet(tl)
}

// Update returns a builder for updating this TweetLike.
// Note that you need to call TweetLike.Unwrap() before calling this method if this TweetLike
// was returned from a transaction, and the transaction was committed or rolled back.
func (tl *TweetLike) Update() *TweetLikeUpdateOne {
	return (&TweetLikeClient{config: tl.config}).UpdateOne(tl)
}

// Unwrap unwraps the TweetLike entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tl *TweetLike) Unwrap() *TweetLike {
	_tx, ok := tl.config.driver.(*txDriver)
	if !ok {
		panic("ent: TweetLike is not a transactional entity")
	}
	tl.config.driver = _tx.drv
	return tl
}

// String implements the fmt.Stringer.
func (tl *TweetLike) String() string {
	var builder strings.Builder
	builder.WriteString("TweetLike(")
	builder.WriteString("liked_at=")
	builder.WriteString(tl.LikedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", tl.UserID))
	builder.WriteString(", ")
	builder.WriteString("tweet_id=")
	builder.WriteString(fmt.Sprintf("%v", tl.TweetID))
	builder.WriteByte(')')
	return builder.String()
}

// TweetLikes is a parsable slice of TweetLike.
type TweetLikes []*TweetLike

func (tl TweetLikes) config(cfg config) {
	for _i := range tl {
		tl[_i].config = cfg
	}
}
