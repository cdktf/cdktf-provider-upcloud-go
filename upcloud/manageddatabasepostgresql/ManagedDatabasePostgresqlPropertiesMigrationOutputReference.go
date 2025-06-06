// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasepostgresql

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/manageddatabasepostgresql/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabasePostgresqlPropertiesMigrationOutputReference interface {
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
	Dbname() *string
	SetDbname(val *string)
	DbnameInput() *string
	// Experimental.
	Fqn() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IgnoreDbs() *string
	SetIgnoreDbs(val *string)
	IgnoreDbsInput() *string
	IgnoreRoles() *string
	SetIgnoreRoles(val *string)
	IgnoreRolesInput() *string
	InternalValue() *ManagedDatabasePostgresqlPropertiesMigration
	SetInternalValue(val *ManagedDatabasePostgresqlPropertiesMigration)
	Method() *string
	SetMethod(val *string)
	MethodInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Ssl() interface{}
	SetSsl(val interface{})
	SslInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetDbname()
	ResetHost()
	ResetIgnoreDbs()
	ResetIgnoreRoles()
	ResetMethod()
	ResetPassword()
	ResetPort()
	ResetSsl()
	ResetUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabasePostgresqlPropertiesMigrationOutputReference
type jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) Dbname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) DbnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) IgnoreDbs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreDbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) IgnoreDbsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreDbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) IgnoreRoles() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) IgnoreRolesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) InternalValue() *ManagedDatabasePostgresqlPropertiesMigration {
	var returns *ManagedDatabasePostgresqlPropertiesMigration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) Ssl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) SslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewManagedDatabasePostgresqlPropertiesMigrationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabasePostgresqlPropertiesMigrationOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabasePostgresqlPropertiesMigrationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresqlPropertiesMigrationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabasePostgresqlPropertiesMigrationOutputReference_Override(m ManagedDatabasePostgresqlPropertiesMigrationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresqlPropertiesMigrationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetDbname(val *string) {
	if err := j.validateSetDbnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbname",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetIgnoreDbs(val *string) {
	if err := j.validateSetIgnoreDbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreDbs",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetIgnoreRoles(val *string) {
	if err := j.validateSetIgnoreRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreRoles",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetInternalValue(val *ManagedDatabasePostgresqlPropertiesMigration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetSsl(val interface{}) {
	if err := j.validateSetSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssl",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ResetDbname() {
	_jsii_.InvokeVoid(
		m,
		"resetDbname",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		m,
		"resetHost",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ResetIgnoreDbs() {
	_jsii_.InvokeVoid(
		m,
		"resetIgnoreDbs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ResetIgnoreRoles() {
	_jsii_.InvokeVoid(
		m,
		"resetIgnoreRoles",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		m,
		"resetMethod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		m,
		"resetPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ResetSsl() {
	_jsii_.InvokeVoid(
		m,
		"resetSsl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesMigrationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

