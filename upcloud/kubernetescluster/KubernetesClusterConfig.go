// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterConfig struct {
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
	// IP addresses or IP ranges in CIDR format which are allowed to access the cluster control plane.
	//
	// To allow access from any source, use `["0.0.0.0/0"]`. To deny access from all sources, use `[]`. Values set here do not restrict access to node groups or exposed Kubernetes services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/kubernetes_cluster#control_plane_ip_filter KubernetesCluster#control_plane_ip_filter}
	ControlPlaneIpFilter *[]*string `field:"required" json:"controlPlaneIpFilter" yaml:"controlPlaneIpFilter"`
	// Cluster name. Needs to be unique within the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/kubernetes_cluster#name KubernetesCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Network ID for the cluster to run in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/kubernetes_cluster#network KubernetesCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Zone in which the Kubernetes cluster will be hosted, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/kubernetes_cluster#zone KubernetesCluster#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// User defined key-value pairs to classify the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/kubernetes_cluster#labels KubernetesCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The pricing plan used for the cluster. You can list available plans with `upctl kubernetes plans`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/kubernetes_cluster#plan KubernetesCluster#plan}
	Plan *string `field:"optional" json:"plan" yaml:"plan"`
	// Enable private node groups. Private node groups requires a network that is routed through NAT gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/kubernetes_cluster#private_node_groups KubernetesCluster#private_node_groups}
	PrivateNodeGroups interface{} `field:"optional" json:"privateNodeGroups" yaml:"privateNodeGroups"`
	// Set default storage encryption strategy for all nodes in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/kubernetes_cluster#storage_encryption KubernetesCluster#storage_encryption}
	StorageEncryption *string `field:"optional" json:"storageEncryption" yaml:"storageEncryption"`
	// The upgrade strategy to use when changing the cluster `version`.
	//
	// If not set, `manual` strategy will be used by default. When using `manual` strategy, you must replace the existing node-groups to update them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/kubernetes_cluster#upgrade_strategy_type KubernetesCluster#upgrade_strategy_type}
	UpgradeStrategyType *string `field:"optional" json:"upgradeStrategyType" yaml:"upgradeStrategyType"`
	// Kubernetes version ID, e.g. `1.31`. You can list available version IDs with `upctl kubernetes versions`.
	//
	// Note that when changing the cluster version, `upgrade_strategy` will be taken into account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/kubernetes_cluster#version KubernetesCluster#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

