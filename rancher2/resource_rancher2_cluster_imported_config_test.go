package rancher2

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
	"github.com/stretchr/testify/assert"
)

func TestResourceRancher2ClusterDiffAllowsImportedConfigForManagedTKE(t *testing.T) {
	resourceConfig := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name": "test-tke",
		"tke_config_v2": []interface{}{
			map[string]interface{}{
				"imported":              false,
				"region":                "ap-guangzhou",
				"tke_credential_secret": "cattle-global-data:cc-tke",
			},
		},
		"imported_config": []interface{}{
			map[string]interface{}{
				"private_registry_url": "registry.example.com",
			},
		},
	})

	_, err := resourceRancher2Cluster().Diff(nil, resourceConfig, nil)
	assert.NoError(t, err)
}

func TestResourceRancher2ClusterDiffRejectsImportedConfigForManagedGKE(t *testing.T) {
	resourceConfig := terraform.NewResourceConfigRaw(map[string]interface{}{
		"name": "test-gke",
		"gke_config_v2": []interface{}{
			map[string]interface{}{
				"name":                     "test-gke",
				"google_credential_secret": "cattle-global-data:cc-gke",
				"project_id":               "project-id",
				"region":                   "us-central1",
				"imported":                 false,
			},
		},
		"imported_config": []interface{}{
			map[string]interface{}{
				"private_registry_url": "registry.example.com",
			},
		},
	})

	_, err := resourceRancher2Cluster().Diff(nil, resourceConfig, nil)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "rancher2_cluster.imported_config")
	}
}

func TestFlattenClusterTKEImportedConfigState(t *testing.T) {
	cases := []struct {
		name           string
		cluster        *Cluster
		initialState   map[string]interface{}
		expectedConfig []interface{}
	}{
		{
			name: "set imported config for tke",
			cluster: &Cluster{
				Cluster: managementClient.Cluster{
					Name:   "test-tke",
					Driver: clusterDriverTKEV2,
					TKEConfig: &managementClient.TKEClusterConfigSpec{
						Region:              "ap-guangzhou",
						TKECredentialSecret: "cattle-global-data:cc-tke",
					},
					ImportedConfig: &managementClient.ImportedConfig{
						PrivateRegistryURL: "registry.example.com",
					},
				},
			},
			initialState: map[string]interface{}{},
			expectedConfig: []interface{}{
				map[string]interface{}{
					"private_registry_url": "registry.example.com",
				},
			},
		},
		{
			name: "clear imported config for tke",
			cluster: &Cluster{
				Cluster: managementClient.Cluster{
					Name:   "test-tke",
					Driver: clusterDriverTKEV2,
					TKEConfig: &managementClient.TKEClusterConfigSpec{
						Region:              "ap-guangzhou",
						TKECredentialSecret: "cattle-global-data:cc-tke",
					},
				},
			},
			initialState: map[string]interface{}{
				"imported_config": []interface{}{
					map[string]interface{}{
						"private_registry_url": "stale.registry.example.com",
					},
				},
			},
			expectedConfig: []interface{}{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resourceData := schema.TestResourceDataRaw(t, clusterFields(), tc.initialState)

			err := flattenCluster(
				resourceData,
				tc.cluster,
				nil,
				&managementClient.GenerateKubeConfigOutput{},
				"",
				"",
			)
			if !assert.NoError(t, err) {
				return
			}

			actual, ok := resourceData.Get("imported_config").([]interface{})
			if !assert.True(t, ok) {
				return
			}
			if len(tc.expectedConfig) == 0 {
				assert.Len(t, actual, 0)
				return
			}
			assert.Equal(t, tc.expectedConfig, actual)
		})
	}
}
