// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package network

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/network/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference interface {
	cdktf.ComplexObject
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExcludeBySource() *[]*string
	SetExcludeBySource(val *[]*string)
	ExcludeBySourceInput() *[]*string
	FilterByDestination() *[]*string
	SetFilterByDestination(val *[]*string)
	FilterByDestinationInput() *[]*string
	FilterByRouteType() *[]*string
	SetFilterByRouteType(val *[]*string)
	FilterByRouteTypeInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetEnabled()
	ResetExcludeBySource()
	ResetFilterByDestination()
	ResetFilterByRouteType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference
type jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) ExcludeBySource() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeBySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) ExcludeBySourceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeBySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) FilterByDestination() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterByDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) FilterByDestinationInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterByDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) FilterByRouteType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterByRouteType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) FilterByRouteTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterByRouteTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.network.NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference_Override(n NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.network.NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference)SetExcludeBySource(val *[]*string) {
	if err := j.validateSetExcludeBySourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeBySource",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference)SetFilterByDestination(val *[]*string) {
	if err := j.validateSetFilterByDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterByDestination",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference)SetFilterByRouteType(val *[]*string) {
	if err := j.validateSetFilterByRouteTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterByRouteType",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) ResetExcludeBySource() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludeBySource",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) ResetFilterByDestination() {
	_jsii_.InvokeVoid(
		n,
		"resetFilterByDestination",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) ResetFilterByRouteType() {
	_jsii_.InvokeVoid(
		n,
		"resetFilterByRouteType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

