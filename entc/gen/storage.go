package gen

import (
	"fmt"
	"reflect"

	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/sql"
)

// A SchemaOption defines what type of schema feature a storage driver support.
type SchemaMode uint

const (
	// Unique defines field and edge uniqueness support.
	Unique SchemaMode = 1 << iota

	// Cascade defines cascading operations (e.g. cascade deletion).
	Cascade

	// Migrate defines static schema and migration support (e.g. SQL-based).
	Migrate
)

// Support reports whether m support the given mode.
func (m SchemaMode) Support(mode SchemaMode) bool { return m&mode != 0 }

// Storage driver type for codegen.
type Storage struct {
	Name       string          // storage name.
	Builder    reflect.Type    // query builder type.
	Dialects   []string        // supported dialects.
	IdentName  string          // identifier name (fields and funcs).
	Imports    []string        // import packages needed.
	SchemaMode SchemaMode      // schema mode support.
	OpCode     func(Op) string // operation code for predicates.
}

// StorageDrivers holds the storage driver options for entc.
var drivers = []*Storage{
	{
		Name:      "sql",
		IdentName: "SQL",
		Builder:   reflect.TypeOf(&sql.Selector{}),
		Dialects:  []string{"dialect.SQLite", "dialect.MySQL"},
		Imports: []string{
			"github.com/facebookincubator/ent/dialect/sql",
		},
		SchemaMode: Unique | Cascade | Migrate,
		OpCode:     opCodes(sqlCode[:]),
	},
	{
		Name:      "gremlin",
		IdentName: "Gremlin",
		Builder:   reflect.TypeOf(&dsl.Traversal{}),
		Dialects:  []string{"dialect.Neptune"},
		Imports: []string{
			"github.com/facebookincubator/ent/dialect/gremlin",
			"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl",
			"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__",
			"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g",
			"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p",
			"github.com/facebookincubator/ent/dialect/gremlin/encoding/graphson",
		},
		SchemaMode: Unique,
		OpCode:     opCodes(gremlinCode[:]),
	},
}

// NewStorage returns a the storage driver type from the given string.
// It fails if the provided string is not a valid option. this function
// is here in order to remove the validation logic from entc command line.
func NewStorage(s string) (*Storage, error) {
	for _, d := range drivers {
		if s == d.Name {
			return d, nil
		}
	}
	return nil, fmt.Errorf("entc/gen: invalid storage driver %q", s)
}

// String implements the fmt.Stringer interface for template usage.
func (s *Storage) String() string { return s.Name }

var (
	// exceptional operation names in sql.
	sqlCode = [...]string{
		IsNil:  "IsNull",
		NotNil: "NotNull",
	}
	// exceptional operation names in gremlin.
	gremlinCode = [...]string{
		IsNil:     "HasNot",
		NotNil:    "Has",
		In:        "Within",
		NotIn:     "Without",
		Contains:  "Containing",
		HasPrefix: "StartingWith",
		HasSuffix: "EndingWith",
	}
)

func opCodes(codes []string) func(Op) string {
	return func(o Op) string {
		if int(o) < len(codes) && codes[o] != "" {
			return codes[o]
		}
		return o.Name()
	}
}
