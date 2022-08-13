// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FloatingIpAddressConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Is address for utility or public network.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/floating_ip_address#access FloatingIpAddress#access}
	Access *string `field:"optional" json:"access" yaml:"access"`
	// The address family of new IP address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/floating_ip_address#family FloatingIpAddress#family}
	Family *string `field:"optional" json:"family" yaml:"family"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/floating_ip_address#id FloatingIpAddress#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// MAC address of server interface to assign address to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/floating_ip_address#mac_address FloatingIpAddress#mac_address}
	MacAddress *string `field:"optional" json:"macAddress" yaml:"macAddress"`
	// Zone of address, required when assigning a detached floating IP address.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/floating_ip_address#zone FloatingIpAddress#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

