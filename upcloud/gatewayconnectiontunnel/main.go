// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gatewayconnectiontunnel

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.gatewayConnectionTunnel.GatewayConnectionTunnel",
		reflect.TypeOf((*GatewayConnectionTunnel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "connectionId", GoGetter: "ConnectionId"},
			_jsii_.MemberProperty{JsiiProperty: "connectionIdInput", GoGetter: "ConnectionIdInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipsecAuthPsk", GoGetter: "IpsecAuthPsk"},
			_jsii_.MemberProperty{JsiiProperty: "ipsecAuthPskInput", GoGetter: "IpsecAuthPskInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipsecProperties", GoGetter: "IpsecProperties"},
			_jsii_.MemberProperty{JsiiProperty: "ipsecPropertiesInput", GoGetter: "IpsecPropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "localAddressName", GoGetter: "LocalAddressName"},
			_jsii_.MemberProperty{JsiiProperty: "localAddressNameInput", GoGetter: "LocalAddressNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "operationalState", GoGetter: "OperationalState"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putIpsecAuthPsk", GoMethod: "PutIpsecAuthPsk"},
			_jsii_.MemberMethod{JsiiMethod: "putIpsecProperties", GoMethod: "PutIpsecProperties"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "remoteAddress", GoGetter: "RemoteAddress"},
			_jsii_.MemberProperty{JsiiProperty: "remoteAddressInput", GoGetter: "RemoteAddressInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpsecProperties", GoMethod: "ResetIpsecProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_GatewayConnectionTunnel{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.gatewayConnectionTunnel.GatewayConnectionTunnelConfig",
		reflect.TypeOf((*GatewayConnectionTunnelConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.gatewayConnectionTunnel.GatewayConnectionTunnelIpsecAuthPsk",
		reflect.TypeOf((*GatewayConnectionTunnelIpsecAuthPsk)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.gatewayConnectionTunnel.GatewayConnectionTunnelIpsecAuthPskOutputReference",
		reflect.TypeOf((*GatewayConnectionTunnelIpsecAuthPskOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "psk", GoGetter: "Psk"},
			_jsii_.MemberProperty{JsiiProperty: "pskInput", GoGetter: "PskInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GatewayConnectionTunnelIpsecAuthPskOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.gatewayConnectionTunnel.GatewayConnectionTunnelIpsecProperties",
		reflect.TypeOf((*GatewayConnectionTunnelIpsecProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.gatewayConnectionTunnel.GatewayConnectionTunnelIpsecPropertiesOutputReference",
		reflect.TypeOf((*GatewayConnectionTunnelIpsecPropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "childRekeyTime", GoGetter: "ChildRekeyTime"},
			_jsii_.MemberProperty{JsiiProperty: "childRekeyTimeInput", GoGetter: "ChildRekeyTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dpdDelay", GoGetter: "DpdDelay"},
			_jsii_.MemberProperty{JsiiProperty: "dpdDelayInput", GoGetter: "DpdDelayInput"},
			_jsii_.MemberProperty{JsiiProperty: "dpdTimeout", GoGetter: "DpdTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "dpdTimeoutInput", GoGetter: "DpdTimeoutInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "ikeLifetime", GoGetter: "IkeLifetime"},
			_jsii_.MemberProperty{JsiiProperty: "ikeLifetimeInput", GoGetter: "IkeLifetimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "phase1Algorithms", GoGetter: "Phase1Algorithms"},
			_jsii_.MemberProperty{JsiiProperty: "phase1AlgorithmsInput", GoGetter: "Phase1AlgorithmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "phase1DhGroupNumbers", GoGetter: "Phase1DhGroupNumbers"},
			_jsii_.MemberProperty{JsiiProperty: "phase1DhGroupNumbersInput", GoGetter: "Phase1DhGroupNumbersInput"},
			_jsii_.MemberProperty{JsiiProperty: "phase1IntegrityAlgorithms", GoGetter: "Phase1IntegrityAlgorithms"},
			_jsii_.MemberProperty{JsiiProperty: "phase1IntegrityAlgorithmsInput", GoGetter: "Phase1IntegrityAlgorithmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "phase2Algorithms", GoGetter: "Phase2Algorithms"},
			_jsii_.MemberProperty{JsiiProperty: "phase2AlgorithmsInput", GoGetter: "Phase2AlgorithmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "phase2DhGroupNumbers", GoGetter: "Phase2DhGroupNumbers"},
			_jsii_.MemberProperty{JsiiProperty: "phase2DhGroupNumbersInput", GoGetter: "Phase2DhGroupNumbersInput"},
			_jsii_.MemberProperty{JsiiProperty: "phase2IntegrityAlgorithms", GoGetter: "Phase2IntegrityAlgorithms"},
			_jsii_.MemberProperty{JsiiProperty: "phase2IntegrityAlgorithmsInput", GoGetter: "Phase2IntegrityAlgorithmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "rekeyTime", GoGetter: "RekeyTime"},
			_jsii_.MemberProperty{JsiiProperty: "rekeyTimeInput", GoGetter: "RekeyTimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetChildRekeyTime", GoMethod: "ResetChildRekeyTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetDpdDelay", GoMethod: "ResetDpdDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetDpdTimeout", GoMethod: "ResetDpdTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetIkeLifetime", GoMethod: "ResetIkeLifetime"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhase1Algorithms", GoMethod: "ResetPhase1Algorithms"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhase1DhGroupNumbers", GoMethod: "ResetPhase1DhGroupNumbers"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhase1IntegrityAlgorithms", GoMethod: "ResetPhase1IntegrityAlgorithms"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhase2Algorithms", GoMethod: "ResetPhase2Algorithms"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhase2DhGroupNumbers", GoMethod: "ResetPhase2DhGroupNumbers"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhase2IntegrityAlgorithms", GoMethod: "ResetPhase2IntegrityAlgorithms"},
			_jsii_.MemberMethod{JsiiMethod: "resetRekeyTime", GoMethod: "ResetRekeyTime"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GatewayConnectionTunnelIpsecPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
