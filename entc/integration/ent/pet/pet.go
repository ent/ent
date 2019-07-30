// Code generated (@generated) by entc, DO NOT EDIT.

package pet

const (
	// Label holds the string label denoting the pet type in the database.
	Label = "pet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// Table holds the table name of the pet in the database.
	Table = "pets"
	// TeamTable is the table the holds the team relation/edge.
	TeamTable = "pets"
	// TeamInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	TeamInverseTable = "users"
	// TeamColumn is the table column denoting the team relation/edge.
	TeamColumn = "team_id"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "pets"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"

	// TeamInverseLabel holds the string label denoting the team inverse edge type in the database.
	TeamInverseLabel = "user_team"
	// OwnerInverseLabel holds the string label denoting the owner inverse edge type in the database.
	OwnerInverseLabel = "user_pets"
)

// Columns holds all SQL columns are pet fields.
var Columns = []string{
	FieldID,
	FieldName,
}
