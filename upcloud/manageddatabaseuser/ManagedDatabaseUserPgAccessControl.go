// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseuser


type ManagedDatabaseUserPgAccessControl struct {
	// Grant replication privilege.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/managed_database_user#allow_replication ManagedDatabaseUser#allow_replication}
	AllowReplication interface{} `field:"optional" json:"allowReplication" yaml:"allowReplication"`
}

