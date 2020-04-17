package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Name     string    `json:"name" binding:"required"`
	Phone    string    `json:"phone" binding:"required"`
	Birthday time.Time `json:"birthday" binding:"required"`
}

func (m *Member) TableName() string {
	return "member"
}
