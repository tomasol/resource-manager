// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/net-auto/resourceManager/ent/poolproperties"
	"github.com/net-auto/resourceManager/ent/predicate"
	"github.com/net-auto/resourceManager/ent/property"
	"github.com/net-auto/resourceManager/ent/resourcepool"
	"github.com/net-auto/resourceManager/ent/resourcetype"
)

// PoolPropertiesUpdate is the builder for updating PoolProperties entities.
type PoolPropertiesUpdate struct {
	config
	hooks    []Hook
	mutation *PoolPropertiesMutation
}

// Where adds a new predicate for the builder.
func (ppu *PoolPropertiesUpdate) Where(ps ...predicate.PoolProperties) *PoolPropertiesUpdate {
	ppu.mutation.predicates = append(ppu.mutation.predicates, ps...)
	return ppu
}

// SetPoolID sets the pool edge to ResourcePool by id.
func (ppu *PoolPropertiesUpdate) SetPoolID(id int) *PoolPropertiesUpdate {
	ppu.mutation.SetPoolID(id)
	return ppu
}

// SetNillablePoolID sets the pool edge to ResourcePool by id if the given value is not nil.
func (ppu *PoolPropertiesUpdate) SetNillablePoolID(id *int) *PoolPropertiesUpdate {
	if id != nil {
		ppu = ppu.SetPoolID(*id)
	}
	return ppu
}

// SetPool sets the pool edge to ResourcePool.
func (ppu *PoolPropertiesUpdate) SetPool(r *ResourcePool) *PoolPropertiesUpdate {
	return ppu.SetPoolID(r.ID)
}

// AddResourceTypeIDs adds the resourceType edge to ResourceType by ids.
func (ppu *PoolPropertiesUpdate) AddResourceTypeIDs(ids ...int) *PoolPropertiesUpdate {
	ppu.mutation.AddResourceTypeIDs(ids...)
	return ppu
}

// AddResourceType adds the resourceType edges to ResourceType.
func (ppu *PoolPropertiesUpdate) AddResourceType(r ...*ResourceType) *PoolPropertiesUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ppu.AddResourceTypeIDs(ids...)
}

// AddPropertyIDs adds the properties edge to Property by ids.
func (ppu *PoolPropertiesUpdate) AddPropertyIDs(ids ...int) *PoolPropertiesUpdate {
	ppu.mutation.AddPropertyIDs(ids...)
	return ppu
}

// AddProperties adds the properties edges to Property.
func (ppu *PoolPropertiesUpdate) AddProperties(p ...*Property) *PoolPropertiesUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppu.AddPropertyIDs(ids...)
}

// Mutation returns the PoolPropertiesMutation object of the builder.
func (ppu *PoolPropertiesUpdate) Mutation() *PoolPropertiesMutation {
	return ppu.mutation
}

// ClearPool clears the "pool" edge to type ResourcePool.
func (ppu *PoolPropertiesUpdate) ClearPool() *PoolPropertiesUpdate {
	ppu.mutation.ClearPool()
	return ppu
}

// ClearResourceType clears all "resourceType" edges to type ResourceType.
func (ppu *PoolPropertiesUpdate) ClearResourceType() *PoolPropertiesUpdate {
	ppu.mutation.ClearResourceType()
	return ppu
}

// RemoveResourceTypeIDs removes the resourceType edge to ResourceType by ids.
func (ppu *PoolPropertiesUpdate) RemoveResourceTypeIDs(ids ...int) *PoolPropertiesUpdate {
	ppu.mutation.RemoveResourceTypeIDs(ids...)
	return ppu
}

// RemoveResourceType removes resourceType edges to ResourceType.
func (ppu *PoolPropertiesUpdate) RemoveResourceType(r ...*ResourceType) *PoolPropertiesUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ppu.RemoveResourceTypeIDs(ids...)
}

// ClearProperties clears all "properties" edges to type Property.
func (ppu *PoolPropertiesUpdate) ClearProperties() *PoolPropertiesUpdate {
	ppu.mutation.ClearProperties()
	return ppu
}

// RemovePropertyIDs removes the properties edge to Property by ids.
func (ppu *PoolPropertiesUpdate) RemovePropertyIDs(ids ...int) *PoolPropertiesUpdate {
	ppu.mutation.RemovePropertyIDs(ids...)
	return ppu
}

