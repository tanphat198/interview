package repo

import (
	"ProjectManagement/config"
	"ProjectManagement/model"
)

func CreateMember(member *model.Member) error {
	if err := config.DB.Create(member).Error; err != nil {
		return err
	}
	return nil
}

func UpdateMember(member *model.Member) error {
	if err := config.DB.Save(member).Error; err != nil {
		return err
	}
	return nil
}
