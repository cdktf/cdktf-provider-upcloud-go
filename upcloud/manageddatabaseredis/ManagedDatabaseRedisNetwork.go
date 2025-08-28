// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseredis


type ManagedDatabaseRedisNetwork struct {
	// Network family. Currently only `IPv4` is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/managed_database_redis#family ManagedDatabaseRedis#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// The name of the network. Must be unique within the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/managed_database_redis#name ManagedDatabaseRedis#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the network. Must be private.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/managed_database_redis#type ManagedDatabaseRedis#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Private network UUID. Must reside in the same zone as the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/managed_database_redis#uuid ManagedDatabaseRedis#uuid}
	Uuid *string `field:"required" json:"uuid" yaml:"uuid"`
}

