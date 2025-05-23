// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fiber-be-template/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// NormalizedUsername holds the value of the "normalized_username" field.
	NormalizedUsername string `json:"normalized_username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// NormalizedEmail holds the value of the "normalized_email" field.
	NormalizedEmail string `json:"normalized_email,omitempty"`
	// EmailConfirmed holds the value of the "email_confirmed" field.
	EmailConfirmed bool `json:"email_confirmed,omitempty"`
	// PasswordHash holds the value of the "password_hash" field.
	PasswordHash string `json:"password_hash,omitempty"`
	// ConcurrencyStamp holds the value of the "concurrency_stamp" field.
	ConcurrencyStamp string `json:"concurrency_stamp,omitempty"`
	// SecurityStamp holds the value of the "security_stamp" field.
	SecurityStamp string `json:"security_stamp,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// PhoneNumberConfirmed holds the value of the "phone_number_confirmed" field.
	PhoneNumberConfirmed bool `json:"phone_number_confirmed,omitempty"`
	// TwoFactorEnabled holds the value of the "two_factor_enabled" field.
	TwoFactorEnabled bool `json:"two_factor_enabled,omitempty"`
	// LockoutEnabled holds the value of the "lockout_enabled" field.
	LockoutEnabled bool `json:"lockout_enabled,omitempty"`
	// AccessFailedCount holds the value of the "access_failed_count" field.
	AccessFailedCount int `json:"access_failed_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserRoles holds the value of the user_roles edge.
	UserRoles []*UserRole `json:"user_roles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserRolesOrErr returns the UserRoles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserRolesOrErr() ([]*UserRole, error) {
	if e.loadedTypes[0] {
		return e.UserRoles, nil
	}
	return nil, &NotLoadedError{edge: "user_roles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldEmailConfirmed, user.FieldPhoneNumberConfirmed, user.FieldTwoFactorEnabled, user.FieldLockoutEnabled:
			values[i] = new(sql.NullBool)
		case user.FieldAccessFailedCount:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldNormalizedUsername, user.FieldEmail, user.FieldNormalizedEmail, user.FieldPasswordHash, user.FieldConcurrencyStamp, user.FieldSecurityStamp, user.FieldPhoneNumber:
			values[i] = new(sql.NullString)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldNormalizedUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field normalized_username", values[i])
			} else if value.Valid {
				u.NormalizedUsername = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldNormalizedEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field normalized_email", values[i])
			} else if value.Valid {
				u.NormalizedEmail = value.String
			}
		case user.FieldEmailConfirmed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_confirmed", values[i])
			} else if value.Valid {
				u.EmailConfirmed = value.Bool
			}
		case user.FieldPasswordHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_hash", values[i])
			} else if value.Valid {
				u.PasswordHash = value.String
			}
		case user.FieldConcurrencyStamp:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field concurrency_stamp", values[i])
			} else if value.Valid {
				u.ConcurrencyStamp = value.String
			}
		case user.FieldSecurityStamp:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field security_stamp", values[i])
			} else if value.Valid {
				u.SecurityStamp = value.String
			}
		case user.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				u.PhoneNumber = value.String
			}
		case user.FieldPhoneNumberConfirmed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number_confirmed", values[i])
			} else if value.Valid {
				u.PhoneNumberConfirmed = value.Bool
			}
		case user.FieldTwoFactorEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field two_factor_enabled", values[i])
			} else if value.Valid {
				u.TwoFactorEnabled = value.Bool
			}
		case user.FieldLockoutEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field lockout_enabled", values[i])
			} else if value.Valid {
				u.LockoutEnabled = value.Bool
			}
		case user.FieldAccessFailedCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field access_failed_count", values[i])
			} else if value.Valid {
				u.AccessFailedCount = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryUserRoles queries the "user_roles" edge of the User entity.
func (u *User) QueryUserRoles() *UserRoleQuery {
	return NewUserClient(u.config).QueryUserRoles(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("normalized_username=")
	builder.WriteString(u.NormalizedUsername)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("normalized_email=")
	builder.WriteString(u.NormalizedEmail)
	builder.WriteString(", ")
	builder.WriteString("email_confirmed=")
	builder.WriteString(fmt.Sprintf("%v", u.EmailConfirmed))
	builder.WriteString(", ")
	builder.WriteString("password_hash=")
	builder.WriteString(u.PasswordHash)
	builder.WriteString(", ")
	builder.WriteString("concurrency_stamp=")
	builder.WriteString(u.ConcurrencyStamp)
	builder.WriteString(", ")
	builder.WriteString("security_stamp=")
	builder.WriteString(u.SecurityStamp)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(u.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("phone_number_confirmed=")
	builder.WriteString(fmt.Sprintf("%v", u.PhoneNumberConfirmed))
	builder.WriteString(", ")
	builder.WriteString("two_factor_enabled=")
	builder.WriteString(fmt.Sprintf("%v", u.TwoFactorEnabled))
	builder.WriteString(", ")
	builder.WriteString("lockout_enabled=")
	builder.WriteString(fmt.Sprintf("%v", u.LockoutEnabled))
	builder.WriteString(", ")
	builder.WriteString("access_failed_count=")
	builder.WriteString(fmt.Sprintf("%v", u.AccessFailedCount))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
