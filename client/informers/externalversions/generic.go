// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1beta1 "github.com/easysoft/quickon-api/qucheng/v1beta1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=qucheng, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("backups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().Backups().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("dbs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().Dbs().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("dbbackups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().DbBackups().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("dbdumpreplicas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().DbDumpReplicas().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("dbrestores"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().DbRestores().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("dbservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().DbServices().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("deletebackuprequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().DeleteBackupRequests().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("resticrepositories"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().ResticRepositories().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("resticsnapshotdumps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().ResticSnapshotDumps().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("restores"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().Restores().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("storageprofiles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().StorageProfiles().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("volumebackups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().VolumeBackups().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("volumerestores"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Qucheng().V1beta1().VolumeRestores().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
