//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedDatabasePostgresqlComponentsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedDatabasePostgresqlComponentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabasePostgresqlComponentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabasePostgresqlComponentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabasePostgresqlComponentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedDatabasePostgresqlComponentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

