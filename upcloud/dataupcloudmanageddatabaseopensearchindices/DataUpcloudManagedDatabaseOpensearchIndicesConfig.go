// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataupcloudmanageddatabaseopensearchindices

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataUpcloudManagedDatabaseOpensearchIndicesConfig struct {
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
	// Service's UUID for which these indices belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.4.0/docs/data-sources/managed_database_opensearch_indices#service DataUpcloudManagedDatabaseOpensearchIndices#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.4.0/docs/data-sources/managed_database_opensearch_indices#id DataUpcloudManagedDatabaseOpensearchIndices#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// indices block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.4.0/docs/data-sources/managed_database_opensearch_indices#indices DataUpcloudManagedDatabaseOpensearchIndices#indices}
	Indices interface{} `field:"optional" json:"indices" yaml:"indices"`
}

