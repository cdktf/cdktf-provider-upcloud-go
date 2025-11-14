// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendrule


type LoadbalancerFrontendRuleMatchersSrcIp struct {
	// IP address. CIDR masks are supported, e.g. `192.168.0.0/24`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/loadbalancer_frontend_rule#value LoadbalancerFrontendRule#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Defines if the condition should be inverted. Works similarly to logical NOT operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.1/docs/resources/loadbalancer_frontend_rule#inverse LoadbalancerFrontendRule#inverse}
	Inverse interface{} `field:"optional" json:"inverse" yaml:"inverse"`
}

