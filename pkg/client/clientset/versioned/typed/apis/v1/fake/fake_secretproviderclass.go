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

// FakeSecretProviderClasses implements SecretProviderClassInterface
type FakeSecretProviderClasses struct {
	Fake *FakeSecretsstoreV1
	ns   string
}

var secretproviderclassesResource = schema.GroupVersionResource{Group: "secrets-store.csi.x-k8s.io", Version: "v1", Resource: "secretproviderclasses"}

var secretproviderclassesKind = schema.GroupVersionKind{Group: "secrets-store.csi.x-k8s.io", Version: "v1", Kind: "SecretProviderClass"}

// Get takes name of the secretProviderClass, and returns the corresponding secretProviderClass object, and an error if there is any.
func (c *FakeSecretProviderClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *apisv1.SecretProviderClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secretproviderclassesResource, c.ns, name), &apisv1.SecretProviderClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.SecretProviderClass), err
}

// List takes label and field selectors, and returns the list of SecretProviderClasses that match those selectors.
func (c *FakeSecretProviderClasses) List(ctx context.Context, opts v1.ListOptions) (result *apisv1.SecretProviderClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secretproviderclassesResource, secretproviderclassesKind, c.ns, opts), &apisv1.SecretProviderClassList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apisv1.SecretProviderClassList{ListMeta: obj.(*apisv1.SecretProviderClassList).ListMeta}
	for _, item := range obj.(*apisv1.SecretProviderClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secretProviderClasses.
func (c *FakeSecretProviderClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secretproviderclassesResource, c.ns, opts))

}

// Create takes the representation of a secretProviderClass and creates it.  Returns the server's representation of the secretProviderClass, and an error, if there is any.
func (c *FakeSecretProviderClasses) Create(ctx context.Context, secretProviderClass *apisv1.SecretProviderClass, opts v1.CreateOptions) (result *apisv1.SecretProviderClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secretproviderclassesResource, c.ns, secretProviderClass), &apisv1.SecretProviderClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.SecretProviderClass), err
}

// Update takes the representation of a secretProviderClass and updates it. Returns the server's representation of the secretProviderClass, and an error, if there is any.
func (c *FakeSecretProviderClasses) Update(ctx context.Context, secretProviderClass *apisv1.SecretProviderClass, opts v1.UpdateOptions) (result *apisv1.SecretProviderClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secretproviderclassesResource, c.ns, secretProviderClass), &apisv1.SecretProviderClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.SecretProviderClass), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecretProviderClasses) UpdateStatus(ctx context.Context, secretProviderClass *apisv1.SecretProviderClass, opts v1.UpdateOptions) (*apisv1.SecretProviderClass, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(secretproviderclassesResource, "status", c.ns, secretProviderClass), &apisv1.SecretProviderClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.SecretProviderClass), err
}

// Delete takes name of the secretProviderClass and deletes it. Returns an error if one occurs.
func (c *FakeSecretProviderClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(secretproviderclassesResource, c.ns, name), &apisv1.SecretProviderClass{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecretProviderClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secretproviderclassesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apisv1.SecretProviderClassList{})
	return err
}

// Patch applies the patch and returns the patched secretProviderClass.
func (c *FakeSecretProviderClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apisv1.SecretProviderClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secretproviderclassesResource, c.ns, name, pt, data, subresources...), &apisv1.SecretProviderClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apisv1.SecretProviderClass), err
}
