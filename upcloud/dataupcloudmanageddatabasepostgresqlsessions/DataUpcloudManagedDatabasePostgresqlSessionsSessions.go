// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataupcloudmanageddatabasepostgresqlsessions


type DataUpcloudManagedDatabasePostgresqlSessionsSessions struct {
	// Top-level transaction identifier of this service, if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/data-sources/managed_database_postgresql_sessions#backend_xid DataUpcloudManagedDatabasePostgresqlSessions#backend_xid}
	BackendXid *float64 `field:"optional" json:"backendXid" yaml:"backendXid"`
	// The current service's xmin horizon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/data-sources/managed_database_postgresql_sessions#backend_xmin DataUpcloudManagedDatabasePostgresqlSessions#backend_xmin}
	BackendXmin *float64 `field:"optional" json:"backendXmin" yaml:"backendXmin"`
	// Host name of the connected client, as reported by a reverse DNS lookup of `client_addr`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/data-sources/managed_database_postgresql_sessions#client_hostname DataUpcloudManagedDatabasePostgresqlSessions#client_hostname}
	ClientHostname *string `field:"optional" json:"clientHostname" yaml:"clientHostname"`
	// Time when this process' current transaction was started, or null if no transaction is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/data-sources/managed_database_postgresql_sessions#xact_start DataUpcloudManagedDatabasePostgresqlSessions#xact_start}
	XactStart *string `field:"optional" json:"xactStart" yaml:"xactStart"`
}

