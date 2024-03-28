// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package objectstorage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObjectStorageConfig struct {
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
	// The access key used to identify user.
	//
	// Can be set to an empty string, which will tell the provider to get the access key from environment variable.
	// 				The environment variable should be "UPCLOUD_OBJECT_STORAGE_ACCESS_KEY_{name}".
	// 				{name} is the name given to object storage instance (so not the resource label), it should be all uppercased
	// 				and all dashes (-) should be replaced with underscores (_). For example, object storage named "my-files" would
	// 				use environment variable named "UPCLOUD_OBJECT_STORAGE_ACCESS_KEY_MY_FILES".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/object_storage#access_key ObjectStorage#access_key}
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// The name of the object storage instance to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/object_storage#name ObjectStorage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The secret key used to authenticate user.
	//
	// Can be set to an empty string, which will tell the provider to get the secret key from environment variable.
	// 				The environment variable should be "UPCLOUD_OBJECT_STORAGE_SECRET_KEY_{name}".
	// 				{name} is the name given to object storage instance (so not the resource label), it should be all uppercased
	// 				and all dashes (-) should be replaced with underscores (_). For example, object storage named "my-files" would
	// 				use environment variable named "UPCLOUD_OBJECT_STORAGE_SECRET_KEY_MY_FILES".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/object_storage#secret_key ObjectStorage#secret_key}
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
	// The size of the object storage instance in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/object_storage#size ObjectStorage#size}
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// The zone in which the object storage instance will be created, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/object_storage#zone ObjectStorage#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/object_storage#bucket ObjectStorage#bucket}
	Bucket interface{} `field:"optional" json:"bucket" yaml:"bucket"`
	// The description of the object storage instance to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/object_storage#description ObjectStorage#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/object_storage#id ObjectStorage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

