package manageddatabasemysql

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysql",
		reflect.TypeOf((*ManagedDatabaseMysql)(nil)).Elem(),
		[]_jsii_.Member{
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowDow", GoGetter: "MaintenanceWindowDow"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowDowInput", GoGetter: "MaintenanceWindowDowInput"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowTime", GoGetter: "MaintenanceWindowTime"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowTimeInput", GoGetter: "MaintenanceWindowTimeInput"},
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
			j := jsiiProxy_ManagedDatabaseMysql{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlComponents",
		reflect.TypeOf((*ManagedDatabaseMysqlComponents)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlComponentsList",
		reflect.TypeOf((*ManagedDatabaseMysqlComponentsList)(nil)).Elem(),
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
			j := jsiiProxy_ManagedDatabaseMysqlComponentsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlComponentsOutputReference",
		reflect.TypeOf((*ManagedDatabaseMysqlComponentsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_ManagedDatabaseMysqlComponentsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlConfig",
		reflect.TypeOf((*ManagedDatabaseMysqlConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlNodeStates",
		reflect.TypeOf((*ManagedDatabaseMysqlNodeStates)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlNodeStatesList",
		reflect.TypeOf((*ManagedDatabaseMysqlNodeStatesList)(nil)).Elem(),
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
			j := jsiiProxy_ManagedDatabaseMysqlNodeStatesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlNodeStatesOutputReference",
		reflect.TypeOf((*ManagedDatabaseMysqlNodeStatesOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_ManagedDatabaseMysqlNodeStatesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlProperties",
		reflect.TypeOf((*ManagedDatabaseMysqlProperties)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlPropertiesMigration",
		reflect.TypeOf((*ManagedDatabaseMysqlPropertiesMigration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlPropertiesMigrationOutputReference",
		reflect.TypeOf((*ManagedDatabaseMysqlPropertiesMigrationOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_ManagedDatabaseMysqlPropertiesMigrationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-upcloud.managedDatabaseMysql.ManagedDatabaseMysqlPropertiesOutputReference",
		reflect.TypeOf((*ManagedDatabaseMysqlPropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "adminPassword", GoGetter: "AdminPassword"},
			_jsii_.MemberProperty{JsiiProperty: "adminPasswordInput", GoGetter: "AdminPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "adminUsername", GoGetter: "AdminUsername"},
			_jsii_.MemberProperty{JsiiProperty: "adminUsernameInput", GoGetter: "AdminUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "automaticUtilityNetworkIpFilter", GoGetter: "AutomaticUtilityNetworkIpFilter"},
			_jsii_.MemberProperty{JsiiProperty: "automaticUtilityNetworkIpFilterInput", GoGetter: "AutomaticUtilityNetworkIpFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "backupHour", GoGetter: "BackupHour"},
			_jsii_.MemberProperty{JsiiProperty: "backupHourInput", GoGetter: "BackupHourInput"},
			_jsii_.MemberProperty{JsiiProperty: "backupMinute", GoGetter: "BackupMinute"},
			_jsii_.MemberProperty{JsiiProperty: "backupMinuteInput", GoGetter: "BackupMinuteInput"},
			_jsii_.MemberProperty{JsiiProperty: "binlogRetentionPeriod", GoGetter: "BinlogRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "binlogRetentionPeriodInput", GoGetter: "BinlogRetentionPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectTimeout", GoGetter: "ConnectTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "connectTimeoutInput", GoGetter: "ConnectTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTimeZone", GoGetter: "DefaultTimeZone"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTimeZoneInput", GoGetter: "DefaultTimeZoneInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "groupConcatMaxLen", GoGetter: "GroupConcatMaxLen"},
			_jsii_.MemberProperty{JsiiProperty: "groupConcatMaxLenInput", GoGetter: "GroupConcatMaxLenInput"},
			_jsii_.MemberProperty{JsiiProperty: "informationSchemaStatsExpiry", GoGetter: "InformationSchemaStatsExpiry"},
			_jsii_.MemberProperty{JsiiProperty: "informationSchemaStatsExpiryInput", GoGetter: "InformationSchemaStatsExpiryInput"},
			_jsii_.MemberProperty{JsiiProperty: "innodbFtMinTokenSize", GoGetter: "InnodbFtMinTokenSize"},
			_jsii_.MemberProperty{JsiiProperty: "innodbFtMinTokenSizeInput", GoGetter: "InnodbFtMinTokenSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "innodbFtServerStopwordTable", GoGetter: "InnodbFtServerStopwordTable"},
			_jsii_.MemberProperty{JsiiProperty: "innodbFtServerStopwordTableInput", GoGetter: "InnodbFtServerStopwordTableInput"},
			_jsii_.MemberProperty{JsiiProperty: "innodbLockWaitTimeout", GoGetter: "InnodbLockWaitTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "innodbLockWaitTimeoutInput", GoGetter: "InnodbLockWaitTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "innodbLogBufferSize", GoGetter: "InnodbLogBufferSize"},
			_jsii_.MemberProperty{JsiiProperty: "innodbLogBufferSizeInput", GoGetter: "InnodbLogBufferSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "innodbOnlineAlterLogMaxSize", GoGetter: "InnodbOnlineAlterLogMaxSize"},
			_jsii_.MemberProperty{JsiiProperty: "innodbOnlineAlterLogMaxSizeInput", GoGetter: "InnodbOnlineAlterLogMaxSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "innodbPrintAllDeadlocks", GoGetter: "InnodbPrintAllDeadlocks"},
			_jsii_.MemberProperty{JsiiProperty: "innodbPrintAllDeadlocksInput", GoGetter: "InnodbPrintAllDeadlocksInput"},
			_jsii_.MemberProperty{JsiiProperty: "innodbRollbackOnTimeout", GoGetter: "InnodbRollbackOnTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "innodbRollbackOnTimeoutInput", GoGetter: "InnodbRollbackOnTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "interactiveTimeout", GoGetter: "InteractiveTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "interactiveTimeoutInput", GoGetter: "InteractiveTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalTmpMemStorageEngine", GoGetter: "InternalTmpMemStorageEngine"},
			_jsii_.MemberProperty{JsiiProperty: "internalTmpMemStorageEngineInput", GoGetter: "InternalTmpMemStorageEngineInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipFilter", GoGetter: "IpFilter"},
			_jsii_.MemberProperty{JsiiProperty: "ipFilterInput", GoGetter: "IpFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "longQueryTime", GoGetter: "LongQueryTime"},
			_jsii_.MemberProperty{JsiiProperty: "longQueryTimeInput", GoGetter: "LongQueryTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxAllowedPacket", GoGetter: "MaxAllowedPacket"},
			_jsii_.MemberProperty{JsiiProperty: "maxAllowedPacketInput", GoGetter: "MaxAllowedPacketInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxHeapTableSize", GoGetter: "MaxHeapTableSize"},
			_jsii_.MemberProperty{JsiiProperty: "maxHeapTableSizeInput", GoGetter: "MaxHeapTableSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "migration", GoGetter: "Migration"},
			_jsii_.MemberProperty{JsiiProperty: "migrationInput", GoGetter: "MigrationInput"},
			_jsii_.MemberProperty{JsiiProperty: "netReadTimeout", GoGetter: "NetReadTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "netReadTimeoutInput", GoGetter: "NetReadTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "netWriteTimeout", GoGetter: "NetWriteTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "netWriteTimeoutInput", GoGetter: "NetWriteTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicAccess", GoGetter: "PublicAccess"},
			_jsii_.MemberProperty{JsiiProperty: "publicAccessInput", GoGetter: "PublicAccessInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMigration", GoMethod: "PutMigration"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminPassword", GoMethod: "ResetAdminPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminUsername", GoMethod: "ResetAdminUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutomaticUtilityNetworkIpFilter", GoMethod: "ResetAutomaticUtilityNetworkIpFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackupHour", GoMethod: "ResetBackupHour"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackupMinute", GoMethod: "ResetBackupMinute"},
			_jsii_.MemberMethod{JsiiMethod: "resetBinlogRetentionPeriod", GoMethod: "ResetBinlogRetentionPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectTimeout", GoMethod: "ResetConnectTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTimeZone", GoMethod: "ResetDefaultTimeZone"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupConcatMaxLen", GoMethod: "ResetGroupConcatMaxLen"},
			_jsii_.MemberMethod{JsiiMethod: "resetInformationSchemaStatsExpiry", GoMethod: "ResetInformationSchemaStatsExpiry"},
			_jsii_.MemberMethod{JsiiMethod: "resetInnodbFtMinTokenSize", GoMethod: "ResetInnodbFtMinTokenSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetInnodbFtServerStopwordTable", GoMethod: "ResetInnodbFtServerStopwordTable"},
			_jsii_.MemberMethod{JsiiMethod: "resetInnodbLockWaitTimeout", GoMethod: "ResetInnodbLockWaitTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetInnodbLogBufferSize", GoMethod: "ResetInnodbLogBufferSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetInnodbOnlineAlterLogMaxSize", GoMethod: "ResetInnodbOnlineAlterLogMaxSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetInnodbPrintAllDeadlocks", GoMethod: "ResetInnodbPrintAllDeadlocks"},
			_jsii_.MemberMethod{JsiiMethod: "resetInnodbRollbackOnTimeout", GoMethod: "ResetInnodbRollbackOnTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetInteractiveTimeout", GoMethod: "ResetInteractiveTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetInternalTmpMemStorageEngine", GoMethod: "ResetInternalTmpMemStorageEngine"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpFilter", GoMethod: "ResetIpFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetLongQueryTime", GoMethod: "ResetLongQueryTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxAllowedPacket", GoMethod: "ResetMaxAllowedPacket"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxHeapTableSize", GoMethod: "ResetMaxHeapTableSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetMigration", GoMethod: "ResetMigration"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetReadTimeout", GoMethod: "ResetNetReadTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetWriteTimeout", GoMethod: "ResetNetWriteTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicAccess", GoMethod: "ResetPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetSlowQueryLog", GoMethod: "ResetSlowQueryLog"},
			_jsii_.MemberMethod{JsiiMethod: "resetSortBufferSize", GoMethod: "ResetSortBufferSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetSqlMode", GoMethod: "ResetSqlMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetSqlRequirePrimaryKey", GoMethod: "ResetSqlRequirePrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetTmpTableSize", GoMethod: "ResetTmpTableSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersion", GoMethod: "ResetVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetWaitTimeout", GoMethod: "ResetWaitTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "slowQueryLog", GoGetter: "SlowQueryLog"},
			_jsii_.MemberProperty{JsiiProperty: "slowQueryLogInput", GoGetter: "SlowQueryLogInput"},
			_jsii_.MemberProperty{JsiiProperty: "sortBufferSize", GoGetter: "SortBufferSize"},
			_jsii_.MemberProperty{JsiiProperty: "sortBufferSizeInput", GoGetter: "SortBufferSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "sqlMode", GoGetter: "SqlMode"},
			_jsii_.MemberProperty{JsiiProperty: "sqlModeInput", GoGetter: "SqlModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "sqlRequirePrimaryKey", GoGetter: "SqlRequirePrimaryKey"},
			_jsii_.MemberProperty{JsiiProperty: "sqlRequirePrimaryKeyInput", GoGetter: "SqlRequirePrimaryKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tmpTableSize", GoGetter: "TmpTableSize"},
			_jsii_.MemberProperty{JsiiProperty: "tmpTableSizeInput", GoGetter: "TmpTableSizeInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "versionInput", GoGetter: "VersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "waitTimeout", GoGetter: "WaitTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "waitTimeoutInput", GoGetter: "WaitTimeoutInput"},
		},
		func() interface{} {
			j := jsiiProxy_ManagedDatabaseMysqlPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
