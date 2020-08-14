// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (as *AllocationStrategyQuery) CollectFields(ctx context.Context, satisfies ...string) *AllocationStrategyQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		as = as.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return as
}

func (as *AllocationStrategyQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *AllocationStrategyQuery {
	return as
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pr *PropertyQuery) CollectFields(ctx context.Context, satisfies ...string) *PropertyQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		pr = pr.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return pr
}

func (pr *PropertyQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *PropertyQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "propertyType":
			pr = pr.WithType(func(query *PropertyTypeQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return pr
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pt *PropertyTypeQuery) CollectFields(ctx context.Context, satisfies ...string) *PropertyTypeQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		pt = pt.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return pt
}

func (pt *PropertyTypeQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *PropertyTypeQuery {
	return pt
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *ResourceQuery) CollectFields(ctx context.Context, satisfies ...string) *ResourceQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		r = r.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return r
}

func (r *ResourceQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ResourceQuery {
	return r
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (rp *ResourcePoolQuery) CollectFields(ctx context.Context, satisfies ...string) *ResourcePoolQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		rp = rp.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return rp
}

func (rp *ResourcePoolQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ResourcePoolQuery {
	return rp
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (rt *ResourceTypeQuery) CollectFields(ctx context.Context, satisfies ...string) *ResourceTypeQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		rt = rt.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return rt
}

func (rt *ResourceTypeQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ResourceTypeQuery {
	return rt
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TagQuery) CollectFields(ctx context.Context, satisfies ...string) *TagQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *TagQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TagQuery {
	return t
}
