// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontend


type LoadbalancerFrontendProperties struct {
	// Enable or disable HTTP/2 support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/resources/loadbalancer_frontend#http2_enabled LoadbalancerFrontend#http2_enabled}
	Http2Enabled interface{} `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// Enable or disable inbound proxy protocol support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/resources/loadbalancer_frontend#inbound_proxy_protocol LoadbalancerFrontend#inbound_proxy_protocol}
	InboundProxyProtocol interface{} `field:"optional" json:"inboundProxyProtocol" yaml:"inboundProxyProtocol"`
	// Client request timeout in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/resources/loadbalancer_frontend#timeout_client LoadbalancerFrontend#timeout_client}
	TimeoutClient *float64 `field:"optional" json:"timeoutClient" yaml:"timeoutClient"`
}

