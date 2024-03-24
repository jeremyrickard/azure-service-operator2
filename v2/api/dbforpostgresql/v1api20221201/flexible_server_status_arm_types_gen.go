// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20221201

type FlexibleServer_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: Describes the identity of the application.
	Identity *UserAssignedIdentity_STATUS_ARM `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *Sku_STATUS_ARM `json:"sku,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties of a server.
type ServerProperties_STATUS_ARM struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// AuthConfig: AuthConfig properties of a server.
	AuthConfig *AuthConfig_STATUS_ARM `json:"authConfig,omitempty"`

	// AvailabilityZone: availability zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	// Backup: Backup properties of a server.
	Backup *Backup_STATUS_ARM `json:"backup,omitempty"`

	// CreateMode: The mode to create a new PostgreSQL server.
	CreateMode *ServerProperties_CreateMode_STATUS `json:"createMode,omitempty"`

	// DataEncryption: Data encryption properties of a server.
	DataEncryption *DataEncryption_STATUS_ARM `json:"dataEncryption,omitempty"`

	// FullyQualifiedDomainName: The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// HighAvailability: High availability properties of a server.
	HighAvailability *HighAvailability_STATUS_ARM `json:"highAvailability,omitempty"`

	// MaintenanceWindow: Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindow_STATUS_ARM `json:"maintenanceWindow,omitempty"`

	// MinorVersion: The minor version of the server.
	MinorVersion *string `json:"minorVersion,omitempty"`

	// Network: Network properties of a server. This Network property is required to be passed only in case you want the server
	// to be Private access server.
	Network *Network_STATUS_ARM `json:"network,omitempty"`

	// PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when
	// 'createMode' is 'PointInTimeRestore' or 'GeoRestore'.
	PointInTimeUTC *string `json:"pointInTimeUTC,omitempty"`

	// ReplicaCapacity: Replicas allowed for a server.
	ReplicaCapacity *int `json:"replicaCapacity,omitempty"`

	// ReplicationRole: Replication role of the server
	ReplicationRole *ReplicationRole_STATUS `json:"replicationRole,omitempty"`

	// SourceServerResourceId: The source server resource ID to restore from. It's required when 'createMode' is
	// 'PointInTimeRestore' or 'GeoRestore' or 'Replica'. This property is returned only for Replica server
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	// State: A state of a server that is visible to user.
	State *ServerProperties_State_STATUS `json:"state,omitempty"`

	// Storage: Storage properties of a server.
	Storage *Storage_STATUS_ARM `json:"storage,omitempty"`

	// Version: PostgreSQL Server version.
	Version *ServerVersion_STATUS `json:"version,omitempty"`
}

// Sku information related properties of a server.
type Sku_STATUS_ARM struct {
	// Name: The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.
	Name *string `json:"name,omitempty"`

	// Tier: The tier of the particular SKU, e.g. Burstable.
	Tier *Sku_Tier_STATUS `json:"tier,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// Information describing the identities associated with this application.
type UserAssignedIdentity_STATUS_ARM struct {
	// TenantId: Tenant id of the server.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: the types of identities associated with this resource; currently restricted to 'None and UserAssigned'
	Type *UserAssignedIdentity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: represents user assigned identities map.
	UserAssignedIdentities map[string]UserIdentity_STATUS_ARM `json:"userAssignedIdentities,omitempty"`
}

// Authentication configuration properties of a server
type AuthConfig_STATUS_ARM struct {
	// ActiveDirectoryAuth: If Enabled, Azure Active Directory authentication is enabled.
	ActiveDirectoryAuth *AuthConfig_ActiveDirectoryAuth_STATUS `json:"activeDirectoryAuth,omitempty"`

	// PasswordAuth: If Enabled, Password authentication is enabled.
	PasswordAuth *AuthConfig_PasswordAuth_STATUS `json:"passwordAuth,omitempty"`

	// TenantId: Tenant id of the server.
	TenantId *string `json:"tenantId,omitempty"`
}

// Backup properties of a server
type Backup_STATUS_ARM struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// EarliestRestoreDate: The earliest restore point time (ISO8601 format) for server.
	EarliestRestoreDate *string `json:"earliestRestoreDate,omitempty"`

	// GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.
	GeoRedundantBackup *Backup_GeoRedundantBackup_STATUS `json:"geoRedundantBackup,omitempty"`
}

// Data encryption properties of a server
type DataEncryption_STATUS_ARM struct {
	// PrimaryKeyURI: URI for the key for data encryption for primary server.
	PrimaryKeyURI *string `json:"primaryKeyURI,omitempty"`

	// PrimaryUserAssignedIdentityId: Resource Id for the User assigned identity to be used for data encryption for primary
	// server.
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	// Type: Data encryption type to depict if it is System Managed vs Azure Key vault.
	Type *DataEncryption_Type_STATUS `json:"type,omitempty"`
}

// High availability properties of a server
type HighAvailability_STATUS_ARM struct {
	// Mode: The HA mode for the server.
	Mode *HighAvailability_Mode_STATUS `json:"mode,omitempty"`

	// StandbyAvailabilityZone: availability zone information of the standby.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`

	// State: A state of a HA server that is visible to user.
	State *HighAvailability_State_STATUS `json:"state,omitempty"`
}

