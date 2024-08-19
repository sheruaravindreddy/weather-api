package main

import (
	"log"
	_ "weather-api/docs"
	"weather-api/internal/config"
	"weather-api/internal/router"
)

//	@title			Weather API
//	@version		1.0
//	@description	This is a sample server for weather conditions.
//	@host			localhost:8080
//	@BasePath		/
func main() {
    // load configuration from local env and config
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Initialize the router with the loaded configuration above
    r := router.SetupRouter(cfg)

    // Start the golang backend weather-api server
    log.Printf("Server running on port %s\n", cfg.Server.Port)
    r.Run(cfg.Server.Port)
}