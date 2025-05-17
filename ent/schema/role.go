package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"  // Add this import
	"entgo.io/ent/schema"          // Add this import
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			StorageKey("Id"),

		field.String("name").NotEmpty().StorageKey("Name"),
		field.String("normalized_name").NotEmpty().StorageKey("NormalizedName"),
		field.String("concurrency_stamp").Optional().StorageKey("ConcurrencyStamp"),
	}
}


// Add this Annotations method to set the table name
func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Role"},
	}
}


func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_roles", UserRole.Type),
	}
}