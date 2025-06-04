// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataupcloudmanageddatabasevalkeysessions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/dataupcloudmanageddatabasevalkeysessions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference interface {
	cdktf.ComplexObject
	ActiveChannelSubscriptions() *float64
	ActiveDatabase() *string
	ActivePatternMatchingChannelSubscriptions() *float64
	ApplicationName() *string
	ClientAddr() *string
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
	ConnectionAge() *float64
	ConnectionIdle() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Flags() *[]*string
	FlagsRaw() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MultiExecCommands() *float64
	OutputBuffer() *float64
	OutputBufferMemory() *float64
	OutputListLength() *float64
	Query() *string
	QueryBuffer() *float64
	QueryBufferFree() *float64
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference
type jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) ActiveChannelSubscriptions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activeChannelSubscriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) ActiveDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) ActivePatternMatchingChannelSubscriptions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activePatternMatchingChannelSubscriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) ClientAddr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAddr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) ConnectionAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) ConnectionIdle() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionIdle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) Flags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"flags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) FlagsRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) MultiExecCommands() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiExecCommands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) OutputBuffer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outputBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) OutputBufferMemory() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outputBufferMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) OutputListLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outputListLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) QueryBuffer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) QueryBufferFree() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryBufferFree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataUpcloudManagedDatabaseValkeySessionsSessionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.dataUpcloudManagedDatabaseValkeySessions.DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference_Override(d DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.dataUpcloudManagedDatabaseValkeySessions.DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabaseValkeySessionsSessionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

