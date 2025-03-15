// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerFrontendRuleConfig struct {
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
	// ID of the load balancer frontend to which the frontend rule is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#frontend LoadbalancerFrontendRule#frontend}
	Frontend *string `field:"required" json:"frontend" yaml:"frontend"`
	// The name of the frontend rule. Must be unique within the frontend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#name LoadbalancerFrontendRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Rule with the higher priority goes first. Rules with the same priority processed in alphabetical order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#priority LoadbalancerFrontendRule#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#actions LoadbalancerFrontendRule#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// matchers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#matchers LoadbalancerFrontendRule#matchers}
	Matchers interface{} `field:"optional" json:"matchers" yaml:"matchers"`
	// Defines boolean operator used to combine multiple matchers. Defaults to `and`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#matching_condition LoadbalancerFrontendRule#matching_condition}
	MatchingCondition *string `field:"optional" json:"matchingCondition" yaml:"matchingCondition"`
}

