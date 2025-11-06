// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataupcloudhosts

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataUpcloudHostsHostsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataUpcloudHostsHostsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataUpcloudHostsHostsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudHostsHostsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudHostsHostsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudHostsHostsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataUpcloudHostsHostsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataUpcloudHostsHostsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

