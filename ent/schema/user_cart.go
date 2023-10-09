package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type UserCart struct {
	ent.Schema
}

func (UserCart) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_cart"},
	}
}

func (UserCart) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
	}
}

func (UserCart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_owner", User.Type).
			Ref("user_cart").
			Field("user_id").
			Unique().
			Required(),
		edge.To("carts", Cart.Type),
	}
}
