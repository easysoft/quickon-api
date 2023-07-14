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
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/easysoft/qucheng-operator/apis/qucheng/v1beta1"
	scheme "github.com/easysoft/qucheng-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RestoresGetter has a method to return a RestoreInterface.
// A group's client should implement this interface.
type RestoresGetter interface {
	Restores(namespace string) RestoreInterface
}

// RestoreInterface has methods to work with Restore resources.
type RestoreInterface interface {
	Create(ctx context.Context, restore *v1beta1.Restore, opts v1.CreateOptions) (*v1beta1.Restore, error)
	Update(ctx context.Context, restore *v1beta1.Restore, opts v1.UpdateOptions) (*v1beta1.Restore, error)
	UpdateStatus(ctx context.Context, restore *v1beta1.Restore, opts v1.UpdateOptions) (*v1beta1.Restore, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Restore, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.RestoreList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Restore, err error)
	RestoreExpansion
}

// restores implements RestoreInterface
type restores struct {
	client rest.Interface
	ns     string
}

// newRestores returns a Restores
func newRestores(c *QuchengV1beta1Client, namespace string) *restores {
	return &restores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the restore, and returns the corresponding restore object, and an error if there is any.
func (c *restores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Restore, err error) {
	result = &v1beta1.Restore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("restores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Restores that match those selectors.
func (c *restores) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.RestoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.RestoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("restores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested restores.
func (c *restores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("restores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a restore and creates it.  Returns the server's representation of the restore, and an error, if there is any.
func (c *restores) Create(ctx context.Context, restore *v1beta1.Restore, opts v1.CreateOptions) (result *v1beta1.Restore, err error) {
	result = &v1beta1.Restore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("restores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(restore).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a restore and updates it. Returns the server's representation of the restore, and an error, if there is any.
func (c *restores) Update(ctx context.Context, restore *v1beta1.Restore, opts v1.UpdateOptions) (result *v1beta1.Restore, err error) {
	result = &v1beta1.Restore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("restores").
		Name(restore.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(restore).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *restores) UpdateStatus(ctx context.Context, restore *v1beta1.Restore, opts v1.UpdateOptions) (result *v1beta1.Restore, err error) {
	result = &v1beta1.Restore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("restores").
		Name(restore.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(restore).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the restore and deletes it. Returns an error if one occurs.
func (c *restores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("restores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *restores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("restores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched restore.
func (c *restores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Restore, err error) {
	result = &v1beta1.Restore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("restores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
