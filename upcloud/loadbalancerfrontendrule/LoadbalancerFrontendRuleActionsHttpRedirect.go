// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendrule


type LoadbalancerFrontendRuleActionsHttpRedirect struct {
	// Target location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/loadbalancer_frontend_rule#location LoadbalancerFrontendRule#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Target scheme.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/loadbalancer_frontend_rule#scheme LoadbalancerFrontendRule#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// HTTP status code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/loadbalancer_frontend_rule#status LoadbalancerFrontendRule#status}
	Status *float64 `field:"optional" json:"status" yaml:"status"`
}

