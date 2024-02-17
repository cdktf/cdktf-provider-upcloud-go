// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v13/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/upcloudltd/upcloud/4.1.0/docs upcloud}.
type UpcloudProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	RawOverrides() interface{}
	RequestTimeoutSec() *float64
	SetRequestTimeoutSec(val *float64)
	RequestTimeoutSecInput() *float64
	RetryMax() *float64
	SetRetryMax(val *float64)
	RetryMaxInput() *float64
	RetryWaitMaxSec() *float64
	SetRetryWaitMaxSec(val *float64)
	RetryWaitMaxSecInput() *float64
	RetryWaitMinSec() *float64
	SetRetryWaitMinSec(val *float64)
	RetryWaitMinSecInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetRequestTimeoutSec()
	ResetRetryMax()
	ResetRetryWaitMaxSec()
	ResetRetryWaitMinSec()
	ResetUsername()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for UpcloudProvider
type jsiiProxy_UpcloudProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_UpcloudProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) RequestTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) RequestTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) RetryMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) RetryMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) RetryWaitMaxSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryWaitMaxSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) RetryWaitMaxSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryWaitMaxSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) RetryWaitMinSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryWaitMinSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) RetryWaitMinSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryWaitMinSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpcloudProvider) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/upcloudltd/upcloud/4.1.0/docs upcloud} Resource.
func NewUpcloudProvider(scope constructs.Construct, id *string, config *UpcloudProviderConfig) UpcloudProvider {
	_init_.Initialize()

	if err := validateNewUpcloudProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_UpcloudProvider{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.provider.UpcloudProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/upcloudltd/upcloud/4.1.0/docs upcloud} Resource.
func NewUpcloudProvider_Override(u UpcloudProvider, scope constructs.Construct, id *string, config *UpcloudProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.provider.UpcloudProvider",
		[]interface{}{scope, id, config},
		u,
	)
}

func (j *jsiiProxy_UpcloudProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_UpcloudProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_UpcloudProvider)SetRequestTimeoutSec(val *float64) {
	_jsii_.Set(
		j,
		"requestTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_UpcloudProvider)SetRetryMax(val *float64) {
	_jsii_.Set(
		j,
		"retryMax",
		val,
	)
}

func (j *jsiiProxy_UpcloudProvider)SetRetryWaitMaxSec(val *float64) {
	_jsii_.Set(
		j,
		"retryWaitMaxSec",
		val,
	)
}

func (j *jsiiProxy_UpcloudProvider)SetRetryWaitMinSec(val *float64) {
	_jsii_.Set(
		j,
		"retryWaitMinSec",
		val,
	)
}

func (j *jsiiProxy_UpcloudProvider)SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a UpcloudProvider resource upon running "cdktf plan <stack-name>".
func UpcloudProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateUpcloudProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.provider.UpcloudProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func UpcloudProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUpcloudProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.provider.UpcloudProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func UpcloudProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUpcloudProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.provider.UpcloudProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func UpcloudProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUpcloudProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.provider.UpcloudProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func UpcloudProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-upcloud.provider.UpcloudProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_UpcloudProvider) AddOverride(path *string, value interface{}) {
	if err := u.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (u *jsiiProxy_UpcloudProvider) OverrideLogicalId(newLogicalId *string) {
	if err := u.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		u,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (u *jsiiProxy_UpcloudProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		u,
		"resetAlias",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpcloudProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		u,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpcloudProvider) ResetPassword() {
	_jsii_.InvokeVoid(
		u,
		"resetPassword",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpcloudProvider) ResetRequestTimeoutSec() {
	_jsii_.InvokeVoid(
		u,
		"resetRequestTimeoutSec",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpcloudProvider) ResetRetryMax() {
	_jsii_.InvokeVoid(
		u,
		"resetRetryMax",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpcloudProvider) ResetRetryWaitMaxSec() {
	_jsii_.InvokeVoid(
		u,
		"resetRetryWaitMaxSec",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpcloudProvider) ResetRetryWaitMinSec() {
	_jsii_.InvokeVoid(
		u,
		"resetRetryWaitMinSec",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpcloudProvider) ResetUsername() {
	_jsii_.InvokeVoid(
		u,
		"resetUsername",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpcloudProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpcloudProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		u,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpcloudProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpcloudProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpcloudProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UpcloudProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

