// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/easysoft/quickon-api/qucheng/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageProfiles implements StorageProfileInterface
type FakeStorageProfiles struct {
	Fake *FakeQuchengV1beta1
	ns   string
}

var storageprofilesResource = schema.GroupVersionResource{Group: "qucheng", Version: "v1beta1", Resource: "storageprofiles"}

var storageprofilesKind = schema.GroupVersionKind{Group: "qucheng", Version: "v1beta1", Kind: "StorageProfile"}

// Get takes name of the storageProfile, and returns the corresponding storageProfile object, and an error if there is any.
func (c *FakeStorageProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.StorageProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storageprofilesResource, c.ns, name), &v1beta1.StorageProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageProfile), err
}

// List takes label and field selectors, and returns the list of StorageProfiles that match those selectors.
func (c *FakeStorageProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.StorageProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storageprofilesResource, storageprofilesKind, c.ns, opts), &v1beta1.StorageProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.StorageProfileList{ListMeta: obj.(*v1beta1.StorageProfileList).ListMeta}
	for _, item := range obj.(*v1beta1.StorageProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageProfiles.
func (c *FakeStorageProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storageprofilesResource, c.ns, opts))

}

// Create takes the representation of a storageProfile and creates it.  Returns the server's representation of the storageProfile, and an error, if there is any.
func (c *FakeStorageProfiles) Create(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.CreateOptions) (result *v1beta1.StorageProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storageprofilesResource, c.ns, storageProfile), &v1beta1.StorageProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageProfile), err
}

// Update takes the representation of a storageProfile and updates it. Returns the server's representation of the storageProfile, and an error, if there is any.
func (c *FakeStorageProfiles) Update(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.UpdateOptions) (result *v1beta1.StorageProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storageprofilesResource, c.ns, storageProfile), &v1beta1.StorageProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageProfiles) UpdateStatus(ctx context.Context, storageProfile *v1beta1.StorageProfile, opts v1.UpdateOptions) (*v1beta1.StorageProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storageprofilesResource, "status", c.ns, storageProfile), &v1beta1.StorageProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageProfile), err
}

// Delete takes name of the storageProfile and deletes it. Returns an error if one occurs.
func (c *FakeStorageProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(storageprofilesResource, c.ns, name, opts), &v1beta1.StorageProfile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storageprofilesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.StorageProfileList{})
	return err
}

// Patch applies the patch and returns the patched storageProfile.
func (c *FakeStorageProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.StorageProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storageprofilesResource, c.ns, name, pt, data, subresources...), &v1beta1.StorageProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageProfile), err
}
