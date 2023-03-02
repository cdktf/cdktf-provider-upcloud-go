package loadbalancerfrontend


type LoadbalancerFrontendNetworks struct {
	// Name of the load balancer network.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend#name LoadbalancerFrontend#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

