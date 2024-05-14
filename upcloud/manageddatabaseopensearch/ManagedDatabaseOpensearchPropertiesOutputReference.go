// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v14/jsii"

	"github.com/cdktf/cdktf-provider-upcloud-go/upcloud/v14/manageddatabaseopensearch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseOpensearchPropertiesOutputReference interface {
	cdktf.ComplexObject
	ActionAutoCreateIndexEnabled() interface{}
	SetActionAutoCreateIndexEnabled(val interface{})
	ActionAutoCreateIndexEnabledInput() interface{}
	ActionDestructiveRequiresName() interface{}
	SetActionDestructiveRequiresName(val interface{})
	ActionDestructiveRequiresNameInput() interface{}
	AuthFailureListeners() ManagedDatabaseOpensearchPropertiesAuthFailureListenersOutputReference
	AuthFailureListenersInput() *ManagedDatabaseOpensearchPropertiesAuthFailureListeners
	AutomaticUtilityNetworkIpFilter() interface{}
	SetAutomaticUtilityNetworkIpFilter(val interface{})
	AutomaticUtilityNetworkIpFilterInput() interface{}
	ClusterMaxShardsPerNode() *float64
	SetClusterMaxShardsPerNode(val *float64)
	ClusterMaxShardsPerNodeInput() *float64
	ClusterRoutingAllocationNodeConcurrentRecoveries() *float64
	SetClusterRoutingAllocationNodeConcurrentRecoveries(val *float64)
	ClusterRoutingAllocationNodeConcurrentRecoveriesInput() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomDomain() *string
	SetCustomDomain(val *string)
	CustomDomainInput() *string
	EmailSenderName() *string
	SetEmailSenderName(val *string)
	EmailSenderNameInput() *string
	EmailSenderPassword() *string
	SetEmailSenderPassword(val *string)
	EmailSenderPasswordInput() *string
	EmailSenderUsername() *string
	SetEmailSenderUsername(val *string)
	EmailSenderUsernameInput() *string
	EnableSecurityAudit() interface{}
	SetEnableSecurityAudit(val interface{})
	EnableSecurityAuditInput() interface{}
	// Experimental.
	Fqn() *string
	HttpMaxContentLength() *float64
	SetHttpMaxContentLength(val *float64)
	HttpMaxContentLengthInput() *float64
	HttpMaxHeaderSize() *float64
	SetHttpMaxHeaderSize(val *float64)
	HttpMaxHeaderSizeInput() *float64
	HttpMaxInitialLineLength() *float64
	SetHttpMaxInitialLineLength(val *float64)
	HttpMaxInitialLineLengthInput() *float64
	IndexPatterns() *[]*string
	SetIndexPatterns(val *[]*string)
	IndexPatternsInput() *[]*string
	IndexTemplate() ManagedDatabaseOpensearchPropertiesIndexTemplateOutputReference
	IndexTemplateInput() *ManagedDatabaseOpensearchPropertiesIndexTemplate
	IndicesFielddataCacheSize() *float64
	SetIndicesFielddataCacheSize(val *float64)
	IndicesFielddataCacheSizeInput() *float64
	IndicesMemoryIndexBufferSize() *float64
	SetIndicesMemoryIndexBufferSize(val *float64)
	IndicesMemoryIndexBufferSizeInput() *float64
	IndicesMemoryMaxIndexBufferSize() *float64
	SetIndicesMemoryMaxIndexBufferSize(val *float64)
	IndicesMemoryMaxIndexBufferSizeInput() *float64
	IndicesMemoryMinIndexBufferSize() *float64
	SetIndicesMemoryMinIndexBufferSize(val *float64)
	IndicesMemoryMinIndexBufferSizeInput() *float64
	IndicesQueriesCacheSize() *float64
	SetIndicesQueriesCacheSize(val *float64)
	IndicesQueriesCacheSizeInput() *float64
	IndicesQueryBoolMaxClauseCount() *float64
	SetIndicesQueryBoolMaxClauseCount(val *float64)
	IndicesQueryBoolMaxClauseCountInput() *float64
	IndicesRecoveryMaxBytesPerSec() *float64
	SetIndicesRecoveryMaxBytesPerSec(val *float64)
	IndicesRecoveryMaxBytesPerSecInput() *float64
	IndicesRecoveryMaxConcurrentFileChunks() *float64
	SetIndicesRecoveryMaxConcurrentFileChunks(val *float64)
	IndicesRecoveryMaxConcurrentFileChunksInput() *float64
	InternalValue() *ManagedDatabaseOpensearchProperties
	SetInternalValue(val *ManagedDatabaseOpensearchProperties)
	IpFilter() *[]*string
	SetIpFilter(val *[]*string)
	IpFilterInput() *[]*string
	IsmEnabled() interface{}
	SetIsmEnabled(val interface{})
	IsmEnabledInput() interface{}
	IsmHistoryEnabled() interface{}
	SetIsmHistoryEnabled(val interface{})
	IsmHistoryEnabledInput() interface{}
	IsmHistoryMaxAge() *float64
	SetIsmHistoryMaxAge(val *float64)
	IsmHistoryMaxAgeInput() *float64
	IsmHistoryMaxDocs() *float64
	SetIsmHistoryMaxDocs(val *float64)
	IsmHistoryMaxDocsInput() *float64
	IsmHistoryRolloverCheckPeriod() *float64
	SetIsmHistoryRolloverCheckPeriod(val *float64)
	IsmHistoryRolloverCheckPeriodInput() *float64
	IsmHistoryRolloverRetentionPeriod() *float64
	SetIsmHistoryRolloverRetentionPeriod(val *float64)
	IsmHistoryRolloverRetentionPeriodInput() *float64
	KeepIndexRefreshInterval() interface{}
	SetKeepIndexRefreshInterval(val interface{})
	KeepIndexRefreshIntervalInput() interface{}
	Openid() ManagedDatabaseOpensearchPropertiesOpenidOutputReference
	OpenidInput() *ManagedDatabaseOpensearchPropertiesOpenid
	OpensearchDashboards() ManagedDatabaseOpensearchPropertiesOpensearchDashboardsOutputReference
	OpensearchDashboardsInput() *ManagedDatabaseOpensearchPropertiesOpensearchDashboards
	OverrideMainResponseVersion() interface{}
	SetOverrideMainResponseVersion(val interface{})
	OverrideMainResponseVersionInput() interface{}
	PluginsAlertingFilterByBackendRoles() interface{}
	SetPluginsAlertingFilterByBackendRoles(val interface{})
	PluginsAlertingFilterByBackendRolesInput() interface{}
	PublicAccess() interface{}
	SetPublicAccess(val interface{})
	PublicAccessInput() interface{}
	ReindexRemoteWhitelist() *[]*string
	SetReindexRemoteWhitelist(val *[]*string)
	ReindexRemoteWhitelistInput() *[]*string
	Saml() ManagedDatabaseOpensearchPropertiesSamlOutputReference
	SamlInput() *ManagedDatabaseOpensearchPropertiesSaml
	ScriptMaxCompilationsRate() *string
	SetScriptMaxCompilationsRate(val *string)
	ScriptMaxCompilationsRateInput() *string
	SearchMaxBuckets() *float64
	SetSearchMaxBuckets(val *float64)
	SearchMaxBucketsInput() *float64
	ServiceLog() interface{}
	SetServiceLog(val interface{})
	ServiceLogInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThreadPoolAnalyzeQueueSize() *float64
	SetThreadPoolAnalyzeQueueSize(val *float64)
	ThreadPoolAnalyzeQueueSizeInput() *float64
	ThreadPoolAnalyzeSize() *float64
	SetThreadPoolAnalyzeSize(val *float64)
	ThreadPoolAnalyzeSizeInput() *float64
	ThreadPoolForceMergeSize() *float64
	SetThreadPoolForceMergeSize(val *float64)
	ThreadPoolForceMergeSizeInput() *float64
	ThreadPoolGetQueueSize() *float64
	SetThreadPoolGetQueueSize(val *float64)
	ThreadPoolGetQueueSizeInput() *float64
	ThreadPoolGetSize() *float64
	SetThreadPoolGetSize(val *float64)
	ThreadPoolGetSizeInput() *float64
	ThreadPoolSearchQueueSize() *float64
	SetThreadPoolSearchQueueSize(val *float64)
	ThreadPoolSearchQueueSizeInput() *float64
	ThreadPoolSearchSize() *float64
	SetThreadPoolSearchSize(val *float64)
	ThreadPoolSearchSizeInput() *float64
	ThreadPoolSearchThrottledQueueSize() *float64
	SetThreadPoolSearchThrottledQueueSize(val *float64)
	ThreadPoolSearchThrottledQueueSizeInput() *float64
	ThreadPoolSearchThrottledSize() *float64
	SetThreadPoolSearchThrottledSize(val *float64)
	ThreadPoolSearchThrottledSizeInput() *float64
	ThreadPoolWriteQueueSize() *float64
	SetThreadPoolWriteQueueSize(val *float64)
	ThreadPoolWriteQueueSizeInput() *float64
	ThreadPoolWriteSize() *float64
	SetThreadPoolWriteSize(val *float64)
	ThreadPoolWriteSizeInput() *float64
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAuthFailureListeners(value *ManagedDatabaseOpensearchPropertiesAuthFailureListeners)
	PutIndexTemplate(value *ManagedDatabaseOpensearchPropertiesIndexTemplate)
	PutOpenid(value *ManagedDatabaseOpensearchPropertiesOpenid)
	PutOpensearchDashboards(value *ManagedDatabaseOpensearchPropertiesOpensearchDashboards)
	PutSaml(value *ManagedDatabaseOpensearchPropertiesSaml)
	ResetActionAutoCreateIndexEnabled()
	ResetActionDestructiveRequiresName()
	ResetAuthFailureListeners()
	ResetAutomaticUtilityNetworkIpFilter()
	ResetClusterMaxShardsPerNode()
	ResetClusterRoutingAllocationNodeConcurrentRecoveries()
	ResetCustomDomain()
	ResetEmailSenderName()
	ResetEmailSenderPassword()
	ResetEmailSenderUsername()
	ResetEnableSecurityAudit()
	ResetHttpMaxContentLength()
	ResetHttpMaxHeaderSize()
	ResetHttpMaxInitialLineLength()
	ResetIndexPatterns()
	ResetIndexTemplate()
	ResetIndicesFielddataCacheSize()
	ResetIndicesMemoryIndexBufferSize()
	ResetIndicesMemoryMaxIndexBufferSize()
	ResetIndicesMemoryMinIndexBufferSize()
	ResetIndicesQueriesCacheSize()
	ResetIndicesQueryBoolMaxClauseCount()
	ResetIndicesRecoveryMaxBytesPerSec()
	ResetIndicesRecoveryMaxConcurrentFileChunks()
	ResetIpFilter()
	ResetIsmEnabled()
	ResetIsmHistoryEnabled()
	ResetIsmHistoryMaxAge()
	ResetIsmHistoryMaxDocs()
	ResetIsmHistoryRolloverCheckPeriod()
	ResetIsmHistoryRolloverRetentionPeriod()
	ResetKeepIndexRefreshInterval()
	ResetOpenid()
	ResetOpensearchDashboards()
	ResetOverrideMainResponseVersion()
	ResetPluginsAlertingFilterByBackendRoles()
	ResetPublicAccess()
	ResetReindexRemoteWhitelist()
	ResetSaml()
	ResetScriptMaxCompilationsRate()
	ResetSearchMaxBuckets()
	ResetServiceLog()
	ResetThreadPoolAnalyzeQueueSize()
	ResetThreadPoolAnalyzeSize()
	ResetThreadPoolForceMergeSize()
	ResetThreadPoolGetQueueSize()
	ResetThreadPoolGetSize()
	ResetThreadPoolSearchQueueSize()
	ResetThreadPoolSearchSize()
	ResetThreadPoolSearchThrottledQueueSize()
	ResetThreadPoolSearchThrottledSize()
	ResetThreadPoolWriteQueueSize()
	ResetThreadPoolWriteSize()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedDatabaseOpensearchPropertiesOutputReference
type jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ActionAutoCreateIndexEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionAutoCreateIndexEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ActionAutoCreateIndexEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionAutoCreateIndexEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ActionDestructiveRequiresName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionDestructiveRequiresName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ActionDestructiveRequiresNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionDestructiveRequiresNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) AuthFailureListeners() ManagedDatabaseOpensearchPropertiesAuthFailureListenersOutputReference {
	var returns ManagedDatabaseOpensearchPropertiesAuthFailureListenersOutputReference
	_jsii_.Get(
		j,
		"authFailureListeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) AuthFailureListenersInput() *ManagedDatabaseOpensearchPropertiesAuthFailureListeners {
	var returns *ManagedDatabaseOpensearchPropertiesAuthFailureListeners
	_jsii_.Get(
		j,
		"authFailureListenersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) AutomaticUtilityNetworkIpFilter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUtilityNetworkIpFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) AutomaticUtilityNetworkIpFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUtilityNetworkIpFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ClusterMaxShardsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterMaxShardsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ClusterMaxShardsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterMaxShardsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ClusterRoutingAllocationNodeConcurrentRecoveries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterRoutingAllocationNodeConcurrentRecoveries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ClusterRoutingAllocationNodeConcurrentRecoveriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clusterRoutingAllocationNodeConcurrentRecoveriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) CustomDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) CustomDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) EmailSenderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSenderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) EmailSenderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSenderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) EmailSenderPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSenderPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) EmailSenderPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSenderPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) EmailSenderUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSenderUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) EmailSenderUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailSenderUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) EnableSecurityAudit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSecurityAudit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) EnableSecurityAuditInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSecurityAuditInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) HttpMaxContentLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxContentLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) HttpMaxContentLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxContentLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) HttpMaxHeaderSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxHeaderSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) HttpMaxHeaderSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxHeaderSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) HttpMaxInitialLineLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxInitialLineLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) HttpMaxInitialLineLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpMaxInitialLineLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndexPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndexPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"indexPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndexTemplate() ManagedDatabaseOpensearchPropertiesIndexTemplateOutputReference {
	var returns ManagedDatabaseOpensearchPropertiesIndexTemplateOutputReference
	_jsii_.Get(
		j,
		"indexTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndexTemplateInput() *ManagedDatabaseOpensearchPropertiesIndexTemplate {
	var returns *ManagedDatabaseOpensearchPropertiesIndexTemplate
	_jsii_.Get(
		j,
		"indexTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesFielddataCacheSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesFielddataCacheSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesFielddataCacheSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesFielddataCacheSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesMemoryIndexBufferSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryIndexBufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesMemoryIndexBufferSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryIndexBufferSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesMemoryMaxIndexBufferSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryMaxIndexBufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesMemoryMaxIndexBufferSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryMaxIndexBufferSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesMemoryMinIndexBufferSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryMinIndexBufferSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesMemoryMinIndexBufferSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesMemoryMinIndexBufferSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesQueriesCacheSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesQueriesCacheSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesQueriesCacheSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesQueriesCacheSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesQueryBoolMaxClauseCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesQueryBoolMaxClauseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesQueryBoolMaxClauseCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesQueryBoolMaxClauseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesRecoveryMaxBytesPerSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesRecoveryMaxBytesPerSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesRecoveryMaxBytesPerSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesRecoveryMaxBytesPerSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesRecoveryMaxConcurrentFileChunks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesRecoveryMaxConcurrentFileChunks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IndicesRecoveryMaxConcurrentFileChunksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"indicesRecoveryMaxConcurrentFileChunksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) InternalValue() *ManagedDatabaseOpensearchProperties {
	var returns *ManagedDatabaseOpensearchProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IpFilter() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IpFilterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ismEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ismEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmHistoryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ismHistoryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmHistoryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ismHistoryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmHistoryMaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryMaxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmHistoryMaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryMaxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmHistoryMaxDocs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryMaxDocs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmHistoryMaxDocsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryMaxDocsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmHistoryRolloverCheckPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryRolloverCheckPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmHistoryRolloverCheckPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryRolloverCheckPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmHistoryRolloverRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryRolloverRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) IsmHistoryRolloverRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ismHistoryRolloverRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) KeepIndexRefreshInterval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepIndexRefreshInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) KeepIndexRefreshIntervalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepIndexRefreshIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) Openid() ManagedDatabaseOpensearchPropertiesOpenidOutputReference {
	var returns ManagedDatabaseOpensearchPropertiesOpenidOutputReference
	_jsii_.Get(
		j,
		"openid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) OpenidInput() *ManagedDatabaseOpensearchPropertiesOpenid {
	var returns *ManagedDatabaseOpensearchPropertiesOpenid
	_jsii_.Get(
		j,
		"openidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) OpensearchDashboards() ManagedDatabaseOpensearchPropertiesOpensearchDashboardsOutputReference {
	var returns ManagedDatabaseOpensearchPropertiesOpensearchDashboardsOutputReference
	_jsii_.Get(
		j,
		"opensearchDashboards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) OpensearchDashboardsInput() *ManagedDatabaseOpensearchPropertiesOpensearchDashboards {
	var returns *ManagedDatabaseOpensearchPropertiesOpensearchDashboards
	_jsii_.Get(
		j,
		"opensearchDashboardsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) OverrideMainResponseVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideMainResponseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) OverrideMainResponseVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideMainResponseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) PluginsAlertingFilterByBackendRoles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pluginsAlertingFilterByBackendRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) PluginsAlertingFilterByBackendRolesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pluginsAlertingFilterByBackendRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) PublicAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) PublicAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ReindexRemoteWhitelist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reindexRemoteWhitelist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ReindexRemoteWhitelistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reindexRemoteWhitelistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) Saml() ManagedDatabaseOpensearchPropertiesSamlOutputReference {
	var returns ManagedDatabaseOpensearchPropertiesSamlOutputReference
	_jsii_.Get(
		j,
		"saml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) SamlInput() *ManagedDatabaseOpensearchPropertiesSaml {
	var returns *ManagedDatabaseOpensearchPropertiesSaml
	_jsii_.Get(
		j,
		"samlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ScriptMaxCompilationsRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptMaxCompilationsRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ScriptMaxCompilationsRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptMaxCompilationsRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) SearchMaxBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"searchMaxBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) SearchMaxBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"searchMaxBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ServiceLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ServiceLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolAnalyzeQueueSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolAnalyzeQueueSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolAnalyzeQueueSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolAnalyzeQueueSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolAnalyzeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolAnalyzeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolAnalyzeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolAnalyzeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolForceMergeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolForceMergeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolForceMergeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolForceMergeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolGetQueueSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolGetQueueSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolGetQueueSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolGetQueueSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolGetSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolGetSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolGetSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolGetSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolSearchQueueSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchQueueSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolSearchQueueSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchQueueSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolSearchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolSearchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolSearchThrottledQueueSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchThrottledQueueSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolSearchThrottledQueueSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchThrottledQueueSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolSearchThrottledSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchThrottledSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolSearchThrottledSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolSearchThrottledSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolWriteQueueSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolWriteQueueSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolWriteQueueSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolWriteQueueSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolWriteSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolWriteSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ThreadPoolWriteSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPoolWriteSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewManagedDatabaseOpensearchPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedDatabaseOpensearchPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewManagedDatabaseOpensearchPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedDatabaseOpensearchPropertiesOutputReference_Override(m ManagedDatabaseOpensearchPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-upcloud.managedDatabaseOpensearch.ManagedDatabaseOpensearchPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetActionAutoCreateIndexEnabled(val interface{}) {
	if err := j.validateSetActionAutoCreateIndexEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionAutoCreateIndexEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetActionDestructiveRequiresName(val interface{}) {
	if err := j.validateSetActionDestructiveRequiresNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionDestructiveRequiresName",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetAutomaticUtilityNetworkIpFilter(val interface{}) {
	if err := j.validateSetAutomaticUtilityNetworkIpFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticUtilityNetworkIpFilter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetClusterMaxShardsPerNode(val *float64) {
	if err := j.validateSetClusterMaxShardsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterMaxShardsPerNode",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetClusterRoutingAllocationNodeConcurrentRecoveries(val *float64) {
	if err := j.validateSetClusterRoutingAllocationNodeConcurrentRecoveriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterRoutingAllocationNodeConcurrentRecoveries",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetCustomDomain(val *string) {
	if err := j.validateSetCustomDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDomain",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetEmailSenderName(val *string) {
	if err := j.validateSetEmailSenderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSenderName",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetEmailSenderPassword(val *string) {
	if err := j.validateSetEmailSenderPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSenderPassword",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetEmailSenderUsername(val *string) {
	if err := j.validateSetEmailSenderUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailSenderUsername",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetEnableSecurityAudit(val interface{}) {
	if err := j.validateSetEnableSecurityAuditParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSecurityAudit",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetHttpMaxContentLength(val *float64) {
	if err := j.validateSetHttpMaxContentLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMaxContentLength",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetHttpMaxHeaderSize(val *float64) {
	if err := j.validateSetHttpMaxHeaderSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMaxHeaderSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetHttpMaxInitialLineLength(val *float64) {
	if err := j.validateSetHttpMaxInitialLineLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMaxInitialLineLength",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIndexPatterns(val *[]*string) {
	if err := j.validateSetIndexPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexPatterns",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIndicesFielddataCacheSize(val *float64) {
	if err := j.validateSetIndicesFielddataCacheSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesFielddataCacheSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIndicesMemoryIndexBufferSize(val *float64) {
	if err := j.validateSetIndicesMemoryIndexBufferSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesMemoryIndexBufferSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIndicesMemoryMaxIndexBufferSize(val *float64) {
	if err := j.validateSetIndicesMemoryMaxIndexBufferSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesMemoryMaxIndexBufferSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIndicesMemoryMinIndexBufferSize(val *float64) {
	if err := j.validateSetIndicesMemoryMinIndexBufferSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesMemoryMinIndexBufferSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIndicesQueriesCacheSize(val *float64) {
	if err := j.validateSetIndicesQueriesCacheSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesQueriesCacheSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIndicesQueryBoolMaxClauseCount(val *float64) {
	if err := j.validateSetIndicesQueryBoolMaxClauseCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesQueryBoolMaxClauseCount",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIndicesRecoveryMaxBytesPerSec(val *float64) {
	if err := j.validateSetIndicesRecoveryMaxBytesPerSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesRecoveryMaxBytesPerSec",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIndicesRecoveryMaxConcurrentFileChunks(val *float64) {
	if err := j.validateSetIndicesRecoveryMaxConcurrentFileChunksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indicesRecoveryMaxConcurrentFileChunks",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetInternalValue(val *ManagedDatabaseOpensearchProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIpFilter(val *[]*string) {
	if err := j.validateSetIpFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipFilter",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIsmEnabled(val interface{}) {
	if err := j.validateSetIsmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIsmHistoryEnabled(val interface{}) {
	if err := j.validateSetIsmHistoryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismHistoryEnabled",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIsmHistoryMaxAge(val *float64) {
	if err := j.validateSetIsmHistoryMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismHistoryMaxAge",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIsmHistoryMaxDocs(val *float64) {
	if err := j.validateSetIsmHistoryMaxDocsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismHistoryMaxDocs",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIsmHistoryRolloverCheckPeriod(val *float64) {
	if err := j.validateSetIsmHistoryRolloverCheckPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismHistoryRolloverCheckPeriod",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetIsmHistoryRolloverRetentionPeriod(val *float64) {
	if err := j.validateSetIsmHistoryRolloverRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ismHistoryRolloverRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetKeepIndexRefreshInterval(val interface{}) {
	if err := j.validateSetKeepIndexRefreshIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keepIndexRefreshInterval",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetOverrideMainResponseVersion(val interface{}) {
	if err := j.validateSetOverrideMainResponseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideMainResponseVersion",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetPluginsAlertingFilterByBackendRoles(val interface{}) {
	if err := j.validateSetPluginsAlertingFilterByBackendRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginsAlertingFilterByBackendRoles",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetPublicAccess(val interface{}) {
	if err := j.validateSetPublicAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAccess",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetReindexRemoteWhitelist(val *[]*string) {
	if err := j.validateSetReindexRemoteWhitelistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reindexRemoteWhitelist",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetScriptMaxCompilationsRate(val *string) {
	if err := j.validateSetScriptMaxCompilationsRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptMaxCompilationsRate",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetSearchMaxBuckets(val *float64) {
	if err := j.validateSetSearchMaxBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchMaxBuckets",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetServiceLog(val interface{}) {
	if err := j.validateSetServiceLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLog",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetThreadPoolAnalyzeQueueSize(val *float64) {
	if err := j.validateSetThreadPoolAnalyzeQueueSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolAnalyzeQueueSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetThreadPoolAnalyzeSize(val *float64) {
	if err := j.validateSetThreadPoolAnalyzeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolAnalyzeSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetThreadPoolForceMergeSize(val *float64) {
	if err := j.validateSetThreadPoolForceMergeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolForceMergeSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetThreadPoolGetQueueSize(val *float64) {
	if err := j.validateSetThreadPoolGetQueueSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolGetQueueSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetThreadPoolGetSize(val *float64) {
	if err := j.validateSetThreadPoolGetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolGetSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetThreadPoolSearchQueueSize(val *float64) {
	if err := j.validateSetThreadPoolSearchQueueSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolSearchQueueSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetThreadPoolSearchSize(val *float64) {
	if err := j.validateSetThreadPoolSearchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolSearchSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetThreadPoolSearchThrottledQueueSize(val *float64) {
	if err := j.validateSetThreadPoolSearchThrottledQueueSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolSearchThrottledQueueSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetThreadPoolSearchThrottledSize(val *float64) {
	if err := j.validateSetThreadPoolSearchThrottledSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolSearchThrottledSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetThreadPoolWriteQueueSize(val *float64) {
	if err := j.validateSetThreadPoolWriteQueueSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolWriteQueueSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetThreadPoolWriteSize(val *float64) {
	if err := j.validateSetThreadPoolWriteSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadPoolWriteSize",
		val,
	)
}

func (j *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) PutAuthFailureListeners(value *ManagedDatabaseOpensearchPropertiesAuthFailureListeners) {
	if err := m.validatePutAuthFailureListenersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAuthFailureListeners",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) PutIndexTemplate(value *ManagedDatabaseOpensearchPropertiesIndexTemplate) {
	if err := m.validatePutIndexTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIndexTemplate",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) PutOpenid(value *ManagedDatabaseOpensearchPropertiesOpenid) {
	if err := m.validatePutOpenidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putOpenid",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) PutOpensearchDashboards(value *ManagedDatabaseOpensearchPropertiesOpensearchDashboards) {
	if err := m.validatePutOpensearchDashboardsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putOpensearchDashboards",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) PutSaml(value *ManagedDatabaseOpensearchPropertiesSaml) {
	if err := m.validatePutSamlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSaml",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetActionAutoCreateIndexEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetActionAutoCreateIndexEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetActionDestructiveRequiresName() {
	_jsii_.InvokeVoid(
		m,
		"resetActionDestructiveRequiresName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetAuthFailureListeners() {
	_jsii_.InvokeVoid(
		m,
		"resetAuthFailureListeners",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetAutomaticUtilityNetworkIpFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetAutomaticUtilityNetworkIpFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetClusterMaxShardsPerNode() {
	_jsii_.InvokeVoid(
		m,
		"resetClusterMaxShardsPerNode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetClusterRoutingAllocationNodeConcurrentRecoveries() {
	_jsii_.InvokeVoid(
		m,
		"resetClusterRoutingAllocationNodeConcurrentRecoveries",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetCustomDomain() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomDomain",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetEmailSenderName() {
	_jsii_.InvokeVoid(
		m,
		"resetEmailSenderName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetEmailSenderPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetEmailSenderPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetEmailSenderUsername() {
	_jsii_.InvokeVoid(
		m,
		"resetEmailSenderUsername",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetEnableSecurityAudit() {
	_jsii_.InvokeVoid(
		m,
		"resetEnableSecurityAudit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetHttpMaxContentLength() {
	_jsii_.InvokeVoid(
		m,
		"resetHttpMaxContentLength",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetHttpMaxHeaderSize() {
	_jsii_.InvokeVoid(
		m,
		"resetHttpMaxHeaderSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetHttpMaxInitialLineLength() {
	_jsii_.InvokeVoid(
		m,
		"resetHttpMaxInitialLineLength",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIndexPatterns() {
	_jsii_.InvokeVoid(
		m,
		"resetIndexPatterns",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIndexTemplate() {
	_jsii_.InvokeVoid(
		m,
		"resetIndexTemplate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIndicesFielddataCacheSize() {
	_jsii_.InvokeVoid(
		m,
		"resetIndicesFielddataCacheSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIndicesMemoryIndexBufferSize() {
	_jsii_.InvokeVoid(
		m,
		"resetIndicesMemoryIndexBufferSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIndicesMemoryMaxIndexBufferSize() {
	_jsii_.InvokeVoid(
		m,
		"resetIndicesMemoryMaxIndexBufferSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIndicesMemoryMinIndexBufferSize() {
	_jsii_.InvokeVoid(
		m,
		"resetIndicesMemoryMinIndexBufferSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIndicesQueriesCacheSize() {
	_jsii_.InvokeVoid(
		m,
		"resetIndicesQueriesCacheSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIndicesQueryBoolMaxClauseCount() {
	_jsii_.InvokeVoid(
		m,
		"resetIndicesQueryBoolMaxClauseCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIndicesRecoveryMaxBytesPerSec() {
	_jsii_.InvokeVoid(
		m,
		"resetIndicesRecoveryMaxBytesPerSec",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIndicesRecoveryMaxConcurrentFileChunks() {
	_jsii_.InvokeVoid(
		m,
		"resetIndicesRecoveryMaxConcurrentFileChunks",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIpFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetIpFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIsmEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetIsmEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIsmHistoryEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetIsmHistoryEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIsmHistoryMaxAge() {
	_jsii_.InvokeVoid(
		m,
		"resetIsmHistoryMaxAge",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIsmHistoryMaxDocs() {
	_jsii_.InvokeVoid(
		m,
		"resetIsmHistoryMaxDocs",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIsmHistoryRolloverCheckPeriod() {
	_jsii_.InvokeVoid(
		m,
		"resetIsmHistoryRolloverCheckPeriod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetIsmHistoryRolloverRetentionPeriod() {
	_jsii_.InvokeVoid(
		m,
		"resetIsmHistoryRolloverRetentionPeriod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetKeepIndexRefreshInterval() {
	_jsii_.InvokeVoid(
		m,
		"resetKeepIndexRefreshInterval",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetOpenid() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetOpensearchDashboards() {
	_jsii_.InvokeVoid(
		m,
		"resetOpensearchDashboards",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetOverrideMainResponseVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideMainResponseVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetPluginsAlertingFilterByBackendRoles() {
	_jsii_.InvokeVoid(
		m,
		"resetPluginsAlertingFilterByBackendRoles",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetPublicAccess() {
	_jsii_.InvokeVoid(
		m,
		"resetPublicAccess",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetReindexRemoteWhitelist() {
	_jsii_.InvokeVoid(
		m,
		"resetReindexRemoteWhitelist",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetSaml() {
	_jsii_.InvokeVoid(
		m,
		"resetSaml",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetScriptMaxCompilationsRate() {
	_jsii_.InvokeVoid(
		m,
		"resetScriptMaxCompilationsRate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetSearchMaxBuckets() {
	_jsii_.InvokeVoid(
		m,
		"resetSearchMaxBuckets",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetServiceLog() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetThreadPoolAnalyzeQueueSize() {
	_jsii_.InvokeVoid(
		m,
		"resetThreadPoolAnalyzeQueueSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetThreadPoolAnalyzeSize() {
	_jsii_.InvokeVoid(
		m,
		"resetThreadPoolAnalyzeSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetThreadPoolForceMergeSize() {
	_jsii_.InvokeVoid(
		m,
		"resetThreadPoolForceMergeSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetThreadPoolGetQueueSize() {
	_jsii_.InvokeVoid(
		m,
		"resetThreadPoolGetQueueSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetThreadPoolGetSize() {
	_jsii_.InvokeVoid(
		m,
		"resetThreadPoolGetSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetThreadPoolSearchQueueSize() {
	_jsii_.InvokeVoid(
		m,
		"resetThreadPoolSearchQueueSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetThreadPoolSearchSize() {
	_jsii_.InvokeVoid(
		m,
		"resetThreadPoolSearchSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetThreadPoolSearchThrottledQueueSize() {
	_jsii_.InvokeVoid(
		m,
		"resetThreadPoolSearchThrottledQueueSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetThreadPoolSearchThrottledSize() {
	_jsii_.InvokeVoid(
		m,
		"resetThreadPoolSearchThrottledSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetThreadPoolWriteQueueSize() {
	_jsii_.InvokeVoid(
		m,
		"resetThreadPoolWriteQueueSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetThreadPoolWriteSize() {
	_jsii_.InvokeVoid(
		m,
		"resetThreadPoolWriteSize",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDatabaseOpensearchPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

