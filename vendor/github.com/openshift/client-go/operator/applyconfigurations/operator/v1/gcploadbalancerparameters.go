// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
)

// GCPLoadBalancerParametersApplyConfiguration represents an declarative configuration of the GCPLoadBalancerParameters type for use
// with apply.
type GCPLoadBalancerParametersApplyConfiguration struct {
	ClientAccess *v1.GCPClientAccess `json:"clientAccess,omitempty"`
}

// GCPLoadBalancerParametersApplyConfiguration constructs an declarative configuration of the GCPLoadBalancerParameters type for use with
// apply.
func GCPLoadBalancerParameters() *GCPLoadBalancerParametersApplyConfiguration {
	return &GCPLoadBalancerParametersApplyConfiguration{}
}

// WithClientAccess sets the ClientAccess field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientAccess field is set to the value of the last call.
func (b *GCPLoadBalancerParametersApplyConfiguration) WithClientAccess(value v1.GCPClientAccess) *GCPLoadBalancerParametersApplyConfiguration {
	b.ClientAccess = &value
	return b
}
