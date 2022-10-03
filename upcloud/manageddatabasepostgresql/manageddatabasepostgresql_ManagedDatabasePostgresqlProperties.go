package manageddatabasepostgresql


type ManagedDatabasePostgresqlProperties struct {
	// Custom password for admin user.
	//
	// Defaults to random string. This must be set only when a new service is being created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#admin_password ManagedDatabasePostgresql#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Custom username for admin user. This must be set only when a new service is being created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#admin_username ManagedDatabasePostgresql#admin_username}
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// Automatic utility network IP Filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#automatic_utility_network_ip_filter ManagedDatabasePostgresql#automatic_utility_network_ip_filter}
	AutomaticUtilityNetworkIpFilter interface{} `field:"optional" json:"automaticUtilityNetworkIpFilter" yaml:"automaticUtilityNetworkIpFilter"`
	// autovacuum_analyze_scale_factor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#autovacuum_analyze_scale_factor ManagedDatabasePostgresql#autovacuum_analyze_scale_factor}
	AutovacuumAnalyzeScaleFactor *float64 `field:"optional" json:"autovacuumAnalyzeScaleFactor" yaml:"autovacuumAnalyzeScaleFactor"`
	// autovacuum_analyze_threshold.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#autovacuum_analyze_threshold ManagedDatabasePostgresql#autovacuum_analyze_threshold}
	AutovacuumAnalyzeThreshold *float64 `field:"optional" json:"autovacuumAnalyzeThreshold" yaml:"autovacuumAnalyzeThreshold"`
	// autovacuum_freeze_max_age.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#autovacuum_freeze_max_age ManagedDatabasePostgresql#autovacuum_freeze_max_age}
	AutovacuumFreezeMaxAge *float64 `field:"optional" json:"autovacuumFreezeMaxAge" yaml:"autovacuumFreezeMaxAge"`
	// autovacuum_max_workers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#autovacuum_max_workers ManagedDatabasePostgresql#autovacuum_max_workers}
	AutovacuumMaxWorkers *float64 `field:"optional" json:"autovacuumMaxWorkers" yaml:"autovacuumMaxWorkers"`
	// autovacuum_naptime.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#autovacuum_naptime ManagedDatabasePostgresql#autovacuum_naptime}
	AutovacuumNaptime *float64 `field:"optional" json:"autovacuumNaptime" yaml:"autovacuumNaptime"`
	// autovacuum_vacuum_cost_delay.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#autovacuum_vacuum_cost_delay ManagedDatabasePostgresql#autovacuum_vacuum_cost_delay}
	AutovacuumVacuumCostDelay *float64 `field:"optional" json:"autovacuumVacuumCostDelay" yaml:"autovacuumVacuumCostDelay"`
	// autovacuum_vacuum_cost_limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#autovacuum_vacuum_cost_limit ManagedDatabasePostgresql#autovacuum_vacuum_cost_limit}
	AutovacuumVacuumCostLimit *float64 `field:"optional" json:"autovacuumVacuumCostLimit" yaml:"autovacuumVacuumCostLimit"`
	// autovacuum_vacuum_scale_factor.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#autovacuum_vacuum_scale_factor ManagedDatabasePostgresql#autovacuum_vacuum_scale_factor}
	AutovacuumVacuumScaleFactor *float64 `field:"optional" json:"autovacuumVacuumScaleFactor" yaml:"autovacuumVacuumScaleFactor"`
	// autovacuum_vacuum_threshold.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#autovacuum_vacuum_threshold ManagedDatabasePostgresql#autovacuum_vacuum_threshold}
	AutovacuumVacuumThreshold *float64 `field:"optional" json:"autovacuumVacuumThreshold" yaml:"autovacuumVacuumThreshold"`
	// The hour of day (in UTC) when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#backup_hour ManagedDatabasePostgresql#backup_hour}
	BackupHour *float64 `field:"optional" json:"backupHour" yaml:"backupHour"`
	// The minute of an hour when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#backup_minute ManagedDatabasePostgresql#backup_minute}
	BackupMinute *float64 `field:"optional" json:"backupMinute" yaml:"backupMinute"`
	// bgwriter_delay.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#bgwriter_delay ManagedDatabasePostgresql#bgwriter_delay}
	BgwriterDelay *float64 `field:"optional" json:"bgwriterDelay" yaml:"bgwriterDelay"`
	// bgwriter_flush_after.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#bgwriter_flush_after ManagedDatabasePostgresql#bgwriter_flush_after}
	BgwriterFlushAfter *float64 `field:"optional" json:"bgwriterFlushAfter" yaml:"bgwriterFlushAfter"`
	// bgwriter_lru_maxpages.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#bgwriter_lru_maxpages ManagedDatabasePostgresql#bgwriter_lru_maxpages}
	BgwriterLruMaxpages *float64 `field:"optional" json:"bgwriterLruMaxpages" yaml:"bgwriterLruMaxpages"`
	// bgwriter_lru_multiplier.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#bgwriter_lru_multiplier ManagedDatabasePostgresql#bgwriter_lru_multiplier}
	BgwriterLruMultiplier *float64 `field:"optional" json:"bgwriterLruMultiplier" yaml:"bgwriterLruMultiplier"`
	// deadlock_timeout.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#deadlock_timeout ManagedDatabasePostgresql#deadlock_timeout}
	DeadlockTimeout *float64 `field:"optional" json:"deadlockTimeout" yaml:"deadlockTimeout"`
	// idle_in_transaction_session_timeout.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#idle_in_transaction_session_timeout ManagedDatabasePostgresql#idle_in_transaction_session_timeout}
	IdleInTransactionSessionTimeout *float64 `field:"optional" json:"idleInTransactionSessionTimeout" yaml:"idleInTransactionSessionTimeout"`
	// IP filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#ip_filter ManagedDatabasePostgresql#ip_filter}
	IpFilter *[]*string `field:"optional" json:"ipFilter" yaml:"ipFilter"`
	// jit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#jit ManagedDatabasePostgresql#jit}
	Jit interface{} `field:"optional" json:"jit" yaml:"jit"`
	// log_autovacuum_min_duration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#log_autovacuum_min_duration ManagedDatabasePostgresql#log_autovacuum_min_duration}
	LogAutovacuumMinDuration *float64 `field:"optional" json:"logAutovacuumMinDuration" yaml:"logAutovacuumMinDuration"`
	// log_error_verbosity.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#log_error_verbosity ManagedDatabasePostgresql#log_error_verbosity}
	LogErrorVerbosity *string `field:"optional" json:"logErrorVerbosity" yaml:"logErrorVerbosity"`
	// log_line_prefix.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#log_line_prefix ManagedDatabasePostgresql#log_line_prefix}
	LogLinePrefix *string `field:"optional" json:"logLinePrefix" yaml:"logLinePrefix"`
	// log_min_duration_statement.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#log_min_duration_statement ManagedDatabasePostgresql#log_min_duration_statement}
	LogMinDurationStatement *float64 `field:"optional" json:"logMinDurationStatement" yaml:"logMinDurationStatement"`
	// max_files_per_process.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_files_per_process ManagedDatabasePostgresql#max_files_per_process}
	MaxFilesPerProcess *float64 `field:"optional" json:"maxFilesPerProcess" yaml:"maxFilesPerProcess"`
	// max_locks_per_transaction.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_locks_per_transaction ManagedDatabasePostgresql#max_locks_per_transaction}
	MaxLocksPerTransaction *float64 `field:"optional" json:"maxLocksPerTransaction" yaml:"maxLocksPerTransaction"`
	// max_logical_replication_workers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_logical_replication_workers ManagedDatabasePostgresql#max_logical_replication_workers}
	MaxLogicalReplicationWorkers *float64 `field:"optional" json:"maxLogicalReplicationWorkers" yaml:"maxLogicalReplicationWorkers"`
	// max_parallel_workers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_parallel_workers ManagedDatabasePostgresql#max_parallel_workers}
	MaxParallelWorkers *float64 `field:"optional" json:"maxParallelWorkers" yaml:"maxParallelWorkers"`
	// max_parallel_workers_per_gather.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_parallel_workers_per_gather ManagedDatabasePostgresql#max_parallel_workers_per_gather}
	MaxParallelWorkersPerGather *float64 `field:"optional" json:"maxParallelWorkersPerGather" yaml:"maxParallelWorkersPerGather"`
	// max_pred_locks_per_transaction.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_pred_locks_per_transaction ManagedDatabasePostgresql#max_pred_locks_per_transaction}
	MaxPredLocksPerTransaction *float64 `field:"optional" json:"maxPredLocksPerTransaction" yaml:"maxPredLocksPerTransaction"`
	// max_prepared_transactions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_prepared_transactions ManagedDatabasePostgresql#max_prepared_transactions}
	MaxPreparedTransactions *float64 `field:"optional" json:"maxPreparedTransactions" yaml:"maxPreparedTransactions"`
	// max_replication_slots.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_replication_slots ManagedDatabasePostgresql#max_replication_slots}
	MaxReplicationSlots *float64 `field:"optional" json:"maxReplicationSlots" yaml:"maxReplicationSlots"`
	// max_stack_depth.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_stack_depth ManagedDatabasePostgresql#max_stack_depth}
	MaxStackDepth *float64 `field:"optional" json:"maxStackDepth" yaml:"maxStackDepth"`
	// max_standby_archive_delay.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_standby_archive_delay ManagedDatabasePostgresql#max_standby_archive_delay}
	MaxStandbyArchiveDelay *float64 `field:"optional" json:"maxStandbyArchiveDelay" yaml:"maxStandbyArchiveDelay"`
	// max_standby_streaming_delay.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_standby_streaming_delay ManagedDatabasePostgresql#max_standby_streaming_delay}
	MaxStandbyStreamingDelay *float64 `field:"optional" json:"maxStandbyStreamingDelay" yaml:"maxStandbyStreamingDelay"`
	// max_wal_senders.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_wal_senders ManagedDatabasePostgresql#max_wal_senders}
	MaxWalSenders *float64 `field:"optional" json:"maxWalSenders" yaml:"maxWalSenders"`
	// max_worker_processes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_worker_processes ManagedDatabasePostgresql#max_worker_processes}
	MaxWorkerProcesses *float64 `field:"optional" json:"maxWorkerProcesses" yaml:"maxWorkerProcesses"`
	// migration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#migration ManagedDatabasePostgresql#migration}
	Migration *ManagedDatabasePostgresqlPropertiesMigration `field:"optional" json:"migration" yaml:"migration"`
	// pgbouncer block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#pgbouncer ManagedDatabasePostgresql#pgbouncer}
	Pgbouncer *ManagedDatabasePostgresqlPropertiesPgbouncer `field:"optional" json:"pgbouncer" yaml:"pgbouncer"`
	// pglookout block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#pglookout ManagedDatabasePostgresql#pglookout}
	Pglookout *ManagedDatabasePostgresqlPropertiesPglookout `field:"optional" json:"pglookout" yaml:"pglookout"`
	// pg_partman_bgw.interval.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#pg_partman_bgw_interval ManagedDatabasePostgresql#pg_partman_bgw_interval}
	PgPartmanBgwInterval *float64 `field:"optional" json:"pgPartmanBgwInterval" yaml:"pgPartmanBgwInterval"`
	// pg_partman_bgw.role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#pg_partman_bgw_role ManagedDatabasePostgresql#pg_partman_bgw_role}
	PgPartmanBgwRole *string `field:"optional" json:"pgPartmanBgwRole" yaml:"pgPartmanBgwRole"`
	// Should the service which is being forked be a read replica.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#pg_read_replica ManagedDatabasePostgresql#pg_read_replica}
	PgReadReplica interface{} `field:"optional" json:"pgReadReplica" yaml:"pgReadReplica"`
	// Name of the PG Service from which to fork (deprecated, use service_to_fork_from).
	//
	// This has effect only when a new service is being created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#pg_service_to_fork_from ManagedDatabasePostgresql#pg_service_to_fork_from}
	PgServiceToForkFrom *string `field:"optional" json:"pgServiceToForkFrom" yaml:"pgServiceToForkFrom"`
	// pg_stat_statements.track.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#pg_stat_statements_track ManagedDatabasePostgresql#pg_stat_statements_track}
	PgStatStatementsTrack *string `field:"optional" json:"pgStatStatementsTrack" yaml:"pgStatStatementsTrack"`
	// Public Access.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#public_access ManagedDatabasePostgresql#public_access}
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// shared_buffers_percentage.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#shared_buffers_percentage ManagedDatabasePostgresql#shared_buffers_percentage}
	SharedBuffersPercentage *float64 `field:"optional" json:"sharedBuffersPercentage" yaml:"sharedBuffersPercentage"`
	// Synchronous replication type. Note that the service plan also needs to support synchronous replication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#synchronous_replication ManagedDatabasePostgresql#synchronous_replication}
	SynchronousReplication *string `field:"optional" json:"synchronousReplication" yaml:"synchronousReplication"`
	// temp_file_limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#temp_file_limit ManagedDatabasePostgresql#temp_file_limit}
	TempFileLimit *float64 `field:"optional" json:"tempFileLimit" yaml:"tempFileLimit"`
	// timescaledb block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#timescaledb ManagedDatabasePostgresql#timescaledb}
	Timescaledb *ManagedDatabasePostgresqlPropertiesTimescaledb `field:"optional" json:"timescaledb" yaml:"timescaledb"`
	// timezone.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#timezone ManagedDatabasePostgresql#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// track_activity_query_size.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#track_activity_query_size ManagedDatabasePostgresql#track_activity_query_size}
	TrackActivityQuerySize *float64 `field:"optional" json:"trackActivityQuerySize" yaml:"trackActivityQuerySize"`
	// track_commit_timestamp.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#track_commit_timestamp ManagedDatabasePostgresql#track_commit_timestamp}
	TrackCommitTimestamp *string `field:"optional" json:"trackCommitTimestamp" yaml:"trackCommitTimestamp"`
	// track_functions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#track_functions ManagedDatabasePostgresql#track_functions}
	TrackFunctions *string `field:"optional" json:"trackFunctions" yaml:"trackFunctions"`
	// track_io_timing.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#track_io_timing ManagedDatabasePostgresql#track_io_timing}
	TrackIoTiming *string `field:"optional" json:"trackIoTiming" yaml:"trackIoTiming"`
	// Variant of the PostgreSQL service, may affect the features that are exposed by default.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#variant ManagedDatabasePostgresql#variant}
	Variant *string `field:"optional" json:"variant" yaml:"variant"`
	// PostgreSQL major version.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#version ManagedDatabasePostgresql#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// wal_sender_timeout.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#wal_sender_timeout ManagedDatabasePostgresql#wal_sender_timeout}
	WalSenderTimeout *float64 `field:"optional" json:"walSenderTimeout" yaml:"walSenderTimeout"`
	// wal_writer_delay.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#wal_writer_delay ManagedDatabasePostgresql#wal_writer_delay}
	WalWriterDelay *float64 `field:"optional" json:"walWriterDelay" yaml:"walWriterDelay"`
	// work_mem.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#work_mem ManagedDatabasePostgresql#work_mem}
	WorkMem *float64 `field:"optional" json:"workMem" yaml:"workMem"`
}

