// Code generated by entc, DO NOT EDIT.

package righttotreatmenttype

const (
	// Label holds the string label denoting the righttotreatmenttype type in the database.
	Label = "right_to_treatment_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTypeName holds the string denoting the typename field in the database.
	FieldTypeName = "type_name"

	// EdgeType holds the string denoting the type edge name in mutations.
	EdgeType = "type"

	// Table holds the table name of the righttotreatmenttype in the database.
	Table = "right_to_treatment_types"
	// TypeTable is the table the holds the type relation/edge.
	TypeTable = "right_to_treatments"
	// TypeInverseTable is the table name for the RightToTreatment entity.
	// It exists in this package in order to avoid circular dependency with the "righttotreatment" package.
	TypeInverseTable = "right_to_treatments"
	// TypeColumn is the table column denoting the type relation/edge.
	TypeColumn = "right_to_treatment_type_type"
)

// Columns holds all SQL columns for righttotreatmenttype fields.
var Columns = []string{
	FieldID,
	FieldTypeName,
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
	// TypeNameValidator is a validator for the "TypeName" field. It is called by the builders before save.
	TypeNameValidator func(string) error
)
