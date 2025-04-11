// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseuser


type ManagedDatabaseUserValkeyAccessControl struct {
	// Set access control to all commands in specified categories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_user#categories ManagedDatabaseUser#categories}
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// Set access control to Pub/Sub channels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_user#channels ManagedDatabaseUser#channels}
	Channels *[]*string `field:"optional" json:"channels" yaml:"channels"`
	// Set access control to commands.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_user#commands ManagedDatabaseUser#commands}
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// Set access control to keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_user#keys ManagedDatabaseUser#keys}
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
}

