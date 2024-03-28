// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gatewayconnection

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GatewayConnectionRemoteRouteList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GatewayConnectionRemoteRouteList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GatewayConnectionRemoteRouteList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GatewayConnectionRemoteRouteList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GatewayConnectionRemoteRouteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GatewayConnectionRemoteRouteList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GatewayConnectionRemoteRouteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGatewayConnectionRemoteRouteListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

