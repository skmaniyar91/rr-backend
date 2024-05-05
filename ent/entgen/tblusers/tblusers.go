// Code generated by ent, DO NOT EDIT.

package tblusers

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the tblusers type in the database.
	Label = "tbl_users"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "Id_ulid"
	// FieldCreatedBy holds the string denoting the createdby field in the database.
	FieldCreatedBy = "CreatedBy"
	// FieldUpdatedBy holds the string denoting the updatedby field in the database.
	FieldUpdatedBy = "UpdatedBy"
	// FieldDeletedBy holds the string denoting the deletedby field in the database.
	FieldDeletedBy = "DeletedBy"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "IP"
	// FieldUserAgent holds the string denoting the useragent field in the database.
	FieldUserAgent = "UserAgent"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "CreatedAt"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "UpdatedAt"
	// FieldDeletedAt holds the string denoting the deletedat field in the database.
	FieldDeletedAt = "DeletedAt"
	// FieldUserName holds the string denoting the username field in the database.
	FieldUserName = "UserName"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "Password"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "Email"
	// Table holds the table name of the tblusers in the database.
	Table = "Tbl_Users"
)

// Columns holds all SQL columns for tblusers fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedBy,
	FieldIP,
	FieldUserAgent,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUserName,
	FieldPassword,
	FieldEmail,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "rr-backend/ent/entgen/runtime"
var (
	Hooks        [4]ent.Hook
	Interceptors [1]ent.Interceptor
	// CreatedByValidator is a validator for the "CreatedBy" field. It is called by the builders before save.
	CreatedByValidator func(string) error
	// UpdatedByValidator is a validator for the "UpdatedBy" field. It is called by the builders before save.
	UpdatedByValidator func(string) error
	// DeletedByValidator is a validator for the "DeletedBy" field. It is called by the builders before save.
	DeletedByValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "CreatedAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "UpdatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "UpdatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
	// UserNameValidator is a validator for the "UserName" field. It is called by the builders before save.
	UserNameValidator func(string) error
	// PasswordValidator is a validator for the "Password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// EmailValidator is a validator for the "Email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the TblUSers queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the CreatedBy field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the UpdatedBy field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDeletedBy orders the results by the DeletedBy field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByIP orders the results by the IP field.
func ByIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIP, opts...).ToFunc()
}

// ByUserAgent orders the results by the UserAgent field.
func ByUserAgent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserAgent, opts...).ToFunc()
}

// ByCreatedAt orders the results by the CreatedAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the UpdatedAt field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the DeletedAt field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByUserName orders the results by the UserName field.
func ByUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserName, opts...).ToFunc()
}

// ByPassword orders the results by the Password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByEmail orders the results by the Email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}