package service

import (
	"encoding/json"
	"log"
	"net/http"
)

// enableCors -> Allow request form localhost. To be deleted
func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
}

func CurrentWeatherHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	w.Header().Set("Content-Type", "application/json")

	data, err := currentWeatherData()
	if err != nil {
		log.Printf("Error while fetching current weather data %v", err)
	}

	json.NewEncoder(w).Encode(data)
}

func FiveDaysForecastHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	w.Header().Set("Content-Type", "application/json")

	data, err := forecastWeatherData()
	if err != nil {
		log.Printf("Error fetching 5 days forecast data %v", err)
	}

	json.NewEncoder(w).Encode(data)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	w.Header().Set("Content-Type", "application/json")

	res := []byte(`{ "alive": true}`)
	var raw map[string]interface{}
	if err := json.Unmarshal(res, &raw); err != nil {
		log.Fatalf("Could not unmarshal raw to json %v", err)
	}
	out, _ := json.Marshal(raw)

	w.Write(out)
}
