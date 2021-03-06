// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/net-auto/resourceManager/ent/predicate"
	"github.com/net-auto/resourceManager/ent/property"
	"github.com/net-auto/resourceManager/ent/resource"
	"github.com/net-auto/resourceManager/ent/resourcepool"
)

// ResourceQuery is the builder for querying Resource entities.
type ResourceQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Resource
	// eager-loading edges.
	withPool       *ResourcePoolQuery
	withProperties *PropertyQuery
	withNestedPool *ResourcePoolQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (rq *ResourceQuery) Where(ps ...predicate.Resource) *ResourceQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *ResourceQuery) Limit(limit int) *ResourceQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *ResourceQuery) Offset(offset int) *ResourceQuery {
	rq.offset = &offset
	return rq
}

// Order adds an order step to the query.
func (rq *ResourceQuery) Order(o ...OrderFunc) *ResourceQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryPool chains the current query on the pool edge.
func (rq *ResourceQuery) QueryPool() *ResourcePoolQuery {
	query := &ResourcePoolQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resource.Table, resource.FieldID, selector),
			sqlgraph.To(resourcepool.Table, resourcepool.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, resource.PoolTable, resource.PoolColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProperties chains the current query on the properties edge.
func (rq *ResourceQuery) QueryProperties() *PropertyQuery {
	query := &PropertyQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resource.Table, resource.FieldID, selector),
			sqlgraph.To(property.Table, property.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, resource.PropertiesTable, resource.PropertiesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNestedPool chains the current query on the nested_pool edge.
func (rq *ResourceQuery) QueryNestedPool() *ResourcePoolQuery {
	query := &ResourcePoolQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resource.Table, resource.FieldID, selector),
			sqlgraph.To(resourcepool.Table, resourcepool.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, resource.NestedPoolTable, resource.NestedPoolColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Resource entity in the query. Returns *NotFoundError when no resource was found.
func (rq *ResourceQuery) First(ctx context.Context) (*Resource, error) {
	nodes, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{resource.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ResourceQuery) FirstX(ctx context.Context) *Resource {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Resource id in the query. Returns *NotFoundError when no id was found.
func (rq *ResourceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{resource.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ResourceQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Resource entity in the query, returns an error if not exactly one entity was returned.
func (rq *ResourceQuery) Only(ctx context.Context) (*Resource, error) {
	nodes, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{resource.Label}
	default:
		return nil, &NotSingularError{resource.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ResourceQuery) OnlyX(ctx context.Context) *Resource {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only Resource id in the query, returns an error if not exactly one id was returned.
func (rq *ResourceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{resource.Label}
	default:
		err = &NotSingularError{resource.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ResourceQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Resources.
func (rq *ResourceQuery) All(ctx context.Context) ([]*Resource, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *ResourceQuery) AllX(ctx context.Context) []*Resource {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Resource ids.
func (rq *ResourceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rq.Select(resource.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ResourceQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ResourceQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ResourceQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ResourceQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ResourceQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ResourceQuery) Clone() *ResourceQuery {
	return &ResourceQuery{
		config:     rq.config,
		limit:      rq.limit,
		offset:     rq.offset,
		order:      append([]OrderFunc{}, rq.order...),
		unique:     append([]string{}, rq.unique...),
		predicates: append([]predicate.Resource{}, rq.predicates...),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

//  WithPool tells the query-builder to eager-loads the nodes that are connected to
// the "pool" edge. The optional arguments used to configure the query builder of the edge.
func (rq *ResourceQuery) WithPool(opts ...func(*ResourcePoolQuery)) *ResourceQuery {
	query := &ResourcePoolQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withPool = query
	return rq
}

//  WithProperties tells the query-builder to eager-loads the nodes that are connected to
// the "properties" edge. The optional arguments used to configure the query builder of the edge.
func (rq *ResourceQuery) WithProperties(opts ...func(*PropertyQuery)) *ResourceQuery {
	query := &PropertyQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withProperties = query
	return rq
}

//  WithNestedPool tells the query-builder to eager-loads the nodes that are connected to
// the "nested_pool" edge. The optional arguments used to configure the query builder of the edge.
func (rq *ResourceQuery) WithNestedPool(opts ...func(*ResourcePoolQuery)) *ResourceQuery {
	query := &ResourcePoolQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withNestedPool = query
	return rq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status resource.Status `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Resource.Query().
//		GroupBy(resource.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rq *ResourceQuery) GroupBy(field string, fields ...string) *ResourceGroupBy {
	group := &ResourceGroupBy{config: rq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Status resource.Status `json:"status,omitempty"`
//	}
//
//	client.Resource.Query().
//		Select(resource.FieldStatus).
//		Scan(ctx, &v)
//
func (rq *ResourceQuery) Select(field string, fields ...string) *ResourceSelect {
	selector := &ResourceSelect{config: rq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return selector
}

func (rq *ResourceQuery) prepareQuery(ctx context.Context) error {
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	if err := resource.Policy.EvalQuery(ctx, rq); err != nil {
		return err
	}
	return nil
}

func (rq *ResourceQuery) sqlAll(ctx context.Context) ([]*Resource, error) {
	var (
		nodes       = []*Resource{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [3]bool{
			rq.withPool != nil,
			rq.withProperties != nil,
			rq.withNestedPool != nil,
		}
	)
	if rq.withPool != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, resource.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Resource{config: rq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rq.withPool; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Resource)
		for i := range nodes {
			if fk := nodes[i].resource_pool_claims; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(resourcepool.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "resource_pool_claims" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Pool = n
			}
		}
	}

	if query := rq.withProperties; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Resource)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Properties = []*Property{}
		}
		query.withFKs = true
		query.Where(predicate.Property(func(s *sql.Selector) {
			s.Where(sql.InValues(resource.PropertiesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.resource_properties
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "resource_properties" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "resource_properties" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Properties = append(node.Edges.Properties, n)
		}
	}

	if query := rq.withNestedPool; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Resource)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.ResourcePool(func(s *sql.Selector) {
			s.Where(sql.InValues(resource.NestedPoolColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.resource_nested_pool
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "resource_nested_pool" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "resource_nested_pool" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.NestedPool = n
		}
	}

	return nodes, nil
}

func (rq *ResourceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ResourceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rq *ResourceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resource.Table,
			Columns: resource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resource.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, resource.ValidColumn)
			}
		}
	}
	return _spec
}

func (rq *ResourceQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(resource.Table)
	selector := builder.Select(t1.Columns(resource.Columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(resource.Columns...)...)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector, resource.ValidColumn)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ResourceGroupBy is the builder for group-by Resource entities.
type ResourceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ResourceGroupBy) Aggregate(fns ...AggregateFunc) *ResourceGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scan the result into the given value.
func (rgb *ResourceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rgb *ResourceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rgb *ResourceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ResourceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rgb *ResourceGroupBy) StringsX(ctx context.Context) []string {
	v, err := rgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rgb *ResourceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resource.Label}
	default:
		err = fmt.Errorf("ent: ResourceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rgb *ResourceGroupBy) StringX(ctx context.Context) string {
	v, err := rgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rgb *ResourceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ResourceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rgb *ResourceGroupBy) IntsX(ctx context.Context) []int {
	v, err := rgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rgb *ResourceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resource.Label}
	default:
		err = fmt.Errorf("ent: ResourceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rgb *ResourceGroupBy) IntX(ctx context.Context) int {
	v, err := rgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rgb *ResourceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ResourceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rgb *ResourceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rgb *ResourceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resource.Label}
	default:
		err = fmt.Errorf("ent: ResourceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rgb *ResourceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rgb *ResourceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ResourceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rgb *ResourceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rgb *ResourceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resource.Label}
	default:
		err = fmt.Errorf("ent: ResourceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rgb *ResourceGroupBy) BoolX(ctx context.Context) bool {
	v, err := rgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rgb *ResourceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rgb.fields {
		if !resource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *ResourceGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql
	columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
	columns = append(columns, rgb.fields...)
	for _, fn := range rgb.fns {
		columns = append(columns, fn(selector, resource.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(rgb.fields...)
}

// ResourceSelect is the builder for select fields of Resource entities.
type ResourceSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (rs *ResourceSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := rs.path(ctx)
	if err != nil {
		return err
	}
	rs.sql = query
	return rs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rs *ResourceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (rs *ResourceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ResourceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rs *ResourceSelect) StringsX(ctx context.Context) []string {
	v, err := rs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (rs *ResourceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resource.Label}
	default:
		err = fmt.Errorf("ent: ResourceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rs *ResourceSelect) StringX(ctx context.Context) string {
	v, err := rs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (rs *ResourceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ResourceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rs *ResourceSelect) IntsX(ctx context.Context) []int {
	v, err := rs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (rs *ResourceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resource.Label}
	default:
		err = fmt.Errorf("ent: ResourceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rs *ResourceSelect) IntX(ctx context.Context) int {
	v, err := rs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (rs *ResourceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ResourceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rs *ResourceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (rs *ResourceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resource.Label}
	default:
		err = fmt.Errorf("ent: ResourceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rs *ResourceSelect) Float64X(ctx context.Context) float64 {
	v, err := rs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (rs *ResourceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ResourceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rs *ResourceSelect) BoolsX(ctx context.Context) []bool {
	v, err := rs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (rs *ResourceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resource.Label}
	default:
		err = fmt.Errorf("ent: ResourceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rs *ResourceSelect) BoolX(ctx context.Context) bool {
	v, err := rs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rs *ResourceSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rs.fields {
		if !resource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := rs.sqlQuery().Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rs *ResourceSelect) sqlQuery() sql.Querier {
	selector := rs.sql
	selector.Select(selector.Columns(rs.fields...)...)
	return selector
}
