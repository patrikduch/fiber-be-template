package models

import (
	"github.com/google/uuid"
)

type Role struct {
	ID   uuid.UUID `gorm:"column:Id;type:uuid;primaryKey" json:"id"`
	Name string    `gorm:"column:Name;uniqueIndex" json:"name"`
}

func (Role) TableName() string {
	return "Role"
}