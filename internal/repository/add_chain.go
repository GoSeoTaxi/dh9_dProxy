package repository

import (
	"fmt"

	"github.com/GoSeoTaxi/dh9_dProxy/internal/model"
)

func (r *Repository) addChain(data string, namePrefix int64) error {

	proxy, err := model.ParseProxy(data)
	if err != nil {
		return err
	}

	gostV3URL := fmt.Sprintf("http://%s:%s/config/chains", r.gostV3Host, r.gostV3Port)

	requestBody := map[string]interface{}{
		"name":     fmt.Sprintf("chain-%d", namePrefix),
		"resolver": "resolver-0",
		"hops": []map[string]interface{}{
			{
				"name": fmt.Sprintf("hop-%d", namePrefix),
				"nodes": []map[string]interface{}{
					{
						"name":     fmt.Sprintf("node-%d", namePrefix),
						"addr":     proxy.Addr,
						"resolver": "resolver-0",
						"connector": map[string]interface{}{
							"type": proxy.Type,
							"auth": map[string]string{
								"username": proxy.Login,
								"password": proxy.Pass,
							},
						},
						"dialer": map[string]interface{}{
							"type": "tcp",
						},
					},
				},
			},
		},
	}

	_, err = r.makeRequestPost(gostV3URL, requestBody)
	if err != nil {
		return err
	}

	return nil
}
