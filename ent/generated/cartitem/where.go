// Code generated by ent, DO NOT EDIT.

package cartitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kx-boutique/ent/generated/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldLTE(FieldID, id))
}

// CartID applies equality check predicate on the "cart_id" field. It's identical to CartIDEQ.
func CartID(v uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldCartID, v))
}

// ProductID applies equality check predicate on the "product_id" field. It's identical to ProductIDEQ.
func ProductID(v uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldProductID, v))
}

// Qty applies equality check predicate on the "qty" field. It's identical to QtyEQ.
func Qty(v int32) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldQty, v))
}

// CheckedOut applies equality check predicate on the "checked_out" field. It's identical to CheckedOutEQ.
func CheckedOut(v bool) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldCheckedOut, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// CartIDEQ applies the EQ predicate on the "cart_id" field.
func CartIDEQ(v uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldCartID, v))
}

// CartIDNEQ applies the NEQ predicate on the "cart_id" field.
func CartIDNEQ(v uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldNEQ(FieldCartID, v))
}

// CartIDIn applies the In predicate on the "cart_id" field.
func CartIDIn(vs ...uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldIn(FieldCartID, vs...))
}

// CartIDNotIn applies the NotIn predicate on the "cart_id" field.
func CartIDNotIn(vs ...uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldNotIn(FieldCartID, vs...))
}

// ProductIDEQ applies the EQ predicate on the "product_id" field.
func ProductIDEQ(v uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldProductID, v))
}

// ProductIDNEQ applies the NEQ predicate on the "product_id" field.
func ProductIDNEQ(v uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldNEQ(FieldProductID, v))
}

// ProductIDIn applies the In predicate on the "product_id" field.
func ProductIDIn(vs ...uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldIn(FieldProductID, vs...))
}

// ProductIDNotIn applies the NotIn predicate on the "product_id" field.
func ProductIDNotIn(vs ...uuid.UUID) predicate.CartItem {
	return predicate.CartItem(sql.FieldNotIn(FieldProductID, vs...))
}

// QtyEQ applies the EQ predicate on the "qty" field.
func QtyEQ(v int32) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldQty, v))
}

// QtyNEQ applies the NEQ predicate on the "qty" field.
func QtyNEQ(v int32) predicate.CartItem {
	return predicate.CartItem(sql.FieldNEQ(FieldQty, v))
}

// QtyIn applies the In predicate on the "qty" field.
func QtyIn(vs ...int32) predicate.CartItem {
	return predicate.CartItem(sql.FieldIn(FieldQty, vs...))
}

// QtyNotIn applies the NotIn predicate on the "qty" field.
func QtyNotIn(vs ...int32) predicate.CartItem {
	return predicate.CartItem(sql.FieldNotIn(FieldQty, vs...))
}

// QtyGT applies the GT predicate on the "qty" field.
func QtyGT(v int32) predicate.CartItem {
	return predicate.CartItem(sql.FieldGT(FieldQty, v))
}

// QtyGTE applies the GTE predicate on the "qty" field.
func QtyGTE(v int32) predicate.CartItem {
	return predicate.CartItem(sql.FieldGTE(FieldQty, v))
}

// QtyLT applies the LT predicate on the "qty" field.
func QtyLT(v int32) predicate.CartItem {
	return predicate.CartItem(sql.FieldLT(FieldQty, v))
}

// QtyLTE applies the LTE predicate on the "qty" field.
func QtyLTE(v int32) predicate.CartItem {
	return predicate.CartItem(sql.FieldLTE(FieldQty, v))
}

// CheckedOutEQ applies the EQ predicate on the "checked_out" field.
func CheckedOutEQ(v bool) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldCheckedOut, v))
}

// CheckedOutNEQ applies the NEQ predicate on the "checked_out" field.
func CheckedOutNEQ(v bool) predicate.CartItem {
	return predicate.CartItem(sql.FieldNEQ(FieldCheckedOut, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CartItem {
	return predicate.CartItem(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasCheckoutItem applies the HasEdge predicate on the "checkout_item" edge.
func HasCheckoutItem() predicate.CartItem {
	return predicate.CartItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, CheckoutItemTable, CheckoutItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCheckoutItemWith applies the HasEdge predicate on the "checkout_item" edge with a given conditions (other predicates).
func HasCheckoutItemWith(preds ...predicate.CheckoutItem) predicate.CartItem {
	return predicate.CartItem(func(s *sql.Selector) {
		step := newCheckoutItemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCartIDOwner applies the HasEdge predicate on the "cart_id_owner" edge.
func HasCartIDOwner() predicate.CartItem {
	return predicate.CartItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CartIDOwnerTable, CartIDOwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCartIDOwnerWith applies the HasEdge predicate on the "cart_id_owner" edge with a given conditions (other predicates).
func HasCartIDOwnerWith(preds ...predicate.Cart) predicate.CartItem {
	return predicate.CartItem(func(s *sql.Selector) {
		step := newCartIDOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProductIDOwner applies the HasEdge predicate on the "product_id_owner" edge.
func HasProductIDOwner() predicate.CartItem {
	return predicate.CartItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductIDOwnerTable, ProductIDOwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductIDOwnerWith applies the HasEdge predicate on the "product_id_owner" edge with a given conditions (other predicates).
func HasProductIDOwnerWith(preds ...predicate.Product) predicate.CartItem {
	return predicate.CartItem(func(s *sql.Selector) {
		step := newProductIDOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CartItem) predicate.CartItem {
	return predicate.CartItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CartItem) predicate.CartItem {
	return predicate.CartItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CartItem) predicate.CartItem {
	return predicate.CartItem(sql.NotPredicates(p))
}
