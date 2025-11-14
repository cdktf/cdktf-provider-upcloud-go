// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filestorage


type FileStorageShareAcl struct {
	// Access level: 'ro' or 'rw'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#permission FileStorage#permission}
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Target IP/CIDR or '*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/file_storage#target FileStorage#target}
	Target *string `field:"required" json:"target" yaml:"target"`
}

