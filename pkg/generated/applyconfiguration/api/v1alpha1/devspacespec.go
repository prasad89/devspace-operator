// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// DevSpaceSpecApplyConfiguration represents a declarative configuration of the DevSpaceSpec type for use
// with apply.
type DevSpaceSpecApplyConfiguration struct {
	Name     *string `json:"name,omitempty"`
	Owner    *string `json:"owner,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
}

// DevSpaceSpecApplyConfiguration constructs a declarative configuration of the DevSpaceSpec type for use with
// apply.
func DevSpaceSpec() *DevSpaceSpecApplyConfiguration {
	return &DevSpaceSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *DevSpaceSpecApplyConfiguration) WithName(value string) *DevSpaceSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithOwner sets the Owner field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Owner field is set to the value of the last call.
func (b *DevSpaceSpecApplyConfiguration) WithOwner(value string) *DevSpaceSpecApplyConfiguration {
	b.Owner = &value
	return b
}

// WithHostname sets the Hostname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hostname field is set to the value of the last call.
func (b *DevSpaceSpecApplyConfiguration) WithHostname(value string) *DevSpaceSpecApplyConfiguration {
	b.Hostname = &value
	return b
}
