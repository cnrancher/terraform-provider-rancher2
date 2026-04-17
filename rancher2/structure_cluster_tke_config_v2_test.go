package rancher2

import (
	"testing"

	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
	"github.com/stretchr/testify/assert"
)

var (
	testClusterTKEConfigV2Conf      *managementClient.TKEClusterConfigSpec
	testClusterTKEConfigV2Interface []interface{}
)

func init() {
	deletionProtection := true

	testClusterTKEConfigV2Conf = &managementClient.TKEClusterConfigSpec{
		Imported:            false,
		ClusterID:           "",
		Region:              "ap-guangzhou",
		TKECredentialSecret: "cattle-global-data:cc-tke",
		ClusterBasicSettings: &managementClient.ClusterBasicSettings{
			ClusterName:    "test",
			ClusterType:    "MANAGED_CLUSTER",
			ClusterVersion: "v1.30.1",
		},
		ClusterAdvancedSettings: &managementClient.ClusterAdvancedSettings{
			ContainerRuntime: "containerd",
			NetworkType:      "VPC-CNI",
		},
		ClusterCIDRSettings: &managementClient.ClusterCIDRSettings{
			ClusterCIDR: "10.0.0.0/16",
			ServiceCIDR: "10.96.0.0/16",
		},
		ClusterEndpoint: &managementClient.ClusterEndpoint{
			Enable: true,
		},
		ExtensionAddon: []managementClient.ExtensionAddon{{
			AddonName:  "abc",
			AddonParam: "x=y",
		}},
		NodePoolList: []managementClient.NodePoolDetail{{
			Name:            "pool-a",
			EnableAutoscale: true,
			UserScript:      "#!/bin/bash\necho hello",
			AutoScalingGroupPara: &managementClient.AutoScalingGroupPara{
				MinSize:         1,
				MaxSize:         3,
				DesiredCapacity: 1,
			},
		}},
		VirtualNodePoolList: []managementClient.VirtualNodePoolDetail{{
			DeletionProtection: &deletionProtection,
			Labels: []managementClient.VirtualNodeLabel{{
				Name:  "env",
				Value: "test",
			}},
			Name:             "virtual-pool-a",
			NodePoolID:       "vnp-123",
			OS:               "tlinux3.1x86_64",
			SecurityGroupIDs: []string{"sg-123"},
			SubnetIDs:        []string{"subnet-123"},
			Taints: []managementClient.VirtualNodeTaint{{
				Effect: "NoSchedule",
				Key:    "dedicated",
				Value:  "virtual",
			}},
			VirtualNodes: []managementClient.VirtualNodeSpec{{
				DisplayName: "virtual-node-a",
				SubnetId:    "subnet-123",
				Tags: []managementClient.VirtualNodeTag{{
					Key:   "owner",
					Value: "terraform",
				}},
			}},
		}},
		RunInstancesForNode: &managementClient.RunInstancesForNode{
			InstanceType:  "S5.MEDIUM2",
			InstanceCount: 1,
		},
	}

	testClusterTKEConfigV2Interface = []interface{}{
		map[string]interface{}{
			"imported":              false,
			"region":                "ap-guangzhou",
			"tke_credential_secret": "cattle-global-data:cc-tke",
			"cluster_basic_settings": []interface{}{
				map[string]interface{}{
					"cluster_name":    "test",
					"cluster_type":    "MANAGED_CLUSTER",
					"cluster_version": "v1.30.1",
					"is_auto_upgrade": false,
				},
			},
			"cluster_advanced_settings": []interface{}{
				map[string]interface{}{
					"as_enabled":                 false,
					"audit_enabled":              false,
					"audit_log_topic_id":         "",
					"audit_logset_id":            "",
					"container_runtime":          "containerd",
					"deletion_protection":        false,
					"enable_customized_pod_cidr": false,
					"ipvs":                       false,
					"is_dual_stack":              false,
					"is_non_static_ip_mode":      false,
					"kube_proxy_mode":            "",
					"network_type":               "VPC-CNI",
					"node_name_type":             "",
					"qgpu_share_enable":          false,
					"runtime_version":            "",
					"vpc_cni_type":               "",
				},
			},
			"cluster_cidr_settings": []interface{}{
				map[string]interface{}{
					"cluster_cidr":                 "10.0.0.0/16",
					"ignore_cluster_cidr_conflict": false,
					"ignore_service_cidr_conflict": false,
					"service_cidr":                 "10.96.0.0/16",
				},
			},
			"cluster_endpoint": []interface{}{
				map[string]interface{}{
					"enable": true,
				},
			},
			"extension_addon": []interface{}{
				map[string]interface{}{
					"addon_name":  "abc",
					"addon_param": "x=y",
				},
			},
			"node_pool_list": []interface{}{
				map[string]interface{}{
					"name":                "pool-a",
					"deletion_protection": false,
					"enable_autoscale":    true,
					"user_script":         "#!/bin/bash\necho hello",
					"auto_scaling_group_para": []interface{}{
						map[string]interface{}{
							"desired_capacity": 1,
							"max_size":         3,
							"min_size":         1,
						},
					},
				},
			},
			"virtual_node_pool_list": []interface{}{
				map[string]interface{}{
					"deletion_protection": true,
					"labels": []interface{}{
						map[string]interface{}{
							"name":  "env",
							"value": "test",
						},
					},
					"name":               "virtual-pool-a",
					"node_pool_id":       "vnp-123",
					"os":                 "tlinux3.1x86_64",
					"security_group_ids": []interface{}{"sg-123"},
					"subnet_ids":         []interface{}{"subnet-123"},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": "NoSchedule",
							"key":    "dedicated",
							"value":  "virtual",
						},
					},
					"virtual_nodes": []interface{}{
						map[string]interface{}{
							"display_name": "virtual-node-a",
							"subnet_id":    "subnet-123",
							"tags": []interface{}{
								map[string]interface{}{
									"key":   "owner",
									"value": "terraform",
								},
							},
						},
					},
				},
			},
			"run_instances_for_node": []interface{}{
				map[string]interface{}{
					"instance_type":      "S5.MEDIUM2",
					"instance_count":     1,
					"monitor_service":    false,
					"public_ip_assigned": false,
					"security_service":   false,
				},
			},
		},
	}
}

