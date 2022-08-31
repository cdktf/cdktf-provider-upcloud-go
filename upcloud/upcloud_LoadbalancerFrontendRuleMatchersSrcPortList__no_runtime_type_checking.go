//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersSrcPortList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersSrcPortList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersSrcPortList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersSrcPortList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersSrcPortList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersSrcPortList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadbalancerFrontendRuleMatchersSrcPortListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

