// Code generated (@generated) by entc, DO NOT EDIT.

package street

const (
	// Label holds the string label denoting the street type in the database.
	Label = "street"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// Table holds the table name of the street in the database.
	Table = "streets"
	// CityTable is the table the holds the city relation/edge.
	CityTable = "streets"
	// CityInverseTable is the table name for the City entity.
	// It exists in this package in order to avoid circular dependency with the "city" package.
	CityInverseTable = "cities"
	// CityColumn is the table column denoting the city relation/edge.
	CityColumn = "city_id"
)

// Columns holds all SQL columns are street fields.
var Columns = []string{
	FieldID,
	FieldName,
}
