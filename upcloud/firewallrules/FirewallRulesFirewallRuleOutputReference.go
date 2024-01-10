// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v12/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v12/firewallrules/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallRulesFirewallRuleOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
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
	DestinationAddressEnd() *string
	SetDestinationAddressEnd(val *string)
	DestinationAddressEndInput() *string
	DestinationAddressStart() *string
	SetDestinationAddressStart(val *string)
	DestinationAddressStartInput() *string
	DestinationPortEnd() *string
	SetDestinationPortEnd(val *string)
	DestinationPortEndInput() *string
	DestinationPortStart() *string
	SetDestinationPortStart(val *string)
	DestinationPortStartInput() *string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	Family() *string
	SetFamily(val *string)
	FamilyInput() *string
	// Experimental.
	Fqn() *string
	IcmpType() *string
	SetIcmpType(val *string)
	IcmpTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	SourceAddressEnd() *string
	SetSourceAddressEnd(val *string)
	SourceAddressEndInput() *string
	SourceAddressStart() *string
	SetSourceAddressStart(val *string)
	SourceAddressStartInput() *string
	SourcePortEnd() *string
	SetSourcePortEnd(val *string)
	SourcePortEndInput() *string
	SourcePortStart() *string
	SetSourcePortStart(val *string)
	SourcePortStartInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetComment()
	ResetDestinationAddressEnd()
	ResetDestinationAddressStart()
	ResetDestinationPortEnd()
	ResetDestinationPortStart()
	ResetFamily()
	ResetIcmpType()
	ResetProtocol()
	ResetSourceAddressEnd()
	ResetSourceAddressStart()
	ResetSourcePortEnd()
	ResetSourcePortStart()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FirewallRulesFirewallRuleOutputReference
type jsiiProxy_FirewallRulesFirewallRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) DestinationAddressEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) DestinationAddressEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) DestinationAddressStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) DestinationAddressStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationAddressStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) DestinationPortEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) DestinationPortEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) DestinationPortStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) DestinationPortStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) FamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) IcmpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) IcmpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"icmpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) SourceAddressEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) SourceAddressEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) SourceAddressStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) SourceAddressStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAddressStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) SourcePortEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) SourcePortEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) SourcePortStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) SourcePortStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFirewallRulesFirewallRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FirewallRulesFirewallRuleOutputReference {
	_init_.Initialize()

	if err := validateNewFirewallRulesFirewallRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirewallRulesFirewallRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.firewallRules.FirewallRulesFirewallRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFirewallRulesFirewallRuleOutputReference_Override(f FirewallRulesFirewallRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.firewallRules.FirewallRulesFirewallRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetComment(val *string) {
	if err := j.validateSetCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetDestinationAddressEnd(val *string) {
	if err := j.validateSetDestinationAddressEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddressEnd",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetDestinationAddressStart(val *string) {
	if err := j.validateSetDestinationAddressStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationAddressStart",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetDestinationPortEnd(val *string) {
	if err := j.validateSetDestinationPortEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPortEnd",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetDestinationPortStart(val *string) {
	if err := j.validateSetDestinationPortStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPortStart",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetFamily(val *string) {
	if err := j.validateSetFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetIcmpType(val *string) {
	if err := j.validateSetIcmpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"icmpType",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetSourceAddressEnd(val *string) {
	if err := j.validateSetSourceAddressEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddressEnd",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetSourceAddressStart(val *string) {
	if err := j.validateSetSourceAddressStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAddressStart",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetSourcePortEnd(val *string) {
	if err := j.validateSetSourcePortEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePortEnd",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetSourcePortStart(val *string) {
	if err := j.validateSetSourcePortStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePortStart",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FirewallRulesFirewallRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetComment() {
	_jsii_.InvokeVoid(
		f,
		"resetComment",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetDestinationAddressEnd() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationAddressEnd",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetDestinationAddressStart() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationAddressStart",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetDestinationPortEnd() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationPortEnd",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetDestinationPortStart() {
	_jsii_.InvokeVoid(
		f,
		"resetDestinationPortStart",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetFamily() {
	_jsii_.InvokeVoid(
		f,
		"resetFamily",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetIcmpType() {
	_jsii_.InvokeVoid(
		f,
		"resetIcmpType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		f,
		"resetProtocol",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetSourceAddressEnd() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceAddressEnd",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetSourceAddressStart() {
	_jsii_.InvokeVoid(
		f,
		"resetSourceAddressStart",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetSourcePortEnd() {
	_jsii_.InvokeVoid(
		f,
		"resetSourcePortEnd",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ResetSourcePortStart() {
	_jsii_.InvokeVoid(
		f,
		"resetSourcePortStart",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirewallRulesFirewallRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

