// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// UpdatePrivateEndpointDetails Update details for a private endpoint.
type UpdatePrivateEndpointDetails struct {

	// The private endpoint display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the private endpoint. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN for the private endpoint.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet within the VCN for the private endpoint.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// An array of network security group (NSG) OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the private endpoint.
	NsgIdList []string `mandatory:"false" json:"nsgIdList"`

	// Free-form tags associated with the resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdatePrivateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePrivateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
