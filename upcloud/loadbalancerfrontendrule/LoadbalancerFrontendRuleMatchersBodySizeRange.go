// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendrule


type LoadbalancerFrontendRuleMatchersBodySizeRange struct {
	// Integer value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/loadbalancer_frontend_rule#range_end LoadbalancerFrontendRule#range_end}
	RangeEnd *float64 `field:"required" json:"rangeEnd" yaml:"rangeEnd"`
	// Integer value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/loadbalancer_frontend_rule#range_start LoadbalancerFrontendRule#range_start}
	RangeStart *float64 `field:"required" json:"rangeStart" yaml:"rangeStart"`
}

