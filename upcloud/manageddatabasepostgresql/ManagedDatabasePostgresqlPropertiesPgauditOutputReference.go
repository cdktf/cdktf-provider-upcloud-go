// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasepostgresql

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v14/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v14/manageddatabasepostgresql/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabasePostgresqlPropertiesPgauditOutputReference interface {
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
	FeatureEnabled() interface{}
	SetFeatureEnabled(val interface{})
	FeatureEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ManagedDatabasePostgresqlPropertiesPgaudit
	SetInternalValue(val *ManagedDatabasePostgresqlPropertiesPgaudit)
	Log() *[]*string
	SetLog(val *[]*string)
	LogCatalog() interface{}
	SetLogCatalog(val interface{})
	LogCatalogInput() interface{}
	LogClient() interface{}
	SetLogClient(val interface{})
	LogClientInput() interface{}
	LogInput() *[]*string
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	LogMaxStringLength() *float64
	SetLogMaxStringLength(val *float64)
	LogMaxStringLengthInput() *float64
	LogNestedStatements() interface{}
	SetLogNestedStatements(val interface{})
	LogNestedStatementsInput() interface{}
	LogParameter() interface{}
	SetLogParameter(val interface{})
	LogParameterInput() interface{}
	LogParameterMaxSize() *float64
	SetLogParameterMaxSize(val *float64)
	LogParameterMaxSizeInput() *float64
	LogRelation() interface{}
	SetLogRelation(val interface{})
	LogRelationInput() interface{}
	LogRows() interface{}
	SetLogRows(val interface{})
	LogRowsInput() interface{}
	LogStatement() interface{}
	SetLogStatement(val interface{})
	LogStatementInput() interface{}
	LogStatementOnce() interface{}
	SetLogStatementOnce(val interface{})
	LogStatementOnceInput() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
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
	ResetFeatureEnabled()
	ResetLog()
	ResetLogCatalog()
	ResetLogClient()
	ResetLogLevel()
	ResetLogMaxStringLength()
	ResetLogNestedStatements()
	ResetLogParameter()
	ResetLogParameterMaxSize()
	ResetLogRelation()
	ResetLogRows()
	ResetLogStatement()
	ResetLogStatementOnce()
	ResetRole()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabasePostgresqlPropertiesPgauditOutputReference
type jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) FeatureEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"featureEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) FeatureEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"featureEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) InternalValue() *ManagedDatabasePostgresqlPropertiesPgaudit {
	var returns *ManagedDatabasePostgresqlPropertiesPgaudit
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) Log() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"log",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogCatalog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogCatalogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logCatalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogClient() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogClientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogMaxStringLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logMaxStringLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogMaxStringLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logMaxStringLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogNestedStatements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logNestedStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogNestedStatementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logNestedStatementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogParameter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogParameterMaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logParameterMaxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogParameterMaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logParameterMaxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogRelation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logRelation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogRelationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logRelationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogRows() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogRowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogStatement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogStatementOnce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStatementOnce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) LogStatementOnceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logStatementOnceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDatabasePostgresqlPropertiesPgauditOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabasePostgresqlPropertiesPgauditOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabasePostgresqlPropertiesPgauditOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresqlPropertiesPgauditOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabasePostgresqlPropertiesPgauditOutputReference_Override(m ManagedDatabasePostgresqlPropertiesPgauditOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresqlPropertiesPgauditOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetFeatureEnabled(val interface{}) {
	if err := j.validateSetFeatureEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetInternalValue(val *ManagedDatabasePostgresqlPropertiesPgaudit) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLog(val *[]*string) {
	if err := j.validateSetLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"log",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLogCatalog(val interface{}) {
	if err := j.validateSetLogCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logCatalog",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLogClient(val interface{}) {
	if err := j.validateSetLogClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logClient",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLogMaxStringLength(val *float64) {
	if err := j.validateSetLogMaxStringLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logMaxStringLength",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLogNestedStatements(val interface{}) {
	if err := j.validateSetLogNestedStatementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logNestedStatements",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLogParameter(val interface{}) {
	if err := j.validateSetLogParameterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logParameter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLogParameterMaxSize(val *float64) {
	if err := j.validateSetLogParameterMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logParameterMaxSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLogRelation(val interface{}) {
	if err := j.validateSetLogRelationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logRelation",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLogRows(val interface{}) {
	if err := j.validateSetLogRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logRows",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLogStatement(val interface{}) {
	if err := j.validateSetLogStatementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStatement",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetLogStatementOnce(val interface{}) {
	if err := j.validateSetLogStatementOnceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logStatementOnce",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetFeatureEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetFeatureEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLog() {
	_jsii_.InvokeVoid(
		m,
		"resetLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLogCatalog() {
	_jsii_.InvokeVoid(
		m,
		"resetLogCatalog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLogClient() {
	_jsii_.InvokeVoid(
		m,
		"resetLogClient",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLogMaxStringLength() {
	_jsii_.InvokeVoid(
		m,
		"resetLogMaxStringLength",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLogNestedStatements() {
	_jsii_.InvokeVoid(
		m,
		"resetLogNestedStatements",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLogParameter() {
	_jsii_.InvokeVoid(
		m,
		"resetLogParameter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLogParameterMaxSize() {
	_jsii_.InvokeVoid(
		m,
		"resetLogParameterMaxSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLogRelation() {
	_jsii_.InvokeVoid(
		m,
		"resetLogRelation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLogRows() {
	_jsii_.InvokeVoid(
		m,
		"resetLogRows",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLogStatement() {
	_jsii_.InvokeVoid(
		m,
		"resetLogStatement",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetLogStatementOnce() {
	_jsii_.InvokeVoid(
		m,
		"resetLogStatementOnce",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ResetRole() {
	_jsii_.InvokeVoid(
		m,
		"resetRole",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesPgauditOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

