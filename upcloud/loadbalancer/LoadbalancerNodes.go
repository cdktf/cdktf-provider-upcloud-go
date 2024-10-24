// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer


type LoadbalancerNodes struct {
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.0/docs/resources/loadbalancer#networks Loadbalancer#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
}

