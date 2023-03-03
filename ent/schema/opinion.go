package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Opinion holds the schema definition for the Opinion entity.
type Opinion struct {
	ent.Schema
}

// Fields of the Opinion.
func (Opinion) Fields() []ent.Field {
	return []ent.Field{
		field.Text("asunto").
			NotEmpty(),
		field.Text("contenido").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Enum("status").
			NamedValues(
				"Activo", "ACTIVO",
				"Inactivo", "INACTIVO",
			).
			Default("ACTIVO"),
	}
}

// Edges of the Opinion.
func (Opinion) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("parent", Opinion.Type).
            Unique().
            From("children"),
    }
}