// RemoveProperties removes properties edges to Property.
func (ppu *PoolPropertiesUpdate) RemoveProperties(p ...*Property) *PoolPropertiesUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppu.RemovePropertyIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ppu *PoolPropertiesUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ppu.hooks) == 0 {
		affected, err = ppu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PoolPropertiesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ppu.mutation = mutation
			affected, err = ppu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ppu.hooks) - 1; i >= 0; i-- {
			mut = ppu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ppu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ppu *PoolPropertiesUpdate) SaveX(ctx context.Context) int {
	affected, err := ppu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ppu *PoolPropertiesUpdate) Exec(ctx context.Context) error {
	_, err := ppu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppu *PoolPropertiesUpdate) ExecX(ctx context.Context) {
	if err := ppu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ppu *PoolPropertiesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   poolproperties.Table,
			Columns: poolproperties.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: poolproperties.FieldID,
			},
		},
	}
	if ps := ppu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ppu.mutation.PoolCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   poolproperties.PoolTable,
			Columns: []string{poolproperties.PoolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resourcepool.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.PoolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   poolproperties.PoolTable,
			Columns: []string{poolproperties.PoolColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ppu.mutation.ResourceTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poolproperties.ResourceTypeTable,
			Columns: poolproperties.ResourceTypePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resourcetype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.RemovedResourceTypeIDs(); len(nodes) > 0 && !ppu.mutation.ResourceTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poolproperties.ResourceTypeTable,
			Columns: poolproperties.ResourceTypePrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.ResourceTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poolproperties.ResourceTypeTable,
			Columns: poolproperties.ResourceTypePrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ppu.mutation.PropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poolproperties.PropertiesTable,
			Columns: []string{poolproperties.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: property.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.RemovedPropertiesIDs(); len(nodes) > 0 && !ppu.mutation.PropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poolproperties.PropertiesTable,
			Columns: []string{poolproperties.PropertiesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.PropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poolproperties.PropertiesTable,
			Columns: []string{poolproperties.PropertiesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ppu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{poolproperties.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PoolPropertiesUpdateOne is the builder for updating a single PoolProperties entity.
type PoolPropertiesUpdateOne struct {
	config
	hooks    []Hook
	mutation *PoolPropertiesMutation
}

// SetPoolID sets the pool edge to ResourcePool by id.
func (ppuo *PoolPropertiesUpdateOne) SetPoolID(id int) *PoolPropertiesUpdateOne {
	ppuo.mutation.SetPoolID(id)
	return ppuo
}

// SetNillablePoolID sets the pool edge to ResourcePool by id if the given value is not nil.
func (ppuo *PoolPropertiesUpdateOne) SetNillablePoolID(id *int) *PoolPropertiesUpdateOne {
	if id != nil {
		ppuo = ppuo.SetPoolID(*id)
	}
	return ppuo
}

// SetPool sets the pool edge to ResourcePool.
func (ppuo *PoolPropertiesUpdateOne) SetPool(r *ResourcePool) *PoolPropertiesUpdateOne {
	return ppuo.SetPoolID(r.ID)
}

// AddResourceTypeIDs adds the resourceType edge to ResourceType by ids.
func (ppuo *PoolPropertiesUpdateOne) AddResourceTypeIDs(ids ...int) *PoolPropertiesUpdateOne {
	ppuo.mutation.AddResourceTypeIDs(ids...)
	return ppuo
}

// AddResourceType adds the resourceType edges to ResourceType.
func (ppuo *PoolPropertiesUpdateOne) AddResourceType(r ...*ResourceType) *PoolPropertiesUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ppuo.AddResourceTypeIDs(ids...)
}

// AddPropertyIDs adds the properties edge to Property by ids.
func (ppuo *PoolPropertiesUpdateOne) AddPropertyIDs(ids ...int) *PoolPropertiesUpdateOne {
	ppuo.mutation.AddPropertyIDs(ids...)
	return ppuo
}

// AddProperties adds the properties edges to Property.
func (ppuo *PoolPropertiesUpdateOne) AddProperties(p ...*Property) *PoolPropertiesUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppuo.AddPropertyIDs(ids...)
}

// Mutation returns the PoolPropertiesMutation object of the builder.
func (ppuo *PoolPropertiesUpdateOne) Mutation() *PoolPropertiesMutation {
	return ppuo.mutation
}

// ClearPool clears the "pool" edge to type ResourcePool.
func (ppuo *PoolPropertiesUpdateOne) ClearPool() *PoolPropertiesUpdateOne {
	ppuo.mutation.ClearPool()
	return ppuo
}

// ClearResourceType clears all "resourceType" edges to type ResourceType.
func (ppuo *PoolPropertiesUpdateOne) ClearResourceType() *PoolPropertiesUpdateOne {
	ppuo.mutation.ClearResourceType()
	return ppuo
}

// RemoveResourceTypeIDs removes the resourceType edge to ResourceType by ids.
func (ppuo *PoolPropertiesUpdateOne) RemoveResourceTypeIDs(ids ...int) *PoolPropertiesUpdateOne {
	ppuo.mutation.RemoveResourceTypeIDs(ids...)
	return ppuo
}

// RemoveResourceType removes resourceType edges to ResourceType.
func (ppuo *PoolPropertiesUpdateOne) RemoveResourceType(r ...*ResourceType) *PoolPropertiesUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ppuo.RemoveResourceTypeIDs(ids...)
}

// ClearProperties clears all "properties" edges to type Property.
func (ppuo *PoolPropertiesUpdateOne) ClearProperties() *PoolPropertiesUpdateOne {
	ppuo.mutation.ClearProperties()
	return ppuo
}

// RemovePropertyIDs removes the properties edge to Property by ids.
func (ppuo *PoolPropertiesUpdateOne) RemovePropertyIDs(ids ...int) *PoolPropertiesUpdateOne {
	ppuo.mutation.RemovePropertyIDs(ids...)
	return ppuo
}

// RemoveProperties removes properties edges to Property.
func (ppuo *PoolPropertiesUpdateOne) RemoveProperties(p ...*Property) *PoolPropertiesUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppuo.RemovePropertyIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ppuo *PoolPropertiesUpdateOne) Save(ctx context.Context) (*PoolProperties, error) {
	var (
		err  error
		node *PoolProperties
	)
	if len(ppuo.hooks) == 0 {
		node, err = ppuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PoolPropertiesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ppuo.mutation = mutation
			node, err = ppuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ppuo.hooks) - 1; i >= 0; i-- {
			mut = ppuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ppuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ppuo *PoolPropertiesUpdateOne) SaveX(ctx context.Context) *PoolProperties {
	node, err := ppuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ppuo *PoolPropertiesUpdateOne) Exec(ctx context.Context) error {
	_, err := ppuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppuo *PoolPropertiesUpdateOne) ExecX(ctx context.Context) {
	if err := ppuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ppuo *PoolPropertiesUpdateOne) sqlSave(ctx context.Context) (_node *PoolProperties, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   poolproperties.Table,
			Columns: poolproperties.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: poolproperties.FieldID,
			},
		},
	}
	id, ok := ppuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PoolProperties.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ppuo.mutation.PoolCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   poolproperties.PoolTable,
			Columns: []string{poolproperties.PoolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resourcepool.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.PoolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   poolproperties.PoolTable,
			Columns: []string{poolproperties.PoolColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ppuo.mutation.ResourceTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poolproperties.ResourceTypeTable,
			Columns: poolproperties.ResourceTypePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resourcetype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.RemovedResourceTypeIDs(); len(nodes) > 0 && !ppuo.mutation.ResourceTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poolproperties.ResourceTypeTable,
			Columns: poolproperties.ResourceTypePrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.ResourceTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poolproperties.ResourceTypeTable,
			Columns: poolproperties.ResourceTypePrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ppuo.mutation.PropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poolproperties.PropertiesTable,
			Columns: []string{poolproperties.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: property.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.RemovedPropertiesIDs(); len(nodes) > 0 && !ppuo.mutation.PropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poolproperties.PropertiesTable,
			Columns: []string{poolproperties.PropertiesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.PropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poolproperties.PropertiesTable,
			Columns: []string{poolproperties.PropertiesColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PoolProperties{config: ppuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ppuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{poolproperties.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
