package loadbalancerfrontendrule


type LoadbalancerFrontendRuleActionsUseBackend struct {
	// The name of the backend where traffic will be routed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.10.0/docs/resources/loadbalancer_frontend_rule#backend_name LoadbalancerFrontendRule#backend_name}
	BackendName *string `field:"required" json:"backendName" yaml:"backendName"`
}

