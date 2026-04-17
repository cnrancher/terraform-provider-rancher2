package rancher2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateClusterTKEImportedConfigV2AllowsManagedFields(t *testing.T) {
	input := []interface{}{
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

	err := validateClusterTKEImportedConfigV2(input)
	assert.NoError(t, err)
}

func TestValidateClusterTKEImportedConfigV2RejectsBlockedBlocks(t *testing.T) {
	input := []interface{}{
		map[string]interface{}{
			"imported": true,
			"cluster_basic_settings": []interface{}{
				map[string]interface{}{
					"cluster_name": "test",
				},
			},
		},
	}

	err := validateClusterTKEImportedConfigV2(input)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "cluster_basic_settings")
	}
}

func TestValidateClusterTKEImportedConfigV2RejectsBlockedEndpointFields(t *testing.T) {
	input := []interface{}{
		map[string]interface{}{
			"imported": true,
			"cluster_endpoint": []interface{}{
				map[string]interface{}{
					"domain": "example.com",
					"enable": true,
				},
			},
		},
	}

	err := validateClusterTKEImportedConfigV2(input)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "cluster_endpoint.domain")
	}
}

func TestValidateClusterTKEImportedConfigV2SkipsNonImportedClusters(t *testing.T) {
	input := []interface{}{
		map[string]interface{}{
			"imported": false,
			"cluster_basic_settings": []interface{}{
				map[string]interface{}{
					"cluster_name": "test",
				},
			},
			"cluster_endpoint": []interface{}{
				map[string]interface{}{
					"domain": "example.com",
				},
			},
		},
	}

	err := validateClusterTKEImportedConfigV2(input)
	assert.NoError(t, err)
}
