// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kx-boutique/ent/generated/cart"
)

// CartCreate is the builder for creating a Cart entity.
type CartCreate struct {
	config
	mutation *CartMutation
	hooks    []Hook
}

// SetCartID sets the "cart_id" field.
func (cc *CartCreate) SetCartID(u uuid.UUID) *CartCreate {
	cc.mutation.SetCartID(u)
	return cc
}

// SetProductID sets the "product_id" field.
func (cc *CartCreate) SetProductID(u uuid.UUID) *CartCreate {
	cc.mutation.SetProductID(u)
	return cc
}

// SetQty sets the "qty" field.
func (cc *CartCreate) SetQty(i int64) *CartCreate {
	cc.mutation.SetQty(i)
	return cc
}

// SetID sets the "id" field.
func (cc *CartCreate) SetID(u uuid.UUID) *CartCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CartCreate) SetNillableID(u *uuid.UUID) *CartCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the CartMutation object of the builder.
func (cc *CartCreate) Mutation() *CartMutation {
	return cc.mutation
}

// Save creates the Cart in the database.
func (cc *CartCreate) Save(ctx context.Context) (*Cart, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CartCreate) SaveX(ctx context.Context) *Cart {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CartCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CartCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CartCreate) defaults() {
	if _, ok := cc.mutation.ID(); !ok {
		v := cart.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CartCreate) check() error {
	if _, ok := cc.mutation.CartID(); !ok {
		return &ValidationError{Name: "cart_id", err: errors.New(`generated: missing required field "Cart.cart_id"`)}
	}
	if _, ok := cc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`generated: missing required field "Cart.product_id"`)}
	}
	if _, ok := cc.mutation.Qty(); !ok {
		return &ValidationError{Name: "qty", err: errors.New(`generated: missing required field "Cart.qty"`)}
	}
	return nil
}

func (cc *CartCreate) sqlSave(ctx context.Context) (*Cart, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CartCreate) createSpec() (*Cart, *sqlgraph.CreateSpec) {
	var (
		_node = &Cart{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(cart.Table, sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CartID(); ok {
		_spec.SetField(cart.FieldCartID, field.TypeUUID, value)
		_node.CartID = value
	}
	if value, ok := cc.mutation.ProductID(); ok {
		_spec.SetField(cart.FieldProductID, field.TypeUUID, value)
		_node.ProductID = value
	}
	if value, ok := cc.mutation.Qty(); ok {
		_spec.SetField(cart.FieldQty, field.TypeInt64, value)
		_node.Qty = value
	}
	return _node, _spec
}

// CartCreateBulk is the builder for creating many Cart entities in bulk.
type CartCreateBulk struct {
	config
	err      error
	builders []*CartCreate
}

// Save creates the Cart entities in the database.
func (ccb *CartCreateBulk) Save(ctx context.Context) ([]*Cart, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Cart, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CartMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CartCreateBulk) SaveX(ctx context.Context) []*Cart {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CartCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CartCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}