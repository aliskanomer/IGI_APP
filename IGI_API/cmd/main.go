package main

import (
	// packages
	"IGI_API/internal/routes"
	"IGI_API/internal/utils"

	// modules
	"fmt"
	"net/http"

	// Swagger
	_ "IGI_API/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title IGI_API
// @version 1.0.1
// @description This is the IGI API server documentation.
// @host localhost:8080
// @BasePath /
func main() {
	// Load environment variables
	utils.LoadEnv()

	// Initialize routes
	routes.InitRoutes()
	utils.Logger("info", "Server", 0, fmt.Sprintf("IGI_API is serving at %s", utils.ReadEnvVar("BASE_URL")))

	// Swagger UI route
	http.Handle("/docs/", httpSwagger.WrapHandler)
	utils.Logger("info", "Server", 0, fmt.Sprintf("Documentation: %s", utils.ReadEnvVar("SWAGGER_URL")))

	// Start the server and log any errors
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		utils.Logger("error", "Server", http.StatusInternalServerError, fmt.Errorf("server failed: %v", err))
	}
}
