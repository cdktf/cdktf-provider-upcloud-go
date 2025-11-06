// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package filestorage

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FileStorageShareAclList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (f *jsiiProxy_FileStorageShareAclList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FileStorageShareAclList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FileStorageShareAclList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FileStorageShareAclList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FileStorageShareAclList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FileStorageShareAclList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFileStorageShareAclListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

