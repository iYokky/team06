// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/team06/app/ent/prefix"
)

// Prefix is the model entity for the Prefix schema.
type Prefix struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PrefixValue holds the value of the "prefixValue" field.
	PrefixValue string `json:"prefixValue,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PrefixQuery when eager-loading is set.
	Edges PrefixEdges `json:"edges"`
}

// PrefixEdges holds the relations/edges for other nodes in the graph.
type PrefixEdges struct {
	// Patient holds the value of the patient edge.
	Patient []*Patient
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading.
func (e PrefixEdges) PatientOrErr() ([]*Patient, error) {
	if e.loadedTypes[0] {
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Prefix) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case prefix.FieldID:
			values[i] = &sql.NullInt64{}
		case prefix.FieldPrefixValue:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Prefix", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Prefix fields.
func (pr *Prefix) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case prefix.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case prefix.FieldPrefixValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field prefixValue", values[i])
			} else if value.Valid {
				pr.PrefixValue = value.String
			}
		}
	}
	return nil
}

// QueryPatient queries the "patient" edge of the Prefix entity.
func (pr *Prefix) QueryPatient() *PatientQuery {
	return (&PrefixClient{config: pr.config}).QueryPatient(pr)
}

// Update returns a builder for updating this Prefix.
// Note that you need to call Prefix.Unwrap() before calling this method if this Prefix
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Prefix) Update() *PrefixUpdateOne {
	return (&PrefixClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Prefix entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Prefix) Unwrap() *Prefix {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Prefix is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Prefix) String() string {
	var builder strings.Builder
	builder.WriteString("Prefix(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", prefixValue=")
	builder.WriteString(pr.PrefixValue)
	builder.WriteByte(')')
	return builder.String()
}

// Prefixes is a parsable slice of Prefix.
type Prefixes []*Prefix

func (pr Prefixes) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
