// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package managedobjectstorage

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManagedObjectStorage) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateImportFromParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (m *jsiiProxy_ManagedObjectStorage) validatePutNetworkParameters(value interface{}) error {
	return nil
}

func validateManagedObjectStorage_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateManagedObjectStorage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateManagedObjectStorage_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateManagedObjectStorage_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorage) validateSetConfiguredStatusParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorage) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorage) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorage) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorage) validateSetLabelsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorage) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorage) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorage) validateSetRegionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManagedObjectStorage) validateSetUsersParameters(val *[]*string) error {
	return nil
}

func validateNewManagedObjectStorageParameters(scope constructs.Construct, id *string, config *ManagedObjectStorageConfig) error {
	return nil
}

