// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/manageddatabaseopensearch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseOpensearchPropertiesOpenidOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	ConnectUrl() *string
	SetConnectUrl(val *string)
	ConnectUrlInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	Header() *string
	SetHeader(val *string)
	HeaderInput() *string
	InternalValue() *ManagedDatabaseOpensearchPropertiesOpenid
	SetInternalValue(val *ManagedDatabaseOpensearchPropertiesOpenid)
	JwtHeader() *string
	SetJwtHeader(val *string)
	JwtHeaderInput() *string
	JwtUrlParameter() *string
	SetJwtUrlParameter(val *string)
	JwtUrlParameterInput() *string
	RefreshRateLimitCount() *float64
	SetRefreshRateLimitCount(val *float64)
	RefreshRateLimitCountInput() *float64
	RefreshRateLimitTimeWindowMs() *float64
	SetRefreshRateLimitTimeWindowMs(val *float64)
	RefreshRateLimitTimeWindowMsInput() *float64
	RolesKey() *string
	SetRolesKey(val *string)
	RolesKeyInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	SubjectKey() *string
	SetSubjectKey(val *string)
	SubjectKeyInput() *string
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
	ResetClientId()
	ResetClientSecret()
	ResetConnectUrl()
	ResetEnabled()
	ResetHeader()
	ResetJwtHeader()
	ResetJwtUrlParameter()
	ResetRefreshRateLimitCount()
	ResetRefreshRateLimitTimeWindowMs()
	ResetRolesKey()
	ResetScope()
	ResetSubjectKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabaseOpensearchPropertiesOpenidOutputReference
type jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ConnectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ConnectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) Header() *string {
	var returns *string
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) HeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) InternalValue() *ManagedDatabaseOpensearchPropertiesOpenid {
	var returns *ManagedDatabaseOpensearchPropertiesOpenid
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) JwtHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) JwtHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) JwtUrlParameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtUrlParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) JwtUrlParameterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtUrlParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) RefreshRateLimitCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshRateLimitCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) RefreshRateLimitCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshRateLimitCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) RefreshRateLimitTimeWindowMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshRateLimitTimeWindowMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) RefreshRateLimitTimeWindowMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshRateLimitTimeWindowMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) RolesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) RolesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) SubjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) SubjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDatabaseOpensearchPropertiesOpenidOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseOpensearchPropertiesOpenidOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseOpensearchPropertiesOpenidOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesOpenidOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseOpensearchPropertiesOpenidOutputReference_Override(m ManagedDatabaseOpensearchPropertiesOpenidOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesOpenidOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetConnectUrl(val *string) {
	if err := j.validateSetConnectUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectUrl",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetHeader(val *string) {
	if err := j.validateSetHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"header",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetInternalValue(val *ManagedDatabaseOpensearchPropertiesOpenid) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetJwtHeader(val *string) {
	if err := j.validateSetJwtHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtHeader",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetJwtUrlParameter(val *string) {
	if err := j.validateSetJwtUrlParameterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtUrlParameter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetRefreshRateLimitCount(val *float64) {
	if err := j.validateSetRefreshRateLimitCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshRateLimitCount",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetRefreshRateLimitTimeWindowMs(val *float64) {
	if err := j.validateSetRefreshRateLimitTimeWindowMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshRateLimitTimeWindowMs",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetRolesKey(val *string) {
	if err := j.validateSetRolesKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rolesKey",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetSubjectKey(val *string) {
	if err := j.validateSetSubjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectKey",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		m,
		"resetClientId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		m,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetConnectUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetConnectUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		m,
		"resetHeader",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetJwtHeader() {
	_jsii_.InvokeVoid(
		m,
		"resetJwtHeader",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetJwtUrlParameter() {
	_jsii_.InvokeVoid(
		m,
		"resetJwtUrlParameter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetRefreshRateLimitCount() {
	_jsii_.InvokeVoid(
		m,
		"resetRefreshRateLimitCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetRefreshRateLimitTimeWindowMs() {
	_jsii_.InvokeVoid(
		m,
		"resetRefreshRateLimitTimeWindowMs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetRolesKey() {
	_jsii_.InvokeVoid(
		m,
		"resetRolesKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		m,
		"resetScope",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ResetSubjectKey() {
	_jsii_.InvokeVoid(
		m,
		"resetSubjectKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOpenidOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

