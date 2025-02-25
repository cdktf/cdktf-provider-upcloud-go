// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedobjectstorageuseraccesskey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedObjectStorageUserAccessKeyConfig struct {
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
	// Managed Object Storage service UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.0/docs/resources/managed_object_storage_user_access_key#service_uuid ManagedObjectStorageUserAccessKey#service_uuid}
	ServiceUuid *string `field:"required" json:"serviceUuid" yaml:"serviceUuid"`
	// Status of the key. Valid values: `Active`|`Inactive`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.0/docs/resources/managed_object_storage_user_access_key#status ManagedObjectStorageUserAccessKey#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// Username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.0/docs/resources/managed_object_storage_user_access_key#username ManagedObjectStorageUserAccessKey#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.0/docs/resources/managed_object_storage_user_access_key#id ManagedObjectStorageUserAccessKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

