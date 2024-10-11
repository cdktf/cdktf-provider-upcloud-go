// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasepostgresql


type ManagedDatabasePostgresqlNetwork struct {
	// Network family. Currently only `IPv4` is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/managed_database_postgresql#family ManagedDatabasePostgresql#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// The name of the network. Must be unique within the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/managed_database_postgresql#name ManagedDatabasePostgresql#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the network. Must be private.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/managed_database_postgresql#type ManagedDatabasePostgresql#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Private network UUID. Must reside in the same zone as the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/managed_database_postgresql#uuid ManagedDatabasePostgresql#uuid}
	Uuid *string `field:"required" json:"uuid" yaml:"uuid"`
}

