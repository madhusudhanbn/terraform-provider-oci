package provider

import (
	tf_ai_anomaly_detection "github.com/terraform-providers/terraform-provider-oci/internal/service/ai_anomaly_detection"
	"github.com/terraform-providers/terraform-provider-oci/internal/service/audit"
	"github.com/terraform-providers/terraform-provider-oci/internal/service/budget"
	tf_core "github.com/terraform-providers/terraform-provider-oci/internal/service/core"
	tf_identity "github.com/terraform-providers/terraform-provider-oci/internal/service/identity"
	tf_kms "github.com/terraform-providers/terraform-provider-oci/internal/service/kms"
	tf_load_balancer "github.com/terraform-providers/terraform-provider-oci/internal/service/load_balancer"
)

func init() {
	// ai anomaly service
	RegisterDatasource("oci_ai_anomaly_detection_ai_private_endpoints", tf_ai_anomaly_detection.AiAnomalyDetectionAiPrivateEndpointsDataSource())
	RegisterDatasource("oci_ai_anomaly_detection_ai_private_endpoint", tf_ai_anomaly_detection.AiAnomalyDetectionAiPrivateEndpointDataSource())
	RegisterDatasource("oci_ai_anomaly_detection_data_asset", tf_ai_anomaly_detection.AiAnomalyDetectionDataAssetDataSource())
	RegisterDatasource("oci_ai_anomaly_detection_data_assets", tf_ai_anomaly_detection.AiAnomalyDetectionDataAssetsDataSource())
	RegisterDatasource("oci_ai_anomaly_detection_model", tf_ai_anomaly_detection.AiAnomalyDetectionModelDataSource())
	RegisterDatasource("oci_ai_anomaly_detection_models", tf_ai_anomaly_detection.AiAnomalyDetectionModelsDataSource())
	RegisterDatasource("oci_ai_anomaly_detection_project", tf_ai_anomaly_detection.AiAnomalyDetectionProjectDataSource())
	RegisterDatasource("oci_ai_anomaly_detection_projects", tf_ai_anomaly_detection.AiAnomalyDetectionProjectsDataSource())

	// identity service
	RegisterDatasource("oci_identity_smtp_credentials", tf_identity.IdentitySmtpCredentialsDataSource())
	RegisterDatasource("oci_budget_alert_rule", budget.BudgetAlertRuleDataSource())
	RegisterDatasource("oci_budget_budget", budget.BudgetBudgetDataSource())
	RegisterDatasource("oci_budget_alert_rules", budget.BudgetAlertRulesDataSource())
	RegisterDatasource("oci_budget_budgets", budget.BudgetBudgetsDataSource())
	RegisterDatasource("oci_identity_compartment", tf_identity.IdentityCompartmentDataSource())
	RegisterDatasource("oci_identity_compartments", tf_identity.IdentityCompartmentsDataSource())
	RegisterDatasource("oci_identity_network_source", tf_identity.IdentityNetworkSourceDataSource())
	RegisterDatasource("oci_identity_region_subscriptions", tf_identity.IdentityRegionSubscriptionsDataSource())
	RegisterDatasource("oci_identity_iam_work_requests", tf_identity.IdentityIamWorkRequestsDataSource())
	RegisterDatasource("oci_identity_groups", tf_identity.IdentityGroupsDataSource())
	RegisterDatasource("oci_identity_cost_tracking_tags", tf_identity.IdentityCostTrackingTagsDataSource())
	RegisterDatasource("oci_identity_users", tf_identity.IdentityUsersDataSource())
	RegisterDatasource("oci_identity_availability_domains", tf_identity.IdentityAvailabilityDomainsDataSource())
	RegisterDatasource("oci_identity_iam_work_request_logs", tf_identity.IdentityIamWorkRequestLogsDataSource())
	RegisterDatasource("oci_identity_tag_defaults", tf_identity.IdentityTagDefaultsDataSource())
	RegisterDatasource("oci_identity_user_group_memberships", tf_identity.IdentityUserGroupMembershipsDataSource())
	RegisterDatasource("oci_identity_swift_passwords", tf_identity.IdentitySwiftPasswordsDataSource())
	RegisterDatasource("oci_identity_group", tf_identity.IdentityGroupDataSource())
	RegisterDatasource("oci_identity_api_keys", tf_identity.IdentityApiKeysDataSource())
	RegisterDatasource("oci_identity_domain", tf_identity.IdentityDomainDataSource())
	RegisterDatasource("oci_identity_tag_default", tf_identity.IdentityTagDefaultDataSource())
	RegisterDatasource("oci_identity_availability_domain", tf_identity.IdentityAvailabilityDomainDataSource())
	RegisterDatasource("oci_identity_customer_secret_keys", tf_identity.IdentityCustomerSecretKeysDataSource())
	RegisterDatasource("oci_identity_iam_work_request", tf_identity.IdentityIamWorkRequestDataSource())
	RegisterDatasource("oci_identity_iam_work_request_errors", tf_identity.IdentityIamWorkRequestErrorsDataSource())
	RegisterDatasource("oci_identity_allowed_domain_license_types", tf_identity.IdentityAllowedDomainLicenseTypesDataSource())
	RegisterDatasource("oci_identity_idp_group_mappings", tf_identity.IdentityIdpGroupMappingsDataSource())
	RegisterDatasource("oci_identity_dynamic_groups", tf_identity.IdentityDynamicGroupsDataSource())
	RegisterDatasource("oci_identity_tenancy", tf_identity.IdentityTenancyDataSource())
	RegisterDatasource("oci_identity_auth_tokens", tf_identity.IdentityAuthTokensDataSource())
	RegisterDatasource("oci_identity_user", tf_identity.IdentityUserDataSource())
	RegisterDatasource("oci_identity_network_sources", tf_identity.IdentityNetworkSourcesDataSource())
	RegisterDatasource("oci_identity_tags", tf_identity.IdentityTagsDataSource())
	RegisterDatasource("oci_identity_policies", tf_identity.IdentityPoliciesDataSource())
	RegisterDatasource("oci_identity_authentication_policy", tf_identity.IdentityAuthenticationPolicyDataSource())
	RegisterDatasource("oci_identity_regions", tf_identity.IdentityRegionsDataSource())
	RegisterDatasource("oci_identity_tag", tf_identity.IdentityTagDataSource())
	RegisterDatasource("oci_identity_db_credentials", tf_identity.IdentityDbCredentialsDataSource())
	RegisterDatasource("oci_identity_identity_providers", tf_identity.IdentityIdentityProvidersDataSource())
	RegisterDatasource("oci_identity_identity_provider_groups", tf_identity.IdentityIdentityProviderGroupsDataSource())
	RegisterDatasource("oci_identity_tag_namespaces", tf_identity.IdentityTagNamespacesDataSource())
	RegisterDatasource("oci_identity_domains", tf_identity.IdentityDomainsDataSource())
	RegisterDatasource("oci_identity_ui_password", tf_identity.IdentityUiPasswordDataSource())
	RegisterDatasource("oci_identity_fault_domains", tf_identity.IdentityFaultDomainsDataSource())
	RegisterDatasource("oci_audit_configuration", audit.AuditConfigurationDataSource())
	RegisterDatasource("oci_audit_events", audit.AuditAuditEventsDataSource())

	//kms service
	RegisterDatasource("oci_kms_key", tf_kms.KmsKeyDataSource())
	RegisterDatasource("oci_kms_replication_status", tf_kms.KmsReplicationStatusDataSource())
	RegisterDatasource("oci_kms_vault_usage", tf_kms.KmsVaultUsageDataSource())
	RegisterDatasource("oci_kms_keys", tf_kms.KmsKeysDataSource())
	RegisterDatasource("oci_kms_key_version", tf_kms.KmsKeyVersionDataSource())
	RegisterDatasource("oci_kms_key_versions", tf_kms.KmsKeyVersionsDataSource())
	RegisterDatasource("oci_kms_decrypted_data", tf_kms.KmsDecryptedDataDataSource())
	RegisterDatasource("oci_kms_vaults", tf_kms.KmsVaultsDataSource())
	RegisterDatasource("oci_kms_encrypted_data", tf_kms.KmsEncryptedDataDataSource())
	RegisterDatasource("oci_kms_vault", tf_kms.KmsVaultDataSource())
	RegisterDatasource("oci_kms_vault_replicas", tf_kms.KmsVaultReplicasDataSource())

	//load_balancer service
	RegisterDatasource("oci_load_balancer_ssl_cipher_suites", tf_load_balancer.LoadBalancerSslCipherSuitesDataSource())
	RegisterDatasource("oci_load_balancer_listener_rules", tf_load_balancer.LoadBalancerListenerRulesDataSource())
	RegisterDatasource("oci_load_balancer_protocols", tf_load_balancer.LoadBalancerLoadBalancerProtocolsDataSource())
	RegisterDatasource("oci_load_balancer_policies", tf_load_balancer.LoadBalancerLoadBalancerPoliciesDataSource())
	RegisterDatasource("oci_load_balancer_backend_sets", tf_load_balancer.LoadBalancerBackendSetsDataSource())
	RegisterDatasource("oci_load_balancer_rule_set", tf_load_balancer.LoadBalancerRuleSetDataSource())
	RegisterDatasource("oci_load_balancer_rule_sets", tf_load_balancer.LoadBalancerRuleSetsDataSource())
	RegisterDatasource("oci_load_balancer_backend_set_health", tf_load_balancer.LoadBalancerBackendSetHealthDataSource())
	RegisterDatasource("oci_load_balancer_backend_health", tf_load_balancer.LoadBalancerBackendHealthDataSource())
	RegisterDatasource("oci_load_balancer_hostnames", tf_load_balancer.LoadBalancerHostnamesDataSource())
	RegisterDatasource("oci_load_balancer_load_balancers", tf_load_balancer.LoadBalancerLoadBalancersDataSource())
	RegisterDatasource("oci_load_balancer_ssl_cipher_suite", tf_load_balancer.LoadBalancerSslCipherSuiteDataSource())
	RegisterDatasource("oci_load_balancer_load_balancer_routing_policies", tf_load_balancer.LoadBalancerLoadBalancerRoutingPoliciesDataSource())
	RegisterDatasource("oci_load_balancer_certificates", tf_load_balancer.LoadBalancerCertificatesDataSource())
	RegisterDatasource("oci_load_balancer_health", tf_load_balancer.LoadBalancerLoadBalancerHealthDataSource())
	RegisterDatasource("oci_load_balancer_shapes", tf_load_balancer.LoadBalancerLoadBalancerShapesDataSource())
	RegisterDatasource("oci_load_balancer_backends", tf_load_balancer.LoadBalancerBackendsDataSource())
	RegisterDatasource("oci_load_balancer_path_route_sets", tf_load_balancer.LoadBalancerPathRouteSetsDataSource())
	RegisterDatasource("oci_load_balancer_load_balancer_routing_policy", tf_load_balancer.LoadBalancerLoadBalancerRoutingPolicyDataSource())

	// core service
	RegisterDatasource("oci_core_tunnel_security_associations", tf_core.CoreTunnelSecurityAssociationsDataSource())
	RegisterDatasource("oci_core_cluster_networks", tf_core.CoreClusterNetworksDataSource())
	RegisterDatasource("oci_core_volume", tf_core.CoreVolumeDataSource())
	RegisterDatasource("oci_core_ipsec_connection_tunnels", tf_core.CoreIpSecConnectionTunnelsDataSource())
	RegisterDatasource("oci_core_drg_route_distribution_statements", tf_core.CoreDrgRouteDistributionStatementsDataSource())
	RegisterDatasource("oci_core_services", tf_core.CoreServicesDataSource())
	RegisterDatasource("oci_core_compute_capacity_reservations", tf_core.CoreComputeCapacityReservationsDataSource())
	RegisterDatasource("oci_core_private_ips", tf_core.CorePrivateIpsDataSource())
	RegisterDatasource("oci_core_drg_route_table", tf_core.CoreDrgRouteTableDataSource())
	RegisterDatasource("oci_core_app_catalog_listing_resource_versions", tf_core.CoreAppCatalogListingResourceVersionsDataSource())
	RegisterDatasource("oci_core_security_lists", tf_core.CoreSecurityListsDataSource())
	RegisterDatasource("oci_core_app_catalog_listings", tf_core.CoreAppCatalogListingsDataSource())
	RegisterDatasource("oci_core_subnet", tf_core.CoreSubnetDataSource())
	RegisterDatasource("oci_core_image", tf_core.CoreImageDataSource())
	RegisterDatasource("oci_core_image_shapes", tf_core.CoreImageShapesDataSource())
	RegisterDatasource("oci_core_ipsec_algorithm", tf_core.CoreIpsecAlgorithmDataSource())
	RegisterDatasource("oci_core_instance_pool_instances", tf_core.CoreInstancePoolInstancesDataSource())
	RegisterDatasource("oci_core_letter_of_authority", tf_core.CoreLetterOfAuthorityDataSource())
	RegisterDatasource("oci_core_cross_connect", tf_core.CoreCrossConnectDataSource())
	RegisterDatasource("oci_core_cross_connect_status", tf_core.CoreCrossConnectStatusDataSource())
	RegisterDatasource("oci_core_instance_pool", tf_core.CoreInstancePoolDataSource())
	RegisterDatasource("oci_core_virtual_circuit", tf_core.CoreVirtualCircuitDataSource())
	RegisterDatasource("oci_core_app_catalog_subscriptions", tf_core.CoreAppCatalogSubscriptionsDataSource())
	RegisterDatasource("oci_core_volume_backup_policies", tf_core.CoreVolumeBackupPoliciesDataSource())
	RegisterDatasource("oci_core_instance_configuration", tf_core.CoreInstanceConfigurationDataSource())
	RegisterDatasource("oci_core_boot_volume_backups", tf_core.CoreBootVolumeBackupsDataSource())
	RegisterDatasource("oci_core_drg_route_tables", tf_core.CoreDrgRouteTablesDataSource())
	RegisterDatasource("oci_core_cross_connect_port_speed_shapes", tf_core.CoreCrossConnectPortSpeedShapesDataSource())
	RegisterDatasource("oci_core_public_ip", tf_core.CorePublicIpDataSource())
	RegisterDatasource("oci_core_boot_volumes", tf_core.CoreBootVolumesDataSource())
	RegisterDatasource("oci_core_compute_global_image_capability_schemas_version", tf_core.CoreComputeGlobalImageCapabilitySchemasVersionDataSource())
	RegisterDatasource("oci_core_dedicated_vm_host", tf_core.CoreDedicatedVmHostDataSource())
	RegisterDatasource("oci_core_block_volume_replica", tf_core.CoreBlockVolumeReplicaDataSource())
	RegisterDatasource("oci_core_vnic_attachments", tf_core.CoreVnicAttachmentsDataSource())
	RegisterDatasource("oci_core_instance_configurations", tf_core.CoreInstanceConfigurationsDataSource())
	RegisterDatasource("oci_core_ipsec_connection_tunnel_error", tf_core.CoreIpsecConnectionTunnelErrorDataSource())
	RegisterDatasource("oci_core_public_ips", tf_core.CorePublicIpsDataSource())
	RegisterDatasource("oci_core_byoip_range", tf_core.CoreByoipRangeDataSource())
	RegisterDatasource("oci_core_vcn", tf_core.CoreVcnDataSource())
	RegisterDatasource("oci_core_drg_route_distributions", tf_core.CoreDrgRouteDistributionsDataSource())
	RegisterDatasource("oci_core_virtual_circuit_bandwidth_shapes", tf_core.CoreVirtualCircuitBandwidthShapesDataSource())
	RegisterDatasource("oci_core_compute_global_image_capability_schemas", tf_core.CoreComputeGlobalImageCapabilitySchemasDataSource())
	RegisterDatasource("oci_core_cpes", tf_core.CoreCpesDataSource())
	RegisterDatasource("oci_core_instance_console_connections", tf_core.CoreInstanceConsoleConnectionsDataSource())
	RegisterDatasource("oci_core_boot_volume_replicas", tf_core.CoreBootVolumeReplicasDataSource())
	RegisterDatasource("oci_core_drgs", tf_core.CoreDrgsDataSource())
	RegisterDatasource("oci_core_subnets", tf_core.CoreSubnetsDataSource())
	RegisterDatasource("oci_core_dedicated_vm_hosts", tf_core.CoreDedicatedVmHostsDataSource())
	RegisterDatasource("oci_core_instance_pools", tf_core.CoreInstancePoolsDataSource())
	RegisterDatasource("oci_core_volume_groups", tf_core.CoreVolumeGroupsDataSource())
	RegisterDatasource("oci_core_cross_connect_group", tf_core.CoreCrossConnectGroupDataSource())
	RegisterDatasource("oci_core_remote_peering_connections", tf_core.CoreRemotePeeringConnectionsDataSource())
	RegisterDatasource("oci_core_cross_connect_locations", tf_core.CoreCrossConnectLocationsDataSource())
	RegisterDatasource("oci_core_network_security_group", tf_core.CoreNetworkSecurityGroupDataSource())
	RegisterDatasource("oci_core_internet_gateways", tf_core.CoreInternetGatewaysDataSource())
	RegisterDatasource("oci_core_dhcp_options", tf_core.CoreDhcpOptionsDataSource())
	RegisterDatasource("oci_core_drg_route_table_route_rules", tf_core.CoreDrgRouteTableRouteRulesDataSource())
	RegisterDatasource("oci_core_cross_connect_groups", tf_core.CoreCrossConnectGroupsDataSource())
	RegisterDatasource("oci_core_vlan", tf_core.CoreVlanDataSource())
	RegisterDatasource("oci_core_byoip_allocated_ranges", tf_core.CoreByoipAllocatedRangesDataSource())
	RegisterDatasource("oci_core_compute_capacity_reservation_instances", tf_core.CoreComputeCapacityReservationInstancesDataSource())
	RegisterDatasource("oci_core_console_history_data", tf_core.CoreConsoleHistoryContentDataSource())
	RegisterDatasource("oci_core_volumes", tf_core.CoreVolumesDataSource())
	RegisterDatasource("oci_core_fast_connect_provider_services", tf_core.CoreFastConnectProviderServicesDataSource())
	RegisterDatasource("oci_core_local_peering_gateways", tf_core.CoreLocalPeeringGatewaysDataSource())
	RegisterDatasource("oci_core_volume_backups", tf_core.CoreVolumeBackupsDataSource())
	RegisterDatasource("oci_core_network_security_group_vnics", tf_core.CoreNetworkSecurityGroupVnicsDataSource())
	RegisterDatasource("oci_core_vcn_dns_resolver_association", tf_core.CoreVcnDnsResolverAssociationDataSource())
	RegisterDatasource("oci_core_images", tf_core.CoreImagesDataSource())
	RegisterDatasource("oci_core_instance_credentials", tf_core.CoreInstanceCredentialDataSource())
	RegisterDatasource("oci_core_cluster_network_instances", tf_core.CoreClusterNetworkInstancesDataSource())
	RegisterDatasource("oci_core_app_catalog_listing", tf_core.CoreAppCatalogListingDataSource())
	RegisterDatasource("oci_core_compute_image_capability_schemas", tf_core.CoreComputeImageCapabilitySchemasDataSource())
	RegisterDatasource("oci_core_ipv6s", tf_core.CoreIpv6sDataSource())
	RegisterDatasource("oci_core_compute_image_capability_schema", tf_core.CoreComputeImageCapabilitySchemaDataSource())
	RegisterDatasource("oci_core_instance", tf_core.CoreInstanceDataSource())
	RegisterDatasource("oci_core_fast_connect_provider_service", tf_core.CoreFastConnectProviderServiceDataSource())
	RegisterDatasource("oci_core_instance_pool_load_balancer_attachment", tf_core.CoreInstancePoolLoadBalancerAttachmentDataSource())
	RegisterDatasource("oci_core_compute_capacity_reservation", tf_core.CoreComputeCapacityReservationDataSource())
	RegisterDatasource("oci_core_vcns", tf_core.CoreVcnsDataSource())
	RegisterDatasource("oci_core_volume_group_backups", tf_core.CoreVolumeGroupBackupsDataSource())
	RegisterDatasource("oci_core_volume_attachments", tf_core.CoreVolumeAttachmentsDataSource())
	RegisterDatasource("oci_core_drg_route_distribution", tf_core.CoreDrgRouteDistributionDataSource())
	RegisterDatasource("oci_core_instances", tf_core.CoreInstancesDataSource())
	RegisterDatasource("oci_core_public_ip_pool", tf_core.CorePublicIpPoolDataSource())
	RegisterDatasource("oci_core_compute_capacity_reservation_instance_shapes", tf_core.CoreComputeCapacityReservationInstanceShapesDataSource())
	RegisterDatasource("oci_core_public_ip_pools", tf_core.CorePublicIpPoolsDataSource())
	RegisterDatasource("oci_core_image_shape", tf_core.CoreImageShapeDataSource())
	RegisterDatasource("oci_core_ipsec_connections", tf_core.CoreIpSecConnectionsDataSource())
	RegisterDatasource("oci_core_shapes", tf_core.CoreShapesDataSource())
	RegisterDatasource("oci_core_ipsec_status", tf_core.CoreIpSecConnectionDeviceStatusDataSource())
	RegisterDatasource("oci_core_instance_devices", tf_core.CoreInstanceDevicesDataSource())
	RegisterDatasource("oci_core_vnic", tf_core.CoreVnicDataSource())
	RegisterDatasource("oci_core_byoip_ranges", tf_core.CoreByoipRangesDataSource())
	RegisterDatasource("oci_core_vlans", tf_core.CoreVlansDataSource())
	RegisterDatasource("oci_core_virtual_circuits", tf_core.CoreVirtualCircuitsDataSource())
	RegisterDatasource("oci_core_private_ip", tf_core.CorePrivateIpDataSource())
	RegisterDatasource("oci_core_dedicated_vm_host_shapes", tf_core.CoreDedicatedVmHostShapesDataSource())
	RegisterDatasource("oci_core_nat_gateway", tf_core.CoreNatGatewayDataSource())
	RegisterDatasource("oci_core_boot_volume_attachments", tf_core.CoreBootVolumeAttachmentsDataSource())
	RegisterDatasource("oci_core_nat_gateways", tf_core.CoreNatGatewaysDataSource())
	RegisterDatasource("oci_core_compute_global_image_capability_schemas_versions", tf_core.CoreComputeGlobalImageCapabilitySchemasVersionsDataSource())
	RegisterDatasource("oci_core_route_tables", tf_core.CoreRouteTablesDataSource())
	RegisterDatasource("oci_core_boot_volume", tf_core.CoreBootVolumeDataSource())
	RegisterDatasource("oci_core_dedicated_vm_host_instance_shapes", tf_core.CoreDedicatedVmHostInstanceShapesDataSource())
	RegisterDatasource("oci_core_cross_connects", tf_core.CoreCrossConnectsDataSource())
	RegisterDatasource("oci_core_app_catalog_listing_resource_version", tf_core.CoreAppCatalogListingResourceVersionDataSource())
	RegisterDatasource("oci_core_cluster_network", tf_core.CoreClusterNetworkDataSource())
	RegisterDatasource("oci_core_service_gateways", tf_core.CoreServiceGatewaysDataSource())
	RegisterDatasource("oci_core_boot_volume_replica", tf_core.CoreBootVolumeReplicaDataSource())
	RegisterDatasource("oci_core_dedicated_vm_hosts_instances", tf_core.CoreDedicatedVmHostsInstancesDataSource())
	RegisterDatasource("oci_core_ipsec_connection_tunnel_routes", tf_core.CoreIpsecConnectionTunnelRoutesDataSource())
	RegisterDatasource("oci_core_cpe_device_shape", tf_core.CoreCpeDeviceShapeDataSource())
	RegisterDatasource("oci_core_ipsec_config", tf_core.CoreIpSecConnectionDeviceConfigDataSource())
	RegisterDatasource("oci_core_compute_global_image_capability_schema", tf_core.CoreComputeGlobalImageCapabilitySchemaDataSource())
	RegisterDatasource("oci_core_network_security_group_security_rules", tf_core.CoreNetworkSecurityGroupSecurityRulesDataSource())
	RegisterDatasource("oci_core_cpe_device_shapes", tf_core.CoreCpeDeviceShapesDataSource())
	RegisterDatasource("oci_core_ipv6", tf_core.CoreIpv6DataSource())
	RegisterDatasource("oci_core_peer_region_for_remote_peerings", tf_core.CorePeerRegionForRemotePeeringsDataSource())
	RegisterDatasource("oci_core_virtual_circuit_public_prefixes", tf_core.CoreVirtualCircuitPublicPrefixesDataSource())
	RegisterDatasource("oci_core_boot_volume_backup", tf_core.CoreBootVolumeBackupDataSource())
	RegisterDatasource("oci_core_fast_connect_provider_service_key", tf_core.CoreFastConnectProviderServiceKeyDataSource())
	RegisterDatasource("oci_core_instance_measured_boot_report", tf_core.CoreInstanceMeasuredBootReportDataSource())
	RegisterDatasource("oci_core_network_security_groups", tf_core.CoreNetworkSecurityGroupsDataSource())
	RegisterDatasource("oci_core_volume_backup_policy_assignments", tf_core.CoreVolumeBackupPolicyAssignmentsDataSource())
	RegisterDatasource("oci_core_block_volume_replicas", tf_core.CoreBlockVolumeReplicasDataSource())
	RegisterDatasource("oci_core_drg_attachments", tf_core.CoreDrgAttachmentsDataSource())
	RegisterDatasource("oci_core_console_histories", tf_core.CoreConsoleHistoriesDataSource())
	RegisterDatasource("oci_core_ipsec_connection_tunnel", tf_core.CoreIpSecConnectionTunnelDataSource())

}
