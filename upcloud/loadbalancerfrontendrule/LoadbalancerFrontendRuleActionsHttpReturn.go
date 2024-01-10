// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendrule


type LoadbalancerFrontendRuleActionsHttpReturn struct {
	// Content type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_frontend_rule#content_type LoadbalancerFrontendRule#content_type}
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// The payload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_frontend_rule#payload LoadbalancerFrontendRule#payload}
	Payload *string `field:"required" json:"payload" yaml:"payload"`
	// HTTP status code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.3.1/docs/resources/loadbalancer_frontend_rule#status LoadbalancerFrontendRule#status}
	Status *float64 `field:"required" json:"status" yaml:"status"`
}

