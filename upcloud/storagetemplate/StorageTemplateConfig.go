// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageTemplateConfig struct {
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
	// The source storage that is used as a base for this storage template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.30.0/docs/resources/storage_template#source_storage StorageTemplate#source_storage}
	SourceStorage *string `field:"required" json:"sourceStorage" yaml:"sourceStorage"`
	// The title of the storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.30.0/docs/resources/storage_template#title StorageTemplate#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// User defined key-value pairs to classify the storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.30.0/docs/resources/storage_template#labels StorageTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

