// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type ManagedDatabaseMysqlProperties struct {
	// Custom password for admin user.
	//
	// Defaults to random string. This must be set only when a new service is being created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#admin_password ManagedDatabaseMysql#admin_password}
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Custom username for admin user. This must be set only when a new service is being created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#admin_username ManagedDatabaseMysql#admin_username}
	AdminUsername *string `field:"optional" json:"adminUsername" yaml:"adminUsername"`
	// Automatic utility network IP Filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#automatic_utility_network_ip_filter ManagedDatabaseMysql#automatic_utility_network_ip_filter}
	AutomaticUtilityNetworkIpFilter interface{} `field:"optional" json:"automaticUtilityNetworkIpFilter" yaml:"automaticUtilityNetworkIpFilter"`
	// The hour of day (in UTC) when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#backup_hour ManagedDatabaseMysql#backup_hour}
	BackupHour *float64 `field:"optional" json:"backupHour" yaml:"backupHour"`
	// The minute of an hour when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#backup_minute ManagedDatabaseMysql#backup_minute}
	BackupMinute *float64 `field:"optional" json:"backupMinute" yaml:"backupMinute"`
	// The minimum amount of time in seconds to keep binlog entries before deletion.
	//
	// This may be extended for services that require binlog entries for longer than the default for example if using the MySQL Debezium Kafka connector.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#binlog_retention_period ManagedDatabaseMysql#binlog_retention_period}
	BinlogRetentionPeriod *float64 `field:"optional" json:"binlogRetentionPeriod" yaml:"binlogRetentionPeriod"`
	// connect_timeout.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#connect_timeout ManagedDatabaseMysql#connect_timeout}
	ConnectTimeout *float64 `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// default_time_zone.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#default_time_zone ManagedDatabaseMysql#default_time_zone}
	DefaultTimeZone *string `field:"optional" json:"defaultTimeZone" yaml:"defaultTimeZone"`
	// group_concat_max_len.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#group_concat_max_len ManagedDatabaseMysql#group_concat_max_len}
	GroupConcatMaxLen *float64 `field:"optional" json:"groupConcatMaxLen" yaml:"groupConcatMaxLen"`
	// information_schema_stats_expiry.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#information_schema_stats_expiry ManagedDatabaseMysql#information_schema_stats_expiry}
	InformationSchemaStatsExpiry *float64 `field:"optional" json:"informationSchemaStatsExpiry" yaml:"informationSchemaStatsExpiry"`
	// innodb_ft_min_token_size.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#innodb_ft_min_token_size ManagedDatabaseMysql#innodb_ft_min_token_size}
	InnodbFtMinTokenSize *float64 `field:"optional" json:"innodbFtMinTokenSize" yaml:"innodbFtMinTokenSize"`
	// innodb_ft_server_stopword_table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#innodb_ft_server_stopword_table ManagedDatabaseMysql#innodb_ft_server_stopword_table}
	InnodbFtServerStopwordTable *string `field:"optional" json:"innodbFtServerStopwordTable" yaml:"innodbFtServerStopwordTable"`
	// innodb_lock_wait_timeout.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#innodb_lock_wait_timeout ManagedDatabaseMysql#innodb_lock_wait_timeout}
	InnodbLockWaitTimeout *float64 `field:"optional" json:"innodbLockWaitTimeout" yaml:"innodbLockWaitTimeout"`
	// innodb_log_buffer_size.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#innodb_log_buffer_size ManagedDatabaseMysql#innodb_log_buffer_size}
	InnodbLogBufferSize *float64 `field:"optional" json:"innodbLogBufferSize" yaml:"innodbLogBufferSize"`
	// innodb_online_alter_log_max_size.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#innodb_online_alter_log_max_size ManagedDatabaseMysql#innodb_online_alter_log_max_size}
	InnodbOnlineAlterLogMaxSize *float64 `field:"optional" json:"innodbOnlineAlterLogMaxSize" yaml:"innodbOnlineAlterLogMaxSize"`
	// innodb_print_all_deadlocks.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#innodb_print_all_deadlocks ManagedDatabaseMysql#innodb_print_all_deadlocks}
	InnodbPrintAllDeadlocks interface{} `field:"optional" json:"innodbPrintAllDeadlocks" yaml:"innodbPrintAllDeadlocks"`
	// innodb_rollback_on_timeout.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#innodb_rollback_on_timeout ManagedDatabaseMysql#innodb_rollback_on_timeout}
	InnodbRollbackOnTimeout interface{} `field:"optional" json:"innodbRollbackOnTimeout" yaml:"innodbRollbackOnTimeout"`
	// interactive_timeout.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#interactive_timeout ManagedDatabaseMysql#interactive_timeout}
	InteractiveTimeout *float64 `field:"optional" json:"interactiveTimeout" yaml:"interactiveTimeout"`
	// internal_tmp_mem_storage_engine.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#internal_tmp_mem_storage_engine ManagedDatabaseMysql#internal_tmp_mem_storage_engine}
	InternalTmpMemStorageEngine *string `field:"optional" json:"internalTmpMemStorageEngine" yaml:"internalTmpMemStorageEngine"`
	// IP filter.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#ip_filter ManagedDatabaseMysql#ip_filter}
	IpFilter *[]*string `field:"optional" json:"ipFilter" yaml:"ipFilter"`
	// long_query_time.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#long_query_time ManagedDatabaseMysql#long_query_time}
	LongQueryTime *float64 `field:"optional" json:"longQueryTime" yaml:"longQueryTime"`
	// max_allowed_packet.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#max_allowed_packet ManagedDatabaseMysql#max_allowed_packet}
	MaxAllowedPacket *float64 `field:"optional" json:"maxAllowedPacket" yaml:"maxAllowedPacket"`
	// max_heap_table_size.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#max_heap_table_size ManagedDatabaseMysql#max_heap_table_size}
	MaxHeapTableSize *float64 `field:"optional" json:"maxHeapTableSize" yaml:"maxHeapTableSize"`
	// migration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#migration ManagedDatabaseMysql#migration}
	Migration *ManagedDatabaseMysqlPropertiesMigration `field:"optional" json:"migration" yaml:"migration"`
	// net_read_timeout.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#net_read_timeout ManagedDatabaseMysql#net_read_timeout}
	NetReadTimeout *float64 `field:"optional" json:"netReadTimeout" yaml:"netReadTimeout"`
	// net_write_timeout.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#net_write_timeout ManagedDatabaseMysql#net_write_timeout}
	NetWriteTimeout *float64 `field:"optional" json:"netWriteTimeout" yaml:"netWriteTimeout"`
	// Public Access.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#public_access ManagedDatabaseMysql#public_access}
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// slow_query_log.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#slow_query_log ManagedDatabaseMysql#slow_query_log}
	SlowQueryLog interface{} `field:"optional" json:"slowQueryLog" yaml:"slowQueryLog"`
	// sort_buffer_size.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#sort_buffer_size ManagedDatabaseMysql#sort_buffer_size}
	SortBufferSize *float64 `field:"optional" json:"sortBufferSize" yaml:"sortBufferSize"`
	// sql_mode.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#sql_mode ManagedDatabaseMysql#sql_mode}
	SqlMode *string `field:"optional" json:"sqlMode" yaml:"sqlMode"`
	// sql_require_primary_key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#sql_require_primary_key ManagedDatabaseMysql#sql_require_primary_key}
	SqlRequirePrimaryKey interface{} `field:"optional" json:"sqlRequirePrimaryKey" yaml:"sqlRequirePrimaryKey"`
	// tmp_table_size.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#tmp_table_size ManagedDatabaseMysql#tmp_table_size}
	TmpTableSize *float64 `field:"optional" json:"tmpTableSize" yaml:"tmpTableSize"`
	// MySQL major version.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#version ManagedDatabaseMysql#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
	// wait_timeout.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql#wait_timeout ManagedDatabaseMysql#wait_timeout}
	WaitTimeout *float64 `field:"optional" json:"waitTimeout" yaml:"waitTimeout"`
}

