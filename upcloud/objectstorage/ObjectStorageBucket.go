// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package objectstorage


type ObjectStorageBucket struct {
	// The name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.1/docs/resources/object_storage#name ObjectStorage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

