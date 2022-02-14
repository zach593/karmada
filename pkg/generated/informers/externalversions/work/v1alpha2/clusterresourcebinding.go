// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	time "time"

	workv1alpha2 "github.com/zach593/karmada/pkg/apis/work/v1alpha2"
	versioned "github.com/zach593/karmada/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/zach593/karmada/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/zach593/karmada/pkg/generated/listers/work/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterResourceBindingInformer provides access to a shared informer and lister for
// ClusterResourceBindings.
type ClusterResourceBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.ClusterResourceBindingLister
}

type clusterResourceBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterResourceBindingInformer constructs a new informer for ClusterResourceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterResourceBindingInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterResourceBindingInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterResourceBindingInformer constructs a new informer for ClusterResourceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterResourceBindingInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkV1alpha2().ClusterResourceBindings().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkV1alpha2().ClusterResourceBindings().Watch(context.TODO(), options)
			},
		},
		&workv1alpha2.ClusterResourceBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterResourceBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterResourceBindingInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterResourceBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&workv1alpha2.ClusterResourceBinding{}, f.defaultInformer)
}

func (f *clusterResourceBindingInformer) Lister() v1alpha2.ClusterResourceBindingLister {
	return v1alpha2.NewClusterResourceBindingLister(f.Informer().GetIndexer())
}
