package rancher2

import (
	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
)

func flattenClusterTKEDataDisk(in *managementClient.DataDisk) []interface{} {
	if in == nil {
		return nil
	}
	obj := map[string]interface{}{}
	if in.DiskSize > 0 {
		obj["disk_size"] = int(in.DiskSize)
	}
	if len(in.DiskType) > 0 {
		obj["disk_type"] = in.DiskType
	}
	return []interface{}{obj}
}

func expandClusterTKEDataDisk(p []interface{}) *managementClient.DataDisk {
	if len(p) == 0 || p[0] == nil {
		return nil
	}
	in := p[0].(map[string]interface{})
	obj := &managementClient.DataDisk{}
	if v, ok := in["disk_size"].(int); ok && v > 0 {
		obj.DiskSize = int64(v)
	}
	if v, ok := in["disk_type"].(string); ok && len(v) > 0 {
		obj.DiskType = v
	}
	return obj
}

func flattenClusterTKERunInstancesForNode(in *managementClient.RunInstancesForNode) []interface{} {
	if in == nil {
		return nil
	}
	obj := map[string]interface{}{}
	if len(in.ImageID) > 0 {
		obj["image_id"] = in.ImageID
	}
	if len(in.InstanceChargeType) > 0 {
		obj["instance_charge_type"] = in.InstanceChargeType
	}
	if in.InstanceCount > 0 {
		obj["instance_count"] = int(in.InstanceCount)
	}
	if len(in.InstanceName) > 0 {
		obj["instance_name"] = in.InstanceName
	}
	if len(in.InstanceType) > 0 {
		obj["instance_type"] = in.InstanceType
	}
	if len(in.InternetChargeType) > 0 {
		obj["internet_charge_type"] = in.InternetChargeType
	}
	if in.InternetMaxBandwidthOut > 0 {
		obj["internet_max_bandwidth_out"] = int(in.InternetMaxBandwidthOut)
	}
	if len(in.KeyIDs) > 0 {
		obj["key_ids"] = toArrayInterfaceSorted(in.KeyIDs)
	}
	obj["monitor_service"] = in.MonitorService
	if len(in.NodeRole) > 0 {
		obj["node_role"] = in.NodeRole
	}
	if in.ProjectID > 0 {
		obj["project_id"] = int(in.ProjectID)
	}
	obj["public_ip_assigned"] = in.PublicIpAssigned
	obj["security_service"] = in.SecurityService
	if len(in.SubnetID) > 0 {
		obj["subnet_id"] = in.SubnetID
	}
	if in.SystemDisk != nil {
		obj["system_disk"] = flattenClusterTKEDataDisk(in.SystemDisk)
	}
	if len(in.UserData) > 0 {
		obj["user_data"] = in.UserData
	}
	if len(in.VpcID) > 0 {
		obj["vpc_id"] = in.VpcID
	}
	if len(in.Zone) > 0 {
		obj["zone"] = in.Zone
	}
	return []interface{}{obj}
}

func expandClusterTKERunInstancesForNode(p []interface{}) *managementClient.RunInstancesForNode {
	if len(p) == 0 || p[0] == nil {
		return nil
	}
	in := p[0].(map[string]interface{})
	obj := &managementClient.RunInstancesForNode{}
	if v, ok := in["image_id"].(string); ok && len(v) > 0 {
		obj.ImageID = v
	}
	if v, ok := in["instance_charge_type"].(string); ok && len(v) > 0 {
		obj.InstanceChargeType = v
	}
	if v, ok := in["instance_count"].(int); ok && v > 0 {
		obj.InstanceCount = int64(v)
	}
	if v, ok := in["instance_name"].(string); ok && len(v) > 0 {
		obj.InstanceName = v
	}
	if v, ok := in["instance_type"].(string); ok && len(v) > 0 {
		obj.InstanceType = v
	}
	if v, ok := in["internet_charge_type"].(string); ok && len(v) > 0 {
		obj.InternetChargeType = v
	}
	if v, ok := in["internet_max_bandwidth_out"].(int); ok && v > 0 {
		obj.InternetMaxBandwidthOut = int64(v)
	}
	if v, ok := in["key_ids"].([]interface{}); ok {
		obj.KeyIDs = toArrayStringSorted(v)
	}
	if v, ok := in["monitor_service"].(bool); ok {
		obj.MonitorService = v
	}
	if v, ok := in["node_role"].(string); ok && len(v) > 0 {
		obj.NodeRole = v
	}
	if v, ok := in["project_id"].(int); ok && v > 0 {
		obj.ProjectID = int64(v)
	}
	if v, ok := in["public_ip_assigned"].(bool); ok {
		obj.PublicIpAssigned = v
	}
	if v, ok := in["security_service"].(bool); ok {
		obj.SecurityService = v
	}
	if v, ok := in["subnet_id"].(string); ok && len(v) > 0 {
		obj.SubnetID = v
	}
	if v, ok := in["system_disk"].([]interface{}); ok && len(v) > 0 {
		obj.SystemDisk = expandClusterTKEDataDisk(v)
	}
	if v, ok := in["user_data"].(string); ok && len(v) > 0 {
		obj.UserData = v
	}
	if v, ok := in["vpc_id"].(string); ok && len(v) > 0 {
		obj.VpcID = v
	}
	if v, ok := in["zone"].(string); ok && len(v) > 0 {
		obj.Zone = v
	}
	return obj
}

