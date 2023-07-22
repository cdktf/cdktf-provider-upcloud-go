package loadbalancerfrontendrule


type LoadbalancerFrontendRuleMatchersSrcPortRange struct {
	// Integer value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.12.0/docs/resources/loadbalancer_frontend_rule#range_end LoadbalancerFrontendRule#range_end}
	RangeEnd *float64 `field:"required" json:"rangeEnd" yaml:"rangeEnd"`
	// Integer value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.12.0/docs/resources/loadbalancer_frontend_rule#range_start LoadbalancerFrontendRule#range_start}
	RangeStart *float64 `field:"required" json:"rangeStart" yaml:"rangeStart"`
}

