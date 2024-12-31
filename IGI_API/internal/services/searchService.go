// Description: This is a concurrent service that runs a search query on multiple resources at once. Please check documentation for more insights.
package services

import (
	// packages
	"IGI_API/internal/models"
	"IGI_API/internal/utils"

	// modules
	"fmt"
	"net/http"
	"sync"
)

// Declare Search service struct
type Search struct {
	BaseURL string // Base URL for SWAPI API
}

// singleton instance of the search service.
var (
	searchInstance *Search
	searchLock     sync.Mutex
)

// Init Search service with base url by singleton pattern.
func NewSearch() *Search {
	if searchInstance == nil {
		searchLock.Lock()
		defer searchLock.Unlock()
		if searchInstance == nil {
			baseURL := utils.ReadEnvVar("SWAPI_BASE_URL")
			searchInstance = &Search{BaseURL: baseURL}
		}
	}
	return searchInstance
}

// Runs concurrent search queries on multiple resources. Concats results on sepereate slices.
func (_searchService *Search) Search(params models.SearchParams) (models.SearchResults, error) {
	utils.Logger("info", "Search", 0, "Initiating concurrent search...")

	var w8grp sync.WaitGroup // wait group to mng concurrent goroutines.
	var mutex sync.Mutex     // mutex to lock results slice while appending.

	results := models.SearchResults{} // results struct to hold the final data.

	errChan := make(chan error, len(params.Source)) // channel for capturing errors. Buffer size (max err capacity) is equal to the number of sources.

	// iterate over the sources...
	for _, source := range params.Source {

		w8grp.Add(1)
		//...concurrently to fetch data from SWAPI
		go func(source string) {
			defer w8grp.Done()
			// construct URL for the search query. eg. /people/?name=luke || /planet/?name=tatooine
			searchURL := fmt.Sprintf("%s%s/?name=%s", _searchService.BaseURL, source, params.Keyword)

			// request data from SWAPI
			utils.Logger("info", "Search", 0, fmt.Sprintf("Fetching data from source: %s", source))
			resp, err := http.Get(searchURL)

			// Connection to SWAPI failed. Log & add Err_channel
			if err != nil {
				utils.Logger("error", "Search", http.StatusInternalServerError, err)
				errChan <- fmt.Errorf("ERR: error fetching data from %s: %v", source, err)
				return
			}
			defer resp.Body.Close()

			// parse response by SWAPISearchResult model
			var swapiSearchResult models.SWAPISearchResult
			if err := utils.ParseJSONResponse(resp, &swapiSearchResult); err != nil {
				utils.Logger("error", "Search", http.StatusInternalServerError, err)
				errChan <- fmt.Errorf("ERR: error parsing response from %s: %v", source, err)
				return
			}

			// SWAPI returned non-success status code. Log & add Err_channel
			if resp.StatusCode != http.StatusOK {
				utils.Logger("error", "Search", resp.StatusCode, fmt.Errorf("unexpected status code: %d", resp.StatusCode))
				errChan <- fmt.Errorf("ERR: Unexpected status code %d from %s", resp.StatusCode, source)
				return
			}

			if swapiSearchResult.Result == nil {
				utils.Logger("info", "Search", 0, fmt.Sprintf("No results found from source: %s", source))
				return
			}

			// Start processing the results based on the source type.
			mutex.Lock()
			defer mutex.Unlock()

			// Query builder ensures the source data safety.
			switch source {
			case "people":
				// Map the result to the People model by utils.
				peopleData := []models.SWAPISearchItemPeople{}
				if err := utils.ConvertToTargetModel(swapiSearchResult.Result, &peopleData); err != nil {
					// Data mapping failed. Log & add Err_channel
					_errMsg := fmt.Errorf("ERR: Failed mapping data from source: %s, error: %w", source, err)
					utils.Logger("error", "Search", http.StatusInternalServerError, _errMsg)
					errChan <- fmt.Errorf("ERR: error mapping people data: %v", err)
					return
				}

				// Append the mapped data to the results arr
				results.People = append(results.People, peopleData...)

			case "planets":
				// Map the result to the Planet model by utils.
				planetData := []models.SWAPISearchItemPlanet{}
				if err := utils.ConvertToTargetModel(swapiSearchResult.Result, &planetData); err != nil {
					// Data mapping failed. Log & add Err_channel
					_errMsg := fmt.Errorf("ERR: Failed mapping data from source: %s, error: %w", source, err)
					utils.Logger("error", "Search", http.StatusInternalServerError, _errMsg)
					errChan <- fmt.Errorf("error mapping planet data: %v", err)
					return
				}

				// Append the mapped data to the results arr
				results.Planets = append(results.Planets, planetData...)
			}
		}(source)

	}

	// Wait for all goroutines to finish.
	w8grp.Wait()

	// Check for errors after closing the channel
	close(errChan)

	for err := range errChan {
		utils.Logger("error", "Search", http.StatusInternalServerError, err)
		return results, fmt.Errorf("ERR: One or more errors occurred during search: %w", err)
	}

	utils.Logger("success", "Search", http.StatusOK, "Search operation completed successfuly!")
	return results, nil
}

// KISS EXPLANATION:
// Service always returns raw data.
// Controller sort, paginates and caches the raw material. It is controllers reponsibility to decide if data should be read from and/or written to cache.