func flattenClusterTKENodePoolAutoScalingGroupPara(in *managementClient.AutoScalingGroupPara) []interface{} {
	if in == nil {
		return nil
	}
	obj := map[string]interface{}{}
	if len(in.AutoScalingGroupName) > 0 {
		obj["auto_scaling_group_name"] = in.AutoScalingGroupName
	}
	if in.DesiredCapacity > 0 {
		obj["desired_capacity"] = int(in.DesiredCapacity)
	}
	if in.MaxSize > 0 {
		obj["max_size"] = int(in.MaxSize)
	}
	if in.MinSize > 0 {
		obj["min_size"] = int(in.MinSize)
	}
	if len(in.SubnetIDs) > 0 {
		obj["subnet_ids"] = toArrayInterfaceSorted(in.SubnetIDs)
	}
	if len(in.VpcID) > 0 {
		obj["vpc_id"] = in.VpcID
	}
	return []interface{}{obj}
}

func expandClusterTKENodePoolAutoScalingGroupPara(p []interface{}) *managementClient.AutoScalingGroupPara {
	if len(p) == 0 || p[0] == nil {
		return nil
	}
	in := p[0].(map[string]interface{})
	obj := &managementClient.AutoScalingGroupPara{}
	if v, ok := in["auto_scaling_group_name"].(string); ok && len(v) > 0 {
		obj.AutoScalingGroupName = v
	}
	if v, ok := in["desired_capacity"].(int); ok && v > 0 {
		obj.DesiredCapacity = int64(v)
	}
	if v, ok := in["max_size"].(int); ok && v > 0 {
		obj.MaxSize = int64(v)
	}
	if v, ok := in["min_size"].(int); ok && v > 0 {
		obj.MinSize = int64(v)
	}
	if v, ok := in["subnet_ids"].([]interface{}); ok {
		obj.SubnetIDs = toArrayStringSorted(v)
	}
	if v, ok := in["vpc_id"].(string); ok && len(v) > 0 {
		obj.VpcID = v
	}
	return obj
}

func flattenClusterTKENodePoolLaunchConfigurePara(in *managementClient.LaunchConfigurePara) []interface{} {
	if in == nil {
		return nil
	}
	obj := map[string]interface{}{}
	if len(in.DataDisks) > 0 {
		out := make([]interface{}, len(in.DataDisks))
		for i := range in.DataDisks {
			dd := in.DataDisks[i]
			item := map[string]interface{}{}
			if dd.DiskSize > 0 {
				item["disk_size"] = int(dd.DiskSize)
			}
			if len(dd.DiskType) > 0 {
				item["disk_type"] = dd.DiskType
			}
			out[i] = item
		}
		obj["data_disks"] = out
	}
	if len(in.InstanceChargeType) > 0 {
		obj["instance_charge_type"] = in.InstanceChargeType
	}
	if len(in.InstanceType) > 0 {
		obj["instance_type"] = in.InstanceType
	}
	if len(in.InternetChargeType) > 0 {
		obj["internet_charge_type"] = in.InternetChargeType
	}
	if in.InternetMaxBandwidthOut > 0 {
		obj["internet_max_bandwidth_out"] = int(in.InternetMaxBandwidthOut)
	}
	if len(in.KeyIDs) > 0 {
		obj["key_ids"] = toArrayInterfaceSorted(in.KeyIDs)
	}
	if len(in.LaunchConfigurationName) > 0 {
		obj["launch_configuration_name"] = in.LaunchConfigurationName
	}
	obj["public_ip_assigned"] = in.PublicIpAssigned
	if len(in.SecurityGroupIDs) > 0 {
		obj["security_group_ids"] = toArrayInterfaceSorted(in.SecurityGroupIDs)
	}
	if in.SystemDisk != nil {
		obj["system_disk"] = flattenClusterTKEDataDisk(in.SystemDisk)
	}
	return []interface{}{obj}
}

func expandClusterTKENodePoolLaunchConfigurePara(p []interface{}) *managementClient.LaunchConfigurePara {
	if len(p) == 0 || p[0] == nil {
		return nil
	}
	in := p[0].(map[string]interface{})
	obj := &managementClient.LaunchConfigurePara{}
	if v, ok := in["data_disks"].([]interface{}); ok && len(v) > 0 {
		obj.DataDisks = make([]managementClient.DataDisk, len(v))
		for i := range v {
			item := v[i].(map[string]interface{})
			if s, ok := item["disk_size"].(int); ok && s > 0 {
				obj.DataDisks[i].DiskSize = int64(s)
			}
			if t, ok := item["disk_type"].(string); ok && len(t) > 0 {
				obj.DataDisks[i].DiskType = t
			}
		}
	}
	if v, ok := in["instance_charge_type"].(string); ok && len(v) > 0 {
		obj.InstanceChargeType = v
	}
	if v, ok := in["instance_type"].(string); ok && len(v) > 0 {
		obj.InstanceType = v
	}
	if v, ok := in["internet_charge_type"].(string); ok && len(v) > 0 {
		obj.InternetChargeType = v
	}
	if v, ok := in["internet_max_bandwidth_out"].(int); ok && v > 0 {
		obj.InternetMaxBandwidthOut = int64(v)
	}
	if v, ok := in["key_ids"].([]interface{}); ok {
		obj.KeyIDs = toArrayStringSorted(v)
	}
	if v, ok := in["launch_configuration_name"].(string); ok && len(v) > 0 {
		obj.LaunchConfigurationName = v
	}
	if v, ok := in["public_ip_assigned"].(bool); ok {
		obj.PublicIpAssigned = v
	}
	if v, ok := in["security_group_ids"].([]interface{}); ok {
		obj.SecurityGroupIDs = toArrayStringSorted(v)
	}
	if v, ok := in["system_disk"].([]interface{}); ok && len(v) > 0 {
		obj.SystemDisk = expandClusterTKEDataDisk(v)
	}
	return obj
}

