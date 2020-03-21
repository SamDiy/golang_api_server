package controllers

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	Name    string
	Age     int
	Hobbies []string
}

func RunInitControllers() {

	http.HandleFunc("/api/init", func(res http.ResponseWriter, req *http.Request) {
		profile := Profile{"Sam", 30, []string{"gaming", "programming"}}
		js, err := json.Marshal(profile)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(js)
		// fmt.Fprintf(res, req.Method)
	})

}
