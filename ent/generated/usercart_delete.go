// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kx-boutique/ent/generated/predicate"
	"github.com/kx-boutique/ent/generated/usercart"
)

// UserCartDelete is the builder for deleting a UserCart entity.
type UserCartDelete struct {
	config
	hooks    []Hook
	mutation *UserCartMutation
}

// Where appends a list predicates to the UserCartDelete builder.
func (ucd *UserCartDelete) Where(ps ...predicate.UserCart) *UserCartDelete {
	ucd.mutation.Where(ps...)
	return ucd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ucd *UserCartDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ucd.sqlExec, ucd.mutation, ucd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ucd *UserCartDelete) ExecX(ctx context.Context) int {
	n, err := ucd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ucd *UserCartDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usercart.Table, sqlgraph.NewFieldSpec(usercart.FieldID, field.TypeUUID))
	if ps := ucd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ucd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ucd.mutation.done = true
	return affected, err
}

// UserCartDeleteOne is the builder for deleting a single UserCart entity.
type UserCartDeleteOne struct {
	ucd *UserCartDelete
}

// Where appends a list predicates to the UserCartDelete builder.
func (ucdo *UserCartDeleteOne) Where(ps ...predicate.UserCart) *UserCartDeleteOne {
	ucdo.ucd.mutation.Where(ps...)
	return ucdo
}

// Exec executes the deletion query.
func (ucdo *UserCartDeleteOne) Exec(ctx context.Context) error {
	n, err := ucdo.ucd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usercart.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdo *UserCartDeleteOne) ExecX(ctx context.Context) {
	if err := ucdo.Exec(ctx); err != nil {
		panic(err)
	}
}
