// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/easysoft/quikon-api/qucheng/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DbBackupLister helps list DbBackups.
// All objects returned here must be treated as read-only.
type DbBackupLister interface {
	// List lists all DbBackups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.DbBackup, err error)
	// DbBackups returns an object that can list and get DbBackups.
	DbBackups(namespace string) DbBackupNamespaceLister
	DbBackupListerExpansion
}

// dbBackupLister implements the DbBackupLister interface.
type dbBackupLister struct {
	indexer cache.Indexer
}

// NewDbBackupLister returns a new DbBackupLister.
func NewDbBackupLister(indexer cache.Indexer) DbBackupLister {
	return &dbBackupLister{indexer: indexer}
}

// List lists all DbBackups in the indexer.
func (s *dbBackupLister) List(selector labels.Selector) (ret []*v1beta1.DbBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.DbBackup))
	})
	return ret, err
}

// DbBackups returns an object that can list and get DbBackups.
func (s *dbBackupLister) DbBackups(namespace string) DbBackupNamespaceLister {
	return dbBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DbBackupNamespaceLister helps list and get DbBackups.
// All objects returned here must be treated as read-only.
type DbBackupNamespaceLister interface {
	// List lists all DbBackups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.DbBackup, err error)
	// Get retrieves the DbBackup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.DbBackup, error)
	DbBackupNamespaceListerExpansion
}

// dbBackupNamespaceLister implements the DbBackupNamespaceLister
// interface.
type dbBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DbBackups in the indexer for a given namespace.
func (s dbBackupNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.DbBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.DbBackup))
	})
	return ret, err
}

// Get retrieves the DbBackup from the indexer for a given namespace and name.
func (s dbBackupNamespaceLister) Get(name string) (*v1beta1.DbBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("dbbackup"), name)
	}
	return obj.(*v1beta1.DbBackup), nil
}