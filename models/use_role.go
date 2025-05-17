package models

import (
	"github.com/google/uuid"
)

type UserRole struct {
	UserId uuid.UUID `gorm:"column:UserId;type:uuid"`
	RoleId uuid.UUID `gorm:"column:RoleId;type:uuid"`
}

func (UserRole) TableName() string {
	return "UserRoles"
}