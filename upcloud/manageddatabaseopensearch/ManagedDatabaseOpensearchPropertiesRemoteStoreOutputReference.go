// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/manageddatabaseopensearch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *ManagedDatabaseOpensearchPropertiesRemoteStore
	SetInternalValue(val *ManagedDatabaseOpensearchPropertiesRemoteStore)
	SegmentPressureBytesLagVarianceFactor() *float64
	SetSegmentPressureBytesLagVarianceFactor(val *float64)
	SegmentPressureBytesLagVarianceFactorInput() *float64
	SegmentPressureConsecutiveFailuresLimit() *float64
	SetSegmentPressureConsecutiveFailuresLimit(val *float64)
	SegmentPressureConsecutiveFailuresLimitInput() *float64
	SegmentPressureEnabled() interface{}
	SetSegmentPressureEnabled(val interface{})
	SegmentPressureEnabledInput() interface{}
	SegmentPressureTimeLagVarianceFactor() *float64
	SetSegmentPressureTimeLagVarianceFactor(val *float64)
	SegmentPressureTimeLagVarianceFactorInput() *float64
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
	ResetSegmentPressureBytesLagVarianceFactor()
	ResetSegmentPressureConsecutiveFailuresLimit()
	ResetSegmentPressureEnabled()
	ResetSegmentPressureTimeLagVarianceFactor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference
type jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) InternalValue() *ManagedDatabaseOpensearchPropertiesRemoteStore {
	var returns *ManagedDatabaseOpensearchPropertiesRemoteStore
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) SegmentPressureBytesLagVarianceFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentPressureBytesLagVarianceFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) SegmentPressureBytesLagVarianceFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentPressureBytesLagVarianceFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) SegmentPressureConsecutiveFailuresLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentPressureConsecutiveFailuresLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) SegmentPressureConsecutiveFailuresLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentPressureConsecutiveFailuresLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) SegmentPressureEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentPressureEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) SegmentPressureEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentPressureEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) SegmentPressureTimeLagVarianceFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentPressureTimeLagVarianceFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) SegmentPressureTimeLagVarianceFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentPressureTimeLagVarianceFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseOpensearchPropertiesRemoteStoreOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference_Override(m ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference)SetInternalValue(val *ManagedDatabaseOpensearchPropertiesRemoteStore) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference)SetSegmentPressureBytesLagVarianceFactor(val *float64) {
	if err := j.validateSetSegmentPressureBytesLagVarianceFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentPressureBytesLagVarianceFactor",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference)SetSegmentPressureConsecutiveFailuresLimit(val *float64) {
	if err := j.validateSetSegmentPressureConsecutiveFailuresLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentPressureConsecutiveFailuresLimit",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference)SetSegmentPressureEnabled(val interface{}) {
	if err := j.validateSetSegmentPressureEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentPressureEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference)SetSegmentPressureTimeLagVarianceFactor(val *float64) {
	if err := j.validateSetSegmentPressureTimeLagVarianceFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentPressureTimeLagVarianceFactor",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) ResetSegmentPressureBytesLagVarianceFactor() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentPressureBytesLagVarianceFactor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) ResetSegmentPressureConsecutiveFailuresLimit() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentPressureConsecutiveFailuresLimit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) ResetSegmentPressureEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentPressureEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) ResetSegmentPressureTimeLagVarianceFactor() {
	_jsii_.InvokeVoid(
		m,
		"resetSegmentPressureTimeLagVarianceFactor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesRemoteStoreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

