// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storage


type StorageImport struct {
	// The mode of the import task. One of `http_import` or `direct_upload`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.32.0/docs/resources/storage#source Storage#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The location of the file to import.
	//
	// For `http_import` an accessible URL. For `direct_upload` a local file. When direct uploading a compressed image, `Content-Type` header of the PUT request is set automatically based on the file extension (`.gz` or `.xz`, case-insensitive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.32.0/docs/resources/storage#source_location Storage#source_location}
	SourceLocation *string `field:"required" json:"sourceLocation" yaml:"sourceLocation"`
	// SHA256 hash of the source content.
	//
	// This hash is used to verify the integrity of the imported data by comparing it to `sha256sum` after the import has completed. Possible filename is automatically removed from the hash before comparison.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.32.0/docs/resources/storage#source_hash Storage#source_hash}
	SourceHash *string `field:"optional" json:"sourceHash" yaml:"sourceHash"`
}

