// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/easysoft/quickon-api/qucheng/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VolumeRestoreLister helps list VolumeRestores.
// All objects returned here must be treated as read-only.
type VolumeRestoreLister interface {
	// List lists all VolumeRestores in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.VolumeRestore, err error)
	// VolumeRestores returns an object that can list and get VolumeRestores.
	VolumeRestores(namespace string) VolumeRestoreNamespaceLister
	VolumeRestoreListerExpansion
}

// volumeRestoreLister implements the VolumeRestoreLister interface.
type volumeRestoreLister struct {
	indexer cache.Indexer
}

// NewVolumeRestoreLister returns a new VolumeRestoreLister.
func NewVolumeRestoreLister(indexer cache.Indexer) VolumeRestoreLister {
	return &volumeRestoreLister{indexer: indexer}
}

// List lists all VolumeRestores in the indexer.
func (s *volumeRestoreLister) List(selector labels.Selector) (ret []*v1beta1.VolumeRestore, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeRestore))
	})
	return ret, err
}

// VolumeRestores returns an object that can list and get VolumeRestores.
func (s *volumeRestoreLister) VolumeRestores(namespace string) VolumeRestoreNamespaceLister {
	return volumeRestoreNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VolumeRestoreNamespaceLister helps list and get VolumeRestores.
// All objects returned here must be treated as read-only.
type VolumeRestoreNamespaceLister interface {
	// List lists all VolumeRestores in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.VolumeRestore, err error)
	// Get retrieves the VolumeRestore from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.VolumeRestore, error)
	VolumeRestoreNamespaceListerExpansion
}

// volumeRestoreNamespaceLister implements the VolumeRestoreNamespaceLister
// interface.
type volumeRestoreNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VolumeRestores in the indexer for a given namespace.
func (s volumeRestoreNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.VolumeRestore, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeRestore))
	})
	return ret, err
}

// Get retrieves the VolumeRestore from the indexer for a given namespace and name.
func (s volumeRestoreNamespaceLister) Get(name string) (*v1beta1.VolumeRestore, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("volumerestore"), name)
	}
	return obj.(*v1beta1.VolumeRestore), nil
}
