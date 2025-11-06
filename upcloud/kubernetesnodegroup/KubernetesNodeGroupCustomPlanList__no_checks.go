// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package kubernetesnodegroup

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesNodeGroupCustomPlanList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (k *jsiiProxy_KubernetesNodeGroupCustomPlanList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KubernetesNodeGroupCustomPlanList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodeGroupCustomPlanList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodeGroupCustomPlanList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodeGroupCustomPlanList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KubernetesNodeGroupCustomPlanList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKubernetesNodeGroupCustomPlanListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

