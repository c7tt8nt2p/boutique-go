// Code generated by ent, DO NOT EDIT.

package cartitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the cartitem type in the database.
	Label = "cart_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCartID holds the string denoting the cart_id field in the database.
	FieldCartID = "cart_id"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldQty holds the string denoting the qty field in the database.
	FieldQty = "qty"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeCheckoutItem holds the string denoting the checkout_item edge name in mutations.
	EdgeCheckoutItem = "checkout_item"
	// EdgeCartIDOwner holds the string denoting the cart_id_owner edge name in mutations.
	EdgeCartIDOwner = "cart_id_owner"
	// EdgeProductIDOwner holds the string denoting the product_id_owner edge name in mutations.
	EdgeProductIDOwner = "product_id_owner"
	// Table holds the table name of the cartitem in the database.
	Table = "cart_items"
	// CheckoutItemTable is the table that holds the checkout_item relation/edge.
	CheckoutItemTable = "checkout_items"
	// CheckoutItemInverseTable is the table name for the CheckoutItem entity.
	// It exists in this package in order to avoid circular dependency with the "checkoutitem" package.
	CheckoutItemInverseTable = "checkout_items"
	// CheckoutItemColumn is the table column denoting the checkout_item relation/edge.
	CheckoutItemColumn = "cart_item_id"
	// CartIDOwnerTable is the table that holds the cart_id_owner relation/edge.
	CartIDOwnerTable = "cart_items"
	// CartIDOwnerInverseTable is the table name for the Cart entity.
	// It exists in this package in order to avoid circular dependency with the "cart" package.
	CartIDOwnerInverseTable = "carts"
	// CartIDOwnerColumn is the table column denoting the cart_id_owner relation/edge.
	CartIDOwnerColumn = "cart_id"
	// ProductIDOwnerTable is the table that holds the product_id_owner relation/edge.
	ProductIDOwnerTable = "cart_items"
	// ProductIDOwnerInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductIDOwnerInverseTable = "products"
	// ProductIDOwnerColumn is the table column denoting the product_id_owner relation/edge.
	ProductIDOwnerColumn = "product_id"
)

// Columns holds all SQL columns for cartitem fields.
var Columns = []string{
	FieldID,
	FieldCartID,
	FieldProductID,
	FieldQty,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the CartItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCartID orders the results by the cart_id field.
func ByCartID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCartID, opts...).ToFunc()
}

// ByProductID orders the results by the product_id field.
func ByProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductID, opts...).ToFunc()
}

// ByQty orders the results by the qty field.
func ByQty(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQty, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCheckoutItemField orders the results by checkout_item field.
func ByCheckoutItemField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCheckoutItemStep(), sql.OrderByField(field, opts...))
	}
}

// ByCartIDOwnerField orders the results by cart_id_owner field.
func ByCartIDOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCartIDOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByProductIDOwnerField orders the results by product_id_owner field.
func ByProductIDOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductIDOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newCheckoutItemStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CheckoutItemInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, CheckoutItemTable, CheckoutItemColumn),
	)
}
func newCartIDOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CartIDOwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CartIDOwnerTable, CartIDOwnerColumn),
	)
}
func newProductIDOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductIDOwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProductIDOwnerTable, ProductIDOwnerColumn),
	)
}
