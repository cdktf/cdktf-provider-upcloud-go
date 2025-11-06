// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kubernetesnodegroup

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesNodeGroupTaintList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KubernetesNodeGroupTaintList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KubernetesNodeGroupTaintList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodeGroupTaintList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodeGroupTaintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodeGroupTaintList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodeGroupTaintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKubernetesNodeGroupTaintListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

