// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/manageddatabaseopensearch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseOpensearchPropertiesJwtOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ManagedDatabaseOpensearchPropertiesJwt
	SetInternalValue(val *ManagedDatabaseOpensearchPropertiesJwt)
	JwtClockSkewToleranceSeconds() *float64
	SetJwtClockSkewToleranceSeconds(val *float64)
	JwtClockSkewToleranceSecondsInput() *float64
	JwtHeader() *string
	SetJwtHeader(val *string)
	JwtHeaderInput() *string
	JwtUrlParameter() *string
	SetJwtUrlParameter(val *string)
	JwtUrlParameterInput() *string
	RequiredAudience() *string
	SetRequiredAudience(val *string)
	RequiredAudienceInput() *string
	RequiredIssuer() *string
	SetRequiredIssuer(val *string)
	RequiredIssuerInput() *string
	RolesKey() *string
	SetRolesKey(val *string)
	RolesKeyInput() *string
	SigningKey() *string
	SetSigningKey(val *string)
	SigningKeyInput() *string
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
	ResetEnabled()
	ResetJwtClockSkewToleranceSeconds()
	ResetJwtHeader()
	ResetJwtUrlParameter()
	ResetRequiredAudience()
	ResetRequiredIssuer()
	ResetRolesKey()
	ResetSigningKey()
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

// The jsii proxy struct for ManagedDatabaseOpensearchPropertiesJwtOutputReference
type jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) InternalValue() *ManagedDatabaseOpensearchPropertiesJwt {
	var returns *ManagedDatabaseOpensearchPropertiesJwt
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) JwtClockSkewToleranceSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jwtClockSkewToleranceSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) JwtClockSkewToleranceSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jwtClockSkewToleranceSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) JwtHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) JwtHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) JwtUrlParameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtUrlParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) JwtUrlParameterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwtUrlParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) RequiredAudience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requiredAudience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) RequiredAudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requiredAudienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) RequiredIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requiredIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) RequiredIssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requiredIssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) RolesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) RolesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) SigningKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) SigningKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) SubjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) SubjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDatabaseOpensearchPropertiesJwtOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseOpensearchPropertiesJwtOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseOpensearchPropertiesJwtOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesJwtOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseOpensearchPropertiesJwtOutputReference_Override(m ManagedDatabaseOpensearchPropertiesJwtOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesJwtOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetInternalValue(val *ManagedDatabaseOpensearchPropertiesJwt) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetJwtClockSkewToleranceSeconds(val *float64) {
	if err := j.validateSetJwtClockSkewToleranceSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtClockSkewToleranceSeconds",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetJwtHeader(val *string) {
	if err := j.validateSetJwtHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtHeader",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetJwtUrlParameter(val *string) {
	if err := j.validateSetJwtUrlParameterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtUrlParameter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetRequiredAudience(val *string) {
	if err := j.validateSetRequiredAudienceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredAudience",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetRequiredIssuer(val *string) {
	if err := j.validateSetRequiredIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredIssuer",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetRolesKey(val *string) {
	if err := j.validateSetRolesKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rolesKey",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetSigningKey(val *string) {
	if err := j.validateSetSigningKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signingKey",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetSubjectKey(val *string) {
	if err := j.validateSetSubjectKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectKey",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ResetJwtClockSkewToleranceSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetJwtClockSkewToleranceSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ResetJwtHeader() {
	_jsii_.InvokeVoid(
		m,
		"resetJwtHeader",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ResetJwtUrlParameter() {
	_jsii_.InvokeVoid(
		m,
		"resetJwtUrlParameter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ResetRequiredAudience() {
	_jsii_.InvokeVoid(
		m,
		"resetRequiredAudience",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ResetRequiredIssuer() {
	_jsii_.InvokeVoid(
		m,
		"resetRequiredIssuer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ResetRolesKey() {
	_jsii_.InvokeVoid(
		m,
		"resetRolesKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ResetSigningKey() {
	_jsii_.InvokeVoid(
		m,
		"resetSigningKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ResetSubjectKey() {
	_jsii_.InvokeVoid(
		m,
		"resetSubjectKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesJwtOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

