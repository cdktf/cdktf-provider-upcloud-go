// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasepostgresql

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v12/manageddatabasepostgresql/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference interface {
	cdktf.ComplexObject
	AutodbIdleTimeout() *float64
	SetAutodbIdleTimeout(val *float64)
	AutodbIdleTimeoutInput() *float64
	AutodbMaxDbConnections() *float64
	SetAutodbMaxDbConnections(val *float64)
	AutodbMaxDbConnectionsInput() *float64
	AutodbPoolMode() *string
	SetAutodbPoolMode(val *string)
	AutodbPoolModeInput() *string
	AutodbPoolSize() *float64
	SetAutodbPoolSize(val *float64)
	AutodbPoolSizeInput() *float64
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
	IgnoreStartupParameters() *[]*string
	SetIgnoreStartupParameters(val *[]*string)
	IgnoreStartupParametersInput() *[]*string
	InternalValue() *ManagedDatabasePostgresqlPropertiesPgbouncer
	SetInternalValue(val *ManagedDatabasePostgresqlPropertiesPgbouncer)
	MinPoolSize() *float64
	SetMinPoolSize(val *float64)
	MinPoolSizeInput() *float64
	ServerIdleTimeout() *float64
	SetServerIdleTimeout(val *float64)
	ServerIdleTimeoutInput() *float64
	ServerLifetime() *float64
	SetServerLifetime(val *float64)
	ServerLifetimeInput() *float64
	ServerResetQueryAlways() interface{}
	SetServerResetQueryAlways(val interface{})
	ServerResetQueryAlwaysInput() interface{}
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
	ResetAutodbIdleTimeout()
	ResetAutodbMaxDbConnections()
	ResetAutodbPoolMode()
	ResetAutodbPoolSize()
	ResetIgnoreStartupParameters()
	ResetMinPoolSize()
	ResetServerIdleTimeout()
	ResetServerLifetime()
	ResetServerResetQueryAlways()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference
type jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) AutodbIdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbIdleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) AutodbIdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbIdleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) AutodbMaxDbConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbMaxDbConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) AutodbMaxDbConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbMaxDbConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) AutodbPoolMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autodbPoolMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) AutodbPoolModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autodbPoolModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) AutodbPoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbPoolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) AutodbPoolSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autodbPoolSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) IgnoreStartupParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreStartupParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) IgnoreStartupParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreStartupParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) InternalValue() *ManagedDatabasePostgresqlPropertiesPgbouncer {
	var returns *ManagedDatabasePostgresqlPropertiesPgbouncer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) MinPoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPoolSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) MinPoolSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minPoolSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ServerIdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverIdleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ServerIdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverIdleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ServerLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ServerLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serverLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ServerResetQueryAlways() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverResetQueryAlways",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ServerResetQueryAlwaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverResetQueryAlwaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDatabasePostgresqlPropertiesPgbouncerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabasePostgresqlPropertiesPgbouncerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabasePostgresqlPropertiesPgbouncerOutputReference_Override(m ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetAutodbIdleTimeout(val *float64) {
	if err := j.validateSetAutodbIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodbIdleTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetAutodbMaxDbConnections(val *float64) {
	if err := j.validateSetAutodbMaxDbConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodbMaxDbConnections",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetAutodbPoolMode(val *string) {
	if err := j.validateSetAutodbPoolModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodbPoolMode",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetAutodbPoolSize(val *float64) {
	if err := j.validateSetAutodbPoolSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autodbPoolSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetIgnoreStartupParameters(val *[]*string) {
	if err := j.validateSetIgnoreStartupParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreStartupParameters",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetInternalValue(val *ManagedDatabasePostgresqlPropertiesPgbouncer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetMinPoolSize(val *float64) {
	if err := j.validateSetMinPoolSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minPoolSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetServerIdleTimeout(val *float64) {
	if err := j.validateSetServerIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverIdleTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetServerLifetime(val *float64) {
	if err := j.validateSetServerLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverLifetime",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetServerResetQueryAlways(val interface{}) {
	if err := j.validateSetServerResetQueryAlwaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverResetQueryAlways",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ResetAutodbIdleTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetAutodbIdleTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ResetAutodbMaxDbConnections() {
	_jsii_.InvokeVoid(
		m,
		"resetAutodbMaxDbConnections",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ResetAutodbPoolMode() {
	_jsii_.InvokeVoid(
		m,
		"resetAutodbPoolMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ResetAutodbPoolSize() {
	_jsii_.InvokeVoid(
		m,
		"resetAutodbPoolSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ResetIgnoreStartupParameters() {
	_jsii_.InvokeVoid(
		m,
		"resetIgnoreStartupParameters",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ResetMinPoolSize() {
	_jsii_.InvokeVoid(
		m,
		"resetMinPoolSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ResetServerIdleTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetServerIdleTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ResetServerLifetime() {
	_jsii_.InvokeVoid(
		m,
		"resetServerLifetime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ResetServerResetQueryAlways() {
	_jsii_.InvokeVoid(
		m,
		"resetServerResetQueryAlways",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

