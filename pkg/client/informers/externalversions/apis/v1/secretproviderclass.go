/*
Copyright 2020 The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	apisv1 "sigs.k8s.io/secrets-store-csi-driver/apis/v1"
	versioned "sigs.k8s.io/secrets-store-csi-driver/pkg/client/clientset/versioned"
	internalinterfaces "sigs.k8s.io/secrets-store-csi-driver/pkg/client/informers/externalversions/internalinterfaces"
	v1 "sigs.k8s.io/secrets-store-csi-driver/pkg/client/listers/apis/v1"
)

// SecretProviderClassInformer provides access to a shared informer and lister for
// SecretProviderClasses.
type SecretProviderClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SecretProviderClassLister
}

type secretProviderClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSecretProviderClassInformer constructs a new informer for SecretProviderClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSecretProviderClassInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSecretProviderClassInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSecretProviderClassInformer constructs a new informer for SecretProviderClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSecretProviderClassInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretsstoreV1().SecretProviderClasses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretsstoreV1().SecretProviderClasses(namespace).Watch(context.TODO(), options)
			},
		},
		&apisv1.SecretProviderClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *secretProviderClassInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSecretProviderClassInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *secretProviderClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisv1.SecretProviderClass{}, f.defaultInformer)
}

func (f *secretProviderClassInformer) Lister() v1.SecretProviderClassLister {
	return v1.NewSecretProviderClassLister(f.Informer().GetIndexer())
}
