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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.1/docs/resources/loadbalancer_frontend#default_backend_name LoadbalancerFrontend#default_backend_name}
	DefaultBackendName *string `field:"required" json:"defaultBackendName" yaml:"defaultBackendName"`
	// UUID of the load balancer to which the frontend is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.1/docs/resources/loadbalancer_frontend#loadbalancer LoadbalancerFrontend#loadbalancer}
	Loadbalancer *string `field:"required" json:"loadbalancer" yaml:"loadbalancer"`
	// When load balancer operating in `tcp` mode it acts as a layer 4 proxy.
	//
	// In `http` mode it acts as a layer 7 proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.1/docs/resources/loadbalancer_frontend#mode LoadbalancerFrontend#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The name of the frontend. Must be unique within the load balancer service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.1/docs/resources/loadbalancer_frontend#name LoadbalancerFrontend#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Port to listen for incoming requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.1/docs/resources/loadbalancer_frontend#port LoadbalancerFrontend#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.1/docs/resources/loadbalancer_frontend#networks LoadbalancerFrontend#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.1/docs/resources/loadbalancer_frontend#properties LoadbalancerFrontend#properties}
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

