// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasevalkey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v14/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v14/manageddatabasevalkey/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseValkeyPropertiesOutputReference interface {
	cdktf.ComplexObject
	AutomaticUtilityNetworkIpFilter() interface{}
	SetAutomaticUtilityNetworkIpFilter(val interface{})
	AutomaticUtilityNetworkIpFilterInput() interface{}
	BackupHour() *float64
	SetBackupHour(val *float64)
	BackupHourInput() *float64
	BackupMinute() *float64
	SetBackupMinute(val *float64)
	BackupMinuteInput() *float64
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
	FrequentSnapshots() interface{}
	SetFrequentSnapshots(val interface{})
	FrequentSnapshotsInput() interface{}
	InternalValue() *ManagedDatabaseValkeyProperties
	SetInternalValue(val *ManagedDatabaseValkeyProperties)
	IpFilter() *[]*string
	SetIpFilter(val *[]*string)
	IpFilterInput() *[]*string
	Migration() ManagedDatabaseValkeyPropertiesMigrationOutputReference
	MigrationInput() *ManagedDatabaseValkeyPropertiesMigration
	PublicAccess() interface{}
	SetPublicAccess(val interface{})
	PublicAccessInput() interface{}
	ServiceLog() interface{}
	SetServiceLog(val interface{})
	ServiceLogInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValkeyAclChannelsDefault() *string
	SetValkeyAclChannelsDefault(val *string)
	ValkeyAclChannelsDefaultInput() *string
	ValkeyActiveExpireEffort() *float64
	SetValkeyActiveExpireEffort(val *float64)
	ValkeyActiveExpireEffortInput() *float64
	ValkeyIoThreads() *float64
	SetValkeyIoThreads(val *float64)
	ValkeyIoThreadsInput() *float64
	ValkeyLfuDecayTime() *float64
	SetValkeyLfuDecayTime(val *float64)
	ValkeyLfuDecayTimeInput() *float64
	ValkeyLfuLogFactor() *float64
	SetValkeyLfuLogFactor(val *float64)
	ValkeyLfuLogFactorInput() *float64
	ValkeyMaxmemoryPolicy() *string
	SetValkeyMaxmemoryPolicy(val *string)
	ValkeyMaxmemoryPolicyInput() *string
	ValkeyNotifyKeyspaceEvents() *string
	SetValkeyNotifyKeyspaceEvents(val *string)
	ValkeyNotifyKeyspaceEventsInput() *string
	ValkeyNumberOfDatabases() *float64
	SetValkeyNumberOfDatabases(val *float64)
	ValkeyNumberOfDatabasesInput() *float64
	ValkeyPersistence() *string
	SetValkeyPersistence(val *string)
	ValkeyPersistenceInput() *string
	ValkeyPubsubClientOutputBufferLimit() *float64
	SetValkeyPubsubClientOutputBufferLimit(val *float64)
	ValkeyPubsubClientOutputBufferLimitInput() *float64
	ValkeySsl() interface{}
	SetValkeySsl(val interface{})
	ValkeySslInput() interface{}
	ValkeyTimeout() *float64
	SetValkeyTimeout(val *float64)
	ValkeyTimeoutInput() *float64
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
	PutMigration(value *ManagedDatabaseValkeyPropertiesMigration)
	ResetAutomaticUtilityNetworkIpFilter()
	ResetBackupHour()
	ResetBackupMinute()
	ResetFrequentSnapshots()
	ResetIpFilter()
	ResetMigration()
	ResetPublicAccess()
	ResetServiceLog()
	ResetValkeyAclChannelsDefault()
	ResetValkeyActiveExpireEffort()
	ResetValkeyIoThreads()
	ResetValkeyLfuDecayTime()
	ResetValkeyLfuLogFactor()
	ResetValkeyMaxmemoryPolicy()
	ResetValkeyNotifyKeyspaceEvents()
	ResetValkeyNumberOfDatabases()
	ResetValkeyPersistence()
	ResetValkeyPubsubClientOutputBufferLimit()
	ResetValkeySsl()
	ResetValkeyTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabaseValkeyPropertiesOutputReference
type jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) AutomaticUtilityNetworkIpFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUtilityNetworkIpFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) AutomaticUtilityNetworkIpFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUtilityNetworkIpFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) BackupHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) BackupHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) BackupMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) BackupMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) FrequentSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frequentSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) FrequentSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frequentSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) InternalValue() *ManagedDatabaseValkeyProperties {
	var returns *ManagedDatabaseValkeyProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) IpFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) IpFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) Migration() ManagedDatabaseValkeyPropertiesMigrationOutputReference {
	var returns ManagedDatabaseValkeyPropertiesMigrationOutputReference
	_jsii_.Get(
		j,
		"migration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) MigrationInput() *ManagedDatabaseValkeyPropertiesMigration {
	var returns *ManagedDatabaseValkeyPropertiesMigration
	_jsii_.Get(
		j,
		"migrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) PublicAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) PublicAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ServiceLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ServiceLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyAclChannelsDefault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valkeyAclChannelsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyAclChannelsDefaultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valkeyAclChannelsDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyActiveExpireEffort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyActiveExpireEffort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyActiveExpireEffortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyActiveExpireEffortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyIoThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyIoThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyIoThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyIoThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyLfuDecayTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyLfuDecayTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyLfuDecayTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyLfuDecayTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyLfuLogFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyLfuLogFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyLfuLogFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyLfuLogFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyMaxmemoryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valkeyMaxmemoryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyMaxmemoryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valkeyMaxmemoryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyNotifyKeyspaceEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valkeyNotifyKeyspaceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyNotifyKeyspaceEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valkeyNotifyKeyspaceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyNumberOfDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyNumberOfDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyNumberOfDatabasesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyNumberOfDatabasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyPersistence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valkeyPersistence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyPersistenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valkeyPersistenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyPubsubClientOutputBufferLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyPubsubClientOutputBufferLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyPubsubClientOutputBufferLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyPubsubClientOutputBufferLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeySsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valkeySsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeySslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valkeySslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ValkeyTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valkeyTimeoutInput",
		&returns,
	)
	return returns
}


func NewManagedDatabaseValkeyPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseValkeyPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseValkeyPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseValkey.ManagedDatabaseValkeyPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseValkeyPropertiesOutputReference_Override(m ManagedDatabaseValkeyPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseValkey.ManagedDatabaseValkeyPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetAutomaticUtilityNetworkIpFilter(val interface{}) {
	if err := j.validateSetAutomaticUtilityNetworkIpFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticUtilityNetworkIpFilter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetBackupHour(val *float64) {
	if err := j.validateSetBackupHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupHour",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetBackupMinute(val *float64) {
	if err := j.validateSetBackupMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupMinute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetFrequentSnapshots(val interface{}) {
	if err := j.validateSetFrequentSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequentSnapshots",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetInternalValue(val *ManagedDatabaseValkeyProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetIpFilter(val *[]*string) {
	if err := j.validateSetIpFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFilter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetPublicAccess(val interface{}) {
	if err := j.validateSetPublicAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAccess",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetServiceLog(val interface{}) {
	if err := j.validateSetServiceLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLog",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeyAclChannelsDefault(val *string) {
	if err := j.validateSetValkeyAclChannelsDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyAclChannelsDefault",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeyActiveExpireEffort(val *float64) {
	if err := j.validateSetValkeyActiveExpireEffortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyActiveExpireEffort",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeyIoThreads(val *float64) {
	if err := j.validateSetValkeyIoThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyIoThreads",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeyLfuDecayTime(val *float64) {
	if err := j.validateSetValkeyLfuDecayTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyLfuDecayTime",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeyLfuLogFactor(val *float64) {
	if err := j.validateSetValkeyLfuLogFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyLfuLogFactor",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeyMaxmemoryPolicy(val *string) {
	if err := j.validateSetValkeyMaxmemoryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyMaxmemoryPolicy",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeyNotifyKeyspaceEvents(val *string) {
	if err := j.validateSetValkeyNotifyKeyspaceEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyNotifyKeyspaceEvents",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeyNumberOfDatabases(val *float64) {
	if err := j.validateSetValkeyNumberOfDatabasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyNumberOfDatabases",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeyPersistence(val *string) {
	if err := j.validateSetValkeyPersistenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyPersistence",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeyPubsubClientOutputBufferLimit(val *float64) {
	if err := j.validateSetValkeyPubsubClientOutputBufferLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyPubsubClientOutputBufferLimit",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeySsl(val interface{}) {
	if err := j.validateSetValkeySslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeySsl",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference)SetValkeyTimeout(val *float64) {
	if err := j.validateSetValkeyTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valkeyTimeout",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) PutMigration(value *ManagedDatabaseValkeyPropertiesMigration) {
	if err := m.validatePutMigrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMigration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetAutomaticUtilityNetworkIpFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetAutomaticUtilityNetworkIpFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetBackupHour() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupHour",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetBackupMinute() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupMinute",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetFrequentSnapshots() {
	_jsii_.InvokeVoid(
		m,
		"resetFrequentSnapshots",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetIpFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetIpFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetMigration() {
	_jsii_.InvokeVoid(
		m,
		"resetMigration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetPublicAccess() {
	_jsii_.InvokeVoid(
		m,
		"resetPublicAccess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetServiceLog() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeyAclChannelsDefault() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeyAclChannelsDefault",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeyActiveExpireEffort() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeyActiveExpireEffort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeyIoThreads() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeyIoThreads",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeyLfuDecayTime() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeyLfuDecayTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeyLfuLogFactor() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeyLfuLogFactor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeyMaxmemoryPolicy() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeyMaxmemoryPolicy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeyNotifyKeyspaceEvents() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeyNotifyKeyspaceEvents",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeyNumberOfDatabases() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeyNumberOfDatabases",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeyPersistence() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeyPersistence",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeyPubsubClientOutputBufferLimit() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeyPubsubClientOutputBufferLimit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeySsl() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeySsl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ResetValkeyTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetValkeyTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabaseValkeyPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

