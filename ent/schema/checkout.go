package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Checkout struct {
	ent.Schema
}

func (Checkout) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.Float("total_price"),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Checkout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("checkout_item", CheckoutItem.Type),
		edge.From("user_id_owner", User.Type).
			Ref("checkout").
			Field("user_id").
			Unique().
			Required(),
	}
}
