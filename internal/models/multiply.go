package models

type MultiplyRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type MultiplyResponse struct {
	Message string `json:"message"`
	Result  int    `json:"product"`
}
