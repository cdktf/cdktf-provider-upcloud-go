// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type LoadbalancerFrontendRuleActionsUseBackend struct {
	// The name of the backend where traffic will be routed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend_rule#backend_name LoadbalancerFrontendRule#backend_name}
	BackendName *string `field:"required" json:"backendName" yaml:"backendName"`
}

