//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package loadbalancerfrontendrule

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersHostList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersHostList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersHostList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersHostList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersHostList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersHostList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadbalancerFrontendRuleMatchersHostListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

