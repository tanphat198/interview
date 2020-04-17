package repo

import (
	"ProjectManagement/server/config"
	"ProjectManagement/server/model"
)

func CreateProject(project *model.Project) error {
	if err := config.GetDB().Create(project).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProject(pId int64, id int64) error {
	if err := config.GetDB().Table("project").Where("id = ?", pId).Update("member_id", id).Error; err != nil {
		return err
	}
	return nil
}

func GetProjectWithName(project *model.Project, name string) error {
	if err := config.GetDB().Where("name = ?", name).First(project).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProject(project *[]model.Project) error {
	if err := config.GetDB().Find(project).Error; err != nil {
		return err
	}
	return nil
}
