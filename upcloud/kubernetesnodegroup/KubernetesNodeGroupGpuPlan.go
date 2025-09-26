// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesnodegroup


type KubernetesNodeGroupGpuPlan struct {
	// The size of the storage device in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/kubernetes_node_group#storage_size KubernetesNodeGroup#storage_size}
	StorageSize *float64 `field:"optional" json:"storageSize" yaml:"storageSize"`
	// The storage tier to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/kubernetes_node_group#storage_tier KubernetesNodeGroup#storage_tier}
	StorageTier *string `field:"optional" json:"storageTier" yaml:"storageTier"`
}

