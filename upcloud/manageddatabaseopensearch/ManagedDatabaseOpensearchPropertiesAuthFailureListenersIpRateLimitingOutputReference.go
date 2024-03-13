// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v14/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v14/manageddatabaseopensearch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference interface {
	cdktf.ComplexObject
	AllowedTries() *float64
	SetAllowedTries(val *float64)
	AllowedTriesInput() *float64
	BlockExpirySeconds() *float64
	SetBlockExpirySeconds(val *float64)
	BlockExpirySecondsInput() *float64
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
	// Experimental.
	Fqn() *string
	InternalValue() *ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimiting
	SetInternalValue(val *ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimiting)
	MaxBlockedClients() *float64
	SetMaxBlockedClients(val *float64)
	MaxBlockedClientsInput() *float64
	MaxTrackedClients() *float64
	SetMaxTrackedClients(val *float64)
	MaxTrackedClientsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeWindowSeconds() *float64
	SetTimeWindowSeconds(val *float64)
	TimeWindowSecondsInput() *float64
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetAllowedTries()
	ResetBlockExpirySeconds()
	ResetMaxBlockedClients()
	ResetMaxTrackedClients()
	ResetTimeWindowSeconds()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference
type jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) AllowedTries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allowedTries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) AllowedTriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allowedTriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) BlockExpirySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockExpirySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) BlockExpirySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockExpirySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) InternalValue() *ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimiting {
	var returns *ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimiting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) MaxBlockedClients() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBlockedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) MaxBlockedClientsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBlockedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) MaxTrackedClients() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTrackedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) MaxTrackedClientsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTrackedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) TimeWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) TimeWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference_Override(m ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference)SetAllowedTries(val *float64) {
	if err := j.validateSetAllowedTriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedTries",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference)SetBlockExpirySeconds(val *float64) {
	if err := j.validateSetBlockExpirySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockExpirySeconds",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference)SetInternalValue(val *ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimiting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference)SetMaxBlockedClients(val *float64) {
	if err := j.validateSetMaxBlockedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBlockedClients",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference)SetMaxTrackedClients(val *float64) {
	if err := j.validateSetMaxTrackedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTrackedClients",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference)SetTimeWindowSeconds(val *float64) {
	if err := j.validateSetTimeWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) ResetAllowedTries() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowedTries",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) ResetBlockExpirySeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetBlockExpirySeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) ResetMaxBlockedClients() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBlockedClients",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) ResetMaxTrackedClients() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxTrackedClients",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) ResetTimeWindowSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeWindowSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		m,
		"resetType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimitingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

