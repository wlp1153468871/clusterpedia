package storage

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	pediainternal "github.com/clusterpedia-io/clusterpedia/pkg/apis/pedia"
)

type StorageFactory interface {
	GetResourceVersions(ctx context.Context, cluster string) (map[schema.GroupVersionResource]map[string]interface{}, error)
	CleanCluster(ctx context.Context, cluster string) error
	CleanClusterResource(ctx context.Context, cluster string, gvr schema.GroupVersionResource) error

	NewResourceStorage(config *ResourceStorageConfig) (ResourceStorage, error)
	NewCollectionResourceStorage(cr *pediainternal.CollectionResource) (CollectionResourceStorage, error)

	GetCollectionResources(ctx context.Context) ([]*pediainternal.CollectionResource, error)
}

type ResourceStorage interface {
	GetStorageConfig() *ResourceStorageConfig

	Get(ctx context.Context, cluster, namespace, name string, obj runtime.Object) error
	List(ctx context.Context, listObj runtime.Object, opts *pediainternal.ListOptions) error

	Create(ctx context.Context, cluster string, obj runtime.Object) error
	Update(ctx context.Context, cluster string, obj runtime.Object) error
	Delete(ctx context.Context, cluster string, obj runtime.Object) error
}

type CollectionResourceStorage interface {
	Get(ctx context.Context, opts *pediainternal.ListOptions) (*pediainternal.CollectionResource, error)
}

type ResourceStorageConfig struct {
	GroupResource        schema.GroupResource
	StorageGroupResource schema.GroupResource

	Codec          runtime.Codec
	StorageVersion schema.GroupVersion
	MemoryVersion  schema.GroupVersion
}
