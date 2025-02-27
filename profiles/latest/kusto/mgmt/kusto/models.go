// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package kusto

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/kusto/mgmt/2020-09-18/kusto"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AzureScaleType = original.AzureScaleType

const (
	Automatic AzureScaleType = original.Automatic
	Manual    AzureScaleType = original.Manual
	None      AzureScaleType = original.None
)

type AzureSkuName = original.AzureSkuName

const (
	DevNoSLAStandardD11V2 AzureSkuName = original.DevNoSLAStandardD11V2
	DevNoSLAStandardE2aV4 AzureSkuName = original.DevNoSLAStandardE2aV4
	StandardD11V2         AzureSkuName = original.StandardD11V2
	StandardD12V2         AzureSkuName = original.StandardD12V2
	StandardD13V2         AzureSkuName = original.StandardD13V2
	StandardD14V2         AzureSkuName = original.StandardD14V2
	StandardDS13V21TBPS   AzureSkuName = original.StandardDS13V21TBPS
	StandardDS13V22TBPS   AzureSkuName = original.StandardDS13V22TBPS
	StandardDS14V23TBPS   AzureSkuName = original.StandardDS14V23TBPS
	StandardDS14V24TBPS   AzureSkuName = original.StandardDS14V24TBPS
	StandardE16asV43TBPS  AzureSkuName = original.StandardE16asV43TBPS
	StandardE16asV44TBPS  AzureSkuName = original.StandardE16asV44TBPS
	StandardE16aV4        AzureSkuName = original.StandardE16aV4
	StandardE2aV4         AzureSkuName = original.StandardE2aV4
	StandardE4aV4         AzureSkuName = original.StandardE4aV4
	StandardE64iV3        AzureSkuName = original.StandardE64iV3
	StandardE8asV41TBPS   AzureSkuName = original.StandardE8asV41TBPS
	StandardE8asV42TBPS   AzureSkuName = original.StandardE8asV42TBPS
	StandardE8aV4         AzureSkuName = original.StandardE8aV4
	StandardL16s          AzureSkuName = original.StandardL16s
	StandardL4s           AzureSkuName = original.StandardL4s
	StandardL8s           AzureSkuName = original.StandardL8s
)

type AzureSkuTier = original.AzureSkuTier

const (
	Basic    AzureSkuTier = original.Basic
	Standard AzureSkuTier = original.Standard
)

type BlobStorageEventType = original.BlobStorageEventType

const (
	MicrosoftStorageBlobCreated BlobStorageEventType = original.MicrosoftStorageBlobCreated
	MicrosoftStorageBlobRenamed BlobStorageEventType = original.MicrosoftStorageBlobRenamed
)

type ClusterPrincipalRole = original.ClusterPrincipalRole

const (
	AllDatabasesAdmin  ClusterPrincipalRole = original.AllDatabasesAdmin
	AllDatabasesViewer ClusterPrincipalRole = original.AllDatabasesViewer
)

type Compression = original.Compression

const (
	CompressionGZip Compression = original.CompressionGZip
	CompressionNone Compression = original.CompressionNone
)

type DatabasePrincipalRole = original.DatabasePrincipalRole

const (
	Admin               DatabasePrincipalRole = original.Admin
	Ingestor            DatabasePrincipalRole = original.Ingestor
	Monitor             DatabasePrincipalRole = original.Monitor
	UnrestrictedViewers DatabasePrincipalRole = original.UnrestrictedViewers
	User                DatabasePrincipalRole = original.User
	Viewer              DatabasePrincipalRole = original.Viewer
)

type DatabasePrincipalType = original.DatabasePrincipalType

const (
	DatabasePrincipalTypeApp   DatabasePrincipalType = original.DatabasePrincipalTypeApp
	DatabasePrincipalTypeGroup DatabasePrincipalType = original.DatabasePrincipalTypeGroup
	DatabasePrincipalTypeUser  DatabasePrincipalType = original.DatabasePrincipalTypeUser
)

type DefaultPrincipalsModificationKind = original.DefaultPrincipalsModificationKind

const (
	DefaultPrincipalsModificationKindNone    DefaultPrincipalsModificationKind = original.DefaultPrincipalsModificationKindNone
	DefaultPrincipalsModificationKindReplace DefaultPrincipalsModificationKind = original.DefaultPrincipalsModificationKindReplace
	DefaultPrincipalsModificationKindUnion   DefaultPrincipalsModificationKind = original.DefaultPrincipalsModificationKindUnion
)

type EngineType = original.EngineType

