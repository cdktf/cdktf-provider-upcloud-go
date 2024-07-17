// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontend


type LoadbalancerFrontendNetworks struct {
	// Name of the load balancer network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.8.0/docs/resources/loadbalancer_frontend#name LoadbalancerFrontend#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

