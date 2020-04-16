package controllers

import (
	"ProjectManagement/server/model"
	"ProjectManagement/server/repo"
	"ProjectManagement/server/utils"
	"log"

	"github.com/gin-gonic/gin"
)

var GetProjectWithName = func(c *gin.Context) {
	name := c.Params.ByName("name")
	var project *model.Project
	err := repo.GetProjectWithName(project, name)
	if err != nil {
		log.Print("can not get all public brands", err)
	}
	resp := utils.Message(true, "success")
	resp["data"] = project
	utils.Respond(c.Writer, resp)
}

var CreateProject = func(c *gin.Context) {
	project := model.Project{}
	err := c.ShouldBindJSON(&project)
	if err != nil {
		utils.Respond(c.Writer, utils.Message(false, "Error while decoding request body"))
		return
	}
	err = repo.CreateProject(&project)
	if err != nil {
		log.Print("can not create project", err)
	}
	resp := utils.Message(true, "success")
	resp["data"] = project
	utils.Respond(c.Writer, resp)
}

var GetAllProjects = func(c *gin.Context) {
	var projects []model.Project
	err := repo.GetAllProject(&projects)
	if err != nil {
		log.Print("can not get all members", err)
	}

	resp := utils.Message(true, "success")
	resp["data"] = projects
	utils.Respond(c.Writer, resp)
}

var UpdateProject = func(c *gin.Context) {
	project := model.Project{}

	err := c.ShouldBindJSON(&project)
	if err != nil {
		utils.Respond(c.Writer, utils.Message(false, "Error while decoding request body"))
		return
	}

	err = repo.UpdateProject(&project)
	if err != nil {
		log.Print("can not update brand", err)
	}
	resp := utils.Message(true, "success")
	resp["data"] = project
	utils.Respond(c.Writer, resp)
}
