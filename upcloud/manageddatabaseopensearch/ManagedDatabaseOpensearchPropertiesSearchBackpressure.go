// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesSearchBackpressure struct {
	// The search backpressure mode. The search backpressure mode. Valid values are monitor_only, enforced, or disabled. Default is monitor_only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#mode ManagedDatabaseOpensearch#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// node_duress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#node_duress ManagedDatabaseOpensearch#node_duress}
	NodeDuress *ManagedDatabaseOpensearchPropertiesSearchBackpressureNodeDuress `field:"optional" json:"nodeDuress" yaml:"nodeDuress"`
	// search_shard_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#search_shard_task ManagedDatabaseOpensearch#search_shard_task}
	SearchShardTask *ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchShardTask `field:"optional" json:"searchShardTask" yaml:"searchShardTask"`
	// search_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#search_task ManagedDatabaseOpensearch#search_task}
	SearchTask *ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTask `field:"optional" json:"searchTask" yaml:"searchTask"`
}

