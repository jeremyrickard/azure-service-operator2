// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type Profiles_SecurityPolicy_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The json object that contains properties required to create a security policy
	Properties *SecurityPolicyProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Profiles_SecurityPolicy_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (policy Profiles_SecurityPolicy_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (policy *Profiles_SecurityPolicy_Spec_ARM) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/securityPolicies"
func (policy *Profiles_SecurityPolicy_Spec_ARM) GetType() string {
	return "Microsoft.Cdn/profiles/securityPolicies"
}

// The json object that contains properties required to create a security policy
type SecurityPolicyProperties_ARM struct {
	// Parameters: object which contains security policy parameters
	Parameters *SecurityPolicyPropertiesParameters_ARM `json:"parameters,omitempty"`
}

type SecurityPolicyPropertiesParameters_ARM struct {
	// WebApplicationFirewall: Mutually exclusive with all other properties
	WebApplicationFirewall *SecurityPolicyWebApplicationFirewallParameters_ARM `json:"webApplicationFirewall,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because SecurityPolicyPropertiesParameters_ARM represents a discriminated union (JSON OneOf)
func (parameters SecurityPolicyPropertiesParameters_ARM) MarshalJSON() ([]byte, error) {
	if parameters.WebApplicationFirewall != nil {
		return json.Marshal(parameters.WebApplicationFirewall)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the SecurityPolicyPropertiesParameters_ARM
func (parameters *SecurityPolicyPropertiesParameters_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "WebApplicationFirewall" {
		parameters.WebApplicationFirewall = &SecurityPolicyWebApplicationFirewallParameters_ARM{}
		return json.Unmarshal(data, parameters.WebApplicationFirewall)
	}

	// No error
	return nil
}

type SecurityPolicyWebApplicationFirewallParameters_ARM struct {
	// Associations: Waf associations
	Associations []SecurityPolicyWebApplicationFirewallAssociation_ARM `json:"associations,omitempty"`

	// Type: The type of the Security policy to create.
	Type SecurityPolicyWebApplicationFirewallParameters_Type `json:"type,omitempty"`

	// WafPolicy: Resource ID.
	WafPolicy *ResourceReference_ARM `json:"wafPolicy,omitempty"`
}

// settings for security policy patterns to match
type SecurityPolicyWebApplicationFirewallAssociation_ARM struct {
	// Domains: List of domains.
	Domains []ActivatedResourceReference_ARM `json:"domains,omitempty"`

	// PatternsToMatch: List of paths
	PatternsToMatch []string `json:"patternsToMatch,omitempty"`
}

// +kubebuilder:validation:Enum={"WebApplicationFirewall"}
type SecurityPolicyWebApplicationFirewallParameters_Type string

const SecurityPolicyWebApplicationFirewallParameters_Type_WebApplicationFirewall = SecurityPolicyWebApplicationFirewallParameters_Type("WebApplicationFirewall")

// Mapping from string to SecurityPolicyWebApplicationFirewallParameters_Type
var securityPolicyWebApplicationFirewallParameters_Type_Values = map[string]SecurityPolicyWebApplicationFirewallParameters_Type{
	"webapplicationfirewall": SecurityPolicyWebApplicationFirewallParameters_Type_WebApplicationFirewall,
}
