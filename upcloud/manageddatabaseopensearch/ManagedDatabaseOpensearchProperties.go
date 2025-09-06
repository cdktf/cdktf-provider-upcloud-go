// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchProperties struct {
	// action.auto_create_index. Explicitly allow or block automatic creation of indices. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#action_auto_create_index_enabled ManagedDatabaseOpensearch#action_auto_create_index_enabled}
	ActionAutoCreateIndexEnabled interface{} `field:"optional" json:"actionAutoCreateIndexEnabled" yaml:"actionAutoCreateIndexEnabled"`
	// Require explicit index names when deleting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#action_destructive_requires_name ManagedDatabaseOpensearch#action_destructive_requires_name}
	ActionDestructiveRequiresName interface{} `field:"optional" json:"actionDestructiveRequiresName" yaml:"actionDestructiveRequiresName"`
	// auth_failure_listeners block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#auth_failure_listeners ManagedDatabaseOpensearch#auth_failure_listeners}
	AuthFailureListeners *ManagedDatabaseOpensearchPropertiesAuthFailureListeners `field:"optional" json:"authFailureListeners" yaml:"authFailureListeners"`
	// Automatic utility network IP Filter. Automatically allow connections from servers in the utility network within the same zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#automatic_utility_network_ip_filter ManagedDatabaseOpensearch#automatic_utility_network_ip_filter}
	AutomaticUtilityNetworkIpFilter interface{} `field:"optional" json:"automaticUtilityNetworkIpFilter" yaml:"automaticUtilityNetworkIpFilter"`
	// The limit of how much total remote data can be referenced.
	//
	// Defines a limit of how much total remote data can be referenced as a ratio of the size of the disk reserved for the file cache. This is designed to be a safeguard to prevent oversubscribing a cluster. Defaults to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#cluster_filecache_remote_data_ratio ManagedDatabaseOpensearch#cluster_filecache_remote_data_ratio}
	ClusterFilecacheRemoteDataRatio *float64 `field:"optional" json:"clusterFilecacheRemoteDataRatio" yaml:"clusterFilecacheRemoteDataRatio"`
	// Controls the number of shards allowed in the cluster per data node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#cluster_max_shards_per_node ManagedDatabaseOpensearch#cluster_max_shards_per_node}
	ClusterMaxShardsPerNode *float64 `field:"optional" json:"clusterMaxShardsPerNode" yaml:"clusterMaxShardsPerNode"`
	// cluster_remote_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#cluster_remote_store ManagedDatabaseOpensearch#cluster_remote_store}
	ClusterRemoteStore *ManagedDatabaseOpensearchPropertiesClusterRemoteStore `field:"optional" json:"clusterRemoteStore" yaml:"clusterRemoteStore"`
	// When set to true, OpenSearch attempts to evenly distribute the primary shards between the cluster nodes.
	//
	// Enabling this setting does not always guarantee an equal number of primary shards on each node, especially in the event of a failover. Changing this setting to false after it was set to true does not invoke redistribution of primary shards. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#cluster_routing_allocation_balance_prefer_primary ManagedDatabaseOpensearch#cluster_routing_allocation_balance_prefer_primary}
	ClusterRoutingAllocationBalancePreferPrimary interface{} `field:"optional" json:"clusterRoutingAllocationBalancePreferPrimary" yaml:"clusterRoutingAllocationBalancePreferPrimary"`
	// Concurrent incoming/outgoing shard recoveries per node.
	//
	// How many concurrent incoming/outgoing shard recoveries (normally replicas) are allowed to happen on a node. Defaults to node cpu count * 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#cluster_routing_allocation_node_concurrent_recoveries ManagedDatabaseOpensearch#cluster_routing_allocation_node_concurrent_recoveries}
	ClusterRoutingAllocationNodeConcurrentRecoveries *float64 `field:"optional" json:"clusterRoutingAllocationNodeConcurrentRecoveries" yaml:"clusterRoutingAllocationNodeConcurrentRecoveries"`
	// cluster_search_request_slowlog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#cluster_search_request_slowlog ManagedDatabaseOpensearch#cluster_search_request_slowlog}
	ClusterSearchRequestSlowlog *ManagedDatabaseOpensearchPropertiesClusterSearchRequestSlowlog `field:"optional" json:"clusterSearchRequestSlowlog" yaml:"clusterSearchRequestSlowlog"`
	// Custom domain. Serve the web frontend using a custom CNAME pointing to the Aiven DNS name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#custom_domain ManagedDatabaseOpensearch#custom_domain}
	CustomDomain *string `field:"optional" json:"customDomain" yaml:"customDomain"`
	// disk_watermarks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#disk_watermarks ManagedDatabaseOpensearch#disk_watermarks}
	DiskWatermarks *ManagedDatabaseOpensearchPropertiesDiskWatermarks `field:"optional" json:"diskWatermarks" yaml:"diskWatermarks"`
	// Elasticsearch major version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#elasticsearch_version ManagedDatabaseOpensearch#elasticsearch_version}
	ElasticsearchVersion *string `field:"optional" json:"elasticsearchVersion" yaml:"elasticsearchVersion"`
	// Sender name placeholder to be used in Opensearch Dashboards and Opensearch keystore.
	//
	// This should be identical to the Sender name defined in Opensearch dashboards.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#email_sender_name ManagedDatabaseOpensearch#email_sender_name}
	EmailSenderName *string `field:"optional" json:"emailSenderName" yaml:"emailSenderName"`
	// Sender password for Opensearch alerts to authenticate with SMTP server.
	//
	// Sender password for Opensearch alerts to authenticate with SMTP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#email_sender_password ManagedDatabaseOpensearch#email_sender_password}
	EmailSenderPassword *string `field:"optional" json:"emailSenderPassword" yaml:"emailSenderPassword"`
	// Sender username for Opensearch alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#email_sender_username ManagedDatabaseOpensearch#email_sender_username}
	EmailSenderUsername *string `field:"optional" json:"emailSenderUsername" yaml:"emailSenderUsername"`
	// Enable remote-backed storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#enable_remote_backed_storage ManagedDatabaseOpensearch#enable_remote_backed_storage}
	EnableRemoteBackedStorage interface{} `field:"optional" json:"enableRemoteBackedStorage" yaml:"enableRemoteBackedStorage"`
	// Enable searchable snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#enable_searchable_snapshots ManagedDatabaseOpensearch#enable_searchable_snapshots}
	EnableSearchableSnapshots interface{} `field:"optional" json:"enableSearchableSnapshots" yaml:"enableSearchableSnapshots"`
	// Enable/Disable security audit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#enable_security_audit ManagedDatabaseOpensearch#enable_security_audit}
	EnableSecurityAudit interface{} `field:"optional" json:"enableSecurityAudit" yaml:"enableSecurityAudit"`
	// Enable/Disable snapshot API. Enable/Disable snapshot API for custom repositories, this requires security management to be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#enable_snapshot_api ManagedDatabaseOpensearch#enable_snapshot_api}
	EnableSnapshotApi interface{} `field:"optional" json:"enableSnapshotApi" yaml:"enableSnapshotApi"`
	// Maximum content length for HTTP requests to the OpenSearch HTTP API, in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#http_max_content_length ManagedDatabaseOpensearch#http_max_content_length}
	HttpMaxContentLength *float64 `field:"optional" json:"httpMaxContentLength" yaml:"httpMaxContentLength"`
	// The max size of allowed headers, in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#http_max_header_size ManagedDatabaseOpensearch#http_max_header_size}
	HttpMaxHeaderSize *float64 `field:"optional" json:"httpMaxHeaderSize" yaml:"httpMaxHeaderSize"`
	// The max length of an HTTP URL, in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#http_max_initial_line_length ManagedDatabaseOpensearch#http_max_initial_line_length}
	HttpMaxInitialLineLength *float64 `field:"optional" json:"httpMaxInitialLineLength" yaml:"httpMaxInitialLineLength"`
	// Index patterns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#index_patterns ManagedDatabaseOpensearch#index_patterns}
	IndexPatterns *[]*string `field:"optional" json:"indexPatterns" yaml:"indexPatterns"`
	// index_rollup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#index_rollup ManagedDatabaseOpensearch#index_rollup}
	IndexRollup *ManagedDatabaseOpensearchPropertiesIndexRollup `field:"optional" json:"indexRollup" yaml:"indexRollup"`
	// index_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#index_template ManagedDatabaseOpensearch#index_template}
	IndexTemplate *ManagedDatabaseOpensearchPropertiesIndexTemplate `field:"optional" json:"indexTemplate" yaml:"indexTemplate"`
	// Relative amount.
	//
	// Maximum amount of heap memory used for field data cache. This is an expert setting; decreasing the value too much will increase overhead of loading field data; too much memory used for field data cache will decrease amount of heap available for other operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#indices_fielddata_cache_size ManagedDatabaseOpensearch#indices_fielddata_cache_size}
	IndicesFielddataCacheSize *float64 `field:"optional" json:"indicesFielddataCacheSize" yaml:"indicesFielddataCacheSize"`
	// Percentage value.
	//
	// Default is 10%. Total amount of heap used for indexing buffer, before writing segments to disk. This is an expert setting. Too low value will slow down indexing; too high value will increase indexing performance but causes performance issues for query performance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#indices_memory_index_buffer_size ManagedDatabaseOpensearch#indices_memory_index_buffer_size}
	IndicesMemoryIndexBufferSize *float64 `field:"optional" json:"indicesMemoryIndexBufferSize" yaml:"indicesMemoryIndexBufferSize"`
	// Absolute value.
	//
	// Default is unbound. Doesn't work without indices.memory.index_buffer_size. Maximum amount of heap used for query cache, an absolute indices.memory.index_buffer_size maximum hard limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#indices_memory_max_index_buffer_size ManagedDatabaseOpensearch#indices_memory_max_index_buffer_size}
	IndicesMemoryMaxIndexBufferSize *float64 `field:"optional" json:"indicesMemoryMaxIndexBufferSize" yaml:"indicesMemoryMaxIndexBufferSize"`
	// Absolute value.
	//
	// Default is 48mb. Doesn't work without indices.memory.index_buffer_size. Minimum amount of heap used for query cache, an absolute indices.memory.index_buffer_size minimal hard limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#indices_memory_min_index_buffer_size ManagedDatabaseOpensearch#indices_memory_min_index_buffer_size}
	IndicesMemoryMinIndexBufferSize *float64 `field:"optional" json:"indicesMemoryMinIndexBufferSize" yaml:"indicesMemoryMinIndexBufferSize"`
	// Percentage value.
	//
	// Default is 10%. Maximum amount of heap used for query cache. This is an expert setting. Too low value will decrease query performance and increase performance for other operations; too high value will cause issues with other OpenSearch functionality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#indices_queries_cache_size ManagedDatabaseOpensearch#indices_queries_cache_size}
	IndicesQueriesCacheSize *float64 `field:"optional" json:"indicesQueriesCacheSize" yaml:"indicesQueriesCacheSize"`
	// Maximum number of clauses Lucene BooleanQuery can have.
	//
	// The default value (1024) is relatively high, and increasing it may cause performance issues. Investigate other approaches first before increasing this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#indices_query_bool_max_clause_count ManagedDatabaseOpensearch#indices_query_bool_max_clause_count}
	IndicesQueryBoolMaxClauseCount *float64 `field:"optional" json:"indicesQueryBoolMaxClauseCount" yaml:"indicesQueryBoolMaxClauseCount"`
	// Limits total inbound and outbound recovery traffic for each node.
	//
	// Applies to both peer recoveries as well as snapshot recoveries (i.e., restores from a snapshot). Defaults to 40mb.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#indices_recovery_max_bytes_per_sec ManagedDatabaseOpensearch#indices_recovery_max_bytes_per_sec}
	IndicesRecoveryMaxBytesPerSec *float64 `field:"optional" json:"indicesRecoveryMaxBytesPerSec" yaml:"indicesRecoveryMaxBytesPerSec"`
	// Number of file chunks sent in parallel for each recovery. Defaults to 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#indices_recovery_max_concurrent_file_chunks ManagedDatabaseOpensearch#indices_recovery_max_concurrent_file_chunks}
	IndicesRecoveryMaxConcurrentFileChunks *float64 `field:"optional" json:"indicesRecoveryMaxConcurrentFileChunks" yaml:"indicesRecoveryMaxConcurrentFileChunks"`
	// IP filter. Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#ip_filter ManagedDatabaseOpensearch#ip_filter}
	IpFilter *[]*string `field:"optional" json:"ipFilter" yaml:"ipFilter"`
	// Specifies whether ISM is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#ism_enabled ManagedDatabaseOpensearch#ism_enabled}
	IsmEnabled interface{} `field:"optional" json:"ismEnabled" yaml:"ismEnabled"`
	// Specifies whether audit history is enabled or not. The logs from ISM are automatically indexed to a logs document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#ism_history_enabled ManagedDatabaseOpensearch#ism_history_enabled}
	IsmHistoryEnabled interface{} `field:"optional" json:"ismHistoryEnabled" yaml:"ismHistoryEnabled"`
	// The maximum age before rolling over the audit history index in hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#ism_history_max_age ManagedDatabaseOpensearch#ism_history_max_age}
	IsmHistoryMaxAge *float64 `field:"optional" json:"ismHistoryMaxAge" yaml:"ismHistoryMaxAge"`
	// The maximum number of documents before rolling over the audit history index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#ism_history_max_docs ManagedDatabaseOpensearch#ism_history_max_docs}
	IsmHistoryMaxDocs *float64 `field:"optional" json:"ismHistoryMaxDocs" yaml:"ismHistoryMaxDocs"`
	// The time between rollover checks for the audit history index in hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#ism_history_rollover_check_period ManagedDatabaseOpensearch#ism_history_rollover_check_period}
	IsmHistoryRolloverCheckPeriod *float64 `field:"optional" json:"ismHistoryRolloverCheckPeriod" yaml:"ismHistoryRolloverCheckPeriod"`
	// How long audit history indices are kept in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#ism_history_rollover_retention_period ManagedDatabaseOpensearch#ism_history_rollover_retention_period}
	IsmHistoryRolloverRetentionPeriod *float64 `field:"optional" json:"ismHistoryRolloverRetentionPeriod" yaml:"ismHistoryRolloverRetentionPeriod"`
	// Don't reset index.refresh_interval to the default value. Aiven automation resets index.refresh_interval to default value for every index to be sure that indices are always visible to search. If it doesn't fit your case, you can disable this by setting up this flag to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#keep_index_refresh_interval ManagedDatabaseOpensearch#keep_index_refresh_interval}
	KeepIndexRefreshInterval interface{} `field:"optional" json:"keepIndexRefreshInterval" yaml:"keepIndexRefreshInterval"`
	// Enable or disable KNN memory circuit breaker. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#knn_memory_circuit_breaker_enabled ManagedDatabaseOpensearch#knn_memory_circuit_breaker_enabled}
	KnnMemoryCircuitBreakerEnabled interface{} `field:"optional" json:"knnMemoryCircuitBreakerEnabled" yaml:"knnMemoryCircuitBreakerEnabled"`
	// Maximum amount of memory that can be used for KNN index. Defaults to 50% of the JVM heap size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#knn_memory_circuit_breaker_limit ManagedDatabaseOpensearch#knn_memory_circuit_breaker_limit}
	KnnMemoryCircuitBreakerLimit *float64 `field:"optional" json:"knnMemoryCircuitBreakerLimit" yaml:"knnMemoryCircuitBreakerLimit"`
	// The limit of how much total remote data can be referenced.
	//
	// Defines a limit of how much total remote data can be referenced as a ratio of the size of the disk reserved for the file cache. This is designed to be a safeguard to prevent oversubscribing a cluster. Defaults to 5gb. Requires restarting all OpenSearch nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#node_search_cache_size ManagedDatabaseOpensearch#node_search_cache_size}
	NodeSearchCacheSize *string `field:"optional" json:"nodeSearchCacheSize" yaml:"nodeSearchCacheSize"`
	// openid block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#openid ManagedDatabaseOpensearch#openid}
	Openid *ManagedDatabaseOpensearchPropertiesOpenid `field:"optional" json:"openid" yaml:"openid"`
	// opensearch_dashboards block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#opensearch_dashboards ManagedDatabaseOpensearch#opensearch_dashboards}
	OpensearchDashboards *ManagedDatabaseOpensearchPropertiesOpensearchDashboards `field:"optional" json:"opensearchDashboards" yaml:"opensearchDashboards"`
	// Compatibility mode sets OpenSearch to report its version as 7.10 so clients continue to work. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#override_main_response_version ManagedDatabaseOpensearch#override_main_response_version}
	OverrideMainResponseVersion interface{} `field:"optional" json:"overrideMainResponseVersion" yaml:"overrideMainResponseVersion"`
	// Enable or disable filtering of alerting by backend roles. Requires Security plugin. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#plugins_alerting_filter_by_backend_roles ManagedDatabaseOpensearch#plugins_alerting_filter_by_backend_roles}
	PluginsAlertingFilterByBackendRoles interface{} `field:"optional" json:"pluginsAlertingFilterByBackendRoles" yaml:"pluginsAlertingFilterByBackendRoles"`
	// Public Access. Allow access to the service from the public Internet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#public_access ManagedDatabaseOpensearch#public_access}
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// Whitelisted addresses for reindexing. Changing this value will cause all OpenSearch instances to restart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#reindex_remote_whitelist ManagedDatabaseOpensearch#reindex_remote_whitelist}
	ReindexRemoteWhitelist *[]*string `field:"optional" json:"reindexRemoteWhitelist" yaml:"reindexRemoteWhitelist"`
	// remote_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#remote_store ManagedDatabaseOpensearch#remote_store}
	RemoteStore *ManagedDatabaseOpensearchPropertiesRemoteStore `field:"optional" json:"remoteStore" yaml:"remoteStore"`
	// saml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#saml ManagedDatabaseOpensearch#saml}
	Saml *ManagedDatabaseOpensearchPropertiesSaml `field:"optional" json:"saml" yaml:"saml"`
	// Script max compilation rate - circuit breaker to prevent/minimize OOMs.
	//
	// Script compilation circuit breaker limits the number of inline script compilations within a period of time. Default is use-context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#script_max_compilations_rate ManagedDatabaseOpensearch#script_max_compilations_rate}
	ScriptMaxCompilationsRate *string `field:"optional" json:"scriptMaxCompilationsRate" yaml:"scriptMaxCompilationsRate"`
	// search_backpressure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#search_backpressure ManagedDatabaseOpensearch#search_backpressure}
	SearchBackpressure *ManagedDatabaseOpensearchPropertiesSearchBackpressure `field:"optional" json:"searchBackpressure" yaml:"searchBackpressure"`
	// search_insights_top_queries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#search_insights_top_queries ManagedDatabaseOpensearch#search_insights_top_queries}
	SearchInsightsTopQueries *ManagedDatabaseOpensearchPropertiesSearchInsightsTopQueries `field:"optional" json:"searchInsightsTopQueries" yaml:"searchInsightsTopQueries"`
	// Maximum number of aggregation buckets allowed in a single response.
	//
	// OpenSearch default value is used when this is not defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#search_max_buckets ManagedDatabaseOpensearch#search_max_buckets}
	SearchMaxBuckets *float64 `field:"optional" json:"searchMaxBuckets" yaml:"searchMaxBuckets"`
	// segrep block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#segrep ManagedDatabaseOpensearch#segrep}
	Segrep *ManagedDatabaseOpensearchPropertiesSegrep `field:"optional" json:"segrep" yaml:"segrep"`
	// Service logging. Store logs for the service so that they are available in the HTTP API and console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#service_log ManagedDatabaseOpensearch#service_log}
	ServiceLog interface{} `field:"optional" json:"serviceLog" yaml:"serviceLog"`
	// shard_indexing_pressure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#shard_indexing_pressure ManagedDatabaseOpensearch#shard_indexing_pressure}
	ShardIndexingPressure *ManagedDatabaseOpensearchPropertiesShardIndexingPressure `field:"optional" json:"shardIndexingPressure" yaml:"shardIndexingPressure"`
	// analyze thread pool queue size. Size for the thread pool queue. See documentation for exact details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#thread_pool_analyze_queue_size ManagedDatabaseOpensearch#thread_pool_analyze_queue_size}
	ThreadPoolAnalyzeQueueSize *float64 `field:"optional" json:"threadPoolAnalyzeQueueSize" yaml:"threadPoolAnalyzeQueueSize"`
	// analyze thread pool size.
	//
	// Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#thread_pool_analyze_size ManagedDatabaseOpensearch#thread_pool_analyze_size}
	ThreadPoolAnalyzeSize *float64 `field:"optional" json:"threadPoolAnalyzeSize" yaml:"threadPoolAnalyzeSize"`
	// force_merge thread pool size.
	//
	// Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#thread_pool_force_merge_size ManagedDatabaseOpensearch#thread_pool_force_merge_size}
	ThreadPoolForceMergeSize *float64 `field:"optional" json:"threadPoolForceMergeSize" yaml:"threadPoolForceMergeSize"`
	// get thread pool queue size. Size for the thread pool queue. See documentation for exact details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#thread_pool_get_queue_size ManagedDatabaseOpensearch#thread_pool_get_queue_size}
	ThreadPoolGetQueueSize *float64 `field:"optional" json:"threadPoolGetQueueSize" yaml:"threadPoolGetQueueSize"`
	// get thread pool size.
	//
	// Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#thread_pool_get_size ManagedDatabaseOpensearch#thread_pool_get_size}
	ThreadPoolGetSize *float64 `field:"optional" json:"threadPoolGetSize" yaml:"threadPoolGetSize"`
	// search thread pool queue size. Size for the thread pool queue. See documentation for exact details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#thread_pool_search_queue_size ManagedDatabaseOpensearch#thread_pool_search_queue_size}
	ThreadPoolSearchQueueSize *float64 `field:"optional" json:"threadPoolSearchQueueSize" yaml:"threadPoolSearchQueueSize"`
	// search thread pool size.
	//
	// Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#thread_pool_search_size ManagedDatabaseOpensearch#thread_pool_search_size}
	ThreadPoolSearchSize *float64 `field:"optional" json:"threadPoolSearchSize" yaml:"threadPoolSearchSize"`
	// search_throttled thread pool queue size. Size for the thread pool queue. See documentation for exact details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#thread_pool_search_throttled_queue_size ManagedDatabaseOpensearch#thread_pool_search_throttled_queue_size}
	ThreadPoolSearchThrottledQueueSize *float64 `field:"optional" json:"threadPoolSearchThrottledQueueSize" yaml:"threadPoolSearchThrottledQueueSize"`
	// search_throttled thread pool size.
	//
	// Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#thread_pool_search_throttled_size ManagedDatabaseOpensearch#thread_pool_search_throttled_size}
	ThreadPoolSearchThrottledSize *float64 `field:"optional" json:"threadPoolSearchThrottledSize" yaml:"threadPoolSearchThrottledSize"`
	// write thread pool queue size. Size for the thread pool queue. See documentation for exact details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#thread_pool_write_queue_size ManagedDatabaseOpensearch#thread_pool_write_queue_size}
	ThreadPoolWriteQueueSize *float64 `field:"optional" json:"threadPoolWriteQueueSize" yaml:"threadPoolWriteQueueSize"`
	// write thread pool size.
	//
	// Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#thread_pool_write_size ManagedDatabaseOpensearch#thread_pool_write_size}
	ThreadPoolWriteSize *float64 `field:"optional" json:"threadPoolWriteSize" yaml:"threadPoolWriteSize"`
	// OpenSearch major version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#version ManagedDatabaseOpensearch#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

