// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storage

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageImportList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StorageImportList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageImportList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageImportList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageImportList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageImportList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageImportList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageImportListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

