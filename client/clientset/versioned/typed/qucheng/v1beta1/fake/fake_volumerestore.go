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

// FakeVolumeRestores implements VolumeRestoreInterface
type FakeVolumeRestores struct {
	Fake *FakeQuchengV1beta1
	ns   string
}

var volumerestoresResource = schema.GroupVersionResource{Group: "qucheng", Version: "v1beta1", Resource: "volumerestores"}

var volumerestoresKind = schema.GroupVersionKind{Group: "qucheng", Version: "v1beta1", Kind: "VolumeRestore"}

// Get takes name of the volumeRestore, and returns the corresponding volumeRestore object, and an error if there is any.
func (c *FakeVolumeRestores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumerestoresResource, c.ns, name), &v1beta1.VolumeRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeRestore), err
}

// List takes label and field selectors, and returns the list of VolumeRestores that match those selectors.
func (c *FakeVolumeRestores) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeRestoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumerestoresResource, volumerestoresKind, c.ns, opts), &v1beta1.VolumeRestoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VolumeRestoreList{ListMeta: obj.(*v1beta1.VolumeRestoreList).ListMeta}
	for _, item := range obj.(*v1beta1.VolumeRestoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeRestores.
func (c *FakeVolumeRestores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumerestoresResource, c.ns, opts))

}

// Create takes the representation of a volumeRestore and creates it.  Returns the server's representation of the volumeRestore, and an error, if there is any.
func (c *FakeVolumeRestores) Create(ctx context.Context, volumeRestore *v1beta1.VolumeRestore, opts v1.CreateOptions) (result *v1beta1.VolumeRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumerestoresResource, c.ns, volumeRestore), &v1beta1.VolumeRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeRestore), err
}

// Update takes the representation of a volumeRestore and updates it. Returns the server's representation of the volumeRestore, and an error, if there is any.
func (c *FakeVolumeRestores) Update(ctx context.Context, volumeRestore *v1beta1.VolumeRestore, opts v1.UpdateOptions) (result *v1beta1.VolumeRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumerestoresResource, c.ns, volumeRestore), &v1beta1.VolumeRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeRestore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeRestores) UpdateStatus(ctx context.Context, volumeRestore *v1beta1.VolumeRestore, opts v1.UpdateOptions) (*v1beta1.VolumeRestore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(volumerestoresResource, "status", c.ns, volumeRestore), &v1beta1.VolumeRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeRestore), err
}

// Delete takes name of the volumeRestore and deletes it. Returns an error if one occurs.
func (c *FakeVolumeRestores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(volumerestoresResource, c.ns, name, opts), &v1beta1.VolumeRestore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeRestores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(volumerestoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VolumeRestoreList{})
	return err
}

// Patch applies the patch and returns the patched volumeRestore.
func (c *FakeVolumeRestores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumerestoresResource, c.ns, name, pt, data, subresources...), &v1beta1.VolumeRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeRestore), err
}