// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesnodegroup


type KubernetesNodeGroupKubeletArgs struct {
	// Kubelet argument key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.0.3/docs/resources/kubernetes_node_group#key KubernetesNodeGroup#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Kubelet argument value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.0.3/docs/resources/kubernetes_node_group#value KubernetesNodeGroup#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

