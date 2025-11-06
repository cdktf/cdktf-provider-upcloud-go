// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/manageddatabaseopensearch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference interface {
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
	InternalValue() *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameter
	SetInternalValue(val *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameter)
	NodeAttribute() ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterNodeOutputReference
	NodeAttributeInput() *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterNode
	Shard() ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterShardOutputReference
	ShardInput() *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterShard
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
	PutNodeAttribute(value *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterNode)
	PutShard(value *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterShard)
	ResetNodeAttribute()
	ResetShard()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference
type jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) InternalValue() *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameter {
	var returns *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) NodeAttribute() ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterNodeOutputReference {
	var returns ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterNodeOutputReference
	_jsii_.Get(
		j,
		"nodeAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) NodeAttributeInput() *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterNode {
	var returns *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterNode
	_jsii_.Get(
		j,
		"nodeAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) Shard() ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterShardOutputReference {
	var returns ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterShardOutputReference
	_jsii_.Get(
		j,
		"shard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) ShardInput() *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterShard {
	var returns *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterShard
	_jsii_.Get(
		j,
		"shardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference_Override(m ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference)SetInternalValue(val *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) PutNodeAttribute(value *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterNode) {
	if err := m.validatePutNodeAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNodeAttribute",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) PutShard(value *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterShard) {
	if err := m.validatePutShardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putShard",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) ResetNodeAttribute() {
	_jsii_.InvokeVoid(
		m,
		"resetNodeAttribute",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) ResetShard() {
	_jsii_.InvokeVoid(
		m,
		"resetShard",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

