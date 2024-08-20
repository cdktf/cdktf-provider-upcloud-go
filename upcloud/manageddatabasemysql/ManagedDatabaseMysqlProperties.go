// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasemysql


type ManagedDatabaseMysqlProperties struct {
	// Custom password for admin user.
	//
	// Defaults to random string. This must be set only when a new service is being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#admin_password ManagedDatabaseMysql#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Custom username for admin user. This must be set only when a new service is being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#admin_username ManagedDatabaseMysql#admin_username}
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// Automatic utility network IP Filter. Automatically allow connections from servers in the utility network within the same zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#automatic_utility_network_ip_filter ManagedDatabaseMysql#automatic_utility_network_ip_filter}
	AutomaticUtilityNetworkIpFilter interface{} `field:"optional" json:"automaticUtilityNetworkIpFilter" yaml:"automaticUtilityNetworkIpFilter"`
	// The hour of day (in UTC) when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#backup_hour ManagedDatabaseMysql#backup_hour}
	BackupHour *float64 `field:"optional" json:"backupHour" yaml:"backupHour"`
	// The minute of an hour when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#backup_minute ManagedDatabaseMysql#backup_minute}
	BackupMinute *float64 `field:"optional" json:"backupMinute" yaml:"backupMinute"`
	// The minimum amount of time in seconds to keep binlog entries before deletion.
	//
	// This may be extended for services that require binlog entries for longer than the default for example if using the MySQL Debezium Kafka connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#binlog_retention_period ManagedDatabaseMysql#binlog_retention_period}
	BinlogRetentionPeriod *float64 `field:"optional" json:"binlogRetentionPeriod" yaml:"binlogRetentionPeriod"`
	// The number of seconds that the mysqld server waits for a connect packet before responding with Bad handshake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#connect_timeout ManagedDatabaseMysql#connect_timeout}
	ConnectTimeout *float64 `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// Default server time zone as an offset from UTC (from -12:00 to +12:00), a time zone name, or 'SYSTEM' to use the MySQL server default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#default_time_zone ManagedDatabaseMysql#default_time_zone}
	DefaultTimeZone *string `field:"optional" json:"defaultTimeZone" yaml:"defaultTimeZone"`
	// The maximum permitted result length in bytes for the GROUP_CONCAT() function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#group_concat_max_len ManagedDatabaseMysql#group_concat_max_len}
	GroupConcatMaxLen *float64 `field:"optional" json:"groupConcatMaxLen" yaml:"groupConcatMaxLen"`
	// The time, in seconds, before cached statistics expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#information_schema_stats_expiry ManagedDatabaseMysql#information_schema_stats_expiry}
	InformationSchemaStatsExpiry *float64 `field:"optional" json:"informationSchemaStatsExpiry" yaml:"informationSchemaStatsExpiry"`
	// Maximum size for the InnoDB change buffer, as a percentage of the total size of the buffer pool.
	//
	// Default is 25.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_change_buffer_max_size ManagedDatabaseMysql#innodb_change_buffer_max_size}
	InnodbChangeBufferMaxSize *float64 `field:"optional" json:"innodbChangeBufferMaxSize" yaml:"innodbChangeBufferMaxSize"`
	// Specifies whether flushing a page from the InnoDB buffer pool also flushes other dirty pages in the same extent (default is 1): 0 - dirty pages in the same extent are not flushed, 1 - flush contiguous dirty pages in the same extent, 2 - flush dirty pages in the same extent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_flush_neighbors ManagedDatabaseMysql#innodb_flush_neighbors}
	InnodbFlushNeighbors *float64 `field:"optional" json:"innodbFlushNeighbors" yaml:"innodbFlushNeighbors"`
	// Minimum length of words that are stored in an InnoDB FULLTEXT index.
	//
	// Changing this parameter will lead to a restart of the MySQL service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_ft_min_token_size ManagedDatabaseMysql#innodb_ft_min_token_size}
	InnodbFtMinTokenSize *float64 `field:"optional" json:"innodbFtMinTokenSize" yaml:"innodbFtMinTokenSize"`
	// This option is used to specify your own InnoDB FULLTEXT index stopword list for all InnoDB tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_ft_server_stopword_table ManagedDatabaseMysql#innodb_ft_server_stopword_table}
	InnodbFtServerStopwordTable *string `field:"optional" json:"innodbFtServerStopwordTable" yaml:"innodbFtServerStopwordTable"`
	// The length of time in seconds an InnoDB transaction waits for a row lock before giving up.
	//
	// Default is 120.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_lock_wait_timeout ManagedDatabaseMysql#innodb_lock_wait_timeout}
	InnodbLockWaitTimeout *float64 `field:"optional" json:"innodbLockWaitTimeout" yaml:"innodbLockWaitTimeout"`
	// The size in bytes of the buffer that InnoDB uses to write to the log files on disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_log_buffer_size ManagedDatabaseMysql#innodb_log_buffer_size}
	InnodbLogBufferSize *float64 `field:"optional" json:"innodbLogBufferSize" yaml:"innodbLogBufferSize"`
	// The upper limit in bytes on the size of the temporary log files used during online DDL operations for InnoDB tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_online_alter_log_max_size ManagedDatabaseMysql#innodb_online_alter_log_max_size}
	InnodbOnlineAlterLogMaxSize *float64 `field:"optional" json:"innodbOnlineAlterLogMaxSize" yaml:"innodbOnlineAlterLogMaxSize"`
	// When enabled, information about all deadlocks in InnoDB user transactions is recorded in the error log. Disabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_print_all_deadlocks ManagedDatabaseMysql#innodb_print_all_deadlocks}
	InnodbPrintAllDeadlocks interface{} `field:"optional" json:"innodbPrintAllDeadlocks" yaml:"innodbPrintAllDeadlocks"`
	// The number of I/O threads for read operations in InnoDB.
	//
	// Default is 4. Changing this parameter will lead to a restart of the MySQL service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_read_io_threads ManagedDatabaseMysql#innodb_read_io_threads}
	InnodbReadIoThreads *float64 `field:"optional" json:"innodbReadIoThreads" yaml:"innodbReadIoThreads"`
	// When enabled a transaction timeout causes InnoDB to abort and roll back the entire transaction.
	//
	// Changing this parameter will lead to a restart of the MySQL service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_rollback_on_timeout ManagedDatabaseMysql#innodb_rollback_on_timeout}
	InnodbRollbackOnTimeout interface{} `field:"optional" json:"innodbRollbackOnTimeout" yaml:"innodbRollbackOnTimeout"`
	// Defines the maximum number of threads permitted inside of InnoDB. Default is 0 (infinite concurrency - no limit).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_thread_concurrency ManagedDatabaseMysql#innodb_thread_concurrency}
	InnodbThreadConcurrency *float64 `field:"optional" json:"innodbThreadConcurrency" yaml:"innodbThreadConcurrency"`
	// The number of I/O threads for write operations in InnoDB.
	//
	// Default is 4. Changing this parameter will lead to a restart of the MySQL service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#innodb_write_io_threads ManagedDatabaseMysql#innodb_write_io_threads}
	InnodbWriteIoThreads *float64 `field:"optional" json:"innodbWriteIoThreads" yaml:"innodbWriteIoThreads"`
	// The number of seconds the server waits for activity on an interactive connection before closing it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#interactive_timeout ManagedDatabaseMysql#interactive_timeout}
	InteractiveTimeout *float64 `field:"optional" json:"interactiveTimeout" yaml:"interactiveTimeout"`
	// The storage engine for in-memory internal temporary tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#internal_tmp_mem_storage_engine ManagedDatabaseMysql#internal_tmp_mem_storage_engine}
	InternalTmpMemStorageEngine *string `field:"optional" json:"internalTmpMemStorageEngine" yaml:"internalTmpMemStorageEngine"`
	// IP filter. Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#ip_filter ManagedDatabaseMysql#ip_filter}
	IpFilter *[]*string `field:"optional" json:"ipFilter" yaml:"ipFilter"`
	// The slow_query_logs work as SQL statements that take more than long_query_time seconds to execute. Default is 10s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#long_query_time ManagedDatabaseMysql#long_query_time}
	LongQueryTime *float64 `field:"optional" json:"longQueryTime" yaml:"longQueryTime"`
	// Size of the largest message in bytes that can be received by the server. Default is 67108864 (64M).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#max_allowed_packet ManagedDatabaseMysql#max_allowed_packet}
	MaxAllowedPacket *float64 `field:"optional" json:"maxAllowedPacket" yaml:"maxAllowedPacket"`
	// Limits the size of internal in-memory tables. Also set tmp_table_size. Default is 16777216 (16M).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#max_heap_table_size ManagedDatabaseMysql#max_heap_table_size}
	MaxHeapTableSize *float64 `field:"optional" json:"maxHeapTableSize" yaml:"maxHeapTableSize"`
	// migration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#migration ManagedDatabaseMysql#migration}
	Migration *ManagedDatabaseMysqlPropertiesMigration `field:"optional" json:"migration" yaml:"migration"`
	// Start sizes of connection buffer and result buffer.
	//
	// Default is 16384 (16K). Changing this parameter will lead to a restart of the MySQL service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#net_buffer_length ManagedDatabaseMysql#net_buffer_length}
	NetBufferLength *float64 `field:"optional" json:"netBufferLength" yaml:"netBufferLength"`
	// The number of seconds to wait for more data from a connection before aborting the read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#net_read_timeout ManagedDatabaseMysql#net_read_timeout}
	NetReadTimeout *float64 `field:"optional" json:"netReadTimeout" yaml:"netReadTimeout"`
	// The number of seconds to wait for a block to be written to a connection before aborting the write.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#net_write_timeout ManagedDatabaseMysql#net_write_timeout}
	NetWriteTimeout *float64 `field:"optional" json:"netWriteTimeout" yaml:"netWriteTimeout"`
	// Public Access. Allow access to the service from the public Internet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#public_access ManagedDatabaseMysql#public_access}
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// Service logging. Store logs for the service so that they are available in the HTTP API and console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#service_log ManagedDatabaseMysql#service_log}
	ServiceLog interface{} `field:"optional" json:"serviceLog" yaml:"serviceLog"`
	// Slow query log enables capturing of slow queries.
	//
	// Setting slow_query_log to false also truncates the mysql.slow_log table. Default is off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#slow_query_log ManagedDatabaseMysql#slow_query_log}
	SlowQueryLog interface{} `field:"optional" json:"slowQueryLog" yaml:"slowQueryLog"`
	// Sort buffer size in bytes for ORDER BY optimization. Default is 262144 (256K).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#sort_buffer_size ManagedDatabaseMysql#sort_buffer_size}
	SortBufferSize *float64 `field:"optional" json:"sortBufferSize" yaml:"sortBufferSize"`
	// Global SQL mode.
	//
	// Set to empty to use MySQL server defaults. When creating a new service and not setting this field Aiven default SQL mode (strict, SQL standard compliant) will be assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#sql_mode ManagedDatabaseMysql#sql_mode}
	SqlMode *string `field:"optional" json:"sqlMode" yaml:"sqlMode"`
	// Require primary key to be defined for new tables or old tables modified with ALTER TABLE and fail if missing.
	//
	// It is recommended to always have primary keys because various functionality may break if any large table is missing them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#sql_require_primary_key ManagedDatabaseMysql#sql_require_primary_key}
	SqlRequirePrimaryKey interface{} `field:"optional" json:"sqlRequirePrimaryKey" yaml:"sqlRequirePrimaryKey"`
	// Limits the size of internal in-memory tables. Also set max_heap_table_size. Default is 16777216 (16M).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#tmp_table_size ManagedDatabaseMysql#tmp_table_size}
	TmpTableSize *float64 `field:"optional" json:"tmpTableSize" yaml:"tmpTableSize"`
	// MySQL major version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#version ManagedDatabaseMysql#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The number of seconds the server waits for activity on a noninteractive connection before closing it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.0/docs/resources/managed_database_mysql#wait_timeout ManagedDatabaseMysql#wait_timeout}
	WaitTimeout *float64 `field:"optional" json:"waitTimeout" yaml:"waitTimeout"`
}

