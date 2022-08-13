// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerBackendConfig struct {
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
	// ID of the load balancer to which the backend is connected.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_backend#loadbalancer LoadbalancerBackend#loadbalancer}
	Loadbalancer *string `field:"required" json:"loadbalancer" yaml:"loadbalancer"`
	// The name of the backend must be unique within the load balancer service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_backend#name LoadbalancerBackend#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_backend#id LoadbalancerBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_backend#properties LoadbalancerBackend#properties}
	Properties *LoadbalancerBackendProperties `field:"optional" json:"properties" yaml:"properties"`
	// Domain Name Resolver used with dynamic type members.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_backend#resolver_name LoadbalancerBackend#resolver_name}
	ResolverName *string `field:"optional" json:"resolverName" yaml:"resolverName"`
}

