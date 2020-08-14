// Code generated by entc, DO NOT EDIT.

package resourcepool

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/net-auto/resourceManager/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// DealocationSafetyPeriod applies equality check predicate on the "dealocation_safety_period" field. It's identical to DealocationSafetyPeriodEQ.
func DealocationSafetyPeriod(v int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDealocationSafetyPeriod), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ResourcePool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResourcePool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ResourcePool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResourcePool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// PoolTypeEQ applies the EQ predicate on the "pool_type" field.
func PoolTypeEQ(v PoolType) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPoolType), v))
	})
}

// PoolTypeNEQ applies the NEQ predicate on the "pool_type" field.
func PoolTypeNEQ(v PoolType) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPoolType), v))
	})
}

// PoolTypeIn applies the In predicate on the "pool_type" field.
func PoolTypeIn(vs ...PoolType) predicate.ResourcePool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResourcePool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPoolType), v...))
	})
}

// PoolTypeNotIn applies the NotIn predicate on the "pool_type" field.
func PoolTypeNotIn(vs ...PoolType) predicate.ResourcePool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResourcePool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPoolType), v...))
	})
}

// DealocationSafetyPeriodEQ applies the EQ predicate on the "dealocation_safety_period" field.
func DealocationSafetyPeriodEQ(v int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDealocationSafetyPeriod), v))
	})
}

// DealocationSafetyPeriodNEQ applies the NEQ predicate on the "dealocation_safety_period" field.
func DealocationSafetyPeriodNEQ(v int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDealocationSafetyPeriod), v))
	})
}

// DealocationSafetyPeriodIn applies the In predicate on the "dealocation_safety_period" field.
func DealocationSafetyPeriodIn(vs ...int) predicate.ResourcePool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResourcePool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDealocationSafetyPeriod), v...))
	})
}

// DealocationSafetyPeriodNotIn applies the NotIn predicate on the "dealocation_safety_period" field.
func DealocationSafetyPeriodNotIn(vs ...int) predicate.ResourcePool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResourcePool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDealocationSafetyPeriod), v...))
	})
}

// DealocationSafetyPeriodGT applies the GT predicate on the "dealocation_safety_period" field.
func DealocationSafetyPeriodGT(v int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDealocationSafetyPeriod), v))
	})
}

// DealocationSafetyPeriodGTE applies the GTE predicate on the "dealocation_safety_period" field.
func DealocationSafetyPeriodGTE(v int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDealocationSafetyPeriod), v))
	})
}

// DealocationSafetyPeriodLT applies the LT predicate on the "dealocation_safety_period" field.
func DealocationSafetyPeriodLT(v int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDealocationSafetyPeriod), v))
	})
}

// DealocationSafetyPeriodLTE applies the LTE predicate on the "dealocation_safety_period" field.
func DealocationSafetyPeriodLTE(v int) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDealocationSafetyPeriod), v))
	})
}

// HasResourceType applies the HasEdge predicate on the "resource_type" edge.
func HasResourceType() predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResourceTypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResourceTypeTable, ResourceTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResourceTypeWith applies the HasEdge predicate on the "resource_type" edge with a given conditions (other predicates).
func HasResourceTypeWith(preds ...predicate.ResourceType) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResourceTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResourceTypeTable, ResourceTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClaims applies the HasEdge predicate on the "claims" edge.
func HasClaims() predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClaimsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClaimsTable, ClaimsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClaimsWith applies the HasEdge predicate on the "claims" edge with a given conditions (other predicates).
func HasClaimsWith(preds ...predicate.Resource) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClaimsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClaimsTable, ClaimsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAllocationStrategy applies the HasEdge predicate on the "allocation_strategy" edge.
func HasAllocationStrategy() predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AllocationStrategyTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AllocationStrategyTable, AllocationStrategyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAllocationStrategyWith applies the HasEdge predicate on the "allocation_strategy" edge with a given conditions (other predicates).
func HasAllocationStrategyWith(preds ...predicate.AllocationStrategy) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AllocationStrategyInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AllocationStrategyTable, AllocationStrategyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.ResourcePool) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.ResourcePool) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ResourcePool) predicate.ResourcePool {
	return predicate.ResourcePool(func(s *sql.Selector) {
		p(s.Not())
	})
}
