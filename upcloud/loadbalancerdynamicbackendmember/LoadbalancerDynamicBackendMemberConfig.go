// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerdynamicbackendmember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerDynamicBackendMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/loadbalancer_dynamic_backend_member#backend LoadbalancerDynamicBackendMember#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Maximum number of sessions before queueing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/loadbalancer_dynamic_backend_member#max_sessions LoadbalancerDynamicBackendMember#max_sessions}
	MaxSessions *float64 `field:"required" json:"maxSessions" yaml:"maxSessions"`
	// The name of the member must be unique within the load balancer backend service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/loadbalancer_dynamic_backend_member#name LoadbalancerDynamicBackendMember#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Used to adjust the server's weight relative to other servers.
	//
	// All servers will receive a load proportional to their weight relative to the sum of all weights, so the higher the weight, the higher the load.
	// 				A value of 0 means the server will not participate in load balancing but will still accept persistent connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/loadbalancer_dynamic_backend_member#weight LoadbalancerDynamicBackendMember#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Indicates if the member is enabled. Disabled members are excluded from load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/loadbalancer_dynamic_backend_member#enabled LoadbalancerDynamicBackendMember#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/loadbalancer_dynamic_backend_member#id LoadbalancerDynamicBackendMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional fallback IP address in case of failure on DNS resolving.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/loadbalancer_dynamic_backend_member#ip LoadbalancerDynamicBackendMember#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// Server port. Port is optional and can be specified in DNS SRV record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/loadbalancer_dynamic_backend_member#port LoadbalancerDynamicBackendMember#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

