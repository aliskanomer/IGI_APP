package controllers

import (
	// packages
	"IGI_API/internal/cache"
	"IGI_API/internal/models"
	"IGI_API/internal/services"
	"IGI_API/internal/utils"

	// modules
	"net/http"
	"time"
)

// @Summary Search resources by keyword
// @Description Search for people or planets using a keyword from SWAPI.
// @Tags Search
// @Accept  json
// @Produce  json
// @Param   keyword query string true "Search Keyword"
// @Param   source query string true "Data Source (people or planets)"
// @Param   page query int false "Page number (default: 1)"
// @Param   limit query int false "Items per page (default: 15)"
// @Param   sortBy query string false "Sort field (name or height | default: name)"
// @Param   sortOrder query string false "Sort order (asc or desc | default: asc)"
// @Success 200 {object} models.APIResponse{response=models.SearchResponse} "Successful search operation"
// @Failure 400 {object} models.APIResponse "Invalid Query Parameters"
// @Failure 500 {object} models.APIResponse "Internal Server Error"
// @Router /search [get]
func Search(resp http.ResponseWriter, req *http.Request) {
	// Build and validate the received query parameters
	params, err := utils.SearchQueryBuilder(req.URL.Query())

	// Invalid query parameters. Inform invoker with error message.
	if err != nil {
		utils.SendErrorResponse(resp, "ERR: Invalid query parameters!", http.StatusBadRequest, models.ErrorInfo{
			Code:    "INVALID_QUERY_PARAMS",
			Details: err.Error(),
		})
		return
	}

	// Results stored in cache or fetched by service is not paginated or sorted
	// and will be stored in here till the source of data is decided
	var rawResults models.SearchResults
	var dataSource string

	// Generate a key from query word and source to check if the data is already cached.
	_cache := cache.NewCache()
	_cacheKey := _cache.KeyGen(params.Keyword, params.Source)

	// Try to read data from the cache
	cacheItem, found := _cache.Get(_cacheKey)

	// Data source decider. If cache contains data and it is not expired, use it. Else, fetch data from service.
	if found && time.Now().Before(cacheItem.ExpiresAt) {
		// Cache contains data and it is not expired (15min_exp_time)
		dataSource = "cache"
		rawResults = cacheItem.Data
	} else {
		// Meaning cache is empty or expired
		dataSource = "service"

		// Fetch data by service layer
		searchService := services.NewSearch()
		rawResults, err = searchService.Search(params)

		if err != nil {
			utils.SendErrorResponse(resp, "ERR: Unable to fetch search data from service!", http.StatusInternalServerError, models.ErrorInfo{
				Code:    "SERVICE_ERROR",
				Details: err.Error(),
			})
			return
		}
		// Service succeeded. Store the results in the cache.
		_cache.Set(_cacheKey, rawResults, 15*time.Minute)
	}

	// Sort whole data set before pagination
	utils.SearchResultSorter(&rawResults, params.SortBy, params.SortOrder)

	// Paginate the data
	paginatedResults := utils.SearchResultPaginator(rawResults, params.Page, params.Limit)

	totalHits := len(rawResults.People) + len(rawResults.Planets)
	people_pages := (len(rawResults.People) + params.Limit - 1) / params.Limit  // min 1 page when there are less hit then limit
	planet_pages := (len(rawResults.Planets) + params.Limit - 1) / params.Limit // min 1 page when there are less hit then limit

	// Create a response payload model. Client might need the totalHit count or dataSource data for UI operations.
	responsePayload := map[string]interface{}{
		"dataSource": dataSource,
		"hitCount": map[string]int{
			"total":           totalHits,
			"totalPagePeople": people_pages,
			"totalPagePlanet": planet_pages,
		},
		"results": paginatedResults,
	}
	utils.SendSuccessResponse(resp, "SCC: Operation success!", responsePayload, http.StatusOK)
}
