// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/aquasecurity/trivy-operator/pkg/apis/aquasecurity/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterComplianceDetailReports implements ClusterComplianceDetailReportInterface
type FakeClusterComplianceDetailReports struct {
	Fake *FakeAquasecurityV1alpha1
	ns   string
}

var clustercompliancedetailreportsResource = schema.GroupVersionResource{Group: "aquasecurity.github.io", Version: "v1alpha1", Resource: "clustercompliancedetailreports"}

var clustercompliancedetailreportsKind = schema.GroupVersionKind{Group: "aquasecurity.github.io", Version: "v1alpha1", Kind: "ClusterComplianceDetailReport"}

// Get takes name of the clusterComplianceDetailReport, and returns the corresponding clusterComplianceDetailReport object, and an error if there is any.
func (c *FakeClusterComplianceDetailReports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterComplianceDetailReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustercompliancedetailreportsResource, c.ns, name), &v1alpha1.ClusterComplianceDetailReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterComplianceDetailReport), err
}

// List takes label and field selectors, and returns the list of ClusterComplianceDetailReports that match those selectors.
func (c *FakeClusterComplianceDetailReports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterComplianceDetailReportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustercompliancedetailreportsResource, clustercompliancedetailreportsKind, c.ns, opts), &v1alpha1.ClusterComplianceDetailReportList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterComplianceDetailReportList{ListMeta: obj.(*v1alpha1.ClusterComplianceDetailReportList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterComplianceDetailReportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterComplianceDetailReports.
func (c *FakeClusterComplianceDetailReports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustercompliancedetailreportsResource, c.ns, opts))

}

// Create takes the representation of a clusterComplianceDetailReport and creates it.  Returns the server's representation of the clusterComplianceDetailReport, and an error, if there is any.
func (c *FakeClusterComplianceDetailReports) Create(ctx context.Context, clusterComplianceDetailReport *v1alpha1.ClusterComplianceDetailReport, opts v1.CreateOptions) (result *v1alpha1.ClusterComplianceDetailReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustercompliancedetailreportsResource, c.ns, clusterComplianceDetailReport), &v1alpha1.ClusterComplianceDetailReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterComplianceDetailReport), err
}

// Update takes the representation of a clusterComplianceDetailReport and updates it. Returns the server's representation of the clusterComplianceDetailReport, and an error, if there is any.
func (c *FakeClusterComplianceDetailReports) Update(ctx context.Context, clusterComplianceDetailReport *v1alpha1.ClusterComplianceDetailReport, opts v1.UpdateOptions) (result *v1alpha1.ClusterComplianceDetailReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustercompliancedetailreportsResource, c.ns, clusterComplianceDetailReport), &v1alpha1.ClusterComplianceDetailReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterComplianceDetailReport), err
}

// Delete takes name of the clusterComplianceDetailReport and deletes it. Returns an error if one occurs.
func (c *FakeClusterComplianceDetailReports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clustercompliancedetailreportsResource, c.ns, name, opts), &v1alpha1.ClusterComplianceDetailReport{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterComplianceDetailReports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustercompliancedetailreportsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterComplianceDetailReportList{})
	return err
}

// Patch applies the patch and returns the patched clusterComplianceDetailReport.
func (c *FakeClusterComplianceDetailReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterComplianceDetailReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustercompliancedetailreportsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterComplianceDetailReport{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterComplianceDetailReport), err
}
