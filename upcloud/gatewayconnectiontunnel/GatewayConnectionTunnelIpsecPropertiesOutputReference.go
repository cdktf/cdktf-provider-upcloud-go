// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gatewayconnectiontunnel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/gatewayconnectiontunnel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GatewayConnectionTunnelIpsecPropertiesOutputReference interface {
	cdktf.ComplexObject
	ChildRekeyTime() *float64
	SetChildRekeyTime(val *float64)
	ChildRekeyTimeInput() *float64
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
	DpdDelay() *float64
	SetDpdDelay(val *float64)
	DpdDelayInput() *float64
	DpdTimeout() *float64
	SetDpdTimeout(val *float64)
	DpdTimeoutInput() *float64
	// Experimental.
	Fqn() *string
	IkeLifetime() *float64
	SetIkeLifetime(val *float64)
	IkeLifetimeInput() *float64
	InternalValue() *GatewayConnectionTunnelIpsecProperties
	SetInternalValue(val *GatewayConnectionTunnelIpsecProperties)
	Phase1Algorithms() *[]*string
	SetPhase1Algorithms(val *[]*string)
	Phase1AlgorithmsInput() *[]*string
	Phase1DhGroupNumbers() *[]*float64
	SetPhase1DhGroupNumbers(val *[]*float64)
	Phase1DhGroupNumbersInput() *[]*float64
	Phase1IntegrityAlgorithms() *[]*string
	SetPhase1IntegrityAlgorithms(val *[]*string)
	Phase1IntegrityAlgorithmsInput() *[]*string
	Phase2Algorithms() *[]*string
	SetPhase2Algorithms(val *[]*string)
	Phase2AlgorithmsInput() *[]*string
	Phase2DhGroupNumbers() *[]*float64
	SetPhase2DhGroupNumbers(val *[]*float64)
	Phase2DhGroupNumbersInput() *[]*float64
	Phase2IntegrityAlgorithms() *[]*string
	SetPhase2IntegrityAlgorithms(val *[]*string)
	Phase2IntegrityAlgorithmsInput() *[]*string
	RekeyTime() *float64
	SetRekeyTime(val *float64)
	RekeyTimeInput() *float64
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
	ResetChildRekeyTime()
	ResetDpdDelay()
	ResetDpdTimeout()
	ResetIkeLifetime()
	ResetPhase1Algorithms()
	ResetPhase1DhGroupNumbers()
	ResetPhase1IntegrityAlgorithms()
	ResetPhase2Algorithms()
	ResetPhase2DhGroupNumbers()
	ResetPhase2IntegrityAlgorithms()
	ResetRekeyTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GatewayConnectionTunnelIpsecPropertiesOutputReference
type jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ChildRekeyTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childRekeyTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ChildRekeyTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"childRekeyTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) DpdDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) DpdDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) DpdTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) DpdTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) IkeLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ikeLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) IkeLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ikeLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) InternalValue() *GatewayConnectionTunnelIpsecProperties {
	var returns *GatewayConnectionTunnelIpsecProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase1Algorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phase1Algorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase1AlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phase1AlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase1DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"phase1DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase1DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"phase1DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase1IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phase1IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase1IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phase1IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase2Algorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phase2Algorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase2AlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phase2AlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase2DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"phase2DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase2DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"phase2DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase2IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phase2IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Phase2IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phase2IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) RekeyTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rekeyTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) RekeyTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rekeyTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGatewayConnectionTunnelIpsecPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GatewayConnectionTunnelIpsecPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGatewayConnectionTunnelIpsecPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.gatewayConnectionTunnel.GatewayConnectionTunnelIpsecPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGatewayConnectionTunnelIpsecPropertiesOutputReference_Override(g GatewayConnectionTunnelIpsecPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.gatewayConnectionTunnel.GatewayConnectionTunnelIpsecPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetChildRekeyTime(val *float64) {
	if err := j.validateSetChildRekeyTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childRekeyTime",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetDpdDelay(val *float64) {
	if err := j.validateSetDpdDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dpdDelay",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetDpdTimeout(val *float64) {
	if err := j.validateSetDpdTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dpdTimeout",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetIkeLifetime(val *float64) {
	if err := j.validateSetIkeLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeLifetime",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetInternalValue(val *GatewayConnectionTunnelIpsecProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetPhase1Algorithms(val *[]*string) {
	if err := j.validateSetPhase1AlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phase1Algorithms",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetPhase1DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetPhase1DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phase1DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetPhase1IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetPhase1IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phase1IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetPhase2Algorithms(val *[]*string) {
	if err := j.validateSetPhase2AlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phase2Algorithms",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetPhase2DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetPhase2DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phase2DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetPhase2IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetPhase2IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phase2IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetRekeyTime(val *float64) {
	if err := j.validateSetRekeyTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rekeyTime",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ResetChildRekeyTime() {
	_jsii_.InvokeVoid(
		g,
		"resetChildRekeyTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ResetDpdDelay() {
	_jsii_.InvokeVoid(
		g,
		"resetDpdDelay",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ResetDpdTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetDpdTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ResetIkeLifetime() {
	_jsii_.InvokeVoid(
		g,
		"resetIkeLifetime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ResetPhase1Algorithms() {
	_jsii_.InvokeVoid(
		g,
		"resetPhase1Algorithms",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ResetPhase1DhGroupNumbers() {
	_jsii_.InvokeVoid(
		g,
		"resetPhase1DhGroupNumbers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ResetPhase1IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		g,
		"resetPhase1IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ResetPhase2Algorithms() {
	_jsii_.InvokeVoid(
		g,
		"resetPhase2Algorithms",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ResetPhase2DhGroupNumbers() {
	_jsii_.InvokeVoid(
		g,
		"resetPhase2DhGroupNumbers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ResetPhase2IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		g,
		"resetPhase2IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ResetRekeyTime() {
	_jsii_.InvokeVoid(
		g,
		"resetRekeyTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