// Maintenance window properties of a server.
type MaintenanceWindow_STATUS_ARM struct {
	// CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	// DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	// StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	// StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

// Network properties of a server.
type Network_STATUS_ARM struct {
	// DelegatedSubnetResourceId: Delegated subnet arm resource id. This is required to be passed during create, in case we
	// want the server to be VNET injected, i.e. Private access server. During update, pass this only if we want to update the
	// value for Private DNS zone.
	DelegatedSubnetResourceId *string `json:"delegatedSubnetResourceId,omitempty"`

	// PrivateDnsZoneArmResourceId: Private dns zone arm resource id. This is required to be passed during create, in case we
	// want the server to be VNET injected, i.e. Private access server. During update, pass this only if we want to update the
	// value for Private DNS zone.
	PrivateDnsZoneArmResourceId *string `json:"privateDnsZoneArmResourceId,omitempty"`

	// PublicNetworkAccess: public network access is enabled or not
	PublicNetworkAccess *Network_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`
}

type Sku_Tier_STATUS string

const (
	Sku_Tier_STATUS_Burstable       = Sku_Tier_STATUS("Burstable")
	Sku_Tier_STATUS_GeneralPurpose  = Sku_Tier_STATUS("GeneralPurpose")
	Sku_Tier_STATUS_MemoryOptimized = Sku_Tier_STATUS("MemoryOptimized")
)

// Mapping from string to Sku_Tier_STATUS
var sku_Tier_STATUS_Values = map[string]Sku_Tier_STATUS{
	"burstable":       Sku_Tier_STATUS_Burstable,
	"generalpurpose":  Sku_Tier_STATUS_GeneralPurpose,
	"memoryoptimized": Sku_Tier_STATUS_MemoryOptimized,
}

// Storage properties of a server
type Storage_STATUS_ARM struct {
	// StorageSizeGB: Max storage allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS
var systemData_CreatedByType_STATUS_Values = map[string]SystemData_CreatedByType_STATUS{
	"application":     SystemData_CreatedByType_STATUS_Application,
	"key":             SystemData_CreatedByType_STATUS_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_User,
}

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS
var systemData_LastModifiedByType_STATUS_Values = map[string]SystemData_LastModifiedByType_STATUS{
	"application":     SystemData_LastModifiedByType_STATUS_Application,
	"key":             SystemData_LastModifiedByType_STATUS_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_User,
}

type UserAssignedIdentity_Type_STATUS string

const (
	UserAssignedIdentity_Type_STATUS_None         = UserAssignedIdentity_Type_STATUS("None")
	UserAssignedIdentity_Type_STATUS_UserAssigned = UserAssignedIdentity_Type_STATUS("UserAssigned")
)

// Mapping from string to UserAssignedIdentity_Type_STATUS
var userAssignedIdentity_Type_STATUS_Values = map[string]UserAssignedIdentity_Type_STATUS{
	"none":         UserAssignedIdentity_Type_STATUS_None,
	"userassigned": UserAssignedIdentity_Type_STATUS_UserAssigned,
}

// Describes a single user-assigned identity associated with the application.
type UserIdentity_STATUS_ARM struct {
	// ClientId: the client identifier of the Service Principal which this identity represents.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: the object identifier of the Service Principal which this identity represents.
	PrincipalId *string `json:"principalId,omitempty"`
}
