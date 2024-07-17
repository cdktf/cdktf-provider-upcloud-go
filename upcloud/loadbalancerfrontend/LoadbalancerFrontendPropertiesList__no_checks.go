// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package loadbalancerfrontend

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadbalancerFrontendPropertiesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoadbalancerFrontendPropertiesList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadbalancerFrontendPropertiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendPropertiesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendPropertiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendPropertiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendPropertiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadbalancerFrontendPropertiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

