// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendrule


type LoadbalancerFrontendRuleMatchers struct {
	// body_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#body_size LoadbalancerFrontendRule#body_size}
	BodySize interface{} `field:"optional" json:"bodySize" yaml:"bodySize"`
	// body_size_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#body_size_range LoadbalancerFrontendRule#body_size_range}
	BodySizeRange interface{} `field:"optional" json:"bodySizeRange" yaml:"bodySizeRange"`
	// cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#cookie LoadbalancerFrontendRule#cookie}
	Cookie interface{} `field:"optional" json:"cookie" yaml:"cookie"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#header LoadbalancerFrontendRule#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#host LoadbalancerFrontendRule#host}
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// http_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#http_method LoadbalancerFrontendRule#http_method}
	HttpMethod interface{} `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// http_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#http_status LoadbalancerFrontendRule#http_status}
	HttpStatus interface{} `field:"optional" json:"httpStatus" yaml:"httpStatus"`
	// http_status_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#http_status_range LoadbalancerFrontendRule#http_status_range}
	HttpStatusRange interface{} `field:"optional" json:"httpStatusRange" yaml:"httpStatusRange"`
	// num_members_up block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#num_members_up LoadbalancerFrontendRule#num_members_up}
	NumMembersUp interface{} `field:"optional" json:"numMembersUp" yaml:"numMembersUp"`
	// path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#path LoadbalancerFrontendRule#path}
	Path interface{} `field:"optional" json:"path" yaml:"path"`
	// request_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#request_header LoadbalancerFrontendRule#request_header}
	RequestHeader interface{} `field:"optional" json:"requestHeader" yaml:"requestHeader"`
	// response_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#response_header LoadbalancerFrontendRule#response_header}
	ResponseHeader interface{} `field:"optional" json:"responseHeader" yaml:"responseHeader"`
	// src_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#src_ip LoadbalancerFrontendRule#src_ip}
	SrcIp interface{} `field:"optional" json:"srcIp" yaml:"srcIp"`
	// src_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#src_port LoadbalancerFrontendRule#src_port}
	SrcPort interface{} `field:"optional" json:"srcPort" yaml:"srcPort"`
	// src_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#src_port_range LoadbalancerFrontendRule#src_port_range}
	SrcPortRange interface{} `field:"optional" json:"srcPortRange" yaml:"srcPortRange"`
	// url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#url LoadbalancerFrontendRule#url}
	Url interface{} `field:"optional" json:"url" yaml:"url"`
	// url_param block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#url_param LoadbalancerFrontendRule#url_param}
	UrlParam interface{} `field:"optional" json:"urlParam" yaml:"urlParam"`
	// url_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/loadbalancer_frontend_rule#url_query LoadbalancerFrontendRule#url_query}
	UrlQuery interface{} `field:"optional" json:"urlQuery" yaml:"urlQuery"`
}

