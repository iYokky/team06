// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/patientdetail"
	"github.com/team06/app/ent/predicate"
)

// PatientDetailDelete is the builder for deleting a PatientDetail entity.
type PatientDetailDelete struct {
	config
	hooks    []Hook
	mutation *PatientDetailMutation
}

// Where adds a new predicate to the PatientDetailDelete builder.
func (pdd *PatientDetailDelete) Where(ps ...predicate.PatientDetail) *PatientDetailDelete {
	pdd.mutation.predicates = append(pdd.mutation.predicates, ps...)
	return pdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pdd *PatientDetailDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pdd.hooks) == 0 {
		affected, err = pdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pdd.mutation = mutation
			affected, err = pdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pdd.hooks) - 1; i >= 0; i-- {
			mut = pdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdd *PatientDetailDelete) ExecX(ctx context.Context) int {
	n, err := pdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pdd *PatientDetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: patientdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientdetail.FieldID,
			},
		},
	}
	if ps := pdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pdd.driver, _spec)
}

// PatientDetailDeleteOne is the builder for deleting a single PatientDetail entity.
type PatientDetailDeleteOne struct {
	pdd *PatientDetailDelete
}

// Exec executes the deletion query.
func (pddo *PatientDetailDeleteOne) Exec(ctx context.Context) error {
	n, err := pddo.pdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{patientdetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pddo *PatientDetailDeleteOne) ExecX(ctx context.Context) {
	pddo.pdd.ExecX(ctx)
}
