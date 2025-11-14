// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filestorage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FileStorageConfig struct {
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
	// The service configured status indicates the service's current intended status. Managed by the customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#configured_status FileStorage#configured_status}
	ConfiguredStatus *string `field:"required" json:"configuredStatus" yaml:"configuredStatus"`
	// Name of the file storage service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#name FileStorage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Size of the file storage in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#size FileStorage#size}
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Zone in which the service will be hosted, e.g. `fi-hel1`. You can list available zones with `upctl zone list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#zone FileStorage#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// User defined key-value pairs to classify the file storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#labels FileStorage#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#network FileStorage#network}
	Network interface{} `field:"optional" json:"network" yaml:"network"`
	// share block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#share FileStorage#share}
	Share interface{} `field:"optional" json:"share" yaml:"share"`
}

