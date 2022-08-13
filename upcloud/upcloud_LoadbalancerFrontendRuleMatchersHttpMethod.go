// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type LoadbalancerFrontendRuleMatchersHttpMethod struct {
	// String value (`GET`, `HEAD`, `POST`, `PUT`, `PATCH`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend_rule#value LoadbalancerFrontendRule#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

