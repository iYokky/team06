// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/predicate"
	"github.com/team06/app/ent/righttotreatment"
)

// RightToTreatmentDelete is the builder for deleting a RightToTreatment entity.
type RightToTreatmentDelete struct {
	config
	hooks    []Hook
	mutation *RightToTreatmentMutation
}

// Where adds a new predicate to the RightToTreatmentDelete builder.
func (rttd *RightToTreatmentDelete) Where(ps ...predicate.RightToTreatment) *RightToTreatmentDelete {
	rttd.mutation.predicates = append(rttd.mutation.predicates, ps...)
	return rttd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rttd *RightToTreatmentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rttd.hooks) == 0 {
		affected, err = rttd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RightToTreatmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rttd.mutation = mutation
			affected, err = rttd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rttd.hooks) - 1; i >= 0; i-- {
			mut = rttd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rttd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rttd *RightToTreatmentDelete) ExecX(ctx context.Context) int {
	n, err := rttd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rttd *RightToTreatmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: righttotreatment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: righttotreatment.FieldID,
			},
		},
	}
	if ps := rttd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rttd.driver, _spec)
}

// RightToTreatmentDeleteOne is the builder for deleting a single RightToTreatment entity.
type RightToTreatmentDeleteOne struct {
	rttd *RightToTreatmentDelete
}

// Exec executes the deletion query.
func (rttdo *RightToTreatmentDeleteOne) Exec(ctx context.Context) error {
	n, err := rttdo.rttd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{righttotreatment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rttdo *RightToTreatmentDeleteOne) ExecX(ctx context.Context) {
	rttdo.rttd.ExecX(ctx)
}