func flattenClusterTKENodePoolDetail(p []managementClient.NodePoolDetail) []interface{} {
	if len(p) == 0 {
		return nil
	}
	out := make([]interface{}, len(p))
	for i, in := range p {
		obj := map[string]interface{}{}
		if len(in.Name) > 0 {
			obj["name"] = in.Name
		}
		if len(in.NodePoolID) > 0 {
			obj["node_pool_id"] = in.NodePoolID
		}
		if len(in.ClusterID) > 0 {
			obj["cluster_id"] = in.ClusterID
		}
		if len(in.NodePoolOs) > 0 {
			obj["node_pool_os"] = in.NodePoolOs
		}
		if len(in.OsCustomizeType) > 0 {
			obj["os_customize_type"] = in.OsCustomizeType
		}
		obj["deletion_protection"] = in.DeletionProtection
		obj["enable_autoscale"] = in.EnableAutoscale
		if len(in.Labels) > 0 {
			obj["labels"] = toArrayInterfaceSorted(in.Labels)
		}
		if len(in.Tags) > 0 {
			obj["tags"] = toArrayInterfaceSorted(in.Tags)
		}
		if len(in.Taints) > 0 {
			obj["taints"] = toArrayInterfaceSorted(in.Taints)
		}
		if len(in.UserScript) > 0 {
			obj["user_script"] = in.UserScript
		}
		if in.AutoScalingGroupPara != nil {
			obj["auto_scaling_group_para"] = flattenClusterTKENodePoolAutoScalingGroupPara(in.AutoScalingGroupPara)
		}
		if in.LaunchConfigurePara != nil {
			obj["launch_configure_para"] = flattenClusterTKENodePoolLaunchConfigurePara(in.LaunchConfigurePara)
		}
		out[i] = obj
	}
	return out
}

func expandClusterTKENodePoolDetail(p []interface{}) []managementClient.NodePoolDetail {
	if len(p) == 0 {
		return nil
	}
	out := make([]managementClient.NodePoolDetail, len(p))
	for i := range p {
		in := p[i].(map[string]interface{})
		obj := managementClient.NodePoolDetail{}
		if v, ok := in["name"].(string); ok && len(v) > 0 {
			obj.Name = v
		}
		if v, ok := in["node_pool_id"].(string); ok && len(v) > 0 {
			obj.NodePoolID = v
		}
		if v, ok := in["cluster_id"].(string); ok && len(v) > 0 {
			obj.ClusterID = v
		}
		if v, ok := in["node_pool_os"].(string); ok && len(v) > 0 {
			obj.NodePoolOs = v
		}
		if v, ok := in["os_customize_type"].(string); ok && len(v) > 0 {
			obj.OsCustomizeType = v
		}
		if v, ok := in["deletion_protection"].(bool); ok {
			obj.DeletionProtection = v
		}
		if v, ok := in["enable_autoscale"].(bool); ok {
			obj.EnableAutoscale = v
		}
		if v, ok := in["labels"].([]interface{}); ok {
			obj.Labels = toArrayStringSorted(v)
		}
		if v, ok := in["tags"].([]interface{}); ok {
			obj.Tags = toArrayStringSorted(v)
		}
		if v, ok := in["taints"].([]interface{}); ok {
			obj.Taints = toArrayStringSorted(v)
		}
		if v, ok := in["user_script"].(string); ok && len(v) > 0 {
			obj.UserScript = v
		}
		if v, ok := in["auto_scaling_group_para"].([]interface{}); ok && len(v) > 0 {
			obj.AutoScalingGroupPara = expandClusterTKENodePoolAutoScalingGroupPara(v)
		}
		if v, ok := in["launch_configure_para"].([]interface{}); ok && len(v) > 0 {
			obj.LaunchConfigurePara = expandClusterTKENodePoolLaunchConfigurePara(v)
		}
		out[i] = obj
	}
	return out
}

func flattenClusterTKEVirtualNodeLabel(p []managementClient.VirtualNodeLabel) []interface{} {
	if len(p) == 0 {
		return nil
	}
	out := make([]interface{}, len(p))
	for i := range p {
		obj := map[string]interface{}{}
		if len(p[i].Name) > 0 {
			obj["name"] = p[i].Name
		}
		if len(p[i].Value) > 0 {
			obj["value"] = p[i].Value
		}
		out[i] = obj
	}
	return out
}

func expandClusterTKEVirtualNodeLabel(p []interface{}) []managementClient.VirtualNodeLabel {
	if len(p) == 0 {
		return nil
	}
	out := make([]managementClient.VirtualNodeLabel, len(p))
	for i := range p {
		in := p[i].(map[string]interface{})
		if v, ok := in["name"].(string); ok && len(v) > 0 {
			out[i].Name = v
		}
		if v, ok := in["value"].(string); ok && len(v) > 0 {
			out[i].Value = v
		}
	}
	return out
}

func flattenClusterTKEVirtualNodeTaint(p []managementClient.VirtualNodeTaint) []interface{} {
	if len(p) == 0 {
		return nil
	}
	out := make([]interface{}, len(p))
	for i := range p {
		obj := map[string]interface{}{}
		if len(p[i].Effect) > 0 {
			obj["effect"] = p[i].Effect
		}
		if len(p[i].Key) > 0 {
			obj["key"] = p[i].Key
		}
		if len(p[i].Value) > 0 {
			obj["value"] = p[i].Value
		}
		out[i] = obj
	}
	return out
}

func expandClusterTKEVirtualNodeTaint(p []interface{}) []managementClient.VirtualNodeTaint {
	if len(p) == 0 {
		return nil
	}
	out := make([]managementClient.VirtualNodeTaint, len(p))
	for i := range p {
		in := p[i].(map[string]interface{})
		if v, ok := in["effect"].(string); ok && len(v) > 0 {
			out[i].Effect = v
		}
		if v, ok := in["key"].(string); ok && len(v) > 0 {
			out[i].Key = v
		}
		if v, ok := in["value"].(string); ok && len(v) > 0 {
			out[i].Value = v
		}
	}
	return out
}

