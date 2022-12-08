// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	scheme "github.com/easysoft/quikon-api/client/clientset/versioned/scheme"
	v1beta1 "github.com/easysoft/quikon-api/qucheng/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeleteBackupRequestsGetter has a method to return a DeleteBackupRequestInterface.
// A group's client should implement this interface.
type DeleteBackupRequestsGetter interface {
	DeleteBackupRequests(namespace string) DeleteBackupRequestInterface
}

// DeleteBackupRequestInterface has methods to work with DeleteBackupRequest resources.
type DeleteBackupRequestInterface interface {
	Create(ctx context.Context, deleteBackupRequest *v1beta1.DeleteBackupRequest, opts v1.CreateOptions) (*v1beta1.DeleteBackupRequest, error)
	Update(ctx context.Context, deleteBackupRequest *v1beta1.DeleteBackupRequest, opts v1.UpdateOptions) (*v1beta1.DeleteBackupRequest, error)
	UpdateStatus(ctx context.Context, deleteBackupRequest *v1beta1.DeleteBackupRequest, opts v1.UpdateOptions) (*v1beta1.DeleteBackupRequest, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.DeleteBackupRequest, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DeleteBackupRequestList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DeleteBackupRequest, err error)
	DeleteBackupRequestExpansion
}

// deleteBackupRequests implements DeleteBackupRequestInterface
type deleteBackupRequests struct {
	client rest.Interface
	ns     string
}

// newDeleteBackupRequests returns a DeleteBackupRequests
func newDeleteBackupRequests(c *QuchengV1beta1Client, namespace string) *deleteBackupRequests {
	return &deleteBackupRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deleteBackupRequest, and returns the corresponding deleteBackupRequest object, and an error if there is any.
func (c *deleteBackupRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DeleteBackupRequest, err error) {
	result = &v1beta1.DeleteBackupRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeleteBackupRequests that match those selectors.
func (c *deleteBackupRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DeleteBackupRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.DeleteBackupRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deleteBackupRequests.
func (c *deleteBackupRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deleteBackupRequest and creates it.  Returns the server's representation of the deleteBackupRequest, and an error, if there is any.
func (c *deleteBackupRequests) Create(ctx context.Context, deleteBackupRequest *v1beta1.DeleteBackupRequest, opts v1.CreateOptions) (result *v1beta1.DeleteBackupRequest, err error) {
	result = &v1beta1.DeleteBackupRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deleteBackupRequest).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deleteBackupRequest and updates it. Returns the server's representation of the deleteBackupRequest, and an error, if there is any.
func (c *deleteBackupRequests) Update(ctx context.Context, deleteBackupRequest *v1beta1.DeleteBackupRequest, opts v1.UpdateOptions) (result *v1beta1.DeleteBackupRequest, err error) {
	result = &v1beta1.DeleteBackupRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		Name(deleteBackupRequest.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deleteBackupRequest).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *deleteBackupRequests) UpdateStatus(ctx context.Context, deleteBackupRequest *v1beta1.DeleteBackupRequest, opts v1.UpdateOptions) (result *v1beta1.DeleteBackupRequest, err error) {
	result = &v1beta1.DeleteBackupRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		Name(deleteBackupRequest.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deleteBackupRequest).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deleteBackupRequest and deletes it. Returns an error if one occurs.
func (c *deleteBackupRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deleteBackupRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deleteBackupRequest.
func (c *deleteBackupRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DeleteBackupRequest, err error) {
	result = &v1beta1.DeleteBackupRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deletebackuprequests").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
