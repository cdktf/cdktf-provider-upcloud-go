//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataUpcloudHostsHostsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataUpcloudHostsHostsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudHostsHostsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudHostsHostsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudHostsHostsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataUpcloudHostsHostsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

