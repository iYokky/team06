// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/team06/app/ent/predicate"
	"github.com/team06/app/ent/treatmenttype"
)

// TreatmentTypeDelete is the builder for deleting a TreatmentType entity.
type TreatmentTypeDelete struct {
	config
	hooks    []Hook
	mutation *TreatmentTypeMutation
}

// Where adds a new predicate to the TreatmentTypeDelete builder.
func (ttd *TreatmentTypeDelete) Where(ps ...predicate.TreatmentType) *TreatmentTypeDelete {
	ttd.mutation.predicates = append(ttd.mutation.predicates, ps...)
	return ttd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ttd *TreatmentTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ttd.hooks) == 0 {
		affected, err = ttd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TreatmentTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ttd.mutation = mutation
			affected, err = ttd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttd.hooks) - 1; i >= 0; i-- {
			mut = ttd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttd *TreatmentTypeDelete) ExecX(ctx context.Context) int {
	n, err := ttd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ttd *TreatmentTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: treatmenttype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: treatmenttype.FieldID,
			},
		},
	}
	if ps := ttd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ttd.driver, _spec)
}

// TreatmentTypeDeleteOne is the builder for deleting a single TreatmentType entity.
type TreatmentTypeDeleteOne struct {
	ttd *TreatmentTypeDelete
}

// Exec executes the deletion query.
func (ttdo *TreatmentTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ttdo.ttd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{treatmenttype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ttdo *TreatmentTypeDeleteOne) ExecX(ctx context.Context) {
	ttdo.ttd.ExecX(ctx)
}