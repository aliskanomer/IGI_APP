package controllers

import (
	// packages
	"IGI_API/internal/models"
	"IGI_API/internal/services"
	"IGI_API/internal/utils"
	"strconv"

	// modules
	"net/http"
)

// @Summary Get all planets
// @Description Retrieve paginated planet data from SWAPI.
// @Tags Planets
// @Accept  json
// @Produce  json
// @Param   page query int false "Page number (default is 1)"
// @Param   limit query int false "Items per page (default is 15)"
// @Success 200 {object} models.APIResponse{response=models.SWAPIList} "List of Planets with pagination details"
// @Failure 400 {object} models.APIResponse "Invalid Query Parameters"
// @Failure 500 {object} models.APIResponse "Internal Server Error"
// @Router /planet [get]
func GetPlanetAll(resp http.ResponseWriter, req *http.Request) {
	// generate and instance of Planet service *1
	planetService := services.NewPlanet()

	// Extract the page and limit from the URL query
	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(req.URL.Query().Get("limit"))

	// Assign default values if not provided or invalid
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 15
	}

	// fetch data by service layer
	data, err := planetService.GetPlanetAll(page, limit) //

	// Service failed. Send error response by utils
	if err != nil {
		utils.SendErrorResponse(resp, "ERR: Unable to fetch planet data!", http.StatusInternalServerError, models.ErrorInfo{
			Code:    "SERVICE_ERROR",
			Details: err.Error(),
		})
		return
	}

	// Data fetched. Send success response by utils
	utils.SendSuccessResponse(resp, "SCC: Planets fetched successfuly!", data, http.StatusOK)
}

// @Summary Get planet by ID
// @Description Retrieve a specific planet's data by its ID from SWAPI.
// @Tags Planets
// @Accept  json
// @Produce  json
// @Param   id path string true "Planet ID"
// @Success 200 {object} models.APIResponse{response=models.PlanetByIDResponse} "Planet Data"
// @Failure 400 {object} models.APIResponse "Invalid Planet ID"
// @Failure 500 {object} models.APIResponse "Internal Server Error"
// @Router /planet/{id} [get]
func GetPlanetById(resp http.ResponseWriter, req *http.Request) {
	// generate and instance of Planet service *1
	planetService := services.NewPlanet()

	// Extract the ID from the URL
	id := req.URL.Path[len("/planet/"):]

	// Validate the ID
	if id == "" {
		utils.SendErrorResponse(resp, "ERR: ID is required!", http.StatusBadRequest, models.ErrorInfo{
			Code:    "INVALID_REQUEST",
			Details: "Missing ID parameter in the URL.",
		})
		return
	}

	// Fetch data by service layer
	data, err := planetService.GetPlanetByID(id)

	// Service failed. Send error response by utils
	if err != nil {
		utils.SendErrorResponse(resp, "ERR: Unable to fetch planet data!", http.StatusInternalServerError, models.ErrorInfo{
			Code:    "SERVICE_ERROR",
			Details: err.Error(),
		})
		return
	}

	// Data fetched. Send success response by utils
	utils.SendSuccessResponse(resp, "Planet fetched successfuly", data, http.StatusOK)
}
