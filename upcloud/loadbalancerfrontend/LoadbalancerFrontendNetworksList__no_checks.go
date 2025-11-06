// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package loadbalancerfrontend

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LoadbalancerFrontendNetworksList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LoadbalancerFrontendNetworksList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LoadbalancerFrontendNetworksList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendNetworksList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendNetworksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendNetworksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LoadbalancerFrontendNetworksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLoadbalancerFrontendNetworksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

