package models

type AddRequest struct {
	Numbers []int `json:"numbers"`
}

type AddResponse struct {
	Message string `json:"message"`
	Result  int    `json:"sum"`
}
