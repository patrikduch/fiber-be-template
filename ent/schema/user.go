package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).
            Default(uuid.New).
            Immutable().
            StorageKey("Id"), // Maps to "Id" column
            
        field.String("username").
            NotEmpty().
            StorageKey("UserName"), // Maps to "UserName" column
            
        field.String("email").
            Unique().
            NotEmpty().
            StorageKey("Email"), // Maps to "Email" column
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return nil
}

// Config returns the configuration for the User schema.
func (User) Config() ent.Config {
    return ent.Config{
        Table: "User",
    }
}