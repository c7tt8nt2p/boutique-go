// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CartsColumns holds the columns for the "carts" table.
	CartsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "item_id", Type: field.TypeInt64},
		{Name: "count", Type: field.TypeInt64},
	}
	// CartsTable holds the schema information for the "carts" table.
	CartsTable = &schema.Table{
		Name:       "carts",
		Columns:    CartsColumns,
		PrimaryKey: []*schema.Column{CartsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CartsTable,
	}
)

func init() {
}