func flattenClusterTKEVirtualNodeTag(p []managementClient.VirtualNodeTag) []interface{} {
	if len(p) == 0 {
		return nil
	}
	out := make([]interface{}, len(p))
	for i := range p {
		obj := map[string]interface{}{}
		if len(p[i].Key) > 0 {
			obj["key"] = p[i].Key
		}
		if len(p[i].Value) > 0 {
			obj["value"] = p[i].Value
		}
		out[i] = obj
	}
	return out
}

func expandClusterTKEVirtualNodeTag(p []interface{}) []managementClient.VirtualNodeTag {
	if len(p) == 0 {
		return nil
	}
	out := make([]managementClient.VirtualNodeTag, len(p))
	for i := range p {
		in := p[i].(map[string]interface{})
		if v, ok := in["key"].(string); ok && len(v) > 0 {
			out[i].Key = v
		}
		if v, ok := in["value"].(string); ok && len(v) > 0 {
			out[i].Value = v
		}
	}
	return out
}

func flattenClusterTKEVirtualNodeSpec(p []managementClient.VirtualNodeSpec) []interface{} {
	if len(p) == 0 {
		return nil
	}
	out := make([]interface{}, len(p))
	for i := range p {
		obj := map[string]interface{}{}
		if len(p[i].DisplayName) > 0 {
			obj["display_name"] = p[i].DisplayName
		}
		if len(p[i].SubnetId) > 0 {
			obj["subnet_id"] = p[i].SubnetId
		}
		if len(p[i].Tags) > 0 {
			obj["tags"] = flattenClusterTKEVirtualNodeTag(p[i].Tags)
		}
		out[i] = obj
	}
	return out
}

func expandClusterTKEVirtualNodeSpec(p []interface{}) []managementClient.VirtualNodeSpec {
	if len(p) == 0 {
		return nil
	}
	out := make([]managementClient.VirtualNodeSpec, len(p))
	for i := range p {
		in := p[i].(map[string]interface{})
		if v, ok := in["display_name"].(string); ok && len(v) > 0 {
			out[i].DisplayName = v
		}
		if v, ok := in["subnet_id"].(string); ok && len(v) > 0 {
			out[i].SubnetId = v
		}
		if v, ok := in["tags"].([]interface{}); ok && len(v) > 0 {
			out[i].Tags = expandClusterTKEVirtualNodeTag(v)
		}
	}
	return out
}

func flattenClusterTKEVirtualNodePoolDetail(p []managementClient.VirtualNodePoolDetail) []interface{} {
	if len(p) == 0 {
		return nil
	}
	out := make([]interface{}, len(p))
	for i := range p {
		obj := map[string]interface{}{}
		if p[i].DeletionProtection != nil {
			obj["deletion_protection"] = *p[i].DeletionProtection
		}
		if len(p[i].Labels) > 0 {
			obj["labels"] = flattenClusterTKEVirtualNodeLabel(p[i].Labels)
		}
		if len(p[i].Name) > 0 {
			obj["name"] = p[i].Name
		}
		if len(p[i].NodePoolID) > 0 {
			obj["node_pool_id"] = p[i].NodePoolID
		}
		if len(p[i].OS) > 0 {
			obj["os"] = p[i].OS
		}
		if len(p[i].SecurityGroupIDs) > 0 {
			obj["security_group_ids"] = toArrayInterfaceSorted(p[i].SecurityGroupIDs)
		}
		if len(p[i].SubnetIDs) > 0 {
			obj["subnet_ids"] = toArrayInterfaceSorted(p[i].SubnetIDs)
		}
		if len(p[i].Taints) > 0 {
			obj["taints"] = flattenClusterTKEVirtualNodeTaint(p[i].Taints)
		}
		if len(p[i].VirtualNodes) > 0 {
			obj["virtual_nodes"] = flattenClusterTKEVirtualNodeSpec(p[i].VirtualNodes)
		}
		out[i] = obj
	}
	return out
}

func expandClusterTKEVirtualNodePoolDetail(p []interface{}) []managementClient.VirtualNodePoolDetail {
	if len(p) == 0 {
		return nil
	}
	out := make([]managementClient.VirtualNodePoolDetail, len(p))
	for i := range p {
		in := p[i].(map[string]interface{})
		if v, ok := in["deletion_protection"].(bool); ok {
			out[i].DeletionProtection = &v
		}
		if v, ok := in["labels"].([]interface{}); ok && len(v) > 0 {
			out[i].Labels = expandClusterTKEVirtualNodeLabel(v)
		}
		if v, ok := in["name"].(string); ok && len(v) > 0 {
			out[i].Name = v
		}
		if v, ok := in["node_pool_id"].(string); ok && len(v) > 0 {
			out[i].NodePoolID = v
		}
		if v, ok := in["os"].(string); ok && len(v) > 0 {
			out[i].OS = v
		}
		if v, ok := in["security_group_ids"].([]interface{}); ok && len(v) > 0 {
			out[i].SecurityGroupIDs = toArrayStringSorted(v)
		}
		if v, ok := in["subnet_ids"].([]interface{}); ok && len(v) > 0 {
			out[i].SubnetIDs = toArrayStringSorted(v)
		}
		if v, ok := in["taints"].([]interface{}); ok && len(v) > 0 {
			out[i].Taints = expandClusterTKEVirtualNodeTaint(v)
		}
		if v, ok := in["virtual_nodes"].([]interface{}); ok && len(v) > 0 {
			out[i].VirtualNodes = expandClusterTKEVirtualNodeSpec(v)
		}
	}
	return out
}

