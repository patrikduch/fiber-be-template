package models

import (
	"github.com/google/uuid"
)

type User struct {
    ID                   uuid.UUID `json:"id"`
    Name                 string    `json:"username"`
    Email                string    `json:"email"`
    NormalizedEmail      string    `json:"normalizedEmail"`
    Password             string    `json:"-"`
    EmailConfirmed       bool      `json:"emailConfirmed"`
    PhoneNumberConfirmed bool      `json:"phoneNumberConfirmed"`
    TwoFactorEnabled     bool      `json:"twoFactorEnabled"`
    LockoutEnabled       bool      `json:"lockoutEnabled"`
    AccessFailedCount    int       `json:"accessFailedCount"`
    Role                 *Role     `json:"role,omitempty"` // 👈 Add this line
}