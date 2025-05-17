package common

type Error403Response struct {
	Error string `json:"error" example:"Access denied"`
}
