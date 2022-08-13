// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type NetworkIpNetwork struct {
	// The CIDR range of the subnet.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/network#address Network#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Is DHCP enabled?
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/network#dhcp Network#dhcp}
	Dhcp interface{} `field:"required" json:"dhcp" yaml:"dhcp"`
	// IP address family.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/network#family Network#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// Is the gateway the DHCP default route?
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/network#dhcp_default_route Network#dhcp_default_route}
	DhcpDefaultRoute interface{} `field:"optional" json:"dhcpDefaultRoute" yaml:"dhcpDefaultRoute"`
	// The DNS servers given by DHCP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/network#dhcp_dns Network#dhcp_dns}
	DhcpDns *[]*string `field:"optional" json:"dhcpDns" yaml:"dhcpDns"`
	// Gateway address given by DHCP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/network#gateway Network#gateway}
	Gateway *string `field:"optional" json:"gateway" yaml:"gateway"`
}

