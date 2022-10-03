package manageddatabasemysql

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-upcloud-go/upcloud/v3/jsii"

	"github.com/hashicorp/cdktf-provider-upcloud-go/upcloud/v3/manageddatabasemysql/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseMysqlPropertiesMigrationOutputReference interface {
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
	InternalValue() *ManagedDatabaseMysqlPropertiesMigration
	SetInternalValue(val *ManagedDatabaseMysqlPropertiesMigration)
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

// The jsii proxy struct for ManagedDatabaseMysqlPropertiesMigrationOutputReference
type jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) Dbname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) DbnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) IgnoreDbs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreDbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) IgnoreDbsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ignoreDbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) InternalValue() *ManagedDatabaseMysqlPropertiesMigration {
	var returns *ManagedDatabaseMysqlPropertiesMigration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) Ssl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) SslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewManagedDatabaseMysqlPropertiesMigrationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseMysqlPropertiesMigrationOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseMysqlPropertiesMigrationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlPropertiesMigrationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseMysqlPropertiesMigrationOutputReference_Override(m ManagedDatabaseMysqlPropertiesMigrationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlPropertiesMigrationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetDbname(val *string) {
	if err := j.validateSetDbnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbname",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetIgnoreDbs(val *string) {
	if err := j.validateSetIgnoreDbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreDbs",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetInternalValue(val *ManagedDatabaseMysqlPropertiesMigration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetSsl(val interface{}) {
	if err := j.validateSetSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssl",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) ResetDbname() {
	_jsii_.InvokeVoid(
		m,
		"resetDbname",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		m,
		"resetHost",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) ResetIgnoreDbs() {
	_jsii_.InvokeVoid(
		m,
		"resetIgnoreDbs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		m,
		"resetPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) ResetSsl() {
	_jsii_.InvokeVoid(
		m,
		"resetSsl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

