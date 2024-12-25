package models

// Struct for common people properties
type PersonProperties struct {
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
	Created   string `json:"created"`
	Edited    string `json:"edited"`
	Name      string `json:"name"`
	Homeworld string `json:"homeworld"`
	URL       string `json:"url"`
}

// Struct for individual results eg. /people/:id
type PeopleByIDResponse struct {
	Message string `json:"message"`
	Result  struct {
		Properties  PersonProperties `json:"properties"`
		Description string           `json:"description"`
		UID         string           `json:"uid"`
	} `json:"result"`
}
