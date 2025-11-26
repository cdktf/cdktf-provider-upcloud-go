// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package network

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkIpNetworkDhcpEffectiveRoutesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkIpNetworkDhcpEffectiveRoutesList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkIpNetworkDhcpEffectiveRoutesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkIpNetworkDhcpEffectiveRoutesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkIpNetworkDhcpEffectiveRoutesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkIpNetworkDhcpEffectiveRoutesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkIpNetworkDhcpEffectiveRoutesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

