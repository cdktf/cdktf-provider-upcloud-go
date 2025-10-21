// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameter struct {
	// node block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#node ManagedDatabaseOpensearch#node}
	NodeAttribute *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterNode `field:"optional" json:"nodeAttribute" yaml:"nodeAttribute"`
	// shard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#shard ManagedDatabaseOpensearch#shard}
	Shard *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameterShard `field:"optional" json:"shard" yaml:"shard"`
}

