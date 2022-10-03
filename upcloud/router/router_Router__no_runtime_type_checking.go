//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package router

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Router) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_Router) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Router) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Router) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Router) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Router) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Router) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Router) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Router) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Router) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Router) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Router) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateRouter_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Router) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Router) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Router) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Router) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Router) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewRouterParameters(scope constructs.Construct, id *string, config *RouterConfig) error {
	return nil
}

