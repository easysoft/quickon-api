// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/easysoft/quikon-api/qucheng/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DbServiceLister helps list DbServices.
// All objects returned here must be treated as read-only.
type DbServiceLister interface {
	// List lists all DbServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.DbService, err error)
	// DbServices returns an object that can list and get DbServices.
	DbServices(namespace string) DbServiceNamespaceLister
	DbServiceListerExpansion
}

// dbServiceLister implements the DbServiceLister interface.
type dbServiceLister struct {
	indexer cache.Indexer
}

// NewDbServiceLister returns a new DbServiceLister.
func NewDbServiceLister(indexer cache.Indexer) DbServiceLister {
	return &dbServiceLister{indexer: indexer}
}

// List lists all DbServices in the indexer.
func (s *dbServiceLister) List(selector labels.Selector) (ret []*v1beta1.DbService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.DbService))
	})
	return ret, err
}

// DbServices returns an object that can list and get DbServices.
func (s *dbServiceLister) DbServices(namespace string) DbServiceNamespaceLister {
	return dbServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DbServiceNamespaceLister helps list and get DbServices.
// All objects returned here must be treated as read-only.
type DbServiceNamespaceLister interface {
	// List lists all DbServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.DbService, err error)
	// Get retrieves the DbService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.DbService, error)
	DbServiceNamespaceListerExpansion
}

// dbServiceNamespaceLister implements the DbServiceNamespaceLister
// interface.
type dbServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DbServices in the indexer for a given namespace.
func (s dbServiceNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.DbService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.DbService))
	})
	return ret, err
}

// Get retrieves the DbService from the indexer for a given namespace and name.
func (s dbServiceNamespaceLister) Get(name string) (*v1beta1.DbService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("dbservice"), name)
	}
	return obj.(*v1beta1.DbService), nil
}
