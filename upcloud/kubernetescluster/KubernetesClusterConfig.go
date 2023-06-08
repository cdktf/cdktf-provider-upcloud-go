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
	// Cluster name. Needs to be unique within the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/kubernetes_cluster#name KubernetesCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Network ID for the cluster to run in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/kubernetes_cluster#network KubernetesCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Zone in which the Kubernetes cluster will be hosted, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/kubernetes_cluster#zone KubernetesCluster#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/kubernetes_cluster#id KubernetesCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The pricing plan used for the cluster.
	//
	// Default plan is `development`. You can list available plans with `upctl kubernetes plans`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/kubernetes_cluster#plan KubernetesCluster#plan}
	Plan *string `field:"optional" json:"plan" yaml:"plan"`
	// Enable private node groups. Private node groups requires a network that is routed through NAT gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/kubernetes_cluster#private_node_groups KubernetesCluster#private_node_groups}
	PrivateNodeGroups interface{} `field:"optional" json:"privateNodeGroups" yaml:"privateNodeGroups"`
}

