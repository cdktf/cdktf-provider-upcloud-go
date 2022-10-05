package loadbalancerfrontendrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v3/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v3/loadbalancerfrontendrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerFrontendRuleActionsOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	HttpRedirect() LoadbalancerFrontendRuleActionsHttpRedirectList
	HttpRedirectInput() interface{}
	HttpReturn() LoadbalancerFrontendRuleActionsHttpReturnList
	HttpReturnInput() interface{}
	InternalValue() *LoadbalancerFrontendRuleActions
	SetInternalValue(val *LoadbalancerFrontendRuleActions)
	SetForwardedHeaders() LoadbalancerFrontendRuleActionsSetForwardedHeadersList
	SetForwardedHeadersInput() interface{}
	TcpReject() LoadbalancerFrontendRuleActionsTcpRejectList
	TcpRejectInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseBackend() LoadbalancerFrontendRuleActionsUseBackendList
	UseBackendInput() interface{}
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
	PutHttpRedirect(value interface{})
	PutHttpReturn(value interface{})
	PutSetForwardedHeaders(value interface{})
	PutTcpReject(value interface{})
	PutUseBackend(value interface{})
	ResetHttpRedirect()
	ResetHttpReturn()
	ResetSetForwardedHeaders()
	ResetTcpReject()
	ResetUseBackend()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LoadbalancerFrontendRuleActionsOutputReference
type jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) HttpRedirect() LoadbalancerFrontendRuleActionsHttpRedirectList {
	var returns LoadbalancerFrontendRuleActionsHttpRedirectList
	_jsii_.Get(
		j,
		"httpRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) HttpRedirectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) HttpReturn() LoadbalancerFrontendRuleActionsHttpReturnList {
	var returns LoadbalancerFrontendRuleActionsHttpReturnList
	_jsii_.Get(
		j,
		"httpReturn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) HttpReturnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpReturnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) InternalValue() *LoadbalancerFrontendRuleActions {
	var returns *LoadbalancerFrontendRuleActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) SetForwardedHeaders() LoadbalancerFrontendRuleActionsSetForwardedHeadersList {
	var returns LoadbalancerFrontendRuleActionsSetForwardedHeadersList
	_jsii_.Get(
		j,
		"setForwardedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) SetForwardedHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"setForwardedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) TcpReject() LoadbalancerFrontendRuleActionsTcpRejectList {
	var returns LoadbalancerFrontendRuleActionsTcpRejectList
	_jsii_.Get(
		j,
		"tcpReject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) TcpRejectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcpRejectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) UseBackend() LoadbalancerFrontendRuleActionsUseBackendList {
	var returns LoadbalancerFrontendRuleActionsUseBackendList
	_jsii_.Get(
		j,
		"useBackend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) UseBackendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBackendInput",
		&returns,
	)
	return returns
}


func NewLoadbalancerFrontendRuleActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LoadbalancerFrontendRuleActionsOutputReference {
	_init_.Initialize()

	if err := validateNewLoadbalancerFrontendRuleActionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.loadbalancerFrontendRule.LoadbalancerFrontendRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLoadbalancerFrontendRuleActionsOutputReference_Override(l LoadbalancerFrontendRuleActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.loadbalancerFrontendRule.LoadbalancerFrontendRuleActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference)SetInternalValue(val *LoadbalancerFrontendRuleActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) PutHttpRedirect(value interface{}) {
	if err := l.validatePutHttpRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHttpRedirect",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) PutHttpReturn(value interface{}) {
	if err := l.validatePutHttpReturnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHttpReturn",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) PutSetForwardedHeaders(value interface{}) {
	if err := l.validatePutSetForwardedHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSetForwardedHeaders",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) PutTcpReject(value interface{}) {
	if err := l.validatePutTcpRejectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTcpReject",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) PutUseBackend(value interface{}) {
	if err := l.validatePutUseBackendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putUseBackend",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) ResetHttpRedirect() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpRedirect",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) ResetHttpReturn() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpReturn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) ResetSetForwardedHeaders() {
	_jsii_.InvokeVoid(
		l,
		"resetSetForwardedHeaders",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) ResetTcpReject() {
	_jsii_.InvokeVoid(
		l,
		"resetTcpReject",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) ResetUseBackend() {
	_jsii_.InvokeVoid(
		l,
		"resetUseBackend",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

