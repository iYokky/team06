// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/department"
	"github.com/team06/app/ent/nurse"
	"github.com/team06/app/ent/patient"
	"github.com/team06/app/ent/predicate"
	"github.com/team06/app/ent/triageresult"
	"github.com/team06/app/ent/urgencylevel"
)

// TriageResultUpdate is the builder for updating TriageResult entities.
type TriageResultUpdate struct {
	config
	hooks    []Hook
	mutation *TriageResultMutation
}

// Where adds a new predicate for the TriageResultUpdate builder.
func (tru *TriageResultUpdate) Where(ps ...predicate.TriageResult) *TriageResultUpdate {
	tru.mutation.predicates = append(tru.mutation.predicates, ps...)
	return tru
}

// SetSymptom sets the "symptom" field.
func (tru *TriageResultUpdate) SetSymptom(s string) *TriageResultUpdate {
	tru.mutation.SetSymptom(s)
	return tru
}

// SetTriageDate sets the "triageDate" field.
func (tru *TriageResultUpdate) SetTriageDate(t time.Time) *TriageResultUpdate {
	tru.mutation.SetTriageDate(t)
	return tru
}

// SetNillableTriageDate sets the "triageDate" field if the given value is not nil.
func (tru *TriageResultUpdate) SetNillableTriageDate(t *time.Time) *TriageResultUpdate {
	if t != nil {
		tru.SetTriageDate(*t)
	}
	return tru
}

// SetTriageResultToUrgencyLevelID sets the "triageResultToUrgencyLevel" edge to the UrgencyLevel entity by ID.
func (tru *TriageResultUpdate) SetTriageResultToUrgencyLevelID(id int) *TriageResultUpdate {
	tru.mutation.SetTriageResultToUrgencyLevelID(id)
	return tru
}

// SetNillableTriageResultToUrgencyLevelID sets the "triageResultToUrgencyLevel" edge to the UrgencyLevel entity by ID if the given value is not nil.
func (tru *TriageResultUpdate) SetNillableTriageResultToUrgencyLevelID(id *int) *TriageResultUpdate {
	if id != nil {
		tru = tru.SetTriageResultToUrgencyLevelID(*id)
	}
	return tru
}

// SetTriageResultToUrgencyLevel sets the "triageResultToUrgencyLevel" edge to the UrgencyLevel entity.
func (tru *TriageResultUpdate) SetTriageResultToUrgencyLevel(u *UrgencyLevel) *TriageResultUpdate {
	return tru.SetTriageResultToUrgencyLevelID(u.ID)
}

// SetTriageResultToDepartmentID sets the "triageResultToDepartment" edge to the Department entity by ID.
func (tru *TriageResultUpdate) SetTriageResultToDepartmentID(id int) *TriageResultUpdate {
	tru.mutation.SetTriageResultToDepartmentID(id)
	return tru
}

// SetNillableTriageResultToDepartmentID sets the "triageResultToDepartment" edge to the Department entity by ID if the given value is not nil.
func (tru *TriageResultUpdate) SetNillableTriageResultToDepartmentID(id *int) *TriageResultUpdate {
	if id != nil {
		tru = tru.SetTriageResultToDepartmentID(*id)
	}
	return tru
}

// SetTriageResultToDepartment sets the "triageResultToDepartment" edge to the Department entity.
func (tru *TriageResultUpdate) SetTriageResultToDepartment(d *Department) *TriageResultUpdate {
	return tru.SetTriageResultToDepartmentID(d.ID)
}

// SetTriageResultToNurseID sets the "triageResultToNurse" edge to the Nurse entity by ID.
func (tru *TriageResultUpdate) SetTriageResultToNurseID(id int) *TriageResultUpdate {
	tru.mutation.SetTriageResultToNurseID(id)
	return tru
}

// SetNillableTriageResultToNurseID sets the "triageResultToNurse" edge to the Nurse entity by ID if the given value is not nil.
func (tru *TriageResultUpdate) SetNillableTriageResultToNurseID(id *int) *TriageResultUpdate {
	if id != nil {
		tru = tru.SetTriageResultToNurseID(*id)
	}
	return tru
}

// SetTriageResultToNurse sets the "triageResultToNurse" edge to the Nurse entity.
func (tru *TriageResultUpdate) SetTriageResultToNurse(n *Nurse) *TriageResultUpdate {
	return tru.SetTriageResultToNurseID(n.ID)
}

