// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storagetemplate

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageTemplate) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_StorageTemplate) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateStorageTemplate_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateStorageTemplate_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStorageTemplate_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateStorageTemplate_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageTemplate) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageTemplate) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageTemplate) validateSetLabelsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_StorageTemplate) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_StorageTemplate) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageTemplate) validateSetSourceStorageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageTemplate) validateSetTitleParameters(val *string) error {
	return nil
}

func validateNewStorageTemplateParameters(scope constructs.Construct, id *string, config *StorageTemplateConfig) error {
	return nil
}

