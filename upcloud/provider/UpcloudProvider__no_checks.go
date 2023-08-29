// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UpcloudProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (u *jsiiProxy_UpcloudProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateUpcloudProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUpcloudProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateUpcloudProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewUpcloudProviderParameters(scope constructs.Construct, id *string, config *UpcloudProviderConfig) error {
	return nil
}

