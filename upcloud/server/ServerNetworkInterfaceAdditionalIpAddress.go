// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server


type ServerNetworkInterfaceAdditionalIpAddress struct {
	// An additional IP address for this interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/server#ip_address Server#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The type of the additional IP address of this interface (one of `IPv4` or `IPv6`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/server#ip_address_family Server#ip_address_family}
	IpAddressFamily *string `field:"optional" json:"ipAddressFamily" yaml:"ipAddressFamily"`
}

