// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha2 "github.com/zach593/karmada/pkg/apis/work/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterResourceBindings implements ClusterResourceBindingInterface
type FakeClusterResourceBindings struct {
	Fake *FakeWorkV1alpha2
}

var clusterresourcebindingsResource = schema.GroupVersionResource{Group: "work.karmada.io", Version: "v1alpha2", Resource: "clusterresourcebindings"}

var clusterresourcebindingsKind = schema.GroupVersionKind{Group: "work.karmada.io", Version: "v1alpha2", Kind: "ClusterResourceBinding"}

// Get takes name of the clusterResourceBinding, and returns the corresponding clusterResourceBinding object, and an error if there is any.
func (c *FakeClusterResourceBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.ClusterResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterresourcebindingsResource, name), &v1alpha2.ClusterResourceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ClusterResourceBinding), err
}

// List takes label and field selectors, and returns the list of ClusterResourceBindings that match those selectors.
func (c *FakeClusterResourceBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.ClusterResourceBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterresourcebindingsResource, clusterresourcebindingsKind, opts), &v1alpha2.ClusterResourceBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.ClusterResourceBindingList{ListMeta: obj.(*v1alpha2.ClusterResourceBindingList).ListMeta}
	for _, item := range obj.(*v1alpha2.ClusterResourceBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterResourceBindings.
func (c *FakeClusterResourceBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterresourcebindingsResource, opts))
}

// Create takes the representation of a clusterResourceBinding and creates it.  Returns the server's representation of the clusterResourceBinding, and an error, if there is any.
func (c *FakeClusterResourceBindings) Create(ctx context.Context, clusterResourceBinding *v1alpha2.ClusterResourceBinding, opts v1.CreateOptions) (result *v1alpha2.ClusterResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterresourcebindingsResource, clusterResourceBinding), &v1alpha2.ClusterResourceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ClusterResourceBinding), err
}

// Update takes the representation of a clusterResourceBinding and updates it. Returns the server's representation of the clusterResourceBinding, and an error, if there is any.
func (c *FakeClusterResourceBindings) Update(ctx context.Context, clusterResourceBinding *v1alpha2.ClusterResourceBinding, opts v1.UpdateOptions) (result *v1alpha2.ClusterResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterresourcebindingsResource, clusterResourceBinding), &v1alpha2.ClusterResourceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ClusterResourceBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterResourceBindings) UpdateStatus(ctx context.Context, clusterResourceBinding *v1alpha2.ClusterResourceBinding, opts v1.UpdateOptions) (*v1alpha2.ClusterResourceBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterresourcebindingsResource, "status", clusterResourceBinding), &v1alpha2.ClusterResourceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ClusterResourceBinding), err
}

// Delete takes name of the clusterResourceBinding and deletes it. Returns an error if one occurs.
func (c *FakeClusterResourceBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterresourcebindingsResource, name), &v1alpha2.ClusterResourceBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterResourceBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterresourcebindingsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.ClusterResourceBindingList{})
	return err
}

// Patch applies the patch and returns the patched clusterResourceBinding.
func (c *FakeClusterResourceBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.ClusterResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterresourcebindingsResource, name, pt, data, subresources...), &v1alpha2.ClusterResourceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.ClusterResourceBinding), err
}
