package controllers

import (
	"ProjectManagement/model"
	"ProjectManagement/repo"
	"capstone/go-contacts/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func GetProjectWithName(c *gin.Context) {
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
