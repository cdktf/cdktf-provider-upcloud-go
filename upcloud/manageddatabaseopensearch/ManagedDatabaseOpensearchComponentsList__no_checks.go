// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package manageddatabaseopensearch

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedDatabaseOpensearchComponentsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedDatabaseOpensearchComponentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseOpensearchComponentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseOpensearchComponentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseOpensearchComponentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedDatabaseOpensearchComponentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

