// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v12/server/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerTemplateOutputReference interface {
	cdktf.ComplexObject
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
	AddressPosition() *string
	SetAddressPosition(val *string)
	AddressPositionInput() *string
	BackupRule() ServerTemplateBackupRuleOutputReference
	BackupRuleInput() *ServerTemplateBackupRule
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeleteAutoresizeBackup() interface{}
	SetDeleteAutoresizeBackup(val interface{})
	DeleteAutoresizeBackupInput() interface{}
	FilesystemAutoresize() interface{}
	SetFilesystemAutoresize(val interface{})
	FilesystemAutoresizeInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *ServerTemplate
	SetInternalValue(val *ServerTemplate)
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	Storage() *string
	SetStorage(val *string)
	StorageInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tier() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutBackupRule(value *ServerTemplateBackupRule)
	ResetAddress()
	ResetAddressPosition()
	ResetBackupRule()
	ResetDeleteAutoresizeBackup()
	ResetFilesystemAutoresize()
	ResetSize()
	ResetTitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServerTemplateOutputReference
type jsiiProxy_ServerTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServerTemplateOutputReference) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) AddressPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) AddressPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) BackupRule() ServerTemplateBackupRuleOutputReference {
	var returns ServerTemplateBackupRuleOutputReference
	_jsii_.Get(
		j,
		"backupRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) BackupRuleInput() *ServerTemplateBackupRule {
	var returns *ServerTemplateBackupRule
	_jsii_.Get(
		j,
		"backupRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) DeleteAutoresizeBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutoresizeBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) DeleteAutoresizeBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutoresizeBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) FilesystemAutoresize() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filesystemAutoresize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) FilesystemAutoresizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filesystemAutoresizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) InternalValue() *ServerTemplate {
	var returns *ServerTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) Storage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) StorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServerTemplateOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}


func NewServerTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServerTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewServerTemplateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServerTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.server.ServerTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServerTemplateOutputReference_Override(s ServerTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.server.ServerTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetAddress(val *string) {
	if err := j.validateSetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetAddressPosition(val *string) {
	if err := j.validateSetAddressPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressPosition",
		val,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetDeleteAutoresizeBackup(val interface{}) {
	if err := j.validateSetDeleteAutoresizeBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAutoresizeBackup",
		val,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetFilesystemAutoresize(val interface{}) {
	if err := j.validateSetFilesystemAutoresizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filesystemAutoresize",
		val,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetInternalValue(val *ServerTemplate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetStorage(val *string) {
	if err := j.validateSetStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storage",
		val,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServerTemplateOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (s *jsiiProxy_ServerTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) PutBackupRule(value *ServerTemplateBackupRule) {
	if err := s.validatePutBackupRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBackupRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServerTemplateOutputReference) ResetAddress() {
	_jsii_.InvokeVoid(
		s,
		"resetAddress",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerTemplateOutputReference) ResetAddressPosition() {
	_jsii_.InvokeVoid(
		s,
		"resetAddressPosition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerTemplateOutputReference) ResetBackupRule() {
	_jsii_.InvokeVoid(
		s,
		"resetBackupRule",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerTemplateOutputReference) ResetDeleteAutoresizeBackup() {
	_jsii_.InvokeVoid(
		s,
		"resetDeleteAutoresizeBackup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerTemplateOutputReference) ResetFilesystemAutoresize() {
	_jsii_.InvokeVoid(
		s,
		"resetFilesystemAutoresize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerTemplateOutputReference) ResetSize() {
	_jsii_.InvokeVoid(
		s,
		"resetSize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerTemplateOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		s,
		"resetTitle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServerTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServerTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

