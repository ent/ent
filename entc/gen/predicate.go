package gen

// Op is a predicate for the where clause.
type Op int

// List of all builtin predicates.
const (
	EQ           Op = iota // =
	NEQ                    // <>
	GT                     // >
	GTE                    // >=
	LT                     // <
	LTE                    // <=
	IsNil                  // IS NULL / has
	NotNil                 // IS NOT NULL / hasNot
	In                     // within
	NotIn                  // without
	Contains               // containing
	ContainsFold           // containing case-insensitive
	HasPrefix              // startingWith
	HasSuffix              // endingWith
)

// Name returns the string representation of an predicate.
func (o Op) Name() string {
	if int(o) < len(opText) {
		return opText[o]
	}
	return "Unknown"
}

// Variadic reports if the predicate is a variadic function.
func (o Op) Variadic() bool {
	return o == In || o == NotIn
}

// Niladic reports if the predicate is a niladic predicate.
func (o Op) Niladic() bool {
	return o == IsNil || o == NotNil
}

var (
	// operations text.
	opText = [...]string{
		EQ:           "EQ",
		NEQ:          "NEQ",
		GT:           "GT",
		GTE:          "GTE",
		LT:           "LT",
		LTE:          "LTE",
		IsNil:        "IsNil",
		NotNil:       "NotNil",
		Contains:     "Contains",
		ContainsFold: "ContainsFold",
		HasPrefix:    "HasPrefix",
		HasSuffix:    "HasSuffix",
		In:           "In",
		NotIn:        "NotIn",
	}
	// operations per type.
	boolOps     = []Op{EQ, NEQ}
	numericOps  = append(boolOps[:], GT, GTE, LT, LTE, In, NotIn)
	stringOps   = append(numericOps[:], Contains, HasPrefix, HasSuffix)
	nillableOps = []Op{IsNil, NotNil}
)
