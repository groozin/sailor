package models

type AddRequest struct {
	Numbers []int `json:"numbers"`
}

type PrimeRequest struct {
	Number int `json:"number"`
}

type MultiplyRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SubtractRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response struct {
	Message string `json:"message"`
	Result  int    `json:"result"`
}

type PrimeResponse struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
}
