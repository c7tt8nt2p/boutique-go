package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type CartItem struct {
	ent.Schema
}

func (CartItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("cart_id", uuid.UUID{}),
		field.UUID("product_id", uuid.UUID{}),
		field.Int32("qty"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

func (CartItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cart_id_owner", Cart.Type).
			Ref("cart_item").
			Field("cart_id").
			Unique().
			Required(),
		edge.From("product_id_owner", Product.Type).
			Ref("cart_item").
			Field("product_id").
			Unique().
			Required(),
	}
}
