// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Manager Proxy API
//
// Use the Service Manager Proxy API to obtain information about SaaS environments provisioned by Service Manager.
// You can get information such as service types and service environment URLs.
//

package servicemanagerproxy

import (
	"github.com/oracle/oci-go-sdk/v54/common"
)

// ErrorEntity The model for the error entity.
type ErrorEntity struct {

	// A short error code that defines the error.
	Code *string `mandatory:"true" json:"code"`

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`
}

func (m ErrorEntity) String() string {
	return common.PointerString(m)
}
