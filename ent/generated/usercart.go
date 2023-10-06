// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kx-boutique/ent/generated/user"
	"github.com/kx-boutique/ent/generated/usercart"
)

// UserCart is the model entity for the UserCart schema.
type UserCart struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserCartQuery when eager-loading is set.
	Edges        UserCartEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserCartEdges holds the relations/edges for other nodes in the graph.
type UserCartEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Carts holds the value of the carts edge.
	Carts []*Cart `json:"carts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserCartEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// CartsOrErr returns the Carts value or an error if the edge
// was not loaded in eager-loading.
func (e UserCartEdges) CartsOrErr() ([]*Cart, error) {
	if e.loadedTypes[1] {
		return e.Carts, nil
	}
	return nil, &NotLoadedError{edge: "carts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserCart) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usercart.FieldID, usercart.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserCart fields.
func (uc *UserCart) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usercart.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				uc.ID = *value
			}
		case usercart.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				uc.UserID = *value
			}
		default:
			uc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserCart.
// This includes values selected through modifiers, order, etc.
func (uc *UserCart) Value(name string) (ent.Value, error) {
	return uc.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the UserCart entity.
func (uc *UserCart) QueryOwner() *UserQuery {
	return NewUserCartClient(uc.config).QueryOwner(uc)
}

// QueryCarts queries the "carts" edge of the UserCart entity.
func (uc *UserCart) QueryCarts() *CartQuery {
	return NewUserCartClient(uc.config).QueryCarts(uc)
}

// Update returns a builder for updating this UserCart.
// Note that you need to call UserCart.Unwrap() before calling this method if this UserCart
// was returned from a transaction, and the transaction was committed or rolled back.
func (uc *UserCart) Update() *UserCartUpdateOne {
	return NewUserCartClient(uc.config).UpdateOne(uc)
}

// Unwrap unwraps the UserCart entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uc *UserCart) Unwrap() *UserCart {
	_tx, ok := uc.config.driver.(*txDriver)
	if !ok {
		panic("generated: UserCart is not a transactional entity")
	}
	uc.config.driver = _tx.drv
	return uc
}

// String implements the fmt.Stringer.
func (uc *UserCart) String() string {
	var builder strings.Builder
	builder.WriteString("UserCart(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uc.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", uc.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// UserCarts is a parsable slice of UserCart.
type UserCarts []*UserCart
