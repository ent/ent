package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
)

// Predicate is a query predicate.
type Predicate struct {
	QueryBuilder
	fns []func() expression.ConditionBuilder
}

// QueryBuilder is the builder for DynamoDB condition expression.
type QueryBuilder struct {
	condBuilder expression.ConditionBuilder
}

// P creates a new predicate.
//
// P().EQ("name", "tk1122").And().EQ("age", 24)
func P() *Predicate {
	return &Predicate{}
}

// Query runs all appended build steps.
func (p *Predicate) Query() expression.ConditionBuilder {
	p.condBuilder = p.fns[0]()
	for _, f := range p.fns[1:] {
		p.condBuilder = expression.And(p.condBuilder, f())
	}
	return p.condBuilder
}

// Append appends match builders to predicate.
func (p *Predicate) Append(fs ...func() expression.ConditionBuilder) *Predicate {
	p.fns = append(p.fns, fs...)
	return p
}

// EQ returns `expression.Name(key).Equal(expression.Value(val))` predicate.
func EQ(key string, val interface{}) *Predicate {
	return P().EQ(key, val)
}

// EQ appends a "expression.Name(key).Equal(expression.Value(val))" predicate.
func (p *Predicate) EQ(key string, val interface{}) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Name(key).Equal(expression.Value(val))
	})
}

// In returns the `expression.Name(key).In(expression.Value(vals))` predicate.
func In(key string, vals ...interface{}) *Predicate {
	return P().In(key, vals...)
}

// In appends the `expression.Name(key).In(expression.Value(vals))` predicate.
func (p *Predicate) In(key string, vals ...interface{}) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Name(key).In(expression.Value(vals))
	})
}

// And combines all given predicates with expression.And between them.
func And(preds ...*Predicate) *Predicate {
	p := P()
	return p.Append(func() expression.ConditionBuilder {
		cond := preds[0].Query()
		for _, pred := range preds[1:] {
			cond = expression.And(cond, pred.Query())
		}
		return cond
	})
}

// NotExist returns the `expression.Not(expression.Name(key).AttributeExists())` predicate.
func NotExist(col string) *Predicate {
	return P().NotExist(col)
}

// NotExist appends the `expression.Not(expression.Name(key).AttributeExists())` predicate.
func (p *Predicate) NotExist(key string) *Predicate {
	return p.Append(func() expression.ConditionBuilder {
		return expression.Not(expression.Name(key).AttributeExists())
	})
}
