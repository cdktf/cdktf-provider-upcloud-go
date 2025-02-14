// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package floatingipaddress

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FloatingIpAddressConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Network access for the floating IP address. Supported value: `public`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.19.0/docs/resources/floating_ip_address#access FloatingIpAddress#access}
	Access *string `field:"optional" json:"access" yaml:"access"`
	// The address family of the floating IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.19.0/docs/resources/floating_ip_address#family FloatingIpAddress#family}
	Family *string `field:"optional" json:"family" yaml:"family"`
	// MAC address of a server interface to assign address to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.19.0/docs/resources/floating_ip_address#mac_address FloatingIpAddress#mac_address}
	MacAddress *string `field:"optional" json:"macAddress" yaml:"macAddress"`
	// Zone of the address, e.g. `de-fra1`. Required when assigning a detached floating IP address. You can list available zones with `upctl zone list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.19.0/docs/resources/floating_ip_address#zone FloatingIpAddress#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

