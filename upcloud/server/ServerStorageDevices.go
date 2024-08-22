// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server


type ServerStorageDevices struct {
	// A valid storage UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.1/docs/resources/server#storage Server#storage}
	Storage *string `field:"required" json:"storage" yaml:"storage"`
	// The device address the storage will be attached to (`scsi`|`virtio`|`ide`).
	//
	// Leave `address_position` field empty to auto-select next available address from that bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.1/docs/resources/server#address Server#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The device position in the given bus (defined via field `address`).
	//
	// Valid values for address `virtio` are `0-15` (`0`, for example). Valid values for `scsi` or `ide` are `0-1:0-1` (`0:0`, for example). Leave empty to auto-select next available address in the given bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.1/docs/resources/server#address_position Server#address_position}
	AddressPosition *string `field:"optional" json:"addressPosition" yaml:"addressPosition"`
	// The device type the storage will be attached as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.10.1/docs/resources/server#type Server#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

