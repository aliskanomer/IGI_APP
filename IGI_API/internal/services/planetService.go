package services

import (
	// packages
	"IGI_API/internal/models"
	"IGI_API/internal/utils"
	"fmt"

	// modules
	"net/http"
	"sync"
)

// Declare Planet struct
type Planet struct {
	BaseURL string
}

// singleton instance of the planet service.
var (
	planetInstance *Planet
	planetLock     sync.Mutex
)

// Init Planet service with base url by singleton pattern.
func NewPlanet() *Planet {
	if planetInstance == nil {
		planetLock.Lock()
		defer planetLock.Unlock()
		if planetInstance == nil {
			baseURL := utils.ReadEnvVar("SWAPI_BASE_URL") + "planets/"
			planetInstance = &Planet{BaseURL: baseURL}
		}
	}
	return planetInstance
}

// GetPlanetAll to get all people data from SWAPI
func (_planet *Planet) GetPlanetAll(page int, limit int) (*models.SWAPIList, error) {
	utils.Logger("info", "GetPlanetAll", 0, "Fetching planet data...")

	// Build the request URL with pagination parameters
	url := fmt.Sprintf("%s?page=%d&limit=%d", _planet.BaseURL, page, limit)

	// request data from SWAPI
	resp, err := http.Get(url)

	// Connection to SWAPI failed. Log & return
	if err != nil {
		utils.Logger("error", "GetPlanetAll", http.StatusInternalServerError, err)
		return nil, fmt.Errorf("ERR: Failed connecting to SWAPI: %w", err)
	}

	// SWAPI returned non-success status code. Log & return
	if resp.StatusCode != http.StatusOK {
		utils.Logger("error", "GetPlanetAll", resp.StatusCode, fmt.Errorf("unexpected status code %d", resp.StatusCode))
		return nil, fmt.Errorf("ERR: unexpected status code %d", resp.StatusCode)
	}

	// parse response by SWAPIList model
	var response models.SWAPIList

	// Could not parse the response. Log & return
	if err := utils.ParseJSONResponse(resp, &response); err != nil {
		utils.Logger("error", "GetPlanetAll", http.StatusInternalServerError, err)
		return nil, fmt.Errorf("ERR: failed to parse planet data: %w", err)
	}

	// Operation successful. Log & return
	utils.Logger("success", "GetPlanetAll", http.StatusOK, "Planet data fetched successfuly!")
	return &response, nil
}

// GetPlanetByID to get singular planet data from SWAPI - requires ID
func (_planet *Planet) GetPlanetByID(id string) (*models.PlanetByIDResponse, error) {
	utils.Logger("info", "GetPlanetByID", 0, fmt.Sprintf("Fetching planet data with ID: %s...", id))

	// build url with ID
	url := _planet.BaseURL + id

	// request data from SWAPI
	resp, err := http.Get(url)

	// Connection to SWAPI failed. Log & return
	if err != nil {
		utils.Logger("error", "GetPlanetByID", http.StatusInternalServerError, err)
		return nil, fmt.Errorf("ERR: Failed connecting to SWAPI: %w", err)
	}

	// SWAPI returned non-success status code. Log & return
	if resp.StatusCode != http.StatusOK {
		utils.Logger("error", "GetPlanetByID", resp.StatusCode, fmt.Errorf("unexpected status code %d", resp.StatusCode))
		return nil, fmt.Errorf("ERR: unexpected status code %d", resp.StatusCode)
	}

	// parse response by PlanetByIDResponse model
	var response models.PlanetByIDResponse

	// Could not parse the response. Log & return
	if err := utils.ParseJSONResponse(resp, &response); err != nil {
		utils.Logger("error", "GetPlanetByID", http.StatusInternalServerError, err)
		return nil, fmt.Errorf("ERR: failed to parse planet data: %w", err)
	}

	// Operation successful. Log & return
	utils.Logger("success", "GetPlanetByID", http.StatusOK, fmt.Sprintf("Planet data with ID: %s fetched successfuly!", id))
	return &response, nil
}
