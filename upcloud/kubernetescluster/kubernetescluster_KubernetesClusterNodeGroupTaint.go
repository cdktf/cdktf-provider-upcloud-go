package kubernetescluster


type KubernetesClusterNodeGroupTaint struct {
	// Taint effect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/kubernetes_cluster#effect KubernetesCluster#effect}
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Taint key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/kubernetes_cluster#key KubernetesCluster#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Taint value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/kubernetes_cluster#value KubernetesCluster#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

