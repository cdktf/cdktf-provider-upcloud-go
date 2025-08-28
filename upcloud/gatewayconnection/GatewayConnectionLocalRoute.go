// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gatewayconnection


type GatewayConnectionLocalRoute struct {
	// Name of the route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/gateway_connection#name GatewayConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Destination prefix of the route; needs to be a valid IPv4 prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/gateway_connection#static_network GatewayConnection#static_network}
	StaticNetwork *string `field:"required" json:"staticNetwork" yaml:"staticNetwork"`
	// Type of route; currently the only supported type is 'static'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/gateway_connection#type GatewayConnection#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

