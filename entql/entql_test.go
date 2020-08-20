package entql_test

import (
	"strconv"
	"testing"

	"github.com/facebook/ent/entql"

	"github.com/stretchr/testify/assert"
)

func TestPString(t *testing.T) {
	tests := []struct {
		P entql.P
		S string
	}{
		{
			P: entql.And(
				entql.FieldEQ("name", "a8m"),
				entql.FieldIn("org", "fb", "ent"),
			),
			S: `name == "a8m" && org in ["fb","ent"]`,
		},
		{
			P: entql.Or(
				entql.Not(entql.FieldEQ("name", "mashraki")),
				entql.FieldIn("org", "fb", "ent"),
			),
			S: `!(name == "mashraki") || org in ["fb","ent"]`,
		},
		{
			P: entql.HasEdgeWith(
				"groups",
				entql.HasEdgeWith(
					"admins",
					entql.Not(entql.FieldEQ("name", "a8m")),
				),
			),
			S: `has_edge(groups, has_edge(admins, !(name == "a8m")))`,
		},
		{
			P: entql.And(
				entql.FieldGT("age", 30),
				entql.FieldContains("workplace", "fb"),
			),
			S: `age > 30 && contains(workplace, "fb")`,
		},
		{
			P: entql.Not(entql.FieldLT("score", 32.23)),
			S: `!(score < 32.23)`,
		},
		{
			P: entql.And(
				entql.FieldNil("active"),
				entql.FieldNotNil("name"),
			),
			S: `active == nil && name != nil`,
		},
		{
			P: entql.Or(
				entql.FieldNotIn("id", 1, 2, 3),
				entql.FieldHasSuffix("name", "admin"),
			),
			S: `id not in [1,2,3] || has_suffix(name, "admin")`,
		},
		{
			P: entql.EQ(entql.F("current"), entql.F("total")).Negate(),
			S: `!(current == total)`,
		},
	}
	for i := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			s := tests[i].P.String()
			assert.Equal(t, tests[i].S, s)
		})
	}
}
