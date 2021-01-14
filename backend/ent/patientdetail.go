// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/team06/app/ent/bloodtype"
	"github.com/team06/app/ent/gender"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/patientdetail"
	"github.com/team06/app/ent/prefix"
)

// PatientDetail is the model entity for the PatientDetail schema.
type PatientDetail struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientDetailQuery when eager-loading is set.
	Edges                      PatientDetailEdges `json:"edges"`
	blood_type_patient_details *int
	gender_patient_details     *int
	patient_patient_details    *int
	prefix_patient_details     *int
}

// PatientDetailEdges holds the relations/edges for other nodes in the graph.
type PatientDetailEdges struct {
	// Prefix holds the value of the prefix edge.
	Prefix *Prefix
	// Gender holds the value of the gender edge.
	Gender *Gender
	// Bloodtype holds the value of the bloodtype edge.
	Bloodtype *BloodType
	// Patient holds the value of the patient edge.
	Patient *Patient
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PrefixOrErr returns the Prefix value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientDetailEdges) PrefixOrErr() (*Prefix, error) {
	if e.loadedTypes[0] {
		if e.Prefix == nil {
			// The edge prefix was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: prefix.Label}
		}
		return e.Prefix, nil
	}
	return nil, &NotLoadedError{edge: "prefix"}
}

// GenderOrErr returns the Gender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientDetailEdges) GenderOrErr() (*Gender, error) {
	if e.loadedTypes[1] {
		if e.Gender == nil {
			// The edge gender was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gender.Label}
		}
		return e.Gender, nil
	}
	return nil, &NotLoadedError{edge: "gender"}
}

// BloodtypeOrErr returns the Bloodtype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientDetailEdges) BloodtypeOrErr() (*BloodType, error) {
	if e.loadedTypes[2] {
		if e.Bloodtype == nil {
			// The edge bloodtype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: bloodtype.Label}
		}
		return e.Bloodtype, nil
	}
	return nil, &NotLoadedError{edge: "bloodtype"}
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PatientDetailEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[3] {
		if e.Patient == nil {
			// The edge patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PatientDetail) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case patientdetail.FieldID:
			values[i] = &sql.NullInt64{}
		case patientdetail.ForeignKeys[0]: // blood_type_patient_details
			values[i] = &sql.NullInt64{}
		case patientdetail.ForeignKeys[1]: // gender_patient_details
			values[i] = &sql.NullInt64{}
		case patientdetail.ForeignKeys[2]: // patient_patient_details
			values[i] = &sql.NullInt64{}
		case patientdetail.ForeignKeys[3]: // prefix_patient_details
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type PatientDetail", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PatientDetail fields.
func (pd *PatientDetail) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case patientdetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pd.ID = int(value.Int64)
		case patientdetail.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field blood_type_patient_details", value)
			} else if value.Valid {
				pd.blood_type_patient_details = new(int)
				*pd.blood_type_patient_details = int(value.Int64)
			}
		case patientdetail.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field gender_patient_details", value)
			} else if value.Valid {
				pd.gender_patient_details = new(int)
				*pd.gender_patient_details = int(value.Int64)
			}
		case patientdetail.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field patient_patient_details", value)
			} else if value.Valid {
				pd.patient_patient_details = new(int)
				*pd.patient_patient_details = int(value.Int64)
			}
		case patientdetail.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field prefix_patient_details", value)
			} else if value.Valid {
				pd.prefix_patient_details = new(int)
				*pd.prefix_patient_details = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPrefix queries the "prefix" edge of the PatientDetail entity.
func (pd *PatientDetail) QueryPrefix() *PrefixQuery {
	return (&PatientDetailClient{config: pd.config}).QueryPrefix(pd)
}

// QueryGender queries the "gender" edge of the PatientDetail entity.
func (pd *PatientDetail) QueryGender() *GenderQuery {
	return (&PatientDetailClient{config: pd.config}).QueryGender(pd)
}

// QueryBloodtype queries the "bloodtype" edge of the PatientDetail entity.
func (pd *PatientDetail) QueryBloodtype() *BloodTypeQuery {
	return (&PatientDetailClient{config: pd.config}).QueryBloodtype(pd)
}

// QueryPatient queries the "patient" edge of the PatientDetail entity.
func (pd *PatientDetail) QueryPatient() *PatientQuery {
	return (&PatientDetailClient{config: pd.config}).QueryPatient(pd)
}

// Update returns a builder for updating this PatientDetail.
// Note that you need to call PatientDetail.Unwrap() before calling this method if this PatientDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (pd *PatientDetail) Update() *PatientDetailUpdateOne {
	return (&PatientDetailClient{config: pd.config}).UpdateOne(pd)
}

// Unwrap unwraps the PatientDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pd *PatientDetail) Unwrap() *PatientDetail {
	tx, ok := pd.config.driver.(*txDriver)
	if !ok {
		panic("ent: PatientDetail is not a transactional entity")
	}
	pd.config.driver = tx.drv
	return pd
}

// String implements the fmt.Stringer.
func (pd *PatientDetail) String() string {
	var builder strings.Builder
	builder.WriteString("PatientDetail(")
	builder.WriteString(fmt.Sprintf("id=%v", pd.ID))
	builder.WriteByte(')')
	return builder.String()
}

// PatientDetails is a parsable slice of PatientDetail.
type PatientDetails []*PatientDetail

func (pd PatientDetails) config(cfg config) {
	for _i := range pd {
		pd[_i].config = cfg
	}
}
