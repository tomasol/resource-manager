// Code generated by entc, DO NOT EDIT.

package propertytype

import (
	"fmt"
	"io"
	"strconv"
)

const (
	// Label holds the string label denoting the propertytype type in the database.
	Label = "property_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldExternalID holds the string denoting the external_id field in the database.
	FieldExternalID = "external_id"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldIntVal holds the string denoting the int_val field in the database.
	FieldIntVal = "int_val"
	// FieldBoolVal holds the string denoting the bool_val field in the database.
	FieldBoolVal = "bool_val"
	// FieldFloatVal holds the string denoting the float_val field in the database.
	FieldFloatVal = "float_val"
	// FieldLatitudeVal holds the string denoting the latitude_val field in the database.
	FieldLatitudeVal = "latitude_val"
	// FieldLongitudeVal holds the string denoting the longitude_val field in the database.
	FieldLongitudeVal = "longitude_val"
	// FieldStringVal holds the string denoting the string_val field in the database.
	FieldStringVal = "string_val"
	// FieldRangeFromVal holds the string denoting the range_from_val field in the database.
	FieldRangeFromVal = "range_from_val"
	// FieldRangeToVal holds the string denoting the range_to_val field in the database.
	FieldRangeToVal = "range_to_val"
	// FieldIsInstanceProperty holds the string denoting the is_instance_property field in the database.
	FieldIsInstanceProperty = "is_instance_property"
	// FieldEditable holds the string denoting the editable field in the database.
	FieldEditable = "editable"
	// FieldMandatory holds the string denoting the mandatory field in the database.
	FieldMandatory = "mandatory"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// FieldNodeType holds the string denoting the nodetype field in the database.
	FieldNodeType = "node_type"

	// EdgeProperties holds the string denoting the properties edge name in mutations.
	EdgeProperties = "properties"

	// Table holds the table name of the propertytype in the database.
	Table = "property_types"
	// PropertiesTable is the table the holds the properties relation/edge.
	PropertiesTable = "properties"
	// PropertiesInverseTable is the table name for the Property entity.
	// It exists in this package in order to avoid circular dependency with the "property" package.
	PropertiesInverseTable = "properties"
	// PropertiesColumn is the table column denoting the properties relation/edge.
	PropertiesColumn = "property_type"
)

// Columns holds all SQL columns for propertytype fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldName,
	FieldExternalID,
	FieldIndex,
	FieldCategory,
	FieldIntVal,
	FieldBoolVal,
	FieldFloatVal,
	FieldLatitudeVal,
	FieldLongitudeVal,
	FieldStringVal,
	FieldRangeFromVal,
	FieldRangeToVal,
	FieldIsInstanceProperty,
	FieldEditable,
	FieldMandatory,
	FieldDeleted,
	FieldNodeType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the PropertyType type.
var ForeignKeys = []string{
	"resource_type_property_types",
}

var (
	// DefaultIsInstanceProperty holds the default value on creation for the is_instance_property field.
	DefaultIsInstanceProperty bool
	// DefaultEditable holds the default value on creation for the editable field.
	DefaultEditable bool
	// DefaultMandatory holds the default value on creation for the mandatory field.
	DefaultMandatory bool
	// DefaultDeleted holds the default value on creation for the deleted field.
	DefaultDeleted bool
)

// Type defines the type for the type enum field.
type Type string

// Type values.
const (
	TypeBool          Type = "bool"
	TypeDate          Type = "date"
	TypeDatetimeLocal Type = "datetime_local"
	TypeEmail         Type = "email"
	TypeEnum          Type = "enum"
	TypeFloat         Type = "float"
	TypeGpsLocation   Type = "gps_location"
	TypeInt           Type = "int"
	TypeNode          Type = "node"
	TypeRange         Type = "range"
	TypeString        Type = "string"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeBool, TypeDate, TypeDatetimeLocal, TypeEmail, TypeEnum, TypeFloat, TypeGpsLocation, TypeInt, TypeNode, TypeRange, TypeString:
		return nil
	default:
		return fmt.Errorf("propertytype: invalid enum value for type field: %q", _type)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (_type Type) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(_type.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (_type *Type) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", v)
	}
	*_type = Type(str)
	if err := TypeValidator(*_type); err != nil {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}
