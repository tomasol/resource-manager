// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package resolver

import (
	"context"
	"fmt"

	"github.com/net-auto/resourceManager/graph/graphql/generated"
	"github.com/net-auto/resourceManager/ent"
	"github.com/net-auto/resourceManager/ent/resourcepool"
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

func (tr txResolver) ClaimResource(ctx context.Context, poolName string) (*ent.Resource, error) {
	var result, zero *ent.Resource
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.ClaimResource(ctx, poolName)
		return
	}); err != nil {
		return zero, err
	}
	if result != nil {
		result = result.Unwrap()
	}
	return result, nil
}

func (tr txResolver) FreeResource(ctx context.Context, input map[string]interface{}, poolName string) (string, error) {
	var result, zero string
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.FreeResource(ctx, input, poolName)
		return
	}); err != nil {
		return zero, err
	}
	return result, nil
}

func (tr txResolver) CreatePool(ctx context.Context, poolType *resourcepool.PoolType, resourceTypeID int, poolName string, poolValues []map[string]interface{}, allocationScript string) (*ent.ResourcePool, error) {
	var result, zero *ent.ResourcePool
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.CreatePool(ctx, poolType, resourceTypeID, poolName, poolValues, allocationScript)
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

func (tr txResolver) UpdateResourcePool(ctx context.Context, resourcePoolID int, poolName string, poolValues []map[string]interface{}, allocationScript string) (string, error) {
	var result, zero string
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.UpdateResourcePool(ctx, resourcePoolID, poolName, poolValues, allocationScript)
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

func (tr txResolver) AddResourceTypeProperty(ctx context.Context, resourceTypeID int, resourceProperties map[string]interface{}) (*ent.ResourceType, error) {
	var result, zero *ent.ResourceType
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.AddResourceTypeProperty(ctx, resourceTypeID, resourceProperties)
		return
	}); err != nil {
		return zero, err
	}
	if result != nil {
		result = result.Unwrap()
	}
	return result, nil
}

func (tr txResolver) AddExistingPropertyToResourceType(ctx context.Context, resourceTypeID int, propertyTypeID int) (int, error) {
	var result, zero int
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.AddExistingPropertyToResourceType(ctx, resourceTypeID, propertyTypeID)
		return
	}); err != nil {
		return zero, err
	}
	return result, nil
}

func (tr txResolver) RemoveResourceTypeProperty(ctx context.Context, resourceTypeID int, propertyTypeID int) (*ent.ResourceType, error) {
	var result, zero *ent.ResourceType
	if err := tr.WithTransaction(ctx, func(ctx context.Context, mr generated.MutationResolver) (err error) {
		result, err = mr.RemoveResourceTypeProperty(ctx, resourceTypeID, propertyTypeID)
		return
	}); err != nil {
		return zero, err
	}
	if result != nil {
		result = result.Unwrap()
	}
	return result, nil
}
