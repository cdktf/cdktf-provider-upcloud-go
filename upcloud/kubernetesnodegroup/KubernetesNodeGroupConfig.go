// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesnodegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesNodeGroupConfig struct {
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
	// Cluster ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#cluster KubernetesNodeGroup#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the node group. Needs to be unique within a cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#name KubernetesNodeGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Amount of nodes to provision in the node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#node_count KubernetesNodeGroup#node_count}
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// The server plan used for the node group. You can list available plans with `upctl server plans`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#plan KubernetesNodeGroup#plan}
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// If set to true, nodes in this group will be placed on separate compute hosts.
	//
	// Please note that anti-affinity policy is considered "best effort" and enabling it does not fully guarantee that the nodes will end up on different hardware.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#anti_affinity KubernetesNodeGroup#anti_affinity}
	AntiAffinity interface{} `field:"optional" json:"antiAffinity" yaml:"antiAffinity"`
	// custom_plan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#custom_plan KubernetesNodeGroup#custom_plan}
	CustomPlan *KubernetesNodeGroupCustomPlan `field:"optional" json:"customPlan" yaml:"customPlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#id KubernetesNodeGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kubelet_args block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#kubelet_args KubernetesNodeGroup#kubelet_args}
	KubeletArgs interface{} `field:"optional" json:"kubeletArgs" yaml:"kubeletArgs"`
	// Key-value pairs to classify the node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#labels KubernetesNodeGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// You can optionally select SSH keys to be added as authorized keys to the nodes in this node group.
	//
	// This allows you to connect to the nodes via SSH once they are running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#ssh_keys KubernetesNodeGroup#ssh_keys}
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Storage encryption strategy for the nodes in this group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#storage_encryption KubernetesNodeGroup#storage_encryption}
	StorageEncryption *string `field:"optional" json:"storageEncryption" yaml:"storageEncryption"`
	// taint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#taint KubernetesNodeGroup#taint}
	Taint interface{} `field:"optional" json:"taint" yaml:"taint"`
	// If set to false, nodes in this group will not have access to utility network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/kubernetes_node_group#utility_network_access KubernetesNodeGroup#utility_network_access}
	UtilityNetworkAccess interface{} `field:"optional" json:"utilityNetworkAccess" yaml:"utilityNetworkAccess"`
}

