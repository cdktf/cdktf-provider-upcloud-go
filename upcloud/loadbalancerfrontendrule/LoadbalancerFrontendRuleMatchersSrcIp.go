package loadbalancerfrontendrule


type LoadbalancerFrontendRuleMatchersSrcIp struct {
	// IP address. CIDR masks are supported, e.g. `192.168.0.0/24`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.12.0/docs/resources/loadbalancer_frontend_rule#value LoadbalancerFrontendRule#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

