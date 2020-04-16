package controllers

import (
	"ProjectManagement/server/model"
	"ProjectManagement/server/repo"
	"ProjectManagement/server/utils"
	"log"

	"github.com/gin-gonic/gin"
)

var CreateMember = func(c *gin.Context) {
	member := model.Member{}
	err := c.ShouldBindJSON(&member)
	if err != nil {
		utils.Respond(c.Writer, utils.Message(false, "Error while decoding request body"))
		return
	}
	err = repo.CreateMember(&member)
	if err != nil {
		log.Print("can not create project", err)
	}
	resp := utils.Message(true, "success")
	resp["data"] = member
	utils.Respond(c.Writer, resp)
}

var GetAllMembers = func(c *gin.Context) {
	var members []model.Member
	err := repo.GetAllMember(&members)
	if err != nil {
		log.Print("can not get all members", err)
	}

	resp := utils.Message(true, "success")
	resp["data"] = members
	utils.Respond(c.Writer, resp)
}
