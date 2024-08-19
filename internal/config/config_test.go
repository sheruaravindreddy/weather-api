package config_test

import (
	"os"
	"testing"
	"weather-api/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("OPENWEATHER_API_KEY", "test-api-key") // mocking env variable

	cfg, err := config.LoadConfig()

	assert.NoError(t, err)
	assert.Equal(t, "test-api-key", cfg.OpenWeather.APIKey)
	assert.Equal(t, "https://api.openweathermap.org/data/2.5/weather", cfg.OpenWeather.BaseURL)
	assert.Equal(t, ":8080", cfg.Server.Port)
}