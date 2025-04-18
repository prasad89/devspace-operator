// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/prasad89/devspace-operator/pkg/generated/clientset/versioned/typed/api/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeApiV1alpha1 struct {
	*testing.Fake
}

func (c *FakeApiV1alpha1) DevSpaces() v1alpha1.DevSpaceInterface {
	return newFakeDevSpaces(c)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeApiV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
