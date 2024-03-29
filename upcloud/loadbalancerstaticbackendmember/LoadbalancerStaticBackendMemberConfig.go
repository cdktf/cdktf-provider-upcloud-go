// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerstaticbackendmember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerStaticBackendMemberConfig struct {
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
	// ID of the load balancer backend to which the member is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.1/docs/resources/loadbalancer_static_backend_member#backend LoadbalancerStaticBackendMember#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Server IP address in the customer private network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.1/docs/resources/loadbalancer_static_backend_member#ip LoadbalancerStaticBackendMember#ip}
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Maximum number of sessions before queueing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.1/docs/resources/loadbalancer_static_backend_member#max_sessions LoadbalancerStaticBackendMember#max_sessions}
	MaxSessions *float64 `field:"required" json:"maxSessions" yaml:"maxSessions"`
	// The name of the member must be unique within the load balancer backend service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.1/docs/resources/loadbalancer_static_backend_member#name LoadbalancerStaticBackendMember#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Server port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.1/docs/resources/loadbalancer_static_backend_member#port LoadbalancerStaticBackendMember#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Used to adjust the server's weight relative to other servers.
	//
	// All servers will receive a load proportional to their weight relative to the sum of all weights, so the higher the weight, the higher the load.
	// 				A value of 0 means the server will not participate in load balancing but will still accept persistent connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.1/docs/resources/loadbalancer_static_backend_member#weight LoadbalancerStaticBackendMember#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Indicates if the member is enabled. Disabled members are excluded from load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.1/docs/resources/loadbalancer_static_backend_member#enabled LoadbalancerStaticBackendMember#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.1/docs/resources/loadbalancer_static_backend_member#id LoadbalancerStaticBackendMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

