// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.provider.UpcloudProvider",
		reflect.TypeOf((*UpcloudProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeoutSec", GoGetter: "RequestTimeoutSec"},
			_jsii_.MemberProperty{JsiiProperty: "requestTimeoutSecInput", GoGetter: "RequestTimeoutSecInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestTimeoutSec", GoMethod: "ResetRequestTimeoutSec"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryMax", GoMethod: "ResetRetryMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryWaitMaxSec", GoMethod: "ResetRetryWaitMaxSec"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryWaitMinSec", GoMethod: "ResetRetryWaitMinSec"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsername", GoMethod: "ResetUsername"},
			_jsii_.MemberProperty{JsiiProperty: "retryMax", GoGetter: "RetryMax"},
			_jsii_.MemberProperty{JsiiProperty: "retryMaxInput", GoGetter: "RetryMaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "retryWaitMaxSec", GoGetter: "RetryWaitMaxSec"},
			_jsii_.MemberProperty{JsiiProperty: "retryWaitMaxSecInput", GoGetter: "RetryWaitMaxSecInput"},
			_jsii_.MemberProperty{JsiiProperty: "retryWaitMinSec", GoGetter: "RetryWaitMinSec"},
			_jsii_.MemberProperty{JsiiProperty: "retryWaitMinSecInput", GoGetter: "RetryWaitMinSecInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameInput", GoGetter: "UsernameInput"},
		},
		func() interface{} {
			j := jsiiProxy_UpcloudProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.provider.UpcloudProviderConfig",
		reflect.TypeOf((*UpcloudProviderConfig)(nil)).Elem(),
	)
}
