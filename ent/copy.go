// Code generated by entc, DO NOT EDIT.

package ent

// Copy existing AllocationStrategy entity with all its fields and (eagerly loaded) edges ! into a new AllocationStrategyCreate builder.
func (c *AllocationStrategyClient) CreateFrom(as *AllocationStrategy) *AllocationStrategyCreate {
	mutation := newAllocationStrategyMutation(c.config, OpCreate, withAllocationStrategy(as))
	mutation.name = &(as.Name)
	mutation.lang = &(as.Lang)
	mutation.script = &(as.Script)

	createBuilder := &AllocationStrategyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}

	if as.Edges.Pools != nil {
		createBuilder.AddPools(as.Edges.Pools...)
	}

	return createBuilder
}

// Copy existing Label entity with all its fields and (eagerly loaded) edges ! into a new LabelCreate builder.
func (c *LabelClient) CreateFrom(l *Label) *LabelCreate {
	mutation := newLabelMutation(c.config, OpCreate, withLabel(l))
	mutation.labl = &(l.Labl)

	createBuilder := &LabelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}

	if l.Edges.Pools != nil {
		createBuilder.AddPools(l.Edges.Pools...)
	}

	return createBuilder
}

// Copy existing Property entity with all its fields and (eagerly loaded) edges ! into a new PropertyCreate builder.
func (c *PropertyClient) CreateFrom(pr *Property) *PropertyCreate {
	mutation := newPropertyMutation(c.config, OpCreate, withProperty(pr))
	mutation.int_val = (pr.IntVal)
	mutation.bool_val = (pr.BoolVal)
	mutation.float_val = (pr.FloatVal)
	mutation.latitude_val = (pr.LatitudeVal)
	mutation.longitude_val = (pr.LongitudeVal)
	mutation.range_from_val = (pr.RangeFromVal)
	mutation.range_to_val = (pr.RangeToVal)
	mutation.string_val = (pr.StringVal)

	createBuilder := &PropertyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}

	if pr.Edges.Type != nil {
		createBuilder.SetType(pr.Edges.Type)
	}

	return createBuilder
}

// Copy existing PropertyType entity with all its fields and (eagerly loaded) edges ! into a new PropertyTypeCreate builder.
func (c *PropertyTypeClient) CreateFrom(pt *PropertyType) *PropertyTypeCreate {
	mutation := newPropertyTypeMutation(c.config, OpCreate, withPropertyType(pt))
	mutation._type = &(pt.Type)
	mutation.name = &(pt.Name)
	mutation.external_id = &(pt.ExternalID)
	mutation.index = &(pt.Index)
	mutation.category = &(pt.Category)
	mutation.int_val = (pt.IntVal)
	mutation.bool_val = (pt.BoolVal)
	mutation.float_val = (pt.FloatVal)
	mutation.latitude_val = (pt.LatitudeVal)
	mutation.longitude_val = (pt.LongitudeVal)
	mutation.string_val = (pt.StringVal)
	mutation.range_from_val = (pt.RangeFromVal)
	mutation.range_to_val = (pt.RangeToVal)
	mutation.is_instance_property = &(pt.IsInstanceProperty)
	mutation.editable = &(pt.Editable)
	mutation.mandatory = &(pt.Mandatory)
	mutation.deleted = &(pt.Deleted)
	mutation.nodeType = &(pt.NodeType)

	createBuilder := &PropertyTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}

	if pt.Edges.Properties != nil {
		createBuilder.AddProperties(pt.Edges.Properties...)
	}

	return createBuilder
}

// Copy existing Resource entity with all its fields and (eagerly loaded) edges ! into a new ResourceCreate builder.
func (c *ResourceClient) CreateFrom(r *Resource) *ResourceCreate {
	mutation := newResourceMutation(c.config, OpCreate, withResource(r))
	mutation.claimed = &(r.Claimed)

	createBuilder := &ResourceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}

	if r.Edges.Pool != nil {
		createBuilder.SetPool(r.Edges.Pool)
	}

	if r.Edges.Properties != nil {
		createBuilder.AddProperties(r.Edges.Properties...)
	}

	return createBuilder
}

// Copy existing ResourcePool entity with all its fields and (eagerly loaded) edges ! into a new ResourcePoolCreate builder.
func (c *ResourcePoolClient) CreateFrom(rp *ResourcePool) *ResourcePoolCreate {
	mutation := newResourcePoolMutation(c.config, OpCreate, withResourcePool(rp))
	mutation.name = &(rp.Name)
	mutation.pool_type = &(rp.PoolType)

	createBuilder := &ResourcePoolCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}

	if rp.Edges.ResourceType != nil {
		createBuilder.SetResourceType(rp.Edges.ResourceType)
	}

	if rp.Edges.Labels != nil {
		createBuilder.SetLabels(rp.Edges.Labels)
	}

	if rp.Edges.Claims != nil {
		createBuilder.AddClaims(rp.Edges.Claims...)
	}

	if rp.Edges.AllocationStrategy != nil {
		createBuilder.SetAllocationStrategy(rp.Edges.AllocationStrategy)
	}

	return createBuilder
}

// Copy existing ResourceType entity with all its fields and (eagerly loaded) edges ! into a new ResourceTypeCreate builder.
func (c *ResourceTypeClient) CreateFrom(rt *ResourceType) *ResourceTypeCreate {
	mutation := newResourceTypeMutation(c.config, OpCreate, withResourceType(rt))
	mutation.name = &(rt.Name)

	createBuilder := &ResourceTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}

	if rt.Edges.PropertyTypes != nil {
		createBuilder.AddPropertyTypes(rt.Edges.PropertyTypes...)
	}

	if rt.Edges.Pools != nil {
		createBuilder.AddPools(rt.Edges.Pools...)
	}

	return createBuilder
}
