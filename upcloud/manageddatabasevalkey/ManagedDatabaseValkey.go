// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasevalkey

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/manageddatabasevalkey/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey upcloud_managed_database_valkey}.
type ManagedDatabaseValkey interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Components() ManagedDatabaseValkeyComponentsList
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceWindowDow() *string
	SetMaintenanceWindowDow(val *string)
	MaintenanceWindowDowInput() *string
	MaintenanceWindowTime() *string
	SetMaintenanceWindowTime(val *string)
	MaintenanceWindowTimeInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() ManagedDatabaseValkeyNetworkList
	NetworkInput() interface{}
	// The tree node.
	Node() constructs.Node
	NodeStates() ManagedDatabaseValkeyNodeStatesList
	Plan() *string
	SetPlan(val *string)
	PlanInput() *string
	Powered() interface{}
	SetPowered(val interface{})
	PoweredInput() interface{}
	PrimaryDatabase() *string
	Properties() ManagedDatabaseValkeyPropertiesOutputReference
	PropertiesInput() *ManagedDatabaseValkeyProperties
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ServiceHost() *string
	ServicePassword() *string
	ServicePort() *string
	ServiceUri() *string
	ServiceUsername() *string
	State() *string
	TerminationProtection() interface{}
	SetTerminationProtection(val interface{})
	TerminationProtectionInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Type() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutNetwork(value interface{})
	PutProperties(value *ManagedDatabaseValkeyProperties)
	ResetId()
	ResetLabels()
	ResetMaintenanceWindowDow()
	ResetMaintenanceWindowTime()
	ResetNetwork()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPowered()
	ResetProperties()
	ResetTerminationProtection()
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

// The jsii proxy struct for ManagedDatabaseValkey
type jsiiProxy_ManagedDatabaseValkey struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ManagedDatabaseValkey) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Components() ManagedDatabaseValkeyComponentsList {
	var returns ManagedDatabaseValkeyComponentsList
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) MaintenanceWindowDow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowDow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) MaintenanceWindowDowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowDowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) MaintenanceWindowTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) MaintenanceWindowTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Network() ManagedDatabaseValkeyNetworkList {
	var returns ManagedDatabaseValkeyNetworkList
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) NetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) NodeStates() ManagedDatabaseValkeyNodeStatesList {
	var returns ManagedDatabaseValkeyNodeStatesList
	_jsii_.Get(
		j,
		"nodeStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Plan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) PlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Powered() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"powered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) PoweredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"poweredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) PrimaryDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Properties() ManagedDatabaseValkeyPropertiesOutputReference {
	var returns ManagedDatabaseValkeyPropertiesOutputReference
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) PropertiesInput() *ManagedDatabaseValkeyProperties {
	var returns *ManagedDatabaseValkeyProperties
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) ServiceHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) ServicePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) ServicePort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) ServiceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) ServiceUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) TerminationProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) TerminationProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseValkey) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey upcloud_managed_database_valkey} Resource.
func NewManagedDatabaseValkey(scope constructs.Construct, id *string, config *ManagedDatabaseValkeyConfig) ManagedDatabaseValkey {
	_init_.Initialize()

	if err := validateNewManagedDatabaseValkeyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseValkey{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseValkey.ManagedDatabaseValkey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_valkey upcloud_managed_database_valkey} Resource.
func NewManagedDatabaseValkey_Override(m ManagedDatabaseValkey, scope constructs.Construct, id *string, config *ManagedDatabaseValkeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseValkey.ManagedDatabaseValkey",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetMaintenanceWindowDow(val *string) {
	if err := j.validateSetMaintenanceWindowDowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindowDow",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetMaintenanceWindowTime(val *string) {
	if err := j.validateSetMaintenanceWindowTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindowTime",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetPlan(val *string) {
	if err := j.validateSetPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plan",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetPowered(val interface{}) {
	if err := j.validateSetPoweredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powered",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetTerminationProtection(val interface{}) {
	if err := j.validateSetTerminationProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationProtection",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseValkey)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a ManagedDatabaseValkey resource upon running "cdktf plan <stack-name>".
func ManagedDatabaseValkey_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateManagedDatabaseValkey_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.managedDatabaseValkey.ManagedDatabaseValkey",
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
func ManagedDatabaseValkey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDatabaseValkey_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.managedDatabaseValkey.ManagedDatabaseValkey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagedDatabaseValkey_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDatabaseValkey_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.managedDatabaseValkey.ManagedDatabaseValkey",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagedDatabaseValkey_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDatabaseValkey_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.managedDatabaseValkey.ManagedDatabaseValkey",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ManagedDatabaseValkey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-upcloud.managedDatabaseValkey.ManagedDatabaseValkey",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) PutNetwork(value interface{}) {
	if err := m.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNetwork",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) PutProperties(value *ManagedDatabaseValkeyProperties) {
	if err := m.validatePutPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putProperties",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) ResetLabels() {
	_jsii_.InvokeVoid(
		m,
		"resetLabels",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) ResetMaintenanceWindowDow() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindowDow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) ResetMaintenanceWindowTime() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindowTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) ResetNetwork() {
	_jsii_.InvokeVoid(
		m,
		"resetNetwork",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) ResetPowered() {
	_jsii_.InvokeVoid(
		m,
		"resetPowered",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) ResetProperties() {
	_jsii_.InvokeVoid(
		m,
		"resetProperties",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) ResetTerminationProtection() {
	_jsii_.InvokeVoid(
		m,
		"resetTerminationProtection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseValkey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseValkey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

