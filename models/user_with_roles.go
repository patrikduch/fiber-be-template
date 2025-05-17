package models

import "github.com/google/uuid"

type UserWithRoles struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"username"`
	Email string    `json:"email"`
	Roles []Role    `json:"roles"`
}