// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-upcloud-go/upcloud/v2/jsii"

	"github.com/hashicorp/cdktf-provider-upcloud-go/upcloud/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerBackendPropertiesOutputReference interface {
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
	HealthCheckExpectedStatus() *float64
	SetHealthCheckExpectedStatus(val *float64)
	HealthCheckExpectedStatusInput() *float64
	HealthCheckFall() *float64
	SetHealthCheckFall(val *float64)
	HealthCheckFallInput() *float64
	HealthCheckInterval() *float64
	SetHealthCheckInterval(val *float64)
	HealthCheckIntervalInput() *float64
	HealthCheckRise() *float64
	SetHealthCheckRise(val *float64)
	HealthCheckRiseInput() *float64
	HealthCheckType() *string
	SetHealthCheckType(val *string)
	HealthCheckTypeInput() *string
	HealthCheckUrl() *string
	SetHealthCheckUrl(val *string)
	HealthCheckUrlInput() *string
	InternalValue() *LoadbalancerBackendProperties
	SetInternalValue(val *LoadbalancerBackendProperties)
	OutboundProxyProtocol() *string
	SetOutboundProxyProtocol(val *string)
	OutboundProxyProtocolInput() *string
	StickySessionCookieName() *string
	SetStickySessionCookieName(val *string)
	StickySessionCookieNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutServer() *float64
	SetTimeoutServer(val *float64)
	TimeoutServerInput() *float64
	TimeoutTunnel() *float64
	SetTimeoutTunnel(val *float64)
	TimeoutTunnelInput() *float64
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
	ResetHealthCheckExpectedStatus()
	ResetHealthCheckFall()
	ResetHealthCheckInterval()
	ResetHealthCheckRise()
	ResetHealthCheckType()
	ResetHealthCheckUrl()
	ResetOutboundProxyProtocol()
	ResetStickySessionCookieName()
	ResetTimeoutServer()
	ResetTimeoutTunnel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LoadbalancerBackendPropertiesOutputReference
type jsiiProxy_LoadbalancerBackendPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckExpectedStatus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckExpectedStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckExpectedStatusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckExpectedStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckFall() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckFall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckFallInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckFallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckRise() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckRise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckRiseInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckRiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) HealthCheckUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) InternalValue() *LoadbalancerBackendProperties {
	var returns *LoadbalancerBackendProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) OutboundProxyProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundProxyProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) OutboundProxyProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundProxyProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) StickySessionCookieName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stickySessionCookieName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) StickySessionCookieNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stickySessionCookieNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) TimeoutServer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) TimeoutServerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) TimeoutTunnel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutTunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) TimeoutTunnelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutTunnelInput",
		&returns,
	)
	return returns
}


func NewLoadbalancerBackendPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LoadbalancerBackendPropertiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_LoadbalancerBackendPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.LoadbalancerBackendPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLoadbalancerBackendPropertiesOutputReference_Override(l LoadbalancerBackendPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.LoadbalancerBackendPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetHealthCheckExpectedStatus(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckExpectedStatus",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetHealthCheckFall(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckFall",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetHealthCheckInterval(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckInterval",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetHealthCheckRise(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckRise",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetHealthCheckType(val *string) {
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetHealthCheckUrl(val *string) {
	_jsii_.Set(
		j,
		"healthCheckUrl",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetInternalValue(val *LoadbalancerBackendProperties) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetOutboundProxyProtocol(val *string) {
	_jsii_.Set(
		j,
		"outboundProxyProtocol",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetStickySessionCookieName(val *string) {
	_jsii_.Set(
		j,
		"stickySessionCookieName",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetTimeoutServer(val *float64) {
	_jsii_.Set(
		j,
		"timeoutServer",
		val,
	)
}

func (j *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) SetTimeoutTunnel(val *float64) {
	_jsii_.Set(
		j,
		"timeoutTunnel",
		val,
	)
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ResetHealthCheckExpectedStatus() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheckExpectedStatus",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ResetHealthCheckFall() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheckFall",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ResetHealthCheckInterval() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheckInterval",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ResetHealthCheckRise() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheckRise",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ResetHealthCheckType() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheckType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ResetHealthCheckUrl() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheckUrl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ResetOutboundProxyProtocol() {
	_jsii_.InvokeVoid(
		l,
		"resetOutboundProxyProtocol",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ResetStickySessionCookieName() {
	_jsii_.InvokeVoid(
		l,
		"resetStickySessionCookieName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ResetTimeoutServer() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeoutServer",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ResetTimeoutTunnel() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeoutTunnel",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoadbalancerBackendPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

