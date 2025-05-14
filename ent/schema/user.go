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
            StorageKey("Id"),

        field.String("username").
            NotEmpty().
            StorageKey("UserName"),

        field.String("email").
            Unique().
            NotEmpty().
            StorageKey("Email"),

        field.String("normalized_email").
            NotEmpty().
            StorageKey("NormalizedEmail"),

        field.String("password_hash").
            Sensitive().
            NotEmpty().
            StorageKey("PasswordHash"),

        field.Bool("email_confirmed").
            Default(false).
            StorageKey("EmailConfirmed"),

        field.Bool("phone_number_confirmed").
            Default(false).
            StorageKey("PhoneNumberConfirmed"),

        field.Bool("two_factor_enabled").
            Default(false).
            StorageKey("TwoFactorEnabled"),

        field.Bool("lockout_enabled").
            Default(false).
            StorageKey("LockoutEnabled"),

        field.Int("access_failed_count").
            Default(0).
            StorageKey("AccessFailedCount"),
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
