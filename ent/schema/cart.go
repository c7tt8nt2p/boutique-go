package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Cart struct {
	ent.Schema
}

func (Cart) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("cart_id", uuid.UUID{}),
		field.UUID("product_id", uuid.UUID{}),
		field.Int64("qty"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}

// Edges of the Cart.
func (Cart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner_cart_id", UserCart.Type).
			Ref("carts").
			Field("cart_id").
			Unique().
			Required(),
		edge.From("owner_product_id", Product.Type).
			Ref("carts").
			Field("product_id").
			Unique().
			Required(),
	}
}
