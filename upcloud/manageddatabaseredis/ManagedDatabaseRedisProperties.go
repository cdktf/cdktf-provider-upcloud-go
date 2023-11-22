// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseredis


type ManagedDatabaseRedisProperties struct {
	// Automatic utility network IP Filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#automatic_utility_network_ip_filter ManagedDatabaseRedis#automatic_utility_network_ip_filter}
	AutomaticUtilityNetworkIpFilter interface{} `field:"optional" json:"automaticUtilityNetworkIpFilter" yaml:"automaticUtilityNetworkIpFilter"`
	// IP filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#ip_filter ManagedDatabaseRedis#ip_filter}
	IpFilter *[]*string `field:"optional" json:"ipFilter" yaml:"ipFilter"`
	// migration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#migration ManagedDatabaseRedis#migration}
	Migration *ManagedDatabaseRedisPropertiesMigration `field:"optional" json:"migration" yaml:"migration"`
	// Public access allows connections to your Managed Database services via the public internet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#public_access ManagedDatabaseRedis#public_access}
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// Default ACL for pub/sub channels used when Redis user is created.
	//
	// Determines default pub/sub channels' ACL for new users if ACL is not supplied. When this option is not defined, all_channels is assumed to keep backward compatibility. This option doesn't affect Redis configuration acl-pubsub-default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#redis_acl_channels_default ManagedDatabaseRedis#redis_acl_channels_default}
	RedisAclChannelsDefault *string `field:"optional" json:"redisAclChannelsDefault" yaml:"redisAclChannelsDefault"`
	// Redis IO thread count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#redis_io_threads ManagedDatabaseRedis#redis_io_threads}
	RedisIoThreads *float64 `field:"optional" json:"redisIoThreads" yaml:"redisIoThreads"`
	// LFU maxmemory-policy counter decay time in minutes. Default is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#redis_lfu_decay_time ManagedDatabaseRedis#redis_lfu_decay_time}
	RedisLfuDecayTime *float64 `field:"optional" json:"redisLfuDecayTime" yaml:"redisLfuDecayTime"`
	// Counter logarithm factor for volatile-lfu and allkeys-lfu maxmemory-policies. Default is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#redis_lfu_log_factor ManagedDatabaseRedis#redis_lfu_log_factor}
	RedisLfuLogFactor *float64 `field:"optional" json:"redisLfuLogFactor" yaml:"redisLfuLogFactor"`
	// Redis maxmemory-policy. Default is `noeviction`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#redis_maxmemory_policy ManagedDatabaseRedis#redis_maxmemory_policy}
	RedisMaxmemoryPolicy *string `field:"optional" json:"redisMaxmemoryPolicy" yaml:"redisMaxmemoryPolicy"`
	// Set notify-keyspace-events option. Default is "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#redis_notify_keyspace_events ManagedDatabaseRedis#redis_notify_keyspace_events}
	RedisNotifyKeyspaceEvents *string `field:"optional" json:"redisNotifyKeyspaceEvents" yaml:"redisNotifyKeyspaceEvents"`
	// Number of redis databases. Set number of redis databases. Changing this will cause a restart of redis service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#redis_number_of_databases ManagedDatabaseRedis#redis_number_of_databases}
	RedisNumberOfDatabases *float64 `field:"optional" json:"redisNumberOfDatabases" yaml:"redisNumberOfDatabases"`
	// Redis persistence.
	//
	// When persistence is 'rdb', Redis does RDB dumps each 10 minutes if any key is changed. Also RDB dumps are done according to backup schedule for backup purposes. When persistence is 'off', no RDB dumps and backups are done, so data can be lost at any moment if service is restarted for any reason, or if service is powered off. Also service can't be forked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#redis_persistence ManagedDatabaseRedis#redis_persistence}
	RedisPersistence *string `field:"optional" json:"redisPersistence" yaml:"redisPersistence"`
	// Pub/sub client output buffer hard limit in MB.
	//
	// Set output buffer limit for pub / sub clients in MB. The value is the hard limit, the soft limit is 1/4 of the hard limit. When setting the limit, be mindful of the available memory in the selected service plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#redis_pubsub_client_output_buffer_limit ManagedDatabaseRedis#redis_pubsub_client_output_buffer_limit}
	RedisPubsubClientOutputBufferLimit *float64 `field:"optional" json:"redisPubsubClientOutputBufferLimit" yaml:"redisPubsubClientOutputBufferLimit"`
	// Require SSL to access Redis. Default is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#redis_ssl ManagedDatabaseRedis#redis_ssl}
	RedisSsl interface{} `field:"optional" json:"redisSsl" yaml:"redisSsl"`
	// Redis idle connection timeout in seconds. Default is 300.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/managed_database_redis#redis_timeout ManagedDatabaseRedis#redis_timeout}
	RedisTimeout *float64 `field:"optional" json:"redisTimeout" yaml:"redisTimeout"`
}