func flattenClusterTKEClusterBasicSettings(in *managementClient.ClusterBasicSettings) []interface{} {
	if in == nil {
		return nil
	}
	obj := map[string]interface{}{}
	if len(in.ClusterDescription) > 0 {
		obj["cluster_description"] = in.ClusterDescription
	}
	if len(in.ClusterLevel) > 0 {
		obj["cluster_level"] = in.ClusterLevel
	}
	if len(in.ClusterName) > 0 {
		obj["cluster_name"] = in.ClusterName
	}
	if len(in.ClusterOs) > 0 {
		obj["cluster_os"] = in.ClusterOs
	}
	if len(in.ClusterType) > 0 {
		obj["cluster_type"] = in.ClusterType
	}
	if len(in.ClusterVersion) > 0 {
		obj["cluster_version"] = in.ClusterVersion
	}
	obj["is_auto_upgrade"] = in.IsAutoUpgrade
	if in.ProjectID > 0 {
		obj["project_id"] = int(in.ProjectID)
	}
	if len(in.Tags) > 0 {
		obj["tags"] = toArrayInterfaceSorted(in.Tags)
	}
	if len(in.VpcID) > 0 {
		obj["vpc_id"] = in.VpcID
	}
	return []interface{}{obj}
}

func expandClusterTKEClusterBasicSettings(p []interface{}) *managementClient.ClusterBasicSettings {
	if len(p) == 0 || p[0] == nil {
		return nil
	}
	in := p[0].(map[string]interface{})
	obj := &managementClient.ClusterBasicSettings{}
	if v, ok := in["cluster_description"].(string); ok && len(v) > 0 {
		obj.ClusterDescription = v
	}
	if v, ok := in["cluster_level"].(string); ok && len(v) > 0 {
		obj.ClusterLevel = v
	}
	if v, ok := in["cluster_name"].(string); ok && len(v) > 0 {
		obj.ClusterName = v
	}
	if v, ok := in["cluster_os"].(string); ok && len(v) > 0 {
		obj.ClusterOs = v
	}
	if v, ok := in["cluster_type"].(string); ok && len(v) > 0 {
		obj.ClusterType = v
	}
	if v, ok := in["cluster_version"].(string); ok && len(v) > 0 {
		obj.ClusterVersion = v
	}
	if v, ok := in["is_auto_upgrade"].(bool); ok {
		obj.IsAutoUpgrade = v
	}
	if v, ok := in["project_id"].(int); ok && v > 0 {
		obj.ProjectID = int64(v)
	}
	if v, ok := in["tags"].([]interface{}); ok {
		obj.Tags = toArrayStringSorted(v)
	}
	if v, ok := in["vpc_id"].(string); ok && len(v) > 0 {
		obj.VpcID = v
	}
	return obj
}

func flattenClusterTKEClusterAdvancedSettings(in *managementClient.ClusterAdvancedSettings) []interface{} {
	if in == nil {
		return nil
	}
	obj := map[string]interface{}{}
	obj["as_enabled"] = in.AsEnabled
	obj["audit_enabled"] = in.AuditEnabled
	obj["audit_log_topic_id"] = in.AuditLogTopicID
	obj["audit_logset_id"] = in.AuditLogsetID
	if in.BasePodNumber > 0 {
		obj["base_pod_number"] = int(in.BasePodNumber)
	}
	obj["cilium_mode"] = in.CiliumMode
	obj["container_runtime"] = in.ContainerRuntime
	obj["deletion_protection"] = in.DeletionProtection
	obj["enable_customized_pod_cidr"] = in.EnableCustomizedPodCIDR
	if len(in.Etcd) > 0 {
		obj["etcd"] = toArrayInterfaceSorted(in.Etcd)
	}
	obj["ipvs"] = in.IPVS
	obj["is_dual_stack"] = in.IsDualStack
	obj["is_non_static_ip_mode"] = in.IsNonStaticIpMode
	if len(in.KubeAPIServer) > 0 {
		obj["kube_api_server"] = toArrayInterfaceSorted(in.KubeAPIServer)
	}
	if len(in.KubeControllerManager) > 0 {
		obj["kube_controller_manager"] = toArrayInterfaceSorted(in.KubeControllerManager)
	}
	obj["kube_proxy_mode"] = in.KubeProxyMode
	if len(in.KubeScheduler) > 0 {
		obj["kube_scheduler"] = toArrayInterfaceSorted(in.KubeScheduler)
	}
	obj["network_type"] = in.NetworkType
	obj["node_name_type"] = in.NodeNameType
	obj["qgpu_share_enable"] = in.QGPUShareEnable
	obj["runtime_version"] = in.RuntimeVersion
	obj["vpc_cni_type"] = in.VpcCniType
	return []interface{}{obj}
}

