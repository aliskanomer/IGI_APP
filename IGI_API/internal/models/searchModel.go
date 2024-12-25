// Description: This file contains the data structures used for search operations. Please check documentation for more details.

package models

// IGIAPI
// Represents the search parameters needed to perform a search operation. Query is built upon this struct.
type SearchParams struct {
	Keyword   string
	Source    []string
	SortBy    string
	SortOrder string
	Page      int
	Limit     int
}

// Represents the search result data structure. This is the base of the cached data structure.
type SearchResults struct {
	People  []SWAPISearchItemPeople `json:"people"`
	Planets []SWAPISearchItemPlanet `json:"planets"`
}

type SearchResponse struct {
	DataSource string `json:"dataSource"`
	HitCount   struct {
		Total           int `json:"total"`
		TotalPagePeople int `json:"totalPagePeople"`
		TotalPagePlanet int `json:"totalPagePlanet"`
	} `json:"hitCount"`
	Results SearchResults `json:"results"`
}
