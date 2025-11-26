// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package filestorage


type FileStorageNetwork struct {
	// IP family, e.g. IPv4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.32.0/docs/resources/file_storage#family FileStorage#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// Attachment name (unique per this service).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.32.0/docs/resources/file_storage#name FileStorage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// UUID of an existing private network to attach.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.32.0/docs/resources/file_storage#uuid FileStorage#uuid}
	Uuid *string `field:"required" json:"uuid" yaml:"uuid"`
	// IP address to assign (optional, auto-assign otherwise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.32.0/docs/resources/file_storage#ip_address FileStorage#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

