// Code generated by ent, DO NOT EDIT.

package project

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/charitan-go/project-server/internal/project/dto"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldReadableID holds the string denoting the readable_id field in the database.
	FieldReadableID = "readable_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldGoal holds the string denoting the goal field in the database.
	FieldGoal = "goal"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldCountryCode holds the string denoting the countrycode field in the database.
	FieldCountryCode = "country_code"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldOwnerCharityReadableID holds the string denoting the owner_charity_readable_id field in the database.
	FieldOwnerCharityReadableID = "owner_charity_readable_id"
	// Table holds the table name of the project in the database.
	Table = "projects"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldReadableID,
	FieldName,
	FieldDescription,
	FieldGoal,
	FieldStartDate,
	FieldEndDate,
	FieldCategory,
	FieldCountryCode,
	FieldStatus,
	FieldOwnerCharityReadableID,
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
	// DefaultReadableID holds the default value on creation for the "readable_id" field.
	DefaultReadableID func() uuid.UUID
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// GoalValidator is a validator for the "goal" field. It is called by the builders before save.
	GoalValidator func(float64) error
	// CountryCodeValidator is a validator for the "countryCode" field. It is called by the builders before save.
	CountryCodeValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// CategoryValidator is a validator for the "category" field enum values. It is called by the builders before save.
func CategoryValidator(c dto.CategoryEnum) error {
	switch c {
	case "FOOD", "HEALTH", "EDUCATION", "ENVIRONMENT", "RELIGION", "HOUSING", "OTHER":
		return nil
	default:
		return fmt.Errorf("project: invalid enum value for category field: %q", c)
	}
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s dto.StatusEnum) error {
	switch s {
	case "PENDING", "APPROVED", "DENIED", "HALTED", "FINISHED", "DELETED":
		return nil
	default:
		return fmt.Errorf("project: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Project queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByReadableID orders the results by the readable_id field.
func ByReadableID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReadableID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByGoal orders the results by the goal field.
func ByGoal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGoal, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
}

// ByCountryCode orders the results by the countryCode field.
func ByCountryCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountryCode, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByOwnerCharityReadableID orders the results by the owner_charity_readable_id field.
func ByOwnerCharityReadableID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerCharityReadableID, opts...).ToFunc()
}
