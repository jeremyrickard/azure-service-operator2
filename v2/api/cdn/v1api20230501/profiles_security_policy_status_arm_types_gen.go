// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import "encoding/json"

type Profiles_SecurityPolicy_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: The json object that contains properties required to create a security policy
	Properties *SecurityPolicyProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: Read only system data
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The json object that contains properties required to create a security policy
type SecurityPolicyProperties_STATUS_ARM struct {
	DeploymentStatus *SecurityPolicyProperties_DeploymentStatus_STATUS `json:"deploymentStatus,omitempty"`

	// Parameters: object which contains security policy parameters
	Parameters *SecurityPolicyPropertiesParameters_STATUS_ARM `json:"parameters,omitempty"`

	// ProfileName: The name of the profile which holds the security policy.
	ProfileName *string `json:"profileName,omitempty"`

	// ProvisioningState: Provisioning status
	ProvisioningState *SecurityPolicyProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

type SecurityPolicyProperties_DeploymentStatus_STATUS string

const (
	SecurityPolicyProperties_DeploymentStatus_STATUS_Failed     = SecurityPolicyProperties_DeploymentStatus_STATUS("Failed")
	SecurityPolicyProperties_DeploymentStatus_STATUS_InProgress = SecurityPolicyProperties_DeploymentStatus_STATUS("InProgress")
	SecurityPolicyProperties_DeploymentStatus_STATUS_NotStarted = SecurityPolicyProperties_DeploymentStatus_STATUS("NotStarted")
	SecurityPolicyProperties_DeploymentStatus_STATUS_Succeeded  = SecurityPolicyProperties_DeploymentStatus_STATUS("Succeeded")
)

// Mapping from string to SecurityPolicyProperties_DeploymentStatus_STATUS
var securityPolicyProperties_DeploymentStatus_STATUS_Values = map[string]SecurityPolicyProperties_DeploymentStatus_STATUS{
	"failed":     SecurityPolicyProperties_DeploymentStatus_STATUS_Failed,
	"inprogress": SecurityPolicyProperties_DeploymentStatus_STATUS_InProgress,
	"notstarted": SecurityPolicyProperties_DeploymentStatus_STATUS_NotStarted,
	"succeeded":  SecurityPolicyProperties_DeploymentStatus_STATUS_Succeeded,
}

type SecurityPolicyProperties_ProvisioningState_STATUS string

const (
	SecurityPolicyProperties_ProvisioningState_STATUS_Creating  = SecurityPolicyProperties_ProvisioningState_STATUS("Creating")
	SecurityPolicyProperties_ProvisioningState_STATUS_Deleting  = SecurityPolicyProperties_ProvisioningState_STATUS("Deleting")
	SecurityPolicyProperties_ProvisioningState_STATUS_Failed    = SecurityPolicyProperties_ProvisioningState_STATUS("Failed")
	SecurityPolicyProperties_ProvisioningState_STATUS_Succeeded = SecurityPolicyProperties_ProvisioningState_STATUS("Succeeded")
	SecurityPolicyProperties_ProvisioningState_STATUS_Updating  = SecurityPolicyProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to SecurityPolicyProperties_ProvisioningState_STATUS
var securityPolicyProperties_ProvisioningState_STATUS_Values = map[string]SecurityPolicyProperties_ProvisioningState_STATUS{
	"creating":  SecurityPolicyProperties_ProvisioningState_STATUS_Creating,
	"deleting":  SecurityPolicyProperties_ProvisioningState_STATUS_Deleting,
	"failed":    SecurityPolicyProperties_ProvisioningState_STATUS_Failed,
	"succeeded": SecurityPolicyProperties_ProvisioningState_STATUS_Succeeded,
	"updating":  SecurityPolicyProperties_ProvisioningState_STATUS_Updating,
}

type SecurityPolicyPropertiesParameters_STATUS_ARM struct {
	// WebApplicationFirewall: Mutually exclusive with all other properties
	WebApplicationFirewall *SecurityPolicyWebApplicationFirewallParameters_STATUS_ARM `json:"webApplicationFirewall,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because SecurityPolicyPropertiesParameters_STATUS_ARM represents a discriminated union (JSON OneOf)
func (parameters SecurityPolicyPropertiesParameters_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if parameters.WebApplicationFirewall != nil {
		return json.Marshal(parameters.WebApplicationFirewall)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the SecurityPolicyPropertiesParameters_STATUS_ARM
func (parameters *SecurityPolicyPropertiesParameters_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "WebApplicationFirewall" {
		parameters.WebApplicationFirewall = &SecurityPolicyWebApplicationFirewallParameters_STATUS_ARM{}
		return json.Unmarshal(data, parameters.WebApplicationFirewall)
	}

	// No error
	return nil
}

type SecurityPolicyWebApplicationFirewallParameters_STATUS_ARM struct {
	// Associations: Waf associations
	Associations []SecurityPolicyWebApplicationFirewallAssociation_STATUS_ARM `json:"associations,omitempty"`

	// Type: The type of the Security policy to create.
	Type SecurityPolicyWebApplicationFirewallParameters_Type_STATUS `json:"type,omitempty"`

	// WafPolicy: Resource ID.
	WafPolicy *ResourceReference_STATUS_ARM `json:"wafPolicy,omitempty"`
}

// settings for security policy patterns to match
type SecurityPolicyWebApplicationFirewallAssociation_STATUS_ARM struct {
	// Domains: List of domains.
	Domains []ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded_ARM `json:"domains,omitempty"`

	// PatternsToMatch: List of paths
	PatternsToMatch []string `json:"patternsToMatch,omitempty"`
}

type SecurityPolicyWebApplicationFirewallParameters_Type_STATUS string

const SecurityPolicyWebApplicationFirewallParameters_Type_STATUS_WebApplicationFirewall = SecurityPolicyWebApplicationFirewallParameters_Type_STATUS("WebApplicationFirewall")

// Mapping from string to SecurityPolicyWebApplicationFirewallParameters_Type_STATUS
var securityPolicyWebApplicationFirewallParameters_Type_STATUS_Values = map[string]SecurityPolicyWebApplicationFirewallParameters_Type_STATUS{
	"webapplicationfirewall": SecurityPolicyWebApplicationFirewallParameters_Type_STATUS_WebApplicationFirewall,
}

// Reference to another resource along with its state.
type ActivatedResourceReference_STATUS_Profiles_SecurityPolicy_SubResourceEmbedded_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}
