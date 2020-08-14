// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package resolver

import (
	"context"
	"fmt"

	"github.com/net-auto/resourceManager/graph/graphql/generated"
	"github.com/net-auto/resourceManager/ent"
	"github.com/net-auto/resourceManager/graph/graphql/model"
)

// txResolver wraps a mutation resolver and executes every mutation under a transaction.
type txResolver struct {
	generated.MutationResolver
}

func (tr txResolver) WithTransaction(ctx context.Context, f func(context.Context, generated.MutationResolver) error) error {
	tx, err := ent.FromContext(ctx).Tx(ctx)
	if err != nil {
		return fmt.Errorf("creating transaction: %w", err)
	}
	ctx = ent.NewTxContext(ctx, tx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()
	ctx = ent.NewContext(ctx, tx.Client())
	if err := f(ctx, tr.MutationResolver); err != nil {
		if r := tx.Rollback(); r != nil {
			err = fmt.Errorf("rolling back transaction: %v", r)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func (tr txResolver) CreateAllocationStrategy(ctx context.Context, name string, script string) (*ent.AllocationStrategy, error) {
	var result, zero *ent.AllocationStrategy
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.CreateAllocationStrategy(ctx, name, script)
		return
	}); err != nil {
		return zero, err
	}
	if result != nil {
		result = result.Unwrap()
	}
	return result, nil
}

func (tr txResolver) DeleteAllocationStrategy(ctx context.Context, allocationStrategyID int) (*ent.AllocationStrategy, error) {
	var result, zero *ent.AllocationStrategy
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.DeleteAllocationStrategy(ctx, allocationStrategyID)
		return
	}); err != nil {
		return zero, err
	}
	if result != nil {
		result = result.Unwrap()
	}
	return result, nil
}

func (tr txResolver) TestAllocationStrategy(ctx context.Context, allocationStrategyID int, resourcePool model.ResourcePoolInput, currentResources []*model.ResourceInput, userInput map[string]interface{}) (map[string]interface{}, error) {
	var result, zero map[string]interface{}
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.TestAllocationStrategy(ctx, allocationStrategyID, resourcePool, currentResources, userInput)
		return
	}); err != nil {
		return zero, err
	}
	return result, nil
}

func (tr txResolver) ClaimResource(ctx context.Context, poolID int, userInput map[string]interface{}) (*ent.Resource, error) {
	var result, zero *ent.Resource
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.ClaimResource(ctx, poolID, userInput)
		return
	}); err != nil {
		return zero, err
	}
	if result != nil {
		result = result.Unwrap()
	}
	return result, nil
}

func (tr txResolver) FreeResource(ctx context.Context, input map[string]interface{}, poolID int) (string, error) {
	var result, zero string
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.FreeResource(ctx, input, poolID)
		return
	}); err != nil {
		return zero, err
	}
	return result, nil
}

func (tr txResolver) CreateSetPool(ctx context.Context, resourceTypeID int, poolName string, poolDealocationSafetyPeriod int, poolValues []map[string]interface{}) (*ent.ResourcePool, error) {
	var result, zero *ent.ResourcePool
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.CreateSetPool(ctx, resourceTypeID, poolName, poolDealocationSafetyPeriod, poolValues)
		return
	}); err != nil {
		return zero, err
	}
	if result != nil {
		result = result.Unwrap()
	}
	return result, nil
}

func (tr txResolver) CreateSingletonPool(ctx context.Context, resourceTypeID int, poolName string, poolValues []map[string]interface{}) (*ent.ResourcePool, error) {
	var result, zero *ent.ResourcePool
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.CreateSingletonPool(ctx, resourceTypeID, poolName, poolValues)
		return
	}); err != nil {
		return zero, err
	}
	if result != nil {
		result = result.Unwrap()
	}
	return result, nil
}

func (tr txResolver) CreateAllocatingPool(ctx context.Context, resourceTypeID int, poolName string, allocationStrategyID int, poolDealocationSafetyPeriod int) (*ent.ResourcePool, error) {
	var result, zero *ent.ResourcePool
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.CreateAllocatingPool(ctx, resourceTypeID, poolName, allocationStrategyID, poolDealocationSafetyPeriod)
		return
	}); err != nil {
		return zero, err
	}
	if result != nil {
		result = result.Unwrap()
	}
	return result, nil
}

func (tr txResolver) DeleteResourcePool(ctx context.Context, resourcePoolID int) (string, error) {
	var result, zero string
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.DeleteResourcePool(ctx, resourcePoolID)
		return
	}); err != nil {
		return zero, err
	}
	return result, nil
}

func (tr txResolver) CreateResourceType(ctx context.Context, resourceName string, resourceProperties map[string]interface{}) (*ent.ResourceType, error) {
	var result, zero *ent.ResourceType
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.CreateResourceType(ctx, resourceName, resourceProperties)
		return
	}); err != nil {
		return zero, err
	}
	if result != nil {
		result = result.Unwrap()
	}
	return result, nil
}

func (tr txResolver) DeleteResourceType(ctx context.Context, resourceTypeID int) (string, error) {
	var result, zero string
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.DeleteResourceType(ctx, resourceTypeID)
		return
	}); err != nil {
		return zero, err
	}
	return result, nil
}

func (tr txResolver) UpdateResourceTypeName(ctx context.Context, resourceTypeID int, resourceName string) (*ent.ResourceType, error) {
	var result, zero *ent.ResourceType
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.UpdateResourceTypeName(ctx, resourceTypeID, resourceName)
		return
	}); err != nil {
		return zero, err
	}
	if result != nil {
		result = result.Unwrap()
	}
	return result, nil
}
