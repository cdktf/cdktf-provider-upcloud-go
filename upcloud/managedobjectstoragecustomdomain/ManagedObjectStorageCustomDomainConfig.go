// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedobjectstoragecustomdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedObjectStorageCustomDomainConfig struct {
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
	// Must be a subdomain and consist of 3 to 5 parts such as objects.example.com. Cannot be root-level domain e.g. example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/resources/managed_object_storage_custom_domain#domain_name ManagedObjectStorageCustomDomain#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Managed Object Storage service UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/resources/managed_object_storage_custom_domain#service_uuid ManagedObjectStorageCustomDomain#service_uuid}
	ServiceUuid *string `field:"required" json:"serviceUuid" yaml:"serviceUuid"`
	// At the moment only `public` is accepted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/resources/managed_object_storage_custom_domain#type ManagedObjectStorageCustomDomain#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

