// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/net-auto/resourceManager/ent/predicate"
	"github.com/net-auto/resourceManager/ent/property"
	"github.com/net-auto/resourceManager/ent/resource"
	"github.com/net-auto/resourceManager/ent/resourcepool"
)

// ResourceUpdate is the builder for updating Resource entities.
type ResourceUpdate struct {
	config
	hooks    []Hook
	mutation *ResourceMutation
}

// Where adds a new predicate for the builder.
func (ru *ResourceUpdate) Where(ps ...predicate.Resource) *ResourceUpdate {
	ru.mutation.predicates = append(ru.mutation.predicates, ps...)
	return ru
}

// SetStatus sets the status field.
func (ru *ResourceUpdate) SetStatus(r resource.Status) *ResourceUpdate {
	ru.mutation.SetStatus(r)
	return ru
}

// SetUpdatedAt sets the updated_at field.
func (ru *ResourceUpdate) SetUpdatedAt(t time.Time) *ResourceUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetPoolID sets the pool edge to ResourcePool by id.
func (ru *ResourceUpdate) SetPoolID(id int) *ResourceUpdate {
	ru.mutation.SetPoolID(id)
	return ru
}

// SetNillablePoolID sets the pool edge to ResourcePool by id if the given value is not nil.
func (ru *ResourceUpdate) SetNillablePoolID(id *int) *ResourceUpdate {
	if id != nil {
		ru = ru.SetPoolID(*id)
	}
	return ru
}

// SetPool sets the pool edge to ResourcePool.
func (ru *ResourceUpdate) SetPool(r *ResourcePool) *ResourceUpdate {
	return ru.SetPoolID(r.ID)
}

// AddPropertyIDs adds the properties edge to Property by ids.
func (ru *ResourceUpdate) AddPropertyIDs(ids ...int) *ResourceUpdate {
	ru.mutation.AddPropertyIDs(ids...)
	return ru
}

// AddProperties adds the properties edges to Property.
func (ru *ResourceUpdate) AddProperties(p ...*Property) *ResourceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPropertyIDs(ids...)
}

// SetNestedPoolID sets the nested_pool edge to ResourcePool by id.
func (ru *ResourceUpdate) SetNestedPoolID(id int) *ResourceUpdate {
	ru.mutation.SetNestedPoolID(id)
	return ru
}

// SetNillableNestedPoolID sets the nested_pool edge to ResourcePool by id if the given value is not nil.
func (ru *ResourceUpdate) SetNillableNestedPoolID(id *int) *ResourceUpdate {
	if id != nil {
		ru = ru.SetNestedPoolID(*id)
	}
	return ru
}

// SetNestedPool sets the nested_pool edge to ResourcePool.
func (ru *ResourceUpdate) SetNestedPool(r *ResourcePool) *ResourceUpdate {
	return ru.SetNestedPoolID(r.ID)
}

// Mutation returns the ResourceMutation object of the builder.
func (ru *ResourceUpdate) Mutation() *ResourceMutation {
	return ru.mutation
}

// ClearPool clears the "pool" edge to type ResourcePool.
func (ru *ResourceUpdate) ClearPool() *ResourceUpdate {
	ru.mutation.ClearPool()
	return ru
}

// ClearProperties clears all "properties" edges to type Property.
func (ru *ResourceUpdate) ClearProperties() *ResourceUpdate {
	ru.mutation.ClearProperties()
	return ru
}

// RemovePropertyIDs removes the properties edge to Property by ids.
func (ru *ResourceUpdate) RemovePropertyIDs(ids ...int) *ResourceUpdate {
	ru.mutation.RemovePropertyIDs(ids...)
	return ru
}

// RemoveProperties removes properties edges to Property.
func (ru *ResourceUpdate) RemoveProperties(p ...*Property) *ResourceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePropertyIDs(ids...)
}

