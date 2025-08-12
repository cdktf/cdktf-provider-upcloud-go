// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesSearchInsightsTopQueriesLatency struct {
	// Enable or disable top N query monitoring by the metric.
	//
	// Enable or disable top N query monitoring by the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#enabled ManagedDatabaseOpensearch#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specify the value of N for the top N queries by the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#top_n_size ManagedDatabaseOpensearch#top_n_size}
	TopNSize *float64 `field:"optional" json:"topNSize" yaml:"topNSize"`
	// The window size of the top N queries by the metric.
	//
	// Configure the window size of the top N queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#window_size ManagedDatabaseOpensearch#window_size}
	WindowSize *string `field:"optional" json:"windowSize" yaml:"windowSize"`
}

