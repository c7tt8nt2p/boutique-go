// Code generated by ent, DO NOT EDIT.

package cart

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the cart type in the database.
	Label = "cart"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldItemID holds the string denoting the item_id field in the database.
	FieldItemID = "item_id"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// Table holds the table name of the cart in the database.
	Table = "carts"
)

// Columns holds all SQL columns for cart fields.
var Columns = []string{
	FieldID,
	FieldItemID,
	FieldCount,
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

// OrderOption defines the ordering options for the Cart queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByItemID orders the results by the item_id field.
func ByItemID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldItemID, opts...).ToFunc()
}

// ByCount orders the results by the count field.
func ByCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCount, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}