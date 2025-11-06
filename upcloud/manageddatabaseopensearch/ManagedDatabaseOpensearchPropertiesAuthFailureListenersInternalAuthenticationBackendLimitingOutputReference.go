// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/manageddatabaseopensearch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference interface {
	cdktf.ComplexObject
	AllowedTries() *float64
	SetAllowedTries(val *float64)
	AllowedTriesInput() *float64
	AuthenticationBackend() *string
	SetAuthenticationBackend(val *string)
	AuthenticationBackendInput() *string
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
	InternalValue() *ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimiting
	SetInternalValue(val *ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimiting)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAllowedTries()
	ResetAuthenticationBackend()
	ResetBlockExpirySeconds()
	ResetMaxBlockedClients()
	ResetMaxTrackedClients()
	ResetTimeWindowSeconds()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference
type jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) AllowedTries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allowedTries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) AllowedTriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allowedTriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) AuthenticationBackend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationBackend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) AuthenticationBackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationBackendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) BlockExpirySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockExpirySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) BlockExpirySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockExpirySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) InternalValue() *ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimiting {
	var returns *ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimiting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) MaxBlockedClients() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBlockedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) MaxBlockedClientsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBlockedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) MaxTrackedClients() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTrackedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) MaxTrackedClientsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTrackedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) TimeWindowSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) TimeWindowSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference_Override(m ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetAllowedTries(val *float64) {
	if err := j.validateSetAllowedTriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedTries",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetAuthenticationBackend(val *string) {
	if err := j.validateSetAuthenticationBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationBackend",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetBlockExpirySeconds(val *float64) {
	if err := j.validateSetBlockExpirySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockExpirySeconds",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetInternalValue(val *ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimiting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetMaxBlockedClients(val *float64) {
	if err := j.validateSetMaxBlockedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBlockedClients",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetMaxTrackedClients(val *float64) {
	if err := j.validateSetMaxTrackedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTrackedClients",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetTimeWindowSeconds(val *float64) {
	if err := j.validateSetTimeWindowSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeWindowSeconds",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) ResetAllowedTries() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowedTries",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) ResetAuthenticationBackend() {
	_jsii_.InvokeVoid(
		m,
		"resetAuthenticationBackend",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) ResetBlockExpirySeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetBlockExpirySeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) ResetMaxBlockedClients() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxBlockedClients",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) ResetMaxTrackedClients() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxTrackedClients",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) ResetTimeWindowSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeWindowSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		m,
		"resetType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimitingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

