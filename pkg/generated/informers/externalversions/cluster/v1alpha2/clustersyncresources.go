// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	time "time"

	clusterv1alpha2 "github.com/clusterpedia-io/api/cluster/v1alpha2"
	versioned "github.com/clusterpedia-io/clusterpedia/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/clusterpedia-io/clusterpedia/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/clusterpedia-io/clusterpedia/pkg/generated/listers/cluster/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterSyncResourcesInformer provides access to a shared informer and lister for
// ClusterSyncResources.
type ClusterSyncResourcesInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.ClusterSyncResourcesLister
}

type clusterSyncResourcesInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterSyncResourcesInformer constructs a new informer for ClusterSyncResources type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterSyncResourcesInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterSyncResourcesInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterSyncResourcesInformer constructs a new informer for ClusterSyncResources type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterSyncResourcesInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV1alpha2().ClusterSyncResources().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV1alpha2().ClusterSyncResources().Watch(context.TODO(), options)
			},
		},
		&clusterv1alpha2.ClusterSyncResources{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterSyncResourcesInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterSyncResourcesInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterSyncResourcesInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clusterv1alpha2.ClusterSyncResources{}, f.defaultInformer)
}

func (f *clusterSyncResourcesInformer) Lister() v1alpha2.ClusterSyncResourcesLister {
	return v1alpha2.NewClusterSyncResourcesLister(f.Informer().GetIndexer())
}
