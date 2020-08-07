// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/net-auto/resourceManager/ent/predicate"
	"github.com/net-auto/resourceManager/ent/property"
	"github.com/net-auto/resourceManager/ent/propertytype"
)

// PropertyUpdate is the builder for updating Property entities.
type PropertyUpdate struct {
	config
	hooks      []Hook
	mutation   *PropertyMutation
	predicates []predicate.Property
}

// Where adds a new predicate for the builder.
func (pu *PropertyUpdate) Where(ps ...predicate.Property) *PropertyUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetIntVal sets the int_val field.
func (pu *PropertyUpdate) SetIntVal(i int) *PropertyUpdate {
	pu.mutation.ResetIntVal()
	pu.mutation.SetIntVal(i)
	return pu
}

// SetNillableIntVal sets the int_val field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableIntVal(i *int) *PropertyUpdate {
	if i != nil {
		pu.SetIntVal(*i)
	}
	return pu
}

// AddIntVal adds i to int_val.
func (pu *PropertyUpdate) AddIntVal(i int) *PropertyUpdate {
	pu.mutation.AddIntVal(i)
	return pu
}

// ClearIntVal clears the value of int_val.
func (pu *PropertyUpdate) ClearIntVal() *PropertyUpdate {
	pu.mutation.ClearIntVal()
	return pu
}

// SetBoolVal sets the bool_val field.
func (pu *PropertyUpdate) SetBoolVal(b bool) *PropertyUpdate {
	pu.mutation.SetBoolVal(b)
	return pu
}

// SetNillableBoolVal sets the bool_val field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableBoolVal(b *bool) *PropertyUpdate {
	if b != nil {
		pu.SetBoolVal(*b)
	}
	return pu
}

// ClearBoolVal clears the value of bool_val.
func (pu *PropertyUpdate) ClearBoolVal() *PropertyUpdate {
	pu.mutation.ClearBoolVal()
	return pu
}

// SetFloatVal sets the float_val field.
func (pu *PropertyUpdate) SetFloatVal(f float64) *PropertyUpdate {
	pu.mutation.ResetFloatVal()
	pu.mutation.SetFloatVal(f)
	return pu
}

// SetNillableFloatVal sets the float_val field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableFloatVal(f *float64) *PropertyUpdate {
	if f != nil {
		pu.SetFloatVal(*f)
	}
	return pu
}

// AddFloatVal adds f to float_val.
func (pu *PropertyUpdate) AddFloatVal(f float64) *PropertyUpdate {
	pu.mutation.AddFloatVal(f)
	return pu
}

// ClearFloatVal clears the value of float_val.
func (pu *PropertyUpdate) ClearFloatVal() *PropertyUpdate {
	pu.mutation.ClearFloatVal()
	return pu
}

// SetLatitudeVal sets the latitude_val field.
func (pu *PropertyUpdate) SetLatitudeVal(f float64) *PropertyUpdate {
	pu.mutation.ResetLatitudeVal()
	pu.mutation.SetLatitudeVal(f)
	return pu
}

// SetNillableLatitudeVal sets the latitude_val field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableLatitudeVal(f *float64) *PropertyUpdate {
	if f != nil {
		pu.SetLatitudeVal(*f)
	}
	return pu
}

// AddLatitudeVal adds f to latitude_val.
func (pu *PropertyUpdate) AddLatitudeVal(f float64) *PropertyUpdate {
	pu.mutation.AddLatitudeVal(f)
	return pu
}

// ClearLatitudeVal clears the value of latitude_val.
func (pu *PropertyUpdate) ClearLatitudeVal() *PropertyUpdate {
	pu.mutation.ClearLatitudeVal()
	return pu
}

// SetLongitudeVal sets the longitude_val field.
func (pu *PropertyUpdate) SetLongitudeVal(f float64) *PropertyUpdate {
	pu.mutation.ResetLongitudeVal()
	pu.mutation.SetLongitudeVal(f)
	return pu
}

// SetNillableLongitudeVal sets the longitude_val field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableLongitudeVal(f *float64) *PropertyUpdate {
	if f != nil {
		pu.SetLongitudeVal(*f)
	}
	return pu
}

