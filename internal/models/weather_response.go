package models

type WeatherResponse struct {
    Coord struct {
        Lon float64 `json:"lon"`
        Lat float64 `json:"lat"`
    } `json:"coord"`
    Weather []struct {
        Main        string `json:"main"`
        Description string `json:"description"`
        Icon        string `json:"icon"`
    } `json:"weather"`
    Main struct {
        Temp     float64 `json:"temp"`
        FeelsLike float64 `json:"feels_like"`
        Pressure int     `json:"pressure"`
        Humidity int     `json:"humidity"`
        TempMin  float64 `json:"temp_min"`
        TempMax  float64 `json:"temp_max"`
    } `json:"main"`
    Wind struct {
        Speed float64 `json:"speed"`
        Deg   int     `json:"deg"`
    } `json:"wind"`
    Clouds struct {
        All int `json:"all"`
    } `json:"clouds"`
    Sys struct {
        Country string `json:"country"`
        Sunrise int64  `json:"sunrise"`
        Sunset  int64  `json:"sunset"`
    } `json:"sys"`
    Name string `json:"name"`
}


// APIResponse is a general structure for successful API responses
type APIResponse struct {
    Status  int         `json:"status"`  // HTTP status code
    Message string      `json:"message"` // Message (success or error)
    Data    interface{} `json:"data,omitempty"` // Response Data
}

// ErrorResponse is a structure for error responses
type ErrorResponse struct {
    Status  int    `json:"status"`  // HTTP status code
    Message string `json:"message"` // Error message
}