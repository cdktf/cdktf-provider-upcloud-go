// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package router


type RouterStaticRoute struct {
	// Next hop address.
	//
	// NOTE: For static route to be active the next hop has to be an address of a reachable running Cloud Server in one of the Private Networks attached to the router.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.1.1/docs/resources/router#nexthop Router#nexthop}
	Nexthop *string `field:"required" json:"nexthop" yaml:"nexthop"`
	// Destination prefix of the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.1.1/docs/resources/router#route Router#route}
	Route *string `field:"required" json:"route" yaml:"route"`
	// Name or description of the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.1.1/docs/resources/router#name Router#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

