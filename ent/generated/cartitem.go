// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kx-boutique/ent/generated/cart"
	"github.com/kx-boutique/ent/generated/cartitem"
	"github.com/kx-boutique/ent/generated/checkoutitem"
	"github.com/kx-boutique/ent/generated/product"
)

// CartItem is the model entity for the CartItem schema.
type CartItem struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CartID holds the value of the "cart_id" field.
	CartID uuid.UUID `json:"cart_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID uuid.UUID `json:"product_id,omitempty"`
	// Qty holds the value of the "qty" field.
	Qty int32 `json:"qty,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CartItemQuery when eager-loading is set.
	Edges        CartItemEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CartItemEdges holds the relations/edges for other nodes in the graph.
type CartItemEdges struct {
	// CheckoutItem holds the value of the checkout_item edge.
	CheckoutItem *CheckoutItem `json:"checkout_item,omitempty"`
	// CartIDOwner holds the value of the cart_id_owner edge.
	CartIDOwner *Cart `json:"cart_id_owner,omitempty"`
	// ProductIDOwner holds the value of the product_id_owner edge.
	ProductIDOwner *Product `json:"product_id_owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CheckoutItemOrErr returns the CheckoutItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CartItemEdges) CheckoutItemOrErr() (*CheckoutItem, error) {
	if e.loadedTypes[0] {
		if e.CheckoutItem == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: checkoutitem.Label}
		}
		return e.CheckoutItem, nil
	}
	return nil, &NotLoadedError{edge: "checkout_item"}
}

// CartIDOwnerOrErr returns the CartIDOwner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CartItemEdges) CartIDOwnerOrErr() (*Cart, error) {
	if e.loadedTypes[1] {
		if e.CartIDOwner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: cart.Label}
		}
		return e.CartIDOwner, nil
	}
	return nil, &NotLoadedError{edge: "cart_id_owner"}
}

// ProductIDOwnerOrErr returns the ProductIDOwner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CartItemEdges) ProductIDOwnerOrErr() (*Product, error) {
	if e.loadedTypes[2] {
		if e.ProductIDOwner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.ProductIDOwner, nil
	}
	return nil, &NotLoadedError{edge: "product_id_owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CartItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cartitem.FieldQty:
			values[i] = new(sql.NullInt64)
		case cartitem.FieldCreatedAt, cartitem.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case cartitem.FieldID, cartitem.FieldCartID, cartitem.FieldProductID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CartItem fields.
func (ci *CartItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cartitem.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ci.ID = *value
			}
		case cartitem.FieldCartID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field cart_id", values[i])
			} else if value != nil {
				ci.CartID = *value
			}
		case cartitem.FieldProductID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value != nil {
				ci.ProductID = *value
			}
		case cartitem.FieldQty:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field qty", values[i])
			} else if value.Valid {
				ci.Qty = int32(value.Int64)
			}
		case cartitem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ci.CreatedAt = value.Time
			}
		case cartitem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ci.UpdatedAt = value.Time
			}
		default:
			ci.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CartItem.
// This includes values selected through modifiers, order, etc.
func (ci *CartItem) Value(name string) (ent.Value, error) {
	return ci.selectValues.Get(name)
}

// QueryCheckoutItem queries the "checkout_item" edge of the CartItem entity.
func (ci *CartItem) QueryCheckoutItem() *CheckoutItemQuery {
	return NewCartItemClient(ci.config).QueryCheckoutItem(ci)
}

// QueryCartIDOwner queries the "cart_id_owner" edge of the CartItem entity.
func (ci *CartItem) QueryCartIDOwner() *CartQuery {
	return NewCartItemClient(ci.config).QueryCartIDOwner(ci)
}

// QueryProductIDOwner queries the "product_id_owner" edge of the CartItem entity.
func (ci *CartItem) QueryProductIDOwner() *ProductQuery {
	return NewCartItemClient(ci.config).QueryProductIDOwner(ci)
}

// Update returns a builder for updating this CartItem.
// Note that you need to call CartItem.Unwrap() before calling this method if this CartItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (ci *CartItem) Update() *CartItemUpdateOne {
	return NewCartItemClient(ci.config).UpdateOne(ci)
}

// Unwrap unwraps the CartItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ci *CartItem) Unwrap() *CartItem {
	_tx, ok := ci.config.driver.(*txDriver)
	if !ok {
		panic("generated: CartItem is not a transactional entity")
	}
	ci.config.driver = _tx.drv
	return ci
}

// String implements the fmt.Stringer.
func (ci *CartItem) String() string {
	var builder strings.Builder
	builder.WriteString("CartItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ci.ID))
	builder.WriteString("cart_id=")
	builder.WriteString(fmt.Sprintf("%v", ci.CartID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", ci.ProductID))
	builder.WriteString(", ")
	builder.WriteString("qty=")
	builder.WriteString(fmt.Sprintf("%v", ci.Qty))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ci.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ci.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CartItems is a parsable slice of CartItem.
type CartItems []*CartItem
