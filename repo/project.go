package repo

import (
	"ProjectManagement/config"
	"ProjectManagement/model"
)

func CreateProject(project *model.Project) error {
	if err := config.DB.Create(project).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProject(project *model.Project) error {
	if err := config.DB.Save(project).Error; err != nil {
		return err
	}
	return nil
}

func GetProjectWithName(project *model.Project, name string) error {
	if err := config.DB.Where("name = ?", name).First(project).Error; err != nil {
		return err
	}
	return nil
}
