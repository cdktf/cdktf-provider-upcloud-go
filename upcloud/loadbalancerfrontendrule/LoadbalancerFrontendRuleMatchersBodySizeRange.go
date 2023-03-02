package loadbalancerfrontendrule


type LoadbalancerFrontendRuleMatchersBodySizeRange struct {
	// Integer value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend_rule#range_end LoadbalancerFrontendRule#range_end}
	RangeEnd *float64 `field:"required" json:"rangeEnd" yaml:"rangeEnd"`
	// Integer value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend_rule#range_start LoadbalancerFrontendRule#range_start}
	RangeStart *float64 `field:"required" json:"rangeStart" yaml:"rangeStart"`
}

