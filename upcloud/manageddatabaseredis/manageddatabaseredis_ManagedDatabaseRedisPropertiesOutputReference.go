package manageddatabaseredis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v5/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v5/manageddatabaseredis/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseRedisPropertiesOutputReference interface {
	cdktf.ComplexObject
	AdditionalBackupRegions() *[]*string
	AutomaticUtilityNetworkIpFilter() interface{}
	SetAutomaticUtilityNetworkIpFilter(val interface{})
	AutomaticUtilityNetworkIpFilterInput() interface{}
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
	InternalValue() *ManagedDatabaseRedisProperties
	SetInternalValue(val *ManagedDatabaseRedisProperties)
	IpFilter() *[]*string
	SetIpFilter(val *[]*string)
	IpFilterInput() *[]*string
	Migration() ManagedDatabaseRedisPropertiesMigrationOutputReference
	MigrationInput() *ManagedDatabaseRedisPropertiesMigration
	PublicAccess() interface{}
	SetPublicAccess(val interface{})
	PublicAccessInput() interface{}
	RecoveryBasebackupName() *string
	SetRecoveryBasebackupName(val *string)
	RecoveryBasebackupNameInput() *string
	RedisAclChannelsDefault() *string
	SetRedisAclChannelsDefault(val *string)
	RedisAclChannelsDefaultInput() *string
	RedisIoThreads() *float64
	SetRedisIoThreads(val *float64)
	RedisIoThreadsInput() *float64
	RedisLfuDecayTime() *float64
	SetRedisLfuDecayTime(val *float64)
	RedisLfuDecayTimeInput() *float64
	RedisLfuLogFactor() *float64
	SetRedisLfuLogFactor(val *float64)
	RedisLfuLogFactorInput() *float64
	RedisMaxmemoryPolicy() *string
	SetRedisMaxmemoryPolicy(val *string)
	RedisMaxmemoryPolicyInput() *string
	RedisNotifyKeyspaceEvents() *string
	SetRedisNotifyKeyspaceEvents(val *string)
	RedisNotifyKeyspaceEventsInput() *string
	RedisNumberOfDatabases() *float64
	SetRedisNumberOfDatabases(val *float64)
	RedisNumberOfDatabasesInput() *float64
	RedisPersistence() *string
	SetRedisPersistence(val *string)
	RedisPersistenceInput() *string
	RedisPubsubClientOutputBufferLimit() *float64
	SetRedisPubsubClientOutputBufferLimit(val *float64)
	RedisPubsubClientOutputBufferLimitInput() *float64
	RedisSsl() interface{}
	SetRedisSsl(val interface{})
	RedisSslInput() interface{}
	RedisTimeout() *float64
	SetRedisTimeout(val *float64)
	RedisTimeoutInput() *float64
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
	PutMigration(value *ManagedDatabaseRedisPropertiesMigration)
	ResetAutomaticUtilityNetworkIpFilter()
	ResetIpFilter()
	ResetMigration()
	ResetPublicAccess()
	ResetRecoveryBasebackupName()
	ResetRedisAclChannelsDefault()
	ResetRedisIoThreads()
	ResetRedisLfuDecayTime()
	ResetRedisLfuLogFactor()
	ResetRedisMaxmemoryPolicy()
	ResetRedisNotifyKeyspaceEvents()
	ResetRedisNumberOfDatabases()
	ResetRedisPersistence()
	ResetRedisPubsubClientOutputBufferLimit()
	ResetRedisSsl()
	ResetRedisTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabaseRedisPropertiesOutputReference
type jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) AdditionalBackupRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalBackupRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) AutomaticUtilityNetworkIpFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUtilityNetworkIpFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) AutomaticUtilityNetworkIpFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUtilityNetworkIpFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) InternalValue() *ManagedDatabaseRedisProperties {
	var returns *ManagedDatabaseRedisProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) IpFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) IpFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) Migration() ManagedDatabaseRedisPropertiesMigrationOutputReference {
	var returns ManagedDatabaseRedisPropertiesMigrationOutputReference
	_jsii_.Get(
		j,
		"migration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) MigrationInput() *ManagedDatabaseRedisPropertiesMigration {
	var returns *ManagedDatabaseRedisPropertiesMigration
	_jsii_.Get(
		j,
		"migrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) PublicAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) PublicAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RecoveryBasebackupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryBasebackupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RecoveryBasebackupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryBasebackupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisAclChannelsDefault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisAclChannelsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisAclChannelsDefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisAclChannelsDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisIoThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisIoThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisIoThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisIoThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisLfuDecayTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisLfuDecayTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisLfuDecayTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisLfuDecayTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisLfuLogFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisLfuLogFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisLfuLogFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisLfuLogFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisMaxmemoryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisMaxmemoryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisMaxmemoryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisMaxmemoryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisNotifyKeyspaceEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisNotifyKeyspaceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisNotifyKeyspaceEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisNotifyKeyspaceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisNumberOfDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisNumberOfDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisNumberOfDatabasesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisNumberOfDatabasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisPersistence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisPersistence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisPersistenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisPersistenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisPubsubClientOutputBufferLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisPubsubClientOutputBufferLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisPubsubClientOutputBufferLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisPubsubClientOutputBufferLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redisSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redisSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) RedisTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"redisTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedDatabaseRedisPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseRedisPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseRedisPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseRedisPropertiesOutputReference_Override(m ManagedDatabaseRedisPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetAutomaticUtilityNetworkIpFilter(val interface{}) {
	if err := j.validateSetAutomaticUtilityNetworkIpFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticUtilityNetworkIpFilter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetInternalValue(val *ManagedDatabaseRedisProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetIpFilter(val *[]*string) {
	if err := j.validateSetIpFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFilter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetPublicAccess(val interface{}) {
	if err := j.validateSetPublicAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAccess",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRecoveryBasebackupName(val *string) {
	if err := j.validateSetRecoveryBasebackupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryBasebackupName",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRedisAclChannelsDefault(val *string) {
	if err := j.validateSetRedisAclChannelsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisAclChannelsDefault",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRedisIoThreads(val *float64) {
	if err := j.validateSetRedisIoThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisIoThreads",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRedisLfuDecayTime(val *float64) {
	if err := j.validateSetRedisLfuDecayTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisLfuDecayTime",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRedisLfuLogFactor(val *float64) {
	if err := j.validateSetRedisLfuLogFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisLfuLogFactor",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRedisMaxmemoryPolicy(val *string) {
	if err := j.validateSetRedisMaxmemoryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisMaxmemoryPolicy",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRedisNotifyKeyspaceEvents(val *string) {
	if err := j.validateSetRedisNotifyKeyspaceEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisNotifyKeyspaceEvents",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRedisNumberOfDatabases(val *float64) {
	if err := j.validateSetRedisNumberOfDatabasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisNumberOfDatabases",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRedisPersistence(val *string) {
	if err := j.validateSetRedisPersistenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisPersistence",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRedisPubsubClientOutputBufferLimit(val *float64) {
	if err := j.validateSetRedisPubsubClientOutputBufferLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisPubsubClientOutputBufferLimit",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRedisSsl(val interface{}) {
	if err := j.validateSetRedisSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisSsl",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetRedisTimeout(val *float64) {
	if err := j.validateSetRedisTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) PutMigration(value *ManagedDatabaseRedisPropertiesMigration) {
	if err := m.validatePutMigrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMigration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetAutomaticUtilityNetworkIpFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetAutomaticUtilityNetworkIpFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetIpFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetIpFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetMigration() {
	_jsii_.InvokeVoid(
		m,
		"resetMigration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetPublicAccess() {
	_jsii_.InvokeVoid(
		m,
		"resetPublicAccess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRecoveryBasebackupName() {
	_jsii_.InvokeVoid(
		m,
		"resetRecoveryBasebackupName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRedisAclChannelsDefault() {
	_jsii_.InvokeVoid(
		m,
		"resetRedisAclChannelsDefault",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRedisIoThreads() {
	_jsii_.InvokeVoid(
		m,
		"resetRedisIoThreads",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRedisLfuDecayTime() {
	_jsii_.InvokeVoid(
		m,
		"resetRedisLfuDecayTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRedisLfuLogFactor() {
	_jsii_.InvokeVoid(
		m,
		"resetRedisLfuLogFactor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRedisMaxmemoryPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetRedisMaxmemoryPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRedisNotifyKeyspaceEvents() {
	_jsii_.InvokeVoid(
		m,
		"resetRedisNotifyKeyspaceEvents",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRedisNumberOfDatabases() {
	_jsii_.InvokeVoid(
		m,
		"resetRedisNumberOfDatabases",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRedisPersistence() {
	_jsii_.InvokeVoid(
		m,
		"resetRedisPersistence",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRedisPubsubClientOutputBufferLimit() {
	_jsii_.InvokeVoid(
		m,
		"resetRedisPubsubClientOutputBufferLimit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRedisSsl() {
	_jsii_.InvokeVoid(
		m,
		"resetRedisSsl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ResetRedisTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetRedisTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

