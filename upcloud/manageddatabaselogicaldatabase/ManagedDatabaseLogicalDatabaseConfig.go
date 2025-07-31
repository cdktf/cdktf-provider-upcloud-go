// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaselogicaldatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseLogicalDatabaseConfig struct {
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
	// Name of the logical database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_logical_database#name ManagedDatabaseLogicalDatabase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Service's UUID for which this user belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_logical_database#service ManagedDatabaseLogicalDatabase#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Default character set for the database (LC_CTYPE).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_logical_database#character_set ManagedDatabaseLogicalDatabase#character_set}
	CharacterSet *string `field:"optional" json:"characterSet" yaml:"characterSet"`
	// Default collation for the database (LC_COLLATE).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_logical_database#collation ManagedDatabaseLogicalDatabase#collation}
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_logical_database#id ManagedDatabaseLogicalDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

