package main

import (
	"docker_project/controller"
	"encoding/json"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", controller.Test)
	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := controller.ShowWeather(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(data)
		})
	http.ListenAndServe(":8090", nil)
}
