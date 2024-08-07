// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storage


type StorageImport struct {
	// The mode of the import task. One of `http_import` or `direct_upload`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/storage#source Storage#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The location of the file to import. For `http_import` an accessible URL for `direct_upload` a local file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/storage#source_location Storage#source_location}
	SourceLocation *string `field:"required" json:"sourceLocation" yaml:"sourceLocation"`
	// For `direct_upload`; an optional hash of the file to upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/storage#source_hash Storage#source_hash}
	SourceHash *string `field:"optional" json:"sourceHash" yaml:"sourceHash"`
}