func TestFlattenClusterTKEConfigV2(t *testing.T) {
	output := flattenClusterTKEConfigV2(testClusterTKEConfigV2Conf, testClusterTKEConfigV2Interface)
	assert.Equal(t, testClusterTKEConfigV2Interface, output, "Unexpected output from flattener.")
}

func TestExpandClusterTKEConfigV2(t *testing.T) {
	output := expandClusterTKEConfigV2(testClusterTKEConfigV2Interface, "test")
	assert.Equal(t, testClusterTKEConfigV2Conf, output, "Unexpected output from expander.")
}

func TestExpandClusterTKEConfigV2ImportedWithClusterEndpoint(t *testing.T) {
	input := []interface{}{
		map[string]interface{}{
			"imported":              true,
			"cluster_id":            "cls-123",
			"region":                "ap-guangzhou",
			"tke_credential_secret": "cattle-global-data:cc-tke",
			"cluster_endpoint": []interface{}{
				map[string]interface{}{
					"domain":               "example.com",
					"enable":               false,
					"extensive_parameters": "foo=bar",
					"security_group":       "sg-123",
					"subnet_id":            "subnet-123",
				},
			},
		},
	}

	output := expandClusterTKEConfigV2(input, "test")

	if assert.NotNil(t, output) {
		assert.True(t, output.Imported)
		assert.Equal(t, "cls-123", output.ClusterID)
		assert.Equal(t, "ap-guangzhou", output.Region)
		assert.Equal(t, "cattle-global-data:cc-tke", output.TKECredentialSecret)
		if assert.NotNil(t, output.ClusterEndpoint) {
			assert.False(t, output.ClusterEndpoint.Enable)
			assert.Empty(t, output.ClusterEndpoint.Domain)
			assert.Empty(t, output.ClusterEndpoint.ExtensiveParameters)
			assert.Empty(t, output.ClusterEndpoint.SecurityGroup)
			assert.Empty(t, output.ClusterEndpoint.SubnetID)
		}
	}
}

