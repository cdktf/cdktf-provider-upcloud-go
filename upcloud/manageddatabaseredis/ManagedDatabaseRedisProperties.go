// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseredis


type ManagedDatabaseRedisProperties struct {
	// Automatic utility network IP Filter. Automatically allow connections from servers in the utility network within the same zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#automatic_utility_network_ip_filter ManagedDatabaseRedis#automatic_utility_network_ip_filter}
	AutomaticUtilityNetworkIpFilter interface{} `field:"optional" json:"automaticUtilityNetworkIpFilter" yaml:"automaticUtilityNetworkIpFilter"`
	// The hour of day (in UTC) when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#backup_hour ManagedDatabaseRedis#backup_hour}
	BackupHour *float64 `field:"optional" json:"backupHour" yaml:"backupHour"`
	// The minute of an hour when backup for the service is started.
	//
	// New backup is only started if previous backup has already completed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#backup_minute ManagedDatabaseRedis#backup_minute}
	BackupMinute *float64 `field:"optional" json:"backupMinute" yaml:"backupMinute"`
	// IP filter. Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#ip_filter ManagedDatabaseRedis#ip_filter}
	IpFilter *[]*string `field:"optional" json:"ipFilter" yaml:"ipFilter"`
	// migration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#migration ManagedDatabaseRedis#migration}
	Migration *ManagedDatabaseRedisPropertiesMigration `field:"optional" json:"migration" yaml:"migration"`
	// Public Access. Allow access to the service from the public Internet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#public_access ManagedDatabaseRedis#public_access}
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// Default ACL for pub/sub channels used when Redis user is created.
	//
	// Determines default pub/sub channels' ACL for new users if ACL is not supplied. When this option is not defined, all_channels is assumed to keep backward compatibility. This option doesn't affect Redis configuration acl-pubsub-default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_acl_channels_default ManagedDatabaseRedis#redis_acl_channels_default}
	RedisAclChannelsDefault *string `field:"optional" json:"redisAclChannelsDefault" yaml:"redisAclChannelsDefault"`
	// Redis IO thread count. Set Redis IO thread count. Changing this will cause a restart of the Redis service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_io_threads ManagedDatabaseRedis#redis_io_threads}
	RedisIoThreads *float64 `field:"optional" json:"redisIoThreads" yaml:"redisIoThreads"`
	// LFU maxmemory-policy counter decay time in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_lfu_decay_time ManagedDatabaseRedis#redis_lfu_decay_time}
	RedisLfuDecayTime *float64 `field:"optional" json:"redisLfuDecayTime" yaml:"redisLfuDecayTime"`
	// Counter logarithm factor for volatile-lfu and allkeys-lfu maxmemory-policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_lfu_log_factor ManagedDatabaseRedis#redis_lfu_log_factor}
	RedisLfuLogFactor *float64 `field:"optional" json:"redisLfuLogFactor" yaml:"redisLfuLogFactor"`
	// Redis maxmemory-policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_maxmemory_policy ManagedDatabaseRedis#redis_maxmemory_policy}
	RedisMaxmemoryPolicy *string `field:"optional" json:"redisMaxmemoryPolicy" yaml:"redisMaxmemoryPolicy"`
	// Set notify-keyspace-events option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_notify_keyspace_events ManagedDatabaseRedis#redis_notify_keyspace_events}
	RedisNotifyKeyspaceEvents *string `field:"optional" json:"redisNotifyKeyspaceEvents" yaml:"redisNotifyKeyspaceEvents"`
	// Number of Redis databases. Set number of Redis databases. Changing this will cause a restart of the Redis service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_number_of_databases ManagedDatabaseRedis#redis_number_of_databases}
	RedisNumberOfDatabases *float64 `field:"optional" json:"redisNumberOfDatabases" yaml:"redisNumberOfDatabases"`
	// Redis persistence.
	//
	// When persistence is 'rdb', Redis does RDB dumps each 10 minutes if any key is changed. Also RDB dumps are done according to the backup schedule for backup purposes. When persistence is 'off', no RDB dumps or backups are done, so data can be lost at any moment if the service is restarted for any reason, or if the service is powered off. Also, the service can't be forked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_persistence ManagedDatabaseRedis#redis_persistence}
	RedisPersistence *string `field:"optional" json:"redisPersistence" yaml:"redisPersistence"`
	// Pub/sub client output buffer hard limit in MB.
	//
	// Set output buffer limit for pub / sub clients in MB. The value is the hard limit, the soft limit is 1/4 of the hard limit. When setting the limit, be mindful of the available memory in the selected service plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_pubsub_client_output_buffer_limit ManagedDatabaseRedis#redis_pubsub_client_output_buffer_limit}
	RedisPubsubClientOutputBufferLimit *float64 `field:"optional" json:"redisPubsubClientOutputBufferLimit" yaml:"redisPubsubClientOutputBufferLimit"`
	// Require SSL to access Redis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_ssl ManagedDatabaseRedis#redis_ssl}
	RedisSsl interface{} `field:"optional" json:"redisSsl" yaml:"redisSsl"`
	// Redis idle connection timeout in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_timeout ManagedDatabaseRedis#redis_timeout}
	RedisTimeout *float64 `field:"optional" json:"redisTimeout" yaml:"redisTimeout"`
	// Redis major version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#redis_version ManagedDatabaseRedis#redis_version}
	RedisVersion *string `field:"optional" json:"redisVersion" yaml:"redisVersion"`
	// Service logging. Store logs for the service so that they are available in the HTTP API and console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_redis#service_log ManagedDatabaseRedis#service_log}
	ServiceLog interface{} `field:"optional" json:"serviceLog" yaml:"serviceLog"`
}

