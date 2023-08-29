// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerfrontendrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v9/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v9/loadbalancerfrontendrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerFrontendRuleMatchersOutputReference interface {
	cdktf.ComplexObject
	BodySize() LoadbalancerFrontendRuleMatchersBodySizeList
	BodySizeInput() interface{}
	BodySizeRange() LoadbalancerFrontendRuleMatchersBodySizeRangeList
	BodySizeRangeInput() interface{}
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
	Cookie() LoadbalancerFrontendRuleMatchersCookieList
	CookieInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Header() LoadbalancerFrontendRuleMatchersHeaderList
	HeaderInput() interface{}
	Host() LoadbalancerFrontendRuleMatchersHostList
	HostInput() interface{}
	HttpMethod() LoadbalancerFrontendRuleMatchersHttpMethodList
	HttpMethodInput() interface{}
	InternalValue() *LoadbalancerFrontendRuleMatchers
	SetInternalValue(val *LoadbalancerFrontendRuleMatchers)
	NumMembersUp() LoadbalancerFrontendRuleMatchersNumMembersUpList
	NumMembersUpInput() interface{}
	Path() LoadbalancerFrontendRuleMatchersPathList
	PathInput() interface{}
	SrcIp() LoadbalancerFrontendRuleMatchersSrcIpList
	SrcIpInput() interface{}
	SrcPort() LoadbalancerFrontendRuleMatchersSrcPortList
	SrcPortInput() interface{}
	SrcPortRange() LoadbalancerFrontendRuleMatchersSrcPortRangeList
	SrcPortRangeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() LoadbalancerFrontendRuleMatchersUrlList
	UrlInput() interface{}
	UrlParam() LoadbalancerFrontendRuleMatchersUrlParamList
	UrlParamInput() interface{}
	UrlQuery() LoadbalancerFrontendRuleMatchersUrlQueryList
	UrlQueryInput() interface{}
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
	PutBodySize(value interface{})
	PutBodySizeRange(value interface{})
	PutCookie(value interface{})
	PutHeader(value interface{})
	PutHost(value interface{})
	PutHttpMethod(value interface{})
	PutNumMembersUp(value interface{})
	PutPath(value interface{})
	PutSrcIp(value interface{})
	PutSrcPort(value interface{})
	PutSrcPortRange(value interface{})
	PutUrl(value interface{})
	PutUrlParam(value interface{})
	PutUrlQuery(value interface{})
	ResetBodySize()
	ResetBodySizeRange()
	ResetCookie()
	ResetHeader()
	ResetHost()
	ResetHttpMethod()
	ResetNumMembersUp()
	ResetPath()
	ResetSrcIp()
	ResetSrcPort()
	ResetSrcPortRange()
	ResetUrl()
	ResetUrlParam()
	ResetUrlQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LoadbalancerFrontendRuleMatchersOutputReference
type jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) BodySize() LoadbalancerFrontendRuleMatchersBodySizeList {
	var returns LoadbalancerFrontendRuleMatchersBodySizeList
	_jsii_.Get(
		j,
		"bodySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) BodySizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) BodySizeRange() LoadbalancerFrontendRuleMatchersBodySizeRangeList {
	var returns LoadbalancerFrontendRuleMatchersBodySizeRangeList
	_jsii_.Get(
		j,
		"bodySizeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) BodySizeRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodySizeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) Cookie() LoadbalancerFrontendRuleMatchersCookieList {
	var returns LoadbalancerFrontendRuleMatchersCookieList
	_jsii_.Get(
		j,
		"cookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) CookieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) Header() LoadbalancerFrontendRuleMatchersHeaderList {
	var returns LoadbalancerFrontendRuleMatchersHeaderList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) Host() LoadbalancerFrontendRuleMatchersHostList {
	var returns LoadbalancerFrontendRuleMatchersHostList
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) HostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) HttpMethod() LoadbalancerFrontendRuleMatchersHttpMethodList {
	var returns LoadbalancerFrontendRuleMatchersHttpMethodList
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) HttpMethodInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) InternalValue() *LoadbalancerFrontendRuleMatchers {
	var returns *LoadbalancerFrontendRuleMatchers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) NumMembersUp() LoadbalancerFrontendRuleMatchersNumMembersUpList {
	var returns LoadbalancerFrontendRuleMatchersNumMembersUpList
	_jsii_.Get(
		j,
		"numMembersUp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) NumMembersUpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numMembersUpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) Path() LoadbalancerFrontendRuleMatchersPathList {
	var returns LoadbalancerFrontendRuleMatchersPathList
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) SrcIp() LoadbalancerFrontendRuleMatchersSrcIpList {
	var returns LoadbalancerFrontendRuleMatchersSrcIpList
	_jsii_.Get(
		j,
		"srcIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) SrcIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srcIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) SrcPort() LoadbalancerFrontendRuleMatchersSrcPortList {
	var returns LoadbalancerFrontendRuleMatchersSrcPortList
	_jsii_.Get(
		j,
		"srcPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) SrcPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srcPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) SrcPortRange() LoadbalancerFrontendRuleMatchersSrcPortRangeList {
	var returns LoadbalancerFrontendRuleMatchersSrcPortRangeList
	_jsii_.Get(
		j,
		"srcPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) SrcPortRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srcPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) Url() LoadbalancerFrontendRuleMatchersUrlList {
	var returns LoadbalancerFrontendRuleMatchersUrlList
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) UrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) UrlParam() LoadbalancerFrontendRuleMatchersUrlParamList {
	var returns LoadbalancerFrontendRuleMatchersUrlParamList
	_jsii_.Get(
		j,
		"urlParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) UrlParamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) UrlQuery() LoadbalancerFrontendRuleMatchersUrlQueryList {
	var returns LoadbalancerFrontendRuleMatchersUrlQueryList
	_jsii_.Get(
		j,
		"urlQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) UrlQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlQueryInput",
		&returns,
	)
	return returns
}


func NewLoadbalancerFrontendRuleMatchersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LoadbalancerFrontendRuleMatchersOutputReference {
	_init_.Initialize()

	if err := validateNewLoadbalancerFrontendRuleMatchersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.loadbalancerFrontendRule.LoadbalancerFrontendRuleMatchersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLoadbalancerFrontendRuleMatchersOutputReference_Override(l LoadbalancerFrontendRuleMatchersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.loadbalancerFrontendRule.LoadbalancerFrontendRuleMatchersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference)SetInternalValue(val *LoadbalancerFrontendRuleMatchers) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutBodySize(value interface{}) {
	if err := l.validatePutBodySizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBodySize",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutBodySizeRange(value interface{}) {
	if err := l.validatePutBodySizeRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBodySizeRange",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutCookie(value interface{}) {
	if err := l.validatePutCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCookie",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutHeader(value interface{}) {
	if err := l.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHeader",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutHost(value interface{}) {
	if err := l.validatePutHostParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHost",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutHttpMethod(value interface{}) {
	if err := l.validatePutHttpMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHttpMethod",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutNumMembersUp(value interface{}) {
	if err := l.validatePutNumMembersUpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putNumMembersUp",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutPath(value interface{}) {
	if err := l.validatePutPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPath",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutSrcIp(value interface{}) {
	if err := l.validatePutSrcIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSrcIp",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutSrcPort(value interface{}) {
	if err := l.validatePutSrcPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSrcPort",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutSrcPortRange(value interface{}) {
	if err := l.validatePutSrcPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSrcPortRange",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutUrl(value interface{}) {
	if err := l.validatePutUrlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putUrl",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutUrlParam(value interface{}) {
	if err := l.validatePutUrlParamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putUrlParam",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) PutUrlQuery(value interface{}) {
	if err := l.validatePutUrlQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putUrlQuery",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetBodySize() {
	_jsii_.InvokeVoid(
		l,
		"resetBodySize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetBodySizeRange() {
	_jsii_.InvokeVoid(
		l,
		"resetBodySizeRange",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetCookie() {
	_jsii_.InvokeVoid(
		l,
		"resetCookie",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		l,
		"resetHeader",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		l,
		"resetHost",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		l,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetNumMembersUp() {
	_jsii_.InvokeVoid(
		l,
		"resetNumMembersUp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		l,
		"resetPath",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetSrcIp() {
	_jsii_.InvokeVoid(
		l,
		"resetSrcIp",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetSrcPort() {
	_jsii_.InvokeVoid(
		l,
		"resetSrcPort",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetSrcPortRange() {
	_jsii_.InvokeVoid(
		l,
		"resetSrcPortRange",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		l,
		"resetUrl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetUrlParam() {
	_jsii_.InvokeVoid(
		l,
		"resetUrlParam",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ResetUrlQuery() {
	_jsii_.InvokeVoid(
		l,
		"resetUrlQuery",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LoadbalancerFrontendRuleMatchersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

