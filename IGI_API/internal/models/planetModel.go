package models

// Struct for common planet properties
type PlanetProperties struct {
	Diameter       string `json:"diameter"`
	RotationPeriod string `json:"rotation_period"`
	OrbitalPeriod  string `json:"orbital_period"`
	Gravity        string `json:"gravity"`
	Population     string `json:"population"`
	Climate        string `json:"climate"`
	Terrain        string `json:"terrain"`
	SurfaceWater   string `json:"surface_water"`
	Created        string `json:"created"`
	Edited         string `json:"edited"`
	Name           string `json:"name"`
	URL            string `json:"url"`
}

// Struct for individual results eg. /planets/:id
type PlanetByIDResponse struct {
	Message string `json:"message"`
	Result  struct {
		Properties  PlanetProperties `json:"properties"`
		Description string           `json:"description"`
		UID         string           `json:"uid"`
	} `json:"result"`
}
