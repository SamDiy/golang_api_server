package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"server/models"
	"server/workdb"
)

func getModels(modelName string) interface{} {
	switch modelName {
	case "user":
		var users []models.User
		return &users
	case "article":
		var articles []models.Article
		return &articles
	case "comment":
		var comments []models.Comment
		return &comments
	}
	return nil
}

func getModel(modelName string) interface{} {
	switch modelName {
	case "user":
		var user models.User
		return &user
	case "article":
		var article models.Article
		return &article
	case "comment":
		var comment models.Comment
		return &comment
	}
	return nil
}

func getMapQuerry(querry url.Values) map[string]interface{} {
	mapQuerry := make(map[string]interface{})
	for key, value := range querry {
		_, _ = key, value
		mapQuerry[key] = value[0]
	}
	return mapQuerry
}

func getRequesResult(modelName string, querryMap map[string]interface{}) []uint8 {
	myModel := getModels(modelName)
	db := workdb.GetDB()
	if len(querryMap) != 0 {
		db.Where(querryMap).Find(myModel)
	} else {
		db.Find(myModel)
	}
	js, _ := json.Marshal(myModel)
	return js
}

func createNewElement(modelName string, body []uint8) []uint8 {
	db := workdb.GetDB()
	myModel := getModel(modelName)
	json.Unmarshal(body, myModel)
	db.Create(myModel)
	js, _ := json.Marshal(myModel)
	return js
}

func updateElement(modelName string, userID []string, body []uint8) []uint8 {
	db := workdb.GetDB()
	myModel := getModel(modelName)
	myNewModel := getModel(modelName)
	json.Unmarshal(body, myNewModel)
	db.First(myModel, "id = ?", userID[0])
	db.Model(myModel).Update(myNewModel)
	js, _ := json.Marshal(myModel)
	return js
}

func deleteElement(modelName string, userID []string) []uint8 {
	db := workdb.GetDB()
	myModel := getModel(modelName)
	db.Where("id = ?", userID[0]).Delete(myModel)
	js, _ := json.Marshal(userID[0])
	return js
}

// CreateServerController create model endpoints
func CreateServerController(modelName string) {

	http.HandleFunc(fmt.Sprintf("/api/%s", modelName), func(res http.ResponseWriter, req *http.Request) {
		var js []uint8
		switch req.Method {
		case "GET":
			querryMap := getMapQuerry(req.URL.Query())
			isEmptyQuerry := len(querryMap) == 0
			_ = isEmptyQuerry
			js = getRequesResult(modelName, querryMap)
		case "POST":
			body, _ := ioutil.ReadAll(req.Body)
			js = createNewElement(modelName, body)
		case "PUT":
			userID, isID := req.URL.Query()["id"]
			body, _ := ioutil.ReadAll(req.Body)
			if !isID {
				http.Error(res, "Not id", 500)
				return
			}
			js = updateElement(modelName, userID, body)
		case "DELETE":
			userID, isID := req.URL.Query()["id"]
			if !isID {
				http.Error(res, "Not id", 500)
				return
			}
			js = deleteElement(modelName, userID)
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(js)
	})
}
