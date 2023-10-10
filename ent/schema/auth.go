package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Auth struct {
	ent.Schema
}

func (Auth) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.String("password_hash"),
	}
}

func (Auth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_id_owner", User.Type).
			Ref("auth").
			Field("user_id").
			Unique().
			Required(),
	}
}
