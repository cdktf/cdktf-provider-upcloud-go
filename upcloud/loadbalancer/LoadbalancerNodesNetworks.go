// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadbalancerNodesNetworks struct {
	// ip_addresses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/loadbalancer#ip_addresses Loadbalancer#ip_addresses}
	IpAddresses interface{} `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
}

