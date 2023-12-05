/*
Copyright The Karmada Authors.

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

	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/work/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResourceBindings implements ResourceBindingInterface
type FakeResourceBindings struct {
	Fake *FakeWorkV1alpha1
	ns   string
}

var resourcebindingsResource = v1alpha1.SchemeGroupVersion.WithResource("resourcebindings")

var resourcebindingsKind = v1alpha1.SchemeGroupVersion.WithKind("ResourceBinding")

// Get takes name of the resourceBinding, and returns the corresponding resourceBinding object, and an error if there is any.
func (c *FakeResourceBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcebindingsResource, c.ns, name), &v1alpha1.ResourceBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceBinding), err
}

// List takes label and field selectors, and returns the list of ResourceBindings that match those selectors.
func (c *FakeResourceBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResourceBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcebindingsResource, resourcebindingsKind, c.ns, opts), &v1alpha1.ResourceBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourceBindingList{ListMeta: obj.(*v1alpha1.ResourceBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourceBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceBindings.
func (c *FakeResourceBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcebindingsResource, c.ns, opts))

}

// Create takes the representation of a resourceBinding and creates it.  Returns the server's representation of the resourceBinding, and an error, if there is any.
func (c *FakeResourceBindings) Create(ctx context.Context, resourceBinding *v1alpha1.ResourceBinding, opts v1.CreateOptions) (result *v1alpha1.ResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcebindingsResource, c.ns, resourceBinding), &v1alpha1.ResourceBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceBinding), err
}

// Update takes the representation of a resourceBinding and updates it. Returns the server's representation of the resourceBinding, and an error, if there is any.
func (c *FakeResourceBindings) Update(ctx context.Context, resourceBinding *v1alpha1.ResourceBinding, opts v1.UpdateOptions) (result *v1alpha1.ResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcebindingsResource, c.ns, resourceBinding), &v1alpha1.ResourceBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceBindings) UpdateStatus(ctx context.Context, resourceBinding *v1alpha1.ResourceBinding, opts v1.UpdateOptions) (*v1alpha1.ResourceBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourcebindingsResource, "status", c.ns, resourceBinding), &v1alpha1.ResourceBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceBinding), err
}

// Delete takes name of the resourceBinding and deletes it. Returns an error if one occurs.
func (c *FakeResourceBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resourcebindingsResource, c.ns, name, opts), &v1alpha1.ResourceBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcebindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourceBindingList{})
	return err
}

// Patch applies the patch and returns the patched resourceBinding.
func (c *FakeResourceBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcebindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResourceBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceBinding), err
}
