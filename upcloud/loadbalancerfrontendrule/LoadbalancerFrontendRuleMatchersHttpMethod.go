// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendrule


type LoadbalancerFrontendRuleMatchersHttpMethod struct {
	// String value (`GET`, `HEAD`, `POST`, `PUT`, `PATCH`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/loadbalancer_frontend_rule#value LoadbalancerFrontendRule#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Defines if the condition should be inverted. Works similarly to logical NOT operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/loadbalancer_frontend_rule#inverse LoadbalancerFrontendRule#inverse}
	Inverse interface{} `field:"optional" json:"inverse" yaml:"inverse"`
}

