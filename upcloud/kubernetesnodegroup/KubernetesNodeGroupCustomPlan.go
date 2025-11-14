// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesnodegroup


type KubernetesNodeGroupCustomPlan struct {
	// The number of CPU cores dedicated to individual node group nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/kubernetes_node_group#cores KubernetesNodeGroup#cores}
	Cores *float64 `field:"required" json:"cores" yaml:"cores"`
	// The amount of memory in megabytes to assign to individual node group node.
	//
	// Value needs to be divisible by 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/kubernetes_node_group#memory KubernetesNodeGroup#memory}
	Memory *float64 `field:"required" json:"memory" yaml:"memory"`
	// The size of the storage device in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/kubernetes_node_group#storage_size KubernetesNodeGroup#storage_size}
	StorageSize *float64 `field:"required" json:"storageSize" yaml:"storageSize"`
	// The storage tier to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/kubernetes_node_group#storage_tier KubernetesNodeGroup#storage_tier}
	StorageTier *string `field:"optional" json:"storageTier" yaml:"storageTier"`
}

