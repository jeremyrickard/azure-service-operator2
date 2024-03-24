// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20221001preview

type Namespace_STATUS_ARM struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Identity: Properties of BYOK Identity description
	Identity *Identity_STATUS_ARM `json:"identity,omitempty"`

	// Location: The Geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the namespace.
	Properties *SBNamespaceProperties_STATUS_ARM `json:"properties,omitempty"`

	// Sku: Properties of SKU
	Sku *SBSku_STATUS_ARM `json:"sku,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

// Properties to configure User Assigned Identities for Bring your Own Keys
type Identity_STATUS_ARM struct {
	// PrincipalId: ObjectId from the KeyVault
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: TenantId from the KeyVault
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Type of managed service identity.
	Type *Identity_Type_STATUS `json:"type,omitempty"`

	// UserAssignedIdentities: Properties for User Assigned Identities
	UserAssignedIdentities map[string]UserAssignedIdentity_STATUS_ARM `json:"userAssignedIdentities,omitempty"`
}

// Properties of the namespace.
type SBNamespaceProperties_STATUS_ARM struct {
	// AlternateName: Alternate name for namespace
	AlternateName *string `json:"alternateName,omitempty"`

	// CreatedAt: The time the namespace was created
	CreatedAt *string `json:"createdAt,omitempty"`

	// DisableLocalAuth: This property disables SAS authentication for the Service Bus namespace.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// Encryption: Properties of BYOK Encryption description
	Encryption *Encryption_STATUS_ARM `json:"encryption,omitempty"`

	// MetricId: Identifier for Azure Insights metrics
	MetricId *string `json:"metricId,omitempty"`

	// MinimumTlsVersion: The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion *SBNamespaceProperties_MinimumTlsVersion_STATUS `json:"minimumTlsVersion,omitempty"`

	// PremiumMessagingPartitions: The number of partitions of a Service Bus namespace. This property is only applicable to
	// Premium SKU namespaces. The default value is 1 and possible values are 1, 2 and 4
	PremiumMessagingPartitions *int `json:"premiumMessagingPartitions,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Provisioning state of the namespace.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// PublicNetworkAccess: This determines if traffic is allowed over public network. By default it is enabled.
	PublicNetworkAccess *SBNamespaceProperties_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// ServiceBusEndpoint: Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint *string `json:"serviceBusEndpoint,omitempty"`

	// Status: Status of the namespace.
	Status *string `json:"status,omitempty"`

	// UpdatedAt: The time the namespace was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`

	// ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

// SKU of the namespace.
type SBSku_STATUS_ARM struct {
	// Capacity: Messaging units for your service bus premium namespace. Valid capacities are {1, 2, 4, 8, 16} multiples of
	// your properties.premiumMessagingPartitions setting. For example, If properties.premiumMessagingPartitions is 1 then
	// possible capacity values are 1, 2, 4, 8, and 16. If properties.premiumMessagingPartitions is 4 then possible capacity
	// values are 4, 8, 16, 32 and 64
	Capacity *int `json:"capacity,omitempty"`

	// Name: Name of this SKU.
	Name *SBSku_Name_STATUS `json:"name,omitempty"`

	// Tier: The billing tier of this particular SKU.
	Tier *SBSku_Tier_STATUS `json:"tier,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The type of identity that last modified the resource.
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// Properties to configure Encryption
type Encryption_STATUS_ARM struct {
	// KeySource: Enumerates the possible value of keySource for Encryption
	KeySource *Encryption_KeySource_STATUS `json:"keySource,omitempty"`

	// KeyVaultProperties: Properties of KeyVault
	KeyVaultProperties []KeyVaultProperties_STATUS_ARM `json:"keyVaultProperties,omitempty"`

	// RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption,omitempty"`
}

type Identity_Type_STATUS string

const (
	Identity_Type_STATUS_None                       = Identity_Type_STATUS("None")
	Identity_Type_STATUS_SystemAssigned             = Identity_Type_STATUS("SystemAssigned")
	Identity_Type_STATUS_SystemAssignedUserAssigned = Identity_Type_STATUS("SystemAssigned, UserAssigned")
	Identity_Type_STATUS_UserAssigned               = Identity_Type_STATUS("UserAssigned")
)

// Mapping from string to Identity_Type_STATUS
var identity_Type_STATUS_Values = map[string]Identity_Type_STATUS{
	"none":                         Identity_Type_STATUS_None,
	"systemassigned":               Identity_Type_STATUS_SystemAssigned,
	"systemassigned, userassigned": Identity_Type_STATUS_SystemAssignedUserAssigned,
	"userassigned":                 Identity_Type_STATUS_UserAssigned,
}

// Properties of the PrivateEndpointConnection.
type PrivateEndpointConnection_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`
}

type SBSku_Name_STATUS string

const (
	SBSku_Name_STATUS_Basic    = SBSku_Name_STATUS("Basic")
	SBSku_Name_STATUS_Premium  = SBSku_Name_STATUS("Premium")
	SBSku_Name_STATUS_Standard = SBSku_Name_STATUS("Standard")
)

// Mapping from string to SBSku_Name_STATUS
var sBSku_Name_STATUS_Values = map[string]SBSku_Name_STATUS{
	"basic":    SBSku_Name_STATUS_Basic,
	"premium":  SBSku_Name_STATUS_Premium,
	"standard": SBSku_Name_STATUS_Standard,
}

type SBSku_Tier_STATUS string

const (
	SBSku_Tier_STATUS_Basic    = SBSku_Tier_STATUS("Basic")
	SBSku_Tier_STATUS_Premium  = SBSku_Tier_STATUS("Premium")
	SBSku_Tier_STATUS_Standard = SBSku_Tier_STATUS("Standard")
)

// Mapping from string to SBSku_Tier_STATUS
var sBSku_Tier_STATUS_Values = map[string]SBSku_Tier_STATUS{
	"basic":    SBSku_Tier_STATUS_Basic,
	"premium":  SBSku_Tier_STATUS_Premium,
	"standard": SBSku_Tier_STATUS_Standard,
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

// Recognized Dictionary value.
type UserAssignedIdentity_STATUS_ARM struct {
	// ClientId: Client Id of user assigned identity
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: Principal Id of user assigned identity
	PrincipalId *string `json:"principalId,omitempty"`
}

// Properties to configure keyVault Properties
type KeyVaultProperties_STATUS_ARM struct {
	Identity *UserAssignedIdentityProperties_STATUS_ARM `json:"identity,omitempty"`

	// KeyName: Name of the Key from KeyVault
	KeyName *string `json:"keyName,omitempty"`

	// KeyVaultUri: Uri of KeyVault
	KeyVaultUri *string `json:"keyVaultUri,omitempty"`

	// KeyVersion: Version of KeyVault
	KeyVersion *string `json:"keyVersion,omitempty"`
}

type UserAssignedIdentityProperties_STATUS_ARM struct {
	// UserAssignedIdentity: ARM ID of user Identity selected for encryption
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}
