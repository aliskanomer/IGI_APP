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

// Declare People struct
type People struct {
	BaseURL string
}

// singleton instance of the people service.
var (
	peopleInstance *People
	peopleLock     sync.Mutex
)

// Init People service with base url by singleton pattern.
func NewPeople() *People {
	if peopleInstance == nil {
		peopleLock.Lock()
		defer peopleLock.Unlock()
		if peopleInstance == nil {
			baseURL := utils.ReadEnvVar("SWAPI_BASE_URL") + "people/"
			peopleInstance = &People{BaseURL: baseURL}
		}
	}
	return peopleInstance
}

func (_people *People) GetPeopleAll(page int, limit int) (*models.SWAPIList, error) {
	utils.Logger("info", "GetPeopleAll", 0, "Fetching people data...")

	// build url with pagination
	url := fmt.Sprintf("%s?page=%d&limit=%d", _people.BaseURL, page, limit)

	// request data from SWAPI
	resp, err := http.Get(url)

	// connection to SWAPI failed. log & return
	if err != nil {
		utils.Logger("error", "GetPeopleAll", http.StatusInternalServerError, err)
		return nil, fmt.Errorf("ERR: Failed connecting the SWAPI: %w", err)
	}

	// SWAPI returned non-success status code. Log & return
	if resp.StatusCode != http.StatusOK {
		utils.Logger("error", "GetPeopleAll", resp.StatusCode, err)
		return nil, fmt.Errorf("ERR: unexpected status code %d", resp.StatusCode)
	}

	// parse response by SWAPIList model
	var response models.SWAPIList

	// Could not parse the response. log & return
	if err := utils.ParseJSONResponse(resp, &response); err != nil {
		utils.Logger("error", "GetPeopleAll", http.StatusInternalServerError, err)
		return nil, fmt.Errorf("ERR: failed to parse people data: %w", err)
	}

	// Operation successful. log & return
	utils.Logger("success", "GetPeopleAll", http.StatusOK, "People data fetched successfuly!")
	return &response, nil
}

// GetPeopleById to get singular person data from SWAPI - requires ID
func (_people *People) GetPeopleById(id string) (*models.PeopleByIDResponse, error) {
	utils.Logger("info", "GetPeopleById", 0, fmt.Sprintf("Fetching person data with ID: %s...", id))

	// build route from base
	url := _people.BaseURL + id

	// request data from SWAPI
	resp, err := http.Get(url)

	// Connection to SWAPI failed. Log & return
	if err != nil {
		utils.Logger("error", "GetPeopleById", http.StatusInternalServerError, err)
		return nil, fmt.Errorf("ERR: Failed connecting to SWAPI: %w", err)
	}

	// SWAPI returned non-success status code. Log & return
	if resp.StatusCode != http.StatusOK {
		utils.Logger("error", "GetPeopleById", resp.StatusCode, fmt.Errorf("unexpected status code %d", resp.StatusCode))
		return nil, fmt.Errorf("ERR: unexpected status code %d", resp.StatusCode)
	}

	// parse response by PeopleByIDResponse model
	var response models.PeopleByIDResponse

	// Could not parse the response. Log & return
	if err := utils.ParseJSONResponse(resp, &response); err != nil {
		utils.Logger("error", "GetPeopleById", http.StatusInternalServerError, err)
		return nil, fmt.Errorf("ERR: failed to parse person data: %w", err)
	}

	// Operation successful. Log & return
	utils.Logger("success", "GetPeopleById", http.StatusOK, fmt.Sprintf("Person data with ID: %s fetched successfuly!", id))
	return &response, nil
}
