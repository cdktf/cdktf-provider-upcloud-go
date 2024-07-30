// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package router

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouterStaticRoutesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RouterStaticRoutesList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RouterStaticRoutesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RouterStaticRoutesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RouterStaticRoutesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RouterStaticRoutesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRouterStaticRoutesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

