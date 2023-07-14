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

// ResticSnapshotDumpsGetter has a method to return a ResticSnapshotDumpInterface.
// A group's client should implement this interface.
type ResticSnapshotDumpsGetter interface {
	ResticSnapshotDumps(namespace string) ResticSnapshotDumpInterface
}

// ResticSnapshotDumpInterface has methods to work with ResticSnapshotDump resources.
type ResticSnapshotDumpInterface interface {
	Create(ctx context.Context, resticSnapshotDump *v1beta1.ResticSnapshotDump, opts v1.CreateOptions) (*v1beta1.ResticSnapshotDump, error)
	Update(ctx context.Context, resticSnapshotDump *v1beta1.ResticSnapshotDump, opts v1.UpdateOptions) (*v1beta1.ResticSnapshotDump, error)
	UpdateStatus(ctx context.Context, resticSnapshotDump *v1beta1.ResticSnapshotDump, opts v1.UpdateOptions) (*v1beta1.ResticSnapshotDump, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ResticSnapshotDump, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ResticSnapshotDumpList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ResticSnapshotDump, err error)
	ResticSnapshotDumpExpansion
}

// resticSnapshotDumps implements ResticSnapshotDumpInterface
type resticSnapshotDumps struct {
	client rest.Interface
	ns     string
}

// newResticSnapshotDumps returns a ResticSnapshotDumps
func newResticSnapshotDumps(c *QuchengV1beta1Client, namespace string) *resticSnapshotDumps {
	return &resticSnapshotDumps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resticSnapshotDump, and returns the corresponding resticSnapshotDump object, and an error if there is any.
func (c *resticSnapshotDumps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ResticSnapshotDump, err error) {
	result = &v1beta1.ResticSnapshotDump{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resticsnapshotdumps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResticSnapshotDumps that match those selectors.
func (c *resticSnapshotDumps) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ResticSnapshotDumpList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ResticSnapshotDumpList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resticsnapshotdumps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resticSnapshotDumps.
func (c *resticSnapshotDumps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resticsnapshotdumps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a resticSnapshotDump and creates it.  Returns the server's representation of the resticSnapshotDump, and an error, if there is any.
func (c *resticSnapshotDumps) Create(ctx context.Context, resticSnapshotDump *v1beta1.ResticSnapshotDump, opts v1.CreateOptions) (result *v1beta1.ResticSnapshotDump, err error) {
	result = &v1beta1.ResticSnapshotDump{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resticsnapshotdumps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resticSnapshotDump).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a resticSnapshotDump and updates it. Returns the server's representation of the resticSnapshotDump, and an error, if there is any.
func (c *resticSnapshotDumps) Update(ctx context.Context, resticSnapshotDump *v1beta1.ResticSnapshotDump, opts v1.UpdateOptions) (result *v1beta1.ResticSnapshotDump, err error) {
	result = &v1beta1.ResticSnapshotDump{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resticsnapshotdumps").
		Name(resticSnapshotDump.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resticSnapshotDump).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *resticSnapshotDumps) UpdateStatus(ctx context.Context, resticSnapshotDump *v1beta1.ResticSnapshotDump, opts v1.UpdateOptions) (result *v1beta1.ResticSnapshotDump, err error) {
	result = &v1beta1.ResticSnapshotDump{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resticsnapshotdumps").
		Name(resticSnapshotDump.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resticSnapshotDump).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the resticSnapshotDump and deletes it. Returns an error if one occurs.
func (c *resticSnapshotDumps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resticsnapshotdumps").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resticSnapshotDumps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resticsnapshotdumps").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched resticSnapshotDump.
func (c *resticSnapshotDumps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ResticSnapshotDump, err error) {
	result = &v1beta1.ResticSnapshotDump{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resticsnapshotdumps").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
