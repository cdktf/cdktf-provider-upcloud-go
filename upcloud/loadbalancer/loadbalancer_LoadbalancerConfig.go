package loadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer#name Loadbalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Private network UUID where traffic will be routed. Must reside in load balancer zone.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer#network Loadbalancer#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Plan which the service will have.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer#plan Loadbalancer#plan}
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// Zone in which the service will be hosted, e.g. `fi-hel1`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer#zone Loadbalancer#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// The service configured status indicates the service's current intended status. Managed by the customer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer#configured_status Loadbalancer#configured_status}
	ConfiguredStatus *string `field:"optional" json:"configuredStatus" yaml:"configuredStatus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer#id Loadbalancer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

