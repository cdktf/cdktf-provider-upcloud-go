// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server


type ServerStorageDevices struct {
	// A valid storage UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.4.0/docs/resources/server#storage Server#storage}
	Storage *string `field:"required" json:"storage" yaml:"storage"`
	// The device address the storage will be attached to (`scsi`|`virtio`|`ide`).
	//
	// Leave `address_position` field empty to auto-select next available address from that bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.4.0/docs/resources/server#address Server#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The device position in the given bus (defined via field `address`).
	//
	// For example `0:0`, or `0`. Leave empty to auto-select next available address in the given bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.4.0/docs/resources/server#address_position Server#address_position}
	AddressPosition *string `field:"optional" json:"addressPosition" yaml:"addressPosition"`
	// The device type the storage will be attached as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.4.0/docs/resources/server#type Server#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

