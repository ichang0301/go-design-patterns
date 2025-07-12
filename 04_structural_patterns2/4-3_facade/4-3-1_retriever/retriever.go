package http_rest_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float64) (Weather, error)
}

type Weather struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Cod   int    `json:"cod"`
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	}
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`

	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`

	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise uint32 `json:"sunrise"`
		Sunset  uint32 `json:"sunset"`
	} `json:"sys"`
}

type CurrentWeatherData struct {
	APIKey string
}

func (d *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (d *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return d.doRequest(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%f,%f&APPID=%s", lat, lon, d.APIKey))
}

func (d *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return d.doRequest(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&APPID=%s", city, countryCode, d.APIKey))
}

func (d *CurrentWeatherData) doRequest(url string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		byt, errMsg := io.ReadAll(resp.Body)
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("status code was %d, aborting. Error message was:\n%s", resp.StatusCode, errMsg)

		return
	}

	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return
	}

	weather, err = d.responseParser(resp.Body)

	return
}
