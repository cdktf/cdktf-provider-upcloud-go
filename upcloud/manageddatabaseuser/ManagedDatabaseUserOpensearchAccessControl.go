// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseuser


type ManagedDatabaseUserOpensearchAccessControl struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/resources/managed_database_user#rules ManagedDatabaseUser#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

