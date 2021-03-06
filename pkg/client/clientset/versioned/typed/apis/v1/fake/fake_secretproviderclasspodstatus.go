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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	apisv1 "sigs.k8s.io/secrets-store-csi-driver/apis/v1"
)

// FakeSecretProviderClassPodStatuses implements SecretProviderClassPodStatusInterface
type FakeSecretProviderClassPodStatuses struct {
	Fake *FakeSecretsstoreV1
	ns   string
}

var secretproviderclasspodstatusesResource = schema.GroupVersionResource{Group: "secrets-store.csi.x-k8s.io", Version: "v1", Resource: "secretproviderclasspodstatuses"}

var secretproviderclasspodstatusesKind = schema.GroupVersionKind{Group: "secrets-store.csi.x-k8s.io", Version: "v1", Kind: "SecretProviderClassPodStatus"}

// Get takes name of the secretProviderClassPodStatus, and returns the corresponding secretProviderClassPodStatus object, and an error if there is any.
func (c *FakeSecretProviderClassPodStatuses) Get(ctx context.Context, name string, options v1.GetOptions) (result *apisv1.SecretProviderClassPodStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secretproviderclasspodstatusesResource, c.ns, name), &apisv1.SecretProviderClassPodStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.SecretProviderClassPodStatus), err
}

// List takes label and field selectors, and returns the list of SecretProviderClassPodStatuses that match those selectors.
func (c *FakeSecretProviderClassPodStatuses) List(ctx context.Context, opts v1.ListOptions) (result *apisv1.SecretProviderClassPodStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secretproviderclasspodstatusesResource, secretproviderclasspodstatusesKind, c.ns, opts), &apisv1.SecretProviderClassPodStatusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apisv1.SecretProviderClassPodStatusList{ListMeta: obj.(*apisv1.SecretProviderClassPodStatusList).ListMeta}
	for _, item := range obj.(*apisv1.SecretProviderClassPodStatusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secretProviderClassPodStatuses.
func (c *FakeSecretProviderClassPodStatuses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secretproviderclasspodstatusesResource, c.ns, opts))

}

// Create takes the representation of a secretProviderClassPodStatus and creates it.  Returns the server's representation of the secretProviderClassPodStatus, and an error, if there is any.
func (c *FakeSecretProviderClassPodStatuses) Create(ctx context.Context, secretProviderClassPodStatus *apisv1.SecretProviderClassPodStatus, opts v1.CreateOptions) (result *apisv1.SecretProviderClassPodStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secretproviderclasspodstatusesResource, c.ns, secretProviderClassPodStatus), &apisv1.SecretProviderClassPodStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.SecretProviderClassPodStatus), err
}

// Update takes the representation of a secretProviderClassPodStatus and updates it. Returns the server's representation of the secretProviderClassPodStatus, and an error, if there is any.
func (c *FakeSecretProviderClassPodStatuses) Update(ctx context.Context, secretProviderClassPodStatus *apisv1.SecretProviderClassPodStatus, opts v1.UpdateOptions) (result *apisv1.SecretProviderClassPodStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secretproviderclasspodstatusesResource, c.ns, secretProviderClassPodStatus), &apisv1.SecretProviderClassPodStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.SecretProviderClassPodStatus), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecretProviderClassPodStatuses) UpdateStatus(ctx context.Context, secretProviderClassPodStatus *apisv1.SecretProviderClassPodStatus, opts v1.UpdateOptions) (*apisv1.SecretProviderClassPodStatus, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(secretproviderclasspodstatusesResource, "status", c.ns, secretProviderClassPodStatus), &apisv1.SecretProviderClassPodStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.SecretProviderClassPodStatus), err
}

// Delete takes name of the secretProviderClassPodStatus and deletes it. Returns an error if one occurs.
func (c *FakeSecretProviderClassPodStatuses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(secretproviderclasspodstatusesResource, c.ns, name), &apisv1.SecretProviderClassPodStatus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecretProviderClassPodStatuses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secretproviderclasspodstatusesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apisv1.SecretProviderClassPodStatusList{})
	return err
}

// Patch applies the patch and returns the patched secretProviderClassPodStatus.
func (c *FakeSecretProviderClassPodStatuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apisv1.SecretProviderClassPodStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secretproviderclasspodstatusesResource, c.ns, name, pt, data, subresources...), &apisv1.SecretProviderClassPodStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.SecretProviderClassPodStatus), err
}
