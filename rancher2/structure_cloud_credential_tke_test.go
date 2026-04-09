package rancher2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCloudCredentialTKEConf      *tkeCredentialConfig
	testCloudCredentialTKEInterface []interface{}
)

func init() {
	testCloudCredentialTKEConf = &tkeCredentialConfig{
		AccessKeyID:     "access_key_id",
		AccessKeySecret: "access_key_secret",
	}
	testCloudCredentialTKEInterface = []interface{}{
		map[string]interface{}{
			"access_key_id":     "access_key_id",
			"access_key_secret": "access_key_secret",
		},
	}
}

func TestFlattenCloudCredentialTKE(t *testing.T) {
	cases := []struct {
		Input          *tkeCredentialConfig
		ExpectedOutput []interface{}
	}{
		{
			testCloudCredentialTKEConf,
			testCloudCredentialTKEInterface,
		},
	}

	for _, tc := range cases {
		output := flattenCloudCredentialTKE(tc.Input, tc.ExpectedOutput)
		assert.Equal(t, tc.ExpectedOutput, output, "Unexpected output from flattener.")
	}
}

func TestExpandCloudCredentialTKE(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *tkeCredentialConfig
	}{
		{
			testCloudCredentialTKEInterface,
			testCloudCredentialTKEConf,
		},
	}

	for _, tc := range cases {
		output := expandCloudCredentialTKE(tc.Input)
		assert.Equal(t, tc.ExpectedOutput, output, "Unexpected output from expander.")
	}
}
