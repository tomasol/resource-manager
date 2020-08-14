// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/net-auto/resourceManager/ent/allocationstrategy"
	"github.com/net-auto/resourceManager/ent/resourcepool"
	"github.com/net-auto/resourceManager/ent/resourcetype"
)

// ResourcePool is the model entity for the ResourcePool schema.
type ResourcePool struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// PoolType holds the value of the "pool_type" field.
	PoolType resourcepool.PoolType `json:"pool_type,omitempty"`
	// DealocationSafetyPeriod holds the value of the "dealocation_safety_period" field.
	DealocationSafetyPeriod int `json:"dealocation_safety_period,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourcePoolQuery when eager-loading is set.
	Edges                             ResourcePoolEdges `json:"edges"`
	resource_pool_allocation_strategy *int
	resource_type_pools               *int
}

// ResourcePoolEdges holds the relations/edges for other nodes in the graph.
type ResourcePoolEdges struct {
	// ResourceType holds the value of the resource_type edge.
	ResourceType *ResourceType
	// Tags holds the value of the tags edge.
	Tags []*Tag
	// Claims holds the value of the claims edge.
	Claims []*Resource
	// AllocationStrategy holds the value of the allocation_strategy edge.
	AllocationStrategy *AllocationStrategy
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ResourceTypeOrErr returns the ResourceType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourcePoolEdges) ResourceTypeOrErr() (*ResourceType, error) {
	if e.loadedTypes[0] {
		if e.ResourceType == nil {
			// The edge resource_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: resourcetype.Label}
		}
		return e.ResourceType, nil
	}
	return nil, &NotLoadedError{edge: "resource_type"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e ResourcePoolEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[1] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// ClaimsOrErr returns the Claims value or an error if the edge
// was not loaded in eager-loading.
func (e ResourcePoolEdges) ClaimsOrErr() ([]*Resource, error) {
	if e.loadedTypes[2] {
		return e.Claims, nil
	}
	return nil, &NotLoadedError{edge: "claims"}
}

// AllocationStrategyOrErr returns the AllocationStrategy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourcePoolEdges) AllocationStrategyOrErr() (*AllocationStrategy, error) {
	if e.loadedTypes[3] {
		if e.AllocationStrategy == nil {
			// The edge allocation_strategy was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: allocationstrategy.Label}
		}
		return e.AllocationStrategy, nil
	}
	return nil, &NotLoadedError{edge: "allocation_strategy"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ResourcePool) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // pool_type
		&sql.NullInt64{},  // dealocation_safety_period
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*ResourcePool) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // resource_pool_allocation_strategy
		&sql.NullInt64{}, // resource_type_pools
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResourcePool fields.
func (rp *ResourcePool) assignValues(values ...interface{}) error {
	if m, n := len(values), len(resourcepool.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	rp.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		rp.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field pool_type", values[1])
	} else if value.Valid {
		rp.PoolType = resourcepool.PoolType(value.String)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field dealocation_safety_period", values[2])
	} else if value.Valid {
		rp.DealocationSafetyPeriod = int(value.Int64)
	}
	values = values[3:]
	if len(values) == len(resourcepool.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field resource_pool_allocation_strategy", value)
		} else if value.Valid {
			rp.resource_pool_allocation_strategy = new(int)
			*rp.resource_pool_allocation_strategy = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field resource_type_pools", value)
		} else if value.Valid {
			rp.resource_type_pools = new(int)
			*rp.resource_type_pools = int(value.Int64)
		}
	}
	return nil
}

// QueryResourceType queries the resource_type edge of the ResourcePool.
func (rp *ResourcePool) QueryResourceType() *ResourceTypeQuery {
	return (&ResourcePoolClient{config: rp.config}).QueryResourceType(rp)
}

// QueryTags queries the tags edge of the ResourcePool.
func (rp *ResourcePool) QueryTags() *TagQuery {
	return (&ResourcePoolClient{config: rp.config}).QueryTags(rp)
}

// QueryClaims queries the claims edge of the ResourcePool.
func (rp *ResourcePool) QueryClaims() *ResourceQuery {
	return (&ResourcePoolClient{config: rp.config}).QueryClaims(rp)
}

// QueryAllocationStrategy queries the allocation_strategy edge of the ResourcePool.
func (rp *ResourcePool) QueryAllocationStrategy() *AllocationStrategyQuery {
	return (&ResourcePoolClient{config: rp.config}).QueryAllocationStrategy(rp)
}

// Update returns a builder for updating this ResourcePool.
// Note that, you need to call ResourcePool.Unwrap() before calling this method, if this ResourcePool
// was returned from a transaction, and the transaction was committed or rolled back.
func (rp *ResourcePool) Update() *ResourcePoolUpdateOne {
	return (&ResourcePoolClient{config: rp.config}).UpdateOne(rp)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (rp *ResourcePool) Unwrap() *ResourcePool {
	tx, ok := rp.config.driver.(*txDriver)
	if !ok {
		panic("ent: ResourcePool is not a transactional entity")
	}
	rp.config.driver = tx.drv
	return rp
}

// String implements the fmt.Stringer.
func (rp *ResourcePool) String() string {
	var builder strings.Builder
	builder.WriteString("ResourcePool(")
	builder.WriteString(fmt.Sprintf("id=%v", rp.ID))
	builder.WriteString(", name=")
	builder.WriteString(rp.Name)
	builder.WriteString(", pool_type=")
	builder.WriteString(fmt.Sprintf("%v", rp.PoolType))
	builder.WriteString(", dealocation_safety_period=")
	builder.WriteString(fmt.Sprintf("%v", rp.DealocationSafetyPeriod))
	builder.WriteByte(')')
	return builder.String()
}

// ResourcePools is a parsable slice of ResourcePool.
type ResourcePools []*ResourcePool

func (rp ResourcePools) config(cfg config) {
	for _i := range rp {
		rp[_i].config = cfg
	}
}
