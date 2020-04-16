package repo

import (
	"ProjectManagement/server/config"
	"ProjectManagement/server/model"
)

func CreateMember(member *model.Member) error {
	if err := config.GetDB().Create(member).Error; err != nil {
		return err
	}
	return nil
}

func UpdateMember(member *model.Member) error {
	if err := config.GetDB().Save(member).Error; err != nil {
		return err
	}
	return nil
}

func GetAllMember(member *[]model.Member) error {
	if err := config.GetDB().Find(member).Error; err != nil {
		return err
	}
	return nil
}
