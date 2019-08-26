// Code generated (@generated) by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age vertex property in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// Table holds the table name of the user in the database.
	Table = "users"
	// GroupsTable is the table the holds the groups relation/edge. The primary key declared below.
	GroupsTable = "group_users"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
)

// Columns holds all SQL columns are user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"group_id", "user_id"}
)
