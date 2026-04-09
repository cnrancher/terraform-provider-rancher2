package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	clusterTKEV2Kind   = "tkeV2"
	clusterDriverTKEV2 = "TKE"
)

func clusterTKEDataDiskFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"disk_size": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Disk size in GB",
		},
		"disk_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Disk type",
		},
	}
}

func clusterTKERunInstancesForNodeFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"image_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Node image ID",
		},
		"instance_charge_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Instance charge type",
		},
		"instance_count": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Number of instances",
		},
		"instance_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Instance name",
		},
		"instance_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Instance type",
		},
		"internet_charge_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Internet charge type",
		},
		"internet_max_bandwidth_out": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Max internet egress bandwidth",
		},
		"key_ids": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "SSH key IDs",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"monitor_service": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Enable monitor service",
		},
		"node_role": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Node role",
		},
		"project_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Project ID",
		},
		"public_ip_assigned": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Assign public IP",
		},
		"security_service": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Enable security service",
		},
		"subnet_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Subnet ID",
		},
		"system_disk": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "System disk",
			Elem:        &schema.Resource{Schema: clusterTKEDataDiskFields()},
		},
		"user_data": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Cloud-init user data",
		},
		"vpc_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "VPC ID",
		},
		"zone": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Availability zone",
		},
	}
}

func clusterTKENodePoolAutoScalingGroupParaFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auto_scaling_group_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Auto scaling group name",
		},
		"desired_capacity": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Desired capacity",
		},
		"max_size": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Max size",
		},
		"min_size": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Min size",
		},
		"subnet_ids": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Subnet IDs",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"vpc_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "VPC ID",
		},
	}
}

func clusterTKENodePoolLaunchConfigureParaFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"data_disks": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Data disks",
			Elem:        &schema.Resource{Schema: clusterTKEDataDiskFields()},
		},
		"instance_charge_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Instance charge type",
		},
		"instance_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Instance type",
		},
		"internet_charge_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Internet charge type",
		},
		"internet_max_bandwidth_out": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Max internet egress bandwidth",
		},
		"key_ids": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "SSH key IDs",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"launch_configuration_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Launch configuration name",
		},
		"public_ip_assigned": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Assign public IP",
		},
		"security_group_ids": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Security group IDs",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"system_disk": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "System disk",
			Elem:        &schema.Resource{Schema: clusterTKEDataDiskFields()},
		},
	}
}

func clusterTKENodePoolDetailFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Node pool name",
		},
		"node_pool_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Node pool ID",
		},
		"cluster_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Cluster ID",
		},
		"node_pool_os": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Node pool OS",
		},
		"os_customize_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "OS customize type",
		},
		"deletion_protection": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Enable deletion protection",
		},
		"enable_autoscale": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Enable autoscaling",
		},
		"labels": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Labels",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"tags": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Tags",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"taints": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Taints",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"user_script": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Custom user script for node pool instances",
		},
		"auto_scaling_group_para": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "Auto scaling group parameters",
			Elem:        &schema.Resource{Schema: clusterTKENodePoolAutoScalingGroupParaFields()},
		},
		"launch_configure_para": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "Launch configuration parameters",
			Elem:        &schema.Resource{Schema: clusterTKENodePoolLaunchConfigureParaFields()},
		},
	}
}

func clusterTKEVirtualNodeLabelFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Virtual node label name",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Virtual node label value",
		},
	}
}

func clusterTKEVirtualNodeTaintFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"effect": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Virtual node taint effect",
		},
		"key": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Virtual node taint key",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Virtual node taint value",
		},
	}
}

func clusterTKEVirtualNodeTagFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Virtual node tag key",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Virtual node tag value",
		},
	}
}

func clusterTKEVirtualNodeSpecFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Virtual node display name",
		},
		"subnet_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Virtual node subnet ID",
		},
		"tags": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Virtual node tags",
			Elem:        &schema.Resource{Schema: clusterTKEVirtualNodeTagFields()},
		},
	}
}

func clusterTKEVirtualNodePoolDetailFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"deletion_protection": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable virtual node pool deletion protection",
		},
		"labels": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Virtual node pool labels",
			Elem:        &schema.Resource{Schema: clusterTKEVirtualNodeLabelFields()},
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Virtual node pool name",
		},
		"node_pool_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Virtual node pool ID",
		},
		"os": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Virtual node pool OS",
		},
		"security_group_ids": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Virtual node pool security group IDs",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"subnet_ids": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Virtual node pool subnet IDs",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"taints": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Virtual node pool taints",
			Elem:        &schema.Resource{Schema: clusterTKEVirtualNodeTaintFields()},
		},
		"virtual_nodes": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Virtual nodes",
			Elem:        &schema.Resource{Schema: clusterTKEVirtualNodeSpecFields()},
		},
	}
}

func clusterTKEClusterBasicSettingsFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Cluster description",
		},
		"cluster_level": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Cluster level",
		},
		"cluster_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Cluster name",
		},
		"cluster_os": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Cluster OS",
		},
		"cluster_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Cluster type",
		},
		"cluster_version": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Cluster version",
		},
		"is_auto_upgrade": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Enable auto upgrade",
		},
		"project_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Project ID",
		},
		"tags": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Tags",
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"vpc_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "VPC ID",
		},
	}
}

func clusterTKEClusterAdvancedSettingsFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"as_enabled":                 {Type: schema.TypeBool, Optional: true},
		"audit_enabled":              {Type: schema.TypeBool, Optional: true},
		"audit_log_topic_id":         {Type: schema.TypeString, Optional: true},
		"audit_logset_id":            {Type: schema.TypeString, Optional: true},
		"base_pod_number":            {Type: schema.TypeInt, Optional: true},
		"cilium_mode":                {Type: schema.TypeString, Optional: true},
		"container_runtime":          {Type: schema.TypeString, Optional: true},
		"deletion_protection":        {Type: schema.TypeBool, Optional: true},
		"enable_customized_pod_cidr": {Type: schema.TypeBool, Optional: true},
		"etcd": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ipvs":                  {Type: schema.TypeBool, Optional: true},
		"is_dual_stack":         {Type: schema.TypeBool, Optional: true},
		"is_non_static_ip_mode": {Type: schema.TypeBool, Optional: true},
		"kube_api_server": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"kube_controller_manager": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"kube_proxy_mode": {Type: schema.TypeString, Optional: true},
		"kube_scheduler": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"network_type":      {Type: schema.TypeString, Optional: true},
		"node_name_type":    {Type: schema.TypeString, Optional: true},
		"qgpu_share_enable": {Type: schema.TypeBool, Optional: true},
		"runtime_version":   {Type: schema.TypeString, Optional: true, Computed: true},
		"vpc_cni_type":      {Type: schema.TypeString, Optional: true},
	}
}

func clusterTKEClusterCIDRSettingsFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"claim_expired_seconds": {Type: schema.TypeInt, Optional: true},
		"cluster_cidr":          {Type: schema.TypeString, Optional: true},
		"eni_subnet_ids": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ignore_cluster_cidr_conflict": {Type: schema.TypeBool, Optional: true},
		"ignore_service_cidr_conflict": {Type: schema.TypeBool, Optional: true},
		"max_cluster_service_num":      {Type: schema.TypeInt, Optional: true},
		"max_node_pod_num":             {Type: schema.TypeInt, Optional: true},
		"os_customize_type":            {Type: schema.TypeString, Optional: true},
		"service_cidr":                 {Type: schema.TypeString, Optional: true, Computed: true},
		"subnet_id":                    {Type: schema.TypeString, Optional: true},
	}
}

func clusterTKEClusterEndpointFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"domain":               {Type: schema.TypeString, Optional: true},
		"enable":               {Type: schema.TypeBool, Optional: true, Computed: true},
		"extensive_parameters": {Type: schema.TypeString, Optional: true},
		"security_group":       {Type: schema.TypeString, Optional: true},
		"subnet_id":            {Type: schema.TypeString, Optional: true},
	}
}

func clusterTKEExtensionAddonFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"addon_name":  {Type: schema.TypeString, Required: true},
		"addon_param": {Type: schema.TypeString, Optional: true},
	}
}

func clusterTKEConfigV2Fields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"imported": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether this is an imported TKE cluster",
		},
		"cluster_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "Existing TKE cluster ID. Required for imported cluster",
		},
		"region": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "TKE region",
		},
		"tke_credential_secret": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Rancher cloud credential secret ID for TKE",
		},
		"cluster_basic_settings": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "TKE cluster basic settings",
			Elem:        &schema.Resource{Schema: clusterTKEClusterBasicSettingsFields()},
		},
		"cluster_advanced_settings": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "TKE cluster advanced settings",
			Elem:        &schema.Resource{Schema: clusterTKEClusterAdvancedSettingsFields()},
		},
		"cluster_cidr_settings": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "TKE cluster CIDR settings",
			Elem:        &schema.Resource{Schema: clusterTKEClusterCIDRSettingsFields()},
		},
		"cluster_endpoint": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "TKE cluster endpoint settings",
			Elem:        &schema.Resource{Schema: clusterTKEClusterEndpointFields()},
		},
		"extension_addon": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "TKE extension addons",
			Elem:        &schema.Resource{Schema: clusterTKEExtensionAddonFields()},
		},
		"node_pool_list": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "TKE node pool list",
			Elem:        &schema.Resource{Schema: clusterTKENodePoolDetailFields()},
		},
		"virtual_node_pool_list": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "TKE virtual node pool list",
			Elem:        &schema.Resource{Schema: clusterTKEVirtualNodePoolDetailFields()},
		},
		"run_instances_for_node": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "RunInstances parameters for initial node creation",
			Elem:        &schema.Resource{Schema: clusterTKERunInstancesForNodeFields()},
		},
	}
}