func TestFlattenClusterTKEConfigV2ImportedKeepsManagedFieldsOnly(t *testing.T) {
	input := &managementClient.TKEClusterConfigSpec{
		Imported:  true,
		ClusterID: "cls-123",
		Region:    "ap-guangzhou",
		ClusterEndpoint: &managementClient.ClusterEndpoint{
			Domain:              "example.com",
			Enable:              false,
			ExtensiveParameters: "foo=bar",
			SecurityGroup:       "sg-123",
			SubnetID:            "subnet-123",
		},
	}

	previous := []interface{}{
		map[string]interface{}{
			"tke_credential_secret": "cattle-global-data:cc-tke",
			"cluster_basic_settings": []interface{}{
				map[string]interface{}{
					"cluster_name": "legacy",
				},
			},
			"node_pool_list": []interface{}{
				map[string]interface{}{
					"name": "legacy-pool",
				},
			},
			"cluster_endpoint": []interface{}{
				map[string]interface{}{
					"domain": "legacy.example.com",
					"enable": true,
				},
			},
		},
	}

	output := flattenClusterTKEConfigV2(input, previous)

	expected := []interface{}{
		map[string]interface{}{
			"imported":              true,
			"cluster_id":            "cls-123",
			"region":                "ap-guangzhou",
			"tke_credential_secret": "cattle-global-data:cc-tke",
			"cluster_endpoint": []interface{}{
				map[string]interface{}{
					"enable": false,
				},
			},
		},
	}

	assert.Equal(t, expected, output)
}

func TestFlattenClusterTKEConfigV2ImportedPreservesPreviousEndpointEnable(t *testing.T) {
	input := &managementClient.TKEClusterConfigSpec{
		Imported:  true,
		ClusterID: "cls-123",
		Region:    "ap-guangzhou",
	}

	previous := []interface{}{
		map[string]interface{}{
			"tke_credential_secret": "cattle-global-data:cc-tke",
			"cluster_endpoint": []interface{}{
				map[string]interface{}{
					"domain": "legacy.example.com",
					"enable": true,
				},
			},
		},
	}

	output := flattenClusterTKEConfigV2(input, previous)

	expected := []interface{}{
		map[string]interface{}{
			"imported":              true,
			"cluster_id":            "cls-123",
			"region":                "ap-guangzhou",
			"tke_credential_secret": "cattle-global-data:cc-tke",
			"cluster_endpoint": []interface{}{
				map[string]interface{}{
					"enable": true,
				},
			},
		},
	}

	assert.Equal(t, expected, output)
}

func TestFixClusterTKEConfigV2PreservesClusterEndpointEnableFalse(t *testing.T) {
	input := []interface{}{
		map[string]interface{}{
			"imported": true,
			"cluster_endpoint": []interface{}{
				map[string]interface{}{
					"domain":               "example.com",
					"enable":               false,
					"extensive_parameters": "foo=bar",
					"security_group":       "sg-123",
					"subnet_id":            "subnet-123",
				},
			},
		},
	}

	values := map[string]interface{}{
		"imported": true,
	}

	output := fixClusterTKEConfigV2(input, values)

	endpoint, ok := output["clusterEndpoint"].(map[string]interface{})
	if assert.True(t, ok) {
		assert.Contains(t, endpoint, "enable")
		assert.Equal(t, false, endpoint["enable"])
		assert.NotContains(t, endpoint, "domain")
		assert.NotContains(t, endpoint, "extensiveParameters")
		assert.NotContains(t, endpoint, "securityGroup")
		assert.NotContains(t, endpoint, "subnetId")
	}
}
