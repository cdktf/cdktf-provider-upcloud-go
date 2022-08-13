// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type ServerNetworkInterface struct {
	// Network interface type. For private network interfaces, a network must be specified with an existing network id.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#type Server#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// `true` if this interface should be used for network booting.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#bootable Server#bootable}
	Bootable interface{} `field:"optional" json:"bootable" yaml:"bootable"`
	// The assigned IP address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#ip_address Server#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// The IP address type of this interface (one of `IPv4` or `IPv6`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#ip_address_family Server#ip_address_family}
	IpAddressFamily *string `field:"optional" json:"ipAddressFamily" yaml:"ipAddressFamily"`
	// The unique ID of a network to attach this network to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#network Server#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// `true` if source IP should be filtered.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#source_ip_filtering Server#source_ip_filtering}
	SourceIpFiltering interface{} `field:"optional" json:"sourceIpFiltering" yaml:"sourceIpFiltering"`
}

