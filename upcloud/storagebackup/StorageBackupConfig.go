// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebackup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageBackupConfig struct {
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
	// The UUID of the storage to back up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/storage_backup#source_storage StorageBackup#source_storage}
	SourceStorage *string `field:"required" json:"sourceStorage" yaml:"sourceStorage"`
	// Title of the backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/storage_backup#title StorageBackup#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// User defined key-value pairs to classify the storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/storage_backup#labels StorageBackup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

