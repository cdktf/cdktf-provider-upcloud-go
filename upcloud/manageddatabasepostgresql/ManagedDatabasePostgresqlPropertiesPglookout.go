// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasepostgresql


type ManagedDatabasePostgresqlPropertiesPglookout struct {
	// Number of seconds of master unavailability before triggering database failover to standby.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.11.1/docs/resources/managed_database_postgresql#max_failover_replication_time_lag ManagedDatabasePostgresql#max_failover_replication_time_lag}
	MaxFailoverReplicationTimeLag *float64 `field:"optional" json:"maxFailoverReplicationTimeLag" yaml:"maxFailoverReplicationTimeLag"`
}

