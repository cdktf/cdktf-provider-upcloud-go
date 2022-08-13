// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerManualCertificateBundleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Certificate within base64 string must be in PEM format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_manual_certificate_bundle#certificate LoadbalancerManualCertificateBundle#certificate}
	Certificate *string `field:"required" json:"certificate" yaml:"certificate"`
	// The name of the bundle must be unique within customer account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_manual_certificate_bundle#name LoadbalancerManualCertificateBundle#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Private key within base64 string must be in PEM format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_manual_certificate_bundle#private_key LoadbalancerManualCertificateBundle#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_manual_certificate_bundle#id LoadbalancerManualCertificateBundle#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Intermediate certificates within base64 string must be in PEM format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/loadbalancer_manual_certificate_bundle#intermediates LoadbalancerManualCertificateBundle#intermediates}
	Intermediates *string `field:"optional" json:"intermediates" yaml:"intermediates"`
}

