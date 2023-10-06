// Code generated by ent, DO NOT EDIT.

package usercart

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the usercart type in the database.
	Label = "user_cart"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeCarts holds the string denoting the carts edge name in mutations.
	EdgeCarts = "carts"
	// Table holds the table name of the usercart in the database.
	Table = "user_cart"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "user_cart"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_id"
	// CartsTable is the table that holds the carts relation/edge.
	CartsTable = "carts"
	// CartsInverseTable is the table name for the Cart entity.
	// It exists in this package in order to avoid circular dependency with the "cart" package.
	CartsInverseTable = "carts"
	// CartsColumn is the table column denoting the carts relation/edge.
	CartsColumn = "cart_id"
)

// Columns holds all SQL columns for usercart fields.
var Columns = []string{
	FieldID,
	FieldUserID,
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the UserCart queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByCartsCount orders the results by carts count.
func ByCartsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCartsStep(), opts...)
	}
}

// ByCarts orders the results by carts terms.
func ByCarts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCartsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, OwnerTable, OwnerColumn),
	)
}
func newCartsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CartsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CartsTable, CartsColumn),
	)
}