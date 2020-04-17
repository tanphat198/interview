package repo

import (
	"ProjectManagement/server/config"
	"ProjectManagement/server/model"
	"time"
)

func CreateMember(member *model.Member) error {
	if err := config.GetDB().Create(member).Error; err != nil {
		return err
	}
	return nil
}

func GetMemberWithId(member *model.Member, id int64) error {
	if err := config.GetDB().Table("member").Where("id = ?", id).First(member).Error; err != nil {
		return err
	}
	return nil
}

func UpdateMemberWithId(member *model.Member, id int64) error {
	if err := config.GetDB().Table("member").Where("id = ?", id).Update("name", member.Name).Update("phone", member.Phone).Update("birthday", member.Birthday).Update("updated_at", time.Now()).Error; err != nil {
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
