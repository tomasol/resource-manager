// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/net-auto/resourceManager/ent/allocationstrategy"
	"github.com/net-auto/resourceManager/ent/poolproperties"
	"github.com/net-auto/resourceManager/ent/resource"
	"github.com/net-auto/resourceManager/ent/resourcepool"
	"github.com/net-auto/resourceManager/ent/resourcetype"
	"github.com/net-auto/resourceManager/ent/tag"
)

// ResourcePoolCreate is the builder for creating a ResourcePool entity.
type ResourcePoolCreate struct {
	config
	mutation *ResourcePoolMutation
	hooks    []Hook
}

// SetName sets the name field.
func (rpc *ResourcePoolCreate) SetName(s string) *ResourcePoolCreate {
	rpc.mutation.SetName(s)
	return rpc
}

// SetDescription sets the description field.
func (rpc *ResourcePoolCreate) SetDescription(s string) *ResourcePoolCreate {
	rpc.mutation.SetDescription(s)
	return rpc
}

// SetNillableDescription sets the description field if the given value is not nil.
func (rpc *ResourcePoolCreate) SetNillableDescription(s *string) *ResourcePoolCreate {
	if s != nil {
		rpc.SetDescription(*s)
	}
	return rpc
}

// SetPoolType sets the pool_type field.
func (rpc *ResourcePoolCreate) SetPoolType(rt resourcepool.PoolType) *ResourcePoolCreate {
	rpc.mutation.SetPoolType(rt)
	return rpc
}

// SetDealocationSafetyPeriod sets the dealocation_safety_period field.
func (rpc *ResourcePoolCreate) SetDealocationSafetyPeriod(i int) *ResourcePoolCreate {
	rpc.mutation.SetDealocationSafetyPeriod(i)
	return rpc
}

// SetNillableDealocationSafetyPeriod sets the dealocation_safety_period field if the given value is not nil.
func (rpc *ResourcePoolCreate) SetNillableDealocationSafetyPeriod(i *int) *ResourcePoolCreate {
	if i != nil {
		rpc.SetDealocationSafetyPeriod(*i)
	}
	return rpc
}

// SetResourceTypeID sets the resource_type edge to ResourceType by id.
func (rpc *ResourcePoolCreate) SetResourceTypeID(id int) *ResourcePoolCreate {
	rpc.mutation.SetResourceTypeID(id)
	return rpc
}

// SetNillableResourceTypeID sets the resource_type edge to ResourceType by id if the given value is not nil.
func (rpc *ResourcePoolCreate) SetNillableResourceTypeID(id *int) *ResourcePoolCreate {
	if id != nil {
		rpc = rpc.SetResourceTypeID(*id)
	}
	return rpc
}

// SetResourceType sets the resource_type edge to ResourceType.
func (rpc *ResourcePoolCreate) SetResourceType(r *ResourceType) *ResourcePoolCreate {
	return rpc.SetResourceTypeID(r.ID)
}

// AddTagIDs adds the tags edge to Tag by ids.
func (rpc *ResourcePoolCreate) AddTagIDs(ids ...int) *ResourcePoolCreate {
	rpc.mutation.AddTagIDs(ids...)
	return rpc
}

// AddTags adds the tags edges to Tag.
func (rpc *ResourcePoolCreate) AddTags(t ...*Tag) *ResourcePoolCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rpc.AddTagIDs(ids...)
}

// AddClaimIDs adds the claims edge to Resource by ids.
func (rpc *ResourcePoolCreate) AddClaimIDs(ids ...int) *ResourcePoolCreate {
	rpc.mutation.AddClaimIDs(ids...)
	return rpc
}

// AddClaims adds the claims edges to Resource.
func (rpc *ResourcePoolCreate) AddClaims(r ...*Resource) *ResourcePoolCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rpc.AddClaimIDs(ids...)
}

// SetPoolPropertiesID sets the poolProperties edge to PoolProperties by id.
func (rpc *ResourcePoolCreate) SetPoolPropertiesID(id int) *ResourcePoolCreate {
	rpc.mutation.SetPoolPropertiesID(id)
	return rpc
}

// SetNillablePoolPropertiesID sets the poolProperties edge to PoolProperties by id if the given value is not nil.
func (rpc *ResourcePoolCreate) SetNillablePoolPropertiesID(id *int) *ResourcePoolCreate {
	if id != nil {
		rpc = rpc.SetPoolPropertiesID(*id)
	}
	return rpc
}

// SetPoolProperties sets the poolProperties edge to PoolProperties.
func (rpc *ResourcePoolCreate) SetPoolProperties(p *PoolProperties) *ResourcePoolCreate {
	return rpc.SetPoolPropertiesID(p.ID)
}

// SetAllocationStrategyID sets the allocation_strategy edge to AllocationStrategy by id.
func (rpc *ResourcePoolCreate) SetAllocationStrategyID(id int) *ResourcePoolCreate {
	rpc.mutation.SetAllocationStrategyID(id)
	return rpc
}

// SetNillableAllocationStrategyID sets the allocation_strategy edge to AllocationStrategy by id if the given value is not nil.
func (rpc *ResourcePoolCreate) SetNillableAllocationStrategyID(id *int) *ResourcePoolCreate {
	if id != nil {
		rpc = rpc.SetAllocationStrategyID(*id)
	}
	return rpc
}

