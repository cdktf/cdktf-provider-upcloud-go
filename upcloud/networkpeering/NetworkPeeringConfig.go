// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpeering

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkPeeringConfig struct {
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
	// Name of the network peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network_peering#name NetworkPeering#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Configured status of the network peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network_peering#configured_status NetworkPeering#configured_status}
	ConfiguredStatus *string `field:"optional" json:"configuredStatus" yaml:"configuredStatus"`
	// User defined key-value pairs to classify the network peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network_peering#labels NetworkPeering#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network_peering#network NetworkPeering#network}
	Network interface{} `field:"optional" json:"network" yaml:"network"`
	// peer_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/network_peering#peer_network NetworkPeering#peer_network}
	PeerNetwork interface{} `field:"optional" json:"peerNetwork" yaml:"peerNetwork"`
}

