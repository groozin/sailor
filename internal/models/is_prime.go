package models

type PrimeRequest struct {
	Number int `json:"number"`
}

type PrimeResponse struct {
	Message string `json:"message"`
	Result  bool   `json:"is_prime"`
}
