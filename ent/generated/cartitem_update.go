// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kx-boutique/ent/generated/cart"
	"github.com/kx-boutique/ent/generated/cartitem"
	"github.com/kx-boutique/ent/generated/predicate"
	"github.com/kx-boutique/ent/generated/product"
)

// CartItemUpdate is the builder for updating CartItem entities.
type CartItemUpdate struct {
	config
	hooks    []Hook
	mutation *CartItemMutation
}

// Where appends a list predicates to the CartItemUpdate builder.
func (ciu *CartItemUpdate) Where(ps ...predicate.CartItem) *CartItemUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetCartID sets the "cart_id" field.
func (ciu *CartItemUpdate) SetCartID(u uuid.UUID) *CartItemUpdate {
	ciu.mutation.SetCartID(u)
	return ciu
}

// SetProductID sets the "product_id" field.
func (ciu *CartItemUpdate) SetProductID(u uuid.UUID) *CartItemUpdate {
	ciu.mutation.SetProductID(u)
	return ciu
}

// SetQty sets the "qty" field.
func (ciu *CartItemUpdate) SetQty(i int32) *CartItemUpdate {
	ciu.mutation.ResetQty()
	ciu.mutation.SetQty(i)
	return ciu
}

// AddQty adds i to the "qty" field.
func (ciu *CartItemUpdate) AddQty(i int32) *CartItemUpdate {
	ciu.mutation.AddQty(i)
	return ciu
}

// SetCreatedAt sets the "created_at" field.
func (ciu *CartItemUpdate) SetCreatedAt(t time.Time) *CartItemUpdate {
	ciu.mutation.SetCreatedAt(t)
	return ciu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciu *CartItemUpdate) SetNillableCreatedAt(t *time.Time) *CartItemUpdate {
	if t != nil {
		ciu.SetCreatedAt(*t)
	}
	return ciu
}

// SetUpdatedAt sets the "updated_at" field.
func (ciu *CartItemUpdate) SetUpdatedAt(t time.Time) *CartItemUpdate {
	ciu.mutation.SetUpdatedAt(t)
	return ciu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ciu *CartItemUpdate) SetNillableUpdatedAt(t *time.Time) *CartItemUpdate {
	if t != nil {
		ciu.SetUpdatedAt(*t)
	}
	return ciu
}

// SetCartIDOwnerID sets the "cart_id_owner" edge to the Cart entity by ID.
func (ciu *CartItemUpdate) SetCartIDOwnerID(id uuid.UUID) *CartItemUpdate {
	ciu.mutation.SetCartIDOwnerID(id)
	return ciu
}

// SetCartIDOwner sets the "cart_id_owner" edge to the Cart entity.
func (ciu *CartItemUpdate) SetCartIDOwner(c *Cart) *CartItemUpdate {
	return ciu.SetCartIDOwnerID(c.ID)
}

// SetProductIDOwnerID sets the "product_id_owner" edge to the Product entity by ID.
func (ciu *CartItemUpdate) SetProductIDOwnerID(id uuid.UUID) *CartItemUpdate {
	ciu.mutation.SetProductIDOwnerID(id)
	return ciu
}

// SetProductIDOwner sets the "product_id_owner" edge to the Product entity.
func (ciu *CartItemUpdate) SetProductIDOwner(p *Product) *CartItemUpdate {
	return ciu.SetProductIDOwnerID(p.ID)
}

// Mutation returns the CartItemMutation object of the builder.
func (ciu *CartItemUpdate) Mutation() *CartItemMutation {
	return ciu.mutation
}

// ClearCartIDOwner clears the "cart_id_owner" edge to the Cart entity.
func (ciu *CartItemUpdate) ClearCartIDOwner() *CartItemUpdate {
	ciu.mutation.ClearCartIDOwner()
	return ciu
}

// ClearProductIDOwner clears the "product_id_owner" edge to the Product entity.
func (ciu *CartItemUpdate) ClearProductIDOwner() *CartItemUpdate {
	ciu.mutation.ClearProductIDOwner()
	return ciu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *CartItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ciu.sqlSave, ciu.mutation, ciu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *CartItemUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *CartItemUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *CartItemUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciu *CartItemUpdate) check() error {
	if _, ok := ciu.mutation.CartIDOwnerID(); ciu.mutation.CartIDOwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "CartItem.cart_id_owner"`)
	}
	if _, ok := ciu.mutation.ProductIDOwnerID(); ciu.mutation.ProductIDOwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "CartItem.product_id_owner"`)
	}
	return nil
}

