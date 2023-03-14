package gateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GatewayConfig struct {
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
	// Features enabled for the gateway.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/gateway#features Gateway#features}
	Features *[]*string `field:"required" json:"features" yaml:"features"`
	// Gateway name. Needs to be unique within the account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/gateway#name Gateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// router block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/gateway#router Gateway#router}
	Router *GatewayRouter `field:"required" json:"router" yaml:"router"`
	// Zone in which the gateway will be hosted, e.g. `de-fra1`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/gateway#zone Gateway#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// The service configured status indicates the service's current intended status. Managed by the customer.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/gateway#configured_status Gateway#configured_status}
	ConfiguredStatus *string `field:"optional" json:"configuredStatus" yaml:"configuredStatus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/gateway#id Gateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key-value pairs to classify the network gateway.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/gateway#labels Gateway#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

