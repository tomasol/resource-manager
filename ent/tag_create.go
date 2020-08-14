// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/net-auto/resourceManager/ent/resourcepool"
	"github.com/net-auto/resourceManager/ent/tag"
)

// TagCreate is the builder for creating a Tag entity.
type TagCreate struct {
	config
	mutation *TagMutation
	hooks    []Hook
}

// SetTag sets the tag field.
func (tc *TagCreate) SetTag(s string) *TagCreate {
	tc.mutation.SetTag(s)
	return tc
}

// AddPoolIDs adds the pools edge to ResourcePool by ids.
func (tc *TagCreate) AddPoolIDs(ids ...int) *TagCreate {
	tc.mutation.AddPoolIDs(ids...)
	return tc
}

// AddPools adds the pools edges to ResourcePool.
func (tc *TagCreate) AddPools(r ...*ResourcePool) *TagCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tc.AddPoolIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tc *TagCreate) Mutation() *TagMutation {
	return tc.mutation
}

// Save creates the Tag in the database.
func (tc *TagCreate) Save(ctx context.Context) (*Tag, error) {
	if err := tc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Tag
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TagCreate) SaveX(ctx context.Context) *Tag {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TagCreate) preSave() error {
	if _, ok := tc.mutation.Tag(); !ok {
		return &ValidationError{Name: "tag", err: errors.New("ent: missing required field \"tag\"")}
	}
	if v, ok := tc.mutation.Tag(); ok {
		if err := tag.TagValidator(v); err != nil {
			return &ValidationError{Name: "tag", err: fmt.Errorf("ent: validator failed for field \"tag\": %w", err)}
		}
	}
	return nil
}

func (tc *TagCreate) sqlSave(ctx context.Context) (*Tag, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TagCreate) createSpec() (*Tag, *sqlgraph.CreateSpec) {
	var (
		t     = &Tag{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tag.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Tag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldTag,
		})
		t.Tag = value
	}
	if nodes := tc.mutation.PoolsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.PoolsTable,
			Columns: tag.PoolsPrimaryKey,
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
	return t, _spec
}

// TagCreateBulk is the builder for creating a bulk of Tag entities.
type TagCreateBulk struct {
	config
	builders []*TagCreate
}

// Save creates the Tag entities in the database.
func (tcb *TagCreateBulk) Save(ctx context.Context) ([]*Tag, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tag, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				builder := tcb.builders[i]
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*TagMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
			for i := len(tcb.builders[i].hooks) - 1; i >= 0; i-- {
				mut = tcb.builders[i].hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
		return nil, err
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (tcb *TagCreateBulk) SaveX(ctx context.Context) []*Tag {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
