// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/manageddatabaseopensearch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseOpensearchPropertiesSegrepOutputReference interface {
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
	InternalValue() *ManagedDatabaseOpensearchPropertiesSegrep
	SetInternalValue(val *ManagedDatabaseOpensearchPropertiesSegrep)
	PressureCheckpointLimit() *float64
	SetPressureCheckpointLimit(val *float64)
	PressureCheckpointLimitInput() *float64
	PressureEnabled() interface{}
	SetPressureEnabled(val interface{})
	PressureEnabledInput() interface{}
	PressureReplicaStaleLimit() *float64
	SetPressureReplicaStaleLimit(val *float64)
	PressureReplicaStaleLimitInput() *float64
	PressureTimeLimit() *string
	SetPressureTimeLimit(val *string)
	PressureTimeLimitInput() *string
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
	ResetPressureCheckpointLimit()
	ResetPressureEnabled()
	ResetPressureReplicaStaleLimit()
	ResetPressureTimeLimit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabaseOpensearchPropertiesSegrepOutputReference
type jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) InternalValue() *ManagedDatabaseOpensearchPropertiesSegrep {
	var returns *ManagedDatabaseOpensearchPropertiesSegrep
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) PressureCheckpointLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pressureCheckpointLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) PressureCheckpointLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pressureCheckpointLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) PressureEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pressureEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) PressureEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pressureEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) PressureReplicaStaleLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pressureReplicaStaleLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) PressureReplicaStaleLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pressureReplicaStaleLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) PressureTimeLimit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pressureTimeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) PressureTimeLimitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pressureTimeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDatabaseOpensearchPropertiesSegrepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseOpensearchPropertiesSegrepOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseOpensearchPropertiesSegrepOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesSegrepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseOpensearchPropertiesSegrepOutputReference_Override(m ManagedDatabaseOpensearchPropertiesSegrepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesSegrepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference)SetInternalValue(val *ManagedDatabaseOpensearchPropertiesSegrep) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference)SetPressureCheckpointLimit(val *float64) {
	if err := j.validateSetPressureCheckpointLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pressureCheckpointLimit",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference)SetPressureEnabled(val interface{}) {
	if err := j.validateSetPressureEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pressureEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference)SetPressureReplicaStaleLimit(val *float64) {
	if err := j.validateSetPressureReplicaStaleLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pressureReplicaStaleLimit",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference)SetPressureTimeLimit(val *string) {
	if err := j.validateSetPressureTimeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pressureTimeLimit",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) ResetPressureCheckpointLimit() {
	_jsii_.InvokeVoid(
		m,
		"resetPressureCheckpointLimit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) ResetPressureEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetPressureEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) ResetPressureReplicaStaleLimit() {
	_jsii_.InvokeVoid(
		m,
		"resetPressureReplicaStaleLimit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) ResetPressureTimeLimit() {
	_jsii_.InvokeVoid(
		m,
		"resetPressureTimeLimit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSegrepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

