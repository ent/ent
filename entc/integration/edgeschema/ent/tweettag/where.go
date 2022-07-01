// Code generated by ent, DO NOT EDIT.

package tweettag

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AddedAt applies equality check predicate on the "added_at" field. It's identical to AddedAtEQ.
func AddedAt(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedAt), v))
	})
}

// TagID applies equality check predicate on the "tag_id" field. It's identical to TagIDEQ.
func TagID(v int) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagID), v))
	})
}

// TweetID applies equality check predicate on the "tweet_id" field. It's identical to TweetIDEQ.
func TweetID(v int) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTweetID), v))
	})
}

// AddedAtEQ applies the EQ predicate on the "added_at" field.
func AddedAtEQ(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddedAt), v))
	})
}

// AddedAtNEQ applies the NEQ predicate on the "added_at" field.
func AddedAtNEQ(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddedAt), v))
	})
}

// AddedAtIn applies the In predicate on the "added_at" field.
func AddedAtIn(vs ...time.Time) predicate.TweetTag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetTag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddedAt), v...))
	})
}

// AddedAtNotIn applies the NotIn predicate on the "added_at" field.
func AddedAtNotIn(vs ...time.Time) predicate.TweetTag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetTag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddedAt), v...))
	})
}

// AddedAtGT applies the GT predicate on the "added_at" field.
func AddedAtGT(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddedAt), v))
	})
}

// AddedAtGTE applies the GTE predicate on the "added_at" field.
func AddedAtGTE(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddedAt), v))
	})
}

// AddedAtLT applies the LT predicate on the "added_at" field.
func AddedAtLT(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddedAt), v))
	})
}

// AddedAtLTE applies the LTE predicate on the "added_at" field.
func AddedAtLTE(v time.Time) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddedAt), v))
	})
}

// TagIDEQ applies the EQ predicate on the "tag_id" field.
func TagIDEQ(v int) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagID), v))
	})
}

// TagIDNEQ applies the NEQ predicate on the "tag_id" field.
func TagIDNEQ(v int) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTagID), v))
	})
}

// TagIDIn applies the In predicate on the "tag_id" field.
func TagIDIn(vs ...int) predicate.TweetTag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetTag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTagID), v...))
	})
}

// TagIDNotIn applies the NotIn predicate on the "tag_id" field.
func TagIDNotIn(vs ...int) predicate.TweetTag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetTag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTagID), v...))
	})
}

// TweetIDEQ applies the EQ predicate on the "tweet_id" field.
func TweetIDEQ(v int) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTweetID), v))
	})
}

// TweetIDNEQ applies the NEQ predicate on the "tweet_id" field.
func TweetIDNEQ(v int) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTweetID), v))
	})
}

// TweetIDIn applies the In predicate on the "tweet_id" field.
func TweetIDIn(vs ...int) predicate.TweetTag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetTag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTweetID), v...))
	})
}

// TweetIDNotIn applies the NotIn predicate on the "tweet_id" field.
func TweetIDNotIn(vs ...int) predicate.TweetTag {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TweetTag(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTweetID), v...))
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTweet applies the HasEdge predicate on the "tweet" edge.
func HasTweet() predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TweetTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TweetTable, TweetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTweetWith applies the HasEdge predicate on the "tweet" edge with a given conditions (other predicates).
func HasTweetWith(preds ...predicate.Tweet) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TweetInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TweetTable, TweetColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TweetTag) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TweetTag) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TweetTag) predicate.TweetTag {
	return predicate.TweetTag(func(s *sql.Selector) {
		p(s.Not())
	})
}
