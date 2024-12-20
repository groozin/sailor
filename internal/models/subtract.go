package models

type SubtractRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SubtractResponse struct {
	Message string `json:"message"`
	Result  int    `json:"difference"`
}
