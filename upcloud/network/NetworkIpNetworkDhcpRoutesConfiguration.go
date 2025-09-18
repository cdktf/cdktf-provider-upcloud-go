// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package network


type NetworkIpNetworkDhcpRoutesConfiguration struct {
	// Automatically populate effective routes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.26.0/docs/resources/network#effective_routes_auto_population Network#effective_routes_auto_population}
	EffectiveRoutesAutoPopulation *NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulation `field:"optional" json:"effectiveRoutesAutoPopulation" yaml:"effectiveRoutesAutoPopulation"`
}

