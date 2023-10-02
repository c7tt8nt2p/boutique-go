// Code generated by ent, DO NOT EDIT.

package cart

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/cart/service/internal/data/ent/generated/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Cart {
	return predicate.Cart(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Cart {
	return predicate.Cart(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Cart {
	return predicate.Cart(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Cart {
	return predicate.Cart(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Cart {
	return predicate.Cart(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Cart {
	return predicate.Cart(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Cart {
	return predicate.Cart(sql.FieldLTE(FieldID, id))
}

// ItemID applies equality check predicate on the "item_id" field. It's identical to ItemIDEQ.
func ItemID(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldItemID, v))
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldCount, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldUserID, v))
}

// ItemIDEQ applies the EQ predicate on the "item_id" field.
func ItemIDEQ(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldItemID, v))
}

// ItemIDNEQ applies the NEQ predicate on the "item_id" field.
func ItemIDNEQ(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldNEQ(FieldItemID, v))
}

// ItemIDIn applies the In predicate on the "item_id" field.
func ItemIDIn(vs ...int64) predicate.Cart {
	return predicate.Cart(sql.FieldIn(FieldItemID, vs...))
}

// ItemIDNotIn applies the NotIn predicate on the "item_id" field.
func ItemIDNotIn(vs ...int64) predicate.Cart {
	return predicate.Cart(sql.FieldNotIn(FieldItemID, vs...))
}

// ItemIDGT applies the GT predicate on the "item_id" field.
func ItemIDGT(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldGT(FieldItemID, v))
}

// ItemIDGTE applies the GTE predicate on the "item_id" field.
func ItemIDGTE(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldGTE(FieldItemID, v))
}

// ItemIDLT applies the LT predicate on the "item_id" field.
func ItemIDLT(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldLT(FieldItemID, v))
}

// ItemIDLTE applies the LTE predicate on the "item_id" field.
func ItemIDLTE(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldLTE(FieldItemID, v))
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldCount, v))
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldNEQ(FieldCount, v))
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int64) predicate.Cart {
	return predicate.Cart(sql.FieldIn(FieldCount, vs...))
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int64) predicate.Cart {
	return predicate.Cart(sql.FieldNotIn(FieldCount, vs...))
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldGT(FieldCount, v))
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldGTE(FieldCount, v))
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldLT(FieldCount, v))
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int64) predicate.Cart {
	return predicate.Cart(sql.FieldLTE(FieldCount, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.Cart {
	return predicate.Cart(sql.FieldLTE(FieldUserID, v))
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