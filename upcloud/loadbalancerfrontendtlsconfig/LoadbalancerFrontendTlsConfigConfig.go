// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendtlsconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerFrontendTlsConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Reference to certificate bundle ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/loadbalancer_frontend_tls_config#certificate_bundle LoadbalancerFrontendTlsConfig#certificate_bundle}
	CertificateBundle *string `field:"required" json:"certificateBundle" yaml:"certificateBundle"`
	// ID of the load balancer frontend to which the TLS config is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/loadbalancer_frontend_tls_config#frontend LoadbalancerFrontendTlsConfig#frontend}
	Frontend *string `field:"required" json:"frontend" yaml:"frontend"`
	// The name of the TLS config must be unique within service frontend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/loadbalancer_frontend_tls_config#name LoadbalancerFrontendTlsConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/loadbalancer_frontend_tls_config#id LoadbalancerFrontendTlsConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