const (
	V2 EngineType = original.V2
	V3 EngineType = original.V3
)

type EventGridDataFormat = original.EventGridDataFormat

const (
	APACHEAVRO EventGridDataFormat = original.APACHEAVRO
	AVRO       EventGridDataFormat = original.AVRO
	CSV        EventGridDataFormat = original.CSV
	JSON       EventGridDataFormat = original.JSON
	MULTIJSON  EventGridDataFormat = original.MULTIJSON
	ORC        EventGridDataFormat = original.ORC
	PARQUET    EventGridDataFormat = original.PARQUET
	PSV        EventGridDataFormat = original.PSV
	RAW        EventGridDataFormat = original.RAW
	SCSV       EventGridDataFormat = original.SCSV
	SINGLEJSON EventGridDataFormat = original.SINGLEJSON
	SOHSV      EventGridDataFormat = original.SOHSV
	TSV        EventGridDataFormat = original.TSV
	TSVE       EventGridDataFormat = original.TSVE
	TXT        EventGridDataFormat = original.TXT
	W3CLOGFILE EventGridDataFormat = original.W3CLOGFILE
)

type EventHubDataFormat = original.EventHubDataFormat

const (
	EventHubDataFormatAPACHEAVRO EventHubDataFormat = original.EventHubDataFormatAPACHEAVRO
	EventHubDataFormatAVRO       EventHubDataFormat = original.EventHubDataFormatAVRO
	EventHubDataFormatCSV        EventHubDataFormat = original.EventHubDataFormatCSV
	EventHubDataFormatJSON       EventHubDataFormat = original.EventHubDataFormatJSON
	EventHubDataFormatMULTIJSON  EventHubDataFormat = original.EventHubDataFormatMULTIJSON
	EventHubDataFormatORC        EventHubDataFormat = original.EventHubDataFormatORC
	EventHubDataFormatPARQUET    EventHubDataFormat = original.EventHubDataFormatPARQUET
	EventHubDataFormatPSV        EventHubDataFormat = original.EventHubDataFormatPSV
	EventHubDataFormatRAW        EventHubDataFormat = original.EventHubDataFormatRAW
	EventHubDataFormatSCSV       EventHubDataFormat = original.EventHubDataFormatSCSV
	EventHubDataFormatSINGLEJSON EventHubDataFormat = original.EventHubDataFormatSINGLEJSON
	EventHubDataFormatSOHSV      EventHubDataFormat = original.EventHubDataFormatSOHSV
	EventHubDataFormatTSV        EventHubDataFormat = original.EventHubDataFormatTSV
	EventHubDataFormatTSVE       EventHubDataFormat = original.EventHubDataFormatTSVE
	EventHubDataFormatTXT        EventHubDataFormat = original.EventHubDataFormatTXT
	EventHubDataFormatW3CLOGFILE EventHubDataFormat = original.EventHubDataFormatW3CLOGFILE
)

type IdentityType = original.IdentityType

const (
	IdentityTypeNone                       IdentityType = original.IdentityTypeNone
	IdentityTypeSystemAssigned             IdentityType = original.IdentityTypeSystemAssigned
	IdentityTypeSystemAssignedUserAssigned IdentityType = original.IdentityTypeSystemAssignedUserAssigned
	IdentityTypeUserAssigned               IdentityType = original.IdentityTypeUserAssigned
)

type IotHubDataFormat = original.IotHubDataFormat

const (
	IotHubDataFormatAPACHEAVRO IotHubDataFormat = original.IotHubDataFormatAPACHEAVRO
	IotHubDataFormatAVRO       IotHubDataFormat = original.IotHubDataFormatAVRO
	IotHubDataFormatCSV        IotHubDataFormat = original.IotHubDataFormatCSV
	IotHubDataFormatJSON       IotHubDataFormat = original.IotHubDataFormatJSON
	IotHubDataFormatMULTIJSON  IotHubDataFormat = original.IotHubDataFormatMULTIJSON
	IotHubDataFormatORC        IotHubDataFormat = original.IotHubDataFormatORC
	IotHubDataFormatPARQUET    IotHubDataFormat = original.IotHubDataFormatPARQUET
	IotHubDataFormatPSV        IotHubDataFormat = original.IotHubDataFormatPSV
	IotHubDataFormatRAW        IotHubDataFormat = original.IotHubDataFormatRAW
	IotHubDataFormatSCSV       IotHubDataFormat = original.IotHubDataFormatSCSV
	IotHubDataFormatSINGLEJSON IotHubDataFormat = original.IotHubDataFormatSINGLEJSON
	IotHubDataFormatSOHSV      IotHubDataFormat = original.IotHubDataFormatSOHSV
	IotHubDataFormatTSV        IotHubDataFormat = original.IotHubDataFormatTSV
	IotHubDataFormatTSVE       IotHubDataFormat = original.IotHubDataFormatTSVE
	IotHubDataFormatTXT        IotHubDataFormat = original.IotHubDataFormatTXT
	IotHubDataFormatW3CLOGFILE IotHubDataFormat = original.IotHubDataFormatW3CLOGFILE
)

