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
	"github.com/kx-boutique/ent/generated/cartitem"
	"github.com/kx-boutique/ent/generated/predicate"
	"github.com/kx-boutique/ent/generated/product"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where appends a list predicates to the ProductUpdate builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProductUpdate) SetName(s string) *ProductUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProductUpdate) SetDescription(s string) *ProductUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetStock sets the "stock" field.
func (pu *ProductUpdate) SetStock(i int32) *ProductUpdate {
	pu.mutation.ResetStock()
	pu.mutation.SetStock(i)
	return pu
}

// AddStock adds i to the "stock" field.
func (pu *ProductUpdate) AddStock(i int32) *ProductUpdate {
	pu.mutation.AddStock(i)
	return pu
}

// SetUnitPrice sets the "unit_price" field.
func (pu *ProductUpdate) SetUnitPrice(f float64) *ProductUpdate {
	pu.mutation.ResetUnitPrice()
	pu.mutation.SetUnitPrice(f)
	return pu
}

// AddUnitPrice adds f to the "unit_price" field.
func (pu *ProductUpdate) AddUnitPrice(f float64) *ProductUpdate {
	pu.mutation.AddUnitPrice(f)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *ProductUpdate) SetCreatedAt(t time.Time) *ProductUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableCreatedAt(t *time.Time) *ProductUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProductUpdate) SetUpdatedAt(t time.Time) *ProductUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableUpdatedAt(t *time.Time) *ProductUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// AddCartItemIDs adds the "cart_item" edge to the CartItem entity by IDs.
func (pu *ProductUpdate) AddCartItemIDs(ids ...uuid.UUID) *ProductUpdate {
	pu.mutation.AddCartItemIDs(ids...)
	return pu
}

// AddCartItem adds the "cart_item" edges to the CartItem entity.
func (pu *ProductUpdate) AddCartItem(c ...*CartItem) *ProductUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCartItemIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// ClearCartItem clears all "cart_item" edges to the CartItem entity.
func (pu *ProductUpdate) ClearCartItem() *ProductUpdate {
	pu.mutation.ClearCartItem()
	return pu
}

// RemoveCartItemIDs removes the "cart_item" edge to CartItem entities by IDs.
func (pu *ProductUpdate) RemoveCartItemIDs(ids ...uuid.UUID) *ProductUpdate {
	pu.mutation.RemoveCartItemIDs(ids...)
	return pu
}

// RemoveCartItem removes "cart_item" edges to CartItem entities.
func (pu *ProductUpdate) RemoveCartItem(c ...*CartItem) *ProductUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCartItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.Stock(); ok {
		_spec.SetField(product.FieldStock, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.AddedStock(); ok {
		_spec.AddField(product.FieldStock, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.UnitPrice(); ok {
		_spec.SetField(product.FieldUnitPrice, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedUnitPrice(); ok {
		_spec.AddField(product.FieldUnitPrice, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(product.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.CartItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CartItemTable,
			Columns: []string{product.CartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedCartItemIDs(); len(nodes) > 0 && !pu.mutation.CartItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CartItemTable,
			Columns: []string{product.CartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CartItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CartItemTable,
			Columns: []string{product.CartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductMutation
}

// SetName sets the "name" field.
func (puo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProductUpdateOne) SetDescription(s string) *ProductUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetStock sets the "stock" field.
func (puo *ProductUpdateOne) SetStock(i int32) *ProductUpdateOne {
	puo.mutation.ResetStock()
	puo.mutation.SetStock(i)
	return puo
}

// AddStock adds i to the "stock" field.
func (puo *ProductUpdateOne) AddStock(i int32) *ProductUpdateOne {
	puo.mutation.AddStock(i)
	return puo
}

// SetUnitPrice sets the "unit_price" field.
func (puo *ProductUpdateOne) SetUnitPrice(f float64) *ProductUpdateOne {
	puo.mutation.ResetUnitPrice()
	puo.mutation.SetUnitPrice(f)
	return puo
}

// AddUnitPrice adds f to the "unit_price" field.
func (puo *ProductUpdateOne) AddUnitPrice(f float64) *ProductUpdateOne {
	puo.mutation.AddUnitPrice(f)
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *ProductUpdateOne) SetCreatedAt(t time.Time) *ProductUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableCreatedAt(t *time.Time) *ProductUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProductUpdateOne) SetUpdatedAt(t time.Time) *ProductUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableUpdatedAt(t *time.Time) *ProductUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// AddCartItemIDs adds the "cart_item" edge to the CartItem entity by IDs.
func (puo *ProductUpdateOne) AddCartItemIDs(ids ...uuid.UUID) *ProductUpdateOne {
	puo.mutation.AddCartItemIDs(ids...)
	return puo
}

// AddCartItem adds the "cart_item" edges to the CartItem entity.
func (puo *ProductUpdateOne) AddCartItem(c ...*CartItem) *ProductUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCartItemIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// ClearCartItem clears all "cart_item" edges to the CartItem entity.
func (puo *ProductUpdateOne) ClearCartItem() *ProductUpdateOne {
	puo.mutation.ClearCartItem()
	return puo
}

// RemoveCartItemIDs removes the "cart_item" edge to CartItem entities by IDs.
func (puo *ProductUpdateOne) RemoveCartItemIDs(ids ...uuid.UUID) *ProductUpdateOne {
	puo.mutation.RemoveCartItemIDs(ids...)
	return puo
}

// RemoveCartItem removes "cart_item" edges to CartItem entities.
func (puo *ProductUpdateOne) RemoveCartItem(c ...*CartItem) *ProductUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCartItemIDs(ids...)
}

// Where appends a list predicates to the ProductUpdate builder.
func (puo *ProductUpdateOne) Where(ps ...predicate.Product) *ProductUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProductUpdateOne) Select(field string, fields ...string) *ProductUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Product entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Product.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for _, f := range fields {
			if !product.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.Stock(); ok {
		_spec.SetField(product.FieldStock, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.AddedStock(); ok {
		_spec.AddField(product.FieldStock, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.UnitPrice(); ok {
		_spec.SetField(product.FieldUnitPrice, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedUnitPrice(); ok {
		_spec.AddField(product.FieldUnitPrice, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(product.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.CartItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CartItemTable,
			Columns: []string{product.CartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedCartItemIDs(); len(nodes) > 0 && !puo.mutation.CartItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CartItemTable,
			Columns: []string{product.CartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CartItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CartItemTable,
			Columns: []string{product.CartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Product{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