func expandClusterTKEClusterAdvancedSettings(p []interface{}) *managementClient.ClusterAdvancedSettings {
	if len(p) == 0 || p[0] == nil {
		return nil
	}
	in := p[0].(map[string]interface{})
	obj := &managementClient.ClusterAdvancedSettings{}
	if v, ok := in["as_enabled"].(bool); ok {
		obj.AsEnabled = v
	}
	if v, ok := in["audit_enabled"].(bool); ok {
		obj.AuditEnabled = v
	}
	if v, ok := in["audit_log_topic_id"].(string); ok && len(v) > 0 {
		obj.AuditLogTopicID = v
	}
	if v, ok := in["audit_logset_id"].(string); ok && len(v) > 0 {
		obj.AuditLogsetID = v
	}
	if v, ok := in["base_pod_number"].(int); ok && v > 0 {
		obj.BasePodNumber = int64(v)
	}
	if v, ok := in["cilium_mode"].(string); ok && len(v) > 0 {
		obj.CiliumMode = v
	}
	if v, ok := in["container_runtime"].(string); ok && len(v) > 0 {
		obj.ContainerRuntime = v
	}
	if v, ok := in["deletion_protection"].(bool); ok {
		obj.DeletionProtection = v
	}
	if v, ok := in["enable_customized_pod_cidr"].(bool); ok {
		obj.EnableCustomizedPodCIDR = v
	}
	if v, ok := in["etcd"].([]interface{}); ok {
		obj.Etcd = toArrayStringSorted(v)
	}
	if v, ok := in["ipvs"].(bool); ok {
		obj.IPVS = v
	}
	if v, ok := in["is_dual_stack"].(bool); ok {
		obj.IsDualStack = v
	}
	if v, ok := in["is_non_static_ip_mode"].(bool); ok {
		obj.IsNonStaticIpMode = v
	}
	if v, ok := in["kube_api_server"].([]interface{}); ok {
		obj.KubeAPIServer = toArrayStringSorted(v)
	}
	if v, ok := in["kube_controller_manager"].([]interface{}); ok {
		obj.KubeControllerManager = toArrayStringSorted(v)
	}
	if v, ok := in["kube_proxy_mode"].(string); ok && len(v) > 0 {
		obj.KubeProxyMode = v
	}
	if v, ok := in["kube_scheduler"].([]interface{}); ok {
		obj.KubeScheduler = toArrayStringSorted(v)
	}
	if v, ok := in["network_type"].(string); ok && len(v) > 0 {
		obj.NetworkType = v
	}
	if v, ok := in["node_name_type"].(string); ok && len(v) > 0 {
		obj.NodeNameType = v
	}
	if v, ok := in["qgpu_share_enable"].(bool); ok {
		obj.QGPUShareEnable = v
	}
	if v, ok := in["runtime_version"].(string); ok && len(v) > 0 {
		obj.RuntimeVersion = v
	}
	if v, ok := in["vpc_cni_type"].(string); ok && len(v) > 0 {
		obj.VpcCniType = v
	}
	return obj
}

func flattenClusterTKEClusterCIDRSettings(in *managementClient.ClusterCIDRSettings) []interface{} {
	if in == nil {
		return nil
	}
	obj := map[string]interface{}{}
	if in.ClaimExpiredSeconds > 0 {
		obj["claim_expired_seconds"] = int(in.ClaimExpiredSeconds)
	}
	if len(in.ClusterCIDR) > 0 {
		obj["cluster_cidr"] = in.ClusterCIDR
	}
	if len(in.EniSubnetIDs) > 0 {
		obj["eni_subnet_ids"] = toArrayInterfaceSorted(in.EniSubnetIDs)
	}
	obj["ignore_cluster_cidr_conflict"] = in.IgnoreClusterCIDRConflict
	obj["ignore_service_cidr_conflict"] = in.IgnoreServiceCIDRConflict
	if in.MaxClusterServiceNum > 0 {
		obj["max_cluster_service_num"] = int(in.MaxClusterServiceNum)
	}
	if in.MaxNodePodNum > 0 {
		obj["max_node_pod_num"] = int(in.MaxNodePodNum)
	}
	if len(in.OsCustomizeType) > 0 {
		obj["os_customize_type"] = in.OsCustomizeType
	}
	if len(in.ServiceCIDR) > 0 {
		obj["service_cidr"] = in.ServiceCIDR
	}
	if len(in.SubnetID) > 0 {
		obj["subnet_id"] = in.SubnetID
	}
	return []interface{}{obj}
}

func expandClusterTKEClusterCIDRSettings(p []interface{}) *managementClient.ClusterCIDRSettings {
	if len(p) == 0 || p[0] == nil {
		return nil
	}
	in := p[0].(map[string]interface{})
	obj := &managementClient.ClusterCIDRSettings{}
	if v, ok := in["claim_expired_seconds"].(int); ok && v > 0 {
		obj.ClaimExpiredSeconds = int64(v)
	}
	if v, ok := in["cluster_cidr"].(string); ok && len(v) > 0 {
		obj.ClusterCIDR = v
	}
	if v, ok := in["eni_subnet_ids"].([]interface{}); ok {
		obj.EniSubnetIDs = toArrayStringSorted(v)
	}
	if v, ok := in["ignore_cluster_cidr_conflict"].(bool); ok {
		obj.IgnoreClusterCIDRConflict = v
	}
	if v, ok := in["ignore_service_cidr_conflict"].(bool); ok {
		obj.IgnoreServiceCIDRConflict = v
	}
	if v, ok := in["max_cluster_service_num"].(int); ok && v > 0 {
		obj.MaxClusterServiceNum = int64(v)
	}
	if v, ok := in["max_node_pod_num"].(int); ok && v > 0 {
		obj.MaxNodePodNum = int64(v)
	}
	if v, ok := in["os_customize_type"].(string); ok && len(v) > 0 {
		obj.OsCustomizeType = v
	}
	if v, ok := in["service_cidr"].(string); ok && len(v) > 0 {
		obj.ServiceCIDR = v
	}
	if v, ok := in["subnet_id"].(string); ok && len(v) > 0 {
		obj.SubnetID = v
	}
	return obj
}

func flattenClusterTKEClusterEndpoint(in *managementClient.ClusterEndpoint) []interface{} {
	if in == nil {
		return nil
	}
	obj := map[string]interface{}{}
	if len(in.Domain) > 0 {
		obj["domain"] = in.Domain
	}
	obj["enable"] = in.Enable
	if len(in.ExtensiveParameters) > 0 {
		obj["extensive_parameters"] = in.ExtensiveParameters
	}
	if len(in.SecurityGroup) > 0 {
		obj["security_group"] = in.SecurityGroup
	}
	if len(in.SubnetID) > 0 {
		obj["subnet_id"] = in.SubnetID
	}
	return []interface{}{obj}
}

