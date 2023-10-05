// Code generated by ent, DO NOT EDIT.

package generated

import (
	"time"

	"github.com/google/uuid"
	"github.com/kx-boutique/ent/generated/cart"
	"github.com/kx-boutique/ent/generated/product"
	"github.com/kx-boutique/ent/generated/user"
	"github.com/kx-boutique/ent/generated/usercart"
	"github.com/kx-boutique/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cartFields := schema.Cart{}.Fields()
	_ = cartFields
	// cartDescID is the schema descriptor for id field.
	cartDescID := cartFields[0].Descriptor()
	// cart.DefaultID holds the default value on creation for the id field.
	cart.DefaultID = cartDescID.Default.(func() uuid.UUID)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescCreatedAt is the schema descriptor for created_at field.
	productDescCreatedAt := productFields[5].Descriptor()
	// product.DefaultCreatedAt holds the default value on creation for the created_at field.
	product.DefaultCreatedAt = productDescCreatedAt.Default.(func() time.Time)
	// productDescUpdatedAt is the schema descriptor for updated_at field.
	productDescUpdatedAt := productFields[6].Descriptor()
	// product.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	product.DefaultUpdatedAt = productDescUpdatedAt.Default.(func() time.Time)
	// productDescID is the schema descriptor for id field.
	productDescID := productFields[0].Descriptor()
	// product.DefaultID holds the default value on creation for the id field.
	product.DefaultID = productDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[2].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[3].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	usercartFields := schema.UserCart{}.Fields()
	_ = usercartFields
	// usercartDescID is the schema descriptor for id field.
	usercartDescID := usercartFields[0].Descriptor()
	// usercart.DefaultID holds the default value on creation for the id field.
	usercart.DefaultID = usercartDescID.Default.(func() uuid.UUID)
}
