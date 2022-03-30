// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_golden_gate "github.com/oracle/oci-go-sdk/v64/goldengate"

	oci_common "github.com/oracle/oci-go-sdk/v64/common"
)

func init() {
	RegisterOracleClient("oci_golden_gate.GoldenGateClient", &OracleClient{InitClientFn: initGoldengateGoldenGateClient})
}

func initGoldengateGoldenGateClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_golden_gate.NewGoldenGateClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) GoldenGateClient() *oci_golden_gate.GoldenGateClient {
	return m.GetClient("oci_golden_gate.GoldenGateClient").(*oci_golden_gate.GoldenGateClient)
}