func flattenClusterTKEImportedClusterEndpoint(in *managementClient.ClusterEndpoint, p []interface{}) []interface{} {
	obj := map[string]interface{}{}

	if in != nil {
		obj["enable"] = in.Enable
		return []interface{}{obj}
	}

	if len(p) == 0 || p[0] == nil {
		return nil
	}

	prev, ok := p[0].(map[string]interface{})
	if !ok {
		return nil
	}

	if v, exists := prev["enable"]; exists {
		if enable, ok := v.(bool); ok {
			obj["enable"] = enable
			return []interface{}{obj}
		}
	}

	return nil
}

func expandClusterTKEClusterEndpoint(p []interface{}) *managementClient.ClusterEndpoint {
	if len(p) == 0 || p[0] == nil {
		return nil
	}
	in := p[0].(map[string]interface{})
	obj := &managementClient.ClusterEndpoint{}
	if v, ok := in["domain"].(string); ok && len(v) > 0 {
		obj.Domain = v
	}
	if v, ok := in["enable"].(bool); ok {
		obj.Enable = v
	}
	if v, ok := in["extensive_parameters"].(string); ok && len(v) > 0 {
		obj.ExtensiveParameters = v
	}
	if v, ok := in["security_group"].(string); ok && len(v) > 0 {
		obj.SecurityGroup = v
	}
	if v, ok := in["subnet_id"].(string); ok && len(v) > 0 {
		obj.SubnetID = v
	}
	return obj
}

func expandClusterTKEImportedClusterEndpoint(p []interface{}) *managementClient.ClusterEndpoint {
	if len(p) == 0 || p[0] == nil {
		return nil
	}

	in := p[0].(map[string]interface{})
	obj := &managementClient.ClusterEndpoint{}

	if v, ok := in["enable"].(bool); ok {
		obj.Enable = v
	}

	return obj
}

func flattenClusterTKEExtensionAddon(p []managementClient.ExtensionAddon) []interface{} {
	if len(p) == 0 {
		return nil
	}
	out := make([]interface{}, len(p))
	for i := range p {
		obj := map[string]interface{}{}
		if len(p[i].AddonName) > 0 {
			obj["addon_name"] = p[i].AddonName
		}
		if len(p[i].AddonParam) > 0 {
			obj["addon_param"] = p[i].AddonParam
		}
		out[i] = obj
	}
	return out
}

func expandClusterTKEExtensionAddon(p []interface{}) []managementClient.ExtensionAddon {
	if len(p) == 0 {
		return nil
	}
	out := make([]managementClient.ExtensionAddon, len(p))
	for i := range p {
		in := p[i].(map[string]interface{})
		if v, ok := in["addon_name"].(string); ok && len(v) > 0 {
			out[i].AddonName = v
		}
		if v, ok := in["addon_param"].(string); ok && len(v) > 0 {
			out[i].AddonParam = v
		}
	}
	return out
}

func flattenClusterTKEConfigV2(in *managementClient.TKEClusterConfigSpec, p []interface{}) []interface{} {
	if in == nil {
		return nil
	}

	if in.Imported {
		obj := map[string]interface{}{
			"imported": in.Imported,
		}

		var previous map[string]interface{}
		if len(p) > 0 && p[0] != nil {
			previous, _ = p[0].(map[string]interface{})
		}

		if len(in.ClusterID) > 0 {
			obj["cluster_id"] = in.ClusterID
		} else if previous != nil {
			if v, ok := previous["cluster_id"].(string); ok && len(v) > 0 {
				obj["cluster_id"] = v
			}
		}
		if len(in.Region) > 0 {
			obj["region"] = in.Region
		} else if previous != nil {
			if v, ok := previous["region"].(string); ok && len(v) > 0 {
				obj["region"] = v
			}
		}
		if len(in.TKECredentialSecret) > 0 {
			obj["tke_credential_secret"] = in.TKECredentialSecret
		} else if previous != nil {
			if v, ok := previous["tke_credential_secret"].(string); ok && len(v) > 0 {
				obj["tke_credential_secret"] = v
			}
		}

		var previousClusterEndpoint []interface{}
		if previous != nil {
			if v, ok := previous["cluster_endpoint"].([]interface{}); ok {
				previousClusterEndpoint = v
			}
		}
		if endpoint := flattenClusterTKEImportedClusterEndpoint(in.ClusterEndpoint, previousClusterEndpoint); len(endpoint) > 0 {
			obj["cluster_endpoint"] = endpoint
		}

		return []interface{}{obj}
	}

	obj := map[string]interface{}{}
	if len(p) > 0 && p[0] != nil {
		obj = p[0].(map[string]interface{})
	}

	obj["imported"] = in.Imported
	if len(in.ClusterID) > 0 {
		obj["cluster_id"] = in.ClusterID
	}
	if len(in.Region) > 0 {
		obj["region"] = in.Region
	}
	if len(in.TKECredentialSecret) > 0 {
		obj["tke_credential_secret"] = in.TKECredentialSecret
	}

	if in.ClusterBasicSettings != nil {
		obj["cluster_basic_settings"] = flattenClusterTKEClusterBasicSettings(in.ClusterBasicSettings)
	}
	if in.ClusterAdvancedSettings != nil {
		obj["cluster_advanced_settings"] = flattenClusterTKEClusterAdvancedSettings(in.ClusterAdvancedSettings)
	}
	if in.ClusterCIDRSettings != nil {
		obj["cluster_cidr_settings"] = flattenClusterTKEClusterCIDRSettings(in.ClusterCIDRSettings)
	}
	if in.ClusterEndpoint != nil {
		obj["cluster_endpoint"] = flattenClusterTKEClusterEndpoint(in.ClusterEndpoint)
	}
	if len(in.ExtensionAddon) > 0 {
		obj["extension_addon"] = flattenClusterTKEExtensionAddon(in.ExtensionAddon)
	}
	if len(in.NodePoolList) > 0 {
		obj["node_pool_list"] = flattenClusterTKENodePoolDetail(in.NodePoolList)
	}
	if len(in.VirtualNodePoolList) > 0 {
		obj["virtual_node_pool_list"] = flattenClusterTKEVirtualNodePoolDetail(in.VirtualNodePoolList)
	}
	if in.RunInstancesForNode != nil {
		obj["run_instances_for_node"] = flattenClusterTKERunInstancesForNode(in.RunInstancesForNode)
	}

	return []interface{}{obj}
}

