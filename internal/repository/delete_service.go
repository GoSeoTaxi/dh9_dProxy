package repository

import "fmt"

func (r *Repository) deleteService(serviceName string) error {
	serviceURL := fmt.Sprintf("http://%s:%s/config/services/%s", r.gostV3Host, r.gostV3Port, serviceName)

	resp, err := r.client.R().
		SetBasicAuth(r.gostV3Username, r.gostV3Password).
		Delete(serviceURL)

	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса DELETE: %w", err)
	}

	return parseResponseOKinBody(resp.Body())
}
