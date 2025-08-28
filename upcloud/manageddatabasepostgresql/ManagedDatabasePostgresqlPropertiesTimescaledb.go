// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasepostgresql


type ManagedDatabasePostgresqlPropertiesTimescaledb struct {
	// The number of background workers for timescaledb operations.
	//
	// You should configure this setting to the sum of your number of databases and the total number of concurrent background workers you want running at any given point in time. Changing this parameter causes a service restart.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/managed_database_postgresql#max_background_workers ManagedDatabasePostgresql#max_background_workers}
	MaxBackgroundWorkers *float64 `field:"optional" json:"maxBackgroundWorkers" yaml:"maxBackgroundWorkers"`
}

