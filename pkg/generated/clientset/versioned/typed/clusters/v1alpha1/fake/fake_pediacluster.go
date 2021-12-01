// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/clusterpedia-io/clusterpedia/pkg/apis/clusters/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePediaClusters implements PediaClusterInterface
type FakePediaClusters struct {
	Fake *FakeClustersV1alpha1
}

var pediaclustersResource = schema.GroupVersionResource{Group: "clusters.clusterpedia.io", Version: "v1alpha1", Resource: "pediaclusters"}

var pediaclustersKind = schema.GroupVersionKind{Group: "clusters.clusterpedia.io", Version: "v1alpha1", Kind: "PediaCluster"}

// Get takes name of the pediaCluster, and returns the corresponding pediaCluster object, and an error if there is any.
func (c *FakePediaClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PediaCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(pediaclustersResource, name), &v1alpha1.PediaCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PediaCluster), err
}

// List takes label and field selectors, and returns the list of PediaClusters that match those selectors.
func (c *FakePediaClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PediaClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(pediaclustersResource, pediaclustersKind, opts), &v1alpha1.PediaClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PediaClusterList{ListMeta: obj.(*v1alpha1.PediaClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.PediaClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pediaClusters.
func (c *FakePediaClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(pediaclustersResource, opts))
}

// Create takes the representation of a pediaCluster and creates it.  Returns the server's representation of the pediaCluster, and an error, if there is any.
func (c *FakePediaClusters) Create(ctx context.Context, pediaCluster *v1alpha1.PediaCluster, opts v1.CreateOptions) (result *v1alpha1.PediaCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(pediaclustersResource, pediaCluster), &v1alpha1.PediaCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PediaCluster), err
}

// Update takes the representation of a pediaCluster and updates it. Returns the server's representation of the pediaCluster, and an error, if there is any.
func (c *FakePediaClusters) Update(ctx context.Context, pediaCluster *v1alpha1.PediaCluster, opts v1.UpdateOptions) (result *v1alpha1.PediaCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(pediaclustersResource, pediaCluster), &v1alpha1.PediaCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PediaCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePediaClusters) UpdateStatus(ctx context.Context, pediaCluster *v1alpha1.PediaCluster, opts v1.UpdateOptions) (*v1alpha1.PediaCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(pediaclustersResource, "status", pediaCluster), &v1alpha1.PediaCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PediaCluster), err
}

// Delete takes name of the pediaCluster and deletes it. Returns an error if one occurs.
func (c *FakePediaClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(pediaclustersResource, name), &v1alpha1.PediaCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePediaClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(pediaclustersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PediaClusterList{})
	return err
}

// Patch applies the patch and returns the patched pediaCluster.
func (c *FakePediaClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PediaCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(pediaclustersResource, name, pt, data, subresources...), &v1alpha1.PediaCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PediaCluster), err
}
