// Code generated (@generated) by entc, DO NOT EDIT.

package groupinfo

import (
	"github.com/facebookincubator/ent/entc/integration/ent/schema"
)

const (
	// Label holds the string label denoting the groupinfo type in the database.
	Label = "group_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDesc holds the string denoting the desc vertex property in the database.
	FieldDesc = "desc"
	// FieldMaxUsers holds the string denoting the max_users vertex property in the database.
	FieldMaxUsers = "max_users"

	// Table holds the table name of the groupinfo in the database.
	Table = "group_infos"
	// GroupsTable is the table the holds the groups relation/edge.
	GroupsTable = "groups"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// GroupsColumn is the table column denoting the groups relation/edge.
	GroupsColumn = "info_id"

	// GroupsInverseLabel holds the string label denoting the groups inverse edge type in the database.
	GroupsInverseLabel = "group_info"
)

// Columns holds all SQL columns are groupinfo fields.
var Columns = []string{
	FieldID,
	FieldDesc,
	FieldMaxUsers,
}

var (
	fields = schema.GroupInfo{}.Fields()
	// descMaxUsers is the schema descriptor for max_users field.
	descMaxUsers = fields[1].Descriptor()
	// DefaultMaxUsers holds the default value on creation for the max_users field.
	DefaultMaxUsers = descMaxUsers.Default.(int)
)
