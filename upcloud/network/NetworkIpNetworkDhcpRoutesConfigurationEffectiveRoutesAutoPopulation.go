// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package network


type NetworkIpNetworkDhcpRoutesConfigurationEffectiveRoutesAutoPopulation struct {
	// Enable or disable route auto-population.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/network#enabled Network#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Exclude routes coming from specific sources (router-connected-networks, static-route).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/network#exclude_by_source Network#exclude_by_source}
	ExcludeBySource *[]*string `field:"optional" json:"excludeBySource" yaml:"excludeBySource"`
	// CIDR destinations to include when auto-populating routes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/network#filter_by_destination Network#filter_by_destination}
	FilterByDestination *[]*string `field:"optional" json:"filterByDestination" yaml:"filterByDestination"`
	// Include only routes of given types (service, user).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/network#filter_by_route_type Network#filter_by_route_type}
	FilterByRouteType *[]*string `field:"optional" json:"filterByRouteType" yaml:"filterByRouteType"`
}

