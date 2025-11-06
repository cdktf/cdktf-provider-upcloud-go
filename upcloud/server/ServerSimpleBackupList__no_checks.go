// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package server

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServerSimpleBackupList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServerSimpleBackupList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServerSimpleBackupList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServerSimpleBackupList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServerSimpleBackupList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServerSimpleBackupList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServerSimpleBackupList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServerSimpleBackupListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

