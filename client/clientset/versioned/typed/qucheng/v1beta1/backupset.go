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

// BackupSetsGetter has a method to return a BackupSetInterface.
// A group's client should implement this interface.
type BackupSetsGetter interface {
	BackupSets(namespace string) BackupSetInterface
}

// BackupSetInterface has methods to work with BackupSet resources.
type BackupSetInterface interface {
	Create(ctx context.Context, backupSet *v1beta1.BackupSet, opts v1.CreateOptions) (*v1beta1.BackupSet, error)
	Update(ctx context.Context, backupSet *v1beta1.BackupSet, opts v1.UpdateOptions) (*v1beta1.BackupSet, error)
	UpdateStatus(ctx context.Context, backupSet *v1beta1.BackupSet, opts v1.UpdateOptions) (*v1beta1.BackupSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.BackupSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.BackupSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackupSet, err error)
	BackupSetExpansion
}

// backupSets implements BackupSetInterface
type backupSets struct {
	client rest.Interface
	ns     string
}

// newBackupSets returns a BackupSets
func newBackupSets(c *QuchengV1beta1Client, namespace string) *backupSets {
	return &backupSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backupSet, and returns the corresponding backupSet object, and an error if there is any.
func (c *backupSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackupSet, err error) {
	result = &v1beta1.BackupSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupSets that match those selectors.
func (c *backupSets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackupSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BackupSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupSets.
func (c *backupSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backupsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backupSet and creates it.  Returns the server's representation of the backupSet, and an error, if there is any.
func (c *backupSets) Create(ctx context.Context, backupSet *v1beta1.BackupSet, opts v1.CreateOptions) (result *v1beta1.BackupSet, err error) {
	result = &v1beta1.BackupSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backupsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backupSet and updates it. Returns the server's representation of the backupSet, and an error, if there is any.
func (c *backupSets) Update(ctx context.Context, backupSet *v1beta1.BackupSet, opts v1.UpdateOptions) (result *v1beta1.BackupSet, err error) {
	result = &v1beta1.BackupSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupsets").
		Name(backupSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *backupSets) UpdateStatus(ctx context.Context, backupSet *v1beta1.BackupSet, opts v1.UpdateOptions) (result *v1beta1.BackupSet, err error) {
	result = &v1beta1.BackupSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupsets").
		Name(backupSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backupSet and deletes it. Returns an error if one occurs.
func (c *backupSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupsets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupsets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backupSet.
func (c *backupSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackupSet, err error) {
	result = &v1beta1.BackupSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backupsets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
