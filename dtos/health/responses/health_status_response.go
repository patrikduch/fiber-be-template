package responses

type HealthStatusResponse struct {
	App string `json:"app" example:"OK"`
	DB  string `json:"db" example:"OK"`
}
