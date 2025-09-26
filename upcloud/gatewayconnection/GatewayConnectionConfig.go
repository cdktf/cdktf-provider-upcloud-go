// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gatewayconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GatewayConnectionConfig struct {
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
	// The ID of the upcloud_gateway resource to which the connection belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/gateway_connection#gateway GatewayConnection#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
	// The name of the connection, should be unique within the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/gateway_connection#name GatewayConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/gateway_connection#id GatewayConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// local_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/gateway_connection#local_route GatewayConnection#local_route}
	LocalRoute interface{} `field:"optional" json:"localRoute" yaml:"localRoute"`
	// remote_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/gateway_connection#remote_route GatewayConnection#remote_route}
	RemoteRoute interface{} `field:"optional" json:"remoteRoute" yaml:"remoteRoute"`
	// The type of the connection; currently the only supported type is 'ipsec'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/gateway_connection#type GatewayConnection#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

