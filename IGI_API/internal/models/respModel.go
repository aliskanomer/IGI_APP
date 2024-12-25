package models

// APIResponse is the base response structure for all IGI_API responses.
type APIResponse struct {
	Message  string      `json:"message"`
	Response interface{} `json:"response"` // this can be any type of data. Data is presented as a property of response
	Status   int         `json:"status"`
}

