// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendrule


type LoadbalancerFrontendRuleMatchersBodySize struct {
	// Match method (`equal`, `greater`, `greater_or_equal`, `less`, `less_or_equal`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.19.0/docs/resources/loadbalancer_frontend_rule#method LoadbalancerFrontendRule#method}
	Method *string `field:"required" json:"method" yaml:"method"`
	// Integer value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.19.0/docs/resources/loadbalancer_frontend_rule#value LoadbalancerFrontendRule#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// Defines if the condition should be inverted. Works similarly to logical NOT operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.19.0/docs/resources/loadbalancer_frontend_rule#inverse LoadbalancerFrontendRule#inverse}
	Inverse interface{} `field:"optional" json:"inverse" yaml:"inverse"`
}

