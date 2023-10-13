package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/resty.v1"
)

type WeatherData struct {
	Temperature string
	Wind        string
	Humidity    string
}

func main() {
	http.HandleFunc("/weather", WeatherHandler)
	http.ListenAndServe(":8080", nil)
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "city is not defined", http.StatusBadRequest)
		return
	}

	// i am going to use free (no auth) wttr api, luckly it has all needed parameters and responce only has what we need
	// in this case i don't need to parse whole json responce and get data what we need.

	url := fmt.Sprintf("https://wttr.in/%s?format=%%t,%%w,%%h", city)
	resp, err := resty.R().Get(url)

	if err != nil {
		http.Error(w, "HTTP error during request", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode() != 200 {
		http.Error(w, "error getting response", resp.StatusCode())
		return
	}

	// simple request to get weather in Kyiv returns three values : temp, wind and humidity
	// https://wttr.in/Kyiv?format=%t,%w,%h
	data := strings.Split(resp.String(), ",")

	weather := WeatherData{
		Temperature: data[0],
		Wind:        data[1],
		Humidity:    data[2],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
