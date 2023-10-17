// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseredis

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedis",
		reflect.TypeOf((*ManagedDatabaseRedis)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "components", GoGetter: "Components"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowDow", GoGetter: "MaintenanceWindowDow"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowDowInput", GoGetter: "MaintenanceWindowDowInput"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowTime", GoGetter: "MaintenanceWindowTime"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowTimeInput", GoGetter: "MaintenanceWindowTimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodeStates", GoGetter: "NodeStates"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "plan", GoGetter: "Plan"},
			_jsii_.MemberProperty{JsiiProperty: "planInput", GoGetter: "PlanInput"},
			_jsii_.MemberProperty{JsiiProperty: "powered", GoGetter: "Powered"},
			_jsii_.MemberProperty{JsiiProperty: "poweredInput", GoGetter: "PoweredInput"},
			_jsii_.MemberProperty{JsiiProperty: "primaryDatabase", GoGetter: "PrimaryDatabase"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putProperties", GoMethod: "PutProperties"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaintenanceWindowDow", GoMethod: "ResetMaintenanceWindowDow"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaintenanceWindowTime", GoMethod: "ResetMaintenanceWindowTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPowered", GoMethod: "ResetPowered"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetTitle", GoMethod: "ResetTitle"},
			_jsii_.MemberProperty{JsiiProperty: "serviceHost", GoGetter: "ServiceHost"},
			_jsii_.MemberProperty{JsiiProperty: "servicePassword", GoGetter: "ServicePassword"},
			_jsii_.MemberProperty{JsiiProperty: "servicePort", GoGetter: "ServicePort"},
			_jsii_.MemberProperty{JsiiProperty: "serviceUri", GoGetter: "ServiceUri"},
			_jsii_.MemberProperty{JsiiProperty: "serviceUsername", GoGetter: "ServiceUsername"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberProperty{JsiiProperty: "titleInput", GoGetter: "TitleInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "zone", GoGetter: "Zone"},
			_jsii_.MemberProperty{JsiiProperty: "zoneInput", GoGetter: "ZoneInput"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedDatabaseRedis{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisComponents",
		reflect.TypeOf((*ManagedDatabaseRedisComponents)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisComponentsList",
		reflect.TypeOf((*ManagedDatabaseRedisComponentsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedDatabaseRedisComponentsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisComponentsOutputReference",
		reflect.TypeOf((*ManagedDatabaseRedisComponentsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "component", GoGetter: "Component"},
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
			_jsii_.MemberProperty{JsiiProperty: "host", GoGetter: "Host"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "route", GoGetter: "Route"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "usage", GoGetter: "Usage"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedDatabaseRedisComponentsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisConfig",
		reflect.TypeOf((*ManagedDatabaseRedisConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisNodeStates",
		reflect.TypeOf((*ManagedDatabaseRedisNodeStates)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisNodeStatesList",
		reflect.TypeOf((*ManagedDatabaseRedisNodeStatesList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedDatabaseRedisNodeStatesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisNodeStatesOutputReference",
		reflect.TypeOf((*ManagedDatabaseRedisNodeStatesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedDatabaseRedisNodeStatesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisProperties",
		reflect.TypeOf((*ManagedDatabaseRedisProperties)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisPropertiesMigration",
		reflect.TypeOf((*ManagedDatabaseRedisPropertiesMigration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisPropertiesMigrationOutputReference",
		reflect.TypeOf((*ManagedDatabaseRedisPropertiesMigrationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dbname", GoGetter: "Dbname"},
			_jsii_.MemberProperty{JsiiProperty: "dbnameInput", GoGetter: "DbnameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "host", GoGetter: "Host"},
			_jsii_.MemberProperty{JsiiProperty: "hostInput", GoGetter: "HostInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreDbs", GoGetter: "IgnoreDbs"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreDbsInput", GoGetter: "IgnoreDbsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDbname", GoMethod: "ResetDbname"},
			_jsii_.MemberMethod{JsiiMethod: "resetHost", GoMethod: "ResetHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreDbs", GoMethod: "ResetIgnoreDbs"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetPort", GoMethod: "ResetPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetSsl", GoMethod: "ResetSsl"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsername", GoMethod: "ResetUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ssl", GoGetter: "Ssl"},
			_jsii_.MemberProperty{JsiiProperty: "sslInput", GoGetter: "SslInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameInput", GoGetter: "UsernameInput"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedDatabaseRedisPropertiesMigrationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseRedis.ManagedDatabaseRedisPropertiesOutputReference",
		reflect.TypeOf((*ManagedDatabaseRedisPropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "automaticUtilityNetworkIpFilter", GoGetter: "AutomaticUtilityNetworkIpFilter"},
			_jsii_.MemberProperty{JsiiProperty: "automaticUtilityNetworkIpFilterInput", GoGetter: "AutomaticUtilityNetworkIpFilterInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "ipFilter", GoGetter: "IpFilter"},
			_jsii_.MemberProperty{JsiiProperty: "ipFilterInput", GoGetter: "IpFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "migration", GoGetter: "Migration"},
			_jsii_.MemberProperty{JsiiProperty: "migrationInput", GoGetter: "MigrationInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicAccess", GoGetter: "PublicAccess"},
			_jsii_.MemberProperty{JsiiProperty: "publicAccessInput", GoGetter: "PublicAccessInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMigration", GoMethod: "PutMigration"},
			_jsii_.MemberProperty{JsiiProperty: "redisAclChannelsDefault", GoGetter: "RedisAclChannelsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "redisAclChannelsDefaultInput", GoGetter: "RedisAclChannelsDefaultInput"},
			_jsii_.MemberProperty{JsiiProperty: "redisIoThreads", GoGetter: "RedisIoThreads"},
			_jsii_.MemberProperty{JsiiProperty: "redisIoThreadsInput", GoGetter: "RedisIoThreadsInput"},
			_jsii_.MemberProperty{JsiiProperty: "redisLfuDecayTime", GoGetter: "RedisLfuDecayTime"},
			_jsii_.MemberProperty{JsiiProperty: "redisLfuDecayTimeInput", GoGetter: "RedisLfuDecayTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "redisLfuLogFactor", GoGetter: "RedisLfuLogFactor"},
			_jsii_.MemberProperty{JsiiProperty: "redisLfuLogFactorInput", GoGetter: "RedisLfuLogFactorInput"},
			_jsii_.MemberProperty{JsiiProperty: "redisMaxmemoryPolicy", GoGetter: "RedisMaxmemoryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "redisMaxmemoryPolicyInput", GoGetter: "RedisMaxmemoryPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "redisNotifyKeyspaceEvents", GoGetter: "RedisNotifyKeyspaceEvents"},
			_jsii_.MemberProperty{JsiiProperty: "redisNotifyKeyspaceEventsInput", GoGetter: "RedisNotifyKeyspaceEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "redisNumberOfDatabases", GoGetter: "RedisNumberOfDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "redisNumberOfDatabasesInput", GoGetter: "RedisNumberOfDatabasesInput"},
			_jsii_.MemberProperty{JsiiProperty: "redisPersistence", GoGetter: "RedisPersistence"},
			_jsii_.MemberProperty{JsiiProperty: "redisPersistenceInput", GoGetter: "RedisPersistenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "redisPubsubClientOutputBufferLimit", GoGetter: "RedisPubsubClientOutputBufferLimit"},
			_jsii_.MemberProperty{JsiiProperty: "redisPubsubClientOutputBufferLimitInput", GoGetter: "RedisPubsubClientOutputBufferLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "redisSsl", GoGetter: "RedisSsl"},
			_jsii_.MemberProperty{JsiiProperty: "redisSslInput", GoGetter: "RedisSslInput"},
			_jsii_.MemberProperty{JsiiProperty: "redisTimeout", GoGetter: "RedisTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "redisTimeoutInput", GoGetter: "RedisTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutomaticUtilityNetworkIpFilter", GoMethod: "ResetAutomaticUtilityNetworkIpFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpFilter", GoMethod: "ResetIpFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetMigration", GoMethod: "ResetMigration"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicAccess", GoMethod: "ResetPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisAclChannelsDefault", GoMethod: "ResetRedisAclChannelsDefault"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisIoThreads", GoMethod: "ResetRedisIoThreads"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisLfuDecayTime", GoMethod: "ResetRedisLfuDecayTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisLfuLogFactor", GoMethod: "ResetRedisLfuLogFactor"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisMaxmemoryPolicy", GoMethod: "ResetRedisMaxmemoryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisNotifyKeyspaceEvents", GoMethod: "ResetRedisNotifyKeyspaceEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisNumberOfDatabases", GoMethod: "ResetRedisNumberOfDatabases"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisPersistence", GoMethod: "ResetRedisPersistence"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisPubsubClientOutputBufferLimit", GoMethod: "ResetRedisPubsubClientOutputBufferLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisSsl", GoMethod: "ResetRedisSsl"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedisTimeout", GoMethod: "ResetRedisTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedDatabaseRedisPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
