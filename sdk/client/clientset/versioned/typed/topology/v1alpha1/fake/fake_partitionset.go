/*
Copyright The KCP Authors.

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
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/topology/v1alpha1"
	topologyv1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/topology/v1alpha1"
)

// FakePartitionSets implements PartitionSetInterface
type FakePartitionSets struct {
	Fake *FakeTopologyV1alpha1
}

var partitionsetsResource = v1alpha1.SchemeGroupVersion.WithResource("partitionsets")

var partitionsetsKind = v1alpha1.SchemeGroupVersion.WithKind("PartitionSet")

// Get takes name of the partitionSet, and returns the corresponding partitionSet object, and an error if there is any.
func (c *FakePartitionSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PartitionSet, err error) {
	emptyResult := &v1alpha1.PartitionSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(partitionsetsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PartitionSet), err
}

// List takes label and field selectors, and returns the list of PartitionSets that match those selectors.
func (c *FakePartitionSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PartitionSetList, err error) {
	emptyResult := &v1alpha1.PartitionSetList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(partitionsetsResource, partitionsetsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PartitionSetList{ListMeta: obj.(*v1alpha1.PartitionSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.PartitionSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested partitionSets.
func (c *FakePartitionSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(partitionsetsResource, opts))
}

// Create takes the representation of a partitionSet and creates it.  Returns the server's representation of the partitionSet, and an error, if there is any.
func (c *FakePartitionSets) Create(ctx context.Context, partitionSet *v1alpha1.PartitionSet, opts v1.CreateOptions) (result *v1alpha1.PartitionSet, err error) {
	emptyResult := &v1alpha1.PartitionSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(partitionsetsResource, partitionSet, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PartitionSet), err
}

// Update takes the representation of a partitionSet and updates it. Returns the server's representation of the partitionSet, and an error, if there is any.
func (c *FakePartitionSets) Update(ctx context.Context, partitionSet *v1alpha1.PartitionSet, opts v1.UpdateOptions) (result *v1alpha1.PartitionSet, err error) {
	emptyResult := &v1alpha1.PartitionSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(partitionsetsResource, partitionSet, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PartitionSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePartitionSets) UpdateStatus(ctx context.Context, partitionSet *v1alpha1.PartitionSet, opts v1.UpdateOptions) (result *v1alpha1.PartitionSet, err error) {
	emptyResult := &v1alpha1.PartitionSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(partitionsetsResource, "status", partitionSet, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PartitionSet), err
}

// Delete takes name of the partitionSet and deletes it. Returns an error if one occurs.
func (c *FakePartitionSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(partitionsetsResource, name, opts), &v1alpha1.PartitionSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePartitionSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(partitionsetsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PartitionSetList{})
	return err
}

// Patch applies the patch and returns the patched partitionSet.
func (c *FakePartitionSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PartitionSet, err error) {
	emptyResult := &v1alpha1.PartitionSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(partitionsetsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PartitionSet), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied partitionSet.
func (c *FakePartitionSets) Apply(ctx context.Context, partitionSet *topologyv1alpha1.PartitionSetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PartitionSet, err error) {
	if partitionSet == nil {
		return nil, fmt.Errorf("partitionSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(partitionSet)
	if err != nil {
		return nil, err
	}
	name := partitionSet.Name
	if name == nil {
		return nil, fmt.Errorf("partitionSet.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.PartitionSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(partitionsetsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PartitionSet), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePartitionSets) ApplyStatus(ctx context.Context, partitionSet *topologyv1alpha1.PartitionSetApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.PartitionSet, err error) {
	if partitionSet == nil {
		return nil, fmt.Errorf("partitionSet provided to Apply must not be nil")
	}
	data, err := json.Marshal(partitionSet)
	if err != nil {
		return nil, err
	}
	name := partitionSet.Name
	if name == nil {
		return nil, fmt.Errorf("partitionSet.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.PartitionSet{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(partitionsetsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PartitionSet), err
}
