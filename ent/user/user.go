// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "Id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "UserName"
	// FieldNormalizedUsername holds the string denoting the normalized_username field in the database.
	FieldNormalizedUsername = "NormalizedUserName"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "Email"
	// FieldNormalizedEmail holds the string denoting the normalized_email field in the database.
	FieldNormalizedEmail = "NormalizedEmail"
	// FieldEmailConfirmed holds the string denoting the email_confirmed field in the database.
	FieldEmailConfirmed = "EmailConfirmed"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "PasswordHash"
	// FieldConcurrencyStamp holds the string denoting the concurrency_stamp field in the database.
	FieldConcurrencyStamp = "ConcurrencyStamp"
	// FieldSecurityStamp holds the string denoting the security_stamp field in the database.
	FieldSecurityStamp = "SecurityStamp"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "PhoneNumber"
	// FieldPhoneNumberConfirmed holds the string denoting the phone_number_confirmed field in the database.
	FieldPhoneNumberConfirmed = "PhoneNumberConfirmed"
	// FieldTwoFactorEnabled holds the string denoting the two_factor_enabled field in the database.
	FieldTwoFactorEnabled = "TwoFactorEnabled"
	// FieldLockoutEnabled holds the string denoting the lockout_enabled field in the database.
	FieldLockoutEnabled = "LockoutEnabled"
	// FieldAccessFailedCount holds the string denoting the access_failed_count field in the database.
	FieldAccessFailedCount = "AccessFailedCount"
	// EdgeUserRoles holds the string denoting the user_roles edge name in mutations.
	EdgeUserRoles = "user_roles"
	// Table holds the table name of the user in the database.
	Table = "User"
	// UserRolesTable is the table that holds the user_roles relation/edge.
	UserRolesTable = "UserRole"
	// UserRolesInverseTable is the table name for the UserRole entity.
	// It exists in this package in order to avoid circular dependency with the "userrole" package.
	UserRolesInverseTable = "UserRole"
	// UserRolesColumn is the table column denoting the user_roles relation/edge.
	UserRolesColumn = "UserId"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldNormalizedUsername,
	FieldEmail,
	FieldNormalizedEmail,
	FieldEmailConfirmed,
	FieldPasswordHash,
	FieldConcurrencyStamp,
	FieldSecurityStamp,
	FieldPhoneNumber,
	FieldPhoneNumberConfirmed,
	FieldTwoFactorEnabled,
	FieldLockoutEnabled,
	FieldAccessFailedCount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// NormalizedUsernameValidator is a validator for the "normalized_username" field. It is called by the builders before save.
	NormalizedUsernameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// NormalizedEmailValidator is a validator for the "normalized_email" field. It is called by the builders before save.
	NormalizedEmailValidator func(string) error
	// DefaultEmailConfirmed holds the default value on creation for the "email_confirmed" field.
	DefaultEmailConfirmed bool
	// PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	PasswordHashValidator func(string) error
	// DefaultPhoneNumberConfirmed holds the default value on creation for the "phone_number_confirmed" field.
	DefaultPhoneNumberConfirmed bool
	// DefaultTwoFactorEnabled holds the default value on creation for the "two_factor_enabled" field.
	DefaultTwoFactorEnabled bool
	// DefaultLockoutEnabled holds the default value on creation for the "lockout_enabled" field.
	DefaultLockoutEnabled bool
	// DefaultAccessFailedCount holds the default value on creation for the "access_failed_count" field.
	DefaultAccessFailedCount int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByNormalizedUsername orders the results by the normalized_username field.
func ByNormalizedUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNormalizedUsername, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByNormalizedEmail orders the results by the normalized_email field.
func ByNormalizedEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNormalizedEmail, opts...).ToFunc()
}

// ByEmailConfirmed orders the results by the email_confirmed field.
func ByEmailConfirmed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailConfirmed, opts...).ToFunc()
}

// ByPasswordHash orders the results by the password_hash field.
func ByPasswordHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordHash, opts...).ToFunc()
}

// ByConcurrencyStamp orders the results by the concurrency_stamp field.
func ByConcurrencyStamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConcurrencyStamp, opts...).ToFunc()
}

// BySecurityStamp orders the results by the security_stamp field.
func BySecurityStamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecurityStamp, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phone_number field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

// ByPhoneNumberConfirmed orders the results by the phone_number_confirmed field.
func ByPhoneNumberConfirmed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumberConfirmed, opts...).ToFunc()
}

// ByTwoFactorEnabled orders the results by the two_factor_enabled field.
func ByTwoFactorEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTwoFactorEnabled, opts...).ToFunc()
}

// ByLockoutEnabled orders the results by the lockout_enabled field.
func ByLockoutEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLockoutEnabled, opts...).ToFunc()
}

// ByAccessFailedCount orders the results by the access_failed_count field.
func ByAccessFailedCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccessFailedCount, opts...).ToFunc()
}

// ByUserRolesCount orders the results by user_roles count.
func ByUserRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserRolesStep(), opts...)
	}
}

// ByUserRoles orders the results by user_roles terms.
func ByUserRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserRolesTable, UserRolesColumn),
	)
}
