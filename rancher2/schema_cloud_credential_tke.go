package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// Types

type tkeCredentialConfig struct {
	AccessKeyID     string `json:"accessKeyId,omitempty" yaml:"accessKeyId,omitempty"`
	AccessKeySecret string `json:"accessKeySecret,omitempty" yaml:"accessKeySecret,omitempty"`
}

// Schemas

func cloudCredentialTKEFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_key_id": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "TKE access key id",
		},
		"access_key_secret": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "TKE access key secret",
		},
	}
}
