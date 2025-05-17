package common

type Error404Response struct {
	Error string `json:"error" example:"Resource not found"`
}
