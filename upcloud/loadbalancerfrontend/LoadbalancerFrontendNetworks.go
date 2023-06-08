package loadbalancerfrontend


type LoadbalancerFrontendNetworks struct {
	// Name of the load balancer network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/loadbalancer_frontend#name LoadbalancerFrontend#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

