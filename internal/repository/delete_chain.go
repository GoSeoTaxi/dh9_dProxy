package repository

import "fmt"

func (r *Repository) deleteProxyChain(chainName string) error {
	gostV3URL := fmt.Sprintf("http://%s:%s/config/chains/%s", r.gostV3Host, r.gostV3Port, chainName)

	resp, err := r.client.R().
		SetBasicAuth(r.gostV3Username, r.gostV3Password).
		Delete(gostV3URL)

	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса DELETE: %w", err)
	}

	return parseResponseOKinBody(resp.Body())
}
