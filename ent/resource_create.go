// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/net-auto/resourceManager/ent/property"
	"github.com/net-auto/resourceManager/ent/resource"
	"github.com/net-auto/resourceManager/ent/resourcepool"
)

// ResourceCreate is the builder for creating a Resource entity.
type ResourceCreate struct {
	config
	mutation *ResourceMutation
	hooks    []Hook
}

// SetClaimed sets the claimed field.
func (rc *ResourceCreate) SetClaimed(b bool) *ResourceCreate {
	rc.mutation.SetClaimed(b)
	return rc
}

// SetNillableClaimed sets the claimed field if the given value is not nil.
func (rc *ResourceCreate) SetNillableClaimed(b *bool) *ResourceCreate {
	if b != nil {
		rc.SetClaimed(*b)
	}
	return rc
}

// SetPoolID sets the pool edge to ResourcePool by id.
func (rc *ResourceCreate) SetPoolID(id int) *ResourceCreate {
	rc.mutation.SetPoolID(id)
	return rc
}

// SetNillablePoolID sets the pool edge to ResourcePool by id if the given value is not nil.
func (rc *ResourceCreate) SetNillablePoolID(id *int) *ResourceCreate {
	if id != nil {
		rc = rc.SetPoolID(*id)
	}
	return rc
}

// SetPool sets the pool edge to ResourcePool.
func (rc *ResourceCreate) SetPool(r *ResourcePool) *ResourceCreate {
	return rc.SetPoolID(r.ID)
}

// AddPropertyIDs adds the properties edge to Property by ids.
func (rc *ResourceCreate) AddPropertyIDs(ids ...int) *ResourceCreate {
	rc.mutation.AddPropertyIDs(ids...)
	return rc
}

// AddProperties adds the properties edges to Property.
func (rc *ResourceCreate) AddProperties(p ...*Property) *ResourceCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPropertyIDs(ids...)
}

// Mutation returns the ResourceMutation object of the builder.
func (rc *ResourceCreate) Mutation() *ResourceMutation {
	return rc.mutation
}

// Save creates the Resource in the database.
func (rc *ResourceCreate) Save(ctx context.Context) (*Resource, error) {
	if err := rc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Resource
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ResourceCreate) SaveX(ctx context.Context) *Resource {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *ResourceCreate) preSave() error {
	if _, ok := rc.mutation.Claimed(); !ok {
		v := resource.DefaultClaimed
		rc.mutation.SetClaimed(v)
	}
	return nil
}

func (rc *ResourceCreate) sqlSave(ctx context.Context) (*Resource, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	r.ID = int(id)
	return r, nil
}

func (rc *ResourceCreate) createSpec() (*Resource, *sqlgraph.CreateSpec) {
	var (
		r     = &Resource{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: resource.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resource.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Claimed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: resource.FieldClaimed,
		})
		r.Claimed = value
	}
	if nodes := rc.mutation.PoolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   resource.PoolTable,
			Columns: []string{resource.PoolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resourcepool.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.PropertiesTable,
			Columns: []string{resource.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: property.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return r, _spec
}

// ResourceCreateBulk is the builder for creating a bulk of Resource entities.
type ResourceCreateBulk struct {
	config
	builders []*ResourceCreate
}

// Save creates the Resource entities in the database.
func (rcb *ResourceCreateBulk) Save(ctx context.Context) ([]*Resource, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Resource, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				builder := rcb.builders[i]
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*ResourceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
			for i := len(rcb.builders[i].hooks) - 1; i >= 0; i-- {
				mut = rcb.builders[i].hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
		return nil, err
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (rcb *ResourceCreateBulk) SaveX(ctx context.Context) []*Resource {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
