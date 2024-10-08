// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerbackend

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.loadbalancerBackend.LoadbalancerBackend",
		reflect.TypeOf((*LoadbalancerBackend)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loadbalancer", GoGetter: "Loadbalancer"},
			_jsii_.MemberProperty{JsiiProperty: "loadbalancerInput", GoGetter: "LoadbalancerInput"},
			_jsii_.MemberProperty{JsiiProperty: "members", GoGetter: "Members"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putProperties", GoMethod: "PutProperties"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetResolverName", GoMethod: "ResetResolverName"},
			_jsii_.MemberProperty{JsiiProperty: "resolverName", GoGetter: "ResolverName"},
			_jsii_.MemberProperty{JsiiProperty: "resolverNameInput", GoGetter: "ResolverNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsConfigs", GoGetter: "TlsConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_LoadbalancerBackend{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.loadbalancerBackend.LoadbalancerBackendConfig",
		reflect.TypeOf((*LoadbalancerBackendConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.loadbalancerBackend.LoadbalancerBackendProperties",
		reflect.TypeOf((*LoadbalancerBackendProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.loadbalancerBackend.LoadbalancerBackendPropertiesList",
		reflect.TypeOf((*LoadbalancerBackendPropertiesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_LoadbalancerBackendPropertiesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.loadbalancerBackend.LoadbalancerBackendPropertiesOutputReference",
		reflect.TypeOf((*LoadbalancerBackendPropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckExpectedStatus", GoGetter: "HealthCheckExpectedStatus"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckExpectedStatusInput", GoGetter: "HealthCheckExpectedStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckFall", GoGetter: "HealthCheckFall"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckFallInput", GoGetter: "HealthCheckFallInput"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckInterval", GoGetter: "HealthCheckInterval"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckIntervalInput", GoGetter: "HealthCheckIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckRise", GoGetter: "HealthCheckRise"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckRiseInput", GoGetter: "HealthCheckRiseInput"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckTlsVerify", GoGetter: "HealthCheckTlsVerify"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckTlsVerifyInput", GoGetter: "HealthCheckTlsVerifyInput"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckType", GoGetter: "HealthCheckType"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckTypeInput", GoGetter: "HealthCheckTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckUrl", GoGetter: "HealthCheckUrl"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckUrlInput", GoGetter: "HealthCheckUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "http2Enabled", GoGetter: "Http2Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "http2EnabledInput", GoGetter: "Http2EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "outboundProxyProtocol", GoGetter: "OutboundProxyProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "outboundProxyProtocolInput", GoGetter: "OutboundProxyProtocolInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckExpectedStatus", GoMethod: "ResetHealthCheckExpectedStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckFall", GoMethod: "ResetHealthCheckFall"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckInterval", GoMethod: "ResetHealthCheckInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckRise", GoMethod: "ResetHealthCheckRise"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckTlsVerify", GoMethod: "ResetHealthCheckTlsVerify"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckType", GoMethod: "ResetHealthCheckType"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckUrl", GoMethod: "ResetHealthCheckUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttp2Enabled", GoMethod: "ResetHttp2Enabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutboundProxyProtocol", GoMethod: "ResetOutboundProxyProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetStickySessionCookieName", GoMethod: "ResetStickySessionCookieName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeoutServer", GoMethod: "ResetTimeoutServer"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeoutTunnel", GoMethod: "ResetTimeoutTunnel"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsEnabled", GoMethod: "ResetTlsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsUseSystemCa", GoMethod: "ResetTlsUseSystemCa"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsVerify", GoMethod: "ResetTlsVerify"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "stickySessionCookieName", GoGetter: "StickySessionCookieName"},
			_jsii_.MemberProperty{JsiiProperty: "stickySessionCookieNameInput", GoGetter: "StickySessionCookieNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutServer", GoGetter: "TimeoutServer"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutServerInput", GoGetter: "TimeoutServerInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutTunnel", GoGetter: "TimeoutTunnel"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutTunnelInput", GoGetter: "TimeoutTunnelInput"},
			_jsii_.MemberProperty{JsiiProperty: "tlsEnabled", GoGetter: "TlsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "tlsEnabledInput", GoGetter: "TlsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "tlsUseSystemCa", GoGetter: "TlsUseSystemCa"},
			_jsii_.MemberProperty{JsiiProperty: "tlsUseSystemCaInput", GoGetter: "TlsUseSystemCaInput"},
			_jsii_.MemberProperty{JsiiProperty: "tlsVerify", GoGetter: "TlsVerify"},
			_jsii_.MemberProperty{JsiiProperty: "tlsVerifyInput", GoGetter: "TlsVerifyInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LoadbalancerBackendPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
