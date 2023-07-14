// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	versioned "github.com/easysoft/quickon-api/client/clientset/versioned"
	internalinterfaces "github.com/easysoft/quickon-api/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/easysoft/quickon-api/client/listers/qucheng/v1beta1"
	quchengv1beta1 "github.com/easysoft/quickon-api/qucheng/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DbDumpReplicaInformer provides access to a shared informer and lister for
// DbDumpReplicas.
type DbDumpReplicaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.DbDumpReplicaLister
}

type dbDumpReplicaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDbDumpReplicaInformer constructs a new informer for DbDumpReplica type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDbDumpReplicaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDbDumpReplicaInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDbDumpReplicaInformer constructs a new informer for DbDumpReplica type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDbDumpReplicaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuchengV1beta1().DbDumpReplicas(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuchengV1beta1().DbDumpReplicas(namespace).Watch(context.TODO(), options)
			},
		},
		&quchengv1beta1.DbDumpReplica{},
		resyncPeriod,
		indexers,
	)
}

func (f *dbDumpReplicaInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDbDumpReplicaInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dbDumpReplicaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&quchengv1beta1.DbDumpReplica{}, f.defaultInformer)
}

func (f *dbDumpReplicaInformer) Lister() v1beta1.DbDumpReplicaLister {
	return v1beta1.NewDbDumpReplicaLister(f.Informer().GetIndexer())
}