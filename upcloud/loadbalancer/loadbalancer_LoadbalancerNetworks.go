package loadbalancer


type LoadbalancerNetworks struct {
	// Network family. Currently only `IPv4` is supported.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer#family Loadbalancer#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// The name of the network must be unique within the service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer#name Loadbalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the network.
	//
	// Only one public network can be attached and at least one private network must be attached.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer#type Loadbalancer#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Private network UUID.
	//
	// Required for private networks and must reside in loadbalancer zone. For public network the field should be omitted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer#network Loadbalancer#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
}