// AddLongitudeVal adds f to longitude_val.
func (pu *PropertyUpdate) AddLongitudeVal(f float64) *PropertyUpdate {
	pu.mutation.AddLongitudeVal(f)
	return pu
}

// ClearLongitudeVal clears the value of longitude_val.
func (pu *PropertyUpdate) ClearLongitudeVal() *PropertyUpdate {
	pu.mutation.ClearLongitudeVal()
	return pu
}

// SetRangeFromVal sets the range_from_val field.
func (pu *PropertyUpdate) SetRangeFromVal(f float64) *PropertyUpdate {
	pu.mutation.ResetRangeFromVal()
	pu.mutation.SetRangeFromVal(f)
	return pu
}

// SetNillableRangeFromVal sets the range_from_val field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableRangeFromVal(f *float64) *PropertyUpdate {
	if f != nil {
		pu.SetRangeFromVal(*f)
	}
	return pu
}

// AddRangeFromVal adds f to range_from_val.
func (pu *PropertyUpdate) AddRangeFromVal(f float64) *PropertyUpdate {
	pu.mutation.AddRangeFromVal(f)
	return pu
}

// ClearRangeFromVal clears the value of range_from_val.
func (pu *PropertyUpdate) ClearRangeFromVal() *PropertyUpdate {
	pu.mutation.ClearRangeFromVal()
	return pu
}

// SetRangeToVal sets the range_to_val field.
func (pu *PropertyUpdate) SetRangeToVal(f float64) *PropertyUpdate {
	pu.mutation.ResetRangeToVal()
	pu.mutation.SetRangeToVal(f)
	return pu
}

// SetNillableRangeToVal sets the range_to_val field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableRangeToVal(f *float64) *PropertyUpdate {
	if f != nil {
		pu.SetRangeToVal(*f)
	}
	return pu
}

// AddRangeToVal adds f to range_to_val.
func (pu *PropertyUpdate) AddRangeToVal(f float64) *PropertyUpdate {
	pu.mutation.AddRangeToVal(f)
	return pu
}

// ClearRangeToVal clears the value of range_to_val.
func (pu *PropertyUpdate) ClearRangeToVal() *PropertyUpdate {
	pu.mutation.ClearRangeToVal()
	return pu
}

// SetStringVal sets the string_val field.
func (pu *PropertyUpdate) SetStringVal(s string) *PropertyUpdate {
	pu.mutation.SetStringVal(s)
	return pu
}

// SetNillableStringVal sets the string_val field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableStringVal(s *string) *PropertyUpdate {
	if s != nil {
		pu.SetStringVal(*s)
	}
	return pu
}

// ClearStringVal clears the value of string_val.
func (pu *PropertyUpdate) ClearStringVal() *PropertyUpdate {
	pu.mutation.ClearStringVal()
	return pu
}

// SetTypeID sets the type edge to PropertyType by id.
func (pu *PropertyUpdate) SetTypeID(id int) *PropertyUpdate {
	pu.mutation.SetTypeID(id)
	return pu
}

// SetType sets the type edge to PropertyType.
func (pu *PropertyUpdate) SetType(p *PropertyType) *PropertyUpdate {
	return pu.SetTypeID(p.ID)
}

// Mutation returns the PropertyMutation object of the builder.
func (pu *PropertyUpdate) Mutation() *PropertyMutation {
	return pu.mutation
}

