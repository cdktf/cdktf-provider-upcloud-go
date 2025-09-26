// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedobjectstoragepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedObjectStoragePolicyConfig struct {
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
	// Policy document, URL-encoded compliant with RFC 3986.
	//
	// Extra whitespace and escapes are ignored when determining if the document has changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/managed_object_storage_policy#document ManagedObjectStoragePolicy#document}
	Document *string `field:"required" json:"document" yaml:"document"`
	// Policy name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/managed_object_storage_policy#name ManagedObjectStoragePolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Managed Object Storage service UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/managed_object_storage_policy#service_uuid ManagedObjectStoragePolicy#service_uuid}
	ServiceUuid *string `field:"required" json:"serviceUuid" yaml:"serviceUuid"`
	// Description of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/managed_object_storage_policy#description ManagedObjectStoragePolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

