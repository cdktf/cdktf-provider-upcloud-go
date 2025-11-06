// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/manageddatabaseopensearch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference interface {
	cdktf.ComplexObject
	CancellationBurst() *float64
	SetCancellationBurst(val *float64)
	CancellationBurstInput() *float64
	CancellationRate() *float64
	SetCancellationRate(val *float64)
	CancellationRateInput() *float64
	CancellationRatio() *float64
	SetCancellationRatio(val *float64)
	CancellationRatioInput() *float64
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
	CpuTimeMillisThreshold() *float64
	SetCpuTimeMillisThreshold(val *float64)
	CpuTimeMillisThresholdInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ElapsedTimeMillisThreshold() *float64
	SetElapsedTimeMillisThreshold(val *float64)
	ElapsedTimeMillisThresholdInput() *float64
	// Experimental.
	Fqn() *string
	HeapMovingAverageWindowSize() *float64
	SetHeapMovingAverageWindowSize(val *float64)
	HeapMovingAverageWindowSizeInput() *float64
	HeapPercentThreshold() *float64
	SetHeapPercentThreshold(val *float64)
	HeapPercentThresholdInput() *float64
	HeapVariance() *float64
	SetHeapVariance(val *float64)
	HeapVarianceInput() *float64
	InternalValue() *ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTask
	SetInternalValue(val *ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTask)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalHeapPercentThreshold() *float64
	SetTotalHeapPercentThreshold(val *float64)
	TotalHeapPercentThresholdInput() *float64
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
	ResetCancellationBurst()
	ResetCancellationRate()
	ResetCancellationRatio()
	ResetCpuTimeMillisThreshold()
	ResetElapsedTimeMillisThreshold()
	ResetHeapMovingAverageWindowSize()
	ResetHeapPercentThreshold()
	ResetHeapVariance()
	ResetTotalHeapPercentThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference
type jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) CancellationBurst() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cancellationBurst",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) CancellationBurstInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cancellationBurstInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) CancellationRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cancellationRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) CancellationRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cancellationRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) CancellationRatio() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cancellationRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) CancellationRatioInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cancellationRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) CpuTimeMillisThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuTimeMillisThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) CpuTimeMillisThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuTimeMillisThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ElapsedTimeMillisThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elapsedTimeMillisThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ElapsedTimeMillisThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elapsedTimeMillisThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) HeapMovingAverageWindowSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heapMovingAverageWindowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) HeapMovingAverageWindowSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heapMovingAverageWindowSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) HeapPercentThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heapPercentThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) HeapPercentThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heapPercentThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) HeapVariance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heapVariance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) HeapVarianceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heapVarianceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) InternalValue() *ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTask {
	var returns *ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTask
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) TotalHeapPercentThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalHeapPercentThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) TotalHeapPercentThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalHeapPercentThresholdInput",
		&returns,
	)
	return returns
}


func NewManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference_Override(m ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetCancellationBurst(val *float64) {
	if err := j.validateSetCancellationBurstParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cancellationBurst",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetCancellationRate(val *float64) {
	if err := j.validateSetCancellationRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cancellationRate",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetCancellationRatio(val *float64) {
	if err := j.validateSetCancellationRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cancellationRatio",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetCpuTimeMillisThreshold(val *float64) {
	if err := j.validateSetCpuTimeMillisThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuTimeMillisThreshold",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetElapsedTimeMillisThreshold(val *float64) {
	if err := j.validateSetElapsedTimeMillisThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elapsedTimeMillisThreshold",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetHeapMovingAverageWindowSize(val *float64) {
	if err := j.validateSetHeapMovingAverageWindowSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heapMovingAverageWindowSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetHeapPercentThreshold(val *float64) {
	if err := j.validateSetHeapPercentThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heapPercentThreshold",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetHeapVariance(val *float64) {
	if err := j.validateSetHeapVarianceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heapVariance",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetInternalValue(val *ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTask) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference)SetTotalHeapPercentThreshold(val *float64) {
	if err := j.validateSetTotalHeapPercentThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalHeapPercentThreshold",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ResetCancellationBurst() {
	_jsii_.InvokeVoid(
		m,
		"resetCancellationBurst",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ResetCancellationRate() {
	_jsii_.InvokeVoid(
		m,
		"resetCancellationRate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ResetCancellationRatio() {
	_jsii_.InvokeVoid(
		m,
		"resetCancellationRatio",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ResetCpuTimeMillisThreshold() {
	_jsii_.InvokeVoid(
		m,
		"resetCpuTimeMillisThreshold",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ResetElapsedTimeMillisThreshold() {
	_jsii_.InvokeVoid(
		m,
		"resetElapsedTimeMillisThreshold",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ResetHeapMovingAverageWindowSize() {
	_jsii_.InvokeVoid(
		m,
		"resetHeapMovingAverageWindowSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ResetHeapPercentThreshold() {
	_jsii_.InvokeVoid(
		m,
		"resetHeapPercentThreshold",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ResetHeapVariance() {
	_jsii_.InvokeVoid(
		m,
		"resetHeapVariance",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ResetTotalHeapPercentThreshold() {
	_jsii_.InvokeVoid(
		m,
		"resetTotalHeapPercentThreshold",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