type Kind = original.Kind

const (
	KindDatabase          Kind = original.KindDatabase
	KindReadOnlyFollowing Kind = original.KindReadOnlyFollowing
	KindReadWrite         Kind = original.KindReadWrite
)

type KindBasicDataConnection = original.KindBasicDataConnection

const (
	KindDataConnection KindBasicDataConnection = original.KindDataConnection
	KindEventGrid      KindBasicDataConnection = original.KindEventGrid
	KindEventHub       KindBasicDataConnection = original.KindEventHub
	KindIotHub         KindBasicDataConnection = original.KindIotHub
)

type LanguageExtensionName = original.LanguageExtensionName

const (
	PYTHON LanguageExtensionName = original.PYTHON
	R      LanguageExtensionName = original.R
)

type PrincipalType = original.PrincipalType

const (
	PrincipalTypeApp   PrincipalType = original.PrincipalTypeApp
	PrincipalTypeGroup PrincipalType = original.PrincipalTypeGroup
	PrincipalTypeUser  PrincipalType = original.PrincipalTypeUser
)

type PrincipalsModificationKind = original.PrincipalsModificationKind

const (
	PrincipalsModificationKindNone    PrincipalsModificationKind = original.PrincipalsModificationKindNone
	PrincipalsModificationKindReplace PrincipalsModificationKind = original.PrincipalsModificationKindReplace
	PrincipalsModificationKindUnion   PrincipalsModificationKind = original.PrincipalsModificationKindUnion
)

type ProvisioningState = original.ProvisioningState

const (
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Moving    ProvisioningState = original.Moving
	Running   ProvisioningState = original.Running
	Succeeded ProvisioningState = original.Succeeded
)

type Reason = original.Reason

const (
	AlreadyExists Reason = original.AlreadyExists
	Invalid       Reason = original.Invalid
)

type State = original.State

const (
	StateCreating    State = original.StateCreating
	StateDeleted     State = original.StateDeleted
	StateDeleting    State = original.StateDeleting
	StateRunning     State = original.StateRunning
	StateStarting    State = original.StateStarting
	StateStopped     State = original.StateStopped
	StateStopping    State = original.StateStopping
	StateUnavailable State = original.StateUnavailable
	StateUpdating    State = original.StateUpdating
)

type Type = original.Type

const (
	MicrosoftKustoclustersattachedDatabaseConfigurations Type = original.MicrosoftKustoclustersattachedDatabaseConfigurations
	MicrosoftKustoclustersdatabases                      Type = original.MicrosoftKustoclustersdatabases
)

