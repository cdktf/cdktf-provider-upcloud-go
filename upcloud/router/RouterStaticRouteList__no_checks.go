// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package router

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouterStaticRouteList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RouterStaticRouteList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RouterStaticRouteList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RouterStaticRouteList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RouterStaticRouteList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RouterStaticRouteList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RouterStaticRouteList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRouterStaticRouteListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

