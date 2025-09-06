// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesSegrep struct {
	// The maximum number of indexing checkpoints that a replica shard can fall behind when copying from primary.
	//
	// Once `segrep.pressure.checkpoint.limit` is breached along with `segrep.pressure.time.limit`, the segment replication backpressure mechanism is initiated. Default is 4 checkpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#pressure_checkpoint_limit ManagedDatabaseOpensearch#pressure_checkpoint_limit}
	PressureCheckpointLimit *float64 `field:"optional" json:"pressureCheckpointLimit" yaml:"pressureCheckpointLimit"`
	// Enables the segment replication backpressure mechanism. Default is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#pressure_enabled ManagedDatabaseOpensearch#pressure_enabled}
	PressureEnabled interface{} `field:"optional" json:"pressureEnabled" yaml:"pressureEnabled"`
	// The maximum number of stale replica shards that can exist in a replication group.
	//
	// Once `segrep.pressure.replica.stale.limit` is breached, the segment replication backpressure mechanism is initiated. Default is .5, which is 50% of a replication group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#pressure_replica_stale_limit ManagedDatabaseOpensearch#pressure_replica_stale_limit}
	PressureReplicaStaleLimit *float64 `field:"optional" json:"pressureReplicaStaleLimit" yaml:"pressureReplicaStaleLimit"`
	// The maximum amount of time that a replica shard can take to copy from the primary shard.
	//
	// Once segrep.pressure.time.limit is breached along with segrep.pressure.checkpoint.limit, the segment replication backpressure mechanism is initiated. Default is 5 minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_opensearch#pressure_time_limit ManagedDatabaseOpensearch#pressure_time_limit}
	PressureTimeLimit *string `field:"optional" json:"pressureTimeLimit" yaml:"pressureTimeLimit"`
}

