//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObjectStorageBucketList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_ObjectStorageBucketList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ObjectStorageBucketList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObjectStorageBucketList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObjectStorageBucketList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ObjectStorageBucketList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewObjectStorageBucketListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

