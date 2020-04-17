package main

import (
	"ProjectManagement/server/config"
	"ProjectManagement/server/controllers"
	"ProjectManagement/server/model"
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config.GetDB().AutoMigrate(model.Project{}, model.Member{})
	r := gin.Default()
	api := r.Group("/api")
	{
		api.POST("/projects/new", controllers.CreateProject)
		api.POST("/members/new", controllers.CreateMember)
		api.GET("/members", controllers.GetAllMembers)
		api.GET("/projects", controllers.GetAllProjects)
		api.PUT("/projects/:projectId", controllers.UpdateProject)
		api.PUT("/members/:memberId", controllers.UpdateMember)
	}

	// Start and run the server
	fmt.Print("running on localhost:9000")
	err := r.Run(":9000")
	if err != nil {
		fmt.Print(err)
	}
}
