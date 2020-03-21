package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server/models"
	"server/workdb"
)

func getJSONUsers(userID []string, isID bool) []uint8 {
	db := workdb.GetDB()
	var users []models.User
	if isID {
		db.Find(&users, "id = ?", userID[0])
	} else {
		db.Find(&users)
	}
	js, _ := json.Marshal(users)
	return js
}

func createNewUser(body []uint8) []uint8 {
	db := workdb.GetDB()
	var user models.User
	json.Unmarshal(body, &user)
	db.Create(&user)
	js, _ := json.Marshal(user)
	return js
}

func updateUser(userID []string, body []uint8) []uint8 {
	db := workdb.GetDB()
	var user models.User
	var newUser models.User
	json.Unmarshal(body, &newUser)
	db.First(&user, "id = ?", userID[0])
	db.Model(&user).Update(newUser)
	js, _ := json.Marshal(user)
	return js
}

func deleteUser(userID []string) []uint8 {
	db := workdb.GetDB()
	db.Where("id = ?", userID[0]).Delete(&models.User{})
	js, _ := json.Marshal(userID[0])
	return js
}

// RunUserControllers create user endpoints
func RunUserControllers() {

	// db := workdb.GetDB()

	http.HandleFunc("/api/user", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			userID, isID := req.URL.Query()["id"]
			js := getJSONUsers(userID, isID)
			res.Header().Set("Content-Type", "application/json")
			res.Write(js)
		case "POST":
			body, _ := ioutil.ReadAll(req.Body)
			js := createNewUser(body)
			res.Header().Set("Content-Type", "application/json")
			res.Write(js)
		case "PUT":
			userID, isID := req.URL.Query()["id"]
			body, _ := ioutil.ReadAll(req.Body)
			if !isID {
				http.Error(res, "Not id", 500)
				return
			}
			js := updateUser(userID, body)
			res.Header().Set("Content-Type", "application/json")
			res.Write(js)
		case "DELETE":
			userID, isID := req.URL.Query()["id"]
			if !isID {
				http.Error(res, "Not id", 500)
				return
			}
			js := deleteUser(userID)
			res.Header().Set("Content-Type", "application/json")
			res.Write(js)
		}
	})
}
