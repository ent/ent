// Code generated by ent, DO NOT EDIT.

package relationship

const (
	// Label holds the string label denoting the relationship type in the database.
	Label = "relationship"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRelativeID holds the string denoting the relative_id field in the database.
	FieldRelativeID = "relative_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeRelative holds the string denoting the relative edge name in mutations.
	EdgeRelative = "relative"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// Table holds the table name of the relationship in the database.
	Table = "relationships"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "relationships"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// RelativeTable is the table that holds the relative relation/edge.
	RelativeTable = "relationships"
	// RelativeInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RelativeInverseTable = "users"
	// RelativeColumn is the table column denoting the relative relation/edge.
	RelativeColumn = "relative_id"
)

// Columns holds all SQL columns for relationship fields.
var Columns = []string{
	FieldWeight,
	FieldUserID,
	FieldRelativeID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultWeight holds the default value on creation for the "weight" field.
	DefaultWeight int
)
