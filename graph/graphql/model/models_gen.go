// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/net-auto/resourceManager/ent"
	"github.com/net-auto/resourceManager/ent/allocationstrategy"
)

type CreateAllocatingPoolInput struct {
	ResourceTypeID              int                    `json:"resourceTypeId"`
	PoolName                    string                 `json:"poolName"`
	Description                 *string                `json:"description"`
	AllocationStrategyID        int                    `json:"allocationStrategyId"`
	PoolDealocationSafetyPeriod int                    `json:"poolDealocationSafetyPeriod"`
	PoolProperties              map[string]interface{} `json:"poolProperties"`
	PoolPropertyTypes           map[string]interface{} `json:"poolPropertyTypes"`
}

type CreateAllocatingPoolPayload struct {
	Pool *ent.ResourcePool `json:"pool"`
}

type CreateAllocationStrategyInput struct {
	Name        string                  `json:"name"`
	Description *string                 `json:"description"`
	Script      string                  `json:"script"`
	Lang        allocationstrategy.Lang `json:"lang"`
}

type CreateAllocationStrategyPayload struct {
	Strategy *ent.AllocationStrategy `json:"strategy"`
}

type CreateNestedAllocatingPoolInput struct {
	ResourceTypeID              int     `json:"resourceTypeId"`
	PoolName                    string  `json:"poolName"`
	Description                 *string `json:"description"`
	AllocationStrategyID        int     `json:"allocationStrategyId"`
	PoolDealocationSafetyPeriod int     `json:"poolDealocationSafetyPeriod"`
	ParentResourceID            int     `json:"parentResourceId"`
}

type CreateNestedAllocatingPoolPayload struct {
	Pool *ent.ResourcePool `json:"pool"`
}

type CreateNestedSetPoolInput struct {
	ResourceTypeID              int                      `json:"resourceTypeId"`
	PoolName                    string                   `json:"poolName"`
	Description                 *string                  `json:"description"`
	PoolDealocationSafetyPeriod int                      `json:"poolDealocationSafetyPeriod"`
	PoolValues                  []map[string]interface{} `json:"poolValues"`
	ParentResourceID            int                      `json:"parentResourceId"`
}

type CreateNestedSetPoolPayload struct {
	Pool *ent.ResourcePool `json:"pool"`
}

type CreateNestedSingletonPoolInput struct {
	ResourceTypeID   int                      `json:"resourceTypeId"`
	PoolName         string                   `json:"poolName"`
	Description      *string                  `json:"description"`
	PoolValues       []map[string]interface{} `json:"poolValues"`
	ParentResourceID int                      `json:"parentResourceId"`
}

type CreateNestedSingletonPoolPayload struct {
	Pool *ent.ResourcePool `json:"pool"`
}

type CreateResourceTypeInput struct {
	ResourceName       string                 `json:"resourceName"`
	ResourceProperties map[string]interface{} `json:"resourceProperties"`
}

type CreateResourceTypePayload struct {
	ResourceType *ent.ResourceType `json:"resourceType"`
}

type CreateSetPoolInput struct {
	ResourceTypeID              int                      `json:"resourceTypeId"`
	PoolName                    string                   `json:"poolName"`
	Description                 *string                  `json:"description"`
	PoolDealocationSafetyPeriod int                      `json:"poolDealocationSafetyPeriod"`
	PoolValues                  []map[string]interface{} `json:"poolValues"`
}

type CreateSetPoolPayload struct {
	Pool *ent.ResourcePool `json:"pool"`
}

type CreateSingletonPoolInput struct {
	ResourceTypeID int                      `json:"resourceTypeId"`
	PoolName       string                   `json:"poolName"`
	Description    *string                  `json:"description"`
	PoolValues     []map[string]interface{} `json:"poolValues"`
}

type CreateSingletonPoolPayload struct {
	Pool *ent.ResourcePool `json:"pool"`
}

type CreateTagInput struct {
	TagText string `json:"tagText"`
}

type CreateTagPayload struct {
	Tag *ent.Tag `json:"tag"`
}

type DeleteAllocationStrategyInput struct {
	AllocationStrategyID int `json:"allocationStrategyId"`
}

type DeleteAllocationStrategyPayload struct {
	Strategy *ent.AllocationStrategy `json:"strategy"`
}

type DeleteResourcePoolInput struct {
	ResourcePoolID int `json:"resourcePoolId"`
}

type DeleteResourcePoolPayload struct {
	ResourcePoolID int `json:"resourcePoolId"`
}

type DeleteResourceTypeInput struct {
	ResourceTypeID int `json:"resourceTypeId"`
}

type DeleteResourceTypePayload struct {
	ResourceTypeID int `json:"resourceTypeId"`
}

type DeleteTagInput struct {
	TagID int `json:"tagId"`
}

type DeleteTagPayload struct {
	TagID int `json:"tagId"`
}

type PageInfo struct {
	HasPreviousEdge bool `json:"hasPreviousEdge"`
	HasNextEdge     bool `json:"hasNextEdge"`
	StartCursor     int  `json:"startCursor"`
	EndCursor       int  `json:"endCursor"`
}

type PoolCapacityPayload struct {
	FreeCapacity     float64 `json:"freeCapacity"`
	UtilizedCapacity float64 `json:"utilizedCapacity"`
}

type ResourceConnection struct {
	PageInfo *PageInfo       `json:"pageInfo"`
	Edges    []*ResourceEdge `json:"edges"`
}

type ResourceEdge struct {
	Node   *ent.Resource `json:"node"`
	Cursor int           `json:"cursor"`
}

type ResourceInput struct {
	Properties map[string]interface{} `json:"Properties"`
	UpdatedAt  string                 `json:"UpdatedAt"`
	Status     string                 `json:"Status"`
}

type ResourcePoolInput struct {
	PoolProperties   map[string]interface{} `json:"poolProperties"`
	ResourcePoolName string                 `json:"ResourcePoolName"`
}

type TagAnd struct {
	MatchesAll []string `json:"matchesAll"`
}

type TagOr struct {
	MatchesAny []*TagAnd `json:"matchesAny"`
}

type TagPoolInput struct {
	TagID  int `json:"tagId"`
	PoolID int `json:"poolId"`
}

type TagPoolPayload struct {
	Tag *ent.Tag `json:"tag"`
}

type UntagPoolInput struct {
	TagID  int `json:"tagId"`
	PoolID int `json:"poolId"`
}

type UntagPoolPayload struct {
	Tag *ent.Tag `json:"tag"`
}

type UpdateResourceTypeNameInput struct {
	ResourceTypeID int    `json:"resourceTypeId"`
	ResourceName   string `json:"resourceName"`
}

type UpdateResourceTypeNamePayload struct {
	ResourceTypeID int `json:"resourceTypeId"`
}

type UpdateTagInput struct {
	TagID   int    `json:"tagId"`
	TagText string `json:"tagText"`
}

type UpdateTagPayload struct {
	Tag *ent.Tag `json:"tag"`
}
