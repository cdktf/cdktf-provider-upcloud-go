// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gateway

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GatewayAddressesList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GatewayAddressesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GatewayAddressesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GatewayAddressesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GatewayAddressesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGatewayAddressesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

