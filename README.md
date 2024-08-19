
# Weather API

This project is a Go HTTP server that retrieves weather conditions based on latitude and longitude coordinates using the OpenWeather API. The API also categorizes temperatures as hot, cold, or moderate.

## Prerequisites

Before running the server, make sure you have the following:

- [Go](https://golang.org/doc/install) installed on your system.
- An API key from [OpenWeather](https://openweathermap.org/api).

## Setup

### 1. Clone the Repository

First, clone the repository and move into the project directory:

```bash
git clone git@github.com:sheruaravindreddy/weather-api.git
cd weather-api
```

### 2. Set Up Environment Variable

Next, set up your OpenWeather API key as an environment variable:

```bash
export OPENWEATHER_API_KEY='***********************'
```

You can validate that your API key is set correctly by running:

```bash
echo $OPENWEATHER_API_KEY
```

### 3. Run the Server

To start the Go HTTP server, run the following command:

```bash
go run cmd/server/main.go
```

The server will start and listen on port `8080`.

## Usage

### Accessing the API Documentation

You can access the Swagger UI to interact with the API by navigating to the following URL in your web browser:

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Making API Calls

1. Once on the Swagger UI, find the `GET /weather` endpoint.
2. Click on **Try it out**.
3. Enter the latitude and longitude values for the location you wish to query.
4. Click **Execute** to view the weather results and temperature categorization.

### Example Request

- **Latitude:** `37.7749`
- **Longitude:** `-122.4194`

### Example Response

```json
{
  "status": 200,
  "message": "Success",
  "data": {
    "city": "San Francisco",
    "humidity": 88,
    "pressure": 1016,
    "temperature": 13.67,
    "temperature_category": "moderate",
    "weather_condition": "Clouds",
    "wind_speed": 6.17
  }
}
```
Note: Although the requirement is just to return temperature_category and weather_condition, I've additionally added metadata returned from OpenWeather API

## Unit Tests
![image](https://github.com/user-attachments/assets/95199272-ce91-4865-a0e0-9bfcfdbc8268)


## Acknowledgments

- [OpenWeather API](https://openweathermap.org/api) for providing weather data.
- [Swagger UI](https://swagger.io/tools/swagger-ui/) for API documentation.
