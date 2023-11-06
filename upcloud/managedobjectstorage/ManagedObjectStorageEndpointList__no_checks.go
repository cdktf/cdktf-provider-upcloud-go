// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package managedobjectstorage

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedObjectStorageEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorageEndpointList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorageEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorageEndpointList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorageEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManagedObjectStorageEndpointListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

