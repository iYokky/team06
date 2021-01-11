package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Diagnosis holds the schema definition for the Diagnosis entity.
type Diagnosis struct {
	ent.Schema
}

// Fields of the section.
func (Diagnosis) Fields() []ent.Field {
	return []ent.Field{

		field.String("symptom").NotEmpty(),
		field.String("Opinionresult").NotEmpty(),
	}
}

// Edges of the section.
func (Diagnosis) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("Doctor_name", Doctor.Type).
			Ref("DoctorToDiagnosis").
			Unique(),
		edge.From("Patient", Patient.Type).
			Ref("PatientToRightToTreatment").
			Unique(),
		edge.From("TreatmentTypeToDiagnosis", TreatmentType.Type).
			Ref("type").
			Unique(),
	}
}
