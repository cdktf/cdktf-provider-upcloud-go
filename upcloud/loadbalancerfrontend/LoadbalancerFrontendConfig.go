// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontend

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerFrontendConfig struct {
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
	// The name of the default backend where traffic will be routed.
	//
	// Note, default backend can be overwritten in frontend rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.0.3/docs/resources/loadbalancer_frontend#default_backend_name LoadbalancerFrontend#default_backend_name}
	DefaultBackendName *string `field:"required" json:"defaultBackendName" yaml:"defaultBackendName"`
	// ID of the load balancer to which the frontend is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.0.3/docs/resources/loadbalancer_frontend#loadbalancer LoadbalancerFrontend#loadbalancer}
	Loadbalancer *string `field:"required" json:"loadbalancer" yaml:"loadbalancer"`
	// When load balancer operating in `tcp` mode it acts as a layer 4 proxy.
	//
	// In `http` mode it acts as a layer 7 proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.0.3/docs/resources/loadbalancer_frontend#mode LoadbalancerFrontend#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The name of the frontend must be unique within the load balancer service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.0.3/docs/resources/loadbalancer_frontend#name LoadbalancerFrontend#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Port to listen incoming requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.0.3/docs/resources/loadbalancer_frontend#port LoadbalancerFrontend#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.0.3/docs/resources/loadbalancer_frontend#id LoadbalancerFrontend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.0.3/docs/resources/loadbalancer_frontend#networks LoadbalancerFrontend#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.0.3/docs/resources/loadbalancer_frontend#properties LoadbalancerFrontend#properties}
	Properties *LoadbalancerFrontendProperties `field:"optional" json:"properties" yaml:"properties"`
}

