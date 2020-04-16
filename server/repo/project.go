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

func UpdateProject(project *model.Project) error {
	if err := config.GetDB().Save(project).Error; err != nil {
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
