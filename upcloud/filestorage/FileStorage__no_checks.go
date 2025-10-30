// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package filestorage

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FileStorage) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateImportFromParameters(id *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateMoveToIdParameters(id *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validatePutNetworkParameters(value interface{}) error {
	return nil
}

func (f *jsiiProxy_FileStorage) validatePutShareParameters(value interface{}) error {
	return nil
}

func validateFileStorage_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateFileStorage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFileStorage_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateFileStorage_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_FileStorage) validateSetConfiguredStatusParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FileStorage) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FileStorage) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FileStorage) validateSetLabelsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_FileStorage) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_FileStorage) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FileStorage) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_FileStorage) validateSetSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_FileStorage) validateSetZoneParameters(val *string) error {
	return nil
}

func validateNewFileStorageParameters(scope constructs.Construct, id *string, config *FileStorageConfig) error {
	return nil
}

