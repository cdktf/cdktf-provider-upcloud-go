// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasevalkey


type ManagedDatabaseValkeyProperties struct {
	// Automatic utility network IP Filter. Automatically allow connections from servers in the utility network within the same zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#automatic_utility_network_ip_filter ManagedDatabaseValkey#automatic_utility_network_ip_filter}
	AutomaticUtilityNetworkIpFilter interface{} `field:"optional" json:"automaticUtilityNetworkIpFilter" yaml:"automaticUtilityNetworkIpFilter"`
	// The hour of day (in UTC) when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#backup_hour ManagedDatabaseValkey#backup_hour}
	BackupHour *float64 `field:"optional" json:"backupHour" yaml:"backupHour"`
	// The minute of an hour when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#backup_minute ManagedDatabaseValkey#backup_minute}
	BackupMinute *float64 `field:"optional" json:"backupMinute" yaml:"backupMinute"`
	// Frequent RDB snapshots.
	//
	// When enabled, Valkey will create frequent local RDB snapshots. When disabled, Valkey will only take RDB snapshots when a backup is created, based on the backup schedule. This setting is ignored when `valkey_persistence` is set to `off`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#frequent_snapshots ManagedDatabaseValkey#frequent_snapshots}
	FrequentSnapshots interface{} `field:"optional" json:"frequentSnapshots" yaml:"frequentSnapshots"`
	// IP filter. Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#ip_filter ManagedDatabaseValkey#ip_filter}
	IpFilter *[]*string `field:"optional" json:"ipFilter" yaml:"ipFilter"`
	// migration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#migration ManagedDatabaseValkey#migration}
	Migration *ManagedDatabaseValkeyPropertiesMigration `field:"optional" json:"migration" yaml:"migration"`
	// Public Access. Allow access to the service from the public Internet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#public_access ManagedDatabaseValkey#public_access}
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// Service logging. Store logs for the service so that they are available in the HTTP API and console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#service_log ManagedDatabaseValkey#service_log}
	ServiceLog interface{} `field:"optional" json:"serviceLog" yaml:"serviceLog"`
	// Default ACL for pub/sub channels used when a Valkey user is created.
	//
	// Determines default pub/sub channels' ACL for new users if ACL is not supplied. When this option is not defined, all_channels is assumed to keep backward compatibility. This option doesn't affect Valkey configuration acl-pubsub-default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_acl_channels_default ManagedDatabaseValkey#valkey_acl_channels_default}
	ValkeyAclChannelsDefault *string `field:"optional" json:"valkeyAclChannelsDefault" yaml:"valkeyAclChannelsDefault"`
	// Active expire effort.
	//
	// Valkey reclaims expired keys both when accessed and in the background. The background process scans for expired keys to free memory. Increasing the active-expire-effort setting (default 1, max 10) uses more CPU to reclaim expired keys faster, reducing memory usage but potentially increasing latency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_active_expire_effort ManagedDatabaseValkey#valkey_active_expire_effort}
	ValkeyActiveExpireEffort *float64 `field:"optional" json:"valkeyActiveExpireEffort" yaml:"valkeyActiveExpireEffort"`
	// Valkey IO thread count. Set Valkey IO thread count. Changing this will cause a restart of the Valkey service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_io_threads ManagedDatabaseValkey#valkey_io_threads}
	ValkeyIoThreads *float64 `field:"optional" json:"valkeyIoThreads" yaml:"valkeyIoThreads"`
	// LFU maxmemory-policy counter decay time in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_lfu_decay_time ManagedDatabaseValkey#valkey_lfu_decay_time}
	ValkeyLfuDecayTime *float64 `field:"optional" json:"valkeyLfuDecayTime" yaml:"valkeyLfuDecayTime"`
	// Counter logarithm factor for volatile-lfu and allkeys-lfu maxmemory-policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_lfu_log_factor ManagedDatabaseValkey#valkey_lfu_log_factor}
	ValkeyLfuLogFactor *float64 `field:"optional" json:"valkeyLfuLogFactor" yaml:"valkeyLfuLogFactor"`
	// Valkey maxmemory-policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_maxmemory_policy ManagedDatabaseValkey#valkey_maxmemory_policy}
	ValkeyMaxmemoryPolicy *string `field:"optional" json:"valkeyMaxmemoryPolicy" yaml:"valkeyMaxmemoryPolicy"`
	// Set notify-keyspace-events option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_notify_keyspace_events ManagedDatabaseValkey#valkey_notify_keyspace_events}
	ValkeyNotifyKeyspaceEvents *string `field:"optional" json:"valkeyNotifyKeyspaceEvents" yaml:"valkeyNotifyKeyspaceEvents"`
	// Number of Valkey databases. Set number of Valkey databases. Changing this will cause a restart of the Valkey service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_number_of_databases ManagedDatabaseValkey#valkey_number_of_databases}
	ValkeyNumberOfDatabases *float64 `field:"optional" json:"valkeyNumberOfDatabases" yaml:"valkeyNumberOfDatabases"`
	// Valkey persistence.
	//
	// When persistence is 'rdb', Valkey does RDB dumps each 10 minutes if any key is changed. Also RDB dumps are done according to backup schedule for backup purposes. When persistence is 'off', no RDB dumps and backups are done, so data can be lost at any moment if service is restarted for any reason, or if service is powered off. Also service can't be forked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_persistence ManagedDatabaseValkey#valkey_persistence}
	ValkeyPersistence *string `field:"optional" json:"valkeyPersistence" yaml:"valkeyPersistence"`
	// Pub/sub client output buffer hard limit in MB.
	//
	// Set output buffer limit for pub / sub clients in MB. The value is the hard limit, the soft limit is 1/4 of the hard limit. When setting the limit, be mindful of the available memory in the selected service plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_pubsub_client_output_buffer_limit ManagedDatabaseValkey#valkey_pubsub_client_output_buffer_limit}
	ValkeyPubsubClientOutputBufferLimit *float64 `field:"optional" json:"valkeyPubsubClientOutputBufferLimit" yaml:"valkeyPubsubClientOutputBufferLimit"`
	// Require SSL to access Valkey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_ssl ManagedDatabaseValkey#valkey_ssl}
	ValkeySsl interface{} `field:"optional" json:"valkeySsl" yaml:"valkeySsl"`
	// Valkey idle connection timeout in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey#valkey_timeout ManagedDatabaseValkey#valkey_timeout}
	ValkeyTimeout *float64 `field:"optional" json:"valkeyTimeout" yaml:"valkeyTimeout"`
}

