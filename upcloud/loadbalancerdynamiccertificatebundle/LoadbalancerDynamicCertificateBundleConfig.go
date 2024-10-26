// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerdynamiccertificatebundle

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerDynamicCertificateBundleConfig struct {
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
	// Certificate hostnames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/loadbalancer_dynamic_certificate_bundle#hostnames LoadbalancerDynamicCertificateBundle#hostnames}
	Hostnames *[]*string `field:"required" json:"hostnames" yaml:"hostnames"`
	// Private key type (`rsa` / `ecdsa`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/loadbalancer_dynamic_certificate_bundle#key_type LoadbalancerDynamicCertificateBundle#key_type}
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The name of the certificate bundle. Must be unique within customer account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/loadbalancer_dynamic_certificate_bundle#name LoadbalancerDynamicCertificateBundle#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