// ClearType clears the type edge to PropertyType.
func (pu *PropertyUpdate) ClearType() *PropertyUpdate {
	pu.mutation.ClearType()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PropertyUpdate) Save(ctx context.Context) (int, error) {

	if _, ok := pu.mutation.TypeID(); pu.mutation.TypeCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"type\"")
	}
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PropertyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PropertyUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PropertyUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PropertyUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PropertyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   property.Table,
			Columns: property.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: property.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.IntVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: property.FieldIntVal,
		})
	}
	if value, ok := pu.mutation.AddedIntVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: property.FieldIntVal,
		})
	}
	if pu.mutation.IntValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: property.FieldIntVal,
		})
	}
	if value, ok := pu.mutation.BoolVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: property.FieldBoolVal,
		})
	}
	if pu.mutation.BoolValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: property.FieldBoolVal,
		})
	}
	if value, ok := pu.mutation.FloatVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldFloatVal,
		})
	}
	if value, ok := pu.mutation.AddedFloatVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldFloatVal,
		})
	}
	if pu.mutation.FloatValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: property.FieldFloatVal,
		})
	}
	if value, ok := pu.mutation.LatitudeVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldLatitudeVal,
		})
	}
	if value, ok := pu.mutation.AddedLatitudeVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldLatitudeVal,
		})
	}
	if pu.mutation.LatitudeValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: property.FieldLatitudeVal,
		})
	}
	if value, ok := pu.mutation.LongitudeVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldLongitudeVal,
		})
	}
	if value, ok := pu.mutation.AddedLongitudeVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldLongitudeVal,
		})
	}
	if pu.mutation.LongitudeValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: property.FieldLongitudeVal,
		})
	}
	if value, ok := pu.mutation.RangeFromVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldRangeFromVal,
		})
	}
	if value, ok := pu.mutation.AddedRangeFromVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldRangeFromVal,
		})
	}
	if pu.mutation.RangeFromValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: property.FieldRangeFromVal,
		})
	}
	if value, ok := pu.mutation.RangeToVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldRangeToVal,
		})
	}
	if value, ok := pu.mutation.AddedRangeToVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldRangeToVal,
		})
	}
	if pu.mutation.RangeToValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: property.FieldRangeToVal,
		})
	}
	if value, ok := pu.mutation.StringVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: property.FieldStringVal,
		})
	}
	if pu.mutation.StringValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: property.FieldStringVal,
		})
	}
	if pu.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   property.TypeTable,
			Columns: []string{property.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: propertytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   property.TypeTable,
			Columns: []string{property.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: propertytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{property.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PropertyUpdateOne is the builder for updating a single Property entity.
type PropertyUpdateOne struct {
	config
	hooks    []Hook
	mutation *PropertyMutation
}

// SetIntVal sets the int_val field.
func (puo *PropertyUpdateOne) SetIntVal(i int) *PropertyUpdateOne {
	puo.mutation.ResetIntVal()
	puo.mutation.SetIntVal(i)
	return puo
}

// SetNillableIntVal sets the int_val field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableIntVal(i *int) *PropertyUpdateOne {
	if i != nil {
		puo.SetIntVal(*i)
	}
	return puo
}

// AddIntVal adds i to int_val.
func (puo *PropertyUpdateOne) AddIntVal(i int) *PropertyUpdateOne {
	puo.mutation.AddIntVal(i)
	return puo
}

// ClearIntVal clears the value of int_val.
func (puo *PropertyUpdateOne) ClearIntVal() *PropertyUpdateOne {
	puo.mutation.ClearIntVal()
	return puo
}

// SetBoolVal sets the bool_val field.
func (puo *PropertyUpdateOne) SetBoolVal(b bool) *PropertyUpdateOne {
	puo.mutation.SetBoolVal(b)
	return puo
}

// SetNillableBoolVal sets the bool_val field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableBoolVal(b *bool) *PropertyUpdateOne {
	if b != nil {
		puo.SetBoolVal(*b)
	}
	return puo
}

// ClearBoolVal clears the value of bool_val.
func (puo *PropertyUpdateOne) ClearBoolVal() *PropertyUpdateOne {
	puo.mutation.ClearBoolVal()
	return puo
}

// SetFloatVal sets the float_val field.
func (puo *PropertyUpdateOne) SetFloatVal(f float64) *PropertyUpdateOne {
	puo.mutation.ResetFloatVal()
	puo.mutation.SetFloatVal(f)
	return puo
}

// SetNillableFloatVal sets the float_val field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableFloatVal(f *float64) *PropertyUpdateOne {
	if f != nil {
		puo.SetFloatVal(*f)
	}
	return puo
}

