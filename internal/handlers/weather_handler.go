package handlers

import (
	"log"
	"net/http"
	"strconv"
	"weather-api/internal/config"
	"weather-api/internal/models"
	"weather-api/internal/services"

	"github.com/gin-gonic/gin"
)

// GetWeatherHandler godoc
// @Summary Get weather conditions
// @Description Get weather conditions and temperature category based on latitude and longitude.
// @Tags weather
// @Accept json
// @Produce json
// @Param lat query number true "Latitude"
// @Param lon query number true "Longitude"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /weather [get]
func GetWeatherHandler(cfg *config.Config) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Parse latitude and validate it to be a float
        latStr := c.Query("lat")
        lat, err := strconv.ParseFloat(latStr, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, models.ErrorResponse{
                Status:  http.StatusBadRequest,
                Message: "Invalid latitude value. It should be a valid float number.",
            })
            return
        }

        // Parse longitude and validate it to be a float
        lonStr := c.Query("lon")
        lon, err := strconv.ParseFloat(lonStr, 64)
        if err != nil {
            c.JSON(http.StatusBadRequest, models.ErrorResponse{
                Status:  http.StatusBadRequest,
                Message: "Invalid longitude value. It should be a valid float number.",
            })
            return
        }

        // Initialize the WeatherService with API key, base URL, and temperature thresholds
        weatherService := services.NewWeatherService(cfg.OpenWeather.APIKey, cfg.OpenWeather.BaseURL, cfg.TemperatureThresholds.ColdMax, cfg.TemperatureThresholds.ModerateMax)
        
        // Call the GetWeather method of WeatherService to fetch weather data for the given lat and lon
        weatherResponse, err := weatherService.GetWeather(lat, lon)
        if err != nil {
            log.Fatalf(err.Error())
            c.JSON(http.StatusInternalServerError, models.ErrorResponse{
                Status:  http.StatusInternalServerError,
                Message: "Failed to fetch weather data",
            })
            return
        }

        // Categorize the temperature into cold, moderate, or hot based on the temperature thresholds
        tempCategory := weatherService.CategorizeTemperature(weatherResponse.Main.Temp)
        weatherCondition := weatherResponse.Weather[0].Main

        // Construct the successful API response with weather details and send it with a 200 OK status.
        // For now, additionally added partial metadata fetched from open weather API
        // Actual required fileds are: temperature_category & weather_condition
        c.JSON(http.StatusOK, models.APIResponse{
            Status:  http.StatusOK,
            Message: "Success",
            Data: map[string]interface{}{
                "temperature_category": tempCategory,
                "weather_condition":    weatherCondition,
                "temperature":          weatherResponse.Main.Temp,
                "humidity":             weatherResponse.Main.Humidity,
                "pressure":             weatherResponse.Main.Pressure,
                "wind_speed":           weatherResponse.Wind.Speed,
                "city":                 weatherResponse.Name,
            },
        })
    }
}