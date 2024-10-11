// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesIndexRollup struct {
	// plugins.rollup.dashboards.enabled. Whether rollups are enabled in OpenSearch Dashboards. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/managed_database_opensearch#rollup_dashboards_enabled ManagedDatabaseOpensearch#rollup_dashboards_enabled}
	RollupDashboardsEnabled interface{} `field:"optional" json:"rollupDashboardsEnabled" yaml:"rollupDashboardsEnabled"`
	// plugins.rollup.enabled. Whether the rollup plugin is enabled. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/managed_database_opensearch#rollup_enabled ManagedDatabaseOpensearch#rollup_enabled}
	RollupEnabled interface{} `field:"optional" json:"rollupEnabled" yaml:"rollupEnabled"`
	// plugins.rollup.search.backoff_count. How many retries the plugin should attempt for failed rollup jobs. Defaults to 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/managed_database_opensearch#rollup_search_backoff_count ManagedDatabaseOpensearch#rollup_search_backoff_count}
	RollupSearchBackoffCount *float64 `field:"optional" json:"rollupSearchBackoffCount" yaml:"rollupSearchBackoffCount"`
	// plugins.rollup.search.backoff_millis. The backoff time between retries for failed rollup jobs. Defaults to 1000ms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/managed_database_opensearch#rollup_search_backoff_millis ManagedDatabaseOpensearch#rollup_search_backoff_millis}
	RollupSearchBackoffMillis *float64 `field:"optional" json:"rollupSearchBackoffMillis" yaml:"rollupSearchBackoffMillis"`
	// plugins.rollup.search.all_jobs. Whether OpenSearch should return all jobs that match all specified search terms. If disabled, OpenSearch returns just one, as opposed to all, of the jobs that matches the search terms. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/managed_database_opensearch#rollup_search_search_all_jobs ManagedDatabaseOpensearch#rollup_search_search_all_jobs}
	RollupSearchSearchAllJobs interface{} `field:"optional" json:"rollupSearchSearchAllJobs" yaml:"rollupSearchSearchAllJobs"`
}

