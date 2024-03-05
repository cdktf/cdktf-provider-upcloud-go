// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerbackendtlsconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerBackendTlsConfigConfig struct {
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
	// ID of the load balancer backend to which the TLS config is connected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.2/docs/resources/loadbalancer_backend_tls_config#backend LoadbalancerBackendTlsConfig#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Reference to certificate bundle ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.2/docs/resources/loadbalancer_backend_tls_config#certificate_bundle LoadbalancerBackendTlsConfig#certificate_bundle}
	CertificateBundle *string `field:"required" json:"certificateBundle" yaml:"certificateBundle"`
	// The name of the TLS config must be unique within service backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.2/docs/resources/loadbalancer_backend_tls_config#name LoadbalancerBackendTlsConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.2/docs/resources/loadbalancer_backend_tls_config#id LoadbalancerBackendTlsConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

