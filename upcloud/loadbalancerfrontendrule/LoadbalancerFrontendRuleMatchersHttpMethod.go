package loadbalancerfrontendrule


type LoadbalancerFrontendRuleMatchersHttpMethod struct {
	// String value (`GET`, `HEAD`, `POST`, `PUT`, `PATCH`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/loadbalancer_frontend_rule#value LoadbalancerFrontendRule#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

