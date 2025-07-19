// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterShard struct {
	// Shard min limit.
	//
	// Specify the minimum assigned quota for a new shard in any role (coordinator, primary, or replica).
	//                             Shard indexing backpressure increases or decreases this allocated quota based on the inflow of traffic for the shard.
	//                             Default is 0.001.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_opensearch#min_limit ManagedDatabaseOpensearch#min_limit}
	MinLimit *float64 `field:"optional" json:"minLimit" yaml:"minLimit"`
}

