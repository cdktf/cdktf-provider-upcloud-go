// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gatewayconnectiontunnel


type GatewayConnectionTunnelIpsecAuthPsk struct {
	// The pre-shared key.
	//
	// This value is only used during resource creation and is not returned in the state. It is not possible to update this value. If you need to update it, delete the connection and create a new one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.1/docs/resources/gateway_connection_tunnel#psk GatewayConnectionTunnel#psk}
	Psk *string `field:"required" json:"psk" yaml:"psk"`
}

