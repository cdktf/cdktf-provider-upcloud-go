// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package managedobjectstorage

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedObjectStorageNetworkList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorageNetworkList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorageNetworkList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorageNetworkList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorageNetworkList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorageNetworkList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorageNetworkList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedObjectStorageNetworkListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

