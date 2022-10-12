//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package server

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Server) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Server) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Server) validatePutLoginParameters(value *ServerLogin) error {
	return nil
}

func (s *jsiiProxy_Server) validatePutNetworkInterfaceParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Server) validatePutSimpleBackupParameters(value *ServerSimpleBackup) error {
	return nil
}

func (s *jsiiProxy_Server) validatePutStorageDevicesParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_Server) validatePutTemplateParameters(value *ServerTemplate) error {
	return nil
}

func validateServer_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetCpuParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetFirewallParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetHostParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetHostnameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetLabelsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetMemParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetMetadataParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetPlanParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetTagsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetTitleParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetUserDataParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Server) validateSetZoneParameters(val *string) error {
	return nil
}

func validateNewServerParameters(scope constructs.Construct, id *string, config *ServerConfig) error {
	return nil
}

