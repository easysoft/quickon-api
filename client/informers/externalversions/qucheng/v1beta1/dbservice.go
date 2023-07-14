/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	quchengv1beta1 "github.com/easysoft/qucheng-operator/apis/qucheng/v1beta1"
	versioned "github.com/easysoft/qucheng-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/easysoft/qucheng-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/easysoft/qucheng-operator/pkg/client/listers/qucheng/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DbServiceInformer provides access to a shared informer and lister for
// DbServices.
type DbServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.DbServiceLister
}

type dbServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDbServiceInformer constructs a new informer for DbService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDbServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDbServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDbServiceInformer constructs a new informer for DbService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDbServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuchengV1beta1().DbServices(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuchengV1beta1().DbServices(namespace).Watch(context.TODO(), options)
			},
		},
		&quchengv1beta1.DbService{},
		resyncPeriod,
		indexers,
	)
}

func (f *dbServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDbServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dbServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&quchengv1beta1.DbService{}, f.defaultInformer)
}

func (f *dbServiceInformer) Lister() v1beta1.DbServiceLister {
	return v1beta1.NewDbServiceLister(f.Informer().GetIndexer())
}
