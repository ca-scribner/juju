// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package schema

import (
	"github.com/juju/collections/set"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/utils/v4"
	_ "github.com/mattn/go-sqlite3"
	gc "gopkg.in/check.v1"

	jujuversion "github.com/juju/juju/core/version"
)

type modelSchemaSuite struct {
	schemaBaseSuite
}

var _ = gc.Suite(&modelSchemaSuite{})

func (s *modelSchemaSuite) TestModelTables(c *gc.C) {
	s.applyDDL(c, ModelDDL())

	// Ensure that each table is present.
	expected := set.NewStrings(
		// Application
		"application",
		"application_channel",
		"application_config",
		"application_constraint",
		"application_endpoint_space",
		"application_endpoint_cidr",
		"application_platform",
		"application_setting",
		"application_scale",
		"cloud_service",

		// Annotations
		"annotation_application",
		"annotation_charm",
		"annotation_machine",
		"annotation_unit",
		"annotation_model",
		"annotation_storage_instance",
		"annotation_storage_filesystem",
		"annotation_storage_volume",

		// Block commands
		"block_command",
		"block_command_type",

		// Life
		"life",

		// Password
		"password_hash_algorithm",

		// Change log
		"change_log",
		"change_log_edit_type",
		"change_log_namespace",
		"change_log_witness",

		// Model
		"model",

		// Model config
		"model_config",

		// Object store metadata
		"object_store_metadata",
		"object_store_metadata_path",

		// Node
		"net_node",
		"instance_tag",
		"fqdn_address",
		"net_node_fqdn_address",
		"hostname_address",
		"net_node_hostname_address",
		"network_address_scope",

		// Link layer device
		"link_layer_device",
		"link_layer_device_type",
		"virtual_port_type",

		// Network address
		"net_node_ip_address",
		"ip_address_scope",
		"ip_address",
		"ip_address_type",
		"ip_address_origin",
		"ip_address_config_type",
		"ip_address_provider",
		"ip_address_subnet",
		"ip_address_gateway",
		"ip_address_dns_search_domain",
		"ip_address_dns_server_address",

		// Unit
		"unit",
		"unit_resolve_kind",
		"unit_state_charm",
		"unit_state_relation",
		"unit_state",
		"unit_agent",
		"unit_principal",
		"cloud_container",
		"cloud_container_port",
		"unit_agent_status",
		"unit_agent_status_data",
		"unit_workload_status",
		"unit_workload_status_data",
		"cloud_container_status",
		"cloud_container_status_data",
		"unit_agent_status_value",
		"unit_workload_status_value",
		"cloud_container_status_value",

		// Constraint
		"constraint",
		"constraint_tag",
		"constraint_space",
		"constraint_zone",

		// Machine
		"machine",
		"machine_parent",
		"machine_constraint",
		"machine_agent",
		"machine_volume",
		"machine_filesystem",
		"machine_requires_reboot",
		"machine_removals",
		"machine_status",
		"machine_status_data",
		"machine_status_value",
		"machine_cloud_instance",
		"machine_cloud_instance_status_value",
		"machine_cloud_instance_status",
		"machine_cloud_instance_status_data",
		"machine_lxd_profile",

		// Charm
		"architecture",
		"charm_action",
		"charm_category",
		"charm_config_type",
		"charm_config",
		"charm_container_mount",
		"charm_container",
		"charm_device",
		"charm_download_info",
		"charm_extra_binding",
		"charm_hash",
		"charm_local_sequence",
		"charm_manifest_base",
		"charm_metadata",
		"charm_payload",
		"charm_provenance",
		"charm_relation_kind",
		"charm_relation_role",
		"charm_relation_scope",
		"charm_relation",
		"charm_resource_kind",
		"charm_resource",
		"charm_run_as_kind",
		"charm_source",
		"charm_storage_kind",
		"charm_storage_property",
		"charm_storage",
		"charm_tag",
		"charm_term",
		"charm",
		"hash_kind",
		"os",

		// Resources
		"application_resource",
		"kubernetes_application_resource",
		"resource",
		"resource_container_image_metadata_store",
		"resource_file_store",
		"resource_image_store",
		"resource_origin_type",
		"resource_retrieved_by",
		"resource_retrieved_by_type",
		"resource_state",
		"unit_resource",

		// Space
		"space",
		"provider_space",

		// Subnet
		"subnet",
		"provider_subnet",
		"provider_network",
		"provider_network_subnet",
		"availability_zone",
		"availability_zone_subnet",

		// Block device
		"block_device",
		"filesystem_type",
		"block_device_link_device",

		// Storage
		"storage_pool",
		"storage_pool_attribute",
		"storage_kind",
		"storage_instance",
		"storage_unit_owner",
		"storage_attachment",
		"application_storage_directive",
		"unit_storage_directive",
		"storage_volume",
		"storage_instance_volume",
		"storage_volume_attachment",
		"storage_filesystem",
		"storage_instance_filesystem",
		"storage_filesystem_attachment",
		"storage_volume_attachment_plan",
		"storage_volume_attachment_plan_attr",
		"storage_provisioning_status",
		"storage_volume_device_type",

		// Secret
		"secret_rotate_policy",
		"secret",
		"secret_reference",
		"secret_metadata",
		"secret_rotation",
		"secret_value_ref",
		"secret_deleted_value_ref",
		"secret_content",
		"secret_revision",
		"secret_revision_obsolete",
		"secret_revision_expire",
		"secret_application_owner",
		"secret_model_owner",
		"secret_unit_owner",
		"secret_unit_consumer",
		"secret_remote_unit_consumer",
		"secret_permission",
		"secret_role",
		"secret_grant_subject_type",
		"secret_grant_scope_type",

		// Opened Ports
		"protocol",
		"port_range",
	)
	got := readEntityNames(c, s.DB(), "table")
	wanted := expected.Union(internalTableNames)
	c.Assert(got, jc.SameContents, wanted.SortedValues(), gc.Commentf(
		"additive: %v, deletion: %v",
		set.NewStrings(got...).Difference(wanted).SortedValues(),
		wanted.Difference(set.NewStrings(got...)).SortedValues(),
	))
}

