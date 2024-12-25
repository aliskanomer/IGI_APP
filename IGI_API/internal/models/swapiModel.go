package models

// Represents a list item on SWAPI resource list. eg. /people/ response model.
type SWAPIListItem struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

//Represents a response on SWAPI resource list. eg. /people response model.
type SWAPIList struct {
	Message      string          `json:"message"`
	TotalPages   int             `json:"total_pages"`
	TotalRecords int             `json:"total_records"`
	Previous     interface{}     `json:"previous"`
	Next         string          `json:"next"`
	Results      []SWAPIListItem `json:"results"`
}

// Represents the search result struructure of SWAPI. /resourece/?name=keyword response model.
type SWAPISearchResult struct {
	Message string      `json:"message"` // API message field.
	Result  interface{} `json:"result"`  // Generic result field.
}

// Represents the fixed properties of people search from SWAPI. eg. /people/?name=e
type SWAPISearchItem struct {
	Properties  PersonProperties `json:"properties"`
	Description string           `json:"description"`
	ID          string           `json:"_id"`
	UID         string           `json:"uid"`
	Version     int              `json:"__v"`
}

type SWAPISearchItemPeople struct {
	Properties  PersonProperties `json:"properties"`
	Description string           `json:"description"`
	UID         string           `json:"uid"`
}
type SWAPISearchItemPlanet struct {
	Properties  PlanetProperties `json:"properties"`
	Description string           `json:"description"`
	UID         string           `json:"uid"`
}
