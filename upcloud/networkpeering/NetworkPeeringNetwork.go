// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpeering


type NetworkPeeringNetwork struct {
	// The UUID of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.0/docs/resources/network_peering#uuid NetworkPeering#uuid}
	Uuid *string `field:"required" json:"uuid" yaml:"uuid"`
}

