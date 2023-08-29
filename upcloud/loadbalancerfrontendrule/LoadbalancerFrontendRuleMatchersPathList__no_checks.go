// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package loadbalancerfrontendrule

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersPathList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersPathList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersPathList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersPathList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersPathList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersPathList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadbalancerFrontendRuleMatchersPathListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

