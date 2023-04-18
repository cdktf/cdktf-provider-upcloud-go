package loadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerConfig struct {
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
	// The name of the service must be unique within customer account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/loadbalancer#name Loadbalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Plan which the service will have. You can list available loadbalancer plans with `upctl loadbalancer plans`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/loadbalancer#plan Loadbalancer#plan}
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// Zone in which the service will be hosted, e.g. `fi-hel1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/loadbalancer#zone Loadbalancer#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// The service configured status indicates the service's current intended status. Managed by the customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/loadbalancer#configured_status Loadbalancer#configured_status}
	ConfiguredStatus *string `field:"optional" json:"configuredStatus" yaml:"configuredStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/loadbalancer#id Loadbalancer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key-value pairs to classify the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/loadbalancer#labels Loadbalancer#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Private network UUID where traffic will be routed. Must reside in load balancer zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/loadbalancer#network Loadbalancer#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/loadbalancer#networks Loadbalancer#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
}

