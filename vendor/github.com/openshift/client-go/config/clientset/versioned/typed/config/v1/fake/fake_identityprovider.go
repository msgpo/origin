// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	config_v1 "github.com/openshift/api/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIdentityProviders implements IdentityProviderInterface
type FakeIdentityProviders struct {
	Fake *FakeConfigV1
}

var identityprovidersResource = schema.GroupVersionResource{Group: "config.openshift.io", Version: "v1", Resource: "identityproviders"}

var identityprovidersKind = schema.GroupVersionKind{Group: "config.openshift.io", Version: "v1", Kind: "IdentityProvider"}

// Get takes name of the identityProvider, and returns the corresponding identityProvider object, and an error if there is any.
func (c *FakeIdentityProviders) Get(name string, options v1.GetOptions) (result *config_v1.IdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(identityprovidersResource, name), &config_v1.IdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.IdentityProvider), err
}

// List takes label and field selectors, and returns the list of IdentityProviders that match those selectors.
func (c *FakeIdentityProviders) List(opts v1.ListOptions) (result *config_v1.IdentityProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(identityprovidersResource, identityprovidersKind, opts), &config_v1.IdentityProviderList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &config_v1.IdentityProviderList{ListMeta: obj.(*config_v1.IdentityProviderList).ListMeta}
	for _, item := range obj.(*config_v1.IdentityProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identityProviders.
func (c *FakeIdentityProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(identityprovidersResource, opts))
}

// Create takes the representation of a identityProvider and creates it.  Returns the server's representation of the identityProvider, and an error, if there is any.
func (c *FakeIdentityProviders) Create(identityProvider *config_v1.IdentityProvider) (result *config_v1.IdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(identityprovidersResource, identityProvider), &config_v1.IdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.IdentityProvider), err
}

// Update takes the representation of a identityProvider and updates it. Returns the server's representation of the identityProvider, and an error, if there is any.
func (c *FakeIdentityProviders) Update(identityProvider *config_v1.IdentityProvider) (result *config_v1.IdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(identityprovidersResource, identityProvider), &config_v1.IdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.IdentityProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIdentityProviders) UpdateStatus(identityProvider *config_v1.IdentityProvider) (*config_v1.IdentityProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(identityprovidersResource, "status", identityProvider), &config_v1.IdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.IdentityProvider), err
}

// Delete takes name of the identityProvider and deletes it. Returns an error if one occurs.
func (c *FakeIdentityProviders) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(identityprovidersResource, name), &config_v1.IdentityProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentityProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(identityprovidersResource, listOptions)

	_, err := c.Fake.Invokes(action, &config_v1.IdentityProviderList{})
	return err
}

// Patch applies the patch and returns the patched identityProvider.
func (c *FakeIdentityProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *config_v1.IdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(identityprovidersResource, name, data, subresources...), &config_v1.IdentityProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*config_v1.IdentityProvider), err
}
