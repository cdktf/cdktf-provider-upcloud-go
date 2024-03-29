// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package firewallrules

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirewallRulesFirewallRuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FirewallRulesFirewallRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FirewallRulesFirewallRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FirewallRulesFirewallRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FirewallRulesFirewallRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FirewallRulesFirewallRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FirewallRulesFirewallRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFirewallRulesFirewallRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

