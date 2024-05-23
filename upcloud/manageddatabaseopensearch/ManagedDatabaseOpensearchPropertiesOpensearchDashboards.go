// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesOpensearchDashboards struct {
	// Enable or disable OpenSearch Dashboards.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_opensearch#enabled ManagedDatabaseOpensearch#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Limits the maximum amount of memory (in MiB) the OpenSearch Dashboards process can use.
	//
	// This sets the max_old_space_size option of the nodejs running the OpenSearch Dashboards. Note: the memory reserved by OpenSearch Dashboards is not available for OpenSearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_opensearch#max_old_space_size ManagedDatabaseOpensearch#max_old_space_size}
	MaxOldSpaceSize *float64 `field:"optional" json:"maxOldSpaceSize" yaml:"maxOldSpaceSize"`
	// Timeout in milliseconds for requests made by OpenSearch Dashboards towards OpenSearch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_opensearch#opensearch_request_timeout ManagedDatabaseOpensearch#opensearch_request_timeout}
	OpensearchRequestTimeout *float64 `field:"optional" json:"opensearchRequestTimeout" yaml:"opensearchRequestTimeout"`
}

