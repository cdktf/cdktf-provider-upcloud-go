package manageddatabasepostgresql

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v7/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v7/manageddatabasepostgresql/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabasePostgresqlPropertiesOutputReference interface {
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
	AutovacuumAnalyzeScaleFactor() *float64
	SetAutovacuumAnalyzeScaleFactor(val *float64)
	AutovacuumAnalyzeScaleFactorInput() *float64
	AutovacuumAnalyzeThreshold() *float64
	SetAutovacuumAnalyzeThreshold(val *float64)
	AutovacuumAnalyzeThresholdInput() *float64
	AutovacuumFreezeMaxAge() *float64
	SetAutovacuumFreezeMaxAge(val *float64)
	AutovacuumFreezeMaxAgeInput() *float64
	AutovacuumMaxWorkers() *float64
	SetAutovacuumMaxWorkers(val *float64)
	AutovacuumMaxWorkersInput() *float64
	AutovacuumNaptime() *float64
	SetAutovacuumNaptime(val *float64)
	AutovacuumNaptimeInput() *float64
	AutovacuumVacuumCostDelay() *float64
	SetAutovacuumVacuumCostDelay(val *float64)
	AutovacuumVacuumCostDelayInput() *float64
	AutovacuumVacuumCostLimit() *float64
	SetAutovacuumVacuumCostLimit(val *float64)
	AutovacuumVacuumCostLimitInput() *float64
	AutovacuumVacuumScaleFactor() *float64
	SetAutovacuumVacuumScaleFactor(val *float64)
	AutovacuumVacuumScaleFactorInput() *float64
	AutovacuumVacuumThreshold() *float64
	SetAutovacuumVacuumThreshold(val *float64)
	AutovacuumVacuumThresholdInput() *float64
	BackupHour() *float64
	SetBackupHour(val *float64)
	BackupHourInput() *float64
	BackupMinute() *float64
	SetBackupMinute(val *float64)
	BackupMinuteInput() *float64
	BgwriterDelay() *float64
	SetBgwriterDelay(val *float64)
	BgwriterDelayInput() *float64
	BgwriterFlushAfter() *float64
	SetBgwriterFlushAfter(val *float64)
	BgwriterFlushAfterInput() *float64
	BgwriterLruMaxpages() *float64
	SetBgwriterLruMaxpages(val *float64)
	BgwriterLruMaxpagesInput() *float64
	BgwriterLruMultiplier() *float64
	SetBgwriterLruMultiplier(val *float64)
	BgwriterLruMultiplierInput() *float64
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
	DeadlockTimeout() *float64
	SetDeadlockTimeout(val *float64)
	DeadlockTimeoutInput() *float64
	DefaultToastCompression() *string
	SetDefaultToastCompression(val *string)
	DefaultToastCompressionInput() *string
	// Experimental.
	Fqn() *string
	IdleInTransactionSessionTimeout() *float64
	SetIdleInTransactionSessionTimeout(val *float64)
	IdleInTransactionSessionTimeoutInput() *float64
	InternalValue() *ManagedDatabasePostgresqlProperties
	SetInternalValue(val *ManagedDatabasePostgresqlProperties)
	IpFilter() *[]*string
	SetIpFilter(val *[]*string)
	IpFilterInput() *[]*string
	Jit() interface{}
	SetJit(val interface{})
	JitInput() interface{}
	LogAutovacuumMinDuration() *float64
	SetLogAutovacuumMinDuration(val *float64)
	LogAutovacuumMinDurationInput() *float64
	LogErrorVerbosity() *string
	SetLogErrorVerbosity(val *string)
	LogErrorVerbosityInput() *string
	LogLinePrefix() *string
	SetLogLinePrefix(val *string)
	LogLinePrefixInput() *string
	LogMinDurationStatement() *float64
	SetLogMinDurationStatement(val *float64)
	LogMinDurationStatementInput() *float64
	LogTempFiles() *float64
	SetLogTempFiles(val *float64)
	LogTempFilesInput() *float64
	MaxFilesPerProcess() *float64
	SetMaxFilesPerProcess(val *float64)
	MaxFilesPerProcessInput() *float64
	MaxLocksPerTransaction() *float64
	SetMaxLocksPerTransaction(val *float64)
	MaxLocksPerTransactionInput() *float64
	MaxLogicalReplicationWorkers() *float64
	SetMaxLogicalReplicationWorkers(val *float64)
	MaxLogicalReplicationWorkersInput() *float64
	MaxParallelWorkers() *float64
	SetMaxParallelWorkers(val *float64)
	MaxParallelWorkersInput() *float64
	MaxParallelWorkersPerGather() *float64
	SetMaxParallelWorkersPerGather(val *float64)
	MaxParallelWorkersPerGatherInput() *float64
	MaxPredLocksPerTransaction() *float64
	SetMaxPredLocksPerTransaction(val *float64)
	MaxPredLocksPerTransactionInput() *float64
	MaxPreparedTransactions() *float64
	SetMaxPreparedTransactions(val *float64)
	MaxPreparedTransactionsInput() *float64
	MaxReplicationSlots() *float64
	SetMaxReplicationSlots(val *float64)
	MaxReplicationSlotsInput() *float64
	MaxSlotWalKeepSize() *float64
	SetMaxSlotWalKeepSize(val *float64)
	MaxSlotWalKeepSizeInput() *float64
	MaxStackDepth() *float64
	SetMaxStackDepth(val *float64)
	MaxStackDepthInput() *float64
	MaxStandbyArchiveDelay() *float64
	SetMaxStandbyArchiveDelay(val *float64)
	MaxStandbyArchiveDelayInput() *float64
	MaxStandbyStreamingDelay() *float64
	SetMaxStandbyStreamingDelay(val *float64)
	MaxStandbyStreamingDelayInput() *float64
	MaxWalSenders() *float64
	SetMaxWalSenders(val *float64)
	MaxWalSendersInput() *float64
	MaxWorkerProcesses() *float64
	SetMaxWorkerProcesses(val *float64)
	MaxWorkerProcessesInput() *float64
	Migration() ManagedDatabasePostgresqlPropertiesMigrationOutputReference
	MigrationInput() *ManagedDatabasePostgresqlPropertiesMigration
	Pgbouncer() ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference
	PgbouncerInput() *ManagedDatabasePostgresqlPropertiesPgbouncer
	Pglookout() ManagedDatabasePostgresqlPropertiesPglookoutOutputReference
	PglookoutInput() *ManagedDatabasePostgresqlPropertiesPglookout
	PgPartmanBgwInterval() *float64
	SetPgPartmanBgwInterval(val *float64)
	PgPartmanBgwIntervalInput() *float64
	PgPartmanBgwRole() *string
	SetPgPartmanBgwRole(val *string)
	PgPartmanBgwRoleInput() *string
	PgReadReplica() interface{}
	SetPgReadReplica(val interface{})
	PgReadReplicaInput() interface{}
	PgServiceToForkFrom() *string
	SetPgServiceToForkFrom(val *string)
	PgServiceToForkFromInput() *string
	PgStatMonitorEnable() interface{}
	SetPgStatMonitorEnable(val interface{})
	PgStatMonitorEnableInput() interface{}
	PgStatMonitorPgsmEnableQueryPlan() interface{}
	SetPgStatMonitorPgsmEnableQueryPlan(val interface{})
	PgStatMonitorPgsmEnableQueryPlanInput() interface{}
	PgStatMonitorPgsmMaxBuckets() *float64
	SetPgStatMonitorPgsmMaxBuckets(val *float64)
	PgStatMonitorPgsmMaxBucketsInput() *float64
	PgStatStatementsTrack() *string
	SetPgStatStatementsTrack(val *string)
	PgStatStatementsTrackInput() *string
	PublicAccess() interface{}
	SetPublicAccess(val interface{})
	PublicAccessInput() interface{}
	SharedBuffersPercentage() *float64
	SetSharedBuffersPercentage(val *float64)
	SharedBuffersPercentageInput() *float64
	SynchronousReplication() *string
	SetSynchronousReplication(val *string)
	SynchronousReplicationInput() *string
	TempFileLimit() *float64
	SetTempFileLimit(val *float64)
	TempFileLimitInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timescaledb() ManagedDatabasePostgresqlPropertiesTimescaledbOutputReference
	TimescaledbInput() *ManagedDatabasePostgresqlPropertiesTimescaledb
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	TrackActivityQuerySize() *float64
	SetTrackActivityQuerySize(val *float64)
	TrackActivityQuerySizeInput() *float64
	TrackCommitTimestamp() *string
	SetTrackCommitTimestamp(val *string)
	TrackCommitTimestampInput() *string
	TrackFunctions() *string
	SetTrackFunctions(val *string)
	TrackFunctionsInput() *string
	TrackIoTiming() *string
	SetTrackIoTiming(val *string)
	TrackIoTimingInput() *string
	Variant() *string
	SetVariant(val *string)
	VariantInput() *string
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	WalSenderTimeout() *float64
	SetWalSenderTimeout(val *float64)
	WalSenderTimeoutInput() *float64
	WalWriterDelay() *float64
	SetWalWriterDelay(val *float64)
	WalWriterDelayInput() *float64
	WorkMem() *float64
	SetWorkMem(val *float64)
	WorkMemInput() *float64
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
	PutMigration(value *ManagedDatabasePostgresqlPropertiesMigration)
	PutPgbouncer(value *ManagedDatabasePostgresqlPropertiesPgbouncer)
	PutPglookout(value *ManagedDatabasePostgresqlPropertiesPglookout)
	PutTimescaledb(value *ManagedDatabasePostgresqlPropertiesTimescaledb)
	ResetAdminPassword()
	ResetAdminUsername()
	ResetAutomaticUtilityNetworkIpFilter()
	ResetAutovacuumAnalyzeScaleFactor()
	ResetAutovacuumAnalyzeThreshold()
	ResetAutovacuumFreezeMaxAge()
	ResetAutovacuumMaxWorkers()
	ResetAutovacuumNaptime()
	ResetAutovacuumVacuumCostDelay()
	ResetAutovacuumVacuumCostLimit()
	ResetAutovacuumVacuumScaleFactor()
	ResetAutovacuumVacuumThreshold()
	ResetBackupHour()
	ResetBackupMinute()
	ResetBgwriterDelay()
	ResetBgwriterFlushAfter()
	ResetBgwriterLruMaxpages()
	ResetBgwriterLruMultiplier()
	ResetDeadlockTimeout()
	ResetDefaultToastCompression()
	ResetIdleInTransactionSessionTimeout()
	ResetIpFilter()
	ResetJit()
	ResetLogAutovacuumMinDuration()
	ResetLogErrorVerbosity()
	ResetLogLinePrefix()
	ResetLogMinDurationStatement()
	ResetLogTempFiles()
	ResetMaxFilesPerProcess()
	ResetMaxLocksPerTransaction()
	ResetMaxLogicalReplicationWorkers()
	ResetMaxParallelWorkers()
	ResetMaxParallelWorkersPerGather()
	ResetMaxPredLocksPerTransaction()
	ResetMaxPreparedTransactions()
	ResetMaxReplicationSlots()
	ResetMaxSlotWalKeepSize()
	ResetMaxStackDepth()
	ResetMaxStandbyArchiveDelay()
	ResetMaxStandbyStreamingDelay()
	ResetMaxWalSenders()
	ResetMaxWorkerProcesses()
	ResetMigration()
	ResetPgbouncer()
	ResetPglookout()
	ResetPgPartmanBgwInterval()
	ResetPgPartmanBgwRole()
	ResetPgReadReplica()
	ResetPgServiceToForkFrom()
	ResetPgStatMonitorEnable()
	ResetPgStatMonitorPgsmEnableQueryPlan()
	ResetPgStatMonitorPgsmMaxBuckets()
	ResetPgStatStatementsTrack()
	ResetPublicAccess()
	ResetSharedBuffersPercentage()
	ResetSynchronousReplication()
	ResetTempFileLimit()
	ResetTimescaledb()
	ResetTimezone()
	ResetTrackActivityQuerySize()
	ResetTrackCommitTimestamp()
	ResetTrackFunctions()
	ResetTrackIoTiming()
	ResetVariant()
	ResetVersion()
	ResetWalSenderTimeout()
	ResetWalWriterDelay()
	ResetWorkMem()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabasePostgresqlPropertiesOutputReference
type jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutomaticUtilityNetworkIpFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUtilityNetworkIpFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutomaticUtilityNetworkIpFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUtilityNetworkIpFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumAnalyzeScaleFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumAnalyzeScaleFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumAnalyzeScaleFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumAnalyzeScaleFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumAnalyzeThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumAnalyzeThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumAnalyzeThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumAnalyzeThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumFreezeMaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumFreezeMaxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumFreezeMaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumFreezeMaxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumMaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumMaxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumMaxWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumMaxWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumNaptime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumNaptime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumNaptimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumNaptimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumVacuumCostDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumCostDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumVacuumCostDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumCostDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumVacuumCostLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumCostLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumVacuumCostLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumCostLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumVacuumScaleFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumScaleFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumVacuumScaleFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumScaleFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumVacuumThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) AutovacuumVacuumThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autovacuumVacuumThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BackupHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BackupHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BackupMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BackupMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BgwriterDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BgwriterDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BgwriterFlushAfter() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterFlushAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BgwriterFlushAfterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterFlushAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BgwriterLruMaxpages() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterLruMaxpages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BgwriterLruMaxpagesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterLruMaxpagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BgwriterLruMultiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterLruMultiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) BgwriterLruMultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bgwriterLruMultiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) DeadlockTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deadlockTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) DeadlockTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deadlockTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) DefaultToastCompression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultToastCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) DefaultToastCompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultToastCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) IdleInTransactionSessionTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleInTransactionSessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) IdleInTransactionSessionTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleInTransactionSessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) InternalValue() *ManagedDatabasePostgresqlProperties {
	var returns *ManagedDatabasePostgresqlProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) IpFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) IpFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) Jit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) JitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) LogAutovacuumMinDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logAutovacuumMinDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) LogAutovacuumMinDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logAutovacuumMinDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) LogErrorVerbosity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logErrorVerbosity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) LogErrorVerbosityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logErrorVerbosityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) LogLinePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLinePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) LogLinePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLinePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) LogMinDurationStatement() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logMinDurationStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) LogMinDurationStatementInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logMinDurationStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) LogTempFiles() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logTempFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) LogTempFilesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logTempFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxFilesPerProcess() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFilesPerProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxFilesPerProcessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFilesPerProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxLocksPerTransaction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLocksPerTransaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxLocksPerTransactionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLocksPerTransactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxLogicalReplicationWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLogicalReplicationWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxLogicalReplicationWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLogicalReplicationWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxParallelWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxParallelWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxParallelWorkersPerGather() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelWorkersPerGather",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxParallelWorkersPerGatherInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelWorkersPerGatherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxPredLocksPerTransaction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPredLocksPerTransaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxPredLocksPerTransactionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPredLocksPerTransactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxPreparedTransactions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPreparedTransactions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxPreparedTransactionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPreparedTransactionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxReplicationSlots() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicationSlots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxReplicationSlotsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicationSlotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxSlotWalKeepSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSlotWalKeepSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxSlotWalKeepSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSlotWalKeepSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxStackDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStackDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxStackDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStackDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxStandbyArchiveDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStandbyArchiveDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxStandbyArchiveDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStandbyArchiveDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxStandbyStreamingDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStandbyStreamingDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxStandbyStreamingDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStandbyStreamingDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxWalSenders() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWalSenders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxWalSendersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWalSendersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxWorkerProcesses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkerProcesses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MaxWorkerProcessesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkerProcessesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) Migration() ManagedDatabasePostgresqlPropertiesMigrationOutputReference {
	var returns ManagedDatabasePostgresqlPropertiesMigrationOutputReference
	_jsii_.Get(
		j,
		"migration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) MigrationInput() *ManagedDatabasePostgresqlPropertiesMigration {
	var returns *ManagedDatabasePostgresqlPropertiesMigration
	_jsii_.Get(
		j,
		"migrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) Pgbouncer() ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference {
	var returns ManagedDatabasePostgresqlPropertiesPgbouncerOutputReference
	_jsii_.Get(
		j,
		"pgbouncer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgbouncerInput() *ManagedDatabasePostgresqlPropertiesPgbouncer {
	var returns *ManagedDatabasePostgresqlPropertiesPgbouncer
	_jsii_.Get(
		j,
		"pgbouncerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) Pglookout() ManagedDatabasePostgresqlPropertiesPglookoutOutputReference {
	var returns ManagedDatabasePostgresqlPropertiesPglookoutOutputReference
	_jsii_.Get(
		j,
		"pglookout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PglookoutInput() *ManagedDatabasePostgresqlPropertiesPglookout {
	var returns *ManagedDatabasePostgresqlPropertiesPglookout
	_jsii_.Get(
		j,
		"pglookoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgPartmanBgwInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pgPartmanBgwInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgPartmanBgwIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pgPartmanBgwIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgPartmanBgwRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgPartmanBgwRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgPartmanBgwRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgPartmanBgwRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgReadReplica() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pgReadReplica",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgReadReplicaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pgReadReplicaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgServiceToForkFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgServiceToForkFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgServiceToForkFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgServiceToForkFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgStatMonitorEnable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pgStatMonitorEnable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgStatMonitorEnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pgStatMonitorEnableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgStatMonitorPgsmEnableQueryPlan() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pgStatMonitorPgsmEnableQueryPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgStatMonitorPgsmEnableQueryPlanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pgStatMonitorPgsmEnableQueryPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgStatMonitorPgsmMaxBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pgStatMonitorPgsmMaxBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgStatMonitorPgsmMaxBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pgStatMonitorPgsmMaxBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgStatStatementsTrack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgStatStatementsTrack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PgStatStatementsTrackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pgStatStatementsTrackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PublicAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PublicAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) SharedBuffersPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedBuffersPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) SharedBuffersPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharedBuffersPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) SynchronousReplication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"synchronousReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) SynchronousReplicationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"synchronousReplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TempFileLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tempFileLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TempFileLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tempFileLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) Timescaledb() ManagedDatabasePostgresqlPropertiesTimescaledbOutputReference {
	var returns ManagedDatabasePostgresqlPropertiesTimescaledbOutputReference
	_jsii_.Get(
		j,
		"timescaledb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TimescaledbInput() *ManagedDatabasePostgresqlPropertiesTimescaledb {
	var returns *ManagedDatabasePostgresqlPropertiesTimescaledb
	_jsii_.Get(
		j,
		"timescaledbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TrackActivityQuerySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trackActivityQuerySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TrackActivityQuerySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"trackActivityQuerySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TrackCommitTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackCommitTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TrackCommitTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackCommitTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TrackFunctions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TrackFunctionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TrackIoTiming() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackIoTiming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) TrackIoTimingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackIoTimingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) Variant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) VariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) WalSenderTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"walSenderTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) WalSenderTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"walSenderTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) WalWriterDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"walWriterDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) WalWriterDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"walWriterDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) WorkMem() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workMem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) WorkMemInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workMemInput",
		&returns,
	)
	return returns
}


func NewManagedDatabasePostgresqlPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabasePostgresqlPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabasePostgresqlPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresqlPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabasePostgresqlPropertiesOutputReference_Override(m ManagedDatabasePostgresqlPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresqlPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAdminUsername(val *string) {
	if err := j.validateSetAdminUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAutomaticUtilityNetworkIpFilter(val interface{}) {
	if err := j.validateSetAutomaticUtilityNetworkIpFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticUtilityNetworkIpFilter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAutovacuumAnalyzeScaleFactor(val *float64) {
	if err := j.validateSetAutovacuumAnalyzeScaleFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumAnalyzeScaleFactor",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAutovacuumAnalyzeThreshold(val *float64) {
	if err := j.validateSetAutovacuumAnalyzeThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumAnalyzeThreshold",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAutovacuumFreezeMaxAge(val *float64) {
	if err := j.validateSetAutovacuumFreezeMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumFreezeMaxAge",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAutovacuumMaxWorkers(val *float64) {
	if err := j.validateSetAutovacuumMaxWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumMaxWorkers",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAutovacuumNaptime(val *float64) {
	if err := j.validateSetAutovacuumNaptimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumNaptime",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAutovacuumVacuumCostDelay(val *float64) {
	if err := j.validateSetAutovacuumVacuumCostDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumVacuumCostDelay",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAutovacuumVacuumCostLimit(val *float64) {
	if err := j.validateSetAutovacuumVacuumCostLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumVacuumCostLimit",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAutovacuumVacuumScaleFactor(val *float64) {
	if err := j.validateSetAutovacuumVacuumScaleFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumVacuumScaleFactor",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetAutovacuumVacuumThreshold(val *float64) {
	if err := j.validateSetAutovacuumVacuumThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autovacuumVacuumThreshold",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetBackupHour(val *float64) {
	if err := j.validateSetBackupHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupHour",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetBackupMinute(val *float64) {
	if err := j.validateSetBackupMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupMinute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetBgwriterDelay(val *float64) {
	if err := j.validateSetBgwriterDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgwriterDelay",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetBgwriterFlushAfter(val *float64) {
	if err := j.validateSetBgwriterFlushAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgwriterFlushAfter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetBgwriterLruMaxpages(val *float64) {
	if err := j.validateSetBgwriterLruMaxpagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgwriterLruMaxpages",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetBgwriterLruMultiplier(val *float64) {
	if err := j.validateSetBgwriterLruMultiplierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgwriterLruMultiplier",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetDeadlockTimeout(val *float64) {
	if err := j.validateSetDeadlockTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deadlockTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetDefaultToastCompression(val *string) {
	if err := j.validateSetDefaultToastCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultToastCompression",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetIdleInTransactionSessionTimeout(val *float64) {
	if err := j.validateSetIdleInTransactionSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleInTransactionSessionTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetInternalValue(val *ManagedDatabasePostgresqlProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetIpFilter(val *[]*string) {
	if err := j.validateSetIpFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFilter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetJit(val interface{}) {
	if err := j.validateSetJitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jit",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetLogAutovacuumMinDuration(val *float64) {
	if err := j.validateSetLogAutovacuumMinDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAutovacuumMinDuration",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetLogErrorVerbosity(val *string) {
	if err := j.validateSetLogErrorVerbosityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logErrorVerbosity",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetLogLinePrefix(val *string) {
	if err := j.validateSetLogLinePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLinePrefix",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetLogMinDurationStatement(val *float64) {
	if err := j.validateSetLogMinDurationStatementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logMinDurationStatement",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetLogTempFiles(val *float64) {
	if err := j.validateSetLogTempFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logTempFiles",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxFilesPerProcess(val *float64) {
	if err := j.validateSetMaxFilesPerProcessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFilesPerProcess",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxLocksPerTransaction(val *float64) {
	if err := j.validateSetMaxLocksPerTransactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLocksPerTransaction",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxLogicalReplicationWorkers(val *float64) {
	if err := j.validateSetMaxLogicalReplicationWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLogicalReplicationWorkers",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxParallelWorkers(val *float64) {
	if err := j.validateSetMaxParallelWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelWorkers",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxParallelWorkersPerGather(val *float64) {
	if err := j.validateSetMaxParallelWorkersPerGatherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelWorkersPerGather",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxPredLocksPerTransaction(val *float64) {
	if err := j.validateSetMaxPredLocksPerTransactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPredLocksPerTransaction",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxPreparedTransactions(val *float64) {
	if err := j.validateSetMaxPreparedTransactionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPreparedTransactions",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxReplicationSlots(val *float64) {
	if err := j.validateSetMaxReplicationSlotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxReplicationSlots",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxSlotWalKeepSize(val *float64) {
	if err := j.validateSetMaxSlotWalKeepSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSlotWalKeepSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxStackDepth(val *float64) {
	if err := j.validateSetMaxStackDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStackDepth",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxStandbyArchiveDelay(val *float64) {
	if err := j.validateSetMaxStandbyArchiveDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStandbyArchiveDelay",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxStandbyStreamingDelay(val *float64) {
	if err := j.validateSetMaxStandbyStreamingDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStandbyStreamingDelay",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxWalSenders(val *float64) {
	if err := j.validateSetMaxWalSendersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWalSenders",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetMaxWorkerProcesses(val *float64) {
	if err := j.validateSetMaxWorkerProcessesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWorkerProcesses",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetPgPartmanBgwInterval(val *float64) {
	if err := j.validateSetPgPartmanBgwIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgPartmanBgwInterval",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetPgPartmanBgwRole(val *string) {
	if err := j.validateSetPgPartmanBgwRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgPartmanBgwRole",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetPgReadReplica(val interface{}) {
	if err := j.validateSetPgReadReplicaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgReadReplica",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetPgServiceToForkFrom(val *string) {
	if err := j.validateSetPgServiceToForkFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgServiceToForkFrom",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetPgStatMonitorEnable(val interface{}) {
	if err := j.validateSetPgStatMonitorEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgStatMonitorEnable",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetPgStatMonitorPgsmEnableQueryPlan(val interface{}) {
	if err := j.validateSetPgStatMonitorPgsmEnableQueryPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgStatMonitorPgsmEnableQueryPlan",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetPgStatMonitorPgsmMaxBuckets(val *float64) {
	if err := j.validateSetPgStatMonitorPgsmMaxBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgStatMonitorPgsmMaxBuckets",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetPgStatStatementsTrack(val *string) {
	if err := j.validateSetPgStatStatementsTrackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pgStatStatementsTrack",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetPublicAccess(val interface{}) {
	if err := j.validateSetPublicAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAccess",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetSharedBuffersPercentage(val *float64) {
	if err := j.validateSetSharedBuffersPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedBuffersPercentage",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetSynchronousReplication(val *string) {
	if err := j.validateSetSynchronousReplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"synchronousReplication",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetTempFileLimit(val *float64) {
	if err := j.validateSetTempFileLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tempFileLimit",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetTrackActivityQuerySize(val *float64) {
	if err := j.validateSetTrackActivityQuerySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackActivityQuerySize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetTrackCommitTimestamp(val *string) {
	if err := j.validateSetTrackCommitTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackCommitTimestamp",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetTrackFunctions(val *string) {
	if err := j.validateSetTrackFunctionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackFunctions",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetTrackIoTiming(val *string) {
	if err := j.validateSetTrackIoTimingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackIoTiming",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetVariant(val *string) {
	if err := j.validateSetVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variant",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetWalSenderTimeout(val *float64) {
	if err := j.validateSetWalSenderTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"walSenderTimeout",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetWalWriterDelay(val *float64) {
	if err := j.validateSetWalWriterDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"walWriterDelay",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference)SetWorkMem(val *float64) {
	if err := j.validateSetWorkMemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workMem",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PutMigration(value *ManagedDatabasePostgresqlPropertiesMigration) {
	if err := m.validatePutMigrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMigration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PutPgbouncer(value *ManagedDatabasePostgresqlPropertiesPgbouncer) {
	if err := m.validatePutPgbouncerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPgbouncer",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PutPglookout(value *ManagedDatabasePostgresqlPropertiesPglookout) {
	if err := m.validatePutPglookoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPglookout",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) PutTimescaledb(value *ManagedDatabasePostgresqlPropertiesTimescaledb) {
	if err := m.validatePutTimescaledbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimescaledb",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAdminUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetAdminUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAutomaticUtilityNetworkIpFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetAutomaticUtilityNetworkIpFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAutovacuumAnalyzeScaleFactor() {
	_jsii_.InvokeVoid(
		m,
		"resetAutovacuumAnalyzeScaleFactor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAutovacuumAnalyzeThreshold() {
	_jsii_.InvokeVoid(
		m,
		"resetAutovacuumAnalyzeThreshold",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAutovacuumFreezeMaxAge() {
	_jsii_.InvokeVoid(
		m,
		"resetAutovacuumFreezeMaxAge",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAutovacuumMaxWorkers() {
	_jsii_.InvokeVoid(
		m,
		"resetAutovacuumMaxWorkers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAutovacuumNaptime() {
	_jsii_.InvokeVoid(
		m,
		"resetAutovacuumNaptime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAutovacuumVacuumCostDelay() {
	_jsii_.InvokeVoid(
		m,
		"resetAutovacuumVacuumCostDelay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAutovacuumVacuumCostLimit() {
	_jsii_.InvokeVoid(
		m,
		"resetAutovacuumVacuumCostLimit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAutovacuumVacuumScaleFactor() {
	_jsii_.InvokeVoid(
		m,
		"resetAutovacuumVacuumScaleFactor",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetAutovacuumVacuumThreshold() {
	_jsii_.InvokeVoid(
		m,
		"resetAutovacuumVacuumThreshold",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetBackupHour() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupHour",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetBackupMinute() {
	_jsii_.InvokeVoid(
		m,
		"resetBackupMinute",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetBgwriterDelay() {
	_jsii_.InvokeVoid(
		m,
		"resetBgwriterDelay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetBgwriterFlushAfter() {
	_jsii_.InvokeVoid(
		m,
		"resetBgwriterFlushAfter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetBgwriterLruMaxpages() {
	_jsii_.InvokeVoid(
		m,
		"resetBgwriterLruMaxpages",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetBgwriterLruMultiplier() {
	_jsii_.InvokeVoid(
		m,
		"resetBgwriterLruMultiplier",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetDeadlockTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetDeadlockTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetDefaultToastCompression() {
	_jsii_.InvokeVoid(
		m,
		"resetDefaultToastCompression",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetIdleInTransactionSessionTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetIdleInTransactionSessionTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetIpFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetIpFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetJit() {
	_jsii_.InvokeVoid(
		m,
		"resetJit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetLogAutovacuumMinDuration() {
	_jsii_.InvokeVoid(
		m,
		"resetLogAutovacuumMinDuration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetLogErrorVerbosity() {
	_jsii_.InvokeVoid(
		m,
		"resetLogErrorVerbosity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetLogLinePrefix() {
	_jsii_.InvokeVoid(
		m,
		"resetLogLinePrefix",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetLogMinDurationStatement() {
	_jsii_.InvokeVoid(
		m,
		"resetLogMinDurationStatement",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetLogTempFiles() {
	_jsii_.InvokeVoid(
		m,
		"resetLogTempFiles",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxFilesPerProcess() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxFilesPerProcess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxLocksPerTransaction() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxLocksPerTransaction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxLogicalReplicationWorkers() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxLogicalReplicationWorkers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxParallelWorkers() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxParallelWorkers",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxParallelWorkersPerGather() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxParallelWorkersPerGather",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxPredLocksPerTransaction() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxPredLocksPerTransaction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxPreparedTransactions() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxPreparedTransactions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxReplicationSlots() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxReplicationSlots",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxSlotWalKeepSize() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxSlotWalKeepSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxStackDepth() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxStackDepth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxStandbyArchiveDelay() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxStandbyArchiveDelay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxStandbyStreamingDelay() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxStandbyStreamingDelay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxWalSenders() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxWalSenders",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMaxWorkerProcesses() {
	_jsii_.InvokeVoid(
		m,
		"resetMaxWorkerProcesses",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetMigration() {
	_jsii_.InvokeVoid(
		m,
		"resetMigration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetPgbouncer() {
	_jsii_.InvokeVoid(
		m,
		"resetPgbouncer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetPglookout() {
	_jsii_.InvokeVoid(
		m,
		"resetPglookout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetPgPartmanBgwInterval() {
	_jsii_.InvokeVoid(
		m,
		"resetPgPartmanBgwInterval",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetPgPartmanBgwRole() {
	_jsii_.InvokeVoid(
		m,
		"resetPgPartmanBgwRole",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetPgReadReplica() {
	_jsii_.InvokeVoid(
		m,
		"resetPgReadReplica",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetPgServiceToForkFrom() {
	_jsii_.InvokeVoid(
		m,
		"resetPgServiceToForkFrom",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetPgStatMonitorEnable() {
	_jsii_.InvokeVoid(
		m,
		"resetPgStatMonitorEnable",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetPgStatMonitorPgsmEnableQueryPlan() {
	_jsii_.InvokeVoid(
		m,
		"resetPgStatMonitorPgsmEnableQueryPlan",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetPgStatMonitorPgsmMaxBuckets() {
	_jsii_.InvokeVoid(
		m,
		"resetPgStatMonitorPgsmMaxBuckets",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetPgStatStatementsTrack() {
	_jsii_.InvokeVoid(
		m,
		"resetPgStatStatementsTrack",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetPublicAccess() {
	_jsii_.InvokeVoid(
		m,
		"resetPublicAccess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetSharedBuffersPercentage() {
	_jsii_.InvokeVoid(
		m,
		"resetSharedBuffersPercentage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetSynchronousReplication() {
	_jsii_.InvokeVoid(
		m,
		"resetSynchronousReplication",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetTempFileLimit() {
	_jsii_.InvokeVoid(
		m,
		"resetTempFileLimit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetTimescaledb() {
	_jsii_.InvokeVoid(
		m,
		"resetTimescaledb",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetTimezone() {
	_jsii_.InvokeVoid(
		m,
		"resetTimezone",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetTrackActivityQuerySize() {
	_jsii_.InvokeVoid(
		m,
		"resetTrackActivityQuerySize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetTrackCommitTimestamp() {
	_jsii_.InvokeVoid(
		m,
		"resetTrackCommitTimestamp",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetTrackFunctions() {
	_jsii_.InvokeVoid(
		m,
		"resetTrackFunctions",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetTrackIoTiming() {
	_jsii_.InvokeVoid(
		m,
		"resetTrackIoTiming",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetVariant() {
	_jsii_.InvokeVoid(
		m,
		"resetVariant",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetWalSenderTimeout() {
	_jsii_.InvokeVoid(
		m,
		"resetWalSenderTimeout",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetWalWriterDelay() {
	_jsii_.InvokeVoid(
		m,
		"resetWalWriterDelay",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ResetWorkMem() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkMem",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ManagedDatabasePostgresqlPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

