// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallrules


type FirewallRulesFirewallRule struct {
	// Action to take if the rule conditions are met.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#action FirewallRules#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The direction of network traffic this rule will be applied to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#direction FirewallRules#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Freeform comment string for the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#comment FirewallRules#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The destination address range ends from this address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#destination_address_end FirewallRules#destination_address_end}
	DestinationAddressEnd *string `field:"optional" json:"destinationAddressEnd" yaml:"destinationAddressEnd"`
	// The destination address range starts from this address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#destination_address_start FirewallRules#destination_address_start}
	DestinationAddressStart *string `field:"optional" json:"destinationAddressStart" yaml:"destinationAddressStart"`
	// The destination port range ends from this port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#destination_port_end FirewallRules#destination_port_end}
	DestinationPortEnd *string `field:"optional" json:"destinationPortEnd" yaml:"destinationPortEnd"`
	// The destination port range starts from this port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#destination_port_start FirewallRules#destination_port_start}
	DestinationPortStart *string `field:"optional" json:"destinationPortStart" yaml:"destinationPortStart"`
	// The address family of new firewall rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#family FirewallRules#family}
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The ICMP type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#icmp_type FirewallRules#icmp_type}
	IcmpType *string `field:"optional" json:"icmpType" yaml:"icmpType"`
	// The protocol this rule will be applied to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#protocol FirewallRules#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The source address range ends from this address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#source_address_end FirewallRules#source_address_end}
	SourceAddressEnd *string `field:"optional" json:"sourceAddressEnd" yaml:"sourceAddressEnd"`
	// The source address range starts from this address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#source_address_start FirewallRules#source_address_start}
	SourceAddressStart *string `field:"optional" json:"sourceAddressStart" yaml:"sourceAddressStart"`
	// The source port range ends from this port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#source_port_end FirewallRules#source_port_end}
	SourcePortEnd *string `field:"optional" json:"sourcePortEnd" yaml:"sourcePortEnd"`
	// The source port range starts from this port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/3.1.1/docs/resources/firewall_rules#source_port_start FirewallRules#source_port_start}
	SourcePortStart *string `field:"optional" json:"sourcePortStart" yaml:"sourcePortStart"`
}

