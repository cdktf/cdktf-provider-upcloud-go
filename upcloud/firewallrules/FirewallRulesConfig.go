// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallrules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallRulesConfig struct {
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
	// The UUID of the server to be protected with the firewall rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/firewall_rules#server_id FirewallRules#server_id}
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// firewall_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/firewall_rules#firewall_rule FirewallRules#firewall_rule}
	FirewallRule interface{} `field:"optional" json:"firewallRule" yaml:"firewallRule"`
}

