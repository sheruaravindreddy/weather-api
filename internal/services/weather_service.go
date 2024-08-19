package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-api/internal/models"
	"weather-api/pkg/httpclient"
)

type WeatherService struct {
    apiKey              string
    baseURL             string
    httpClient          *http.Client
    coldMax             float64
    moderateMax         float64
}

// NewWeatherService creates a new instance of WeatherService.
// It initializes the service with the provided API key, base URL, and temperature thresholds.
func NewWeatherService(apiKey, baseURL string, coldMax, moderateMax float64) *WeatherService {
    return &WeatherService{
        apiKey:      apiKey,
        baseURL:     baseURL,
        httpClient:  httpclient.NewHTTPClient(),
        coldMax:     coldMax,
        moderateMax: moderateMax,
    }
}

// GetWeather fetches the current weather data from the OpenWeather API
// based on the provided latitude and longitude coordinates.
// It returns the weather data or an error if the request fails.
func (ws *WeatherService) GetWeather(lat, lon float64) (*models.WeatherResponse, error) {
    url := fmt.Sprintf("%s?lat=%f&lon=%f&units=metric&appid=%s", ws.baseURL, lat, lon, ws.apiKey)

    resp, err := ws.httpClient.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var weatherResponse models.WeatherResponse
    if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
        return nil, err
    }

    return &weatherResponse, nil
}

// CategorizeTemperature categorizes the temperature into "cold," "moderate," or "hot"
// based on the thresholds provided during the WeatherService initialization.
// Thresholds can be updated in conf.yaml
func (ws *WeatherService) CategorizeTemperature(temp float64) string {
    switch {
    case temp <= ws.coldMax:
        return "cold"
    case temp > ws.coldMax && temp <= ws.moderateMax:
        return "moderate"
    default:
        return "hot"
    }
}