func (ciu *CartItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ciu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(cartitem.Table, cartitem.Columns, sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID))
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.Qty(); ok {
		_spec.SetField(cartitem.FieldQty, field.TypeInt32, value)
	}
	if value, ok := ciu.mutation.AddedQty(); ok {
		_spec.AddField(cartitem.FieldQty, field.TypeInt32, value)
	}
	if value, ok := ciu.mutation.CreatedAt(); ok {
		_spec.SetField(cartitem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ciu.mutation.UpdatedAt(); ok {
		_spec.SetField(cartitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if ciu.mutation.CartIDOwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.CartIDOwnerTable,
			Columns: []string{cartitem.CartIDOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.CartIDOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.CartIDOwnerTable,
			Columns: []string{cartitem.CartIDOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciu.mutation.ProductIDOwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.ProductIDOwnerTable,
			Columns: []string{cartitem.ProductIDOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.ProductIDOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.ProductIDOwnerTable,
			Columns: []string{cartitem.ProductIDOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cartitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ciu.mutation.done = true
	return n, nil
}

// CartItemUpdateOne is the builder for updating a single CartItem entity.
type CartItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CartItemMutation
}

// SetCartID sets the "cart_id" field.
func (ciuo *CartItemUpdateOne) SetCartID(u uuid.UUID) *CartItemUpdateOne {
	ciuo.mutation.SetCartID(u)
	return ciuo
}

// SetProductID sets the "product_id" field.
func (ciuo *CartItemUpdateOne) SetProductID(u uuid.UUID) *CartItemUpdateOne {
	ciuo.mutation.SetProductID(u)
	return ciuo
}

// SetQty sets the "qty" field.
func (ciuo *CartItemUpdateOne) SetQty(i int32) *CartItemUpdateOne {
	ciuo.mutation.ResetQty()
	ciuo.mutation.SetQty(i)
	return ciuo
}

// AddQty adds i to the "qty" field.
func (ciuo *CartItemUpdateOne) AddQty(i int32) *CartItemUpdateOne {
	ciuo.mutation.AddQty(i)
	return ciuo
}

// SetCreatedAt sets the "created_at" field.
func (ciuo *CartItemUpdateOne) SetCreatedAt(t time.Time) *CartItemUpdateOne {
	ciuo.mutation.SetCreatedAt(t)
	return ciuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciuo *CartItemUpdateOne) SetNillableCreatedAt(t *time.Time) *CartItemUpdateOne {
	if t != nil {
		ciuo.SetCreatedAt(*t)
	}
	return ciuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ciuo *CartItemUpdateOne) SetUpdatedAt(t time.Time) *CartItemUpdateOne {
	ciuo.mutation.SetUpdatedAt(t)
	return ciuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ciuo *CartItemUpdateOne) SetNillableUpdatedAt(t *time.Time) *CartItemUpdateOne {
	if t != nil {
		ciuo.SetUpdatedAt(*t)
	}
	return ciuo
}

// SetCartIDOwnerID sets the "cart_id_owner" edge to the Cart entity by ID.
func (ciuo *CartItemUpdateOne) SetCartIDOwnerID(id uuid.UUID) *CartItemUpdateOne {
	ciuo.mutation.SetCartIDOwnerID(id)
	return ciuo
}

// SetCartIDOwner sets the "cart_id_owner" edge to the Cart entity.
func (ciuo *CartItemUpdateOne) SetCartIDOwner(c *Cart) *CartItemUpdateOne {
	return ciuo.SetCartIDOwnerID(c.ID)
}

// SetProductIDOwnerID sets the "product_id_owner" edge to the Product entity by ID.
func (ciuo *CartItemUpdateOne) SetProductIDOwnerID(id uuid.UUID) *CartItemUpdateOne {
	ciuo.mutation.SetProductIDOwnerID(id)
	return ciuo
}

// SetProductIDOwner sets the "product_id_owner" edge to the Product entity.
func (ciuo *CartItemUpdateOne) SetProductIDOwner(p *Product) *CartItemUpdateOne {
	return ciuo.SetProductIDOwnerID(p.ID)
}

// Mutation returns the CartItemMutation object of the builder.
func (ciuo *CartItemUpdateOne) Mutation() *CartItemMutation {
	return ciuo.mutation
}

// ClearCartIDOwner clears the "cart_id_owner" edge to the Cart entity.
func (ciuo *CartItemUpdateOne) ClearCartIDOwner() *CartItemUpdateOne {
	ciuo.mutation.ClearCartIDOwner()
	return ciuo
}

// ClearProductIDOwner clears the "product_id_owner" edge to the Product entity.
func (ciuo *CartItemUpdateOne) ClearProductIDOwner() *CartItemUpdateOne {
	ciuo.mutation.ClearProductIDOwner()
	return ciuo
}

// Where appends a list predicates to the CartItemUpdate builder.
func (ciuo *CartItemUpdateOne) Where(ps ...predicate.CartItem) *CartItemUpdateOne {
	ciuo.mutation.Where(ps...)
	return ciuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *CartItemUpdateOne) Select(field string, fields ...string) *CartItemUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated CartItem entity.
func (ciuo *CartItemUpdateOne) Save(ctx context.Context) (*CartItem, error) {
	return withHooks(ctx, ciuo.sqlSave, ciuo.mutation, ciuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *CartItemUpdateOne) SaveX(ctx context.Context) *CartItem {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *CartItemUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *CartItemUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciuo *CartItemUpdateOne) check() error {
	if _, ok := ciuo.mutation.CartIDOwnerID(); ciuo.mutation.CartIDOwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "CartItem.cart_id_owner"`)
	}
	if _, ok := ciuo.mutation.ProductIDOwnerID(); ciuo.mutation.ProductIDOwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "CartItem.product_id_owner"`)
	}
	return nil
}

func (ciuo *CartItemUpdateOne) sqlSave(ctx context.Context) (_node *CartItem, err error) {
	if err := ciuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(cartitem.Table, cartitem.Columns, sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID))
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "CartItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cartitem.FieldID)
		for _, f := range fields {
			if !cartitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != cartitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.Qty(); ok {
		_spec.SetField(cartitem.FieldQty, field.TypeInt32, value)
	}
	if value, ok := ciuo.mutation.AddedQty(); ok {
		_spec.AddField(cartitem.FieldQty, field.TypeInt32, value)
	}
	if value, ok := ciuo.mutation.CreatedAt(); ok {
		_spec.SetField(cartitem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ciuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cartitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if ciuo.mutation.CartIDOwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.CartIDOwnerTable,
			Columns: []string{cartitem.CartIDOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.CartIDOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.CartIDOwnerTable,
			Columns: []string{cartitem.CartIDOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciuo.mutation.ProductIDOwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.ProductIDOwnerTable,
			Columns: []string{cartitem.ProductIDOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.ProductIDOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.ProductIDOwnerTable,
			Columns: []string{cartitem.ProductIDOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CartItem{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cartitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ciuo.mutation.done = true
	return _node, nil
}
