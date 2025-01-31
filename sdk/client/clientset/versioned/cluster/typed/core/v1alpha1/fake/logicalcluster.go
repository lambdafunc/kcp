//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	corev1alpha1 "github.com/kcp-dev/kcp/sdk/apis/core/v1alpha1"
	applyconfigurationscorev1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/core/v1alpha1"
	corev1alpha1client "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/core/v1alpha1"
)

var logicalClustersResource = schema.GroupVersionResource{Group: "core.kcp.io", Version: "v1alpha1", Resource: "logicalclusters"}
var logicalClustersKind = schema.GroupVersionKind{Group: "core.kcp.io", Version: "v1alpha1", Kind: "LogicalCluster"}

type logicalClustersClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *logicalClustersClusterClient) Cluster(clusterPath logicalcluster.Path) corev1alpha1client.LogicalClusterInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &logicalClustersClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of LogicalClusters that match those selectors across all clusters.
func (c *logicalClustersClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1alpha1.LogicalClusterList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(logicalClustersResource, logicalClustersKind, logicalcluster.Wildcard, opts), &corev1alpha1.LogicalClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1alpha1.LogicalClusterList{ListMeta: obj.(*corev1alpha1.LogicalClusterList).ListMeta}
	for _, item := range obj.(*corev1alpha1.LogicalClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested LogicalClusters across all clusters.
func (c *logicalClustersClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(logicalClustersResource, logicalcluster.Wildcard, opts))
}

type logicalClustersClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *logicalClustersClient) Create(ctx context.Context, logicalCluster *corev1alpha1.LogicalCluster, opts metav1.CreateOptions) (*corev1alpha1.LogicalCluster, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(logicalClustersResource, c.ClusterPath, logicalCluster), &corev1alpha1.LogicalCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1alpha1.LogicalCluster), err
}

func (c *logicalClustersClient) Update(ctx context.Context, logicalCluster *corev1alpha1.LogicalCluster, opts metav1.UpdateOptions) (*corev1alpha1.LogicalCluster, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(logicalClustersResource, c.ClusterPath, logicalCluster), &corev1alpha1.LogicalCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1alpha1.LogicalCluster), err
}

func (c *logicalClustersClient) UpdateStatus(ctx context.Context, logicalCluster *corev1alpha1.LogicalCluster, opts metav1.UpdateOptions) (*corev1alpha1.LogicalCluster, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(logicalClustersResource, c.ClusterPath, "status", logicalCluster), &corev1alpha1.LogicalCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1alpha1.LogicalCluster), err
}

func (c *logicalClustersClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(logicalClustersResource, c.ClusterPath, name, opts), &corev1alpha1.LogicalCluster{})
	return err
}

func (c *logicalClustersClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(logicalClustersResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &corev1alpha1.LogicalClusterList{})
	return err
}

func (c *logicalClustersClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*corev1alpha1.LogicalCluster, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(logicalClustersResource, c.ClusterPath, name), &corev1alpha1.LogicalCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1alpha1.LogicalCluster), err
}

// List takes label and field selectors, and returns the list of LogicalClusters that match those selectors.
func (c *logicalClustersClient) List(ctx context.Context, opts metav1.ListOptions) (*corev1alpha1.LogicalClusterList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(logicalClustersResource, logicalClustersKind, c.ClusterPath, opts), &corev1alpha1.LogicalClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1alpha1.LogicalClusterList{ListMeta: obj.(*corev1alpha1.LogicalClusterList).ListMeta}
	for _, item := range obj.(*corev1alpha1.LogicalClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *logicalClustersClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(logicalClustersResource, c.ClusterPath, opts))
}

func (c *logicalClustersClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*corev1alpha1.LogicalCluster, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(logicalClustersResource, c.ClusterPath, name, pt, data, subresources...), &corev1alpha1.LogicalCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1alpha1.LogicalCluster), err
}

func (c *logicalClustersClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationscorev1alpha1.LogicalClusterApplyConfiguration, opts metav1.ApplyOptions) (*corev1alpha1.LogicalCluster, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(logicalClustersResource, c.ClusterPath, *name, types.ApplyPatchType, data), &corev1alpha1.LogicalCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1alpha1.LogicalCluster), err
}

func (c *logicalClustersClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationscorev1alpha1.LogicalClusterApplyConfiguration, opts metav1.ApplyOptions) (*corev1alpha1.LogicalCluster, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(logicalClustersResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &corev1alpha1.LogicalCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*corev1alpha1.LogicalCluster), err
}
