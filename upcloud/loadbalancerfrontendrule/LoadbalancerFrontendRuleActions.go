// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendrule


type LoadbalancerFrontendRuleActions struct {
	// http_redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#http_redirect LoadbalancerFrontendRule#http_redirect}
	HttpRedirect interface{} `field:"optional" json:"httpRedirect" yaml:"httpRedirect"`
	// http_return block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#http_return LoadbalancerFrontendRule#http_return}
	HttpReturn interface{} `field:"optional" json:"httpReturn" yaml:"httpReturn"`
	// set_forwarded_headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#set_forwarded_headers LoadbalancerFrontendRule#set_forwarded_headers}
	SetForwardedHeaders interface{} `field:"optional" json:"setForwardedHeaders" yaml:"setForwardedHeaders"`
	// set_request_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#set_request_header LoadbalancerFrontendRule#set_request_header}
	SetRequestHeader interface{} `field:"optional" json:"setRequestHeader" yaml:"setRequestHeader"`
	// set_response_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#set_response_header LoadbalancerFrontendRule#set_response_header}
	SetResponseHeader interface{} `field:"optional" json:"setResponseHeader" yaml:"setResponseHeader"`
	// tcp_reject block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#tcp_reject LoadbalancerFrontendRule#tcp_reject}
	TcpReject interface{} `field:"optional" json:"tcpReject" yaml:"tcpReject"`
	// use_backend block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.4/docs/resources/loadbalancer_frontend_rule#use_backend LoadbalancerFrontendRule#use_backend}
	UseBackend interface{} `field:"optional" json:"useBackend" yaml:"useBackend"`
}

