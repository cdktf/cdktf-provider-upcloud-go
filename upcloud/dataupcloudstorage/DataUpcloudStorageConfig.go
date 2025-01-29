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
	// The access type of the storage, `public` or `private`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/data-sources/storage#access_type DataUpcloudStorage#access_type}
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// UUID of the storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/data-sources/storage#id DataUpcloudStorage#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If more than one result is returned, use the most recent storage.
	//
	// This is only useful with private storages. Public storages might give unpredictable results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/data-sources/storage#most_recent DataUpcloudStorage#most_recent}
	MostRecent interface{} `field:"optional" json:"mostRecent" yaml:"mostRecent"`
	// Exact name of the storage (same as title). Deprecated, use `title` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/data-sources/storage#name DataUpcloudStorage#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Use regular expression to match storage name. Deprecated, use exact title or UUID instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/data-sources/storage#name_regex DataUpcloudStorage#name_regex}
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
	// The title of the storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/data-sources/storage#title DataUpcloudStorage#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The type of the storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/data-sources/storage#type DataUpcloudStorage#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The zone the storage is in, e.g. `de-fra1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.17.0/docs/data-sources/storage#zone DataUpcloudStorage#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

