// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kx-boutique/app/cart/service/internal/data/ent/generated/cart"
	"github.com/kx-boutique/app/cart/service/internal/data/ent/generated/predicate"
	"github.com/google/uuid"
)

// CartUpdate is the builder for updating Cart entities.
type CartUpdate struct {
	config
	hooks    []Hook
	mutation *CartMutation
}

// Where appends a list predicates to the CartUpdate builder.
func (cu *CartUpdate) Where(ps ...predicate.Cart) *CartUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetItemID sets the "item_id" field.
func (cu *CartUpdate) SetItemID(i int64) *CartUpdate {
	cu.mutation.ResetItemID()
	cu.mutation.SetItemID(i)
	return cu
}

// AddItemID adds i to the "item_id" field.
func (cu *CartUpdate) AddItemID(i int64) *CartUpdate {
	cu.mutation.AddItemID(i)
	return cu
}

// SetCount sets the "count" field.
func (cu *CartUpdate) SetCount(i int64) *CartUpdate {
	cu.mutation.ResetCount()
	cu.mutation.SetCount(i)
	return cu
}

// AddCount adds i to the "count" field.
func (cu *CartUpdate) AddCount(i int64) *CartUpdate {
	cu.mutation.AddCount(i)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CartUpdate) SetUserID(u uuid.UUID) *CartUpdate {
	cu.mutation.SetUserID(u)
	return cu
}

// Mutation returns the CartMutation object of the builder.
func (cu *CartUpdate) Mutation() *CartMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CartUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CartUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CartUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CartUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CartUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(cart.Table, cart.Columns, sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.ItemID(); ok {
		_spec.SetField(cart.FieldItemID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedItemID(); ok {
		_spec.AddField(cart.FieldItemID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.Count(); ok {
		_spec.SetField(cart.FieldCount, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedCount(); ok {
		_spec.AddField(cart.FieldCount, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.UserID(); ok {
		_spec.SetField(cart.FieldUserID, field.TypeUUID, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CartUpdateOne is the builder for updating a single Cart entity.
type CartUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CartMutation
}

// SetItemID sets the "item_id" field.
func (cuo *CartUpdateOne) SetItemID(i int64) *CartUpdateOne {
	cuo.mutation.ResetItemID()
	cuo.mutation.SetItemID(i)
	return cuo
}

// AddItemID adds i to the "item_id" field.
func (cuo *CartUpdateOne) AddItemID(i int64) *CartUpdateOne {
	cuo.mutation.AddItemID(i)
	return cuo
}

// SetCount sets the "count" field.
func (cuo *CartUpdateOne) SetCount(i int64) *CartUpdateOne {
	cuo.mutation.ResetCount()
	cuo.mutation.SetCount(i)
	return cuo
}

// AddCount adds i to the "count" field.
func (cuo *CartUpdateOne) AddCount(i int64) *CartUpdateOne {
	cuo.mutation.AddCount(i)
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *CartUpdateOne) SetUserID(u uuid.UUID) *CartUpdateOne {
	cuo.mutation.SetUserID(u)
	return cuo
}

// Mutation returns the CartMutation object of the builder.
func (cuo *CartUpdateOne) Mutation() *CartMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CartUpdate builder.
func (cuo *CartUpdateOne) Where(ps ...predicate.Cart) *CartUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CartUpdateOne) Select(field string, fields ...string) *CartUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cart entity.
func (cuo *CartUpdateOne) Save(ctx context.Context) (*Cart, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CartUpdateOne) SaveX(ctx context.Context) *Cart {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CartUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CartUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CartUpdateOne) sqlSave(ctx context.Context) (_node *Cart, err error) {
	_spec := sqlgraph.NewUpdateSpec(cart.Table, cart.Columns, sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Cart.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cart.FieldID)
		for _, f := range fields {
			if !cart.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != cart.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.ItemID(); ok {
		_spec.SetField(cart.FieldItemID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedItemID(); ok {
		_spec.AddField(cart.FieldItemID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.Count(); ok {
		_spec.SetField(cart.FieldCount, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedCount(); ok {
		_spec.AddField(cart.FieldCount, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.UserID(); ok {
		_spec.SetField(cart.FieldUserID, field.TypeUUID, value)
	}
	_node = &Cart{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
