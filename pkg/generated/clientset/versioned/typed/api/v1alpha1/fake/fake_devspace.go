/*
Copyright 2025.

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
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/prasad89/devspace-operator/api/v1alpha1"
	apiv1alpha1 "github.com/prasad89/devspace-operator/pkg/generated/applyconfiguration/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDevSpaces implements DevSpaceInterface
type FakeDevSpaces struct {
	Fake *FakeApiV1alpha1
}

var devspacesResource = v1alpha1.SchemeGroupVersion.WithResource("devspaces")

var devspacesKind = v1alpha1.SchemeGroupVersion.WithKind("DevSpace")

// Get takes name of the devSpace, and returns the corresponding devSpace object, and an error if there is any.
func (c *FakeDevSpaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DevSpace, err error) {
	emptyResult := &v1alpha1.DevSpace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(devspacesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DevSpace), err
}

// List takes label and field selectors, and returns the list of DevSpaces that match those selectors.
func (c *FakeDevSpaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DevSpaceList, err error) {
	emptyResult := &v1alpha1.DevSpaceList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(devspacesResource, devspacesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DevSpaceList{ListMeta: obj.(*v1alpha1.DevSpaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.DevSpaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devSpaces.
func (c *FakeDevSpaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(devspacesResource, opts))
}

// Create takes the representation of a devSpace and creates it.  Returns the server's representation of the devSpace, and an error, if there is any.
func (c *FakeDevSpaces) Create(ctx context.Context, devSpace *v1alpha1.DevSpace, opts v1.CreateOptions) (result *v1alpha1.DevSpace, err error) {
	emptyResult := &v1alpha1.DevSpace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(devspacesResource, devSpace, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DevSpace), err
}

// Update takes the representation of a devSpace and updates it. Returns the server's representation of the devSpace, and an error, if there is any.
func (c *FakeDevSpaces) Update(ctx context.Context, devSpace *v1alpha1.DevSpace, opts v1.UpdateOptions) (result *v1alpha1.DevSpace, err error) {
	emptyResult := &v1alpha1.DevSpace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(devspacesResource, devSpace, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DevSpace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDevSpaces) UpdateStatus(ctx context.Context, devSpace *v1alpha1.DevSpace, opts v1.UpdateOptions) (result *v1alpha1.DevSpace, err error) {
	emptyResult := &v1alpha1.DevSpace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(devspacesResource, "status", devSpace, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DevSpace), err
}

// Delete takes name of the devSpace and deletes it. Returns an error if one occurs.
func (c *FakeDevSpaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(devspacesResource, name, opts), &v1alpha1.DevSpace{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevSpaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(devspacesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DevSpaceList{})
	return err
}

// Patch applies the patch and returns the patched devSpace.
func (c *FakeDevSpaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DevSpace, err error) {
	emptyResult := &v1alpha1.DevSpace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(devspacesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DevSpace), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied devSpace.
func (c *FakeDevSpaces) Apply(ctx context.Context, devSpace *apiv1alpha1.DevSpaceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DevSpace, err error) {
	if devSpace == nil {
		return nil, fmt.Errorf("devSpace provided to Apply must not be nil")
	}
	data, err := json.Marshal(devSpace)
	if err != nil {
		return nil, err
	}
	name := devSpace.Name
	if name == nil {
		return nil, fmt.Errorf("devSpace.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.DevSpace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(devspacesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DevSpace), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeDevSpaces) ApplyStatus(ctx context.Context, devSpace *apiv1alpha1.DevSpaceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DevSpace, err error) {
	if devSpace == nil {
		return nil, fmt.Errorf("devSpace provided to Apply must not be nil")
	}
	data, err := json.Marshal(devSpace)
	if err != nil {
		return nil, err
	}
	name := devSpace.Name
	if name == nil {
		return nil, fmt.Errorf("devSpace.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.DevSpace{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(devspacesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DevSpace), err
}
