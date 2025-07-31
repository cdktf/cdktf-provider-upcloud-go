// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendrule


type LoadbalancerFrontendRuleMatchersUrl struct {
	// Match method (`exact`, `substring`, `regexp`, `starts`, `ends`, `domain`, `ip`, `exists`).
	//
	// Matcher with `exists` and `ip` methods must be used without `value` and `ignore_case` fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/loadbalancer_frontend_rule#method LoadbalancerFrontendRule#method}
	Method *string `field:"required" json:"method" yaml:"method"`
	// Defines if case should be ignored. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/loadbalancer_frontend_rule#ignore_case LoadbalancerFrontendRule#ignore_case}
	IgnoreCase interface{} `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// Defines if the condition should be inverted. Works similarly to logical NOT operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/loadbalancer_frontend_rule#inverse LoadbalancerFrontendRule#inverse}
	Inverse interface{} `field:"optional" json:"inverse" yaml:"inverse"`
	// String value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/loadbalancer_frontend_rule#value LoadbalancerFrontendRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

