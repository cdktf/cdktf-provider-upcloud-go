// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadbalancerIpAddresses struct {
	// Floating IP address to attach to the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.30.0/docs/resources/loadbalancer#address Loadbalancer#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// Name of the network where to attach the IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.30.0/docs/resources/loadbalancer#network_name Loadbalancer#network_name}
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
}

