// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package manageddatabasemysql

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedDatabaseMysqlComponentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ManagedDatabaseMysqlComponentsList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedDatabaseMysqlComponentsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseMysqlComponentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseMysqlComponentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseMysqlComponentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedDatabaseMysqlComponentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

