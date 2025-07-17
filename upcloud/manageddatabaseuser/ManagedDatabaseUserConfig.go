// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseUserConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Service's UUID for which this user belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_user#service ManagedDatabaseUser#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Name of the database user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_user#username ManagedDatabaseUser#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// MySQL only, authentication type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_user#authentication ManagedDatabaseUser#authentication}
	Authentication *string `field:"optional" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_user#id ManagedDatabaseUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// opensearch_access_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_user#opensearch_access_control ManagedDatabaseUser#opensearch_access_control}
	OpensearchAccessControl *ManagedDatabaseUserOpensearchAccessControl `field:"optional" json:"opensearchAccessControl" yaml:"opensearchAccessControl"`
	// Password for the database user. Defaults to a random value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_user#password ManagedDatabaseUser#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// pg_access_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_user#pg_access_control ManagedDatabaseUser#pg_access_control}
	PgAccessControl *ManagedDatabaseUserPgAccessControl `field:"optional" json:"pgAccessControl" yaml:"pgAccessControl"`
	// redis_access_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_user#redis_access_control ManagedDatabaseUser#redis_access_control}
	RedisAccessControl *ManagedDatabaseUserRedisAccessControl `field:"optional" json:"redisAccessControl" yaml:"redisAccessControl"`
	// valkey_access_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_user#valkey_access_control ManagedDatabaseUser#valkey_access_control}
	ValkeyAccessControl *ManagedDatabaseUserValkeyAccessControl `field:"optional" json:"valkeyAccessControl" yaml:"valkeyAccessControl"`
}