// AddFloatVal adds f to float_val.
func (puo *PropertyUpdateOne) AddFloatVal(f float64) *PropertyUpdateOne {
	puo.mutation.AddFloatVal(f)
	return puo
}

// ClearFloatVal clears the value of float_val.
func (puo *PropertyUpdateOne) ClearFloatVal() *PropertyUpdateOne {
	puo.mutation.ClearFloatVal()
	return puo
}

// SetLatitudeVal sets the latitude_val field.
func (puo *PropertyUpdateOne) SetLatitudeVal(f float64) *PropertyUpdateOne {
	puo.mutation.ResetLatitudeVal()
	puo.mutation.SetLatitudeVal(f)
	return puo
}

// SetNillableLatitudeVal sets the latitude_val field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableLatitudeVal(f *float64) *PropertyUpdateOne {
	if f != nil {
		puo.SetLatitudeVal(*f)
	}
	return puo
}

// AddLatitudeVal adds f to latitude_val.
func (puo *PropertyUpdateOne) AddLatitudeVal(f float64) *PropertyUpdateOne {
	puo.mutation.AddLatitudeVal(f)
	return puo
}

// ClearLatitudeVal clears the value of latitude_val.
func (puo *PropertyUpdateOne) ClearLatitudeVal() *PropertyUpdateOne {
	puo.mutation.ClearLatitudeVal()
	return puo
}

// SetLongitudeVal sets the longitude_val field.
func (puo *PropertyUpdateOne) SetLongitudeVal(f float64) *PropertyUpdateOne {
	puo.mutation.ResetLongitudeVal()
	puo.mutation.SetLongitudeVal(f)
	return puo
}

// SetNillableLongitudeVal sets the longitude_val field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableLongitudeVal(f *float64) *PropertyUpdateOne {
	if f != nil {
		puo.SetLongitudeVal(*f)
	}
	return puo
}

// AddLongitudeVal adds f to longitude_val.
func (puo *PropertyUpdateOne) AddLongitudeVal(f float64) *PropertyUpdateOne {
	puo.mutation.AddLongitudeVal(f)
	return puo
}

// ClearLongitudeVal clears the value of longitude_val.
func (puo *PropertyUpdateOne) ClearLongitudeVal() *PropertyUpdateOne {
	puo.mutation.ClearLongitudeVal()
	return puo
}

// SetRangeFromVal sets the range_from_val field.
func (puo *PropertyUpdateOne) SetRangeFromVal(f float64) *PropertyUpdateOne {
	puo.mutation.ResetRangeFromVal()
	puo.mutation.SetRangeFromVal(f)
	return puo
}

// SetNillableRangeFromVal sets the range_from_val field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableRangeFromVal(f *float64) *PropertyUpdateOne {
	if f != nil {
		puo.SetRangeFromVal(*f)
	}
	return puo
}

// AddRangeFromVal adds f to range_from_val.
func (puo *PropertyUpdateOne) AddRangeFromVal(f float64) *PropertyUpdateOne {
	puo.mutation.AddRangeFromVal(f)
	return puo
}

// ClearRangeFromVal clears the value of range_from_val.
func (puo *PropertyUpdateOne) ClearRangeFromVal() *PropertyUpdateOne {
	puo.mutation.ClearRangeFromVal()
	return puo
}

// SetRangeToVal sets the range_to_val field.
func (puo *PropertyUpdateOne) SetRangeToVal(f float64) *PropertyUpdateOne {
	puo.mutation.ResetRangeToVal()
	puo.mutation.SetRangeToVal(f)
	return puo
}

// SetNillableRangeToVal sets the range_to_val field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableRangeToVal(f *float64) *PropertyUpdateOne {
	if f != nil {
		puo.SetRangeToVal(*f)
	}
	return puo
}

// AddRangeToVal adds f to range_to_val.
func (puo *PropertyUpdateOne) AddRangeToVal(f float64) *PropertyUpdateOne {
	puo.mutation.AddRangeToVal(f)
	return puo
}

// ClearRangeToVal clears the value of range_to_val.
func (puo *PropertyUpdateOne) ClearRangeToVal() *PropertyUpdateOne {
	puo.mutation.ClearRangeToVal()
	return puo
}

