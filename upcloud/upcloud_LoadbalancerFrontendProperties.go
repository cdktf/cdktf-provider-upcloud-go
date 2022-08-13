// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type LoadbalancerFrontendProperties struct {
	// Enable or disable inbound proxy protocol support.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend#inbound_proxy_protocol LoadbalancerFrontend#inbound_proxy_protocol}
	InboundProxyProtocol interface{} `field:"optional" json:"inboundProxyProtocol" yaml:"inboundProxyProtocol"`
	// Client request timeout in seconds.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_frontend#timeout_client LoadbalancerFrontend#timeout_client}
	TimeoutClient *float64 `field:"optional" json:"timeoutClient" yaml:"timeoutClient"`
}

