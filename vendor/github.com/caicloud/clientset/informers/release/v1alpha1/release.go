/*
Copyright 2017 caicloud authors. All rights reserved.
*/

// This file was automatically generated by informer-gen

package v1alpha1

import (
	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1alpha1 "github.com/caicloud/clientset/listers/release/v1alpha1"
	release_v1alpha1 "github.com/caicloud/clientset/pkg/apis/release/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	client_go_kubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ReleaseInformer provides access to a shared informer and lister for
// Releases.
type ReleaseInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReleaseLister
}

type releaseInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newReleaseInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.ReleaseV1alpha1().Releases(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.ReleaseV1alpha1().Releases(v1.NamespaceAll).Watch(options)
			},
		},
		&release_v1alpha1.Release{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *releaseInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&release_v1alpha1.Release{}, func(client client_go_kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
		// panic if client is not *kubernetes.Clientset
		return newReleaseInformer(client.(kubernetes.Interface), resyncPeriod)
	})
}

func (f *releaseInformer) Lister() v1alpha1.ReleaseLister {
	return v1alpha1.NewReleaseLister(f.Informer().GetIndexer())
}
