//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataUpcloudNetworksNetworksList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataUpcloudNetworksNetworksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudNetworksNetworksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudNetworksNetworksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudNetworksNetworksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataUpcloudNetworksNetworksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

