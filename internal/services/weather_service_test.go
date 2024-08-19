package services_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"weather-api/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestGetWeather(t *testing.T) {
	// Create a mock server to simulate OpenWeather API
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
			"weather": [{"main": "Clear"}],
			"main": {"temp": 20.0, "pressure": 1013, "humidity": 70},
			"wind": {"speed": 5.0},
			"sys": {"country": "US"},
			"name": "Test City"
		}`))
	}))
	defer mockServer.Close()

	// Initialize the service with the mock server URL
	weatherService := services.NewWeatherService("fake-api-key", mockServer.URL, 10, 25)

	weather, err := weatherService.GetWeather(35.0, 139.0)
	assert.NoError(t, err)
	assert.Equal(t, "Clear", weather.Weather[0].Main)
	assert.Equal(t, 20.0, weather.Main.Temp)
	assert.Equal(t, "Test City", weather.Name)
}

func TestCategorizeTemperature(t *testing.T) {
	weatherService := services.NewWeatherService("fake-api-key", "http://fake-url", 10, 25)

	tests := []struct {
		temp     float64
		expected string
	}{
		{5, "cold"},
		{15, "moderate"},
		{30, "hot"},
	}

	for _, test := range tests {
		result := weatherService.CategorizeTemperature(test.temp)
		assert.Equal(t, test.expected, result)
	}
}