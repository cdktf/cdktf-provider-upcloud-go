// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package manageddatabasevalkey

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedDatabaseValkeyNetworkList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ManagedDatabaseValkeyNetworkList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedDatabaseValkeyNetworkList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseValkeyNetworkList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseValkeyNetworkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseValkeyNetworkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedDatabaseValkeyNetworkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedDatabaseValkeyNetworkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

