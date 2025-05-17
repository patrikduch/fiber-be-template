package responses

type AuthMeResponseDto struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Role      string `json:"role,omitempty"`      
	CreatedAt string `json:"createdAt,omitempty"`
}