func expandClusterTKEConfigV2(p []interface{}, clusterName string) *managementClient.TKEClusterConfigSpec {
	obj := &managementClient.TKEClusterConfigSpec{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}

	in := p[0].(map[string]interface{})
	if v, ok := in["imported"].(bool); ok {
		obj.Imported = v
	}
	if v, ok := in["cluster_id"].(string); ok && len(v) > 0 {
		obj.ClusterID = v
	}
	if v, ok := in["region"].(string); ok && len(v) > 0 {
		obj.Region = v
	}
	if v, ok := in["tke_credential_secret"].(string); ok && len(v) > 0 {
		obj.TKECredentialSecret = v
	}

	if obj.Imported {
		if v, ok := in["cluster_endpoint"].([]interface{}); ok && len(v) > 0 {
			obj.ClusterEndpoint = expandClusterTKEImportedClusterEndpoint(v)
		}
		return obj
	}

	if v, ok := in["cluster_endpoint"].([]interface{}); ok && len(v) > 0 {
		obj.ClusterEndpoint = expandClusterTKEClusterEndpoint(v)
	}

	if v, ok := in["cluster_basic_settings"].([]interface{}); ok && len(v) > 0 {
		obj.ClusterBasicSettings = expandClusterTKEClusterBasicSettings(v)
	}
	if obj.ClusterBasicSettings == nil {
		obj.ClusterBasicSettings = &managementClient.ClusterBasicSettings{}
	}
	if len(obj.ClusterBasicSettings.ClusterName) == 0 && len(clusterName) > 0 {
		obj.ClusterBasicSettings.ClusterName = clusterName
	}

	if v, ok := in["cluster_advanced_settings"].([]interface{}); ok && len(v) > 0 {
		obj.ClusterAdvancedSettings = expandClusterTKEClusterAdvancedSettings(v)
	}
	if v, ok := in["cluster_cidr_settings"].([]interface{}); ok && len(v) > 0 {
		obj.ClusterCIDRSettings = expandClusterTKEClusterCIDRSettings(v)
	}
	if v, ok := in["extension_addon"].([]interface{}); ok && len(v) > 0 {
		obj.ExtensionAddon = expandClusterTKEExtensionAddon(v)
	}
	if v, ok := in["node_pool_list"].([]interface{}); ok && len(v) > 0 {
		obj.NodePoolList = expandClusterTKENodePoolDetail(v)
	}
	if v, ok := in["virtual_node_pool_list"].([]interface{}); ok && len(v) > 0 {
		obj.VirtualNodePoolList = expandClusterTKEVirtualNodePoolDetail(v)
	}
	if v, ok := in["run_instances_for_node"].([]interface{}); ok && len(v) > 0 {
		obj.RunInstancesForNode = expandClusterTKERunInstancesForNode(v)
	}

	return obj
}

func fixClusterTKEConfigV2(p []interface{}, values map[string]interface{}) map[string]interface{} {
	if len(p) == 0 || p[0] == nil {
		return values
	}
	if values == nil {
		values = map[string]interface{}{}
	}

	in := p[0].(map[string]interface{})
	imported, _ := in["imported"].(bool)
	if v, ok := in["cluster_endpoint"].([]interface{}); ok && len(v) > 0 {
		if imported {
			values["clusterEndpoint"] = fixClusterTKEImportedClusterEndpoint(v)
		} else {
			values["clusterEndpoint"] = fixClusterTKEClusterEndpoint(v)
		}
	}

	return values
}

func fixClusterTKEImportedClusterEndpoint(p []interface{}) map[string]interface{} {
	if len(p) == 0 || p[0] == nil {
		return nil
	}

	in := p[0].(map[string]interface{})
	obj := map[string]interface{}{}

	if v, exists := in["enable"]; exists {
		if enable, ok := v.(bool); ok {
			obj["enable"] = enable
		}
	}

	return obj
}

func fixClusterTKEClusterEndpoint(p []interface{}) map[string]interface{} {
	if len(p) == 0 || p[0] == nil {
		return nil
	}

	in := p[0].(map[string]interface{})
	obj := map[string]interface{}{}

	if v, ok := in["domain"].(string); ok && len(v) > 0 {
		obj["domain"] = v
	}
	if v, exists := in["enable"]; exists {
		if enable, ok := v.(bool); ok {
			obj["enable"] = enable
		}
	}
	if v, ok := in["extensive_parameters"].(string); ok && len(v) > 0 {
		obj["extensiveParameters"] = v
	}
	if v, ok := in["security_group"].(string); ok && len(v) > 0 {
		obj["securityGroup"] = v
	}
	if v, ok := in["subnet_id"].(string); ok && len(v) > 0 {
		obj["subnetId"] = v
	}

	return obj
}
