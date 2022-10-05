package manageddatabasemysql

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v3/manageddatabasemysql/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql upcloud_managed_database_mysql}.
type ManagedDatabaseMysql interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Components() ManagedDatabaseMysqlComponentsList
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	// The tree node.
	Node() constructs.Node
	NodeStates() ManagedDatabaseMysqlNodeStatesList
	Plan() *string
	SetPlan(val *string)
	PlanInput() *string
	Powered() interface{}
	SetPowered(val interface{})
	PoweredInput() interface{}
	PrimaryDatabase() *string
	Properties() ManagedDatabaseMysqlPropertiesOutputReference
	PropertiesInput() *ManagedDatabaseMysqlProperties
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutProperties(value *ManagedDatabaseMysqlProperties)
	ResetId()
	ResetMaintenanceWindowDow()
	ResetMaintenanceWindowTime()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPowered()
	ResetProperties()
	ResetTitle()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ManagedDatabaseMysql
type jsiiProxy_ManagedDatabaseMysql struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ManagedDatabaseMysql) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Components() ManagedDatabaseMysqlComponentsList {
	var returns ManagedDatabaseMysqlComponentsList
	_jsii_.Get(
		j,
		"components",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) MaintenanceWindowDow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowDow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) MaintenanceWindowDowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowDowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) MaintenanceWindowTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) MaintenanceWindowTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) NodeStates() ManagedDatabaseMysqlNodeStatesList {
	var returns ManagedDatabaseMysqlNodeStatesList
	_jsii_.Get(
		j,
		"nodeStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Plan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) PlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Powered() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"powered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) PoweredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"poweredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) PrimaryDatabase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Properties() ManagedDatabaseMysqlPropertiesOutputReference {
	var returns ManagedDatabaseMysqlPropertiesOutputReference
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) PropertiesInput() *ManagedDatabaseMysqlProperties {
	var returns *ManagedDatabaseMysqlProperties
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) ServiceHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) ServicePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) ServicePort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) ServiceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) ServiceUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseMysql) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql upcloud_managed_database_mysql} Resource.
func NewManagedDatabaseMysql(scope constructs.Construct, id *string, config *ManagedDatabaseMysqlConfig) ManagedDatabaseMysql {
	_init_.Initialize()

	if err := validateNewManagedDatabaseMysqlParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseMysql{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysql",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_mysql upcloud_managed_database_mysql} Resource.
func NewManagedDatabaseMysql_Override(m ManagedDatabaseMysql, scope constructs.Construct, id *string, config *ManagedDatabaseMysqlConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysql",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetMaintenanceWindowDow(val *string) {
	if err := j.validateSetMaintenanceWindowDowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindowDow",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetMaintenanceWindowTime(val *string) {
	if err := j.validateSetMaintenanceWindowTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindowTime",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetPlan(val *string) {
	if err := j.validateSetPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plan",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetPowered(val interface{}) {
	if err := j.validateSetPoweredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"powered",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseMysql)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
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
func ManagedDatabaseMysql_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDatabaseMysql_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysql",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ManagedDatabaseMysql_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysql",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ManagedDatabaseMysql) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_ManagedDatabaseMysql) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseMysql) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseMysql) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseMysql) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseMysql) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseMysql) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseMysql) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseMysql) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseMysql) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseMysql) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseMysql) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_ManagedDatabaseMysql) PutProperties(value *ManagedDatabaseMysqlProperties) {
	if err := m.validatePutPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putProperties",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseMysql) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysql) ResetMaintenanceWindowDow() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindowDow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysql) ResetMaintenanceWindowTime() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindowTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysql) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysql) ResetPowered() {
	_jsii_.InvokeVoid(
		m,
		"resetPowered",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysql) ResetProperties() {
	_jsii_.InvokeVoid(
		m,
		"resetProperties",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysql) ResetTitle() {
	_jsii_.InvokeVoid(
		m,
		"resetTitle",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseMysql) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseMysql) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseMysql) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseMysql) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