// ClearNestedPool clears the "nested_pool" edge to type ResourcePool.
func (ru *ResourceUpdate) ClearNestedPool() *ResourceUpdate {
	ru.mutation.ClearNestedPool()
	return ru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *ResourceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResourceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResourceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ResourceUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := resource.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ResourceUpdate) check() error {
	if v, ok := ru.mutation.Status(); ok {
		if err := resource.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ru *ResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resource.Table,
			Columns: resource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resource.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: resource.FieldStatus,
		})
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: resource.FieldUpdatedAt,
		})
	}
	if ru.mutation.PoolCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PoolIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.PropertiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedPropertiesIDs(); len(nodes) > 0 && !ru.mutation.PropertiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PropertiesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.NestedPoolCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resource.NestedPoolTable,
			Columns: []string{resource.NestedPoolColumn},
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
	if nodes := ru.mutation.NestedPoolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resource.NestedPoolTable,
			Columns: []string{resource.NestedPoolColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ResourceUpdateOne is the builder for updating a single Resource entity.
type ResourceUpdateOne struct {
	config
	hooks    []Hook
	mutation *ResourceMutation
}

// SetStatus sets the status field.
func (ruo *ResourceUpdateOne) SetStatus(r resource.Status) *ResourceUpdateOne {
	ruo.mutation.SetStatus(r)
	return ruo
}

// SetUpdatedAt sets the updated_at field.
func (ruo *ResourceUpdateOne) SetUpdatedAt(t time.Time) *ResourceUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetPoolID sets the pool edge to ResourcePool by id.
func (ruo *ResourceUpdateOne) SetPoolID(id int) *ResourceUpdateOne {
	ruo.mutation.SetPoolID(id)
	return ruo
}

// SetNillablePoolID sets the pool edge to ResourcePool by id if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillablePoolID(id *int) *ResourceUpdateOne {
	if id != nil {
		ruo = ruo.SetPoolID(*id)
	}
	return ruo
}

// SetPool sets the pool edge to ResourcePool.
func (ruo *ResourceUpdateOne) SetPool(r *ResourcePool) *ResourceUpdateOne {
	return ruo.SetPoolID(r.ID)
}

// AddPropertyIDs adds the properties edge to Property by ids.
func (ruo *ResourceUpdateOne) AddPropertyIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.AddPropertyIDs(ids...)
	return ruo
}

// AddProperties adds the properties edges to Property.
func (ruo *ResourceUpdateOne) AddProperties(p ...*Property) *ResourceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPropertyIDs(ids...)
}

// SetNestedPoolID sets the nested_pool edge to ResourcePool by id.
func (ruo *ResourceUpdateOne) SetNestedPoolID(id int) *ResourceUpdateOne {
	ruo.mutation.SetNestedPoolID(id)
	return ruo
}

// SetNillableNestedPoolID sets the nested_pool edge to ResourcePool by id if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillableNestedPoolID(id *int) *ResourceUpdateOne {
	if id != nil {
		ruo = ruo.SetNestedPoolID(*id)
	}
	return ruo
}

// SetNestedPool sets the nested_pool edge to ResourcePool.
func (ruo *ResourceUpdateOne) SetNestedPool(r *ResourcePool) *ResourceUpdateOne {
	return ruo.SetNestedPoolID(r.ID)
}

// Mutation returns the ResourceMutation object of the builder.
func (ruo *ResourceUpdateOne) Mutation() *ResourceMutation {
	return ruo.mutation
}

// ClearPool clears the "pool" edge to type ResourcePool.
func (ruo *ResourceUpdateOne) ClearPool() *ResourceUpdateOne {
	ruo.mutation.ClearPool()
	return ruo
}

// ClearProperties clears all "properties" edges to type Property.
func (ruo *ResourceUpdateOne) ClearProperties() *ResourceUpdateOne {
	ruo.mutation.ClearProperties()
	return ruo
}

// RemovePropertyIDs removes the properties edge to Property by ids.
func (ruo *ResourceUpdateOne) RemovePropertyIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.RemovePropertyIDs(ids...)
	return ruo
}

// RemoveProperties removes properties edges to Property.
func (ruo *ResourceUpdateOne) RemoveProperties(p ...*Property) *ResourceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePropertyIDs(ids...)
}

// ClearNestedPool clears the "nested_pool" edge to type ResourcePool.
func (ruo *ResourceUpdateOne) ClearNestedPool() *ResourceUpdateOne {
	ruo.mutation.ClearNestedPool()
	return ruo
}

// Save executes the query and returns the updated entity.
func (ruo *ResourceUpdateOne) Save(ctx context.Context) (*Resource, error) {
	var (
		err  error
		node *Resource
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResourceUpdateOne) SaveX(ctx context.Context) *Resource {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResourceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ResourceUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := resource.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ResourceUpdateOne) check() error {
	if v, ok := ruo.mutation.Status(); ok {
		if err := resource.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (ruo *ResourceUpdateOne) sqlSave(ctx context.Context) (_node *Resource, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resource.Table,
			Columns: resource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resource.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Resource.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: resource.FieldStatus,
		})
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: resource.FieldUpdatedAt,
		})
	}
	if ruo.mutation.PoolCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PoolIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.PropertiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedPropertiesIDs(); len(nodes) > 0 && !ruo.mutation.PropertiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PropertiesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.NestedPoolCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resource.NestedPoolTable,
			Columns: []string{resource.NestedPoolColumn},
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
	if nodes := ruo.mutation.NestedPoolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   resource.NestedPoolTable,
			Columns: []string{resource.NestedPoolColumn},
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
	_node = &Resource{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