// SetTriageResultToPatientID sets the "triageResultToPatient" edge to the Patient entity by ID.
func (tru *TriageResultUpdate) SetTriageResultToPatientID(id int) *TriageResultUpdate {
	tru.mutation.SetTriageResultToPatientID(id)
	return tru
}

// SetNillableTriageResultToPatientID sets the "triageResultToPatient" edge to the Patient entity by ID if the given value is not nil.
func (tru *TriageResultUpdate) SetNillableTriageResultToPatientID(id *int) *TriageResultUpdate {
	if id != nil {
		tru = tru.SetTriageResultToPatientID(*id)
	}
	return tru
}

// SetTriageResultToPatient sets the "triageResultToPatient" edge to the Patient entity.
func (tru *TriageResultUpdate) SetTriageResultToPatient(p *Patient) *TriageResultUpdate {
	return tru.SetTriageResultToPatientID(p.ID)
}

// Mutation returns the TriageResultMutation object of the builder.
func (tru *TriageResultUpdate) Mutation() *TriageResultMutation {
	return tru.mutation
}

// ClearTriageResultToUrgencyLevel clears the "triageResultToUrgencyLevel" edge to the UrgencyLevel entity.
func (tru *TriageResultUpdate) ClearTriageResultToUrgencyLevel() *TriageResultUpdate {
	tru.mutation.ClearTriageResultToUrgencyLevel()
	return tru
}

// ClearTriageResultToDepartment clears the "triageResultToDepartment" edge to the Department entity.
func (tru *TriageResultUpdate) ClearTriageResultToDepartment() *TriageResultUpdate {
	tru.mutation.ClearTriageResultToDepartment()
	return tru
}

// ClearTriageResultToNurse clears the "triageResultToNurse" edge to the Nurse entity.
func (tru *TriageResultUpdate) ClearTriageResultToNurse() *TriageResultUpdate {
	tru.mutation.ClearTriageResultToNurse()
	return tru
}

// ClearTriageResultToPatient clears the "triageResultToPatient" edge to the Patient entity.
func (tru *TriageResultUpdate) ClearTriageResultToPatient() *TriageResultUpdate {
	tru.mutation.ClearTriageResultToPatient()
	return tru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tru *TriageResultUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tru.hooks) == 0 {
		if err = tru.check(); err != nil {
			return 0, err
		}
		affected, err = tru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TriageResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tru.check(); err != nil {
				return 0, err
			}
			tru.mutation = mutation
			affected, err = tru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tru.hooks) - 1; i >= 0; i-- {
			mut = tru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tru *TriageResultUpdate) SaveX(ctx context.Context) int {
	affected, err := tru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tru *TriageResultUpdate) Exec(ctx context.Context) error {
	_, err := tru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tru *TriageResultUpdate) ExecX(ctx context.Context) {
	if err := tru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tru *TriageResultUpdate) check() error {
	if v, ok := tru.mutation.Symptom(); ok {
		if err := triageresult.SymptomValidator(v); err != nil {
			return &ValidationError{Name: "symptom", err: fmt.Errorf("ent: validator failed for field \"symptom\": %w", err)}
		}
	}
	return nil
}

