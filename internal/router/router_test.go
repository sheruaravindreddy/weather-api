package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"weather-api/internal/config"
	"weather-api/internal/router"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
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

	// Mock configuration
	cfg := &config.Config{
		Server: struct {
			Port string "mapstructure:\"port\""
		}{
			Port: ":8080",
		},
		OpenWeather: struct {
			APIKey  string "mapstructure:\"api_key\""
			BaseURL string "mapstructure:\"base_url\""
		}{
			APIKey:  "fake-api-key",
			BaseURL: mockServer.URL,
		},
		TemperatureThresholds: struct {
			ColdMax     float64 "mapstructure:\"cold_max\""
			ModerateMax float64 "mapstructure:\"moderate_max\""
		}{
			ColdMax:     10,
			ModerateMax: 25,
		},
	}

	// Create a new Gin engine for testing
	gin.SetMode(gin.TestMode)
	r := router.SetupRouter(cfg)

	// Test the /weather route with valid lat and lon
	req, _ := http.NewRequest("GET", "/weather?lat=35.0&lon=139.0", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Clear")
}