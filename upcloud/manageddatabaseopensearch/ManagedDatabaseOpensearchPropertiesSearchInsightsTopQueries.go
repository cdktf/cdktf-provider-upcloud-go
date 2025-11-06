// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesSearchInsightsTopQueries struct {
	// cpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#cpu ManagedDatabaseOpensearch#cpu}
	Cpu *ManagedDatabaseOpensearchPropertiesSearchInsightsTopQueriesCpu `field:"optional" json:"cpu" yaml:"cpu"`
	// latency block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#latency ManagedDatabaseOpensearch#latency}
	Latency *ManagedDatabaseOpensearchPropertiesSearchInsightsTopQueriesLatency `field:"optional" json:"latency" yaml:"latency"`
	// memory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#memory ManagedDatabaseOpensearch#memory}
	Memory *ManagedDatabaseOpensearchPropertiesSearchInsightsTopQueriesMemory `field:"optional" json:"memory" yaml:"memory"`
}

