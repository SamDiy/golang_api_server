package main

import (
	"fmt"
	"net/http"
	"server/controllers"
	"server/models"
	"server/workdb"
)

func initAllModels() {
	db := workdb.GetDB()
	db.AutoMigrate(&models.User{}, &models.Article{}, &models.Comment{})
}

func main() {

	workdb.InitDB()
	db := workdb.GetDB()
	initAllModels()
	defer db.Close()

	controllers.CreateServerController("user")
	controllers.CreateServerController("article")
	controllers.CreateServerController("comment")

	fmt.Println("Server run on port 8080")
	http.ListenAndServe(":8080", nil)
}
