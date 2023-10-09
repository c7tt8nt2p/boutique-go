package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Cart struct {
	ent.Schema
}

func (Cart) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
	}
}

func (Cart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_id_owner", User.Type).
			Ref("cart").
			Field("user_id").
			Unique().
			Required(),
		edge.To("cart_item", CartItem.Type),
	}
}
