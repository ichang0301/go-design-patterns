package http_rest_api

import (
	"bytes"
	"io"
	"testing"
)

func getMockData() io.Reader {
	response := `{
		"coord": {
			"lon": -3.7,
			"lat": 40.42
		},
		"weather": [
			{
				"id": 803,
				"main": "Clouds",
				"description": "broken clouds",
				"icon": "04n"
			}
		],
		"base": "stations",
		"main": {
			"temp": 303.56,
			"pressure": 1016.46,
			"humidity": 26.8,
			"temp_min": 300.95,
			"temp_max": 305.93
		},
		"wind": {
			"speed": 3.17,
			"deg": 151.001
		},
		"rain": {
			"3h": 0.0075
		},
		"clouds": {
			"all": 68
		},
		"dt": 1471295823,
		"sys": {
			"type": 3,
			"id": 1442829648,
			"message": 0.0278,
			"country": "ES",
			"sunrise": 1471238808,
			"sunset": 1471288232
		},
		"id": 3117735,
		"name": "Madrid",
		"cod": 200
	}`

	r := bytes.NewReader([]byte(response))
	return r
}

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{
		APIKey: "",
	}

	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}

	if weather.ID != 3117735 {
		t.Errorf("Expected ID 3117735, not %d\n", weather.ID)
	}
}

// func TestOpenWeatherMap_retrieveTemperature(t *testing.T) {
// 	apiKey := "your_api_key_here"	// Replace with your actual API key
// 	weatherMap := CurrentWeatherData{apiKey}
// 	weather, err := weatherMap.GetByCityAndCountryCode("Madrid", "ES")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Printf("Temperature in Madrid is %f celsius\n", weather.Main.Temp-273.15)
// }
