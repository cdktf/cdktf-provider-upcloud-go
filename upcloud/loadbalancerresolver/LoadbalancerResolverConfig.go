// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerresolver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerResolverConfig struct {
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
	// Time in seconds to cache invalid results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_resolver#cache_invalid LoadbalancerResolver#cache_invalid}
	CacheInvalid *float64 `field:"required" json:"cacheInvalid" yaml:"cacheInvalid"`
	// Time in seconds to cache valid results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_resolver#cache_valid LoadbalancerResolver#cache_valid}
	CacheValid *float64 `field:"required" json:"cacheValid" yaml:"cacheValid"`
	// ID of the load balancer to which the resolver is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_resolver#loadbalancer LoadbalancerResolver#loadbalancer}
	Loadbalancer *string `field:"required" json:"loadbalancer" yaml:"loadbalancer"`
	// The name of the resolver must be unique within the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_resolver#name LoadbalancerResolver#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of nameserver IP addresses.
	//
	// Nameserver can reside in public internet or in customer private network.
	// 				Port is optional, if missing then default 53 will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_resolver#nameservers LoadbalancerResolver#nameservers}
	Nameservers *[]*string `field:"required" json:"nameservers" yaml:"nameservers"`
	// Number of retries on failure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_resolver#retries LoadbalancerResolver#retries}
	Retries *float64 `field:"required" json:"retries" yaml:"retries"`
	// Timeout for the query in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_resolver#timeout LoadbalancerResolver#timeout}
	Timeout *float64 `field:"required" json:"timeout" yaml:"timeout"`
	// Timeout for the query retries in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_resolver#timeout_retry LoadbalancerResolver#timeout_retry}
	TimeoutRetry *float64 `field:"required" json:"timeoutRetry" yaml:"timeoutRetry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_resolver#id LoadbalancerResolver#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

