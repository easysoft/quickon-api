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

// ResticRepositoryInformer provides access to a shared informer and lister for
// ResticRepositories.
type ResticRepositoryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ResticRepositoryLister
}

type resticRepositoryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResticRepositoryInformer constructs a new informer for ResticRepository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResticRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResticRepositoryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResticRepositoryInformer constructs a new informer for ResticRepository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResticRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuchengV1beta1().ResticRepositories(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.QuchengV1beta1().ResticRepositories(namespace).Watch(context.TODO(), options)
			},
		},
		&quchengv1beta1.ResticRepository{},
		resyncPeriod,
		indexers,
	)
}

func (f *resticRepositoryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResticRepositoryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resticRepositoryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&quchengv1beta1.ResticRepository{}, f.defaultInformer)
}

func (f *resticRepositoryInformer) Lister() v1beta1.ResticRepositoryLister {
	return v1beta1.NewResticRepositoryLister(f.Informer().GetIndexer())
}