func (tru *TriageResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   triageresult.Table,
			Columns: triageresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: triageresult.FieldID,
			},
		},
	}
	if ps := tru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tru.mutation.Symptom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: triageresult.FieldSymptom,
		})
	}
	if value, ok := tru.mutation.TriageDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: triageresult.FieldTriageDate,
		})
	}
	if tru.mutation.TriageResultToUrgencyLevelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToUrgencyLevelTable,
			Columns: []string{triageresult.TriageResultToUrgencyLevelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: urgencylevel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.TriageResultToUrgencyLevelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToUrgencyLevelTable,
			Columns: []string{triageresult.TriageResultToUrgencyLevelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: urgencylevel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tru.mutation.TriageResultToDepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToDepartmentTable,
			Columns: []string{triageresult.TriageResultToDepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.TriageResultToDepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToDepartmentTable,
			Columns: []string{triageresult.TriageResultToDepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tru.mutation.TriageResultToNurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToNurseTable,
			Columns: []string{triageresult.TriageResultToNurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.TriageResultToNurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToNurseTable,
			Columns: []string{triageresult.TriageResultToNurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tru.mutation.TriageResultToPatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToPatientTable,
			Columns: []string{triageresult.TriageResultToPatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.TriageResultToPatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToPatientTable,
			Columns: []string{triageresult.TriageResultToPatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{triageresult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TriageResultUpdateOne is the builder for updating a single TriageResult entity.
type TriageResultUpdateOne struct {
	config
	hooks    []Hook
	mutation *TriageResultMutation
}

// SetSymptom sets the "symptom" field.
func (truo *TriageResultUpdateOne) SetSymptom(s string) *TriageResultUpdateOne {
	truo.mutation.SetSymptom(s)
	return truo
}

// SetTriageDate sets the "triageDate" field.
func (truo *TriageResultUpdateOne) SetTriageDate(t time.Time) *TriageResultUpdateOne {
	truo.mutation.SetTriageDate(t)
	return truo
}

// SetNillableTriageDate sets the "triageDate" field if the given value is not nil.
func (truo *TriageResultUpdateOne) SetNillableTriageDate(t *time.Time) *TriageResultUpdateOne {
	if t != nil {
		truo.SetTriageDate(*t)
	}
	return truo
}

// SetTriageResultToUrgencyLevelID sets the "triageResultToUrgencyLevel" edge to the UrgencyLevel entity by ID.
func (truo *TriageResultUpdateOne) SetTriageResultToUrgencyLevelID(id int) *TriageResultUpdateOne {
	truo.mutation.SetTriageResultToUrgencyLevelID(id)
	return truo
}

// SetNillableTriageResultToUrgencyLevelID sets the "triageResultToUrgencyLevel" edge to the UrgencyLevel entity by ID if the given value is not nil.
func (truo *TriageResultUpdateOne) SetNillableTriageResultToUrgencyLevelID(id *int) *TriageResultUpdateOne {
	if id != nil {
		truo = truo.SetTriageResultToUrgencyLevelID(*id)
	}
	return truo
}

// SetTriageResultToUrgencyLevel sets the "triageResultToUrgencyLevel" edge to the UrgencyLevel entity.
func (truo *TriageResultUpdateOne) SetTriageResultToUrgencyLevel(u *UrgencyLevel) *TriageResultUpdateOne {
	return truo.SetTriageResultToUrgencyLevelID(u.ID)
}

// SetTriageResultToDepartmentID sets the "triageResultToDepartment" edge to the Department entity by ID.
func (truo *TriageResultUpdateOne) SetTriageResultToDepartmentID(id int) *TriageResultUpdateOne {
	truo.mutation.SetTriageResultToDepartmentID(id)
	return truo
}

// SetNillableTriageResultToDepartmentID sets the "triageResultToDepartment" edge to the Department entity by ID if the given value is not nil.
func (truo *TriageResultUpdateOne) SetNillableTriageResultToDepartmentID(id *int) *TriageResultUpdateOne {
	if id != nil {
		truo = truo.SetTriageResultToDepartmentID(*id)
	}
	return truo
}

// SetTriageResultToDepartment sets the "triageResultToDepartment" edge to the Department entity.
func (truo *TriageResultUpdateOne) SetTriageResultToDepartment(d *Department) *TriageResultUpdateOne {
	return truo.SetTriageResultToDepartmentID(d.ID)
}

// SetTriageResultToNurseID sets the "triageResultToNurse" edge to the Nurse entity by ID.
func (truo *TriageResultUpdateOne) SetTriageResultToNurseID(id int) *TriageResultUpdateOne {
	truo.mutation.SetTriageResultToNurseID(id)
	return truo
}

// SetNillableTriageResultToNurseID sets the "triageResultToNurse" edge to the Nurse entity by ID if the given value is not nil.
func (truo *TriageResultUpdateOne) SetNillableTriageResultToNurseID(id *int) *TriageResultUpdateOne {
	if id != nil {
		truo = truo.SetTriageResultToNurseID(*id)
	}
	return truo
}

// SetTriageResultToNurse sets the "triageResultToNurse" edge to the Nurse entity.
func (truo *TriageResultUpdateOne) SetTriageResultToNurse(n *Nurse) *TriageResultUpdateOne {
	return truo.SetTriageResultToNurseID(n.ID)
}

// SetTriageResultToPatientID sets the "triageResultToPatient" edge to the Patient entity by ID.
func (truo *TriageResultUpdateOne) SetTriageResultToPatientID(id int) *TriageResultUpdateOne {
	truo.mutation.SetTriageResultToPatientID(id)
	return truo
}

// SetNillableTriageResultToPatientID sets the "triageResultToPatient" edge to the Patient entity by ID if the given value is not nil.
func (truo *TriageResultUpdateOne) SetNillableTriageResultToPatientID(id *int) *TriageResultUpdateOne {
	if id != nil {
		truo = truo.SetTriageResultToPatientID(*id)
	}
	return truo
}

// SetTriageResultToPatient sets the "triageResultToPatient" edge to the Patient entity.
func (truo *TriageResultUpdateOne) SetTriageResultToPatient(p *Patient) *TriageResultUpdateOne {
	return truo.SetTriageResultToPatientID(p.ID)
}

// Mutation returns the TriageResultMutation object of the builder.
func (truo *TriageResultUpdateOne) Mutation() *TriageResultMutation {
	return truo.mutation
}

// ClearTriageResultToUrgencyLevel clears the "triageResultToUrgencyLevel" edge to the UrgencyLevel entity.
func (truo *TriageResultUpdateOne) ClearTriageResultToUrgencyLevel() *TriageResultUpdateOne {
	truo.mutation.ClearTriageResultToUrgencyLevel()
	return truo
}

// ClearTriageResultToDepartment clears the "triageResultToDepartment" edge to the Department entity.
func (truo *TriageResultUpdateOne) ClearTriageResultToDepartment() *TriageResultUpdateOne {
	truo.mutation.ClearTriageResultToDepartment()
	return truo
}

// ClearTriageResultToNurse clears the "triageResultToNurse" edge to the Nurse entity.
func (truo *TriageResultUpdateOne) ClearTriageResultToNurse() *TriageResultUpdateOne {
	truo.mutation.ClearTriageResultToNurse()
	return truo
}

// ClearTriageResultToPatient clears the "triageResultToPatient" edge to the Patient entity.
func (truo *TriageResultUpdateOne) ClearTriageResultToPatient() *TriageResultUpdateOne {
	truo.mutation.ClearTriageResultToPatient()
	return truo
}

// Save executes the query and returns the updated TriageResult entity.
func (truo *TriageResultUpdateOne) Save(ctx context.Context) (*TriageResult, error) {
	var (
		err  error
		node *TriageResult
	)
	if len(truo.hooks) == 0 {
		if err = truo.check(); err != nil {
			return nil, err
		}
		node, err = truo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TriageResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = truo.check(); err != nil {
				return nil, err
			}
			truo.mutation = mutation
			node, err = truo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(truo.hooks) - 1; i >= 0; i-- {
			mut = truo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, truo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (truo *TriageResultUpdateOne) SaveX(ctx context.Context) *TriageResult {
	node, err := truo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (truo *TriageResultUpdateOne) Exec(ctx context.Context) error {
	_, err := truo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (truo *TriageResultUpdateOne) ExecX(ctx context.Context) {
	if err := truo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (truo *TriageResultUpdateOne) check() error {
	if v, ok := truo.mutation.Symptom(); ok {
		if err := triageresult.SymptomValidator(v); err != nil {
			return &ValidationError{Name: "symptom", err: fmt.Errorf("ent: validator failed for field \"symptom\": %w", err)}
		}
	}
	return nil
}

func (truo *TriageResultUpdateOne) sqlSave(ctx context.Context) (_node *TriageResult, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   triageresult.Table,
			Columns: triageresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: triageresult.FieldID,
			},
		},
	}
	id, ok := truo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TriageResult.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := truo.mutation.Symptom(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: triageresult.FieldSymptom,
		})
	}
	if value, ok := truo.mutation.TriageDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: triageresult.FieldTriageDate,
		})
	}
	if truo.mutation.TriageResultToUrgencyLevelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToUrgencyLevelTable,
			Columns: []string{triageresult.TriageResultToUrgencyLevelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: urgencylevel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.TriageResultToUrgencyLevelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToUrgencyLevelTable,
			Columns: []string{triageresult.TriageResultToUrgencyLevelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: urgencylevel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if truo.mutation.TriageResultToDepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToDepartmentTable,
			Columns: []string{triageresult.TriageResultToDepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.TriageResultToDepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToDepartmentTable,
			Columns: []string{triageresult.TriageResultToDepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if truo.mutation.TriageResultToNurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToNurseTable,
			Columns: []string{triageresult.TriageResultToNurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.TriageResultToNurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToNurseTable,
			Columns: []string{triageresult.TriageResultToNurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if truo.mutation.TriageResultToPatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToPatientTable,
			Columns: []string{triageresult.TriageResultToPatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.TriageResultToPatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triageresult.TriageResultToPatientTable,
			Columns: []string{triageresult.TriageResultToPatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TriageResult{config: truo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, truo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{triageresult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}