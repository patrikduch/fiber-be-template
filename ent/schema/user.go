package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"  // Add this import for entsql
	"entgo.io/ent/schema"          // Add this import for schema
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
            Default(uuid.New).
            Immutable().
            StorageKey("Id"),

		field.String("username").NotEmpty().StorageKey("UserName"),
		field.String("normalized_username").NotEmpty().StorageKey("NormalizedUserName"),
		field.String("email").NotEmpty().StorageKey("Email"),
		field.String("normalized_email").NotEmpty().StorageKey("NormalizedEmail"),
		field.Bool("email_confirmed").Default(false).StorageKey("EmailConfirmed"),
		field.String("password_hash").NotEmpty().StorageKey("PasswordHash"),
		field.String("concurrency_stamp").Optional().StorageKey("ConcurrencyStamp"),
		field.String("security_stamp").Optional().StorageKey("SecurityStamp"),
		field.String("phone_number").Optional().StorageKey("PhoneNumber"),
		field.Bool("phone_number_confirmed").Default(false).StorageKey("PhoneNumberConfirmed"),
		field.Bool("two_factor_enabled").Default(false).StorageKey("TwoFactorEnabled"),
		field.Bool("lockout_enabled").Default(false).StorageKey("LockoutEnabled"),
		field.Int("access_failed_count").Default(0).StorageKey("AccessFailedCount"),
	}
}

func (User) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "User"}, // Set table name to "User" with capital U
    }
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_roles", UserRole.Type),
	}
}
