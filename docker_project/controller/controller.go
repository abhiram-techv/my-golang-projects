package controller

import (
	"docker_project/config"
	"docker_project/model"
	"encoding/json"
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from weather app \n"))
}

func ShowWeather(city string) (model.WeatherData, error) {
	apiConfig, err := config.LoadApiConfig(".apiConfig")
	if err != nil {
		return model.WeatherData{}, err
	}

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiConfig.OpenWeatherMapApiKey)

	if err != nil {
		return model.WeatherData{}, err
	}
	defer resp.Body.Close()

	var d model.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return model.WeatherData{}, err
	}
	db, err := config.Connect()
	if err != nil {
		log.Print(err)
		return model.WeatherData{}, err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO weather(name) VALUES(?)", d.Name)
	if err != nil {
		log.Print(err)
		return model.WeatherData{}, err
	}
	return d, nil
}
