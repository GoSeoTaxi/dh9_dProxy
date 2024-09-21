package repository

import (
	"encoding/json"
	"fmt"
	"strings"
)

type responseBody struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg"`
}

func parseResponseOKinBody(body []byte) error {
	var responseBodyjson responseBody
	err := json.Unmarshal(body, &responseBodyjson)
	if err != nil {
		return fmt.Errorf("ошибка при разборе ответа сервера: %w", err)
	}

	if responseBodyjson.Msg == "OK" || strings.Contains(responseBodyjson.Msg, "not found") {
		return nil
	}

	return fmt.Errorf("неожиданный ответ от сервера: %s", string(body))
}
