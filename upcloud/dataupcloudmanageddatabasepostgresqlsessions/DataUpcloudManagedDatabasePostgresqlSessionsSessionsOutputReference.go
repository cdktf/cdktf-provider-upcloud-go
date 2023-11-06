// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataupcloudmanageddatabasepostgresqlsessions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v11/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v11/dataupcloudmanageddatabasepostgresqlsessions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference interface {
	cdktf.ComplexObject
	ApplicationName() *string
	BackendStart() *string
	BackendType() *string
	BackendXid() *float64
	SetBackendXid(val *float64)
	BackendXidInput() *float64
	BackendXmin() *float64
	SetBackendXmin(val *float64)
	BackendXminInput() *float64
	ClientAddr() *string
	ClientHostname() *string
	SetClientHostname(val *string)
	ClientHostnameInput() *string
	ClientPort() *float64
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
	Datid() *float64
	Datname() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Query() *string
	QueryDuration() *string
	QueryStart() *string
	State() *string
	StateChange() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Usename() *string
	Usesysid() *float64
	WaitEvent() *string
	WaitEventType() *string
	XactStart() *string
	SetXactStart(val *string)
	XactStartInput() *string
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
	ResetBackendXid()
	ResetBackendXmin()
	ResetClientHostname()
	ResetXactStart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference
type jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) BackendStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) BackendType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) BackendXid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendXid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) BackendXidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendXidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) BackendXmin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendXmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) BackendXminInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendXminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ClientAddr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientAddr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ClientHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ClientHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ClientPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) Datid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"datid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) Datname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) QueryDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) QueryStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) StateChange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) Usename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) Usesysid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usesysid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) WaitEvent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) WaitEventType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitEventType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) XactStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xactStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) XactStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xactStartInput",
		&returns,
	)
	return returns
}


func NewDataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.dataUpcloudManagedDatabasePostgresqlSessions.DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference_Override(d DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.dataUpcloudManagedDatabasePostgresqlSessions.DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference)SetBackendXid(val *float64) {
	if err := j.validateSetBackendXidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendXid",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference)SetBackendXmin(val *float64) {
	if err := j.validateSetBackendXminParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendXmin",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference)SetClientHostname(val *string) {
	if err := j.validateSetClientHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientHostname",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference)SetXactStart(val *string) {
	if err := j.validateSetXactStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xactStart",
		val,
	)
}

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ResetBackendXid() {
	_jsii_.InvokeVoid(
		d,
		"resetBackendXid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ResetBackendXmin() {
	_jsii_.InvokeVoid(
		d,
		"resetBackendXmin",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ResetClientHostname() {
	_jsii_.InvokeVoid(
		d,
		"resetClientHostname",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ResetXactStart() {
	_jsii_.InvokeVoid(
		d,
		"resetXactStart",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataUpcloudManagedDatabasePostgresqlSessionsSessionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

