// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package network

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkIpNetworkList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkIpNetworkList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkIpNetworkList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkIpNetworkList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkIpNetworkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkIpNetworkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkIpNetworkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkIpNetworkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

