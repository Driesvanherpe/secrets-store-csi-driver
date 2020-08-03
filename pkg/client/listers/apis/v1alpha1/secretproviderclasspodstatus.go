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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/secrets-store-csi-driver/apis/v1alpha1"
)

// SecretProviderClassPodStatusLister helps list SecretProviderClassPodStatuses.
type SecretProviderClassPodStatusLister interface {
	// List lists all SecretProviderClassPodStatuses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SecretProviderClassPodStatus, err error)
	// SecretProviderClassPodStatuses returns an object that can list and get SecretProviderClassPodStatuses.
	SecretProviderClassPodStatuses(namespace string) SecretProviderClassPodStatusNamespaceLister
	SecretProviderClassPodStatusListerExpansion
}

// secretProviderClassPodStatusLister implements the SecretProviderClassPodStatusLister interface.
type secretProviderClassPodStatusLister struct {
	indexer cache.Indexer
}

// NewSecretProviderClassPodStatusLister returns a new SecretProviderClassPodStatusLister.
func NewSecretProviderClassPodStatusLister(indexer cache.Indexer) SecretProviderClassPodStatusLister {
	return &secretProviderClassPodStatusLister{indexer: indexer}
}

// List lists all SecretProviderClassPodStatuses in the indexer.
func (s *secretProviderClassPodStatusLister) List(selector labels.Selector) (ret []*v1alpha1.SecretProviderClassPodStatus, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretProviderClassPodStatus))
	})
	return ret, err
}

// SecretProviderClassPodStatuses returns an object that can list and get SecretProviderClassPodStatuses.
func (s *secretProviderClassPodStatusLister) SecretProviderClassPodStatuses(namespace string) SecretProviderClassPodStatusNamespaceLister {
	return secretProviderClassPodStatusNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecretProviderClassPodStatusNamespaceLister helps list and get SecretProviderClassPodStatuses.
type SecretProviderClassPodStatusNamespaceLister interface {
	// List lists all SecretProviderClassPodStatuses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SecretProviderClassPodStatus, err error)
	// Get retrieves the SecretProviderClassPodStatus from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SecretProviderClassPodStatus, error)
	SecretProviderClassPodStatusNamespaceListerExpansion
}

// secretProviderClassPodStatusNamespaceLister implements the SecretProviderClassPodStatusNamespaceLister
// interface.
type secretProviderClassPodStatusNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecretProviderClassPodStatuses in the indexer for a given namespace.
func (s secretProviderClassPodStatusNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecretProviderClassPodStatus, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretProviderClassPodStatus))
	})
	return ret, err
}

// Get retrieves the SecretProviderClassPodStatus from the indexer for a given namespace and name.
func (s secretProviderClassPodStatusNamespaceLister) Get(name string) (*v1alpha1.SecretProviderClassPodStatus, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("secretproviderclasspodstatus"), name)
	}
	return obj.(*v1alpha1.SecretProviderClassPodStatus), nil
}
