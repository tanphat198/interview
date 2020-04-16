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
