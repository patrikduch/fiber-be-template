package requests

type CreateUserRequestDto struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}