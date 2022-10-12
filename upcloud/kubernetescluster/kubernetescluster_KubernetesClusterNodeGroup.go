package kubernetescluster


type KubernetesClusterNodeGroup struct {
	// The name of the node group. Needs to be unique within a cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/kubernetes_cluster#name KubernetesCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Amount of nodes to provision in the node group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/kubernetes_cluster#count KubernetesCluster#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Additional arguments for kubelet for the nodes in this group.
	//
	// WARNING - those arguments will be passed directly to kubelet CLI on each worker node without any validation. Passing invalid arguments can break your whole cluster. Be extra careful when adding kubelet args.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/kubernetes_cluster#kubelet_args KubernetesCluster#kubelet_args}
	KubeletArgs *map[string]*string `field:"optional" json:"kubeletArgs" yaml:"kubeletArgs"`
	// Key-value pairs to classify the node group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/kubernetes_cluster#labels KubernetesCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The pricing plan used for the node group. Valid values available via `upcloud_kubernetes_plan` datasource field `description`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/kubernetes_cluster#plan KubernetesCluster#plan}
	Plan *string `field:"optional" json:"plan" yaml:"plan"`
	// You can optionally select SSH keys to be added as authorized keys to the nodes in this node group.
	//
	// This allows you to connect to the nodes via SSH once they are running.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/kubernetes_cluster#ssh_keys KubernetesCluster#ssh_keys}
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// taint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/kubernetes_cluster#taint KubernetesCluster#taint}
	Taint interface{} `field:"optional" json:"taint" yaml:"taint"`
}

