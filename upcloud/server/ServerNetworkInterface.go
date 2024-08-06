// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server


type ServerNetworkInterface struct {
	// Network interface type. For private network interfaces, a network must be specified with an existing network id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/server#type Server#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// additional_ip_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/server#additional_ip_address Server#additional_ip_address}
	AdditionalIpAddress interface{} `field:"optional" json:"additionalIpAddress" yaml:"additionalIpAddress"`
	// `true` if this interface should be used for network booting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/server#bootable Server#bootable}
	Bootable interface{} `field:"optional" json:"bootable" yaml:"bootable"`
	// The assigned primary IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/server#ip_address Server#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The type of the primary IP address of this interface (one of `IPv4` or `IPv6`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/server#ip_address_family Server#ip_address_family}
	IpAddressFamily *string `field:"optional" json:"ipAddressFamily" yaml:"ipAddressFamily"`
	// The unique ID of a network to attach this network to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/server#network Server#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// `true` if source IP should be filtered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/server#source_ip_filtering Server#source_ip_filtering}
	SourceIpFiltering interface{} `field:"optional" json:"sourceIpFiltering" yaml:"sourceIpFiltering"`
}

