package common

type Error500Response struct {
	Error string `json:"error" example:"Internal server error"`
}