type AttachedDatabaseConfiguration = original.AttachedDatabaseConfiguration
type AttachedDatabaseConfigurationListResult = original.AttachedDatabaseConfigurationListResult
type AttachedDatabaseConfigurationProperties = original.AttachedDatabaseConfigurationProperties
type AttachedDatabaseConfigurationsClient = original.AttachedDatabaseConfigurationsClient
type AttachedDatabaseConfigurationsCreateOrUpdateFuture = original.AttachedDatabaseConfigurationsCreateOrUpdateFuture
type AttachedDatabaseConfigurationsDeleteFuture = original.AttachedDatabaseConfigurationsDeleteFuture
type AzureCapacity = original.AzureCapacity
type AzureEntityResource = original.AzureEntityResource
type AzureResourceSku = original.AzureResourceSku
type AzureSku = original.AzureSku
type BaseClient = original.BaseClient
type BasicDataConnection = original.BasicDataConnection
type BasicDatabase = original.BasicDatabase
type CheckNameRequest = original.CheckNameRequest
type CheckNameResult = original.CheckNameResult
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Cluster = original.Cluster
type ClusterCheckNameRequest = original.ClusterCheckNameRequest
type ClusterListResult = original.ClusterListResult
type ClusterPrincipalAssignment = original.ClusterPrincipalAssignment
type ClusterPrincipalAssignmentCheckNameRequest = original.ClusterPrincipalAssignmentCheckNameRequest
type ClusterPrincipalAssignmentListResult = original.ClusterPrincipalAssignmentListResult
type ClusterPrincipalAssignmentsClient = original.ClusterPrincipalAssignmentsClient
type ClusterPrincipalAssignmentsCreateOrUpdateFuture = original.ClusterPrincipalAssignmentsCreateOrUpdateFuture
type ClusterPrincipalAssignmentsDeleteFuture = original.ClusterPrincipalAssignmentsDeleteFuture
type ClusterPrincipalProperties = original.ClusterPrincipalProperties
type ClusterProperties = original.ClusterProperties
type ClusterUpdate = original.ClusterUpdate
type ClustersAddLanguageExtensionsFuture = original.ClustersAddLanguageExtensionsFuture
type ClustersClient = original.ClustersClient
type ClustersCreateOrUpdateFuture = original.ClustersCreateOrUpdateFuture
type ClustersDeleteFuture = original.ClustersDeleteFuture
type ClustersDetachFollowerDatabasesFuture = original.ClustersDetachFollowerDatabasesFuture
type ClustersDiagnoseVirtualNetworkFuture = original.ClustersDiagnoseVirtualNetworkFuture
type ClustersRemoveLanguageExtensionsFuture = original.ClustersRemoveLanguageExtensionsFuture
type ClustersStartFuture = original.ClustersStartFuture
type ClustersStopFuture = original.ClustersStopFuture
type ClustersUpdateFuture = original.ClustersUpdateFuture
type DataConnection = original.DataConnection
type DataConnectionCheckNameRequest = original.DataConnectionCheckNameRequest
type DataConnectionListResult = original.DataConnectionListResult
type DataConnectionModel = original.DataConnectionModel
type DataConnectionValidation = original.DataConnectionValidation
type DataConnectionValidationListResult = original.DataConnectionValidationListResult
type DataConnectionValidationResult = original.DataConnectionValidationResult
type DataConnectionsClient = original.DataConnectionsClient
type DataConnectionsCreateOrUpdateFuture = original.DataConnectionsCreateOrUpdateFuture
type DataConnectionsDataConnectionValidationMethodFuture = original.DataConnectionsDataConnectionValidationMethodFuture
type DataConnectionsDeleteFuture = original.DataConnectionsDeleteFuture
type DataConnectionsUpdateFuture = original.DataConnectionsUpdateFuture
type Database = original.Database
type DatabaseListResult = original.DatabaseListResult
type DatabaseModel = original.DatabaseModel
type DatabasePrincipal = original.DatabasePrincipal
type DatabasePrincipalAssignment = original.DatabasePrincipalAssignment
type DatabasePrincipalAssignmentCheckNameRequest = original.DatabasePrincipalAssignmentCheckNameRequest
type DatabasePrincipalAssignmentListResult = original.DatabasePrincipalAssignmentListResult
type DatabasePrincipalAssignmentsClient = original.DatabasePrincipalAssignmentsClient
type DatabasePrincipalAssignmentsCreateOrUpdateFuture = original.DatabasePrincipalAssignmentsCreateOrUpdateFuture
type DatabasePrincipalAssignmentsDeleteFuture = original.DatabasePrincipalAssignmentsDeleteFuture
type DatabasePrincipalListRequest = original.DatabasePrincipalListRequest
type DatabasePrincipalListResult = original.DatabasePrincipalListResult
type DatabasePrincipalProperties = original.DatabasePrincipalProperties
type DatabaseStatistics = original.DatabaseStatistics
type DatabasesClient = original.DatabasesClient
type DatabasesCreateOrUpdateFuture = original.DatabasesCreateOrUpdateFuture
type DatabasesDeleteFuture = original.DatabasesDeleteFuture
type DatabasesUpdateFuture = original.DatabasesUpdateFuture
type DiagnoseVirtualNetworkResult = original.DiagnoseVirtualNetworkResult
type EventGridConnectionProperties = original.EventGridConnectionProperties
type EventGridDataConnection = original.EventGridDataConnection
type EventHubConnectionProperties = original.EventHubConnectionProperties
type EventHubDataConnection = original.EventHubDataConnection
type FollowerDatabaseDefinition = original.FollowerDatabaseDefinition
type FollowerDatabaseListResult = original.FollowerDatabaseListResult
type Identity = original.Identity
type IdentityUserAssignedIdentitiesValue = original.IdentityUserAssignedIdentitiesValue
type IotHubConnectionProperties = original.IotHubConnectionProperties
type IotHubDataConnection = original.IotHubDataConnection
type KeyVaultProperties = original.KeyVaultProperties
type LanguageExtension = original.LanguageExtension
type LanguageExtensionsList = original.LanguageExtensionsList
type ListResourceSkusResult = original.ListResourceSkusResult
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type OptimizedAutoscale = original.OptimizedAutoscale
type ProxyResource = original.ProxyResource
type ReadOnlyFollowingDatabase = original.ReadOnlyFollowingDatabase
type ReadOnlyFollowingDatabaseProperties = original.ReadOnlyFollowingDatabaseProperties
type ReadWriteDatabase = original.ReadWriteDatabase
type ReadWriteDatabaseProperties = original.ReadWriteDatabaseProperties
type Resource = original.Resource
type SkuDescription = original.SkuDescription
type SkuDescriptionList = original.SkuDescriptionList
type SkuLocationInfoItem = original.SkuLocationInfoItem
type TrackedResource = original.TrackedResource
type TrustedExternalTenant = original.TrustedExternalTenant
type VirtualNetworkConfiguration = original.VirtualNetworkConfiguration

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAttachedDatabaseConfigurationsClient(subscriptionID string) AttachedDatabaseConfigurationsClient {
	return original.NewAttachedDatabaseConfigurationsClient(subscriptionID)
}
func NewAttachedDatabaseConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) AttachedDatabaseConfigurationsClient {
	return original.NewAttachedDatabaseConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClusterPrincipalAssignmentsClient(subscriptionID string) ClusterPrincipalAssignmentsClient {
	return original.NewClusterPrincipalAssignmentsClient(subscriptionID)
}
func NewClusterPrincipalAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) ClusterPrincipalAssignmentsClient {
	return original.NewClusterPrincipalAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataConnectionsClient(subscriptionID string) DataConnectionsClient {
	return original.NewDataConnectionsClient(subscriptionID)
}
func NewDataConnectionsClientWithBaseURI(baseURI string, subscriptionID string) DataConnectionsClient {
	return original.NewDataConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasePrincipalAssignmentsClient(subscriptionID string) DatabasePrincipalAssignmentsClient {
	return original.NewDatabasePrincipalAssignmentsClient(subscriptionID)
}
func NewDatabasePrincipalAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) DatabasePrincipalAssignmentsClient {
	return original.NewDatabasePrincipalAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAzureScaleTypeValues() []AzureScaleType {
	return original.PossibleAzureScaleTypeValues()
}
func PossibleAzureSkuNameValues() []AzureSkuName {
	return original.PossibleAzureSkuNameValues()
}
func PossibleAzureSkuTierValues() []AzureSkuTier {
	return original.PossibleAzureSkuTierValues()
}
func PossibleBlobStorageEventTypeValues() []BlobStorageEventType {
	return original.PossibleBlobStorageEventTypeValues()
}
func PossibleClusterPrincipalRoleValues() []ClusterPrincipalRole {
	return original.PossibleClusterPrincipalRoleValues()
}
func PossibleCompressionValues() []Compression {
	return original.PossibleCompressionValues()
}
func PossibleDatabasePrincipalRoleValues() []DatabasePrincipalRole {
	return original.PossibleDatabasePrincipalRoleValues()
}
func PossibleDatabasePrincipalTypeValues() []DatabasePrincipalType {
	return original.PossibleDatabasePrincipalTypeValues()
}
func PossibleDefaultPrincipalsModificationKindValues() []DefaultPrincipalsModificationKind {
	return original.PossibleDefaultPrincipalsModificationKindValues()
}
func PossibleEngineTypeValues() []EngineType {
	return original.PossibleEngineTypeValues()
}
func PossibleEventGridDataFormatValues() []EventGridDataFormat {
	return original.PossibleEventGridDataFormatValues()
}
func PossibleEventHubDataFormatValues() []EventHubDataFormat {
	return original.PossibleEventHubDataFormatValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossibleIotHubDataFormatValues() []IotHubDataFormat {
	return original.PossibleIotHubDataFormatValues()
}
func PossibleKindBasicDataConnectionValues() []KindBasicDataConnection {
	return original.PossibleKindBasicDataConnectionValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLanguageExtensionNameValues() []LanguageExtensionName {
	return original.PossibleLanguageExtensionNameValues()
}
func PossiblePrincipalTypeValues() []PrincipalType {
	return original.PossiblePrincipalTypeValues()
}
func PossiblePrincipalsModificationKindValues() []PrincipalsModificationKind {
	return original.PossiblePrincipalsModificationKindValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
