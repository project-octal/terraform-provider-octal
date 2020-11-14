package octal

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/project-octal/linkderd-cli/cmd"
)

func translateInstallOptions(options cmd.InstallOptions, d *schema.ResourceData) {
	_ = d.Set("cluster_domain", options.ClusterDomain)
	_ = d.Set("add_on_config", options.AddOnConfig)
}

func linkderd2Schema() map[string]*schema.Schema {
	thisSchema := map[string]*schema.Schema{
		"cluster_domain": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"add_on_config": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"control_plane_version": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"controller_replicas": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"controller_log_level": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"prometheus_image": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"high_availability": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"controller_uid": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"disable_h2_upgrade": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"disable_heartbeat": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"cni_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"skip_checks": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"omit_webhook_side_effects": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"restrict_dashboard_privileges": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"control_plane_tracing": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"smi_metrics_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"smi_metrics_image": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"identity_options": &schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"replicas": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"trust_domain": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"issuance_lifetime": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"clock_skew_allowance": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"trust_pem_file": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"crt_pem_file": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"key_pem_file": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"identity_external_issuer": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		"proxy_config_options": &schema.Schema{
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"proxy_version": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"proxy_image": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"init_image": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"init_image_version": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"debug_image": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"debug_image_version": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"docker_registry": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"image_pull_policy": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"destination_get_networks": {
						Type:     schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Optional: true,
					},
					"ignore_inbound_ports": {
						Type:     schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Optional: true,
					},
					"ignore_outbound_ports": {
						Type:     schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Optional: true,
					},
					"proxy_uid": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"proxy_log_level": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"proxy_inbound_port": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"proxy_outbound_port": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"proxy_control_port": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"proxy_admin_port": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"proxy_cpu_request": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"proxy_memory_request": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"proxy_cpu_limit": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"proxy_memory_limit": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"enable_external_profiles": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"trace_collector": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"trace_collector_svc_account": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"wait_before_exit_seconds": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"ignore_cluster": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"disable_identity": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"require_identity_on_inbound_ports": {
						Type:     schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Optional: true,
					},
					"disable_tap": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
	}
	return thisSchema
}