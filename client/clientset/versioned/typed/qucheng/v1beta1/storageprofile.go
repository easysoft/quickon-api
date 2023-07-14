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

// StorageProfilesGetter has a method to return a StorageProfileInterface.
// A group's client should implement this interface.
type StorageProfilesGetter interface {
	StorageProfiles(namespace string) StorageProfileInterface
}

// StorageProfileInterface has methods to work with StorageProfile resources.
type StorageProfileInterface interface {
	Create(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.CreateOptions) (*v1beta1.StorageProfile, error)
	Update(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.UpdateOptions) (*v1beta1.StorageProfile, error)
	UpdateStatus(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.UpdateOptions) (*v1beta1.StorageProfile, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.StorageProfile, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.StorageProfileList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.StorageProfile, err error)
	StorageProfileExpansion
}

// storageProfiles implements StorageProfileInterface
type storageProfiles struct {
	client rest.Interface
	ns     string
}

// newStorageProfiles returns a StorageProfiles
func newStorageProfiles(c *QuchengV1beta1Client, namespace string) *storageProfiles {
	return &storageProfiles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the storageProfile, and returns the corresponding storageProfile object, and an error if there is any.
func (c *storageProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.StorageProfile, err error) {
	result = &v1beta1.StorageProfile{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storageprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageProfiles that match those selectors.
func (c *storageProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.StorageProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.StorageProfileList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storageprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageProfiles.
func (c *storageProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("storageprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a storageProfile and creates it.  Returns the server's representation of the storageProfile, and an error, if there is any.
func (c *storageProfiles) Create(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.CreateOptions) (result *v1beta1.StorageProfile, err error) {
	result = &v1beta1.StorageProfile{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("storageprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageProfile).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a storageProfile and updates it. Returns the server's representation of the storageProfile, and an error, if there is any.
func (c *storageProfiles) Update(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.UpdateOptions) (result *v1beta1.StorageProfile, err error) {
	result = &v1beta1.StorageProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storageprofiles").
		Name(storageProfile.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageProfile).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *storageProfiles) UpdateStatus(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.UpdateOptions) (result *v1beta1.StorageProfile, err error) {
	result = &v1beta1.StorageProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storageprofiles").
		Name(storageProfile.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageProfile).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the storageProfile and deletes it. Returns an error if one occurs.
func (c *storageProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storageprofiles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storageprofiles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched storageProfile.
func (c *storageProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.StorageProfile, err error) {
	result = &v1beta1.StorageProfile{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("storageprofiles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
