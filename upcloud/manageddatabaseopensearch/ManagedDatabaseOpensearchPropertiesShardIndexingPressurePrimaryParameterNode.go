// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterNode struct {
	// Node soft limit.
	//
	// Define the percentage of the node-level memory
	//                             threshold that acts as a soft indicator for strain on a node.
	//                             Default is 0.7.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#soft_limit ManagedDatabaseOpensearch#soft_limit}
	SoftLimit *float64 `field:"optional" json:"softLimit" yaml:"softLimit"`
}

