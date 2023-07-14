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

// FakeResticRepositories implements ResticRepositoryInterface
type FakeResticRepositories struct {
	Fake *FakeQuchengV1beta1
	ns   string
}

var resticrepositoriesResource = schema.GroupVersionResource{Group: "qucheng", Version: "v1beta1", Resource: "resticrepositories"}

var resticrepositoriesKind = schema.GroupVersionKind{Group: "qucheng", Version: "v1beta1", Kind: "ResticRepository"}

// Get takes name of the resticRepository, and returns the corresponding resticRepository object, and an error if there is any.
func (c *FakeResticRepositories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ResticRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resticrepositoriesResource, c.ns, name), &v1beta1.ResticRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ResticRepository), err
}

// List takes label and field selectors, and returns the list of ResticRepositories that match those selectors.
func (c *FakeResticRepositories) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ResticRepositoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resticrepositoriesResource, resticrepositoriesKind, c.ns, opts), &v1beta1.ResticRepositoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ResticRepositoryList{ListMeta: obj.(*v1beta1.ResticRepositoryList).ListMeta}
	for _, item := range obj.(*v1beta1.ResticRepositoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resticRepositories.
func (c *FakeResticRepositories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resticrepositoriesResource, c.ns, opts))

}

// Create takes the representation of a resticRepository and creates it.  Returns the server's representation of the resticRepository, and an error, if there is any.
func (c *FakeResticRepositories) Create(ctx context.Context, resticRepository *v1beta1.ResticRepository, opts v1.CreateOptions) (result *v1beta1.ResticRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resticrepositoriesResource, c.ns, resticRepository), &v1beta1.ResticRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ResticRepository), err
}

// Update takes the representation of a resticRepository and updates it. Returns the server's representation of the resticRepository, and an error, if there is any.
func (c *FakeResticRepositories) Update(ctx context.Context, resticRepository *v1beta1.ResticRepository, opts v1.UpdateOptions) (result *v1beta1.ResticRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resticrepositoriesResource, c.ns, resticRepository), &v1beta1.ResticRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ResticRepository), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResticRepositories) UpdateStatus(ctx context.Context, resticRepository *v1beta1.ResticRepository, opts v1.UpdateOptions) (*v1beta1.ResticRepository, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resticrepositoriesResource, "status", c.ns, resticRepository), &v1beta1.ResticRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ResticRepository), err
}

// Delete takes name of the resticRepository and deletes it. Returns an error if one occurs.
func (c *FakeResticRepositories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resticrepositoriesResource, c.ns, name, opts), &v1beta1.ResticRepository{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResticRepositories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resticrepositoriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ResticRepositoryList{})
	return err
}

// Patch applies the patch and returns the patched resticRepository.
func (c *FakeResticRepositories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ResticRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resticrepositoriesResource, c.ns, name, pt, data, subresources...), &v1beta1.ResticRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ResticRepository), err
}