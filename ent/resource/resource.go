// Code generated by entc, DO NOT EDIT.

package resource

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/facebook/ent"
)

const (
	// Label holds the string label denoting the resource type in the database.
	Label = "resource"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"

	// EdgePool holds the string denoting the pool edge name in mutations.
	EdgePool = "pool"
	// EdgeProperties holds the string denoting the properties edge name in mutations.
	EdgeProperties = "properties"
	// EdgeNestedPool holds the string denoting the nested_pool edge name in mutations.
	EdgeNestedPool = "nested_pool"

	// Table holds the table name of the resource in the database.
	Table = "resources"
	// PoolTable is the table the holds the pool relation/edge.
	PoolTable = "resources"
	// PoolInverseTable is the table name for the ResourcePool entity.
	// It exists in this package in order to avoid circular dependency with the "resourcepool" package.
	PoolInverseTable = "resource_pools"
	// PoolColumn is the table column denoting the pool relation/edge.
	PoolColumn = "resource_pool_claims"
	// PropertiesTable is the table the holds the properties relation/edge.
	PropertiesTable = "properties"
	// PropertiesInverseTable is the table name for the Property entity.
	// It exists in this package in order to avoid circular dependency with the "property" package.
	PropertiesInverseTable = "properties"
	// PropertiesColumn is the table column denoting the properties relation/edge.
	PropertiesColumn = "resource_properties"
	// NestedPoolTable is the table the holds the nested_pool relation/edge.
	NestedPoolTable = "resource_pools"
	// NestedPoolInverseTable is the table name for the ResourcePool entity.
	// It exists in this package in order to avoid circular dependency with the "resourcepool" package.
	NestedPoolInverseTable = "resource_pools"
	// NestedPoolColumn is the table column denoting the nested_pool relation/edge.
	NestedPoolColumn = "resource_nested_pool"
)

// Columns holds all SQL columns for resource fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Resource type.
var ForeignKeys = []string{
	"resource_pool_claims",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/net-auto/resourceManager/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the status enum field.
type Status string

// Status values.
const (
	StatusFree    Status = "free"
	StatusClaimed Status = "claimed"
	StatusRetired Status = "retired"
	StatusBench   Status = "bench"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusFree, StatusClaimed, StatusRetired, StatusBench:
		return nil
	default:
		return fmt.Errorf("resource: invalid enum value for status field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (s Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(s.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (s *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*s = Status(str)
	if err := StatusValidator(*s); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}
