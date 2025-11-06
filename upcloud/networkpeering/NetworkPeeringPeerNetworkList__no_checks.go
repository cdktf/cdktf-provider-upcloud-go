// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkpeering

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkPeeringPeerNetworkList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPeeringPeerNetworkList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkPeeringPeerNetworkList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkPeeringPeerNetworkList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkPeeringPeerNetworkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkPeeringPeerNetworkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkPeeringPeerNetworkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkPeeringPeerNetworkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

