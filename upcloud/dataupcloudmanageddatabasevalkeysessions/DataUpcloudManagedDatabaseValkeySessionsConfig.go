// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataupcloudmanageddatabasevalkeysessions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataUpcloudManagedDatabaseValkeySessionsConfig struct {
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
	// Service's UUID for which these sessions belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.18.0/docs/data-sources/managed_database_valkey_sessions#service DataUpcloudManagedDatabaseValkeySessions#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.18.0/docs/data-sources/managed_database_valkey_sessions#id DataUpcloudManagedDatabaseValkeySessions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Number of entries to receive at most.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.18.0/docs/data-sources/managed_database_valkey_sessions#limit DataUpcloudManagedDatabaseValkeySessions#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Offset for retrieved results based on sort order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.18.0/docs/data-sources/managed_database_valkey_sessions#offset DataUpcloudManagedDatabaseValkeySessions#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Order by session field and sort retrieved results. Limited variables can be used for ordering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.18.0/docs/data-sources/managed_database_valkey_sessions#order DataUpcloudManagedDatabaseValkeySessions#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// sessions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.18.0/docs/data-sources/managed_database_valkey_sessions#sessions DataUpcloudManagedDatabaseValkeySessions#sessions}
	Sessions interface{} `field:"optional" json:"sessions" yaml:"sessions"`
}

