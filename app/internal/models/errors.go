package models

import (
	"encoding/json"
	"net/http"
)

func NotFoundBody() string {
	body, _ := json.Marshal(ApiResponse{Code: http.StatusNotFound, Type: "error", Message: "not found"})
	return string(body)
}
