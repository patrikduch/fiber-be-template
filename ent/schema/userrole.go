package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type UserRole struct {
	ent.Schema
}

func (UserRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "UserRole"},
	}
}

// Define the fields exactly as they appear in your database
func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		// This is the missing ID field (UUID or int, depending on your DB)
		field.UUID("id", uuid.UUID{}).
			StorageKey("Id").
			Default(uuid.New), // or omit `.Default()` if DB handles it

		field.UUID("user_id", uuid.UUID{}).
			StorageKey("UserId"),
		field.UUID("role_id", uuid.UUID{}).
			StorageKey("RoleId"),
	}
}

func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_roles").
			Field("user_id").
			Unique().
			Required(),
		edge.From("role", Role.Type).
			Ref("user_roles").
			Field("role_id").
			Unique().
			Required(),
	}
}

// Tell Ent that this entity uses a composite primary key
func (UserRole) ID() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("role_id", uuid.UUID{}),
	}
}

// Define a unique index on the composite key
func (UserRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "role_id").
			Unique(),
	}
}