package service

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	model "github.com/matthausen/weather-service/model"
)

var (
	API_KEY            = os.Getenv("OPENWEATHER_API_KEY")
	city               = os.Getenv("OPENWEATHER_CITY")
	currentWeatherURL  = "https://api.openweathermap.org/data/2.5/weather?q=" + city + ",uk&APPID=" + API_KEY
	forecastWeatherURL = "https://api.openweathermap.org/data/2.5/forecast?q=" + city + "&appid=" + API_KEY
)

func currentWeatherData() (model.CurrentWeatherData, error) {
	resp, err := http.Get(currentWeatherURL)

	var currentWeatherData model.CurrentWeatherData

	if err != nil {
		log.Fatalf("Could not fetch weather data %v", err)
		return currentWeatherData, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&currentWeatherData); err != nil {
		log.Printf("Could not decode response body: %v", err)
		return currentWeatherData, err
	}

	return currentWeatherData, nil
}

func forecastWeatherData() (model.ForecastData, error) {
	resp, err := http.Get(forecastWeatherURL)

	var forecastWeatherData model.ForecastData

	if err != nil {
		log.Fatalf("Could not fetch forecasted data %v", err)
		return forecastWeatherData, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&forecastWeatherData); err != nil {
		log.Printf("There was an error decoding the response %v", err)
		return forecastWeatherData, err
	}

	return forecastWeatherData, nil
}
