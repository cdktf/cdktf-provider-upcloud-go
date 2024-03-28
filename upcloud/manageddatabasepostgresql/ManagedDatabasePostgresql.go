// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasepostgresql

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v14/manageddatabasepostgresql/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/managed_database_postgresql upcloud_managed_database_postgresql}.
type ManagedDatabasePostgresql interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Components() ManagedDatabasePostgresqlComponentsList
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
	Network() ManagedDatabasePostgresqlNetworkList
	NetworkInput() interface{}
	// The tree node.
	Node() constructs.Node
	NodeStates() ManagedDatabasePostgresqlNodeStatesList
	Plan() *string
	SetPlan(val *string)
	PlanInput() *string
	Powered() interface{}
	SetPowered(val interface{})
	PoweredInput() interface{}
	PrimaryDatabase() *string
	Properties() ManagedDatabasePostgresqlPropertiesOutputReference
	PropertiesInput() *ManagedDatabasePostgresqlProperties
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
	Sslmode() *string
	State() *string
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
	PutProperties(value *ManagedDatabasePostgresqlProperties)
	ResetId()
	ResetMaintenanceWindowDow()
	ResetMaintenanceWindowTime()
	ResetNetwork()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPowered()
	ResetProperties()
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

// The jsii proxy struct for ManagedDatabasePostgresql
type jsiiProxy_ManagedDatabasePostgresql struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ManagedDatabasePostgresql) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Components() ManagedDatabasePostgresqlComponentsList {
	var returns ManagedDatabasePostgresqlComponentsList
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) MaintenanceWindowDow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowDow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) MaintenanceWindowDowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowDowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) MaintenanceWindowTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) MaintenanceWindowTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Network() ManagedDatabasePostgresqlNetworkList {
	var returns ManagedDatabasePostgresqlNetworkList
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) NetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) NodeStates() ManagedDatabasePostgresqlNodeStatesList {
	var returns ManagedDatabasePostgresqlNodeStatesList
	_jsii_.Get(
		j,
		"nodeStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Plan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) PlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Powered() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"powered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) PoweredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"poweredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) PrimaryDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Properties() ManagedDatabasePostgresqlPropertiesOutputReference {
	var returns ManagedDatabasePostgresqlPropertiesOutputReference
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) PropertiesInput() *ManagedDatabasePostgresqlProperties {
	var returns *ManagedDatabasePostgresqlProperties
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) ServiceHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) ServicePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) ServicePort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) ServiceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) ServiceUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Sslmode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslmode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabasePostgresql) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/managed_database_postgresql upcloud_managed_database_postgresql} Resource.
func NewManagedDatabasePostgresql(scope constructs.Construct, id *string, config *ManagedDatabasePostgresqlConfig) ManagedDatabasePostgresql {
	_init_.Initialize()

	if err := validateNewManagedDatabasePostgresqlParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabasePostgresql{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresql",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.0/docs/resources/managed_database_postgresql upcloud_managed_database_postgresql} Resource.
func NewManagedDatabasePostgresql_Override(m ManagedDatabasePostgresql, scope constructs.Construct, id *string, config *ManagedDatabasePostgresqlConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresql",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetMaintenanceWindowDow(val *string) {
	if err := j.validateSetMaintenanceWindowDowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindowDow",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetMaintenanceWindowTime(val *string) {
	if err := j.validateSetMaintenanceWindowTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindowTime",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetPlan(val *string) {
	if err := j.validateSetPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plan",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetPowered(val interface{}) {
	if err := j.validateSetPoweredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powered",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabasePostgresql)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a ManagedDatabasePostgresql resource upon running "cdktf plan <stack-name>".
func ManagedDatabasePostgresql_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateManagedDatabasePostgresql_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresql",
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
func ManagedDatabasePostgresql_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDatabasePostgresql_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresql",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagedDatabasePostgresql_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDatabasePostgresql_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresql",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagedDatabasePostgresql_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDatabasePostgresql_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresql",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ManagedDatabasePostgresql_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-upcloud.managedDatabasePostgresql.ManagedDatabasePostgresql",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresql) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabasePostgresql) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabasePostgresql) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabasePostgresql) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabasePostgresql) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresql) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresql) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabasePostgresql) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabasePostgresql) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabasePostgresql) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabasePostgresql) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) PutNetwork(value interface{}) {
	if err := m.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNetwork",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) PutProperties(value *ManagedDatabasePostgresqlProperties) {
	if err := m.validatePutPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putProperties",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ResetMaintenanceWindowDow() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindowDow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ResetMaintenanceWindowTime() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindowTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ResetNetwork() {
	_jsii_.InvokeVoid(
		m,
		"resetNetwork",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ResetPowered() {
	_jsii_.InvokeVoid(
		m,
		"resetPowered",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ResetProperties() {
	_jsii_.InvokeVoid(
		m,
		"resetProperties",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabasePostgresql) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresql) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabasePostgresql) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

