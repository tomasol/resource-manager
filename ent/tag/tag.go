// Code generated by entc, DO NOT EDIT.

package tag

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"

	// EdgePools holds the string denoting the pools edge name in mutations.
	EdgePools = "pools"

	// Table holds the table name of the tag in the database.
	Table = "tags"
	// PoolsTable is the table the holds the pools relation/edge. The primary key declared below.
	PoolsTable = "tag_pools"
	// PoolsInverseTable is the table name for the ResourcePool entity.
	// It exists in this package in order to avoid circular dependency with the "resourcepool" package.
	PoolsInverseTable = "resource_pools"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldTag,
}

var (
	// PoolsPrimaryKey and PoolsColumn2 are the table columns denoting the
	// primary key for the pools relation (M2M).
	PoolsPrimaryKey = []string{"tag_id", "resource_pool_id"}
)

var (
	// TagValidator is a validator for the "tag" field. It is called by the builders before save.
	TagValidator func(string) error
)