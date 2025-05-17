package common

type Error400Response struct {
	Error string `json:"error" example:"Bad request"`
}
