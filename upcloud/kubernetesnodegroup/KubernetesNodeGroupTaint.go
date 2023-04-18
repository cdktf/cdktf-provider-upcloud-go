package kubernetesnodegroup


type KubernetesNodeGroupTaint struct {
	// Taint effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/kubernetes_node_group#effect KubernetesNodeGroup#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Taint key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/kubernetes_node_group#key KubernetesNodeGroup#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Taint value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/kubernetes_node_group#value KubernetesNodeGroup#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

