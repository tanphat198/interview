package main

import (
	"ProjectManagement/config"
	"ProjectManagement/model"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
)

func main() {
	DB := config.GetDB()
	DB.AutoMigrate(&model.Member{}, &model.Project{})

	// Initialize
	Admin := admin.New(&admin.AdminConfig{DB: DB})

	// Allow to use Admin to manage Project, Member
	Admin.AddResource(&model.Project{})
	Admin.AddResource(&model.Member{})

	// initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// Mount admin interface to mux
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000, go to http://localhost:9000/admin")
	http.ListenAndServe(":9000", mux)
}
