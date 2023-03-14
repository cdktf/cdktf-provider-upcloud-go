//go:build no_runtime_type_checking

package gateway

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_Gateway) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (g *jsiiProxy_Gateway) validatePutRouterParameters(value *GatewayRouter) error {
	return nil
}

func validateGateway_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGateway_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGateway_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Gateway) validateSetConfiguredStatusParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Gateway) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Gateway) validateSetFeaturesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Gateway) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Gateway) validateSetLabelsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Gateway) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Gateway) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Gateway) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Gateway) validateSetZoneParameters(val *string) error {
	return nil
}

func validateNewGatewayParameters(scope constructs.Construct, id *string, config *GatewayConfig) error {
	return nil
}

