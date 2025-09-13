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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.25.0/docs/resources/loadbalancer_static_backend_member#backend LoadbalancerStaticBackendMember#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Maximum number of sessions before queueing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.25.0/docs/resources/loadbalancer_static_backend_member#max_sessions LoadbalancerStaticBackendMember#max_sessions}
	MaxSessions *float64 `field:"required" json:"maxSessions" yaml:"maxSessions"`
	// The name of the member. Must be unique within within the load balancer backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.25.0/docs/resources/loadbalancer_static_backend_member#name LoadbalancerStaticBackendMember#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Weight of the member. The higher the weight, the more traffic the member receives.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.25.0/docs/resources/loadbalancer_static_backend_member#weight LoadbalancerStaticBackendMember#weight}
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Indicates if the member is enabled. Disabled members are excluded from load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.25.0/docs/resources/loadbalancer_static_backend_member#enabled LoadbalancerStaticBackendMember#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Optional fallback IP address in case of failure on DNS resolving.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.25.0/docs/resources/loadbalancer_static_backend_member#ip LoadbalancerStaticBackendMember#ip}
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// Server port. Port is optional and can be specified in DNS SRV record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.25.0/docs/resources/loadbalancer_static_backend_member#port LoadbalancerStaticBackendMember#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

