// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesShardIndexingPressureOperatingFactor struct {
	// Lower occupancy limit of the allocated quota of memory for the shard.
	//
	// Specify the lower occupancy limit of the allocated quota of memory for the shard.
	//                     If the total memory usage of a shard is below this limit,
	//                     shard indexing backpressure decreases the current allocated memory for that shard.
	//                     Default is 0.75.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.0/docs/resources/managed_database_opensearch#lower ManagedDatabaseOpensearch#lower}
	Lower *float64 `field:"optional" json:"lower" yaml:"lower"`
	// Optimal occupancy of the allocated quota of memory for the shard.
	//
	// Specify the optimal occupancy of the allocated quota of memory for the shard.
	//                     If the total memory usage of a shard is at this level,
	//                     shard indexing backpressure doesn’t change the current allocated memory for that shard.
	//                     Default is 0.85.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.0/docs/resources/managed_database_opensearch#optimal ManagedDatabaseOpensearch#optimal}
	Optimal *float64 `field:"optional" json:"optimal" yaml:"optimal"`
	// Upper occupancy limit of the allocated quota of memory for the shard.
	//
	// Specify the upper occupancy limit of the allocated quota of memory for the shard.
	//                     If the total memory usage of a shard is above this limit,
	//                     shard indexing backpressure increases the current allocated memory for that shard.
	//                     Default is 0.95.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.0/docs/resources/managed_database_opensearch#upper ManagedDatabaseOpensearch#upper}
	Upper *float64 `field:"optional" json:"upper" yaml:"upper"`
}

