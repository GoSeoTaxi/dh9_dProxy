package repository

import (
	"fmt"
)

func (r *Repository) addService(serviceName string, chainName string, namePrefix string, typeProxy string, authUser string, authPass string) error {

	gostV3URL := fmt.Sprintf("http://%s:%s/config/services", r.gostV3Host, r.gostV3Port)

	requestBody := map[string]interface{}{
		"name": serviceName,
		"addr": fmt.Sprintf(":%v", namePrefix),
		"handler": map[string]interface{}{
			"type":  typeProxy,
			"chain": chainName,
			"auth": map[string]string{
				"username": authUser,
				"password": authPass,
			},
		},
		"listener": map[string]interface{}{
			"type": "tcp",
		},
		"forwarder": map[string]interface{}{
			"nodes": []map[string]interface{}{
				{
					"name": "node-" + namePrefix,
					"addr": ":0",
				},
			},
		},
	}

	_, err := r.makeRequestPost(gostV3URL, requestBody)
	if err != nil {
		return err
	}

	return nil
}
