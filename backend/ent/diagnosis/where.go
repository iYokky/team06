// Code generated by entc, DO NOT EDIT.

package diagnosis

import (
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/team06/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Symptom applies equality check predicate on the "symptom" field. It's identical to SymptomEQ.
func Symptom(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymptom), v))
	})
}

// Opinionresult applies equality check predicate on the "Opinionresult" field. It's identical to OpinionresultEQ.
func Opinionresult(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOpinionresult), v))
	})
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// DiagnosisDate applies equality check predicate on the "diagnosisDate" field. It's identical to DiagnosisDateEQ.
func DiagnosisDate(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiagnosisDate), v))
	})
}

// SymptomEQ applies the EQ predicate on the "symptom" field.
func SymptomEQ(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymptom), v))
	})
}

// SymptomNEQ applies the NEQ predicate on the "symptom" field.
func SymptomNEQ(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSymptom), v))
	})
}

// SymptomIn applies the In predicate on the "symptom" field.
func SymptomIn(vs ...string) predicate.Diagnosis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Diagnosis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSymptom), v...))
	})
}

// SymptomNotIn applies the NotIn predicate on the "symptom" field.
func SymptomNotIn(vs ...string) predicate.Diagnosis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Diagnosis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSymptom), v...))
	})
}

// SymptomGT applies the GT predicate on the "symptom" field.
func SymptomGT(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSymptom), v))
	})
}

// SymptomGTE applies the GTE predicate on the "symptom" field.
func SymptomGTE(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSymptom), v))
	})
}

// SymptomLT applies the LT predicate on the "symptom" field.
func SymptomLT(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSymptom), v))
	})
}

// SymptomLTE applies the LTE predicate on the "symptom" field.
func SymptomLTE(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSymptom), v))
	})
}

// SymptomContains applies the Contains predicate on the "symptom" field.
func SymptomContains(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSymptom), v))
	})
}

// SymptomHasPrefix applies the HasPrefix predicate on the "symptom" field.
func SymptomHasPrefix(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSymptom), v))
	})
}

// SymptomHasSuffix applies the HasSuffix predicate on the "symptom" field.
func SymptomHasSuffix(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSymptom), v))
	})
}

// SymptomEqualFold applies the EqualFold predicate on the "symptom" field.
func SymptomEqualFold(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSymptom), v))
	})
}

// SymptomContainsFold applies the ContainsFold predicate on the "symptom" field.
func SymptomContainsFold(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSymptom), v))
	})
}

// OpinionresultEQ applies the EQ predicate on the "Opinionresult" field.
func OpinionresultEQ(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOpinionresult), v))
	})
}

// OpinionresultNEQ applies the NEQ predicate on the "Opinionresult" field.
func OpinionresultNEQ(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOpinionresult), v))
	})
}

// OpinionresultIn applies the In predicate on the "Opinionresult" field.
func OpinionresultIn(vs ...string) predicate.Diagnosis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Diagnosis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOpinionresult), v...))
	})
}

// OpinionresultNotIn applies the NotIn predicate on the "Opinionresult" field.
func OpinionresultNotIn(vs ...string) predicate.Diagnosis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Diagnosis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOpinionresult), v...))
	})
}

// OpinionresultGT applies the GT predicate on the "Opinionresult" field.
func OpinionresultGT(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOpinionresult), v))
	})
}

// OpinionresultGTE applies the GTE predicate on the "Opinionresult" field.
func OpinionresultGTE(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOpinionresult), v))
	})
}

// OpinionresultLT applies the LT predicate on the "Opinionresult" field.
func OpinionresultLT(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOpinionresult), v))
	})
}

// OpinionresultLTE applies the LTE predicate on the "Opinionresult" field.
func OpinionresultLTE(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOpinionresult), v))
	})
}

// OpinionresultContains applies the Contains predicate on the "Opinionresult" field.
func OpinionresultContains(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOpinionresult), v))
	})
}

// OpinionresultHasPrefix applies the HasPrefix predicate on the "Opinionresult" field.
func OpinionresultHasPrefix(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOpinionresult), v))
	})
}

// OpinionresultHasSuffix applies the HasSuffix predicate on the "Opinionresult" field.
func OpinionresultHasSuffix(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOpinionresult), v))
	})
}

// OpinionresultEqualFold applies the EqualFold predicate on the "Opinionresult" field.
func OpinionresultEqualFold(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOpinionresult), v))
	})
}

// OpinionresultContainsFold applies the ContainsFold predicate on the "Opinionresult" field.
func OpinionresultContainsFold(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOpinionresult), v))
	})
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNote), v))
	})
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.Diagnosis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Diagnosis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNote), v...))
	})
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.Diagnosis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Diagnosis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNote), v...))
	})
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNote), v))
	})
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNote), v))
	})
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNote), v))
	})
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNote), v))
	})
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNote), v))
	})
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNote), v))
	})
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNote), v))
	})
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNote), v))
	})
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNote), v))
	})
}

// DiagnosisDateEQ applies the EQ predicate on the "diagnosisDate" field.
func DiagnosisDateEQ(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiagnosisDate), v))
	})
}

// DiagnosisDateNEQ applies the NEQ predicate on the "diagnosisDate" field.
func DiagnosisDateNEQ(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiagnosisDate), v))
	})
}

// DiagnosisDateIn applies the In predicate on the "diagnosisDate" field.
func DiagnosisDateIn(vs ...time.Time) predicate.Diagnosis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Diagnosis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiagnosisDate), v...))
	})
}

// DiagnosisDateNotIn applies the NotIn predicate on the "diagnosisDate" field.
func DiagnosisDateNotIn(vs ...time.Time) predicate.Diagnosis {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Diagnosis(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiagnosisDate), v...))
	})
}

// DiagnosisDateGT applies the GT predicate on the "diagnosisDate" field.
func DiagnosisDateGT(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiagnosisDate), v))
	})
}

// DiagnosisDateGTE applies the GTE predicate on the "diagnosisDate" field.
func DiagnosisDateGTE(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiagnosisDate), v))
	})
}

// DiagnosisDateLT applies the LT predicate on the "diagnosisDate" field.
func DiagnosisDateLT(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiagnosisDate), v))
	})
}

// DiagnosisDateLTE applies the LTE predicate on the "diagnosisDate" field.
func DiagnosisDateLTE(v time.Time) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiagnosisDate), v))
	})
}

// HasDoctorName applies the HasEdge predicate on the "Doctor_name" edge.
func HasDoctorName() predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorNameTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DoctorNameTable, DoctorNameColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDoctorNameWith applies the HasEdge predicate on the "Doctor_name" edge with a given conditions (other predicates).
func HasDoctorNameWith(preds ...predicate.Doctor) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorNameInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DoctorNameTable, DoctorNameColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatient applies the HasEdge predicate on the "Patient" edge.
func HasPatient() predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "Patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.Patient) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasType applies the HasEdge predicate on the "type" edge.
func HasType() predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TypeTable, TypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTypeWith applies the HasEdge predicate on the "type" edge with a given conditions (other predicates).
func HasTypeWith(preds ...predicate.TreatmentType) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TypeTable, TypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Diagnosis) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Diagnosis) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Diagnosis) predicate.Diagnosis {
	return predicate.Diagnosis(func(s *sql.Selector) {
		p(s.Not())
	})
}
