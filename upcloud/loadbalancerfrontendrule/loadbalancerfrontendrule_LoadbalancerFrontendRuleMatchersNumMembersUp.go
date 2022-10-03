package loadbalancerfrontendrule


type LoadbalancerFrontendRuleMatchersNumMembersUp struct {
	// The name of the `backend` which members will be monitored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend_rule#backend_name LoadbalancerFrontendRule#backend_name}
	BackendName *string `field:"required" json:"backendName" yaml:"backendName"`
	// Match method (`equal`, `greater`, `greater_or_equal`, `less`, `less_or_equal`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend_rule#method LoadbalancerFrontendRule#method}
	Method *string `field:"required" json:"method" yaml:"method"`
	// Integer value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend_rule#value LoadbalancerFrontendRule#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

