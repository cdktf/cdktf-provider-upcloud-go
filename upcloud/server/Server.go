// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v15/server/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/server upcloud_server}.
type Server interface {
	cdktf.TerraformResource
	BootOrder() *string
	SetBootOrder(val *string)
	BootOrderInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Cpu() *float64
	SetCpu(val *float64)
	CpuInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Firewall() interface{}
	SetFirewall(val interface{})
	FirewallInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Host() *float64
	SetHost(val *float64)
	HostInput() *float64
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	HotResize() interface{}
	SetHotResize(val interface{})
	HotResizeInput() interface{}
	Id() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Login() ServerLoginList
	LoginInput() interface{}
	Mem() *float64
	SetMem(val *float64)
	MemInput() *float64
	Metadata() interface{}
	SetMetadata(val interface{})
	MetadataInput() interface{}
	NetworkInterface() ServerNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	NicModel() *string
	SetNicModel(val *string)
	NicModelInput() *string
	// The tree node.
	Node() constructs.Node
	Plan() *string
	SetPlan(val *string)
	PlanInput() *string
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
	ServerGroup() *string
	SetServerGroup(val *string)
	ServerGroupInput() *string
	SimpleBackup() ServerSimpleBackupList
	SimpleBackupInput() interface{}
	StorageDevices() ServerStorageDevicesList
	StorageDevicesInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	Template() ServerTemplateList
	TemplateInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	VideoModel() *string
	SetVideoModel(val *string)
	VideoModelInput() *string
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
	PutLogin(value interface{})
	PutNetworkInterface(value interface{})
	PutSimpleBackup(value interface{})
	PutStorageDevices(value interface{})
	PutTemplate(value interface{})
	ResetBootOrder()
	ResetCpu()
	ResetFirewall()
	ResetHost()
	ResetHotResize()
	ResetLabels()
	ResetLogin()
	ResetMem()
	ResetMetadata()
	ResetNetworkInterface()
	ResetNicModel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlan()
	ResetServerGroup()
	ResetSimpleBackup()
	ResetStorageDevices()
	ResetTags()
	ResetTemplate()
	ResetTimezone()
	ResetTitle()
	ResetUserData()
	ResetVideoModel()
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

// The jsii proxy struct for Server
type jsiiProxy_Server struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Server) BootOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) BootOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Cpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) CpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Firewall() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) FirewallInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firewallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Host() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) HostInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) HotResize() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hotResize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) HotResizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hotResizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Login() ServerLoginList {
	var returns ServerLoginList
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) LoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Mem() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) MemInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Metadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) NetworkInterface() ServerNetworkInterfaceList {
	var returns ServerNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) NicModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) NicModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Plan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) PlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) ServerGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) ServerGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) SimpleBackup() ServerSimpleBackupList {
	var returns ServerSimpleBackupList
	_jsii_.Get(
		j,
		"simpleBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) SimpleBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"simpleBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) StorageDevices() ServerStorageDevicesList {
	var returns ServerStorageDevicesList
	_jsii_.Get(
		j,
		"storageDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) StorageDevicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageDevicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Template() ServerTemplateList {
	var returns ServerTemplateList
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) VideoModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) VideoModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Server) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/server upcloud_server} Resource.
func NewServer(scope constructs.Construct, id *string, config *ServerConfig) Server {
	_init_.Initialize()

	if err := validateNewServerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Server{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.server.Server",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/server upcloud_server} Resource.
func NewServer_Override(s Server, scope constructs.Construct, id *string, config *ServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.server.Server",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_Server)SetBootOrder(val *string) {
	if err := j.validateSetBootOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootOrder",
		val,
	)
}

func (j *jsiiProxy_Server)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Server)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Server)SetCpu(val *float64) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_Server)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Server)SetFirewall(val interface{}) {
	if err := j.validateSetFirewallParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firewall",
		val,
	)
}

func (j *jsiiProxy_Server)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Server)SetHost(val *float64) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_Server)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_Server)SetHotResize(val interface{}) {
	if err := j.validateSetHotResizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hotResize",
		val,
	)
}

func (j *jsiiProxy_Server)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_Server)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Server)SetMem(val *float64) {
	if err := j.validateSetMemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mem",
		val,
	)
}

func (j *jsiiProxy_Server)SetMetadata(val interface{}) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_Server)SetNicModel(val *string) {
	if err := j.validateSetNicModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nicModel",
		val,
	)
}

func (j *jsiiProxy_Server)SetPlan(val *string) {
	if err := j.validateSetPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"plan",
		val,
	)
}

func (j *jsiiProxy_Server)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Server)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Server)SetServerGroup(val *string) {
	if err := j.validateSetServerGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverGroup",
		val,
	)
}

func (j *jsiiProxy_Server)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Server)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_Server)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_Server)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_Server)SetVideoModel(val *string) {
	if err := j.validateSetVideoModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoModel",
		val,
	)
}

func (j *jsiiProxy_Server)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a Server resource upon running "cdktf plan <stack-name>".
func Server_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.server.Server",
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
func Server_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.server.Server",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Server_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.server.Server",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Server_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-upcloud.server.Server",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Server_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-upcloud.server.Server",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_Server) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_Server) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_Server) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_Server) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Server) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_Server) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_Server) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_Server) PutLogin(value interface{}) {
	if err := s.validatePutLoginParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLogin",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Server) PutNetworkInterface(value interface{}) {
	if err := s.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Server) PutSimpleBackup(value interface{}) {
	if err := s.validatePutSimpleBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSimpleBackup",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Server) PutStorageDevices(value interface{}) {
	if err := s.validatePutStorageDevicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStorageDevices",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Server) PutTemplate(value interface{}) {
	if err := s.validatePutTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTemplate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_Server) ResetBootOrder() {
	_jsii_.InvokeVoid(
		s,
		"resetBootOrder",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetCpu() {
	_jsii_.InvokeVoid(
		s,
		"resetCpu",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetFirewall() {
	_jsii_.InvokeVoid(
		s,
		"resetFirewall",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetHost() {
	_jsii_.InvokeVoid(
		s,
		"resetHost",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetHotResize() {
	_jsii_.InvokeVoid(
		s,
		"resetHotResize",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetLabels() {
	_jsii_.InvokeVoid(
		s,
		"resetLabels",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetLogin() {
	_jsii_.InvokeVoid(
		s,
		"resetLogin",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetMem() {
	_jsii_.InvokeVoid(
		s,
		"resetMem",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetMetadata() {
	_jsii_.InvokeVoid(
		s,
		"resetMetadata",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetNicModel() {
	_jsii_.InvokeVoid(
		s,
		"resetNicModel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetPlan() {
	_jsii_.InvokeVoid(
		s,
		"resetPlan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetServerGroup() {
	_jsii_.InvokeVoid(
		s,
		"resetServerGroup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetSimpleBackup() {
	_jsii_.InvokeVoid(
		s,
		"resetSimpleBackup",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetStorageDevices() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageDevices",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetTemplate() {
	_jsii_.InvokeVoid(
		s,
		"resetTemplate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetTimezone() {
	_jsii_.InvokeVoid(
		s,
		"resetTimezone",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetTitle() {
	_jsii_.InvokeVoid(
		s,
		"resetTitle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetUserData() {
	_jsii_.InvokeVoid(
		s,
		"resetUserData",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) ResetVideoModel() {
	_jsii_.InvokeVoid(
		s,
		"resetVideoModel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Server) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Server) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

