package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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
	return []ent.Edge{}
}