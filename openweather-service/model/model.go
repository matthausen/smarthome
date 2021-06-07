package model

type (
	CurrentWeatherData struct {
		Coordinates Coordinates `json:"coord"`
		Weather     []Weather   `json:"weather"`
		Main        Main        `json:"main"`
		Wind        Wind        `json:"wind"`
		Cod         int32       `json:"cod"`
	}

	Coordinates struct {
		Longitude float32 `json:"lon"`
		Latitude  float32 `json:"lat"`
	}

	Weather struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	}

	Main struct {
		Temp     float32 `json:"temp"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
		Humidity int32   `json:"humidity"`
	}

	Wind struct {
		Speed     float32 `json:"speed"`
		Direction int32   `json:"deg"`
	}

	ForecastData struct {
		List []CurrentWeatherData `json:"list"`
	}
)
