// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	m3dboperatorv1 "github.com/m3db/m3db-operator/pkg/apis/m3dboperator/v1"
	versioned "github.com/m3db/m3db-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/m3db/m3db-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/m3db/m3db-operator/pkg/client/listers/m3dboperator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// M3DBClusterInformer provides access to a shared informer and lister for
// M3DBClusters.
type M3DBClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.M3DBClusterLister
}

type m3DBClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewM3DBClusterInformer constructs a new informer for M3DBCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewM3DBClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredM3DBClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredM3DBClusterInformer constructs a new informer for M3DBCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredM3DBClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().M3DBClusters(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().M3DBClusters(namespace).Watch(options)
			},
		},
		&m3dboperatorv1.M3DBCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *m3DBClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredM3DBClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *m3DBClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&m3dboperatorv1.M3DBCluster{}, f.defaultInformer)
}

func (f *m3DBClusterInformer) Lister() v1.M3DBClusterLister {
	return v1.NewM3DBClusterLister(f.Informer().GetIndexer())
}