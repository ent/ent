// Code generated (@generated) by entc, DO NOT EDIT.

package city

const (
	// Label holds the string label denoting the city type in the database.
	Label = "city"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// Table holds the table name of the city in the database.
	Table = "cities"
	// StreetsTable is the table the holds the streets relation/edge.
	StreetsTable = "streets"
	// StreetsInverseTable is the table name for the Street entity.
	// It exists in this package in order to avoid circular dependency with the "street" package.
	StreetsInverseTable = "streets"
	// StreetsColumn is the table column denoting the streets relation/edge.
	StreetsColumn = "city_id"
)

// Columns holds all SQL columns are city fields.
var Columns = []string{
	FieldID,
	FieldName,
}
