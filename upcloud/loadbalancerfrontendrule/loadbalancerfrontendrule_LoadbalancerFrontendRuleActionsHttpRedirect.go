package loadbalancerfrontendrule


type LoadbalancerFrontendRuleActionsHttpRedirect struct {
	// Target location.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend_rule#location LoadbalancerFrontendRule#location}
	Location *string `field:"required" json:"location" yaml:"location"`
}

