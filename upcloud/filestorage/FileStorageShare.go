// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filestorage


type FileStorageShare struct {
	// Unique name of the share (1–64 chars).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#name FileStorage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Absolute path exported by the share (e.g. `/public`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#path FileStorage#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// acl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#acl FileStorage#acl}
	Acl interface{} `field:"optional" json:"acl" yaml:"acl"`
}