func (s *modelSchemaSuite) TestModelViews(c *gc.C) {
	c.Logf("Committing schema DDL")

	s.applyDDL(c, ModelDDL())

	// Ensure that each view is present.
	expected := set.NewStrings(
		"v_address",
		"v_application_charm_download_info",
		"v_application_resource",
		"v_charm_annotation_index",
		"v_charm_config",
		"v_charm_container",
		"v_charm_locator",
		"v_charm_manifest",
		"v_charm_metadata",
		"v_charm_relation",
		"v_charm_resource",
		"v_charm_storage",
		"v_endpoint",
		"v_hardware_characteristics",
		"v_object_store_metadata",
		"v_port_range",
		"v_resource",
		"v_secret_permission",
		"v_space_subnet",
	)
	got := readEntityNames(c, s.DB(), "view")
	c.Assert(got, jc.SameContents, expected.SortedValues(), gc.Commentf(
		"additive: %v, deletion: %v",
		set.NewStrings(got...).Difference(expected).SortedValues(),
		expected.Difference(set.NewStrings(got...)).SortedValues(),
	))
}

func (s *modelSchemaSuite) TestModelTriggers(c *gc.C) {
	s.applyDDL(c, ModelDDL())

	// Expected changelog triggers. Additional triggers are not included and
	// can be added to the addition list.
	expected := set.NewStrings(
		"trg_log_application_delete",
		"trg_log_application_insert",
		"trg_log_application_update",

		"trg_log_application_scale_delete",
		"trg_log_application_scale_insert",
		"trg_log_application_scale_update",

		"trg_log_block_device_delete",
		"trg_log_block_device_insert",
		"trg_log_block_device_update",

		"trg_log_charm_delete",
		"trg_log_charm_insert",
		"trg_log_charm_update",

		"trg_log_machine_cloud_instance_delete",
		"trg_log_machine_cloud_instance_insert",
		"trg_log_machine_cloud_instance_update",

		"trg_log_machine_delete",
		"trg_log_machine_insert",
		"trg_log_machine_update",

		"trg_log_machine_lxd_profile_delete",
		"trg_log_machine_lxd_profile_insert",
		"trg_log_machine_lxd_profile_update",

		"trg_log_machine_requires_reboot_delete",
		"trg_log_machine_requires_reboot_insert",
		"trg_log_machine_requires_reboot_update",

		"trg_log_model_config_delete",
		"trg_log_model_config_insert",
		"trg_log_model_config_update",

		"trg_log_object_store_metadata_path_delete",
		"trg_log_object_store_metadata_path_insert",
		"trg_log_object_store_metadata_path_update",

		"trg_log_port_range_delete",
		"trg_log_port_range_insert",
		"trg_log_port_range_update",

		"trg_log_secret_deleted_value_ref_delete",
		"trg_log_secret_deleted_value_ref_insert",
		"trg_log_secret_deleted_value_ref_update",

		"trg_log_secret_metadata_delete",
		"trg_log_secret_metadata_insert",
		"trg_log_secret_metadata_update",

		"trg_log_secret_reference_delete",
		"trg_log_secret_reference_insert",
		"trg_log_secret_reference_update",

		"trg_log_secret_revision_delete",
		"trg_log_secret_revision_insert",
		"trg_log_secret_revision_update",

		"trg_log_secret_revision_expire_delete",
		"trg_log_secret_revision_expire_insert",
		"trg_log_secret_revision_expire_update",

		"trg_log_secret_revision_obsolete_delete",
		"trg_log_secret_revision_obsolete_insert",
		"trg_log_secret_revision_obsolete_update",

		"trg_log_secret_rotation_delete",
		"trg_log_secret_rotation_insert",
		"trg_log_secret_rotation_update",

		"trg_log_storage_attachment_delete",
		"trg_log_storage_attachment_insert",
		"trg_log_storage_attachment_update",

		"trg_log_storage_filesystem_attachment_delete",
		"trg_log_storage_filesystem_attachment_insert",
		"trg_log_storage_filesystem_attachment_update",

		"trg_log_storage_filesystem_delete",
		"trg_log_storage_filesystem_insert",
		"trg_log_storage_filesystem_update",

		"trg_log_storage_volume_attachment_delete",
		"trg_log_storage_volume_attachment_insert",
		"trg_log_storage_volume_attachment_update",

		"trg_log_storage_volume_attachment_plan_delete",
		"trg_log_storage_volume_attachment_plan_insert",
		"trg_log_storage_volume_attachment_plan_update",

		"trg_log_storage_volume_delete",
		"trg_log_storage_volume_insert",
		"trg_log_storage_volume_update",

		"trg_log_subnet_delete",
		"trg_log_subnet_insert",
		"trg_log_subnet_update",

		"trg_log_unit_delete",
		"trg_log_unit_insert",
		"trg_log_unit_update",
	)

	// These are additional triggers that are not change log triggers, but
	// will be present in the schema.
	additional := set.NewStrings(
		"trg_model_immutable_delete",
		"trg_model_immutable_update",

		"trg_secret_permission_guard_update",
		"trg_charm_local_sequence_guard_update",
	)

	got := readEntityNames(c, s.DB(), "trigger")
	wanted := expected.Union(additional)
	c.Assert(got, jc.SameContents, wanted.SortedValues(), gc.Commentf(
		"additive: %v, deletion: %v",
		set.NewStrings(got...).Difference(wanted).SortedValues(),
		wanted.Difference(set.NewStrings(got...)).SortedValues(),
	))
}

func (s *modelSchemaSuite) TestModelTriggersForImmutableTables(c *gc.C) {
	s.applyDDL(c, ModelDDL())

	modelUUID := utils.MustNewUUID().String()
	controllerUUID := utils.MustNewUUID().String()
	s.assertExecSQL(c, `
INSERT INTO model (uuid, controller_uuid, target_agent_version, name, type, cloud, cloud_type, cloud_region)
VALUES (?, ?, ?, 'my-model', 'caas', 'cloud-1', 'kubernetes', 'cloud-region-1');`,
		modelUUID, controllerUUID, jujuversion.Current.String())
	s.assertExecSQLError(c,
		"UPDATE model SET name = 'new-name' WHERE uuid = ?",
		"model table is immutable", modelUUID)

	s.assertExecSQLError(c,
		"DELETE FROM model WHERE uuid = ?;",
		"model table is immutable", modelUUID)
}
