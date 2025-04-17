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

package v1alpha1

import (
	"context"
	scheme "pkg/client/clientset/versioned/scheme"
	"time"

	v1alpha1 "github.com/prasad89/devspace-operator/pkg/apis/platform.devspace.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DevSpaceSpecsGetter has a method to return a DevSpaceSpecInterface.
// A group's client should implement this interface.
type DevSpaceSpecsGetter interface {
	DevSpaceSpecs(namespace string) DevSpaceSpecInterface
}

// DevSpaceSpecInterface has methods to work with DevSpaceSpec resources.
type DevSpaceSpecInterface interface {
	Create(ctx context.Context, devSpaceSpec *v1alpha1.DevSpaceSpec, opts v1.CreateOptions) (*v1alpha1.DevSpaceSpec, error)
	Update(ctx context.Context, devSpaceSpec *v1alpha1.DevSpaceSpec, opts v1.UpdateOptions) (*v1alpha1.DevSpaceSpec, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DevSpaceSpec, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DevSpaceSpecList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DevSpaceSpec, err error)
	DevSpaceSpecExpansion
}

// devSpaceSpecs implements DevSpaceSpecInterface
type devSpaceSpecs struct {
	client rest.Interface
	ns     string
}

// newDevSpaceSpecs returns a DevSpaceSpecs
func newDevSpaceSpecs(c *PlatformV1alpha1Client, namespace string) *devSpaceSpecs {
	return &devSpaceSpecs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the devSpaceSpec, and returns the corresponding devSpaceSpec object, and an error if there is any.
func (c *devSpaceSpecs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DevSpaceSpec, err error) {
	result = &v1alpha1.DevSpaceSpec{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devspacespecs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DevSpaceSpecs that match those selectors.
func (c *devSpaceSpecs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DevSpaceSpecList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DevSpaceSpecList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devspacespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested devSpaceSpecs.
func (c *devSpaceSpecs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("devspacespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a devSpaceSpec and creates it.  Returns the server's representation of the devSpaceSpec, and an error, if there is any.
func (c *devSpaceSpecs) Create(ctx context.Context, devSpaceSpec *v1alpha1.DevSpaceSpec, opts v1.CreateOptions) (result *v1alpha1.DevSpaceSpec, err error) {
	result = &v1alpha1.DevSpaceSpec{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("devspacespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(devSpaceSpec).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a devSpaceSpec and updates it. Returns the server's representation of the devSpaceSpec, and an error, if there is any.
func (c *devSpaceSpecs) Update(ctx context.Context, devSpaceSpec *v1alpha1.DevSpaceSpec, opts v1.UpdateOptions) (result *v1alpha1.DevSpaceSpec, err error) {
	result = &v1alpha1.DevSpaceSpec{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devspacespecs").
		Name(devSpaceSpec.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(devSpaceSpec).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the devSpaceSpec and deletes it. Returns an error if one occurs.
func (c *devSpaceSpecs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devspacespecs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *devSpaceSpecs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devspacespecs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched devSpaceSpec.
func (c *devSpaceSpecs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DevSpaceSpec, err error) {
	result = &v1alpha1.DevSpaceSpec{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("devspacespecs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
