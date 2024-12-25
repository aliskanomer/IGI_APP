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

// @Summary Get all people
// @Description Retrieve all people data from SWAPI with pagination support.
// @Tags People
// @Accept  json
// @Produce  json
// @Param   page query int false "Page number (default: 1)"
// @Param   limit query int false "Items per page (default: 15)"
// @Success 200 {object}  models.APIResponse{response=models.SWAPIList}  "List of People with pagination"
// @Failure 500 {object} models.APIResponse "Internal Server Error"
// @Router /people [get]
func GetPeopleAll(resp http.ResponseWriter, req *http.Request) {
	// generate an instance of People service
	peopleService := services.NewPeople()

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
	data, err := peopleService.GetPeopleAll(page, limit)

	// Service failed. Send error response by utils
	if err != nil {
		utils.SendErrorResponse(resp, "ERR: Unable to fetch people data!", http.StatusInternalServerError, models.ErrorInfo{
			Code:    "SERVICE_ERROR",
			Details: err.Error(),
		})
		return
	}

	// Data fetched. Send success response by utils
	utils.SendSuccessResponse(resp, "SCC: People fetched successfully!", data, http.StatusOK)
}

// @Summary Get person by ID
// @Description Retrieve a specific person's data by their ID from SWAPI.
// @Tags People
// @Accept  json
// @Produce  json
// @Param   id path string true "Person ID"
// @Success 200 {object} models.APIResponse{response=models.PeopleByIDResponse} "Person Data"
// @Failure 400 {object} models.APIResponse "Invalid Person ID"
// @Failure 500 {object} models.APIResponse "Internal Server Error"
// @Router /people/{id} [get]
func GetPeopleById(resp http.ResponseWriter, req *http.Request) {
	// generate and instance of People service *1
	peopleService := services.NewPeople()

	// Extract the ID from the URL
	id := req.URL.Path[len("/people/"):]

	// send error when id is not provided by the client
	if id == "" {
		utils.SendErrorResponse(resp, "ERR: ID is required!", http.StatusBadRequest, models.ErrorInfo{
			Code:    "INVALID_REQUEST",
			Details: "Missing ID parameter in the URL.",
		})
		return
	}

	// Fetch data from the service layer
	data, err := peopleService.GetPeopleById(id)

	// Service failed. Send error response by utils
	if err != nil {
		utils.SendErrorResponse(resp, "ERR: Unable to fetch person data!", http.StatusInternalServerError, models.ErrorInfo{
			Code:    "SERVICE_ERROR",
			Details: err.Error(),
		})
		return
	}

	// Data fetched. Send success response by utils
	utils.SendSuccessResponse(resp, "SCC: Person fetched successfully!", data, http.StatusOK)
}
