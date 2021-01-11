// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/team06/app/ent/diagnosis"
	"github.com/team06/app/ent/doctor"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/treatmenttype"
)

// Diagnosis is the model entity for the Diagnosis schema.
type Diagnosis struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Symptom holds the value of the "symptom" field.
	Symptom string `json:"symptom,omitempty"`
	// Opinionresult holds the value of the "Opinionresult" field.
	Opinionresult string `json:"Opinionresult,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiagnosisQuery when eager-loading is set.
	Edges                                      DiagnosisEdges `json:"edges"`
	doctor_doctor_to_diagnosis                 *int
	patient_patient_to_diagnosis               *int
	treatment_type_treatment_type_to_diagnosis *int
}

// DiagnosisEdges holds the relations/edges for other nodes in the graph.
type DiagnosisEdges struct {
	// DoctorName holds the value of the Doctor_name edge.
	DoctorName *Doctor
	// Patient holds the value of the Patient edge.
	Patient *Patient
	// Type holds the value of the type edge.
	Type *TreatmentType
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DoctorNameOrErr returns the DoctorName value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DiagnosisEdges) DoctorNameOrErr() (*Doctor, error) {
	if e.loadedTypes[0] {
		if e.DoctorName == nil {
			// The edge Doctor_name was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: doctor.Label}
		}
		return e.DoctorName, nil
	}
	return nil, &NotLoadedError{edge: "Doctor_name"}
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DiagnosisEdges) PatientOrErr() (*Patient, error) {
	if e.loadedTypes[1] {
		if e.Patient == nil {
			// The edge Patient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "Patient"}
}

// TypeOrErr returns the Type value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DiagnosisEdges) TypeOrErr() (*TreatmentType, error) {
	if e.loadedTypes[2] {
		if e.Type == nil {
			// The edge type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: treatmenttype.Label}
		}
		return e.Type, nil
	}
	return nil, &NotLoadedError{edge: "type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Diagnosis) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case diagnosis.FieldID:
			values[i] = &sql.NullInt64{}
		case diagnosis.FieldSymptom, diagnosis.FieldOpinionresult:
			values[i] = &sql.NullString{}
		case diagnosis.ForeignKeys[0]: // doctor_doctor_to_diagnosis
			values[i] = &sql.NullInt64{}
		case diagnosis.ForeignKeys[1]: // patient_patient_to_diagnosis
			values[i] = &sql.NullInt64{}
		case diagnosis.ForeignKeys[2]: // treatment_type_treatment_type_to_diagnosis
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Diagnosis", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Diagnosis fields.
func (d *Diagnosis) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case diagnosis.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case diagnosis.FieldSymptom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field symptom", values[i])
			} else if value.Valid {
				d.Symptom = value.String
			}
		case diagnosis.FieldOpinionresult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Opinionresult", values[i])
			} else if value.Valid {
				d.Opinionresult = value.String
			}
		case diagnosis.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field doctor_doctor_to_diagnosis", value)
			} else if value.Valid {
				d.doctor_doctor_to_diagnosis = new(int)
				*d.doctor_doctor_to_diagnosis = int(value.Int64)
			}
		case diagnosis.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field patient_patient_to_diagnosis", value)
			} else if value.Valid {
				d.patient_patient_to_diagnosis = new(int)
				*d.patient_patient_to_diagnosis = int(value.Int64)
			}
		case diagnosis.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field treatment_type_treatment_type_to_diagnosis", value)
			} else if value.Valid {
				d.treatment_type_treatment_type_to_diagnosis = new(int)
				*d.treatment_type_treatment_type_to_diagnosis = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryDoctorName queries the "Doctor_name" edge of the Diagnosis entity.
func (d *Diagnosis) QueryDoctorName() *DoctorQuery {
	return (&DiagnosisClient{config: d.config}).QueryDoctorName(d)
}

// QueryPatient queries the "Patient" edge of the Diagnosis entity.
func (d *Diagnosis) QueryPatient() *PatientQuery {
	return (&DiagnosisClient{config: d.config}).QueryPatient(d)
}

// QueryType queries the "type" edge of the Diagnosis entity.
func (d *Diagnosis) QueryType() *TreatmentTypeQuery {
	return (&DiagnosisClient{config: d.config}).QueryType(d)
}

// Update returns a builder for updating this Diagnosis.
// Note that you need to call Diagnosis.Unwrap() before calling this method if this Diagnosis
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Diagnosis) Update() *DiagnosisUpdateOne {
	return (&DiagnosisClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Diagnosis entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Diagnosis) Unwrap() *Diagnosis {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Diagnosis is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Diagnosis) String() string {
	var builder strings.Builder
	builder.WriteString("Diagnosis(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", symptom=")
	builder.WriteString(d.Symptom)
	builder.WriteString(", Opinionresult=")
	builder.WriteString(d.Opinionresult)
	builder.WriteByte(')')
	return builder.String()
}

// Diagnoses is a parsable slice of Diagnosis.
type Diagnoses []*Diagnosis

func (d Diagnoses) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}