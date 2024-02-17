// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataupcloudstorage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataUpcloudStorageConfig struct {
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
	// Storage type (normal, backup, cdrom, template). Use 'favorite' as type to filter storages on the list of favorites.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/4.1.0/docs/data-sources/storage#type DataUpcloudStorage#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Storage access type (public, private).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/4.1.0/docs/data-sources/storage#access_type DataUpcloudStorage#access_type}
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/4.1.0/docs/data-sources/storage#id DataUpcloudStorage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If more than one result is returned, use the most recent storage.
	//
	// This is only useful with private storages. Public storages might give unpredictable results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/4.1.0/docs/data-sources/storage#most_recent DataUpcloudStorage#most_recent}
	MostRecent interface{} `field:"optional" json:"mostRecent" yaml:"mostRecent"`
	// Exact name of the storage (same as title).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/4.1.0/docs/data-sources/storage#name DataUpcloudStorage#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Use regular expression to match storage name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/4.1.0/docs/data-sources/storage#name_regex DataUpcloudStorage#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
	// The zone in which the storage resides, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/4.1.0/docs/data-sources/storage#zone DataUpcloudStorage#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

