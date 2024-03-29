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

// FakeDbRestores implements DbRestoreInterface
type FakeDbRestores struct {
	Fake *FakeQuchengV1beta1
	ns   string
}

var dbrestoresResource = schema.GroupVersionResource{Group: "qucheng", Version: "v1beta1", Resource: "dbrestores"}

var dbrestoresKind = schema.GroupVersionKind{Group: "qucheng", Version: "v1beta1", Kind: "DbRestore"}

// Get takes name of the dbRestore, and returns the corresponding dbRestore object, and an error if there is any.
func (c *FakeDbRestores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DbRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dbrestoresResource, c.ns, name), &v1beta1.DbRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DbRestore), err
}

// List takes label and field selectors, and returns the list of DbRestores that match those selectors.
func (c *FakeDbRestores) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DbRestoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dbrestoresResource, dbrestoresKind, c.ns, opts), &v1beta1.DbRestoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DbRestoreList{ListMeta: obj.(*v1beta1.DbRestoreList).ListMeta}
	for _, item := range obj.(*v1beta1.DbRestoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dbRestores.
func (c *FakeDbRestores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dbrestoresResource, c.ns, opts))

}

// Create takes the representation of a dbRestore and creates it.  Returns the server's representation of the dbRestore, and an error, if there is any.
func (c *FakeDbRestores) Create(ctx context.Context, dbRestore *v1beta1.DbRestore, opts v1.CreateOptions) (result *v1beta1.DbRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dbrestoresResource, c.ns, dbRestore), &v1beta1.DbRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DbRestore), err
}

// Update takes the representation of a dbRestore and updates it. Returns the server's representation of the dbRestore, and an error, if there is any.
func (c *FakeDbRestores) Update(ctx context.Context, dbRestore *v1beta1.DbRestore, opts v1.UpdateOptions) (result *v1beta1.DbRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dbrestoresResource, c.ns, dbRestore), &v1beta1.DbRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DbRestore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDbRestores) UpdateStatus(ctx context.Context, dbRestore *v1beta1.DbRestore, opts v1.UpdateOptions) (*v1beta1.DbRestore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dbrestoresResource, "status", c.ns, dbRestore), &v1beta1.DbRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DbRestore), err
}

// Delete takes name of the dbRestore and deletes it. Returns an error if one occurs.
func (c *FakeDbRestores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dbrestoresResource, c.ns, name, opts), &v1beta1.DbRestore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDbRestores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dbrestoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DbRestoreList{})
	return err
}

// Patch applies the patch and returns the patched dbRestore.
func (c *FakeDbRestores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DbRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dbrestoresResource, c.ns, name, pt, data, subresources...), &v1beta1.DbRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DbRestore), err
}
