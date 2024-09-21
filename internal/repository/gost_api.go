package repository

func (r *Repository) makeRequestPost(url string, requestBody map[string]interface{}) (string, error) {
	resp, err := r.client.R().
		SetHeader("Content-Type", "application/json").
		SetBasicAuth(r.gostV3Username, r.gostV3Password).
		SetBody(requestBody).
		Post(url)

	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil
}
