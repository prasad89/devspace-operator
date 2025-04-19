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

	v1alpha1 "github.com/prasad89/devspace-operator/api/v1alpha1"
	apiv1alpha1 "github.com/prasad89/devspace-operator/pkg/generated/applyconfiguration/api/v1alpha1"
	scheme "github.com/prasad89/devspace-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// DevSpacesGetter has a method to return a DevSpaceInterface.
// A group's client should implement this interface.
type DevSpacesGetter interface {
	DevSpaces(namespace string) DevSpaceInterface
}

// DevSpaceInterface has methods to work with DevSpace resources.
type DevSpaceInterface interface {
	Create(ctx context.Context, devSpace *v1alpha1.DevSpace, opts v1.CreateOptions) (*v1alpha1.DevSpace, error)
	Update(ctx context.Context, devSpace *v1alpha1.DevSpace, opts v1.UpdateOptions) (*v1alpha1.DevSpace, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, devSpace *v1alpha1.DevSpace, opts v1.UpdateOptions) (*v1alpha1.DevSpace, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DevSpace, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DevSpaceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DevSpace, err error)
	Apply(ctx context.Context, devSpace *apiv1alpha1.DevSpaceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DevSpace, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, devSpace *apiv1alpha1.DevSpaceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DevSpace, err error)
	DevSpaceExpansion
}

// devSpaces implements DevSpaceInterface
type devSpaces struct {
	*gentype.ClientWithListAndApply[*v1alpha1.DevSpace, *v1alpha1.DevSpaceList, *apiv1alpha1.DevSpaceApplyConfiguration]
}

// newDevSpaces returns a DevSpaces
func newDevSpaces(c *ApiV1alpha1Client, namespace string) *devSpaces {
	return &devSpaces{
		gentype.NewClientWithListAndApply[*v1alpha1.DevSpace, *v1alpha1.DevSpaceList, *apiv1alpha1.DevSpaceApplyConfiguration](
			"devspaces",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.DevSpace { return &v1alpha1.DevSpace{} },
			func() *v1alpha1.DevSpaceList { return &v1alpha1.DevSpaceList{} }),
	}
}
