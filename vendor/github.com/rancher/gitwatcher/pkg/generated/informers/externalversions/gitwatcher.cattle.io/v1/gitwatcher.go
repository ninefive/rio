/*
Copyright The Kubernetes Authors.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	time "time"

	gitwatchercattleiov1 "github.com/rancher/gitwatcher/pkg/apis/gitwatcher.cattle.io/v1"
	versioned "github.com/rancher/gitwatcher/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/rancher/gitwatcher/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/rancher/gitwatcher/pkg/generated/listers/gitwatcher.cattle.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GitWatcherInformer provides access to a shared informer and lister for
// GitWatchers.
type GitWatcherInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.GitWatcherLister
}

type gitWatcherInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGitWatcherInformer constructs a new informer for GitWatcher type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGitWatcherInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGitWatcherInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGitWatcherInformer constructs a new informer for GitWatcher type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGitWatcherInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GitwatcherV1().GitWatchers(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GitwatcherV1().GitWatchers(namespace).Watch(options)
			},
		},
		&gitwatchercattleiov1.GitWatcher{},
		resyncPeriod,
		indexers,
	)
}

func (f *gitWatcherInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGitWatcherInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *gitWatcherInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&gitwatchercattleiov1.GitWatcher{}, f.defaultInformer)
}

func (f *gitWatcherInformer) Lister() v1.GitWatcherLister {
	return v1.NewGitWatcherLister(f.Informer().GetIndexer())
}
