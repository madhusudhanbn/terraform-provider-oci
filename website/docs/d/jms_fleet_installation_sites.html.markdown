---
subcategory: "Jms"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_jms_fleet_installation_sites"
sidebar_current: "docs-oci-datasource-jms-fleet_installation_sites"
description: |-
  Provides the list of Fleet Installation Sites in Oracle Cloud Infrastructure Jms service
---

# Data Source: oci_jms_fleet_installation_sites
This data source provides the list of Fleet Installation Sites in Oracle Cloud Infrastructure Jms service.

List Java installation sites in a Fleet filtered by query parameters.

## Example Usage

```hcl
data "oci_jms_fleet_installation_sites" "test_fleet_installation_sites" {
	#Required
	fleet_id = oci_jms_fleet.test_fleet.id

	#Optional
	application_id = oci_dataflow_application.test_application.id
	installation_path = var.fleet_installation_site_installation_path
	jre_distribution = var.fleet_installation_site_jre_distribution
	jre_security_status = var.fleet_installation_site_jre_security_status
	jre_vendor = var.fleet_installation_site_jre_vendor
	jre_version = var.fleet_installation_site_jre_version
	managed_instance_id = oci_osmanagement_managed_instance.test_managed_instance.id
	os_family = var.fleet_installation_site_os_family
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Optional) The Fleet-unique identifier of the related application.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `installation_path` - (Optional) The file system path of the installation.
* `jre_distribution` - (Optional) The distribution of the related Java Runtime.
* `jre_security_status` - (Optional) The security status of the Java Runtime.
* `jre_vendor` - (Optional) The vendor of the related Java Runtime.
* `jre_version` - (Optional) The version of the related Java Runtime.
* `managed_instance_id` - (Optional) The Fleet-unique identifier of the related managed instance.
* `os_family` - (Optional) The operating system type.


## Attributes Reference

The following attributes are exported:

* `installation_site_collection` - The list of installation_site_collection.

### FleetInstallationSite Reference

The following attributes are exported:

* `items` - A list of Java installation sites.
	* `approximate_application_count` - The approximate count of applications running on this installation
	* `blocklist` - The list of operations that are blocklisted.
		* `operation` - The operation type.
		* `reason` - The reason why the operation is blocklisted.
	* `installation_key` - The unique identifier for the installation of Java Runtime at a specific path on a specific operating system.
	* `jre` - The essential properties to identify a Java Runtime.
		* `distribution` - The distribution of a Java Runtime is the name of the lineage of product to which it belongs, for example _Java(TM) SE Runtime Environment_.
		* `jre_key` - The unique identifier for a Java Runtime.
		* `vendor` - The vendor of the Java Runtime.
		* `version` - The version of the Java Runtime.
	* `managed_instance_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related managed instance. 
	* `operating_system` - Operating System of the platform on which the Java Runtime was reported. 
		* `architecture` - The architecture of the operating system as provided by the Java system property os.arch.
		* `family` - The operating system type, such as Windows or Linux
		* `name` - The name of the operating system as provided by the Java system property os.name.
		* `version` - The version of the operating system as provided by the Java system property os.version.
	* `path` - The file system path of the installation.
	* `security_status` - The security status of the Java Runtime.
	* `state` - The lifecycle state of the installation site.
	* `time_last_seen` - The date and time the resource was _last_ reported to JMS. This is potentially _after_ the specified time period provided by the filters. For example, a resource can be last reported to JMS before the start of a specified time period, if it is also reported during the time period. 

