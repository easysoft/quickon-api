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

package fake

import (
	"context"

	v1beta1 "github.com/easysoft/qucheng-operator/apis/qucheng/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeleteBackupRequests implements DeleteBackupRequestInterface
type FakeDeleteBackupRequests struct {
	Fake *FakeQuchengV1beta1
	ns   string
}

var deletebackuprequestsResource = schema.GroupVersionResource{Group: "qucheng", Version: "v1beta1", Resource: "deletebackuprequests"}

var deletebackuprequestsKind = schema.GroupVersionKind{Group: "qucheng", Version: "v1beta1", Kind: "DeleteBackupRequest"}

// Get takes name of the deleteBackupRequest, and returns the corresponding deleteBackupRequest object, and an error if there is any.
func (c *FakeDeleteBackupRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DeleteBackupRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deletebackuprequestsResource, c.ns, name), &v1beta1.DeleteBackupRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DeleteBackupRequest), err
}

// List takes label and field selectors, and returns the list of DeleteBackupRequests that match those selectors.
func (c *FakeDeleteBackupRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DeleteBackupRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deletebackuprequestsResource, deletebackuprequestsKind, c.ns, opts), &v1beta1.DeleteBackupRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DeleteBackupRequestList{ListMeta: obj.(*v1beta1.DeleteBackupRequestList).ListMeta}
	for _, item := range obj.(*v1beta1.DeleteBackupRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deleteBackupRequests.
func (c *FakeDeleteBackupRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deletebackuprequestsResource, c.ns, opts))

}

// Create takes the representation of a deleteBackupRequest and creates it.  Returns the server's representation of the deleteBackupRequest, and an error, if there is any.
func (c *FakeDeleteBackupRequests) Create(ctx context.Context, deleteBackupRequest *v1beta1.DeleteBackupRequest, opts v1.CreateOptions) (result *v1beta1.DeleteBackupRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deletebackuprequestsResource, c.ns, deleteBackupRequest), &v1beta1.DeleteBackupRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DeleteBackupRequest), err
}

// Update takes the representation of a deleteBackupRequest and updates it. Returns the server's representation of the deleteBackupRequest, and an error, if there is any.
func (c *FakeDeleteBackupRequests) Update(ctx context.Context, deleteBackupRequest *v1beta1.DeleteBackupRequest, opts v1.UpdateOptions) (result *v1beta1.DeleteBackupRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deletebackuprequestsResource, c.ns, deleteBackupRequest), &v1beta1.DeleteBackupRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DeleteBackupRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeleteBackupRequests) UpdateStatus(ctx context.Context, deleteBackupRequest *v1beta1.DeleteBackupRequest, opts v1.UpdateOptions) (*v1beta1.DeleteBackupRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deletebackuprequestsResource, "status", c.ns, deleteBackupRequest), &v1beta1.DeleteBackupRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DeleteBackupRequest), err
}

// Delete takes name of the deleteBackupRequest and deletes it. Returns an error if one occurs.
func (c *FakeDeleteBackupRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(deletebackuprequestsResource, c.ns, name), &v1beta1.DeleteBackupRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeleteBackupRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deletebackuprequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DeleteBackupRequestList{})
	return err
}

// Patch applies the patch and returns the patched deleteBackupRequest.
func (c *FakeDeleteBackupRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DeleteBackupRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deletebackuprequestsResource, c.ns, name, pt, data, subresources...), &v1beta1.DeleteBackupRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DeleteBackupRequest), err
}
