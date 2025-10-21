// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesShardIndexingPressure struct {
	// Enable or disable shard indexing backpressure. Enable or disable shard indexing backpressure. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#enabled ManagedDatabaseOpensearch#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Run shard indexing backpressure in shadow mode or enforced mode.
	//
	// Run shard indexing backpressure in shadow mode or enforced mode.
	//             In shadow mode (value set as false), shard indexing backpressure tracks all granular-level metrics,
	//             but it doesn’t actually reject any indexing requests.
	//             In enforced mode (value set as true),
	//             shard indexing backpressure rejects any requests to the cluster that might cause a dip in its performance.
	//             Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#enforced ManagedDatabaseOpensearch#enforced}
	Enforced interface{} `field:"optional" json:"enforced" yaml:"enforced"`
	// operating_factor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#operating_factor ManagedDatabaseOpensearch#operating_factor}
	OperatingFactor *ManagedDatabaseOpensearchPropertiesShardIndexingPressureOperatingFactor `field:"optional" json:"operatingFactor" yaml:"operatingFactor"`
	// primary_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#primary_parameter ManagedDatabaseOpensearch#primary_parameter}
	PrimaryParameter *ManagedDatabaseOpensearchPropertiesShardIndexingPressurePrimaryParameter `field:"optional" json:"primaryParameter" yaml:"primaryParameter"`
}

