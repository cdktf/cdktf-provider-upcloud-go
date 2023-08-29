// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package manageddatabasepostgresql

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedDatabasePostgresqlNodeStatesList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedDatabasePostgresqlNodeStatesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabasePostgresqlNodeStatesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabasePostgresqlNodeStatesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabasePostgresqlNodeStatesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedDatabasePostgresqlNodeStatesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

