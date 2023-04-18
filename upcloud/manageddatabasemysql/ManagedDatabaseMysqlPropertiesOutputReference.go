package manageddatabasemysql

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v6/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v6/manageddatabasemysql/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseMysqlPropertiesOutputReference interface {
	cdktf.ComplexObject
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
	AutomaticUtilityNetworkIpFilter() interface{}
	SetAutomaticUtilityNetworkIpFilter(val interface{})
	AutomaticUtilityNetworkIpFilterInput() interface{}
	BackupHour() *float64
	SetBackupHour(val *float64)
	BackupHourInput() *float64
	BackupMinute() *float64
	SetBackupMinute(val *float64)
	BackupMinuteInput() *float64
	BinlogRetentionPeriod() *float64
	SetBinlogRetentionPeriod(val *float64)
	BinlogRetentionPeriodInput() *float64
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
	ConnectTimeout() *float64
	SetConnectTimeout(val *float64)
	ConnectTimeoutInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultTimeZone() *string
	SetDefaultTimeZone(val *string)
	DefaultTimeZoneInput() *string
	// Experimental.
	Fqn() *string
	GroupConcatMaxLen() *float64
	SetGroupConcatMaxLen(val *float64)
	GroupConcatMaxLenInput() *float64
	InformationSchemaStatsExpiry() *float64
	SetInformationSchemaStatsExpiry(val *float64)
	InformationSchemaStatsExpiryInput() *float64
	InnodbChangeBufferMaxSize() *float64
	SetInnodbChangeBufferMaxSize(val *float64)
	InnodbChangeBufferMaxSizeInput() *float64
	InnodbFlushNeighbors() *float64
	SetInnodbFlushNeighbors(val *float64)
	InnodbFlushNeighborsInput() *float64
	InnodbFtMinTokenSize() *float64
	SetInnodbFtMinTokenSize(val *float64)
	InnodbFtMinTokenSizeInput() *float64
	InnodbFtServerStopwordTable() *string
	SetInnodbFtServerStopwordTable(val *string)
	InnodbFtServerStopwordTableInput() *string
	InnodbLockWaitTimeout() *float64
	SetInnodbLockWaitTimeout(val *float64)
	InnodbLockWaitTimeoutInput() *float64
	InnodbLogBufferSize() *float64
	SetInnodbLogBufferSize(val *float64)
	InnodbLogBufferSizeInput() *float64
	InnodbOnlineAlterLogMaxSize() *float64
	SetInnodbOnlineAlterLogMaxSize(val *float64)
	InnodbOnlineAlterLogMaxSizeInput() *float64
	InnodbPrintAllDeadlocks() interface{}
	SetInnodbPrintAllDeadlocks(val interface{})
	InnodbPrintAllDeadlocksInput() interface{}
	InnodbReadIoThreads() *float64
	SetInnodbReadIoThreads(val *float64)
	InnodbReadIoThreadsInput() *float64
	InnodbRollbackOnTimeout() interface{}
	SetInnodbRollbackOnTimeout(val interface{})
	InnodbRollbackOnTimeoutInput() interface{}
	InnodbThreadConcurrency() *float64
	SetInnodbThreadConcurrency(val *float64)
	InnodbThreadConcurrencyInput() *float64
	InnodbWriteIoThreads() *float64
	SetInnodbWriteIoThreads(val *float64)
	InnodbWriteIoThreadsInput() *float64
	InteractiveTimeout() *float64
	SetInteractiveTimeout(val *float64)
	InteractiveTimeoutInput() *float64
	InternalTmpMemStorageEngine() *string
	SetInternalTmpMemStorageEngine(val *string)
	InternalTmpMemStorageEngineInput() *string
	InternalValue() *ManagedDatabaseMysqlProperties
	SetInternalValue(val *ManagedDatabaseMysqlProperties)
	IpFilter() *[]*string
	SetIpFilter(val *[]*string)
	IpFilterInput() *[]*string
	LongQueryTime() *float64
	SetLongQueryTime(val *float64)
	LongQueryTimeInput() *float64
	MaxAllowedPacket() *float64
	SetMaxAllowedPacket(val *float64)
	MaxAllowedPacketInput() *float64
	MaxHeapTableSize() *float64
	SetMaxHeapTableSize(val *float64)
	MaxHeapTableSizeInput() *float64
	Migration() ManagedDatabaseMysqlPropertiesMigrationOutputReference
	MigrationInput() *ManagedDatabaseMysqlPropertiesMigration
	NetBufferLength() *float64
	SetNetBufferLength(val *float64)
	NetBufferLengthInput() *float64
	NetReadTimeout() *float64
	SetNetReadTimeout(val *float64)
	NetReadTimeoutInput() *float64
	NetWriteTimeout() *float64
	SetNetWriteTimeout(val *float64)
	NetWriteTimeoutInput() *float64
	PublicAccess() interface{}
	SetPublicAccess(val interface{})
	PublicAccessInput() interface{}
	SlowQueryLog() interface{}
	SetSlowQueryLog(val interface{})
	SlowQueryLogInput() interface{}
	SortBufferSize() *float64
	SetSortBufferSize(val *float64)
	SortBufferSizeInput() *float64
	SqlMode() *string
	SetSqlMode(val *string)
	SqlModeInput() *string
	SqlRequirePrimaryKey() interface{}
	SetSqlRequirePrimaryKey(val interface{})
	SqlRequirePrimaryKeyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TmpTableSize() *float64
	SetTmpTableSize(val *float64)
	TmpTableSizeInput() *float64
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	WaitTimeout() *float64
	SetWaitTimeout(val *float64)
	WaitTimeoutInput() *float64
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
	PutMigration(value *ManagedDatabaseMysqlPropertiesMigration)
	ResetAdminPassword()
	ResetAdminUsername()
	ResetAutomaticUtilityNetworkIpFilter()
	ResetBackupHour()
	ResetBackupMinute()
	ResetBinlogRetentionPeriod()
	ResetConnectTimeout()
	ResetDefaultTimeZone()
	ResetGroupConcatMaxLen()
	ResetInformationSchemaStatsExpiry()
	ResetInnodbChangeBufferMaxSize()
	ResetInnodbFlushNeighbors()
	ResetInnodbFtMinTokenSize()
	ResetInnodbFtServerStopwordTable()
	ResetInnodbLockWaitTimeout()
	ResetInnodbLogBufferSize()
	ResetInnodbOnlineAlterLogMaxSize()
	ResetInnodbPrintAllDeadlocks()
	ResetInnodbReadIoThreads()
	ResetInnodbRollbackOnTimeout()
	ResetInnodbThreadConcurrency()
	ResetInnodbWriteIoThreads()
	ResetInteractiveTimeout()
	ResetInternalTmpMemStorageEngine()
	ResetIpFilter()
	ResetLongQueryTime()
	ResetMaxAllowedPacket()
	ResetMaxHeapTableSize()
	ResetMigration()
	ResetNetBufferLength()
	ResetNetReadTimeout()
	ResetNetWriteTimeout()
	ResetPublicAccess()
	ResetSlowQueryLog()
	ResetSortBufferSize()
	ResetSqlMode()
	ResetSqlRequirePrimaryKey()
	ResetTmpTableSize()
	ResetVersion()
	ResetWaitTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabaseMysqlPropertiesOutputReference
type jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) AutomaticUtilityNetworkIpFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUtilityNetworkIpFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) AutomaticUtilityNetworkIpFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUtilityNetworkIpFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) BackupHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) BackupHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) BackupMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) BackupMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) BinlogRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"binlogRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) BinlogRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"binlogRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ConnectTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ConnectTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) DefaultTimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTimeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) DefaultTimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTimeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) GroupConcatMaxLen() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupConcatMaxLen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) GroupConcatMaxLenInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"groupConcatMaxLenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InformationSchemaStatsExpiry() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"informationSchemaStatsExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InformationSchemaStatsExpiryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"informationSchemaStatsExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbChangeBufferMaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbChangeBufferMaxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbChangeBufferMaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbChangeBufferMaxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbFlushNeighbors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbFlushNeighbors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbFlushNeighborsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbFlushNeighborsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbFtMinTokenSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbFtMinTokenSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbFtMinTokenSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbFtMinTokenSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbFtServerStopwordTable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"innodbFtServerStopwordTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbFtServerStopwordTableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"innodbFtServerStopwordTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbLockWaitTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbLockWaitTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbLockWaitTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbLockWaitTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbLogBufferSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbLogBufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbLogBufferSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbLogBufferSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbOnlineAlterLogMaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbOnlineAlterLogMaxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbOnlineAlterLogMaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbOnlineAlterLogMaxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbPrintAllDeadlocks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"innodbPrintAllDeadlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbPrintAllDeadlocksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"innodbPrintAllDeadlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbReadIoThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbReadIoThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbReadIoThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbReadIoThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbRollbackOnTimeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"innodbRollbackOnTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbRollbackOnTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"innodbRollbackOnTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbThreadConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbThreadConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbThreadConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbThreadConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbWriteIoThreads() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbWriteIoThreads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InnodbWriteIoThreadsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"innodbWriteIoThreadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InteractiveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interactiveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InteractiveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interactiveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InternalTmpMemStorageEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalTmpMemStorageEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InternalTmpMemStorageEngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internalTmpMemStorageEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InternalValue() *ManagedDatabaseMysqlProperties {
	var returns *ManagedDatabaseMysqlProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) IpFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) IpFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) LongQueryTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longQueryTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) LongQueryTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"longQueryTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) MaxAllowedPacket() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedPacket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) MaxAllowedPacketInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedPacketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) MaxHeapTableSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHeapTableSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) MaxHeapTableSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxHeapTableSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) Migration() ManagedDatabaseMysqlPropertiesMigrationOutputReference {
	var returns ManagedDatabaseMysqlPropertiesMigrationOutputReference
	_jsii_.Get(
		j,
		"migration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) MigrationInput() *ManagedDatabaseMysqlPropertiesMigration {
	var returns *ManagedDatabaseMysqlPropertiesMigration
	_jsii_.Get(
		j,
		"migrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) NetBufferLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netBufferLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) NetBufferLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netBufferLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) NetReadTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netReadTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) NetReadTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netReadTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) NetWriteTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netWriteTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) NetWriteTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netWriteTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) PublicAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) PublicAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) SlowQueryLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slowQueryLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) SlowQueryLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slowQueryLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) SortBufferSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sortBufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) SortBufferSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sortBufferSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) SqlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) SqlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) SqlRequirePrimaryKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlRequirePrimaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) SqlRequirePrimaryKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlRequirePrimaryKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) TmpTableSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tmpTableSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) TmpTableSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tmpTableSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) WaitTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) WaitTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitTimeoutInput",
		&returns,
	)
	return returns
}


func NewManagedDatabaseMysqlPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseMysqlPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseMysqlPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseMysqlPropertiesOutputReference_Override(m ManagedDatabaseMysqlPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetAdminUsername(val *string) {
	if err := j.validateSetAdminUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetAutomaticUtilityNetworkIpFilter(val interface{}) {
	if err := j.validateSetAutomaticUtilityNetworkIpFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticUtilityNetworkIpFilter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetBackupHour(val *float64) {
	if err := j.validateSetBackupHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupHour",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetBackupMinute(val *float64) {
	if err := j.validateSetBackupMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupMinute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetBinlogRetentionPeriod(val *float64) {
	if err := j.validateSetBinlogRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binlogRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetConnectTimeout(val *float64) {
	if err := j.validateSetConnectTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetDefaultTimeZone(val *string) {
	if err := j.validateSetDefaultTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTimeZone",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetGroupConcatMaxLen(val *float64) {
	if err := j.validateSetGroupConcatMaxLenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupConcatMaxLen",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInformationSchemaStatsExpiry(val *float64) {
	if err := j.validateSetInformationSchemaStatsExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"informationSchemaStatsExpiry",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbChangeBufferMaxSize(val *float64) {
	if err := j.validateSetInnodbChangeBufferMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbChangeBufferMaxSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbFlushNeighbors(val *float64) {
	if err := j.validateSetInnodbFlushNeighborsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbFlushNeighbors",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbFtMinTokenSize(val *float64) {
	if err := j.validateSetInnodbFtMinTokenSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbFtMinTokenSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbFtServerStopwordTable(val *string) {
	if err := j.validateSetInnodbFtServerStopwordTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbFtServerStopwordTable",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbLockWaitTimeout(val *float64) {
	if err := j.validateSetInnodbLockWaitTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbLockWaitTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbLogBufferSize(val *float64) {
	if err := j.validateSetInnodbLogBufferSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbLogBufferSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbOnlineAlterLogMaxSize(val *float64) {
	if err := j.validateSetInnodbOnlineAlterLogMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbOnlineAlterLogMaxSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbPrintAllDeadlocks(val interface{}) {
	if err := j.validateSetInnodbPrintAllDeadlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbPrintAllDeadlocks",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbReadIoThreads(val *float64) {
	if err := j.validateSetInnodbReadIoThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbReadIoThreads",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbRollbackOnTimeout(val interface{}) {
	if err := j.validateSetInnodbRollbackOnTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbRollbackOnTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbThreadConcurrency(val *float64) {
	if err := j.validateSetInnodbThreadConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbThreadConcurrency",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInnodbWriteIoThreads(val *float64) {
	if err := j.validateSetInnodbWriteIoThreadsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"innodbWriteIoThreads",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInteractiveTimeout(val *float64) {
	if err := j.validateSetInteractiveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interactiveTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInternalTmpMemStorageEngine(val *string) {
	if err := j.validateSetInternalTmpMemStorageEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalTmpMemStorageEngine",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetInternalValue(val *ManagedDatabaseMysqlProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetIpFilter(val *[]*string) {
	if err := j.validateSetIpFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFilter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetLongQueryTime(val *float64) {
	if err := j.validateSetLongQueryTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longQueryTime",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetMaxAllowedPacket(val *float64) {
	if err := j.validateSetMaxAllowedPacketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllowedPacket",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetMaxHeapTableSize(val *float64) {
	if err := j.validateSetMaxHeapTableSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxHeapTableSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetNetBufferLength(val *float64) {
	if err := j.validateSetNetBufferLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netBufferLength",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetNetReadTimeout(val *float64) {
	if err := j.validateSetNetReadTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netReadTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetNetWriteTimeout(val *float64) {
	if err := j.validateSetNetWriteTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netWriteTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetPublicAccess(val interface{}) {
	if err := j.validateSetPublicAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAccess",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetSlowQueryLog(val interface{}) {
	if err := j.validateSetSlowQueryLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slowQueryLog",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetSortBufferSize(val *float64) {
	if err := j.validateSetSortBufferSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortBufferSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetSqlMode(val *string) {
	if err := j.validateSetSqlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlMode",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetSqlRequirePrimaryKey(val interface{}) {
	if err := j.validateSetSqlRequirePrimaryKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlRequirePrimaryKey",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetTmpTableSize(val *float64) {
	if err := j.validateSetTmpTableSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tmpTableSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference)SetWaitTimeout(val *float64) {
	if err := j.validateSetWaitTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitTimeout",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) PutMigration(value *ManagedDatabaseMysqlPropertiesMigration) {
	if err := m.validatePutMigrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMigration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetAdminUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetAdminUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetAutomaticUtilityNetworkIpFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetAutomaticUtilityNetworkIpFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetBackupHour() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupHour",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetBackupMinute() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupMinute",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetBinlogRetentionPeriod() {
	_jsii_.InvokeVoid(
		m,
		"resetBinlogRetentionPeriod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetConnectTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetConnectTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetDefaultTimeZone() {
	_jsii_.InvokeVoid(
		m,
		"resetDefaultTimeZone",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetGroupConcatMaxLen() {
	_jsii_.InvokeVoid(
		m,
		"resetGroupConcatMaxLen",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInformationSchemaStatsExpiry() {
	_jsii_.InvokeVoid(
		m,
		"resetInformationSchemaStatsExpiry",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbChangeBufferMaxSize() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbChangeBufferMaxSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbFlushNeighbors() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbFlushNeighbors",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbFtMinTokenSize() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbFtMinTokenSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbFtServerStopwordTable() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbFtServerStopwordTable",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbLockWaitTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbLockWaitTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbLogBufferSize() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbLogBufferSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbOnlineAlterLogMaxSize() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbOnlineAlterLogMaxSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbPrintAllDeadlocks() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbPrintAllDeadlocks",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbReadIoThreads() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbReadIoThreads",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbRollbackOnTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbRollbackOnTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbThreadConcurrency() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbThreadConcurrency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInnodbWriteIoThreads() {
	_jsii_.InvokeVoid(
		m,
		"resetInnodbWriteIoThreads",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInteractiveTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetInteractiveTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetInternalTmpMemStorageEngine() {
	_jsii_.InvokeVoid(
		m,
		"resetInternalTmpMemStorageEngine",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetIpFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetIpFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetLongQueryTime() {
	_jsii_.InvokeVoid(
		m,
		"resetLongQueryTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetMaxAllowedPacket() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxAllowedPacket",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetMaxHeapTableSize() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxHeapTableSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetMigration() {
	_jsii_.InvokeVoid(
		m,
		"resetMigration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetNetBufferLength() {
	_jsii_.InvokeVoid(
		m,
		"resetNetBufferLength",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetNetReadTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetNetReadTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetNetWriteTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetNetWriteTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetPublicAccess() {
	_jsii_.InvokeVoid(
		m,
		"resetPublicAccess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetSlowQueryLog() {
	_jsii_.InvokeVoid(
		m,
		"resetSlowQueryLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetSortBufferSize() {
	_jsii_.InvokeVoid(
		m,
		"resetSortBufferSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetSqlMode() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlMode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetSqlRequirePrimaryKey() {
	_jsii_.InvokeVoid(
		m,
		"resetSqlRequirePrimaryKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetTmpTableSize() {
	_jsii_.InvokeVoid(
		m,
		"resetTmpTableSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ResetWaitTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetWaitTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

