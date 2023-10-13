package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type CheckoutItem struct {
	ent.Schema
}

func (CheckoutItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("checkout_id", uuid.UUID{}),
		field.UUID("cart_item_id", uuid.UUID{}),
	}
}

func (CheckoutItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("checkout_id_owner", Checkout.Type).
			Ref("checkout_item").
			Field("checkout_id").
			Unique().
			Required(),
		edge.From("cart_item_id_owner", CartItem.Type).
			Ref("checkout_item").
			Field("cart_item_id").
			Unique().
			Required(),
	}
}
