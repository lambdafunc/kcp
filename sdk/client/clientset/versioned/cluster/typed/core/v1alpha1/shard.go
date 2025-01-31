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

package v1alpha1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	corev1alpha1 "github.com/kcp-dev/kcp/sdk/apis/core/v1alpha1"
	corev1alpha1client "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/core/v1alpha1"
)

// ShardsClusterGetter has a method to return a ShardClusterInterface.
// A group's cluster client should implement this interface.
type ShardsClusterGetter interface {
	Shards() ShardClusterInterface
}

// ShardClusterInterface can operate on Shards across all clusters,
// or scope down to one cluster and return a corev1alpha1client.ShardInterface.
type ShardClusterInterface interface {
	Cluster(logicalcluster.Path) corev1alpha1client.ShardInterface
	List(ctx context.Context, opts metav1.ListOptions) (*corev1alpha1.ShardList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type shardsClusterInterface struct {
	clientCache kcpclient.Cache[*corev1alpha1client.CoreV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *shardsClusterInterface) Cluster(clusterPath logicalcluster.Path) corev1alpha1client.ShardInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).Shards()
}

// List returns the entire collection of all Shards across all clusters.
func (c *shardsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1alpha1.ShardList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Shards().List(ctx, opts)
}

// Watch begins to watch all Shards across all clusters.
func (c *shardsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Shards().Watch(ctx, opts)
}
