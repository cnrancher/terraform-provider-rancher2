package rancher2

// Flatteners

func flattenCloudCredentialTKE(in *tkeCredentialConfig, p []interface{}) []interface{} {
	var obj map[string]interface{}
	if len(p) == 0 || p[0] == nil {
		obj = make(map[string]interface{})
	} else {
		obj = p[0].(map[string]interface{})
	}

	if in == nil {
		return []interface{}{}
	}

	if len(in.AccessKeyID) > 0 {
		obj["access_key_id"] = in.AccessKeyID
	}
	if len(in.AccessKeySecret) > 0 {
		obj["access_key_secret"] = in.AccessKeySecret
	}

	return []interface{}{obj}
}

// Expanders

func expandCloudCredentialTKE(p []interface{}) *tkeCredentialConfig {
	obj := &tkeCredentialConfig{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["access_key_id"].(string); ok && len(v) > 0 {
		obj.AccessKeyID = v
	}
	if v, ok := in["access_key_secret"].(string); ok && len(v) > 0 {
		obj.AccessKeySecret = v
	}

	return obj
}
