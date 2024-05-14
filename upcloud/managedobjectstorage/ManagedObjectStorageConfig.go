// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedobjectstorage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedObjectStorageConfig struct {
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
	// Service status managed by the end user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.3.0/docs/resources/managed_object_storage#configured_status ManagedObjectStorage#configured_status}
	ConfiguredStatus *string `field:"required" json:"configuredStatus" yaml:"configuredStatus"`
	// Name of the Managed Object Storage service. Must be unique within account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.3.0/docs/resources/managed_object_storage#name ManagedObjectStorage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Region in which the service will be hosted, see `upcloud_managed_object_storage_regions` data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.3.0/docs/resources/managed_object_storage#region ManagedObjectStorage#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.3.0/docs/resources/managed_object_storage#id ManagedObjectStorage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key-value pairs to classify the managed object storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.3.0/docs/resources/managed_object_storage#labels ManagedObjectStorage#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.3.0/docs/resources/managed_object_storage#network ManagedObjectStorage#network}
	Network interface{} `field:"optional" json:"network" yaml:"network"`
}

