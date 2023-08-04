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

// RestoreSetInformer provides access to a shared informer and lister for
// RestoreSets.
type RestoreSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.RestoreSetLister
}

type restoreSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRestoreSetInformer constructs a new informer for RestoreSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRestoreSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRestoreSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRestoreSetInformer constructs a new informer for RestoreSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRestoreSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuchengV1beta1().RestoreSets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuchengV1beta1().RestoreSets(namespace).Watch(context.TODO(), options)
			},
		},
		&quchengv1beta1.RestoreSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *restoreSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRestoreSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *restoreSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&quchengv1beta1.RestoreSet{}, f.defaultInformer)
}

func (f *restoreSetInformer) Lister() v1beta1.RestoreSetLister {
	return v1beta1.NewRestoreSetLister(f.Informer().GetIndexer())
}
