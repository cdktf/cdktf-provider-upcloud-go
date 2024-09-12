// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gatewayconnectiontunnel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GatewayConnectionTunnelConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// ID of the upcloud_gateway_connection resource to which the tunnel belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.11.0/docs/resources/gateway_connection_tunnel#connection_id GatewayConnectionTunnel#connection_id}
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// ipsec_auth_psk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.11.0/docs/resources/gateway_connection_tunnel#ipsec_auth_psk GatewayConnectionTunnel#ipsec_auth_psk}
	IpsecAuthPsk *GatewayConnectionTunnelIpsecAuthPsk `field:"required" json:"ipsecAuthPsk" yaml:"ipsecAuthPsk"`
	// Public (UpCloud) endpoint address of this tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.11.0/docs/resources/gateway_connection_tunnel#local_address_name GatewayConnectionTunnel#local_address_name}
	LocalAddressName *string `field:"required" json:"localAddressName" yaml:"localAddressName"`
	// The name of the tunnel, should be unique within the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.11.0/docs/resources/gateway_connection_tunnel#name GatewayConnectionTunnel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Remote public IP address of the tunnel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.11.0/docs/resources/gateway_connection_tunnel#remote_address GatewayConnectionTunnel#remote_address}
	RemoteAddress *string `field:"required" json:"remoteAddress" yaml:"remoteAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.11.0/docs/resources/gateway_connection_tunnel#id GatewayConnectionTunnel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ipsec_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.11.0/docs/resources/gateway_connection_tunnel#ipsec_properties GatewayConnectionTunnel#ipsec_properties}
	IpsecProperties *GatewayConnectionTunnelIpsecProperties `field:"optional" json:"ipsecProperties" yaml:"ipsecProperties"`
}

