// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseuser


type ManagedDatabaseUserOpensearchAccessControlRules struct {
	// Set index name, pattern or top level API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_user#index ManagedDatabaseUser#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// Set permission access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_user#permission ManagedDatabaseUser#permission}
	Permission *string `field:"required" json:"permission" yaml:"permission"`
}

