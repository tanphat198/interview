package model

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Member      *Member `json:"-"`
	MemberId    int64   `json:"memberId"`
}

func (p *Project) TableName() string {
	return "project"
}
