/*
Copyright 2023 The Kubernetes Authors.

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

// This package imports things required by build scripts, to force `go mod` to see them as dependencies

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	networktopologyv1alpha1 "github.com/diktyo-io/networktopology-api/pkg/apis/networktopology/v1alpha1"
	versioned "github.com/diktyo-io/networktopology-api/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/diktyo-io/networktopology-api/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/diktyo-io/networktopology-api/pkg/generated/listers/networktopology/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NetworkTopologyInformer provides access to a shared informer and lister for
// NetworkTopologies.
type NetworkTopologyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NetworkTopologyLister
}

type networkTopologyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNetworkTopologyInformer constructs a new informer for NetworkTopology type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkTopologyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNetworkTopologyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNetworkTopologyInformer constructs a new informer for NetworkTopology type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkTopologyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworktopologyV1alpha1().NetworkTopologies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworktopologyV1alpha1().NetworkTopologies(namespace).Watch(context.TODO(), options)
			},
		},
		&networktopologyv1alpha1.NetworkTopology{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkTopologyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNetworkTopologyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *networkTopologyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networktopologyv1alpha1.NetworkTopology{}, f.defaultInformer)
}

func (f *networkTopologyInformer) Lister() v1alpha1.NetworkTopologyLister {
	return v1alpha1.NewNetworkTopologyLister(f.Informer().GetIndexer())
}
