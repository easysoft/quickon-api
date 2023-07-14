// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	scheme "github.com/easysoft/quickon-api/client/clientset/versioned/scheme"
	v1beta1 "github.com/easysoft/quickon-api/qucheng/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DbDumpReplicasGetter has a method to return a DbDumpReplicaInterface.
// A group's client should implement this interface.
type DbDumpReplicasGetter interface {
	DbDumpReplicas(namespace string) DbDumpReplicaInterface
}

// DbDumpReplicaInterface has methods to work with DbDumpReplica resources.
type DbDumpReplicaInterface interface {
	Create(ctx context.Context, dbDumpReplica *v1beta1.DbDumpReplica, opts v1.CreateOptions) (*v1beta1.DbDumpReplica, error)
	Update(ctx context.Context, dbDumpReplica *v1beta1.DbDumpReplica, opts v1.UpdateOptions) (*v1beta1.DbDumpReplica, error)
	UpdateStatus(ctx context.Context, dbDumpReplica *v1beta1.DbDumpReplica, opts v1.UpdateOptions) (*v1beta1.DbDumpReplica, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.DbDumpReplica, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DbDumpReplicaList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DbDumpReplica, err error)
	DbDumpReplicaExpansion
}

// dbDumpReplicas implements DbDumpReplicaInterface
type dbDumpReplicas struct {
	client rest.Interface
	ns     string
}

// newDbDumpReplicas returns a DbDumpReplicas
func newDbDumpReplicas(c *QuchengV1beta1Client, namespace string) *dbDumpReplicas {
	return &dbDumpReplicas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dbDumpReplica, and returns the corresponding dbDumpReplica object, and an error if there is any.
func (c *dbDumpReplicas) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DbDumpReplica, err error) {
	result = &v1beta1.DbDumpReplica{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbdumpreplicas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DbDumpReplicas that match those selectors.
func (c *dbDumpReplicas) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DbDumpReplicaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.DbDumpReplicaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbdumpreplicas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dbDumpReplicas.
func (c *dbDumpReplicas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dbdumpreplicas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dbDumpReplica and creates it.  Returns the server's representation of the dbDumpReplica, and an error, if there is any.
func (c *dbDumpReplicas) Create(ctx context.Context, dbDumpReplica *v1beta1.DbDumpReplica, opts v1.CreateOptions) (result *v1beta1.DbDumpReplica, err error) {
	result = &v1beta1.DbDumpReplica{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dbdumpreplicas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dbDumpReplica).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dbDumpReplica and updates it. Returns the server's representation of the dbDumpReplica, and an error, if there is any.
func (c *dbDumpReplicas) Update(ctx context.Context, dbDumpReplica *v1beta1.DbDumpReplica, opts v1.UpdateOptions) (result *v1beta1.DbDumpReplica, err error) {
	result = &v1beta1.DbDumpReplica{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbdumpreplicas").
		Name(dbDumpReplica.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dbDumpReplica).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dbDumpReplicas) UpdateStatus(ctx context.Context, dbDumpReplica *v1beta1.DbDumpReplica, opts v1.UpdateOptions) (result *v1beta1.DbDumpReplica, err error) {
	result = &v1beta1.DbDumpReplica{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbdumpreplicas").
		Name(dbDumpReplica.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dbDumpReplica).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dbDumpReplica and deletes it. Returns an error if one occurs.
func (c *dbDumpReplicas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbdumpreplicas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dbDumpReplicas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbdumpreplicas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dbDumpReplica.
func (c *dbDumpReplicas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DbDumpReplica, err error) {
	result = &v1beta1.DbDumpReplica{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dbdumpreplicas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
