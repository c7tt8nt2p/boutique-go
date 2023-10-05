// Code generated by ent, DO NOT EDIT.

package cart

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kx-boutique/ent/generated/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldLTE(FieldID, id))
}

// CartID applies equality check predicate on the "cart_id" field. It's identical to CartIDEQ.
func CartID(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldCartID, v))
}

// ProductID applies equality check predicate on the "product_id" field. It's identical to ProductIDEQ.
func ProductID(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldProductID, v))
}

// Qty applies equality check predicate on the "qty" field. It's identical to QtyEQ.
func Qty(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldQty, v))
}

// CartIDEQ applies the EQ predicate on the "cart_id" field.
func CartIDEQ(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldCartID, v))
}

// CartIDNEQ applies the NEQ predicate on the "cart_id" field.
func CartIDNEQ(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldNEQ(FieldCartID, v))
}

// CartIDIn applies the In predicate on the "cart_id" field.
func CartIDIn(vs ...uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldIn(FieldCartID, vs...))
}

// CartIDNotIn applies the NotIn predicate on the "cart_id" field.
func CartIDNotIn(vs ...uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldNotIn(FieldCartID, vs...))
}

// CartIDGT applies the GT predicate on the "cart_id" field.
func CartIDGT(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldGT(FieldCartID, v))
}

// CartIDGTE applies the GTE predicate on the "cart_id" field.
func CartIDGTE(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldGTE(FieldCartID, v))
}

// CartIDLT applies the LT predicate on the "cart_id" field.
func CartIDLT(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldLT(FieldCartID, v))
}

// CartIDLTE applies the LTE predicate on the "cart_id" field.
func CartIDLTE(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldLTE(FieldCartID, v))
}

// ProductIDEQ applies the EQ predicate on the "product_id" field.
func ProductIDEQ(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldProductID, v))
}

// ProductIDNEQ applies the NEQ predicate on the "product_id" field.
func ProductIDNEQ(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldNEQ(FieldProductID, v))
}

// ProductIDIn applies the In predicate on the "product_id" field.
func ProductIDIn(vs ...uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldIn(FieldProductID, vs...))
}

// ProductIDNotIn applies the NotIn predicate on the "product_id" field.
func ProductIDNotIn(vs ...uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldNotIn(FieldProductID, vs...))
}

// ProductIDGT applies the GT predicate on the "product_id" field.
func ProductIDGT(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldGT(FieldProductID, v))
}

// ProductIDGTE applies the GTE predicate on the "product_id" field.
func ProductIDGTE(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldGTE(FieldProductID, v))
}

// ProductIDLT applies the LT predicate on the "product_id" field.
func ProductIDLT(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldLT(FieldProductID, v))
}

// ProductIDLTE applies the LTE predicate on the "product_id" field.
func ProductIDLTE(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldLTE(FieldProductID, v))
}

// QtyEQ applies the EQ predicate on the "qty" field.
func QtyEQ(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldQty, v))
}

// QtyNEQ applies the NEQ predicate on the "qty" field.
func QtyNEQ(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldNEQ(FieldQty, v))
}

// QtyIn applies the In predicate on the "qty" field.
func QtyIn(vs ...int64) predicate.Cart {
	return predicate.Cart(sql.FieldIn(FieldQty, vs...))
}

// QtyNotIn applies the NotIn predicate on the "qty" field.
func QtyNotIn(vs ...int64) predicate.Cart {
	return predicate.Cart(sql.FieldNotIn(FieldQty, vs...))
}

// QtyGT applies the GT predicate on the "qty" field.
func QtyGT(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldGT(FieldQty, v))
}

// QtyGTE applies the GTE predicate on the "qty" field.
func QtyGTE(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldGTE(FieldQty, v))
}

// QtyLT applies the LT predicate on the "qty" field.
func QtyLT(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldLT(FieldQty, v))
}

// QtyLTE applies the LTE predicate on the "qty" field.
func QtyLTE(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldLTE(FieldQty, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Cart) predicate.Cart {
	return predicate.Cart(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Cart) predicate.Cart {
	return predicate.Cart(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Cart) predicate.Cart {
	return predicate.Cart(sql.NotPredicates(p))
}
