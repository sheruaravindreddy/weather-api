package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
    Server struct {
        Port string `mapstructure:"port"`
    } `mapstructure:"server"`
    OpenWeather struct {
        APIKey  string `mapstructure:"api_key"`
        BaseURL string `mapstructure:"base_url"`
    } `mapstructure:"openweather"`
    TemperatureThresholds struct {
        ColdMax     float64 `mapstructure:"cold_max"`
        ModerateMax float64 `mapstructure:"moderate_max"`
    } `mapstructure:"temperature_thresholds"`
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config") // name of config file (without extension)
    viper.SetConfigType("yaml")   // config file extension
    viper.AddConfigPath(".")      // look for config in the current directory
    viper.AddConfigPath("internal/config") // look for config in the internal/config directory

    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("error reading config file: %w", err)
    }

    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf("unable to decode into struct: %w", err)
    }

    // Load the API key from the environment variable
    cfg.OpenWeather.APIKey = os.Getenv(viper.GetString("openweather.api_key_env"))
    if cfg.OpenWeather.APIKey == "" {
        return nil, fmt.Errorf("API key not set in environment")
    }

    return &cfg, nil
}