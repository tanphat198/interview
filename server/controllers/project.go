package controllers

import (
	"ProjectManagement/server/model"
	"ProjectManagement/server/repo"
	"ProjectManagement/server/utils"
	"fmt"
	"log"
	"strconv"

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
	memberId := c.Params.ByName("memberId")
	id, err := strconv.ParseInt(memberId, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", id, id)
	}
	project := model.Project{}
	err = c.ShouldBindJSON(&project)
	if err != nil {
		utils.Respond(c.Writer, utils.Message(false, "Error while decoding request body"))
		return
	}
	if memberId != "" {
		member := model.Member{}
		err := repo.GetMemberWithId(&member, id)
		if err != nil {
			utils.Respond(c.Writer, utils.Message(false, "Can not find this member"))
			return
		}
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
	//memberId := c.Params.ByName("memberId")
	projectId := c.Params.ByName("projectId")
	pId, err := strconv.ParseInt(projectId, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", pId, pId)
	}
	//id, err := strconv.ParseInt(memberId, 10, 64)
	//if err == nil {
	//	fmt.Printf("%d of type %T", id, id)
	//}
	project := model.Project{}

	err = c.ShouldBindJSON(&project)
	if err != nil {
		utils.Respond(c.Writer, utils.Message(false, "Error while decoding request body"))
		return
	}

	if project.MemberId != -1 {
		member := model.Member{}
		err := repo.GetMemberWithId(&member, project.MemberId)
		if err != nil {
			utils.Respond(c.Writer, utils.Message(false, "Can not find this member"))
			return
		}
	}
	project.ID = uint(pId)
	err = repo.UpdateProject(pId, project.MemberId)
	if err != nil {
		log.Print("can not update project", err)
	}
	resp := utils.Message(true, "success")
	resp["data"] = project
	utils.Respond(c.Writer, resp)
}
