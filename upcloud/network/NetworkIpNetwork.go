// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package network


type NetworkIpNetwork struct {
	// The CIDR range of the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network#address Network#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Is DHCP enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network#dhcp Network#dhcp}
	Dhcp interface{} `field:"required" json:"dhcp" yaml:"dhcp"`
	// IP address family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network#family Network#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// Is the gateway the DHCP default route?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network#dhcp_default_route Network#dhcp_default_route}
	DhcpDefaultRoute interface{} `field:"optional" json:"dhcpDefaultRoute" yaml:"dhcpDefaultRoute"`
	// The DNS servers given by DHCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network#dhcp_dns Network#dhcp_dns}
	DhcpDns *[]*string `field:"optional" json:"dhcpDns" yaml:"dhcpDns"`
	// The additional DHCP classless static routes given by DHCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network#dhcp_routes Network#dhcp_routes}
	DhcpRoutes *[]*string `field:"optional" json:"dhcpRoutes" yaml:"dhcpRoutes"`
	// Gateway address given by DHCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network#gateway Network#gateway}
	Gateway *string `field:"optional" json:"gateway" yaml:"gateway"`
}

