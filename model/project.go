package model

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Member      Member `json:"-"`
	MemberID    uint   `json:"memberId" binding:"required"`
}

func (p *Project) TableName() string {
	return "project"
}
