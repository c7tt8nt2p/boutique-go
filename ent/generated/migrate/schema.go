// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthsColumns holds the columns for the "auths" table.
	AuthsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeUUID, Unique: true},
	}
	// AuthsTable holds the schema information for the "auths" table.
	AuthsTable = &schema.Table{
		Name:       "auths",
		Columns:    AuthsColumns,
		PrimaryKey: []*schema.Column{AuthsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "auths_users_auth",
				Columns:    []*schema.Column{AuthsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
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
		{Name: "checked_out", Type: field.TypeBool, Default: false},
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
				Columns:    []*schema.Column{CartItemsColumns[5]},
				RefColumns: []*schema.Column{CartsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "cart_items_products_cart_item",
				Columns:    []*schema.Column{CartItemsColumns[6]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// CheckoutsColumns holds the columns for the "checkouts" table.
	CheckoutsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "total_price", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// CheckoutsTable holds the schema information for the "checkouts" table.
	CheckoutsTable = &schema.Table{
		Name:       "checkouts",
		Columns:    CheckoutsColumns,
		PrimaryKey: []*schema.Column{CheckoutsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "checkouts_users_checkout",
				Columns:    []*schema.Column{CheckoutsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// CheckoutItemsColumns holds the columns for the "checkout_items" table.
	CheckoutItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "cart_item_id", Type: field.TypeUUID, Unique: true},
		{Name: "checkout_id", Type: field.TypeUUID},
	}
	// CheckoutItemsTable holds the schema information for the "checkout_items" table.
	CheckoutItemsTable = &schema.Table{
		Name:       "checkout_items",
		Columns:    CheckoutItemsColumns,
		PrimaryKey: []*schema.Column{CheckoutItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "checkout_items_cart_items_checkout_item",
				Columns:    []*schema.Column{CheckoutItemsColumns[1]},
				RefColumns: []*schema.Column{CartItemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "checkout_items_checkouts_checkout_item",
				Columns:    []*schema.Column{CheckoutItemsColumns[2]},
				RefColumns: []*schema.Column{CheckoutsColumns[0]},
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
		AuthsTable,
		CartsTable,
		CartItemsTable,
		CheckoutsTable,
		CheckoutItemsTable,
		ProductsTable,
		UsersTable,
	}
)

func init() {
	AuthsTable.ForeignKeys[0].RefTable = UsersTable
	CartsTable.ForeignKeys[0].RefTable = UsersTable
	CartItemsTable.ForeignKeys[0].RefTable = CartsTable
	CartItemsTable.ForeignKeys[1].RefTable = ProductsTable
	CheckoutsTable.ForeignKeys[0].RefTable = UsersTable
	CheckoutItemsTable.ForeignKeys[0].RefTable = CartItemsTable
	CheckoutItemsTable.ForeignKeys[1].RefTable = CheckoutsTable
}
