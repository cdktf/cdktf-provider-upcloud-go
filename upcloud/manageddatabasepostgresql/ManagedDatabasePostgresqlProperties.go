// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasepostgresql


type ManagedDatabasePostgresqlProperties struct {
	// Custom password for admin user.
	//
	// Defaults to random string. This must be set only when a new service is being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#admin_password ManagedDatabasePostgresql#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Custom username for admin user. This must be set only when a new service is being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#admin_username ManagedDatabasePostgresql#admin_username}
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// Automatic utility network IP Filter. Automatically allow connections from servers in the utility network within the same zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#automatic_utility_network_ip_filter ManagedDatabasePostgresql#automatic_utility_network_ip_filter}
	AutomaticUtilityNetworkIpFilter interface{} `field:"optional" json:"automaticUtilityNetworkIpFilter" yaml:"automaticUtilityNetworkIpFilter"`
	// Specifies a fraction of the table size to add to autovacuum_analyze_threshold when deciding whether to trigger an ANALYZE.
	//
	// The default is 0.2 (20% of table size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#autovacuum_analyze_scale_factor ManagedDatabasePostgresql#autovacuum_analyze_scale_factor}
	AutovacuumAnalyzeScaleFactor *float64 `field:"optional" json:"autovacuumAnalyzeScaleFactor" yaml:"autovacuumAnalyzeScaleFactor"`
	// Specifies the minimum number of inserted, updated or deleted tuples needed to trigger an ANALYZE in any one table.
	//
	// The default is 50 tuples.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#autovacuum_analyze_threshold ManagedDatabasePostgresql#autovacuum_analyze_threshold}
	AutovacuumAnalyzeThreshold *float64 `field:"optional" json:"autovacuumAnalyzeThreshold" yaml:"autovacuumAnalyzeThreshold"`
	// Specifies the maximum age (in transactions) that a table's pg_class.relfrozenxid field can attain before a VACUUM operation is forced to prevent transaction ID wraparound within the table. Note that the system will launch autovacuum processes to prevent wraparound even when autovacuum is otherwise disabled. This parameter will cause the server to be restarted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#autovacuum_freeze_max_age ManagedDatabasePostgresql#autovacuum_freeze_max_age}
	AutovacuumFreezeMaxAge *float64 `field:"optional" json:"autovacuumFreezeMaxAge" yaml:"autovacuumFreezeMaxAge"`
	// Specifies the maximum number of autovacuum processes (other than the autovacuum launcher) that may be running at any one time.
	//
	// The default is three. This parameter can only be set at server start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#autovacuum_max_workers ManagedDatabasePostgresql#autovacuum_max_workers}
	AutovacuumMaxWorkers *float64 `field:"optional" json:"autovacuumMaxWorkers" yaml:"autovacuumMaxWorkers"`
	// Specifies the minimum delay between autovacuum runs on any given database.
	//
	// The delay is measured in seconds, and the default is one minute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#autovacuum_naptime ManagedDatabasePostgresql#autovacuum_naptime}
	AutovacuumNaptime *float64 `field:"optional" json:"autovacuumNaptime" yaml:"autovacuumNaptime"`
	// Specifies the cost delay value that will be used in automatic VACUUM operations.
	//
	// If -1 is specified, the regular vacuum_cost_delay value will be used. The default value is 20 milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#autovacuum_vacuum_cost_delay ManagedDatabasePostgresql#autovacuum_vacuum_cost_delay}
	AutovacuumVacuumCostDelay *float64 `field:"optional" json:"autovacuumVacuumCostDelay" yaml:"autovacuumVacuumCostDelay"`
	// Specifies the cost limit value that will be used in automatic VACUUM operations.
	//
	// If -1 is specified (which is the default), the regular vacuum_cost_limit value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#autovacuum_vacuum_cost_limit ManagedDatabasePostgresql#autovacuum_vacuum_cost_limit}
	AutovacuumVacuumCostLimit *float64 `field:"optional" json:"autovacuumVacuumCostLimit" yaml:"autovacuumVacuumCostLimit"`
	// Specifies a fraction of the table size to add to autovacuum_vacuum_threshold when deciding whether to trigger a VACUUM.
	//
	// The default is 0.2 (20% of table size).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#autovacuum_vacuum_scale_factor ManagedDatabasePostgresql#autovacuum_vacuum_scale_factor}
	AutovacuumVacuumScaleFactor *float64 `field:"optional" json:"autovacuumVacuumScaleFactor" yaml:"autovacuumVacuumScaleFactor"`
	// Specifies the minimum number of updated or deleted tuples needed to trigger a VACUUM in any one table.
	//
	// The default is 50 tuples.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#autovacuum_vacuum_threshold ManagedDatabasePostgresql#autovacuum_vacuum_threshold}
	AutovacuumVacuumThreshold *float64 `field:"optional" json:"autovacuumVacuumThreshold" yaml:"autovacuumVacuumThreshold"`
	// The hour of day (in UTC) when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#backup_hour ManagedDatabasePostgresql#backup_hour}
	BackupHour *float64 `field:"optional" json:"backupHour" yaml:"backupHour"`
	// The minute of an hour when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#backup_minute ManagedDatabasePostgresql#backup_minute}
	BackupMinute *float64 `field:"optional" json:"backupMinute" yaml:"backupMinute"`
	// Specifies the delay between activity rounds for the background writer in milliseconds. Default is 200.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#bgwriter_delay ManagedDatabasePostgresql#bgwriter_delay}
	BgwriterDelay *float64 `field:"optional" json:"bgwriterDelay" yaml:"bgwriterDelay"`
	// Whenever more than bgwriter_flush_after bytes have been written by the background writer, attempt to force the OS to issue these writes to the underlying storage.
	//
	// Specified in kilobytes, default is 512. Setting of 0 disables forced writeback.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#bgwriter_flush_after ManagedDatabasePostgresql#bgwriter_flush_after}
	BgwriterFlushAfter *float64 `field:"optional" json:"bgwriterFlushAfter" yaml:"bgwriterFlushAfter"`
	// In each round, no more than this many buffers will be written by the background writer.
	//
	// Setting this to zero disables background writing. Default is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#bgwriter_lru_maxpages ManagedDatabasePostgresql#bgwriter_lru_maxpages}
	BgwriterLruMaxpages *float64 `field:"optional" json:"bgwriterLruMaxpages" yaml:"bgwriterLruMaxpages"`
	// The average recent need for new buffers is multiplied by bgwriter_lru_multiplier to arrive at an estimate of the number that will be needed during the next round, (up to bgwriter_lru_maxpages).
	//
	// 1.0 represents a “just in time” policy of writing exactly the number of buffers predicted to be needed. Larger values provide some cushion against spikes in demand, while smaller values intentionally leave writes to be done by server processes. The default is 2.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#bgwriter_lru_multiplier ManagedDatabasePostgresql#bgwriter_lru_multiplier}
	BgwriterLruMultiplier *float64 `field:"optional" json:"bgwriterLruMultiplier" yaml:"bgwriterLruMultiplier"`
	// This is the amount of time, in milliseconds, to wait on a lock before checking to see if there is a deadlock condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#deadlock_timeout ManagedDatabasePostgresql#deadlock_timeout}
	DeadlockTimeout *float64 `field:"optional" json:"deadlockTimeout" yaml:"deadlockTimeout"`
	// Specifies the default TOAST compression method for values of compressible columns (the default is lz4).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#default_toast_compression ManagedDatabasePostgresql#default_toast_compression}
	DefaultToastCompression *string `field:"optional" json:"defaultToastCompression" yaml:"defaultToastCompression"`
	// Time out sessions with open transactions after this number of milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#idle_in_transaction_session_timeout ManagedDatabasePostgresql#idle_in_transaction_session_timeout}
	IdleInTransactionSessionTimeout *float64 `field:"optional" json:"idleInTransactionSessionTimeout" yaml:"idleInTransactionSessionTimeout"`
	// IP filter. Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#ip_filter ManagedDatabasePostgresql#ip_filter}
	IpFilter *[]*string `field:"optional" json:"ipFilter" yaml:"ipFilter"`
	// Controls system-wide use of Just-in-Time Compilation (JIT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#jit ManagedDatabasePostgresql#jit}
	Jit interface{} `field:"optional" json:"jit" yaml:"jit"`
	// Causes each action executed by autovacuum to be logged if it ran for at least the specified number of milliseconds.
	//
	// Setting this to zero logs all autovacuum actions. Minus-one (the default) disables logging autovacuum actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#log_autovacuum_min_duration ManagedDatabasePostgresql#log_autovacuum_min_duration}
	LogAutovacuumMinDuration *float64 `field:"optional" json:"logAutovacuumMinDuration" yaml:"logAutovacuumMinDuration"`
	// Controls the amount of detail written in the server log for each message that is logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#log_error_verbosity ManagedDatabasePostgresql#log_error_verbosity}
	LogErrorVerbosity *string `field:"optional" json:"logErrorVerbosity" yaml:"logErrorVerbosity"`
	// Choose from one of the available log formats.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#log_line_prefix ManagedDatabasePostgresql#log_line_prefix}
	LogLinePrefix *string `field:"optional" json:"logLinePrefix" yaml:"logLinePrefix"`
	// Log statements that take more than this number of milliseconds to run, -1 disables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#log_min_duration_statement ManagedDatabasePostgresql#log_min_duration_statement}
	LogMinDurationStatement *float64 `field:"optional" json:"logMinDurationStatement" yaml:"logMinDurationStatement"`
	// Log statements for each temporary file created larger than this number of kilobytes, -1 disables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#log_temp_files ManagedDatabasePostgresql#log_temp_files}
	LogTempFiles *float64 `field:"optional" json:"logTempFiles" yaml:"logTempFiles"`
	// PostgreSQL maximum number of files that can be open per process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_files_per_process ManagedDatabasePostgresql#max_files_per_process}
	MaxFilesPerProcess *float64 `field:"optional" json:"maxFilesPerProcess" yaml:"maxFilesPerProcess"`
	// PostgreSQL maximum locks per transaction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_locks_per_transaction ManagedDatabasePostgresql#max_locks_per_transaction}
	MaxLocksPerTransaction *float64 `field:"optional" json:"maxLocksPerTransaction" yaml:"maxLocksPerTransaction"`
	// PostgreSQL maximum logical replication workers (taken from the pool of max_parallel_workers).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_logical_replication_workers ManagedDatabasePostgresql#max_logical_replication_workers}
	MaxLogicalReplicationWorkers *float64 `field:"optional" json:"maxLogicalReplicationWorkers" yaml:"maxLogicalReplicationWorkers"`
	// Sets the maximum number of workers that the system can support for parallel queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_parallel_workers ManagedDatabasePostgresql#max_parallel_workers}
	MaxParallelWorkers *float64 `field:"optional" json:"maxParallelWorkers" yaml:"maxParallelWorkers"`
	// Sets the maximum number of workers that can be started by a single Gather or Gather Merge node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_parallel_workers_per_gather ManagedDatabasePostgresql#max_parallel_workers_per_gather}
	MaxParallelWorkersPerGather *float64 `field:"optional" json:"maxParallelWorkersPerGather" yaml:"maxParallelWorkersPerGather"`
	// PostgreSQL maximum predicate locks per transaction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_pred_locks_per_transaction ManagedDatabasePostgresql#max_pred_locks_per_transaction}
	MaxPredLocksPerTransaction *float64 `field:"optional" json:"maxPredLocksPerTransaction" yaml:"maxPredLocksPerTransaction"`
	// PostgreSQL maximum prepared transactions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_prepared_transactions ManagedDatabasePostgresql#max_prepared_transactions}
	MaxPreparedTransactions *float64 `field:"optional" json:"maxPreparedTransactions" yaml:"maxPreparedTransactions"`
	// PostgreSQL maximum replication slots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_replication_slots ManagedDatabasePostgresql#max_replication_slots}
	MaxReplicationSlots *float64 `field:"optional" json:"maxReplicationSlots" yaml:"maxReplicationSlots"`
	// PostgreSQL maximum WAL size (MB) reserved for replication slots.
	//
	// Default is -1 (unlimited). wal_keep_size minimum WAL size setting takes precedence over this.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_slot_wal_keep_size ManagedDatabasePostgresql#max_slot_wal_keep_size}
	MaxSlotWalKeepSize *float64 `field:"optional" json:"maxSlotWalKeepSize" yaml:"maxSlotWalKeepSize"`
	// Maximum depth of the stack in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_stack_depth ManagedDatabasePostgresql#max_stack_depth}
	MaxStackDepth *float64 `field:"optional" json:"maxStackDepth" yaml:"maxStackDepth"`
	// Max standby archive delay in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_standby_archive_delay ManagedDatabasePostgresql#max_standby_archive_delay}
	MaxStandbyArchiveDelay *float64 `field:"optional" json:"maxStandbyArchiveDelay" yaml:"maxStandbyArchiveDelay"`
	// Max standby streaming delay in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_standby_streaming_delay ManagedDatabasePostgresql#max_standby_streaming_delay}
	MaxStandbyStreamingDelay *float64 `field:"optional" json:"maxStandbyStreamingDelay" yaml:"maxStandbyStreamingDelay"`
	// PostgreSQL maximum WAL senders.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_wal_senders ManagedDatabasePostgresql#max_wal_senders}
	MaxWalSenders *float64 `field:"optional" json:"maxWalSenders" yaml:"maxWalSenders"`
	// Sets the maximum number of background processes that the system can support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#max_worker_processes ManagedDatabasePostgresql#max_worker_processes}
	MaxWorkerProcesses *float64 `field:"optional" json:"maxWorkerProcesses" yaml:"maxWorkerProcesses"`
	// migration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#migration ManagedDatabasePostgresql#migration}
	Migration *ManagedDatabasePostgresqlPropertiesMigration `field:"optional" json:"migration" yaml:"migration"`
	// Chooses the algorithm for encrypting passwords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#password_encryption ManagedDatabasePostgresql#password_encryption}
	PasswordEncryption *string `field:"optional" json:"passwordEncryption" yaml:"passwordEncryption"`
	// pgaudit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#pgaudit ManagedDatabasePostgresql#pgaudit}
	Pgaudit *ManagedDatabasePostgresqlPropertiesPgaudit `field:"optional" json:"pgaudit" yaml:"pgaudit"`
	// pgbouncer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#pgbouncer ManagedDatabasePostgresql#pgbouncer}
	Pgbouncer *ManagedDatabasePostgresqlPropertiesPgbouncer `field:"optional" json:"pgbouncer" yaml:"pgbouncer"`
	// pglookout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#pglookout ManagedDatabasePostgresql#pglookout}
	Pglookout *ManagedDatabasePostgresqlPropertiesPglookout `field:"optional" json:"pglookout" yaml:"pglookout"`
	// Sets the time interval to run pg_partman's scheduled tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#pg_partman_bgw_interval ManagedDatabasePostgresql#pg_partman_bgw_interval}
	PgPartmanBgwInterval *float64 `field:"optional" json:"pgPartmanBgwInterval" yaml:"pgPartmanBgwInterval"`
	// Controls which role to use for pg_partman's scheduled background tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#pg_partman_bgw_role ManagedDatabasePostgresql#pg_partman_bgw_role}
	PgPartmanBgwRole *string `field:"optional" json:"pgPartmanBgwRole" yaml:"pgPartmanBgwRole"`
	// Enable pg_stat_monitor extension if available for the current cluster.
	//
	// Enable the pg_stat_monitor extension. Enabling this extension will cause the cluster to be restarted.When this extension is enabled, pg_stat_statements results for utility commands are unreliable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#pg_stat_monitor_enable ManagedDatabasePostgresql#pg_stat_monitor_enable}
	PgStatMonitorEnable interface{} `field:"optional" json:"pgStatMonitorEnable" yaml:"pgStatMonitorEnable"`
	// Enables or disables query plan monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#pg_stat_monitor_pgsm_enable_query_plan ManagedDatabasePostgresql#pg_stat_monitor_pgsm_enable_query_plan}
	PgStatMonitorPgsmEnableQueryPlan interface{} `field:"optional" json:"pgStatMonitorPgsmEnableQueryPlan" yaml:"pgStatMonitorPgsmEnableQueryPlan"`
	// Sets the maximum number of buckets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#pg_stat_monitor_pgsm_max_buckets ManagedDatabasePostgresql#pg_stat_monitor_pgsm_max_buckets}
	PgStatMonitorPgsmMaxBuckets *float64 `field:"optional" json:"pgStatMonitorPgsmMaxBuckets" yaml:"pgStatMonitorPgsmMaxBuckets"`
	// Controls which statements are counted.
	//
	// Specify top to track top-level statements (those issued directly by clients), all to also track nested statements (such as statements invoked within functions), or none to disable statement statistics collection. The default value is top.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#pg_stat_statements_track ManagedDatabasePostgresql#pg_stat_statements_track}
	PgStatStatementsTrack *string `field:"optional" json:"pgStatStatementsTrack" yaml:"pgStatStatementsTrack"`
	// Public Access. Allow access to the service from the public Internet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#public_access ManagedDatabasePostgresql#public_access}
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// Service logging. Store logs for the service so that they are available in the HTTP API and console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#service_log ManagedDatabasePostgresql#service_log}
	ServiceLog interface{} `field:"optional" json:"serviceLog" yaml:"serviceLog"`
	// Percentage of total RAM that the database server uses for shared memory buffers.
	//
	// Valid range is 20-60 (float), which corresponds to 20% - 60%. This setting adjusts the shared_buffers configuration value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#shared_buffers_percentage ManagedDatabasePostgresql#shared_buffers_percentage}
	SharedBuffersPercentage *float64 `field:"optional" json:"sharedBuffersPercentage" yaml:"sharedBuffersPercentage"`
	// Synchronous replication type. Note that the service plan also needs to support synchronous replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#synchronous_replication ManagedDatabasePostgresql#synchronous_replication}
	SynchronousReplication *string `field:"optional" json:"synchronousReplication" yaml:"synchronousReplication"`
	// PostgreSQL temporary file limit in KiB, -1 for unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#temp_file_limit ManagedDatabasePostgresql#temp_file_limit}
	TempFileLimit *float64 `field:"optional" json:"tempFileLimit" yaml:"tempFileLimit"`
	// timescaledb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#timescaledb ManagedDatabasePostgresql#timescaledb}
	Timescaledb *ManagedDatabasePostgresqlPropertiesTimescaledb `field:"optional" json:"timescaledb" yaml:"timescaledb"`
	// PostgreSQL service timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#timezone ManagedDatabasePostgresql#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Specifies the number of bytes reserved to track the currently executing command for each active session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#track_activity_query_size ManagedDatabasePostgresql#track_activity_query_size}
	TrackActivityQuerySize *float64 `field:"optional" json:"trackActivityQuerySize" yaml:"trackActivityQuerySize"`
	// Record commit time of transactions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#track_commit_timestamp ManagedDatabasePostgresql#track_commit_timestamp}
	TrackCommitTimestamp *string `field:"optional" json:"trackCommitTimestamp" yaml:"trackCommitTimestamp"`
	// Enables tracking of function call counts and time used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#track_functions ManagedDatabasePostgresql#track_functions}
	TrackFunctions *string `field:"optional" json:"trackFunctions" yaml:"trackFunctions"`
	// Enables timing of database I/O calls.
	//
	// This parameter is off by default, because it will repeatedly query the operating system for the current time, which may cause significant overhead on some platforms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#track_io_timing ManagedDatabasePostgresql#track_io_timing}
	TrackIoTiming *string `field:"optional" json:"trackIoTiming" yaml:"trackIoTiming"`
	// Variant of the PostgreSQL service, may affect the features that are exposed by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#variant ManagedDatabasePostgresql#variant}
	Variant *string `field:"optional" json:"variant" yaml:"variant"`
	// PostgreSQL major version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#version ManagedDatabasePostgresql#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Terminate replication connections that are inactive for longer than this amount of time, in milliseconds.
	//
	// Setting this value to zero disables the timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#wal_sender_timeout ManagedDatabasePostgresql#wal_sender_timeout}
	WalSenderTimeout *float64 `field:"optional" json:"walSenderTimeout" yaml:"walSenderTimeout"`
	// WAL flush interval in milliseconds.
	//
	// Note that setting this value to lower than the default 200ms may negatively impact performance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#wal_writer_delay ManagedDatabasePostgresql#wal_writer_delay}
	WalWriterDelay *float64 `field:"optional" json:"walWriterDelay" yaml:"walWriterDelay"`
	// Sets the maximum amount of memory to be used by a query operation (such as a sort or hash table) before writing to temporary disk files, in MB.
	//
	// Default is 1MB + 0.075% of total RAM (up to 32MB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.22.1/docs/resources/managed_database_postgresql#work_mem ManagedDatabasePostgresql#work_mem}
	WorkMem *float64 `field:"optional" json:"workMem" yaml:"workMem"`
}

