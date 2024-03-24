// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230701

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisEnterprise_Spec_ARM struct {
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Other properties of the cluster.
	Properties *ClusterProperties_ARM `json:"properties,omitempty"`

	// Sku: The SKU to create, which affects price, performance, and features.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Zones: The Availability Zones where this cluster will be deployed.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisEnterprise_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-07-01"
func (enterprise RedisEnterprise_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (enterprise *RedisEnterprise_Spec_ARM) GetName() string {
	return enterprise.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise"
func (enterprise *RedisEnterprise_Spec_ARM) GetType() string {
	return "Microsoft.Cache/redisEnterprise"
}

// Properties of RedisEnterprise clusters, as opposed to general resource properties like location, tags
type ClusterProperties_ARM struct {
	// MinimumTlsVersion: The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTlsVersion *ClusterProperties_MinimumTlsVersion `json:"minimumTlsVersion,omitempty"`
}

// SKU parameters supplied to the create RedisEnterprise operation.
type Sku_ARM struct {
	// Capacity: The size of the RedisEnterprise cluster. Defaults to 2 or 3 depending on SKU. Valid values are (2, 4, 6, ...)
	// for Enterprise SKUs and (3, 9, 15, ...) for Flash SKUs.
	Capacity *int `json:"capacity,omitempty"`

	// Name: The type of RedisEnterprise cluster to deploy. Possible values: (Enterprise_E10, EnterpriseFlash_F300 etc.)
	Name *Sku_Name `json:"name,omitempty"`
}

// +kubebuilder:validation:Enum={"EnterpriseFlash_F1500","EnterpriseFlash_F300","EnterpriseFlash_F700","Enterprise_E10","Enterprise_E100","Enterprise_E20","Enterprise_E50"}
type Sku_Name string

const (
	Sku_Name_EnterpriseFlash_F1500 = Sku_Name("EnterpriseFlash_F1500")
	Sku_Name_EnterpriseFlash_F300  = Sku_Name("EnterpriseFlash_F300")
	Sku_Name_EnterpriseFlash_F700  = Sku_Name("EnterpriseFlash_F700")
	Sku_Name_Enterprise_E10        = Sku_Name("Enterprise_E10")
	Sku_Name_Enterprise_E100       = Sku_Name("Enterprise_E100")
	Sku_Name_Enterprise_E20        = Sku_Name("Enterprise_E20")
	Sku_Name_Enterprise_E50        = Sku_Name("Enterprise_E50")
)

// Mapping from string to Sku_Name
var sku_Name_Values = map[string]Sku_Name{
	"enterpriseflash_f1500": Sku_Name_EnterpriseFlash_F1500,
	"enterpriseflash_f300":  Sku_Name_EnterpriseFlash_F300,
	"enterpriseflash_f700":  Sku_Name_EnterpriseFlash_F700,
	"enterprise_e10":        Sku_Name_Enterprise_E10,
	"enterprise_e100":       Sku_Name_Enterprise_E100,
	"enterprise_e20":        Sku_Name_Enterprise_E20,
	"enterprise_e50":        Sku_Name_Enterprise_E50,
}
