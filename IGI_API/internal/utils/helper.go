package utils

import (
	// packages
	"IGI_API/internal/models"

	// modules
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

// Check to see if a string is in a slice
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// ConvertToTargetModel converts a generic `interface{}` result into a specific model slice.
func ConvertToTargetModel(source interface{}, target interface{}) error {
	// Encode the source into JSON bytes
	jsonBytes, err := json.Marshal(source)
	if err != nil {
		return fmt.Errorf("ERR: error marshaling source data: %v", err)
	}

	// Decode the JSON bytes into the target model
	if err := json.Unmarshal(jsonBytes, target); err != nil {
		return fmt.Errorf("ERR: error unmarshaling into target model: %v", err)
	}

	return nil
}

// Builds the search query from URL by extracting,parsing and validating the query parameters.
// Returns a SearchParams when success, informs invoker with related error message else.
func SearchQueryBuilder(query url.Values) (models.SearchParams, error) {
	// Defaults
	sortBy := "name"
	sortOrder := "asc"
	page := 1
	limit := 15

	// PARAM: Keyword - required & max100chars
	keyword := strings.ToLower(query.Get("keyword"))
	if len(keyword) == 0 {
		return models.SearchParams{}, errors.New("keyword is required")
	}
	if len(keyword) > 100 {
		return models.SearchParams{}, errors.New("keyword must be at most 100 characters")
	}

	// PARAM: source - required & <people, planet>
	sourceStr := query.Get("source")
	if sourceStr == "" {
		return models.SearchParams{}, errors.New("source is required")
	}
	validSources := []string{"people", "planets"}
	source := strings.Split(sourceStr, ",")
	for _, s := range source {
		if !Contains(validSources, s) {
			return models.SearchParams{}, errors.New("source must be 'people', 'planets', or both")
		}
	}

	// PARAM: page - def1 & min1 --PAGINATION
	if pageStr := query.Get("page"); pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err != nil || parsedPage < 1 {
			return models.SearchParams{}, errors.New("page must be at least 1")
		}
		page = parsedPage
	}
	// PARAM: limit - def15 & min1 --PAGINATION
	if limitStr := query.Get("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit < 1 {
			return models.SearchParams{}, errors.New("limit must be at least 1")
		}
		limit = parsedLimit
	}

	// PARAM: sortBy - defName & <name, created> --SORTING
	if sortByParam := query.Get("sortBy"); sortByParam != "" {
		sortBy = sortByParam
	}
	// PARAM: sortOrder - defAsc & <asc, desc> --SORTING
	if sortOrderParam := query.Get("sortOrder"); sortOrderParam != "" {
		sortOrder = sortOrderParam
	}

	// Query is successfuly built. Return the validated values.
	return models.SearchParams{
		Keyword:   keyword,
		Source:    source,
		SortBy:    sortBy,
		SortOrder: sortOrder,
		Page:      page,
		Limit:     limit,
	}, nil
}

// Sort the search result in their respective arrays based on the sortBy and sortOrder parameters.
// Returns the original arrays sorted. eg. result:{people[SORTED], planets[SORTED]}
func SearchResultSorter(results *models.SearchResults, sortBy, sortOrder string) {
	// Sort People
	sort.Slice(results.People, func(i, j int) bool {
		if sortBy == "created" {
			if sortOrder == "asc" {
				return results.People[i].Properties.Created < results.People[j].Properties.Created
			}
			return results.People[i].Properties.Created > results.People[j].Properties.Created
		}
		// Default to sorting by name
		if sortOrder == "asc" {
			return results.People[i].Properties.Name < results.People[j].Properties.Name
		}
		return results.People[i].Properties.Name > results.People[j].Properties.Name
	})

	// Sort Planets
	sort.Slice(results.Planets, func(i, j int) bool {
		if sortBy == "created" {
			if sortOrder == "asc" {
				return results.Planets[i].Properties.Created < results.Planets[j].Properties.Created
			}
			return results.Planets[i].Properties.Created > results.Planets[j].Properties.Created
		}
		// Default to sorting by name
		if sortOrder == "asc" {
			return results.Planets[i].Properties.Name < results.Planets[j].Properties.Name
		}
		return results.Planets[i].Properties.Name > results.Planets[j].Properties.Name
	})
}

// Paginate the search result in their respective arrays based on the page and limit parameters.
// Returns the original arrays paginated. eg. result:{people[page(i):limit(n)], planets[page(i):limit(n)]}
func SearchResultPaginator(results models.SearchResults, page int, limit int) models.SearchResults {
	start := (page - 1) * limit
	end := start + limit

	// People için pagination
	peopleCount := len(results.People)
	if peopleCount != 0 { // array is not empty*
		if start > peopleCount {
			results.People = []models.SWAPISearchItemPeople{}
		} else {
			if end > peopleCount {
				end = peopleCount
			}
			results.People = results.People[start:end]
		}
	}

	// Planets için pagination
	planetCount := len(results.Planets)
	if planetCount != 0 { // array is empty*
		if start > planetCount && planetCount > 0 {
			results.Planets = []models.SWAPISearchItemPlanet{}
		} else {
			if end > planetCount {
				end = planetCount
			}
			results.Planets = results.Planets[start:end]
		}
	}

	return results
}

// * !prevent index out of range when single source is empty