// SetStringVal sets the string_val field.
func (puo *PropertyUpdateOne) SetStringVal(s string) *PropertyUpdateOne {
	puo.mutation.SetStringVal(s)
	return puo
}

// SetNillableStringVal sets the string_val field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableStringVal(s *string) *PropertyUpdateOne {
	if s != nil {
		puo.SetStringVal(*s)
	}
	return puo
}

// ClearStringVal clears the value of string_val.
func (puo *PropertyUpdateOne) ClearStringVal() *PropertyUpdateOne {
	puo.mutation.ClearStringVal()
	return puo
}

// SetTypeID sets the type edge to PropertyType by id.
func (puo *PropertyUpdateOne) SetTypeID(id int) *PropertyUpdateOne {
	puo.mutation.SetTypeID(id)
	return puo
}

// SetType sets the type edge to PropertyType.
func (puo *PropertyUpdateOne) SetType(p *PropertyType) *PropertyUpdateOne {
	return puo.SetTypeID(p.ID)
}

// Mutation returns the PropertyMutation object of the builder.
func (puo *PropertyUpdateOne) Mutation() *PropertyMutation {
	return puo.mutation
}

// ClearType clears the type edge to PropertyType.
func (puo *PropertyUpdateOne) ClearType() *PropertyUpdateOne {
	puo.mutation.ClearType()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PropertyUpdateOne) Save(ctx context.Context) (*Property, error) {

	if _, ok := puo.mutation.TypeID(); puo.mutation.TypeCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"type\"")
	}
	var (
		err  error
		node *Property
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PropertyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PropertyUpdateOne) SaveX(ctx context.Context) *Property {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *PropertyUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PropertyUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PropertyUpdateOne) sqlSave(ctx context.Context) (pr *Property, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   property.Table,
			Columns: property.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: property.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Property.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.IntVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: property.FieldIntVal,
		})
	}
	if value, ok := puo.mutation.AddedIntVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: property.FieldIntVal,
		})
	}
	if puo.mutation.IntValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: property.FieldIntVal,
		})
	}
	if value, ok := puo.mutation.BoolVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: property.FieldBoolVal,
		})
	}
	if puo.mutation.BoolValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: property.FieldBoolVal,
		})
	}
	if value, ok := puo.mutation.FloatVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldFloatVal,
		})
	}
	if value, ok := puo.mutation.AddedFloatVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldFloatVal,
		})
	}
	if puo.mutation.FloatValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: property.FieldFloatVal,
		})
	}
	if value, ok := puo.mutation.LatitudeVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldLatitudeVal,
		})
	}
	if value, ok := puo.mutation.AddedLatitudeVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldLatitudeVal,
		})
	}
	if puo.mutation.LatitudeValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: property.FieldLatitudeVal,
		})
	}
	if value, ok := puo.mutation.LongitudeVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldLongitudeVal,
		})
	}
	if value, ok := puo.mutation.AddedLongitudeVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldLongitudeVal,
		})
	}
	if puo.mutation.LongitudeValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: property.FieldLongitudeVal,
		})
	}
	if value, ok := puo.mutation.RangeFromVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldRangeFromVal,
		})
	}
	if value, ok := puo.mutation.AddedRangeFromVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldRangeFromVal,
		})
	}
	if puo.mutation.RangeFromValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: property.FieldRangeFromVal,
		})
	}
	if value, ok := puo.mutation.RangeToVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldRangeToVal,
		})
	}
	if value, ok := puo.mutation.AddedRangeToVal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: property.FieldRangeToVal,
		})
	}
	if puo.mutation.RangeToValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: property.FieldRangeToVal,
		})
	}
	if value, ok := puo.mutation.StringVal(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: property.FieldStringVal,
		})
	}
	if puo.mutation.StringValCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: property.FieldStringVal,
		})
	}
	if puo.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   property.TypeTable,
			Columns: []string{property.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: propertytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   property.TypeTable,
			Columns: []string{property.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: propertytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Property{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{property.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
