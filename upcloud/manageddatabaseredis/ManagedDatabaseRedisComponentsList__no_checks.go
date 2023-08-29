// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package manageddatabaseredis

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedDatabaseRedisComponentsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedDatabaseRedisComponentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseRedisComponentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseRedisComponentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseRedisComponentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedDatabaseRedisComponentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

