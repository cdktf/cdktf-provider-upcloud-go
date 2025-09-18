// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancermanualcertificatebundle

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerManualCertificateBundleConfig struct {
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
	// Certificate as base64 encoded string. Must be in PEM format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.26.0/docs/resources/loadbalancer_manual_certificate_bundle#certificate LoadbalancerManualCertificateBundle#certificate}
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// The name of the certificate bundle. Must be unique within customer account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.26.0/docs/resources/loadbalancer_manual_certificate_bundle#name LoadbalancerManualCertificateBundle#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Private key as base64 encoded string. Must be in PEM format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.26.0/docs/resources/loadbalancer_manual_certificate_bundle#private_key LoadbalancerManualCertificateBundle#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Intermediate certificates as base64 encoded string. Must be in PEM format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.26.0/docs/resources/loadbalancer_manual_certificate_bundle#intermediates LoadbalancerManualCertificateBundle#intermediates}
	Intermediates *string `field:"optional" json:"intermediates" yaml:"intermediates"`
}

