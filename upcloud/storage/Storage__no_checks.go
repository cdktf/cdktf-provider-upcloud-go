// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storage

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Storage) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Storage) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Storage) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Storage) validatePutBackupRuleParameters(value *StorageBackupRule) error {
	return nil
}

func (s *jsiiProxy_Storage) validatePutCloneParameters(value *StorageClone) error {
	return nil
}

func (s *jsiiProxy_Storage) validatePutImportParameters(value *StorageImport) error {
	return nil
}

func validateStorage_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateStorage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStorage_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateStorage_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Storage) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Storage) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Storage) validateSetDeleteAutoresizeBackupParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Storage) validateSetFilesystemAutoresizeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Storage) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Storage) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Storage) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Storage) validateSetSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Storage) validateSetTierParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Storage) validateSetTitleParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Storage) validateSetZoneParameters(val *string) error {
	return nil
}

func validateNewStorageParameters(scope constructs.Construct, id *string, config *StorageConfig) error {
	return nil
}

