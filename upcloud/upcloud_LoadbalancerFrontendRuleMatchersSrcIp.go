// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type LoadbalancerFrontendRuleMatchersSrcIp struct {
	// IP address. CIDR masks are supported, e.g. `192.168.0.0/24`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend_rule#value LoadbalancerFrontendRule#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

