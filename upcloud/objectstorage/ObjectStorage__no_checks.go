// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package objectstorage

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObjectStorage) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateImportFromParameters(id *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateMoveToIdParameters(id *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (o *jsiiProxy_ObjectStorage) validatePutBucketParameters(value interface{}) error {
	return nil
}

func validateObjectStorage_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateObjectStorage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateObjectStorage_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateObjectStorage_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ObjectStorage) validateSetAccessKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObjectStorage) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObjectStorage) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ObjectStorage) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObjectStorage) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObjectStorage) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ObjectStorage) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObjectStorage) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_ObjectStorage) validateSetSecretKeyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ObjectStorage) validateSetSizeParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ObjectStorage) validateSetZoneParameters(val *string) error {
	return nil
}

func validateNewObjectStorageParameters(scope constructs.Construct, id *string, config *ObjectStorageConfig) error {
	return nil
}