// SetAllocationStrategy sets the allocation_strategy edge to AllocationStrategy.
func (rpc *ResourcePoolCreate) SetAllocationStrategy(a *AllocationStrategy) *ResourcePoolCreate {
	return rpc.SetAllocationStrategyID(a.ID)
}

// SetParentResourceID sets the parent_resource edge to Resource by id.
func (rpc *ResourcePoolCreate) SetParentResourceID(id int) *ResourcePoolCreate {
	rpc.mutation.SetParentResourceID(id)
	return rpc
}

// SetNillableParentResourceID sets the parent_resource edge to Resource by id if the given value is not nil.
func (rpc *ResourcePoolCreate) SetNillableParentResourceID(id *int) *ResourcePoolCreate {
	if id != nil {
		rpc = rpc.SetParentResourceID(*id)
	}
	return rpc
}

// SetParentResource sets the parent_resource edge to Resource.
func (rpc *ResourcePoolCreate) SetParentResource(r *Resource) *ResourcePoolCreate {
	return rpc.SetParentResourceID(r.ID)
}

// Mutation returns the ResourcePoolMutation object of the builder.
func (rpc *ResourcePoolCreate) Mutation() *ResourcePoolMutation {
	return rpc.mutation
}

// Save creates the ResourcePool in the database.
func (rpc *ResourcePoolCreate) Save(ctx context.Context) (*ResourcePool, error) {
	var (
		err  error
		node *ResourcePool
	)
	rpc.defaults()
	if len(rpc.hooks) == 0 {
		if err = rpc.check(); err != nil {
			return nil, err
		}
		node, err = rpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResourcePoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rpc.check(); err != nil {
				return nil, err
			}
			rpc.mutation = mutation
			node, err = rpc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rpc.hooks) - 1; i >= 0; i-- {
			mut = rpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rpc *ResourcePoolCreate) SaveX(ctx context.Context) *ResourcePool {
	v, err := rpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (rpc *ResourcePoolCreate) defaults() {
	if _, ok := rpc.mutation.DealocationSafetyPeriod(); !ok {
		v := resourcepool.DefaultDealocationSafetyPeriod
		rpc.mutation.SetDealocationSafetyPeriod(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpc *ResourcePoolCreate) check() error {
	if _, ok := rpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := rpc.mutation.Name(); ok {
		if err := resourcepool.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := rpc.mutation.PoolType(); !ok {
		return &ValidationError{Name: "pool_type", err: errors.New("ent: missing required field \"pool_type\"")}
	}
	if v, ok := rpc.mutation.PoolType(); ok {
		if err := resourcepool.PoolTypeValidator(v); err != nil {
			return &ValidationError{Name: "pool_type", err: fmt.Errorf("ent: validator failed for field \"pool_type\": %w", err)}
		}
	}
	if _, ok := rpc.mutation.DealocationSafetyPeriod(); !ok {
		return &ValidationError{Name: "dealocation_safety_period", err: errors.New("ent: missing required field \"dealocation_safety_period\"")}
	}
	return nil
}

func (rpc *ResourcePoolCreate) sqlSave(ctx context.Context) (*ResourcePool, error) {
	_node, _spec := rpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rpc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rpc *ResourcePoolCreate) createSpec() (*ResourcePool, *sqlgraph.CreateSpec) {
	var (
		_node = &ResourcePool{config: rpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: resourcepool.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resourcepool.FieldID,
			},
		}
	)
	if value, ok := rpc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resourcepool.FieldName,
		})
		_node.Name = value
	}
	if value, ok := rpc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resourcepool.FieldDescription,
		})
		_node.Description = &value
	}
	if value, ok := rpc.mutation.PoolType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: resourcepool.FieldPoolType,
		})
		_node.PoolType = value
	}
	if value, ok := rpc.mutation.DealocationSafetyPeriod(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resourcepool.FieldDealocationSafetyPeriod,
		})
		_node.DealocationSafetyPeriod = value
	}
	if nodes := rpc.mutation.ResourceTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   resourcepool.ResourceTypeTable,
			Columns: []string{resourcepool.ResourceTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resourcetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rpc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   resourcepool.TagsTable,
			Columns: resourcepool.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rpc.mutation.ClaimsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resourcepool.ClaimsTable,
			Columns: []string{resourcepool.ClaimsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resource.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rpc.mutation.PoolPropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resourcepool.PoolPropertiesTable,
			Columns: []string{resourcepool.PoolPropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: poolproperties.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rpc.mutation.AllocationStrategyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   resourcepool.AllocationStrategyTable,
			Columns: []string{resourcepool.AllocationStrategyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: allocationstrategy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rpc.mutation.ParentResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   resourcepool.ParentResourceTable,
			Columns: []string{resourcepool.ParentResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resource.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ResourcePoolCreateBulk is the builder for creating a bulk of ResourcePool entities.
type ResourcePoolCreateBulk struct {
	config
	builders []*ResourcePoolCreate
}

// Save creates the ResourcePool entities in the database.
func (rpcb *ResourcePoolCreateBulk) Save(ctx context.Context) ([]*ResourcePool, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rpcb.builders))
	nodes := make([]*ResourcePool, len(rpcb.builders))
	mutators := make([]Mutator, len(rpcb.builders))
	for i := range rpcb.builders {
		func(i int, root context.Context) {
			builder := rpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResourcePoolMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rpcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rpcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (rpcb *ResourcePoolCreateBulk) SaveX(ctx context.Context) []*ResourcePool {
	v, err := rpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
