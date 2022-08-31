//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataUpcloudNetworksNetworksIpNetworkList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataUpcloudNetworksNetworksIpNetworkList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudNetworksNetworksIpNetworkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudNetworksNetworksIpNetworkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudNetworksNetworksIpNetworkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataUpcloudNetworksNetworksIpNetworkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

