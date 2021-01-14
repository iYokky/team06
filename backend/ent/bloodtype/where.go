// Code generated by entc, DO NOT EDIT.

package bloodtype

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/team06/app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
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
func IDGT(id int) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BloodValue applies equality check predicate on the "bloodValue" field. It's identical to BloodValueEQ.
func BloodValue(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBloodValue), v))
	})
}

// BloodValueEQ applies the EQ predicate on the "bloodValue" field.
func BloodValueEQ(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBloodValue), v))
	})
}

// BloodValueNEQ applies the NEQ predicate on the "bloodValue" field.
func BloodValueNEQ(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBloodValue), v))
	})
}

// BloodValueIn applies the In predicate on the "bloodValue" field.
func BloodValueIn(vs ...string) predicate.BloodType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BloodType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBloodValue), v...))
	})
}

// BloodValueNotIn applies the NotIn predicate on the "bloodValue" field.
func BloodValueNotIn(vs ...string) predicate.BloodType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BloodType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBloodValue), v...))
	})
}

// BloodValueGT applies the GT predicate on the "bloodValue" field.
func BloodValueGT(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBloodValue), v))
	})
}

// BloodValueGTE applies the GTE predicate on the "bloodValue" field.
func BloodValueGTE(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBloodValue), v))
	})
}

// BloodValueLT applies the LT predicate on the "bloodValue" field.
func BloodValueLT(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBloodValue), v))
	})
}

// BloodValueLTE applies the LTE predicate on the "bloodValue" field.
func BloodValueLTE(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBloodValue), v))
	})
}

// BloodValueContains applies the Contains predicate on the "bloodValue" field.
func BloodValueContains(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBloodValue), v))
	})
}

// BloodValueHasPrefix applies the HasPrefix predicate on the "bloodValue" field.
func BloodValueHasPrefix(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBloodValue), v))
	})
}

// BloodValueHasSuffix applies the HasSuffix predicate on the "bloodValue" field.
func BloodValueHasSuffix(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBloodValue), v))
	})
}

// BloodValueEqualFold applies the EqualFold predicate on the "bloodValue" field.
func BloodValueEqualFold(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBloodValue), v))
	})
}

// BloodValueContainsFold applies the ContainsFold predicate on the "bloodValue" field.
func BloodValueContainsFold(v string) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBloodValue), v))
	})
}

// HasPatientDetails applies the HasEdge predicate on the "patient_details" edge.
func HasPatientDetails() predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientDetailsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientDetailsTable, PatientDetailsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientDetailsWith applies the HasEdge predicate on the "patient_details" edge with a given conditions (other predicates).
func HasPatientDetailsWith(preds ...predicate.PatientDetail) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientDetailsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientDetailsTable, PatientDetailsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BloodType) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BloodType) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
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
func Not(p predicate.BloodType) predicate.BloodType {
	return predicate.BloodType(func(s *sql.Selector) {
		p(s.Not())
	})
}
