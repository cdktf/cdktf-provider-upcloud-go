//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kubernetescluster

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesClusterNodeGroupTaintList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KubernetesClusterNodeGroupTaintList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterNodeGroupTaintList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterNodeGroupTaintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterNodeGroupTaintList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KubernetesClusterNodeGroupTaintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKubernetesClusterNodeGroupTaintListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

