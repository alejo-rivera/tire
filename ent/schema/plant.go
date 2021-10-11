package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Plant holds the schema definition for the Plant entity.
type Plant struct {
	ent.Schema
}

// Fields of the Plant.
func (Plant) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),
		field.Time("birthdate"),
	}
}

// Edges of the Plant.
func (Plant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("plants").Unique(),
	}
}
