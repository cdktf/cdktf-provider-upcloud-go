// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package server

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServerNetworkInterfaceAdditionalIpAddressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServerNetworkInterfaceAdditionalIpAddressList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServerNetworkInterfaceAdditionalIpAddressList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServerNetworkInterfaceAdditionalIpAddressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServerNetworkInterfaceAdditionalIpAddressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServerNetworkInterfaceAdditionalIpAddressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServerNetworkInterfaceAdditionalIpAddressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServerNetworkInterfaceAdditionalIpAddressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

