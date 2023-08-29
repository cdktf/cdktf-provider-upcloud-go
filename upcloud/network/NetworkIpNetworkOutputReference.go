// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package network

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v9/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v9/network/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkIpNetworkOutputReference interface {
	cdktf.ComplexObject
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dhcp() interface{}
	SetDhcp(val interface{})
	DhcpDefaultRoute() interface{}
	SetDhcpDefaultRoute(val interface{})
	DhcpDefaultRouteInput() interface{}
	DhcpDns() *[]*string
	SetDhcpDns(val *[]*string)
	DhcpDnsInput() *[]*string
	DhcpInput() interface{}
	Family() *string
	SetFamily(val *string)
	FamilyInput() *string
	// Experimental.
	Fqn() *string
	Gateway() *string
	SetGateway(val *string)
	GatewayInput() *string
	InternalValue() *NetworkIpNetwork
	SetInternalValue(val *NetworkIpNetwork)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDhcpDefaultRoute()
	ResetDhcpDns()
	ResetGateway()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkIpNetworkOutputReference
type jsiiProxy_NetworkIpNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) Dhcp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) DhcpDefaultRoute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpDefaultRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) DhcpDefaultRouteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpDefaultRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) DhcpDns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dhcpDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) DhcpDnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dhcpDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) DhcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dhcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) FamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) Gateway() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) GatewayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) InternalValue() *NetworkIpNetwork {
	var returns *NetworkIpNetwork
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkIpNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkIpNetworkOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkIpNetworkOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkIpNetworkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.network.NetworkIpNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkIpNetworkOutputReference_Override(n NetworkIpNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.network.NetworkIpNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference)SetAddress(val *string) {
	if err := j.validateSetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference)SetDhcp(val interface{}) {
	if err := j.validateSetDhcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcp",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference)SetDhcpDefaultRoute(val interface{}) {
	if err := j.validateSetDhcpDefaultRouteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcpDefaultRoute",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference)SetDhcpDns(val *[]*string) {
	if err := j.validateSetDhcpDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhcpDns",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference)SetFamily(val *string) {
	if err := j.validateSetFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference)SetGateway(val *string) {
	if err := j.validateSetGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gateway",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference)SetInternalValue(val *NetworkIpNetwork) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) ResetDhcpDefaultRoute() {
	_jsii_.InvokeVoid(
		n,
		"resetDhcpDefaultRoute",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) ResetDhcpDns() {
	_jsii_.InvokeVoid(
		n,
		"resetDhcpDns",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) ResetGateway() {
	_jsii_.InvokeVoid(
		n,
		"resetGateway",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

