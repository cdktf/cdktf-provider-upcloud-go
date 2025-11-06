// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesnodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/kubernetesnodegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesNodeGroupGpuPlanOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StorageSize() *float64
	SetStorageSize(val *float64)
	StorageSizeInput() *float64
	StorageTier() *string
	SetStorageTier(val *string)
	StorageTierInput() *string
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
	ResetStorageSize()
	ResetStorageTier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesNodeGroupGpuPlanOutputReference
type jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) StorageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) StorageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) StorageTier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) StorageTierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesNodeGroupGpuPlanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) KubernetesNodeGroupGpuPlanOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesNodeGroupGpuPlanOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.kubernetesNodeGroup.KubernetesNodeGroupGpuPlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewKubernetesNodeGroupGpuPlanOutputReference_Override(k KubernetesNodeGroupGpuPlanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.kubernetesNodeGroup.KubernetesNodeGroupGpuPlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference)SetStorageSize(val *float64) {
	if err := j.validateSetStorageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageSize",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference)SetStorageTier(val *string) {
	if err := j.validateSetStorageTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageTier",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) ResetStorageSize() {
	_jsii_.InvokeVoid(
		k,
		"resetStorageSize",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) ResetStorageTier() {
	_jsii_.InvokeVoid(
		k,
		"resetStorageTier",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesNodeGroupGpuPlanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

