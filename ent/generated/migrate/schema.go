// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CartsColumns holds the columns for the "carts" table.
	CartsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID, Unique: true},
	}
	// CartsTable holds the schema information for the "carts" table.
	CartsTable = &schema.Table{
		Name:       "carts",
		Columns:    CartsColumns,
		PrimaryKey: []*schema.Column{CartsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "carts_users_cart",
				Columns:    []*schema.Column{CartsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// CartItemsColumns holds the columns for the "cart_items" table.
	CartItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "qty", Type: field.TypeInt32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "cart_id", Type: field.TypeUUID},
		{Name: "product_id", Type: field.TypeUUID},
	}
	// CartItemsTable holds the schema information for the "cart_items" table.
	CartItemsTable = &schema.Table{
		Name:       "cart_items",
		Columns:    CartItemsColumns,
		PrimaryKey: []*schema.Column{CartItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cart_items_carts_cart_item",
				Columns:    []*schema.Column{CartItemsColumns[4]},
				RefColumns: []*schema.Column{CartsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "cart_items_products_cart_item",
				Columns:    []*schema.Column{CartItemsColumns[5]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "stock", Type: field.TypeInt32},
		{Name: "unit_price", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CartsTable,
		CartItemsTable,
		ProductsTable,
		UsersTable,
	}
)

func init() {
	CartsTable.ForeignKeys[0].RefTable = UsersTable
	CartItemsTable.ForeignKeys[0].RefTable = CartsTable
	CartItemsTable.ForeignKeys[1].RefTable = ProductsTable
}
