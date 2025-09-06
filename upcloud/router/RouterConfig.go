// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package router

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RouterConfig struct {
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
	// Name of the router.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/router#name Router#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// User defined key-value pairs to classify the router.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/router#labels Router#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// static_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.2/docs/resources/router#static_route Router#static_route}
	StaticRoute interface{} `field:"optional" json:"staticRoute" yaml:"staticRoute"`
}

